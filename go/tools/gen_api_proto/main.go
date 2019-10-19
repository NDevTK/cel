// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Command gen_api_proto takes as input an API Discovery Document and produces
// ProtoBuf messages for the types defined in the API's schema.
//
// This command is invoked via the `generate.go` file at the top of the source
// repository, and shouldn't need to be invoked manually.
//
// For documentation, run:
//     ./gen_api_proto -h
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"text/template"
	"unicode"
)

type RestDescription struct {
	Kind             string `json:"kind"`
	DiscoveryVersion string `json:"discoveryVersion"`
	Id               string `json:"id"`

	Name              string `json:"name"`
	Version           string `json:"version"`
	Revision          string `json:"revision"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	DocumentationLink string `json:"documentationLink"`
	BaseUrl           string `json:"baseUrl"`

	Schemas map[string]*Property `json:"schemas"`

	GoPackage string `json:"-"`
}

// Property represents both a Property in a Google Discovery Document and also
// a JsonSchema object as described in
// https://tools.ietf.org/html/draft-zyp-json-schema-03.
type Property struct {
	Id                   string               `json:"id"`
	Type                 string               `json:"type"`
	Ref                  string               `json:"$ref"`
	Description          string               `json:"description"`
	Default              string               `json:"default"`
	Required             bool                 `json:"required"`
	Format               string               `json:"format"`
	Pattern              string               `json:"pattern"`
	Enum                 []string             `json:"enum"`
	EnumDescriptions     []string             `json:"enumDescriptions"`
	Properties           map[string]*Property `json:"properties"`
	AdditionalProperties *Property            `json:"additionalProperties"`
	Items                *Property            `json:"items"`
}

type Arguments struct {
	DiscoveryJsonFile   string
	OutputProtoFile     string
	OutputValidatorFile string
	GoPackage           string
}

const kIndentSpaces = 2

func writeln(w io.Writer, indent int, format string, args ...interface{}) error {
	_, err := w.Write([]byte(strings.Repeat(" ", indent*kIndentSpaces)))
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, format, args...)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\n"))
	return err
}

func wrapLines(lines []string) (out []string) {
	const kColumns = 77

	for _, line := range lines {
		line = strings.TrimRightFunc(line, unicode.IsSpace)
		for {
			if len(line) < kColumns {
				out = append(out, line)
				break
			}

			index := strings.LastIndexFunc(line[:kColumns], unicode.IsSpace)
			if index == -1 {
				index = strings.IndexFunc(line, unicode.IsSpace)
				if index == -1 {
					out = append(out, line)
					break
				}
			}

			out = append(out, line[:index])
			line = line[index+1:]
		}
	}

	return
}

func emitDescription(level int, description string, w io.Writer) error {
	if description == "" {
		return nil
	}
	lines := wrapLines(strings.Split(description, "\n"))
	for _, line := range lines {
		writeln(w, level, "// %s", line)
	}
	return nil
}

func emitField(indent int, fieldName string, fieldIdx int, p *Property, w io.Writer) error {
	emitDescription(indent, p.Description, w)
	if len(p.Enum) > 0 {
		writeln(w, indent, "// Valid values:")
		for _, e := range p.Enum {
			writeln(w, indent, "//     %s", e)
		}
	}

	typeName, err := emitMessage(indent, strings.Title(fieldName), p, w)
	if err != nil {
		return err
	}

	return writeln(w, indent, "%s %s = %d;", typeName, fieldName, fieldIdx)
}

// emitMessage emits a ProtoBuf message definition for the type
// described by a JsonSchema. If the type is a reference or a well known type,
// then no message definition is emitted. In either case, the function returns
// the ProtoBuf type name corresponding to the JsonSchema.
func emitMessage(level int, parentId string, p *Property, w io.Writer) (string, error) {
	if p.Ref != "" {
		return p.Ref, nil
	}

	// No message is necessary for well known or basic types.
	switch p.Type {
	case "string":
		return "string", nil

	case "boolean":
		return "bool", nil

	case "integer":
		return "int32", nil

	case "number":
		return "double", nil

	case "any":
		return "google.protobuf.Any", nil

	case "array":
		if p.Items == nil {
			return "", fmt.Errorf("array does not have an inner type")
		}

		innerType, err := emitMessage(level, parentId, p.Items, w)
		if err != nil {
			return "", err
		}
		return "repeated " + innerType, nil

	case "object":
		break

	default:
		return "", fmt.Errorf("unsupported type \"%s\" as a top level field type or map value type", p.Type)
	}

	name := p.Id
	if name == "" {
		name = parentId
	}
	if name == "" {
		return "", fmt.Errorf("JsonSchema has no \"id\" and has no parent schema name")
	}

	// A pure map from string to whichever type is described in AdditionalProperties.
	if len(p.Properties) == 0 && p.AdditionalProperties != nil {
		mapped_type, _ := emitMessage(level, name, p.AdditionalProperties, w)
		return fmt.Sprintf("map<string, %s>", mapped_type), nil
	}

	emitDescription(level, p.Description, w)
	writeln(w, level, "message %s {", name)

	index := 1
	isFirst := true

	properties := make([]string, 0, len(p.Properties))
	for name, _ := range p.Properties {
		properties = append(properties, name)
	}
	sort.Strings(properties)

	for _, name := range properties {
		pp := p.Properties[name]
		if isFirst {
			isFirst = false
		} else {
			w.Write([]byte("\n"))
		}
		err := emitField(level+1, name, index, pp, w)
		if err != nil {
			return "", err
		}
		index += 1
	}
	writeln(w, level, "}")
	return name, nil
}

func GenerateProtoFile(desc *RestDescription, args *Arguments) error {
	desc.GoPackage = args.GoPackage + "/" + desc.Name

	f, err := os.Create(args.OutputProtoFile)
	if err != nil {
		return err
	}
	defer f.Close()

	const kTemplate = `// This is a generated file. Do not modify directly.

// This ProtoBuf source file is based on the REST protocol defintion for the
// service at {{.BaseUrl}}.
//
// {{.Title}}: {{.Description}}
// 
// API Name      : {{.Name}} ({{.Id}})
// Version       : {{.Version}}
// Revision      : {{.Revision}}
// Documentation : {{.DocumentationLink}}

syntax="proto3";

// Generated protobuf for {{.Name}}
//
// --- Skip validation ---
package {{.Name}};
option go_package="{{.GoPackage}}";

import "google/protobuf/any.proto";


`

	err = template.Must(template.New("heading").Parse(kTemplate)).Execute(f, desc)
	if err != nil {
		return err
	}

	schemas := make([]string, 0, len(desc.Schemas))
	for s, _ := range desc.Schemas {
		schemas = append(schemas, s)
	}
	sort.Strings(schemas)

	for _, s := range schemas {
		_, err = emitMessage(0, "", desc.Schemas[s], f)
		if err != nil {
			return err
		}
		f.WriteString("\n")
	}
	return nil
}

func GenerateValidatorFile(desc *RestDescription, args *Arguments) error {
	// No need to generate a validator file.
	if args.OutputValidatorFile == "" {
		return nil
	}

	f, err := os.Create(args.OutputValidatorFile)
	if err != nil {
		return err
	}
	defer f.Close()

	var messages []string
	for _, s := range desc.Schemas {
		messages = append(messages, s.Id)
		for k, p := range s.Properties {
			if (p.Type == "object" && len(p.Properties) > 0) ||
				(p.Type == "array" && p.Items.Type == "object" && len(p.Items.Properties) > 0) {
				messages = append(messages, s.Id+"_"+strings.Title(k))
			}
		}
	}

	sort.Strings(messages)

	const kHeading = `// This is a generated file. Do not modify directly.

// This ProtoBuf source file is based on the REST protocol defintion for the
// service at {{.BaseUrl}}.
//
// {{.Title}}: {{.Description}}
// 
// API Name      : {{.Name}} ({{.Id}})
// Version       : {{.Version}}
// Revision      : {{.Revision}}
// Documentation : {{.DocumentationLink}}

package {{.Name}}

// All trivial validators.

`
	err = template.Must(template.New("").Parse(kHeading)).Execute(f, desc)
	if err != nil {
		return err
	}

	const kValidators = `{{range . -}}
func (*{{.}}) Validate() error { return nil }
{{end}}`
	return template.Must(template.New("").Parse(kValidators)).Execute(f, messages)
}

func DoIt(args *Arguments) error {
	if args.DiscoveryJsonFile == "" {
		return fmt.Errorf("no input filename specified")
	}
	if args.OutputProtoFile == "" {
		return fmt.Errorf("no output filename specified")
	}

	contents, err := ioutil.ReadFile(args.DiscoveryJsonFile)
	if err != nil {
		return err
	}

	var desc RestDescription
	err = json.Unmarshal(contents, &desc)
	if err != nil {
		return err
	}

	err = GenerateProtoFile(&desc, args)
	if err != nil {
		return err
	}

	return GenerateValidatorFile(&desc, args)
}

func PrintUsage() {
	const kUsage = `Usage of gen_api_proto:

This tool can be used to generate ProtoBuf message definitions based on Google
Cloud Platform API Discovery Document [1]. The tool understands a minimal and
possibly incorrect subset of the description schema and produces ProtoBuf
definitions that are compatible with Google API Client Libraries [2].

API Discovery Documents for each API can be found at the Google API Client Go
library respository [2]. Every API and version specific directory contains a
*-api.json file.

E.g.: The discovery document for the compute/v1 API is at
https://github.com/google/google-api-go-client/tree/master/compute/v1

The only reason for the existence of this tool is to generate ProtoBuf message
syntaxes for some of the Google Compute Engine resources that don't yet have
corresponding .proto definitions in Google Cloud Client Libraries [3]. Once
GCCL fully supports Google Compute Engine, this tool can be retired.

[1]: https://developers.google.com/discovery/v1/reference/apis
[2]: https://developers.google.com/api-client-library/
[3]: https://cloud.google.com/apis/docs/cloud-client-libraries

Synopsis: gen_api_proto -i <name of API Discovery Document> -o <output filename> -g <go validator filename>
`

	os.Stderr.WriteString(kUsage)
}

func main() {
	var args Arguments

	flagset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flagset.StringVar(&args.DiscoveryJsonFile, "i", "", "discovery JSON filename")
	flagset.StringVar(&args.OutputProtoFile, "o", "", "output .proto filename")
	flagset.StringVar(&args.OutputValidatorFile, "g", "", "ouput .go validator filename")
	flagset.StringVar(&args.GoPackage, "p", "chromium.googlesource.com/enterprise/cel/go/gcp",
		"root of go_package option to emit")
	flagset.Usage = func() {
		PrintUsage()
		flagset.PrintDefaults()
	}

	flagset.Parse(os.Args[1:])

	if flagset.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Unrecognized arguments: %#v\n\n", flagset.Args())
		flagset.Usage()
		return
	}

	err := DoIt(&args)
	if err != nil {
		fmt.Print("Error:", err, "\n")
	}
}
