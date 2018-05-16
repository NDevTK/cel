// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"io"
	"io/ioutil"
	"sort"
	"strings"
	"text/template"

	"chromium.googlesource.com/enterprise/cel/go/common"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// text/template compatible template for a Markdown document. Each package is
// described in a single .md file.
//
// The fields in this template are based on a pipeline containing a PackageData
// object.
//
// The flavor of Markdown we use is Gitiles. See
// https://gerrit.googlesource.com/gitiles/+/master/Documentation/markdown.md
const DocTemplate = `
{{- /* Field declaration. */}}
{{- define "field" -}}

* {{ if .IsRepeated -}}
    {{"repeated" | code}} {{end -}}

  {{ if .TypeLink -}}
	 [{{code .TypeName}}]({{.TypeLink}})
  {{- else -}}
	 {{code .TypeName}}
  {{- end }}{{" " -}}

  [{{code .Name}}](#{{.Anchor}}) = {{.Number}}

  {{- if .IsRequired }} (**Required**){{end -}}

{{end -}}

# Schema {{code .Package}} {#{{.Anchor}}}

{{/* Package documentation */ -}}
{{.Doc}}

Messages that are valid in package {{ code .Package }} are as follows:

*** note
Note that this document uses the term "message" to refer to the same concept as
a "message" in Protocol Buffers. Hence every asset and host resource
description is a *message*. So is their embedded structures.
***

{{ range .Messages -}}

## Message {{code .Name}} {#{{.Anchor}}}

{{.Doc}}

{{ if .HasInputs -}}

### Inputs for {{code .Name}}

{{ range .Fields -}}
  {{ if .IsInput -}}
    {{ template "field" . }}
{{ end -}}
{{ end }}
{{ end }}{{/* .HasInputs */ -}}

{{ if .HasOutputs -}}

### Outputs for {{code .Name}}

{{ range .Fields -}}
  {{ if .IsOutput -}}
	{{ template "field" . }}
{{ end -}}
{{ end }}
{{ end }}{{/* .HasOutputs */ -}}

{{ if .HasRuntime -}}

## Runtime fields for {{code .Name}}

{{ range .Fields -}}
  {{ if .IsRuntime -}}
	{{ template "field" . }}
{{ end -}}
{{ end }}
{{ end }}{{/* .HasRuntime */ -}}

{{ range .Fields -}}
### {{code .Name}} {#{{.Anchor}}}

| Property | Comments |
|----------|----------|
| Field Name | {{code .Name}} |
{{if .TypeLink -}}
  | Type | [{{code .TypeName}}]({{.TypeLink}}) |
{{else -}}
  | Type | {{code .TypeName}} |
{{end -}}

{{if .IsRepeated -}}
  | Repeated | Any number of instances of this type is allowed in the schema. |
{{end -}}

{{if .IsRequired -}}
  | Required | This field is required. It is an error to omit this field. |
{{end -}}

{{if .IsReference -}}
  | Reference | The value of this field is a named reference to a {{code .Reference}} |
{{end -}}
{{if .Doc}}
{{.Doc}}
{{- end }}
{{ end }}
{{- end }}

# Enumerations

{{range .Enums -}}
## Enumeration {{code .Name}} {#{{.Anchor}}}

{{.Doc}}

Values:
{{ range .Values -}}
  {{ if .Doc -}}
	* ["{{code .Name}}"](#{{.Anchor}})
  {{- else -}}
	* "{{code .Name}}"
  {{- end }}
{{ end -}}

{{ range .Values }}{{ if .Doc }}
### "{{code .Name}}" {#{{.Anchor}}}

{{ .Doc }}
{{ end }}
{{- end }}
{{ end }}
---
Generated from {{range $i, $e := .SourceFiles -}}
{{if $i}}, {{end}}{{code $e}}
{{- end}}.
`

// PackageData contains the documentation for a single package file.
type PackageData struct {
	Package     string              // Package name. All symbols in this file are contained in this package.
	Anchor      string              // Document local named anchor.
	Doc         string              // Package documentation.
	SourceFiles []string            // Relative path to the source .proto files.
	Messages    []MessageData       // List of messages.
	Enums       map[string]EnumData // All enumerations contained within this package.
}

