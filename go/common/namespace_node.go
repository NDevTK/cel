// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"encoding/json"
	"reflect"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

// namespaceNode is a node in a reference graph.
//
// Each node can represent a proto.Message or a field thereof. Each addressable
// object or field in a namespace will have a corresponding namespaceNode
// object. Within each namespace, there's a 1:1 correspondence between a
// location in the namespace and the namespaceNode corresponding to that path.
// I.e. once a location is found to map to a namespaceNode, that location will
// always map to the same namespaceNode as long as the enclosing References
// object is alive. This holds true even if the corresponding value changes.
//
// namespaceNode objects must be uniquely owned by Namespace objects.
type namespaceNode struct {
	// Where is it?
	location RefPath

	// true if this node is a placeholder that was created to represent a
	// dependency. Such a node will be elevated to the status of concrete value
	// if a proto.Message is seen with a field that maps to this `location`.
	isPlaceholder bool

	// Value contained at this node in the namespace.
	value reflect.Value

	// if true, the value is the output of an external operation. The value
	// should be supplied via Publish() on the References object containing
	// this refValue.
	isOutput bool

	// if true, the value is only known at runtime. These values are considered
	// unavailable during the deployment stage.
	isRuntime bool

	// if true, this node represents a top level collection.
	isTopLevelCollection bool

	// if true, the value at this node is currently available. Implies that
	// isValueAvailable is true for all elemments in dependsOn.
	isValueAvailable bool

	// if true, this node should be considered removed from the namespace. Note
	// that the node itself isn't removed since it shouldn't be possible to
	// rebind this location to another node. That would violate the requirement
	// that each location in a namespace is bound to a unique refNode.
	isNodeRemoved bool

	// if true, the value has been propagated to all its dependents. Note that
	// any change to the value (communicated via assing()), or any change to
	// the list of dependents (communicated via addDependent()) should result
	// in this field being reset to false.
	propagated bool

	// If this value is a reference to another value, then this field will
	// contain the refpath that should be the root of the reference. I.e.: for
	// a field that is a referenence to a WindowsMachine, the referenceRoot
	// would be ["asset", "windows_machine"].
	referenceRoot RefPath

	// set of nodes that this node depends on.
	dependsOn namespaceNodeSet

	// set of nodes that depend on this node.
	dependents namespaceNodeSet

	// ID is unique within its parent (i.e. a References object), and thus
	// should be assigned by the parent. This is used for implementing
	// gonum.org/v1/gonum/graph#Node
	id int64
}

// namespaceNodeSet contains a set of nodes. The key is a pointer to a the
// object.
type namespaceNodeSet map[*namespaceNode]bool

// bind binds this namespaceNode to the specified value. For all practical
// purposes, a node can only be bound once.
func (v *namespaceNode) bind(rv reflect.Value, validation *Validation) error {
	if v.isNodeRemoved {
		return errors.Errorf("attempt to bind a value to a deleted location at %s", v.location)
	}

	if !rv.IsValid() {
		return errors.Errorf("attempting to bind node at %s to an invalid value", v.location)
	}

	if v.value.IsValid() && v.value == rv {
		// already bound to rv.
		return nil
	}

	if v.value.IsValid() && v.value.CanInterface() && rv.CanInterface() && v.value.Interface() == rv.Interface() {
		return nil
	}

	// The only two cases where we allow bind() is if:
	// - the value is a placeholder (i.e. we haven't bound anything to this node yet).
	// - the value is an OUTPUT that isn't available.
	if !v.isPlaceholder && (!v.isOutput || v.isValueAvailable) {
		return errors.Errorf("attempt to rebind value at %s to a new value %#v (old value %#v)",
			v.location.String(), rv, v.value)
	}

	v.value = rv
	v.isPlaceholder = false
	if validation != nil {
		v.isOutput = validation.IsOutput()
		v.isRuntime = validation.IsRuntime()
		v.isTopLevelCollection = validation.IsTopLevelCollection()
	}

	wasAvailable := v.isValueAvailable

	v.isValueAvailable = !v.isOutput

	// Output values that have a non-zero or non-empty value should be
	// considered available on bind.
	if v.isOutput && v.value.IsValid() {
		switch v.value.Kind() {
		case reflect.Interface, reflect.Ptr, reflect.UnsafePointer:
			v.isValueAvailable = !v.value.IsNil()

		case reflect.Chan, reflect.Func:
			// Do nothing

		case reflect.Array, reflect.Slice, reflect.String, reflect.Map:
			v.isValueAvailable = (v.value.Len() != 0)

		case reflect.Struct:
			v.isValueAvailable = true

		default:
			v.isValueAvailable = (v.value.Interface() != reflect.Zero(v.value.Type()).Interface())
		}
	}

	if !wasAvailable && v.isValueAvailable {
		return v.propagate()
	}

	return nil
}

// addParent adds a node that which node depends on. This is the equivalent of
// adding a dependant to the parent node via addDependent().
func (v *namespaceNode) addParent(o *namespaceNode) error {
	return o.addDependent(v)
}

