// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"reflect"
	"sync"
)

var gRegistry resolverRegistry

func RegisterResolverClass(resolver interface{}) {
	err := gRegistry.addWithConversion(resolver)
	if err != nil {
		panic(err)
	}
}

func RegisterResolverFunc(kind ResolverKind, function interface{}) {
	err := gRegistry.addResolverFunction(kind, reflect.ValueOf(function))
	if err != nil {
		panic(err)
	}
}

func getResolvers(target interface{}, kind ResolverKind) ([]resolverFunc, error) {
	return gRegistry.get(target, kind)
}

var ErrorType = reflect.TypeOf((*error)(nil)).Elem()

var ResolverNotFoundError = errors.New("resolver not found")

type resolverRegistry struct {
	m map[resolverKey][]resolverFunc
}

type resolverKey struct {
	targetType   reflect.Type
	resolverType reflect.Type
}

func (r *resolverRegistry) addWithConversion(resolver interface{}) error {
	rv := reflect.ValueOf(resolver)
	rt := rv.Type()
	if rt.Kind() != reflect.Ptr || rt.NumMethod() == 0 {
		return errors.New(`invalid resolver. See documentation in resolvers.go for more details on kinds of acceptable resolvers`)
	}

	foundAny := false
	for i := 0; i < rt.NumMethod(); i++ {
		mt := rt.Method(i)
		resolverType, ok := allResolverTypes[mt.Name]
		if !ok {
			continue
		}

		boundFunc := rv.Method(i)
		err := r.addResolverFunction(resolverType, boundFunc)
		if err != nil {
			return err
		}
		foundAny = true
	}

	if !foundAny {
		return errors.New("none of the methods of the resolver type could be recognized")
	}
	return nil
}

func (r *resolverRegistry) add(target reflect.Type, kind ResolverKind, resolver resolverFunc) error {
	if r.m == nil {
		r.m = make(map[resolverKey][]resolverFunc)
	}

	rk := resolverKey{targetType: target, resolverType: kind}
	rl, _ := r.m[rk]
	rl = append(rl, resolver)
	r.m[rk] = rl
	return nil
}

func (r *resolverRegistry) get(target interface{}, kind ResolverKind) ([]resolverFunc, error) {
	tt := reflect.TypeOf(target)
	if err := isValidTargetType(tt); err != nil {
		return nil, err
	}

	rk := resolverKey{targetType: tt, resolverType: kind}
	l, ok := r.m[rk]
	if !ok || len(l) == 0 {
		return nil, ResolverNotFoundError
	}
	return l, nil
}

func (r *resolverRegistry) addResolverFunction(kind ResolverKind, function reflect.Value) error {
	fType := function.Type()
	if fType.NumIn() != 2 {
		return errors.New("incorrect number of arguments for resolver function")
	}
	targetType := fType.In(1)

	if !ContextType.AssignableTo(fType.In(0)) || !targetType.Implements(ProtoMessageType) ||
		fType.NumOut() != 1 || !fType.Out(0).Implements(ErrorType) {
		return errors.New("invalid resolver function")
	}

	return r.add(targetType, kind, wrapGenericResolverFunction(function, targetType))
}

// isValidTargetType returns true if `tt` represents a valid resolvable type.
// Currently only message.Proto substrates are considered resolvable, and only
// if the passed in type is a pointer type which points to a structure.
func isValidTargetType(tt reflect.Type) error {
	if !tt.Implements(ProtoMessageType) {
		return errors.Errorf("target type %#v does not implement proto.Message", tt)
	}

	if tt.Kind() != reflect.Ptr || tt.Elem().Kind() != reflect.Struct {
		return errors.Errorf(`target type is not a pointer to a struct. Got %#v.

It's not safe to use an interface type as the target since another message type
with identical field names could satisfy the target type. If so, the resolver
may get applied to unexpected message types.
`, tt)
	}

	return nil
}

func wrapGenericResolverFunction(function reflect.Value, targetType reflect.Type) resolverFunc {
	return func(ctx Context, m proto.Message) error {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					panic(errors.Wrapf(err, "wrapping resolving %+v for type %v", function, targetType))
				} else {
					panic(errors.Errorf("wrapping resolver %+v for type %v", function, targetType))
				}
			}
		}()

		ctxValue := reflect.ValueOf(ctx)
		targetValue := reflect.ValueOf(m)

		if targetType != ProtoMessageType {
			targetValue = targetValue.Convert(targetType)
		}

		results := function.Call([]reflect.Value{ctxValue, targetValue})
		if results[0].IsNil() {
			return nil
		}
		return results[0].Interface().(error)
	}
}

type resolverExtractor func(interface{}) resolverFunc

var extractorCache sync.Map

func getResolverExtractor(t ResolverKind) resolverExtractor {
	if existing, ok := extractorCache.Load(t); ok {
		return existing.(resolverExtractor)
	}

	kind := t.(reflect.Type)
	if kind.Kind() != reflect.Interface || kind.NumMethod() != 1 {
		panic("GetResolverFunc called on a type that is not an interface with a single method.")
	}
	m := kind.Method(0)
	e := func(o interface{}) resolverFunc {
		v := reflect.ValueOf(o)
		vm := v.MethodByName(m.Name)
		return wrapGenericResolverFunction(vm, kind)
	}
	extractorCache.Store(t, e)
	return e
}