// MessageData contains the documentation and link information for a single
// ProtoBuf message.
type MessageData struct {
	Name       string      // Name of message.
	Anchor     string      // Document local named anchor that identifies this message.
	Doc        string      // Toplevel documentation for message.
	HasInputs  bool        // True if at least one of Fields is an input.
	HasOutputs bool        // True if at least one of Fields is an output.
	HasRuntime bool        // True if at least one of Fields is a runtime property.
	Fields     []FieldData // Documentation for individual fields.
}

// FieldData contains the documentation for single field in a ProtoBuf message.
type FieldData struct {
	Name        string // Name of field.
	Anchor      string // Document local named anchor that identifies this field.
	TypeName    string // Type name of field.
	TypeLink    string // Link to type definition.
	IsRequired  bool   // Field is required.
	IsInput     bool   // Field is an input field.
	IsOutput    bool   // Field is an output field.
	IsRuntime   bool   // Runtime field.
	IsFQDN      bool   // Is FQDN
	IsLabel     bool   // Is a label.
	IsOrgLabel  bool   // Organization + label.
	IsReference bool   // Is reference
	IsRepeated  bool   // Is repeated
	IsBytes     bool   // The type is bytes.
	Reference   string // Named reference
	Doc         string // Documentation for this field.
	Number      int32  // Field ordinal
}

// Enumeration. Used for both nested (within a message) enumerations and top
// level enumerations.
type EnumData struct {
	Name   string
	Anchor string
	Doc    string
	Values []EnumValueData
}

// Individual value in an enumeration.
type EnumValueData struct {
	Name   string // Name of enumeration value.
	Anchor string
	Doc    string // Documentation for value.
}

// generator encapsulates the state for generating documentation for a single
// package.
type generator struct {
	Descriptors descriptor.FileDescriptorSet // Merged set of file descriptors.
}

// MergeFileDescriptorSet reads and merges a binary encoded FileDescriptorSet from a file. Can
// be called repeatedly.
func (g *generator) MergeFileDescriptorSet(fname string) error {
	var fdset descriptor.FileDescriptorSet
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(data, &fdset)
	if err != nil {
		return err
	}

	proto.Merge(&g.Descriptors, &fdset)
	return nil
}

// AnchorFromTypeName returns the full link text for a possibly cross package
// refrence. The input is a typename as expected to be found in a message
// descriptor.
//
// These types start with a period if they are "full" references. Otherwise we
// assume that returning just the reference as is is probably fine.
func (g *generator) AnchorFromTypeName(tn string) (typeName, anchor string) {
	switch {
	case strings.HasPrefix(tn, ".asset."):
		return tn[1:], "asset.md#" + tn[7:]

	case strings.HasPrefix(tn, ".host."):
		return tn[1:], "host.md#" + tn[6:]

	case strings.HasPrefix(tn, ".common."):
		return tn[1:], "common.md#" + tn[8:]

	case strings.HasPrefix(tn, ".compute."):
		return tn[1:], "gcp_compute.md#" + tn[9:]
	}

	return tn, ""
}

