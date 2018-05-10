// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/golang/protobuf/proto"
	pd "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
	"gonum.org/v1/gonum/graph/traverse"
	"reflect"
)

// Namespace represents a hierarchy of proto.Message objects possibly with
// internal cross references.
//
// This is the primary mechanism used by the CEL toolchain for organizing all
// the resources that go into a lab. At the heart of the Namespace is a Trie
// which maps every addressable path to a corresponding namespaceNode object.
//
//
// Invariants
//
// The following characteristics can be assumed to be true:
//
//   * For every addressable and existing path in the Namespace, there will be
//     a unique corresponding namespaceNode. I.e. there's a 1:1 correspondence
//     between a reference path in the namespace and namespaceNodes.
//
//   * The relationship between a *path* and its corresponding namespaceNode
//     holds for the entire lifetime of the Namespace. I.e. if "a.b.c" maps to
//     namespaceNode nn, then nn will always refer to whichever value exists at
//     that path. This holds true even if the value changes.
//
//   * The namespace can be extended by grafting new proto.Message objects, but
//     it cannot be reduced. Once an object is grafted it cannot be removed for
//     the lifetime of the Namespace.
//
//   * The dependency graph as understood by the Namespace object is always
//     up-to-date as long as all consumers adhere to the API contract as
//     spelled out below.
//
//
// Requirements
//
// The following requirements should be met by all consumers of Namespace:
//
//   * Once an object is grafted to the Namespace via Graft(), that object is
//     owned by the namespace. The object should not be modified directly (See
//     below).
//
//   * An object once grafted should only be modified via calls to Publish()
//     and PublishOutput() methods of Namespace. Doing so ensures referential
//     integrity across the namespace.
//
//
// Why This Is Needed
//
// It is expected that the CEL toolchain will need to work with asset manifests
// with hundreds, if not thousands of resources with interdependencies.
// Therefore it is important for us to be able to efficiently manage these
// dependencies.
//
// All resources are represented using proto.Message objects. Protocol buffers
// allow us the luxury if not having t o deal with representation formats and
// serialization issues. But the resource's external representation -- i.e. the
// one that's described in the Protobuf schema -- is small subset of the
// internal state that needs to be managed in order to properly deploy the
// resource.
//
// The mechanism by which we consume the Go types generated by the Protocol
// Buffers compiler and extend it to be used by the CEL toolchain during
// deployment and on-host configuration is via Namespaces.
//
// It provides us with mechanisms for :
//
//   * ... locating an object or a field given a path regardless of
//         the types of the objects that need t o be traversed.
//
//   * ... determining a path given an object.
//
//   * ... maintaining an on-going graph of dependencies across
//         objects and fields.
//
//   * ... propagating a change to a value such that value-based dependencies
//         across objects are taken into account.
//
//   * ... interpreting the objects in the namespace as a dependency graph for
//         the purpose of analyzing and resolving said dependencies.
//
//
// On Dependency Management
//
// Broady, Namespace manages two types of dependencies: Inter-object
// dependencies, and value-based dependencies.
//
// "value-based dependencies" in this context refer to dependencies introduced
// by inline references. E.g. If you have the following assets:
//
//     machine {
//       name: "foo"
//       label: "${machine.bar.label}"
//     }
//
//     machine {
//       name: "bar"
//       label: "baz"
//     }
//
// Then "machine.foo.label" can be said to have a dependency on
// "machine.bar.label". By extension, we can also say that "machine.foo"
// depends on "machine.bar".
//
// If, instead, "machine.bar.label" was an OUTPUT field, (i.e. one whose value
// is provided by a resolver rather than by the input asset description), then
// that value would not be known at the onset. Thus the Namespace will track a
// dependency from "machine.foo.label" to "machine.bar.label". Once the value
// for "machine.bar.label" is provided via Publish() or PublishOutput(), then
// Namespace can automatically propagate the new value to "machine.foo.label".
//
// Inter-object dependencies are similar, except they don't propagate known
// values the way value-based dependencies do. These inter-object dependencies
// can be annotated in the input schema as explicit named references across the
// namespace or added explicitly by a resolver using the PublishDependency()
// method on the Namespace. Either way, Namespace will track the added
// dependency and use it for generating dependency information during asset
// resolution.
type Namespace struct {
	// The entire namespace.
	// - Value: proto.Message objects
	// - Key: location in ns where Value should be placed.
	ns Trie

	// All proto.Message objects in namespaces show up here as keys. This way
	// we can quickly go from a proto.Message to important attributes like it's
	// location in the namespace and it's resolution state.
	messages map[proto.Message]*namespaceNode

	// nextId is the ID that should be assigned to the next refNode in this
	// object. Unique numerical identifiers are required for using the gonum
	// libraries with our reference graphs.
	nextId int64
}

