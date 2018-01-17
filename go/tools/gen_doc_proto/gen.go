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

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// text/template compatible template for a Markdown document. Each package is
// described in a single .md file.
//
// The fields in this template are based on a pipeline containing a PackageData
// object.
const DocTemplate = `<!-- Generated file. Do not edit -->

# Package {{code .Package}} {#{{.Anchor}}}

{{.Doc}}

Messages that are valid in package {{code .Package}} are as follows:
{{range .Messages -}}
## {{code .Name}} {#{{.Anchor}}}

{{.Doc}}

Fields for {{code .Name}}:
{{range .Fields}}
### {{code .Name}} {#{{.Anchor}}}

{{.Doc}}
{{- end}}
{{end}}

Generated from {{clist .SourceFiles | code}}.
`

// PackageData contains the documentation for a single package file.
type PackageData struct {
	Package     string        // Package name. All symbols in this file are contained in this package.
	Anchor      string        // Document local named anchor.
	Doc         string        // Package documentation.
	SourceFiles []string      // Relative path to the source .proto files.
	Messages    []MessageData // List of messages.
}

// MessageData contains the documentation and link information for a single
// ProtoBuf message.
type MessageData struct {
	Name   string      // Name of message.
	Anchor string      // Document local named anchor that identifies this message.
	Doc    string      // Toplevel documentation for message.
	Fields []FieldData // Documentation for individual fields.
}

// FieldData contains the documentation for single field in a ProtoBuf message.
type FieldData struct {
	Name   string // Name of field.
	Anchor string // Document local named anchor that identifies this field.
	Doc    string // Documentation for this field.
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
	fd *descriptor.FieldDescriptorProto, parent string) {

	var fdata FieldData
	fdata.Name = fd.GetName()
	fdata.Doc = lookupDocs(p, f)
	fdata.Anchor = concat(parent, fdata.Name)

	md.Fields = append(md.Fields, fdata)
}

// extractFromMessageType extracts documentation for single ProtoBuf message.
//
// Note that the message may be nested.
//
// Arguments:
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
func (g *generator) extractFromMessageType(fd *PackageData, p []int32, f *descriptor.FileDescriptorProto,
	m *descriptor.DescriptorProto, parent string) {

	var md MessageData
	md.Name = m.GetName()
	md.Doc = lookupDocs(p, f)
	md.Anchor = concat(parent, md.Name)

	nested_ns := concat(parent, md.Name)
	field_p := append(p, 2)
	for i, fd := range m.Field {
		g.extractFromField(&md, append(field_p, int32(i)), f, fd, nested_ns)
	}
	sort.Slice(md.Fields, func(i, j int) bool { return md.Fields[i].Name < md.Fields[j].Name })

	fd.Messages = append(fd.Messages, md)

	nested_p := append(p, 3)
	for i, nm := range m.NestedType {
		g.extractFromMessageType(fd, append(nested_p, int32(i)), f, nm, nested_ns)
	}
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

	// Only deal with top-level messages for now. We'd need to add other top
	// level types if we used them in the schema.
	for i, m := range f.MessageType {
		g.extractFromMessageType(pd, []int32{4, int32(i)}, f, m, pd.Package)
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
		"code":  func(s string) string { return "`" + s + "`" },
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
		return l.GetLeadingComments()
	}
	return ""
}
