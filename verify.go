/*
Copyright 2024-present jishaocong0910

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vfy

import (
	"context"
	"reflect"
	"strconv"
)

const (
	SUCCESS = "SUCCESS"
	FAIL    = "FAIL"
)

// Verifiable 待验证的结构体实现此接口，在Checklist方法中编写校验过程
type Verifiable interface {
	Checklist(vc *VContext)
}

// Check 校验入口
func Check[V Verifiable](ctx context.Context, v V, opts ...StructOption) (code string, msg string, msgs []string) {
	if !reflect.Indirect(reflect.ValueOf(v)).IsValid() {
		return
	}

	vc := &VContext{Context: ctx, fieldInfo: &FieldInfo{}}
	for _, o := range opts {
		o(vc)
	}
	v.Checklist(vc)

	if !vc.hasWrong {
		code = SUCCESS
	} else {
		code = vc.code
		if code == "" {
			code = FAIL
		}
	}
	msgs = vc.msgs
	if len(msgs) > 0 {
		msg = msgs[0]
	}
	return
}

// Bool 校验bool类型的字段
func Bool(vc *VContext, b *bool, fieldName string, opts ...FieldOption) *checkBool {
	return &checkBool{vc: vc.beforeCheckField(fieldName, opts), b: b}
}

// Byte 校验byte类型的字段
func Byte(vc *VContext, b *byte, fieldName string, opts ...FieldOption) *checkNumber[byte] {
	return &checkNumber[byte]{vc: vc.beforeCheckField(fieldName, opts), n: b, ntsf: uint8ToStr}
}

// Int 校验int类型的字段
func Int(vc *VContext, i *int, fieldName string, opts ...FieldOption) *checkNumber[int] {
	return &checkNumber[int]{vc: vc.beforeCheckField(fieldName, opts), n: i, ntsf: intToStr}
}

// Int8 校验int8类型的字段
func Int8(vc *VContext, i *int8, fieldName string, opts ...FieldOption) *checkNumber[int8] {
	return &checkNumber[int8]{vc: vc.beforeCheckField(fieldName, opts), n: i, ntsf: int8ToStr}
}

// Int16 校验int16类型的字段
func Int16(vc *VContext, i *int16, fieldName string, opts ...FieldOption) *checkNumber[int16] {
	return &checkNumber[int16]{vc: vc.beforeCheckField(fieldName, opts), n: i, ntsf: int16ToStr}
}

// Int32 校验int32类型的字段
func Int32(vc *VContext, i *int32, fieldName string, opts ...FieldOption) *checkNumber[int32] {
	return &checkNumber[int32]{vc: vc.beforeCheckField(fieldName, opts), n: i, ntsf: int32ToStr}
}

// Int64 校验int64类型的字段
func Int64(vc *VContext, i *int64, fieldName string, opts ...FieldOption) *checkNumber[int64] {
	return &checkNumber[int64]{vc: vc.beforeCheckField(fieldName, opts), n: i, ntsf: int64ToStr}
}

// Uint 校验uint类型的字段
func Uint(vc *VContext, u *uint, fieldName string, opts ...FieldOption) *checkNumber[uint] {
	return &checkNumber[uint]{vc: vc.beforeCheckField(fieldName, opts), n: u, ntsf: uintToStr}
}

// Uint8 校验uint8类型的字段
func Uint8(vc *VContext, u *uint8, fieldName string, opts ...FieldOption) *checkNumber[uint8] {
	return &checkNumber[uint8]{vc: vc.beforeCheckField(fieldName, opts), n: u, ntsf: uint8ToStr}
}

// Uint16 校验uint16类型的字段
func Uint16(vc *VContext, u *uint16, fieldName string, opts ...FieldOption) *checkNumber[uint16] {
	return &checkNumber[uint16]{vc: vc.beforeCheckField(fieldName, opts), n: u, ntsf: uint16ToStr}
}

// Uint32 校验uint32类型的字段
func Uint32(vc *VContext, u *uint32, fieldName string, opts ...FieldOption) *checkNumber[uint32] {
	return &checkNumber[uint32]{vc: vc.beforeCheckField(fieldName, opts), n: u, ntsf: uint32ToStr}
}

// Uint64 校验uint64类型的字段
func Uint64(vc *VContext, u *uint64, fieldName string, opts ...FieldOption) *checkNumber[uint64] {
	return &checkNumber[uint64]{vc: vc.beforeCheckField(fieldName, opts), n: u, ntsf: uint64ToStr}
}

// Float32 校验float32类型的字段
func Float32(vc *VContext, f *float32, fieldName string, opts ...FieldOption) *checkNumber[float32] {
	return &checkNumber[float32]{vc: vc.beforeCheckField(fieldName, opts), n: f, ntsf: float32ToStr}
}

// Float64 校验float64类型的字段
func Float64(vc *VContext, f *float64, fieldName string, opts ...FieldOption) *checkNumber[float64] {
	return &checkNumber[float64]{vc: vc.beforeCheckField(fieldName, opts), n: f, ntsf: float64ToStr}
}

// String 校验字符串类型的字段
func String(vc *VContext, s *string, fieldName string, opts ...FieldOption) *checkString {
	return &checkString{vc: vc.beforeCheckField(fieldName, opts), s: s}
}

// Struct 校验结构体类型的字段（必须实现 [Verifiable] 接口）
func Struct[V Verifiable](vc *VContext, s *V, fieldName string, opts ...FieldOption) *checkStruct[V] {
	return &checkStruct[V]{vc: vc.beforeCheckField(fieldName, opts), s: s}
}

// Embed 校验内嵌结构体
func Embed[V Verifiable](vc *VContext, s *V) *checkEmbed[V] {
	c := &checkEmbed[V]{vc: vc.beforeCheckEmbed(), s: s}
	c.dive()
	return c
}

// Slice 校验切片类型的字段
func Slice[T any](vc *VContext, s []T, fieldName string, opts ...FieldOption) *checkSlice[T] {
	return &checkSlice[T]{vc: vc.beforeCheckField(fieldName, opts), s: s}
}

// Map 校验map类型的字段
func Map[K comparable, V any](vc *VContext, m map[K]V, fieldName string, opts ...FieldOption) *checkMap[K, V] {
	return &checkMap[K, V]{vc: vc.beforeCheckField(fieldName, opts), m: m}
}

// Any 校验任意类型的字段
func Any[T any](vc *VContext, a *T, fieldName string, opts ...FieldOption) *checkAny[T] {
	return &checkAny[T]{vc: vc.beforeCheckField(fieldName, opts), a: a}
}

// 断言
func predicate[C comparable, A any, T *A | []A | map[C]A](vc *VContext, t T, opts []RuleOption, mbf msgBuildFunc, confineFunc func() []string, predicateNil func() bool, predicateNoNil func() bool) {
	if vc.interrupt() {
		return
	}
	for _, opt := range opts {
		opt(vc.fieldInfo)
	}
	var fail bool
	if t == nil {
		if !vc.fieldInfo.omittable {
			fail = !predicateNil()
		}
	} else if predicateNoNil != nil {
		fail = !predicateNoNil()
	}
	if fail {
		var confines []string
		if confineFunc != nil {
			confines = confineFunc()
		}
		vc.fail(mbf, confines)
	}
}

func intToStr(ii ...int) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func int8ToStr(ii ...int8) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func int16ToStr(ii ...int16) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func int32ToStr(ii ...int32) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func int64ToStr(ii ...int64) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(i, 10))
	}
	return result
}

func uintToStr(uu ...uint) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func uint8ToStr(uu ...uint8) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func uint16ToStr(uu ...uint16) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func uint32ToStr(uu ...uint32) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func uint64ToStr(uu ...uint64) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(u, 10))
	}
	return result
}

func float32ToStr(ff ...float32) []string {
	result := make([]string, 0, len(ff))
	for _, f := range ff {
		result = append(result, strconv.FormatFloat(float64(f), 'f', -1, 32))
	}
	return result
}

func float64ToStr(ff ...float64) []string {
	result := make([]string, 0, len(ff))
	for _, f := range ff {
		result = append(result, strconv.FormatFloat(f, 'f', -1, 32))
	}
	return result
}