// Graft adds a proto.Message object at the specified RefPath. One way
// to think about it is that calling this function mounts or anchors the object
// `m` at path `root`.
//
// Only a single object can be added at a specific path. However, the object
// being added may overlap with an object that's enclosing the `root` RefPath.
//
// Once grafted, `m` cannot be un-grafted. Any objects included in `m` will
// become part of the namespace.
//
// This part is important: Once an object is grafted, it is considered "owned"
// by the Namespace. The object should not be mutated directly since those
// changes will not be reflected
func (r *Namespace) Graft(m proto.Message, root RefPath) error {
	if len(root) == 0 {
		return errors.New("objects cannot be placed at the root of the namespace")
	}

	err := r.collectFrom(root, m)
	if err != nil {
		return err
	}

	// TODO(asanka): This isn't super efficient. It's probably better to wait
	// until all subtrees are grafted and anneal the combined namespace.
	return r.anneal()
}

// Prune removes all nodes from References that aren't under any path in
// `anchors` and aren't referenced by any node under `anchors`.
func (r *Namespace) Prune(anchors []RefPath) error {
	if len(anchors) == 0 {
		return errors.New("attempting to prune with no anchors")
	}

	var t traverse.BreadthFirst
	g, err := r.asAssetGraph()
	if err != nil {
		return err
	}

	for _, a := range anchors {
		// Promote a to its top level node. Otherwise, "a" stays as is.
		if ta, err := r.TopLevel(a); err == nil {
			a = ta
		}

		// Visit all the top level nodes starting at a:
		r.ns.VisitFrom(a, func(p RefPath, o interface{}) bool {
			if !r.IsTopLevel(p) {
				return true
			}

			n := o.(*namespaceNode)
			if n.isNodeRemoved {
				return true
			}

			if t.Visited(n) {
				return true
			}

			t.Walk(g, n, nil)
			return true
		})
	}

	toPrune := make(namespaceNodeSet)

	// t should have visited all the nodes that we are interested in keeping.
	// We now need to remove all top level nodes that aren't in t.
	r.ns.Visit(func(p RefPath, o interface{}) bool {
		if !r.IsTopLevel(p) {
			return true
		}
		n := o.(*namespaceNode)
		if n.isNodeRemoved {
			return true
		}

		if !t.Visited(n) {
			toPrune[n] = true
		}
		return true
	})

	// and finally, for each node in toPrune, mark them as removed
	for n, _ := range toPrune {
		r.ns.VisitFrom(n.location, func(p RefPath, o interface{}) bool {
			n := o.(*namespaceNode)
			n.isNodeRemoved = true
			return true
		})
	}
	return nil
}