// addDependent adds this node as a dependent of another node.
func (v *namespaceNode) addDependent(o *namespaceNode) error {
	if v.isNodeRemoved || o.isNodeRemoved {
		return errors.Errorf("dependency %s -> %s is invalid. One or both endpoints were removed from the namespace",
			v.location, o.location)
	}

	if v.dependents == nil {
		v.dependents = make(namespaceNodeSet)
	}
	v.dependents[o] = true
	v.propagated = false

	if o.dependsOn == nil {
		o.dependsOn = make(namespaceNodeSet)
	}
	o.dependsOn[v] = true
	return nil
}

// assign a value to this node. Note that the value must already be bound to an
// object or a field (i.e. a value). The bound value must also support setting
// (i.e. no R-values).
//
// If assigning the value causes the node to transition from !isValueAvailable
// to isValueAvailable, then the value of this node is propagated to all its
// children.
//
// assign() can only be used on OUTPUT type fields. All other types are
// considered read-only at deployment time.
func (v *namespaceNode) assign(i interface{}) (err error) {
	if v.isNodeRemoved {
		return errors.Errorf("attempting to assign a value to property at %s which was removed from the namespace",
			v.location)
	}

	if !v.isOutput {
		return errors.Errorf("attempting at assign a value to property at %s which is not an OUTPUT."+
			" Use the Publish methods from assocated Namespace object to assign a value to an OUTPUT field.", v.location)
	}

	if v.isValueAvailable {
		return errors.Errorf("attempting to assign a value to property at %s which already has a value", v.location)
	}

	if s, ok := i.(string); ok && ContainsInlineReferences(s) {
		return errors.Errorf("the value published at %s contains new references. Value is \"%s\"", v.location, s)
	}

	if !v.value.CanSet() {
		return errors.Errorf("the property at %s cannot be modified. Value is %#v", v.location, v.value)
	}

	// v.value.Set() will panic() if something is wrong. Thanks golang for
	// having two ways of indicating failure.
	defer func() {
		if e := recover(); e != nil {
			switch c := e.(type) {
			case error:
				err = errors.Wrapf(c, "while assigning value to %s. Value is %#v", v.location, i)
			case string:
				err = errors.Errorf("%s: while assigning value to %s. Value is %#v", c, v.location, i)
			default:
				err = errors.Errorf("failed to assign value to %s. Value is %#v", v.location, i)
			}
		}
	}()
	v.value.Set(reflect.ValueOf(i))
	v.isValueAvailable = true
	return v.propagate()
}

// propagate visits all the dependents of this node and invokes interpolate()
// on them. This should be done after the value of this field changes so that
// dependent nodes are incrementally resolved.
func (v *namespaceNode) propagate() error {
	if v.propagated || v.value.Kind() != reflect.String {
		return nil
	}

	if !v.isValueAvailable {
		return errors.Errorf("propagate() called before value is available. Path: %s",
			v.location)
	}

	value := v.value.String()
	var errList []error
	for d, _ := range v.dependents {
		errList = AppendErrorList(errList, d.interpolate(v.location, value))
	}
	if len(errList) == 0 {
		v.propagated = true
	}
	return WrapErrorList(errList)
}

// interpolate uses the string value |s| at location |l| to expand inline
// references at this node.
//
// It is not an error to call this function on a node that is not a pending
// string value. If not applicable, the function will simply return nil.
//
// If the interpolation results in the value being well formed and available
// (i.e. no more inline references exist), then the value of this node will be
// propagated to any dependents.
func (v *namespaceNode) interpolate(l RefPath, s string) error {
	if v.isPlaceholder || v.isValueAvailable || v.value.Kind() != reflect.String {
		// nothing to do.
		return nil
	}

	target := v.value.String()
	target, err := expandSingleReference(target, l, s)
	if err != nil {
		return errors.Wrapf(err, "while expanding references in string at %s", v.location)
	}

	v.value.SetString(target)
	v.isValueAvailable = !ContainsInlineReferences(target)

	// v is newly avilable. It's value should be propagated to its dependents.
	if v.isValueAvailable {
		return v.propagate()
	}

	return nil
}

// ID returns an integer ID that's unique within its parent. Implements
// github.com/gonum/graph#Node.
func (v *namespaceNode) ID() int64 {
	return v.id
}

// DOTID returns a label for use when serializing the reference graph as a DOT
// graph.
func (v *namespaceNode) DOTID() string {
	return "\"" + v.location.String() + "\""
}

// MarshalJSON correctly marshals the node contents into JSON. Unsatisfyingly,
// just calling json.Marshal() on proto.Message doesn't always work correctly.
// See documentation for jsonpb for details.
func (v *namespaceNode) MarshalJSON() ([]byte, error) {
	if v.isValueAvailable && v.value.CanInterface() {
		i := v.value.Interface()
		if p, ok := i.(proto.Message); ok {
			m := jsonpb.Marshaler{OrigName: true}
			s, err := m.MarshalToString(p)
			if err != nil {
				return nil, err
			}
			return []byte(s), nil
		}
		return json.Marshal(v.value.Interface())
	}
	return nil, errors.Errorf("value at %s is not available", v.location)
}

// nodeSetToRefPathList takes a set of namespace nodes and returns the
// corresponding list of RefPath. The paths are in no particular order.
func nodeSetToRefPathList(v namespaceNodeSet) []RefPath {
	var o []RefPath
	for iv, _ := range v {
		o = append(o, iv.location)
	}
	return o
}