// extractFromField extracts documentation for a single ProtoBuf field in a message.
//
// Arguments:
//    md : The MessageData object that should receive the generated documentation.
//
//    p  : The path within the FileDescriptorProto to this field. See the |p|
//         argument description in handleMessage for more details.
//
//    fd : The FieldDescriptorProto of the field to be documented.
//
//    parent : The parent message or namespace. Non-empty if this message is
//         nested.
func (g *generator) extractFromField(md *MessageData, p []int32, f *descriptor.FileDescriptorProto,
	d *descriptor.FieldDescriptorProto, parent string) {

	var fd FieldData
	fd.Name = d.GetName()
	fd.Doc = lookupDocs(p, f)
	fd.Anchor = concat(parent, fd.Name)
	fd.Number = d.GetNumber()

	switch d.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_GROUP:
		fallthrough
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		fd.TypeName, fd.TypeLink = g.AnchorFromTypeName(d.GetTypeName())

	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		fd.TypeName, fd.TypeLink = g.AnchorFromTypeName(d.GetTypeName())

	case descriptor.FieldDescriptorProto_TYPE_STRING:
		fd.TypeName = "string"

	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		fd.TypeName = "bool"

	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		fd.TypeName = "bytes"
		fd.IsBytes = true

	case descriptor.FieldDescriptorProto_TYPE_INT32:
		fd.TypeName = "int32"

	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		fd.TypeName = "uint32"

	case descriptor.FieldDescriptorProto_TYPE_INT64:
		fd.TypeName = "int64"

	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		fd.TypeName = "uint64"

	default:
		fd.TypeName = d.GetType().String()
	}

	v := common.GetValidationForField(d)
	fd.IsOutput = v.IsOutput()
	fd.IsRuntime = v.IsRuntime()
	fd.IsLabel = (v.Type == common.Validation_LABEL)
	fd.IsRequired = (v.Type == common.Validation_REQUIRED || (fd.IsLabel && !v.GetOptional()))
	fd.IsOrgLabel = (v.Type == common.Validation_ORGLABEL)
	fd.IsFQDN = (v.Type == common.Validation_FQDN)
	fd.IsReference = v.IsNamedReference()
	fd.Reference = v.Ref

	if d.Label != nil && *d.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
		fd.IsRepeated = true
	}
	fd.IsInput = !fd.IsOutput && !fd.IsRuntime

	md.Fields = append(md.Fields, fd)

	if fd.IsInput {
		md.HasInputs = true
	}
	if fd.IsOutput {
		md.HasOutputs = true
	}
	if fd.IsRuntime {
		md.HasRuntime = true
	}
}

// extractFromMessageType extracts documentation for single ProtoBuf message.
//
// Note that the message may be nested.
//
// Arguments:
//
//    fd : The FileData object which should receive the per-message
//         documentation.
//
//    p  : The path within the FileDescriptorProto to this message. See the
//         documentation of Location.path field in FileDescriptorProto. [1]
//
//         Basically, this is a list of int32s that traverse the
//         FileDescriptorProto message. When traversing from a message to a
//         field, the field number is added to the list. When traversing into a
//         repeated field, the index of the element being traversed to is added
//         to the list.
//
//    m  : The DescriptorProto for the message to be documented.
//
//    parent : The parent message or namespace. Non-empty if this message is
//         nested.
//
// [1]: https://github.com/golang/protobuf/blob/master/protoc-gen-go/descriptor/descriptor.proto
func (g *generator) extractFromMessageType(fd *PackageData, p []int32,
	f *descriptor.FileDescriptorProto,
	m *descriptor.DescriptorProto,
	parent string) {

	var md MessageData
	md.Name = m.GetName()
	md.Doc = lookupDocs(p, f)
	md.Anchor = concat(parent, md.Name)

	nested_ns := concat(parent, md.Name)
	field_p := append(p, 2)
	for i, fd := range m.Field {
		g.extractFromField(&md, append(field_p, int32(i)), f, fd, nested_ns)
	}
	sort.Slice(md.Fields, func(i, j int) bool { return md.Fields[i].Number < md.Fields[j].Number })

	fd.Messages = append(fd.Messages, md)

	enum_p := append(p, 4)
	for i, e := range m.EnumType {
		g.extractFromEnumDesc(fd, append(enum_p, int32(i)), f, e, nested_ns)
	}

	nested_p := append(p, 3)
	for i, nm := range m.NestedType {
		g.extractFromMessageType(fd, append(nested_p, int32(i)), f, nm, nested_ns)
	}
}

// extractFromEnumDesc extracts an enumeration from a EnumDescriptorProto.
//
// The extracted enum is placed in a PackageData structure since all enums are
// placed there.
func (g *generator) extractFromEnumDesc(fd *PackageData,
	p []int32,
	f *descriptor.FileDescriptorProto,
	e *descriptor.EnumDescriptorProto,
	parent string) {

	var ed EnumData
	ed.Name = e.GetName()
	ed.Doc = lookupDocs(p, f)
	ed.Anchor = concat(parent, ed.Name)

	value_p := append(p, 2)
	for i, ev := range e.Value {
		var v EnumValueData
		v.Name = ev.GetName()
		v.Doc = lookupDocs(append(value_p, int32(i)), f)
		v.Anchor = concat(ed.Anchor, v.Name)
		ed.Values = append(ed.Values, v)
	}

	if fd.Enums == nil {
		fd.Enums = make(map[string]EnumData)
	}
	fd.Enums[ed.Anchor] = ed
}