// TopoVisit visits each proto message in the namespace in topological order.
//
// The function invokes `visitor` on each top-level proto.Message with the
// guarantee that the message's dependents have already been successfully
// visited. A visit is considered complete when the `visitor` function returns.
//
// If the visit fails (i.e. the visitor function returns a non-nil error), then
// the node is considered unresolved and the resolution process stops.
//
// It is recommended that a resolver call InputsReady() to ensure that there
// are no inputs with internal references still waiting to be resolved. Even
// though ancestor nodes are resolved prior to resolving a node, it is still
// possible for ancestor nodes not publish all their outputs, or a specific
// resolution run to not resolve a message completely.
//
// Do not modify the messages that are passed in. In addition, namespace
// changing functions like Graft() and Purge() should also not be called during
// a visit.
//
// The `visitor` function may be called using farmed out goroutines. It is safe
// for a visitor to access any fields in `m` including embedded messages
// therein. It is also safe for `visitor` to access messages that `m`
// explicitly depends on since those messages are guaranteed to have been
// visited prior to the invocation of `visitor` on `m`. All other accesses of
// the namespace should be considered unsafe. If resolving `m` requires
// accessing other messages, they should be listed as explicit or implicit
// dependencies of `m`.
//
// The return value of TopoVisit() is the aggregated errors of the entire
// visit. Aggregation is performed using common.WrapErrorList().
func (r *Namespace) TopoVisit(visitor func(m proto.Message) error) error {
	tl, err := r.asTopologicalList()
	if err != nil {
		return err
	}

	// TODO(asanka): This section can be parallelized since tl[i] only needs to
	// wait on its dependencies. As written it waits for all of tl[0:i] to be
	// done.
	//
	// TODO(asanka): This section must deal with panics in the visitor()
	// function.
	for _, n := range tl {
		err = visitor(n.value.Interface().(proto.Message))
		if err != nil {
			return err
		}
	}
	return nil
}

// InputsReady returns true if the namespace rooted at `root` does not have any
// unresolved inputs or placeholders. It is still valid for an output field to
// be unresolved.
func (r *Namespace) InputsReady(root RefPath) bool {
	return r.ns.VisitFrom(root, func(p RefPath, o interface{}) bool {
		n := o.(*namespaceNode)
		// note that removed nodes are vacuously ready.
		if n.isRuntime || n.isOutput || n.isNodeRemoved {
			return true
		}

		if n.isPlaceholder {
			return false
		}

		return n.isValueAvailable
	})
}

// Ready returns true if the namespace rooted at `root` does not have any
// unresolved values. It looks at all values, including outputs.
func (r *Namespace) Ready(root RefPath) bool {
	return r.ns.VisitFrom(root, func(p RefPath, o interface{}) bool {
		n := o.(*namespaceNode)
		// note that removed nodes are vacuously ready.
		if n.isRuntime || n.isNodeRemoved {
			return true
		}

		if n.isPlaceholder {
			return false
		}

		return n.isValueAvailable
	})
}

// PublishOutput adds a new value to the namespace. The new value must be a known
// OUTPUT. It may itself have references or OUTPUT parameters. The latter is
// only possible if the object being added is a `proto.Message`.
//
// It is an error to add a value to the namespace in a location that is not
// marked as OUTPUT.
func (r *Namespace) PublishOutput(l RefPath, newValue interface{}) error {
	v, ok := r.getNode(l)
	if !ok {
		return errors.Errorf("attempt to publish a value at path %s which is not part of the current namespace", l)
	}

	return r.assign(v, newValue)
}

// Publish adds a new value relative to proto.Message. This is equivalent to
// using PublishOutput() with the full path, but is much more convenient if one
// only has the message.
func (r *Namespace) Publish(m proto.Message, field string, value interface{}) error {
	v, ok := r.messages[m]
	if !ok {
		return errors.Errorf("failed to publish value. The input message was not found in the namespace: %#v", m)
	}

	if field != "" {
		v, ok = r.getNode(v.location.Append(field))
		if !ok {
			return errors.Errorf("failed to publish value. The message (%#v) does not have a field named %s", m, field)
		}
	}

	return r.assign(v, value)
}

// Publish an additional dependency.
func (r *Namespace) PublishDependency(m proto.Message, dependsOn RefPath) error {
	from, ok := r.messages[m]
	if !ok {
		return errors.Errorf("failed to publish dependency. The dependent node was not found in the namespace: %#v", m)
	}

	to, ok := r.getNode(dependsOn)
	if !ok {
		return errors.Errorf("failed to publish dependency from \"%s\" to \"%s\". The dependent path (\"%s\") was not found in the namespace.", from.location, dependsOn, dependsOn)
	}

	return to.addDependent(from)
}

