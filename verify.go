// Copyright 2024 jishaocong0910/@163.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vfy

import (
	"context"
	"reflect"
)

const (
	SUCCESS = "SUCCESS"
	ERROR   = "ERROR"
)

type Verifiable interface {
	Checklist(ctx *Context)
}

func Check[V Verifiable](ctx context.Context, v V) (code string, msg string) {
	code, msg, _ = Check_(ctx, v, false)
	return
}

func Check_[V Verifiable](ctx context.Context, v V, all bool) (code string, first string, msgs []string) {
	var target Verifiable
	if rv := reflect.ValueOf(v); rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			target = reflect.New(rv.Type().Elem()).Interface().(Verifiable)
		} else {
			target = v
		}
	} else {
		target = v
	}

	c := &Context{Context: ctx, all: all}
	target.Checklist(c)

	if len(c.msgs) > 0 {
		first = c.msgs[0]
	}
	if !c.wronged {
		c.code = SUCCESS
	} else if c.code == "" {
		c.code = ERROR
	}

	return c.code, first, c.msgs
}

func Bool(ctx *Context, b *bool, fieldName string) *checkBool {
	return &checkBool{ctx: ctx.reset(fieldName), b: b}
}

func Byte(ctx *Context, b *byte, fieldName string) *checkByte {
	return &checkByte{ctx: ctx.reset(fieldName), b: b}
}

func Int(ctx *Context, i *int, fieldName string) *checkInt {
	return &checkInt{ctx: ctx.reset(fieldName), i: i}
}

func Int8(ctx *Context, i *int8, fieldName string) *checkInt8 {
	return &checkInt8{ctx: ctx.reset(fieldName), i: i}
}

func Int16(ctx *Context, i *int16, fieldName string) *checkInt16 {
	return &checkInt16{ctx: ctx.reset(fieldName), i: i}
}

func Int32(ctx *Context, i *int32, fieldName string) *checkInt32 {
	return &checkInt32{ctx: ctx.reset(fieldName), i: i}
}

func Int64(ctx *Context, i *int64, fieldName string) *checkInt64 {
	return &checkInt64{ctx: ctx.reset(fieldName), i: i}
}

func Uint(ctx *Context, u *uint, fieldName string) *checkUint {
	return &checkUint{ctx: ctx.reset(fieldName), u: u}
}

func Uint8(ctx *Context, u *uint8, fieldName string) *checkUint8 {
	return &checkUint8{ctx: ctx.reset(fieldName), u: u}
}

func Uint16(ctx *Context, u *uint16, fieldName string) *checkUint16 {
	return &checkUint16{ctx: ctx.reset(fieldName), u: u}
}

func Uint32(ctx *Context, u *uint32, fieldName string) *checkUint32 {
	return &checkUint32{ctx: ctx.reset(fieldName), u: u}
}

func Uint64(ctx *Context, u *uint64, fieldName string) *checkUint64 {
	return &checkUint64{ctx: ctx.reset(fieldName), u: u}
}

func Float32(ctx *Context, f *float32, fieldName string) *checkFloat32 {
	return &checkFloat32{ctx: ctx.reset(fieldName), f: f}
}

func Float64(ctx *Context, f *float64, fieldName string) *checkFloat64 {
	return &checkFloat64{ctx: ctx.reset(fieldName), f: f}
}

func String(ctx *Context, s *string, fieldName string) *checkString {
	return &checkString{ctx: ctx.reset(fieldName), s: s}
}

func Struct[V Verifiable](ctx *Context, v *V, fieldName string) *checkStruct[V] {
	return &checkStruct[V]{ctx: ctx.reset(fieldName), v: v}
}

func Slices[T any](ctx *Context, s []T, fieldName string) *checkSlices[T] {
	return &checkSlices[T]{ctx: ctx.reset(fieldName), s: s}
}

func Map[K comparable, V any](ctx *Context, m map[K]V, fieldName string) *checkMap[K, V] {
	return &checkMap[K, V]{ctx: ctx.reset(fieldName), m: m}
}

func Any[T any](ctx *Context, t *T, fieldName string) *checkAny[T] {
	return &checkAny[T]{ctx: ctx.reset(fieldName), t: t}
}