// extractFromFileDesc extracts documentation for a single .proto file as described in
// the FileDescriptorProto that's passed in as a parameter.
func (g *generator) extractFromFileDesc(pd *PackageData, f *descriptor.FileDescriptorProto) {
	pd.SourceFiles = append(pd.SourceFiles, f.GetName())
	if pd.Package == "" {
		pd.Package = f.GetPackage()
		pd.Anchor = pd.Package
	} else if pd.Package != f.GetPackage() {
		fatalf(
			nil,
			"all source files specified in a single session should be for a single package. Found %s and %s",
			pd.Package, f.GetPackage())
	}

	for i, e := range f.EnumType {
		g.extractFromEnumDesc(pd, []int32{5, int32(i)}, f, e, "")
	}

	// Only deal with top-level messages for now. We'd need to add other top
	// level types if we used them in the schema.
	for i, m := range f.MessageType {
		g.extractFromMessageType(pd, []int32{4, int32(i)}, f, m, "")
	}

	if doc := lookupDocs([]int32{2}, f); doc != "" {
		pd.Doc = doc
	}
}

// Gen generates the documentation. Call once all the input files have been
// successfully slurped.
func (g *generator) Gen(writer io.Writer) error {
	var pd PackageData

	for _, f := range g.Descriptors.File {
		g.extractFromFileDesc(&pd, f)
	}
	sort.Slice(pd.Messages, func(i, j int) bool { return pd.Messages[i].Name < pd.Messages[j].Name })

	// Template functions.
	functions := template.FuncMap{
		// "code" wraps its argument in backticks.
		"code": func(s string) string {
			if s != "" {
				return "`" + s + "`"
			} else {
				return ""
			}
		},
		"clist": func(s []string) string { return strings.Join(s, ", ") },
	}

	t := template.Must(template.New("file").Funcs(functions).Parse(DocTemplate))

	return t.Execute(writer, pd)
}

// concat joins a list of strings with '.' in the middle. The first string is
// ignored if its empty.
//
// E.g.:
//     concat("a", "b", "c") == "a.b.c"
//     concat("a", "b")      == "a.b"
//     concat("", "b")       == "b"
//     concat("b")           == "b"
//     concat("")            == ""
//     concat()              == ""
func concat(s ...string) string {
	if len(s) == 0 {
		return ""
	}

	if s[0] == "" {
		s = s[1:]
	}
	return strings.Join(s, ".")
}

// samePath returns trus if |a| and |b| are the same path.
func samePath(a []int32, b []int32) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// dedent removes any common whitespace from each line in a newline delimited
// list of lines.
func dedent(s string) string {
	ls := strings.Split(s, "\n")
	if len(ls) == 1 {
		return s
	}

	sc := -1
	for _, l := range ls {
		if len(l) == 0 || strings.TrimSpace(s) == "" {
			continue
		}
		lc := strings.IndexFunc(l, func(c rune) bool { return c != ' ' })
		if lc == -1 {
			continue
		}
		if lc < sc || sc == -1 {
			sc = lc
		}
	}

	if sc < 1 {
		return s
	}

	f := []string{}
	for _, l := range ls {
		if len(l) < sc {
			f = append(f, l)
			continue
		}
		f = append(f, l[sc:])
	}
	return strings.Join(f, "\n")
}

// lookupDocs extracts the documentation based on an entity location. The value
// |p| is as documented in SourceCodeInfo in [1]. The FileDescriptorProto being
// passed in must have its SourceCodeInfo field intact.
//
// [1]: https://github.com/golang/protobuf/blob/master/protoc-gen-go/descriptor/descriptor.proto
func lookupDocs(p []int32, f *descriptor.FileDescriptorProto) string {
	for _, l := range f.SourceCodeInfo.Location {
		if !samePath(p, l.Path) {
			continue
		}

		// only one path really makes sense here.
		return dedent(l.GetLeadingComments())
	}
	return ""
}