// Get returns the value a location specified as a RefPath. If the location
// doesn't exist or the value is not known, the function returns a non-nil
// error. Otherwise the returned `interface{}` is the value found at the
// location. Resolved values shadow original values.
func (r *Namespace) Get(p RefPath) (interface{}, error) {
	vr, ok := r.getNode(p)
	if !ok {
		return nil, errors.Errorf("reference not found \"%s\"", p.String())
	}
	if !vr.isValueAvailable {
		return nil, errors.Errorf("value at \"%s\" is not currently available", p.String())
	}
	return vr.value.Interface(), nil
}

// Has returns true if the value at a RefPath is available.  For a value to be
// available, it must exist and have a known value. This excludes OUTPUT fields
// that have not been resolved, and string fields with unresolved references.
func (r *Namespace) Has(p RefPath) bool {
	if len(p) <= 1 {
		// no values are allowed at the root level.
		return false
	}

	vr, ok := r.getNode(p)
	if !ok {
		return false
	}

	return vr.isValueAvailable && !vr.isNodeRemoved
}

// Indirect indirects (or follows) a named reference.
//
// It's function is best explained via an example. Suppose you have a field
// that is a named reference to another as below:
//
//   windows_user {
//     ...
//     container {
//       ad_domain: "mydomain.example"
//     }
//   }
//
// A resolver will aquire a reference to the outer WindowsUser message, and in
// turn the inner WindowsContainer message. However, it will also likely need
// access to the ActiveDirectoryDomain object that corresponds to the
// "ad_domain" field in WindowsContainer.
//
// This can be resolved as follows:
//
//    v, err := namespace.Indirect(u.Container, "ad_domain")
//    if err != nil {
//      // handle error
//    }
//    domain, ok := v.(*ActiveDirectoryDomain)
//
// This pattern effectively indirects or follows a reference across objects in
// a single namespace.
//
// If there's an error, the returned "err" will be non-nil.
func (r *Namespace) Indirect(m proto.Message, field string) (interface{}, error) {
	p, ok := r.PathFor(m)
	if !ok {
		return nil, errors.New("reference not found for message")
	}
	return r.IndirectReference(p.Append(field))
}

// PathFor returns a RefPath given a proto.Message. Returns EmptyPath and false
// if the message is not known.
func (r *Namespace) PathFor(m proto.Message) (RefPath, bool) {
	n, ok := r.messages[m]
	if !ok {
		return EmptyPath, false
	}
	return n.location, true
}

// TopLevel returns the path to the nearest top level ancestor of a given path.
// Returns the path itself if the path refers to a top level node. On error,
// returns nil along with a non-nil error. Note that if there are no top level
// ancestors of the node, that is considered an error.
func (r *Namespace) TopLevel(p RefPath) (RefPath, error) {
	if p.Empty() {
		return p, nil
	}

	for pp := p.Parent(); !pp.Empty(); {
		n, ok := r.getNode(pp)
		if !ok {
			return nil, errors.Errorf("\"%s\" is not part of the namespace", pp)
		}

		if n.isTopLevelCollection {
			return p, nil
		}

		p = pp
		pp = p.Parent()
	}
	return nil, errors.Errorf("\"%s\" has no top level ancestors")
}

// IsTopLevel returns true if a given path refers to a top level node.
func (r *Namespace) IsTopLevel(p RefPath) bool {
	if p.Empty() {
		return false
	}
	pp := p.Parent()
	pn, ok := r.getNode(pp)
	if !ok {
		return false
	}
	return pn.isTopLevelCollection
}

