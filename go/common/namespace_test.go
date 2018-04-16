// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"strings"
	"testing"
)

func TestNamespace_Graft_basic(t *testing.T) {
	m := TestMessageWithOptions{
		Name:        "from.here",
		Key:         "foo",
		OptionalKey: "bar",
		Fqdn:        "${a.b.with_types.xyz} ${a.b.with_types.abc}"}
	var r Namespace

	err := r.Graft(&m, EmptyPath)
	if err == nil {
		t.Fatal("references should not allow placing objects at the root")
	}

	if r.Has(RefPathMust("a.b.with_options")) {
		t.Fatal("unresolved OUTPUT value reported as available")
	}

	err = r.Graft(&m, RefPathMust("a.b.with_options"))
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if !r.Has(RefPathMust("a.b.with_options")) {
		t.Fatal("resolved OUTPUT value reported as unavailable")
	}

	if !r.HasDirectDependency(RefPathMust("a.b.with_options.key"), RefPathMust("a.b.with_types.repeated_field.foo")) {
		t.Fatal("don't find expected dependency")
	}

	if !r.HasDirectDependency(RefPathMust("a.b.with_options.optional_key"), RefPathMust("a.b.with_types.repeated_field.bar")) {
		t.Fatal("don't find expected dependency")
	}

	if !r.HasDirectDependency(RefPathMust("a.b.with_options.fqdn"), RefPathMust("a.b.with_types.xyz")) {
		t.Fatal("don't find expected dependency")
	}

	if !r.HasDirectDependency(RefPathMust("a.b.with_options.fqdn"), RefPathMust("a.b.with_types.abc")) {
		t.Fatal("don't find expected dependency")
	}
}

