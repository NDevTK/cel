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

	BaseUrl           string `json:"baseUrl"`
	Description       string `json:"description"`
	DocumentationLink string `json:"documentationLink"`
	Name              string `json:"name"`
	Revision          string `json:"revision"`
	Title             string `json:"title"`
	Version           string `json:"version"`

	Schemas map[string]*Property `json:"schemas"`

	GoPackage string `json:"-"`
}

type Property struct {
	Id                   string               `json:"id"`
	Type                 string               `json:"type"`
	Description          string               `json:"description"`
	Format               string               `json:"format"`
	Default              string               `json:"default"`
	Enum                 []string             `json:"enum"`
	Properties           map[string]*Property `json:"properties"`
	AdditionalProperties *Property            `json:"additionalProperties"`
	Items                *Property            `json:"items"`
	Ref                  string               `json:"$ref"`
}

type Arguments struct {
	DiscoveryJsonFile string
	OutputProtoFile   string
	GoPackage         string
}

const kIndent = 2

func Indent(level int, w io.Writer) error {
	_, err := w.Write([]byte(strings.Repeat(" ", level*kIndent)))
	return err
}

func IndentedLine(w io.Writer, level int, format string, args ...interface{}) error {
	err := Indent(level, w)
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

func WrapLines(lines []string) (out []string) {
	const kColumns = 70

	for _, line := range lines {
		line = strings.TrimRightFunc(line, unicode.IsSpace)
		for {
			if len(line) < kColumns {
				out = append(out, line)
				break
			}

			index := strings.LastIndexFunc(line[:kColumns], unicode.IsSpace)
			if index == -1 {
				index := strings.IndexFunc(line, unicode.IsSpace)
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

func EmitDescription(level int, description string, w io.Writer) error {
	if description == "" {
		return nil
	}
	lines := WrapLines(strings.Split(description, "\n"))
	for _, line := range lines {
		IndentedLine(w, level, "// %s", line)
	}
	return nil
}

type FieldType int

const (
	SINGULAR FieldType = iota
	REPEATED
	MAP
)

func TypeWithDecoration(name string, field_type FieldType) string {
	switch field_type {
	case SINGULAR:
		return name

	case REPEATED:
		return "repeated " + name

	case MAP:
		return "map<string, " + name + ">"
	}
	return name
}

func EmitField(level int, name string, index int, p *Property, field_type FieldType, w io.Writer) error {
	EmitDescription(level, p.Description, w)
	switch {
	case p.Ref != "":
		return IndentedLine(w, level, "%s %s = %d;", TypeWithDecoration(p.Ref, field_type), name, index)

	case p.Type == "object":
		type_name := strings.Title(name)
		type_name, _ = EmitProtoForProperty(level, type_name, p, w)
		return IndentedLine(w, level, "%s %s = %d;", TypeWithDecoration(type_name, field_type), name, index)

	case p.Type == "string":
		if len(p.Enum) > 0 {
			IndentedLine(w, level, "// Valid values:")
			for _, e := range p.Enum {
				IndentedLine(w, level, "//     %s", e)
			}
		}
		return IndentedLine(w, level, "%s %s = %d;", TypeWithDecoration("string", field_type), name, index)

	case p.Type == "boolean":
		return IndentedLine(w, level, "%s %s = %d;", TypeWithDecoration("bool", field_type), name, index)

	case p.Type == "integer" || p.Type == "number":
		int_type := "int32"
		if p.Format != "" {
			int_type = p.Format
		}
		return IndentedLine(w, level, "%s %s = %d;", TypeWithDecoration(int_type, field_type), name, index)

	case p.Type == "array":
		if field_type != SINGULAR {
			return fmt.Errorf("repeated repeated? %s: %#v", name, p)
		}

		if p.Items == nil {
			return fmt.Errorf("Array does not have a type for field %s: %#v", name, p)
		}

		return EmitField(level, name, index, p.Items, REPEATED, w)
	}
	return fmt.Errorf("Unknown type for field %s: %#v", name, p)
}

func EmitProtoForProperty(level int, name string, p *Property, w io.Writer) (string, error) {
	if p.Id == "" && p.Ref == "" && name == "" {
		return "", fmt.Errorf("Required property \"Id\" missing")
	}

	if p.Ref != "" {
		return p.Ref, nil
	}

	if p.Type != "object" {
		return p.Type, nil
	}

	if p.Id != "" {
		name = p.Id
	}

	if len(p.Properties) == 0 && p.AdditionalProperties != nil {
		mapped_type, _ := EmitProtoForProperty(level, name, p.AdditionalProperties, w)
		return TypeWithDecoration(mapped_type, MAP), nil
	}

	EmitDescription(level, p.Description, w)
	IndentedLine(w, level, "message %s {", name)

	index := 1
	is_first := true

	properties := make([]string, 0, len(p.Properties))
	for name, _ := range p.Properties {
		properties = append(properties, name)
	}
	sort.Strings(properties)

	for _, name := range properties {
		pp := p.Properties[name]
		if is_first {
			is_first = false
		} else {
			w.Write([]byte("\n"))
		}
		err := EmitField(level+1, name, index, pp, SINGULAR, w)
		if err != nil {
			return "", err
		}
		index += 1
	}
	IndentedLine(w, level, "}")
	return name, nil
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
// {{.Title}}:{{.Description}}
// 
// API Name      : {{.Name}} ({{.Id}})
// Version       : {{.Version}}
// Revision      : {{.Revision}}
// Documentation : {{.DocumentationLink}}

syntax="proto3";
package {{.Name}};
option go_package="{{.GoPackage}}";


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
		_, err = EmitProtoForProperty(0, "", desc.Schemas[s], f)
		if err != nil {
			return err
		}
		f.WriteString("\n")
	}
	return nil
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

Synopsis: gen_api_proto -i <name of API Discovery Document> -o <output filename>
`

	os.Stderr.WriteString(kUsage)
}

func main() {
	var args Arguments

	flagset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flagset.StringVar(&args.DiscoveryJsonFile, "i", "", "discovery JSON filename")
	flagset.StringVar(&args.OutputProtoFile, "o", "", "output .proto filename")
	flagset.StringVar(&args.GoPackage, "g", "chromium.googlesource.com/enterprise/cel/go/gcp",
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