// IndirectReference returns the object referred to by a field. It's an error
// to call this on a field that is not a named reference.
func (r *Namespace) IndirectReference(p RefPath) (interface{}, error) {
	n, ok := r.getNode(p)
	if !ok {
		return nil, errors.Errorf("\"%s\" is not a part of the namespace", p)
	}
	if !n.isValueAvailable {
		return nil, errors.Errorf("value at \"%s\" is not currently available", p.String())
	}
	if n.referenceRoot.Empty() {
		return nil, errors.Errorf("value at \"%s\" is not a reference", p.String())
	}
	return r.Get(n.referenceRoot.Append(n.value.String()))
}

// ExpandString string expands the resolved string references in the string
// `s`. For the call to be successful, all references in `s` must either have
// already been resolved or they should be unresolved OUTPUTs.
//
// Any unresolved OUTPUT values in `s` will be left untouched. These are not
// considered errors.
//
// It is an error if `s` contains references that are unknown in the namespace.
// In such cases, or if any string references are invalid, the function returns
// an error.
func (r *Namespace) ExpandString(s string) (string, error) {
	refs, err := extractInlineReferences(s)
	if err != nil {
		return "", err
	}
	if len(refs) == 0 {
		return s, nil
	}

	var errList []error

	for _, ref := range refs {
		rv, err := r.Get(ref.Ref)
		if err != nil {
			if iu, ok := r.getNode(ref.Ref); ok && iu.isOutput {
				// Unresolved output field. Leave it as-is.
				continue
			}
			errList = append(errList, &UnresolvedReferenceError{
				To:        ref.Ref,
				InlineRef: s,
				Reason:    "target is not a part of this namespace"})
			continue
		}

		repls, ok := rv.(string)
		if !ok {
			errList = append(errList, &UnresolvedReferenceError{
				To:        ref.Ref,
				InlineRef: s,
				Reason:    "target of reference is not a string"})
			continue
		}

		s = s[:ref.Offset] + repls + s[ref.Offset+ref.Length:]
	}

	return s, WrapErrorList(errList)
}

// VisitUnresolved invokes a function over all unresolved dependencies in this
// namespace.
//
// The argument to the visitor function is an UnresolvedValue object.
//
// The function should return true to continue the traversal. A return value of
// false aborts the traversal.
//
// `root` identifies the starting point of the traversal. All nodes in the
// namespace at and below root are part of the traversal, which will be
// performed with an unspecified order.
func (r *Namespace) VisitUnresolved(root RefPath, visitor func(v UnresolvedValue) bool) {
	r.ns.VisitFrom(root, func(p RefPath, o interface{}) bool {
		vr := o.(*namespaceNode)

		if vr.isValueAvailable {
			return true
		}

		if vr.isPlaceholder {
			visitor(UnresolvedValue{Location: vr.location,
				Value: UnresolvedValue_Placeholder{
					Referrers: nodeSetToRefPathList(vr.dependents),
				}})
			return true
		}

		if vr.isOutput {
			visitor(UnresolvedValue{Location: vr.location,
				Value: UnresolvedValue_Output{
					Dependents: nodeSetToRefPathList(vr.dependents),
				}})
			return true
		}

		visitor(UnresolvedValue{Location: vr.location,
			Value: UnresolvedValue_String{
				DependsOn: nodeSetToRefPathList(vr.dependsOn),
			}})
		return true
	})
}

// HasDirectDependency returns true if the value at `from` directly depends on
// the value at `to.
func (r Namespace) HasDirectDependency(from RefPath, to RefPath) bool {
	vfrom, ok := r.getNode(from)
	if !ok {
		return false
	}

	vto, ok := r.getNode(to)
	if !ok {
		return false
	}

	_, ok = vto.dependents[vfrom]
	return ok
}