func TestNamespace_Graft_again(t *testing.T) {
	m := TestMessageWithOptions{
		Name:  "x",
		Label: "${foo.output}",
	}

	var r Namespace
	err := r.Graft(&m, RefPathMust("foo"))
	if err != nil {
		t.Fatal(err)
	}

	err = r.Graft(&m, RefPathMust("foo"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestNamespace_Graft_recursive(t *testing.T) {
	m := TestMessageWithOptions{
		Name:  "x",
		Label: "${foo.label}",
	}

	var r Namespace
	err := r.Graft(&m, RefPathMust("foo"))
	if err == nil || !strings.Contains(err.Error(), "foo.label has a recursive reference") {
		t.Fatal(err)
	}
}

func TestNamespace_ExpandString_basic(t *testing.T) {
	m := TestMessageWithOptions{Fqdn: "${bar.repeated_field.a.name}${bar.repeated_field.b.name}${bar.repeated_field.c}${bar.repeated_field.c.name}"}
	w := TestMessageWithTypes{
		Name: "to",
		RepeatedField: []*TestGoodProto{
			&TestGoodProto{"a"},
			&TestGoodProto{"b"},
			&TestGoodProto{"c"},
		}}

	var r Namespace
	err := r.Graft(&m, RefPathMust("foo"))
	if err != nil {
		t.Fatal(err)
	}
	err = r.Graft(&w, RefPathMust("bar"))
	if err != nil {
		t.Fatal(err)
	}

	if s, err := r.ExpandString(""); err != nil || s != "" {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("abc"); err != nil || s != "abc" {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("\\${abc}"); err != nil || s != "\\${abc}" {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("abc${abc"); err == nil || !strings.Contains(err.Error(), "mismatched braces") {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("abc\\${abc}"); err != nil || s != "abc\\${abc}" {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("${abc}"); err == nil || !strings.Contains(err.Error(), "could not be resolved") {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("${bar.repeated_field.c}"); err == nil || !strings.Contains(err.Error(), "is not a string") {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("a${bar.repeated_field.a.name}c"); err != nil || s != "aac" {
		t.Fatal(s, err)
	}

	if s, err := r.ExpandString("a${bar.repeated_field.a.name}${bar.repeated_field.b.name}${bar.repeated_field.c.name}"); err != nil || s != "aabc" {
		t.Fatal(s, err)
	}
}

func TestNamespace_SetOutput_good(t *testing.T) {
	w := TestMessageWithOptions{
		Name:  "my-name-is-foo",
		Label: "${a.b.foo.output}",
	}

	var r Namespace
	err := r.Graft(&w, RefPathMust("a.b.foo"))
	if err != nil {
		t.Fatal(err)
	}

	// In all the following cases, the test verifies that the published value
	// is reported as unavailable prior to publishing, and is shown as
	// available after publishing.

	// First try a string value.

	if r.Has(RefPathMust("a.b.foo.output")) {
		t.Error("unresolved output field reported as available")
	}

	err = r.PublishOutput(RefPathMust("a.b.foo.output"), "hello")
	if err != nil {
		t.Fatal(err)
	}

	if !r.Has(RefPathMust("a.b.foo.output")) {
		t.Error("resolved output field reported as unavailable")
	}

	// An int value.

	if r.Has(RefPathMust("a.b.foo.output_int")) {
		t.Error("unresolved output field reported as available")
	}

	err = r.PublishOutput(RefPathMust("a.b.foo.output_int"), int32(1))
	if err != nil {
		t.Fatal(err)
	}

	if !r.Has(RefPathMust("a.b.foo.output_int")) {
		t.Error("resolved output field reported as unavailable")
	}

	// A proto.Message value

	if r.Has(RefPathMust("a.b.foo.output_proto")) {
		t.Error("unresolved output field reported as available")
	}

	err = r.PublishOutput(RefPathMust("a.b.foo.output_proto"), &TestGoodProto{})
	if err != nil {
		t.Fatal(err)
	}

	if !r.Has(RefPathMust("a.b.foo.output_proto")) {
		t.Error("resolved output field reported as unavailable")
	}

	// This node also needs to be resolved since all OUTPUT nodes must be
	// available.
	r.PublishOutput(RefPathMust("a.b.foo.output_alt"), "hello")

	// Make sure there are no stragglers.
	r.VisitUnresolved(EmptyPath, func(v UnresolvedValue) bool {
		t.Errorf("node at %s is unresolved.", v.Location)
		return true
	})

	if w.Label != "hello" {
		t.Fatalf("Dependent value not updated correctly. Got \"%s\". Want \"hello\"", w.Label)
	}
}

func TestNamespace_SetOutput_partial(t *testing.T) {
	w := TestMessageWithOptions{
		Label: "${a.b.c.output}",
		Fqdn:  "${a.b.c.output} ${a.b.c.output_alt}!",
	}

	var r Namespace
	r.Graft(&w, RefPathMust("a.b.c"))

	if r.Has(RefPathMust("a.b.c.fqdn")) {
		t.Error("value with unresolved references is reported as available")
	}

	err := r.PublishOutput(RefPathMust("a.b.c.output"), "Hello")
	if err != nil {
		t.Error(err)
	}

	if !r.HasDirectDependency(RefPathMust("a.b.c.fqdn"), RefPathMust("a.b.c.output_alt")) {
		t.Error("missing dependency from static field")
	}

	if w.Label != "Hello" {
		t.Error("dependent value not updated correctly")
	}

	if w.Fqdn != "Hello ${a.b.c.output_alt}!" {
		t.Error("dependent partial value not updated correctly")
	}

	if r.Has(RefPathMust("a.b.c.fqdn")) {
		t.Error("value with unresolved references is reported as available")
	}

	err = r.PublishOutput(RefPathMust("a.b.c.output_alt"), "World")
	if err != nil {
		t.Error(err)
	}

	if w.Fqdn != "Hello World!" {
		t.Error("dependent value not updated correctly")
	}

	if !r.Has(RefPathMust("a.b.c.fqdn")) {
		t.Error("value with resolved references is reported as unavailable")
	}
}

func TestNamespace_SetOutput_duplicate(t *testing.T) {
	w := TestMessageWithOptions{}

	var r Namespace
	r.Graft(&w, RefPathMust("a.b.c"))

	err := r.PublishOutput(RefPathMust("a.b.c.output"), "Hello")
	if err != nil {
		t.Error(err)
	}

	err = r.PublishOutput(RefPathMust("a.b.c.output"), "World")
	if err == nil || !strings.Contains(err.Error(), "at a.b.c.output which already has a value") {
		t.Error(err)
	}
}

func TestNamespace_SetOutput_otherTypes(t *testing.T) {
	w := TestMessageWithOptions{}

	var r Namespace
	r.Graft(&w, RefPathMust("a.b.c"))

	err := r.PublishOutput(RefPathMust("a.b.c.output_proto"), &TestGoodProto{
		Name: "just-added",
	})

	if err != nil {
		t.Error(err)
	}

	err = r.PublishOutput(RefPathMust("a.b.c.output_int"), int32(1))
	if err != nil {
		t.Error(err)
	}

	v, err := r.Get(RefPathMust("a.b.c.output_proto.name"))
	if err != nil {
		t.Fatal(err)
	}

	if s, ok := v.(string); ok {
		if s != "just-added" {
			t.Errorf("newly added field has unexpected value. Got \"%s\". Want \"just-added\"", s)
		}
	} else {
		t.Error("unexpected type while looking up newly added Message")
	}
}

func TestNamespace_SetOutput_notOutput(t *testing.T) {
	w := TestMessageWithOptions{
		Name:  "foo",
		Label: "${foo.output}",
	}

	var r Namespace
	r.Graft(&w, RefPathMust("foo"))

	err := r.PublishOutput(RefPathMust("foo.label"), "hello")
	if err == nil || !strings.Contains(err.Error(), "at foo.label which is not an OUTPUT") {
		t.Fatal(err)
	}
}

func TestNamespace_SetOutput_notKnown(t *testing.T) {
	w := TestMessageWithOptions{
		Name:  "foo",
		Label: "${foo.output}",
	}

	var r Namespace
	r.Graft(&w, RefPathMust("foo"))

	err := r.PublishOutput(RefPathMust("bar.label"), "hello")
	if err == nil || !strings.Contains(err.Error(), "not part of the current namespace") {
		t.Fatal(err)
	}
}

func TestNamespace_SetOutput_badType(t *testing.T) {
	w := TestMessageWithOptions{
		Name:  "foo",
		Label: "${foo.output}",
	}

	var r Namespace
	r.Graft(&w, RefPathMust("foo"))

	err := r.PublishOutput(RefPathMust("foo.output"), 9)
	if err == nil ||
		!strings.Contains(err.Error(), "type int is not assignable to type string") ||
		!strings.Contains(err.Error(), "while assigning value to foo.output. Value is 9") {
		t.Fatal(err)
	}
}

func TestNamespace_VisitUnresolved_basic(t *testing.T) {
	w := TestMessageWithOptions{
		Label: "${a.b.c.output}",
		Fqdn:  "${a.b.c.output} ${a.b.c.output_alt}!",
	}

	var r Namespace
	r.Graft(&w, RefPathMust("a.b.c"))

	seen := make(map[string]bool)

	r.VisitUnresolved(EmptyPath, func(v UnresolvedValue) bool {
		seen[v.Location.String()] = true
		switch v.Location.String() {
		case "a.b.c.output":
			if u, ok := v.Value.(UnresolvedValue_Output); ok {
				if len(u.Dependents) != 2 {
					t.Errorf("incorrect dependents for OUTPUT field. Got %s. Want two fields", u)
				}
			} else {
				t.Error("incorrect unresolved value type")
			}

		case "a.b.c.fqdn":
			if u, ok := v.Value.(UnresolvedValue_String); ok {
				if len(u.DependsOn) != 2 {
					t.Error("incorrect number of dependencies")
				}
			} else {
				t.Error("incorrect unresolved value type")
			}
		}
		return true
	})

	has := func(s string) bool {
		_, ok := seen[s]
		return ok
	}

	if len(seen) != 6 ||
		!has("a.b.c.output_proto") ||
		!has("a.b.c.output") ||
		!has("a.b.c.output_alt") ||
		!has("a.b.c.output_int") ||
		!has("a.b.c.fqdn") ||
		!has("a.b.c.label") {
		t.Errorf("unexpected number/kinds of unresolved values observed. Seen %#v", seen)
	}
}

func TestNamespace_asAssetGraph_sanityCheck(t *testing.T) {
	var r Namespace
	xPath := RefPathMust("a.b.x")
	yPath := RefPathMust("a.b.y")
	zPath := RefPathMust("a.b.z")
	nPath := RefPathMust("a.b.z.output_proto")

	err := r.Graft(&TestMessageWithOptions{Name: "x", Label: "${a.b.y.label}"}, xPath)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Graft(&TestMessageWithOptions{Name: "y", Label: "y-label"}, yPath)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Graft(&TestMessageWithOptions{Name: "z", OutputProto: &TestGoodProto{Name: "nested"}}, zPath)
	if err != nil {
		t.Fatal(err)
	}

	xNode, _ := r.getNode(xPath)
	yNode, _ := r.getNode(yPath)
	zNode, _ := r.getNode(zPath)
	nNode, _ := r.getNode(nPath)
	if xNode == nil || yNode == nil || zNode == nil || nNode == nil {
		t.Fatal("nodes not in namespace graph")
	}

	ids := make(map[int64]bool)
	ids[xNode.ID()] = true
	ids[yNode.ID()] = true
	ids[zNode.ID()] = true
	ids[nNode.ID()] = true
	if len(ids) != 4 {
		t.Fatal("nodes don't have unique ids")
	}

	g, err := r.asAssetGraph()
	if err != nil {
		t.Fatal(err)
	}

	if !g.Has(xNode) || !g.Has(yNode) || !g.Has(zNode) || !g.Has(nNode) {
		t.Errorf("nodes missing from graph: x(%v), y(%v), z(%v), n(%v)",
			g.Has(xNode), g.Has(yNode), g.Has(zNode), g.Has(nNode))
	}

	if !g.HasEdgeFromTo(xNode, yNode) {
		t.Fatal("dependency edge missing for inner reference")
	}

	if !g.HasEdgeFromTo(zNode, nNode) {
		t.Fatal("parent->child edge missing")
	}
}

func TestNamespace_asTopologicalList_noCycles(t *testing.T) {
	var r Namespace
	xPath := RefPathMust("a.b.x")
	yPath := RefPathMust("a.b.y")
	zPath := RefPathMust("a.b.z")

	r.Graft(&TestMessageWithOptions{Name: "x", Label: "${a.b.y.label}"}, xPath)
	r.Graft(&TestMessageWithOptions{Name: "y", Label: "${a.b.z.output_proto.name}"}, yPath)
	r.Graft(&TestMessageWithOptions{Name: "z", Label: "${a.b.x.label}",
		OutputProto: &TestGoodProto{Name: "nested"}}, zPath)

	l, err := r.asTopologicalList()
	if err != nil {
		t.Fatal(err)
	}

	if len(l) != 4 {
		for _, v := range l {
			t.Error(v.location.String())
		}
		t.Fatalf("wrong length for tolopolgical list %d", len(l))
	}

	if l[0].location.String() != "a.b.z.output_proto" ||
		l[1].location.String() != "a.b.y" ||
		l[2].location.String() != "a.b.x" ||
		l[3].location.String() != "a.b.z" {
		t.Error(l[0].location, l[1].location, l[2].location, l[3].location)
	}
}

func TestNamespace_asTopologicalList_cycles(t *testing.T) {
	var r Namespace
	xPath := RefPathMust("a.b.x")
	yPath := RefPathMust("a.b.y")
	zPath := RefPathMust("a.b.z")

	r.Graft(&TestMessageWithOptions{Name: "x", Label: "${a.b.y.label}"}, xPath)
	r.Graft(&TestMessageWithOptions{Name: "y", Label: "${a.b.z.label}"}, yPath)
	r.Graft(&TestMessageWithOptions{Name: "z", Label: "${a.b.x.label}",
		OutputProto: &TestGoodProto{Name: "nested"}}, zPath)

	_, err := r.asTopologicalList()
	if err == nil {
		if d, err := r.AsSerializedDOTGraph(); err == nil {
			t.Errorf("asTopologicalList didn't detect cycles\n%s", d)
		}
	} else {
		if !strings.Contains(err.Error(), "no topological ordering") {
			t.Errorf("unexpected error: %s", err.Error())
		}
	}
}

func TestNamespace_nearestMessage(t *testing.T) {
	var r Namespace
	xPath := RefPathMust("a.b.x")
	yPath := RefPathMust("a.b.y")
	zPath := RefPathMust("a.b.z")

	r.Graft(&TestMessageWithOptions{Name: "x", Label: "${a.b.y.label}"}, xPath)
	r.Graft(&TestMessageWithOptions{Name: "y", Label: "${a.b.x.label}"}, yPath)
	r.Graft(&TestMessageWithOptions{Name: "z", Label: "${a.b.n.label}"}, zPath)

	xLabel, ok := r.getNode(xPath.Append("label"))
	if !ok {
		t.Fatal("couldn't get node for label field")
	}

	xMsg, err := r.enclosingMessage(xLabel)
	if err != nil || xMsg.location.String() != "a.b.x" {
		t.Error("nearestMessage failed to ascend to parent message")
	}

	xMsg, ok = r.getNode(xPath)
	if !ok {
		t.Error("Can't get existing node")
	}

	xMsg, err = r.enclosingMessage(xMsg)
	if err != nil || xMsg.location.String() != "a.b.x" {
		t.Error("nearestMessage failed to match self")
	}

	if len(xLabel.dependsOn) != 1 {
		t.Errorf("incorrect number of dependencies. Got %d. Want 1", len(xLabel.dependsOn))
	}
}

func TestNamespace_Prune_basic(t *testing.T) {
	var r Namespace
	xPath := RefPathMust("a.b.x")
	yPath := RefPathMust("a.b.y")
	zPath := RefPathMust("a.b.z")

	r.Graft(&TestMessageWithOptions{Name: "x", Label: "${a.b.y.label}"}, xPath)
	r.Graft(&TestMessageWithOptions{Name: "y", Label: "y-label"}, yPath)
	r.Graft(&TestMessageWithOptions{Name: "z"}, zPath)

	if !r.Has(xPath) || !r.Has(yPath) || !r.Has(zPath) {
		t.Fatal()
	}

	err := r.Prune([]RefPath{xPath})
	if err != nil {
		t.Fatal(err)
	}

	if !r.Has(xPath) || !r.Has(yPath) {
		t.Error("required paths missing")
	}

	if r.Has(zPath) {
		t.Error("failed to prune unnecessary node")
	}
}

func TestNamespace_anneal(t *testing.T) {
	var r Namespace
	xPath := RefPathMust("a.b.x")
	yPath := RefPathMust("a.b.y")
	zPath := RefPathMust("a.b.z")

	r.Graft(&TestMessageWithOptions{Name: "x", Label: "x-${a.b.y.label}"}, xPath)
	r.Graft(&TestMessageWithOptions{Name: "y", Label: "y-label"}, yPath)
	r.Graft(&TestMessageWithOptions{Name: "z", Label: "z-${a.b.x.label}"}, zPath)

	if o, err := r.Get(xPath.Append("label")); err == nil {
		if o.(string) != "x-y-label" {
			t.Error(o)
		}
	} else {
		t.Error(err)
	}

	if o, err := r.Get(zPath.Append("label")); err == nil {
		if o.(string) != "z-x-y-label" {
			t.Error(o)
		}
	} else {
		t.Error(err)
	}

	if !r.InputsReady(xPath) || !r.InputsReady(yPath) || !r.InputsReady(zPath) || !r.InputsReady(EmptyPath) {
		t.Error("all subtrees should be resolved by now")
	}
}