// collectFrom scans the proto.Message at RefPath `l` for any unresolved
// values. These can be unresolved due to output dependencies (in the case of
// named references), or these can be due to internal string references, or
// these can be due to the value being an OUTPUT field.
//
// This function is idempotent.
func (r *Namespace) collectFrom(l RefPath, m proto.Message) error {
	if r.messages == nil {
		r.messages = make(map[proto.Message]*namespaceNode)
	}

	return WalkProtoMessage(m, l, func(av reflect.Value, p RefPath, f *pd.FieldDescriptorProto) error {
		node := r.newNode(p)

		if f == nil { // a message
			err := node.bind(av, nil)
			if err != nil {
				return err
			}
			node.isValueAvailable = true // message types are always considered available.

			if m, ok := av.Interface().(proto.Message); ok && m != nil {
				r.messages[m] = node
			} else {
				return errors.Errorf("unexpected field type while parsing node at %s", p)
			}
			return nil
		}

		v := GetValidationForField(f)
		err := node.bind(av, &v)
		if err != nil {
			return err
		}

		if av.Kind() == reflect.String {
			refs, err := extractInlineReferences(av.String())
			if err != nil {
				return err
			}

			for _, ref := range refs {
				if ref.Ref.Equals(p) {
					return errors.Errorf("value at %s has a recursive reference", p.String())
				}

				node.addParent(r.newNode(ref.Ref))
			}

			if len(refs) != 0 {
				node.isValueAvailable = false
			}
		}

		// If the field does not contain a named reference, we are done with this field.
		if !v.IsNamedReference() {
			return nil
		}

		node.referenceRoot, err = v.ReferenceRoot()
		if err != nil {
			return errors.Wrapf(err, "invalid reference in validation string %#v", v.Ref)
		}

		// Also if the value is empty or is not a string, we are done. Here we
		// are assuming that the proto is valid. I.e. Validate(m) will return
		// true. Thus the empty value is not indicative of an error. It's
		// likely that the field was optional.
		if av.Kind() != reflect.String || av.Len() == 0 {
			return nil
		}

		// Named references cannot contain internal references.
		if ContainsInlineReferences(av.String()) {
			return errors.Errorf("named reference contains internal references at %s : value is \"%s\"",
				p.String(), av.String())
		}

		target := node.referenceRoot.Append(av.String())
		node.addParent(r.newNode(target))
		return nil
	})
}

func (r *Namespace) getNextNodeId() int64 {
	id := r.nextId
	r.nextId += 1
	return id
}

func (r *Namespace) newNode(l RefPath) *namespaceNode {
	if v, ok := r.getNode(l); ok {
		return v
	}
	v := &namespaceNode{
		location:      l,
		isPlaceholder: true,
		id:            r.getNextNodeId(),
	}
	r.ns.Set(l, v, false)
	return v
}

func (r Namespace) getNode(l RefPath) (*namespaceNode, bool) {
	if uo, ok := r.ns.Get(l); ok {
		return uo.(*namespaceNode), true
	}
	return nil, false
}

// enclosingMessage returns the nearest proto.Message containing the specified
// node. This is almost always the parent of the node, but not always.
func (r *Namespace) enclosingMessage(n *namespaceNode) (*namespaceNode, error) {
	if n.isPlaceholder {
		return nil, errors.Errorf("node at %s is a placeholder", n.location)
	}

	if n.isNodeRemoved {
		return nil, errors.Errorf("node at %s not not a part of the current namespace", n.location)
	}

	if n.value.Type().Implements(ProtoMessageType) {
		return n, nil
	}

	p := n.location.Parent()
	for len(p) >= 1 {
		t, ok := r.getNode(p)
		if ok {
			return r.enclosingMessage(t)
		}
		p = p.Parent()
	}
	return nil, errors.Errorf("can't determine top level node for %s", n.location)
}

// asAssetGraph returns a directed graph that represents the dependency
// relationship between known proto.Message objects in the namespace. If such a
// graph cannot be generated, for example due to bad or non-existent nodes,
// returns a nil graph along with a non-nil error indicating what went wrong.
func (r *Namespace) asAssetGraph() (graph.Directed, error) {
	g := simple.NewDirectedGraph()
	var errList []error

	r.ns.Visit(func(l RefPath, o interface{}) bool {
		n := o.(*namespaceNode)

		if n.isNodeRemoved || n.isRuntime {
			return true
		}

		from, err := r.enclosingMessage(n)
		if err != nil {
			errList = AppendErrorList(errList, err)
			return false
		}

		// This condition applies to output messages that are not available yet.
		if from == n && n.isOutput && !n.isValueAvailable {
			return true
		}

		// This step only needs to happen once per message. So we only do it
		// when we are at the top of a message. I.e. when from == n.
		if from == n {
			r.ns.DescendUntil(l, func(l RefPath, o interface{}) bool {
				parent, err := r.enclosingMessage(o.(*namespaceNode))
				if err != nil {
					return true
				}

				if parent != from {
					// Note that we are setting up the edges so that parent
					// depends on children. This should make sense if you think
					// about it for a minute, but basically in order to resolve
					// a higher level object, we need to have resolved its
					// children first. E.g.: in order to resolve a
					// windows_machine, we should already have resolved the
					// network to which its network interface is connected.
					g.SetEdge(g.NewEdge(parent, from))
				}
				return true
			})
		}

		if !g.Has(from) {
			g.AddNode(from)
		}

		for d, _ := range n.dependsOn {
			to, err := r.enclosingMessage(d)
			if err != nil {
				errList = AppendErrorList(errList, err)
				return false
			}

			// if from == to, we are looking at an intra-message dependency.
			// These may be unresolvable if there are cycles therein. However,
			// they are not considered here.
			// TODO(asanka): Check for cycles in intra-asset dependencies.
			if from != to {
				g.SetEdge(g.NewEdge(from, to))
			}
		}

		return true
	})

	if len(errList) != 0 {
		return nil, WrapErrorList(errList)
	}

	return g, nil
}

// asTopologicalList returns the list of referenceNodes in topological order.
// The node at index 𝒊 only depends on nodes at indices 𝒋 where 𝒋 < 𝒊. Returns
// an error if the dependency graph cannot be be topologically sorted, i.e.
// because of a dependency cycle.
func (r *Namespace) asTopologicalList() ([]*namespaceNode, error) {
	g, err := r.asAssetGraph()
	if err != nil {
		return nil, err
	}

	sorted, err := topo.Sort(g)
	if err != nil {
		return nil, err
	}

	// once we get here, it is guaranteed that a node at sorted[i] only depends
	// on sorted[j] where j > i. So we need to visit the nodes in reverse
	// order.
	for i, j := 0, len(sorted)-1; i < j; i, j = i+1, j-1 {
		sorted[i], sorted[j] = sorted[j], sorted[i]
	}

	rv := make([]*namespaceNode, len(sorted))
	for i, s := range sorted {
		rv[i] = s.(*namespaceNode)
	}
	return rv, nil
}

func (r *Namespace) AsSerializedDOTGraph() (string, error) {
	g, err := r.asAssetGraph()
	if err != nil {
		return "", err
	}

	b, err := dot.Marshal(g, "", "", "  ", false)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (r *Namespace) MarshalJSON() ([]byte, error) {
	return r.ns.MarshalJSON()
}

// anneal propagates available values down the dependency graph.
func (r *Namespace) anneal() error {
	candidates := make(namespaceNodeSet)
	r.ns.Visit(func(p RefPath, o interface{}) bool {
		n := o.(*namespaceNode)
		if !n.isValueAvailable || n.isNodeRemoved || n.propagated {
			return true
		}

		if n.value.Kind() != reflect.String {
			return true
		}

		if len(n.dependents) == 0 {
			return true
		}

		candidates[n] = true
		return true
	})

	var errList []error
	for n, _ := range candidates {
		err := n.propagate()
		errList = AppendErrorList(errList, err)
	}
	return WrapErrorList(errList)
}

func (r *Namespace) assign(v *namespaceNode, newValue interface{}) error {
	err := v.assign(newValue)
	if err != nil {
		return err
	}

	if m, ok := newValue.(proto.Message); ok {
		return r.collectFrom(v.location, m)
	}

	return nil
}
