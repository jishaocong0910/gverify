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

type Verifiable interface {
	Checklist(vc *VContext)
}

func Check[V Verifiable](ctx context.Context, s V, opts ...StructOption) (code string, msg string, msgs []string) {
	var target Verifiable
	if rv := reflect.ValueOf(s); rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			target = reflect.New(rv.Type().Elem()).Interface().(Verifiable)
		} else {
			target = s
		}
	} else {
		target = s
	}

	vc := &VContext{Context: ctx, fieldInfo: &FieldInfo{}}
	for _, o := range opts {
		o(vc)
	}
	target.Checklist(vc)

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

func Bool(vc *VContext, b *bool, fieldName string, opts ...FieldOption) *checkBool {
	return &checkBool{vc: vc.beforeCheckField(fieldName, opts), b: b}
}

func Byte(vc *VContext, b *byte, fieldName string, opts ...FieldOption) *checkNumber[byte] {
	return &checkNumber[byte]{vc: vc.beforeCheckField(fieldName, opts), t: b, ntcf: uint8ToConfine}
}

func Int(vc *VContext, i *int, fieldName string, opts ...FieldOption) *checkNumber[int] {
	return &checkNumber[int]{vc: vc.beforeCheckField(fieldName, opts), t: i, ntcf: intToConfine}
}

func Int8(vc *VContext, i *int8, fieldName string, opts ...FieldOption) *checkNumber[int8] {
	return &checkNumber[int8]{vc: vc.beforeCheckField(fieldName, opts), t: i, ntcf: int8ToConfine}
}

func Int16(vc *VContext, i *int16, fieldName string, opts ...FieldOption) *checkNumber[int16] {
	return &checkNumber[int16]{vc: vc.beforeCheckField(fieldName, opts), t: i, ntcf: int16ToConfine}
}

func Int32(vc *VContext, i *int32, fieldName string, opts ...FieldOption) *checkNumber[int32] {
	return &checkNumber[int32]{vc: vc.beforeCheckField(fieldName, opts), t: i, ntcf: int32ToConfine}
}

func Int64(vc *VContext, i *int64, fieldName string, opts ...FieldOption) *checkNumber[int64] {
	return &checkNumber[int64]{vc: vc.beforeCheckField(fieldName, opts), t: i, ntcf: int64ToConfine}
}

func Uint(vc *VContext, u *uint, fieldName string, opts ...FieldOption) *checkNumber[uint] {
	return &checkNumber[uint]{vc: vc.beforeCheckField(fieldName, opts), t: u, ntcf: uintToConfine}
}

func Uint8(vc *VContext, u *uint8, fieldName string, opts ...FieldOption) *checkNumber[uint8] {
	return &checkNumber[uint8]{vc: vc.beforeCheckField(fieldName, opts), t: u, ntcf: uint8ToConfine}
}

func Uint16(vc *VContext, u *uint16, fieldName string, opts ...FieldOption) *checkNumber[uint16] {
	return &checkNumber[uint16]{vc: vc.beforeCheckField(fieldName, opts), t: u, ntcf: uint16ToConfine}
}

func Uint32(vc *VContext, u *uint32, fieldName string, opts ...FieldOption) *checkNumber[uint32] {
	return &checkNumber[uint32]{vc: vc.beforeCheckField(fieldName, opts), t: u, ntcf: uint32ToConfine}
}

func Uint64(vc *VContext, u *uint64, fieldName string, opts ...FieldOption) *checkNumber[uint64] {
	return &checkNumber[uint64]{vc: vc.beforeCheckField(fieldName, opts), t: u, ntcf: uint64ToConfine}
}

func Float32(vc *VContext, f *float32, fieldName string, opts ...FieldOption) *checkNumber[float32] {
	return &checkNumber[float32]{vc: vc.beforeCheckField(fieldName, opts), t: f, ntcf: float32ToConfine}
}

func Float64(vc *VContext, f *float64, fieldName string, opts ...FieldOption) *checkNumber[float64] {
	return &checkNumber[float64]{vc: vc.beforeCheckField(fieldName, opts), t: f, ntcf: float64ToConfine}
}

func String(vc *VContext, s *string, fieldName string, opts ...FieldOption) *checkString {
	return &checkString{vc: vc.beforeCheckField(fieldName, opts), s: s}
}

func Struct[V Verifiable](vc *VContext, s *V, fieldName string, opts ...FieldOption) *checkStruct[V] {
	return &checkStruct[V]{vc: vc.beforeCheckField(fieldName, opts), s: s}
}

func Slice[T any](vc *VContext, s []T, fieldName string, opts ...FieldOption) *checkSlice[T] {
	return &checkSlice[T]{vc: vc.beforeCheckField(fieldName, opts), s: s}
}

func Map[K comparable, V any](vc *VContext, m map[K]V, fieldName string, opts ...FieldOption) *checkMap[K, V] {
	return &checkMap[K, V]{vc: vc.beforeCheckField(fieldName, opts), m: m}
}

func Any[T any](vc *VContext, t *T, fieldName string, opts ...FieldOption) *checkAny[T] {
	return &checkAny[T]{vc: vc.beforeCheckField(fieldName, opts), t: t}
}

func checkRequired[C comparable, A any, T *A | []A | map[C]A](vc *VContext, t T, opts []ItemOption) {
	if vc.interrupt() {
		return
	}
	for _, opt := range opts {
		opt(vc.fieldInfo)
	}
	if t == nil {
		vc.fail(msg_key_required, nil)
	}
}

func checkPredicate[C comparable, A any, T *A | []A | map[C]A](vc *VContext, t T, opts []ItemOption, k msgKey, confineFunc func() []string, predicate func() bool) {
	if vc.interrupt() {
		return
	}
	for _, opt := range opts {
		opt(vc.fieldInfo)
	}
	var fail bool
	if t == nil {
		if !vc.fieldInfo.omittable {
			fail = true
		}
	} else if !predicate() {
		fail = true
	}
	if fail {
		var confine []string
		if confineFunc != nil {
			confine = confineFunc()
		}
		vc.fail(k, confine)
	}
}

func intToConfine(confines ...int) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(int64(confine), 10))
	}
	return result
}

func int8ToConfine(confines ...int8) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(int64(confine), 10))
	}
	return result
}

func int16ToConfine(confines ...int16) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(int64(confine), 10))
	}
	return result
}

func int32ToConfine(confines ...int32) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(int64(confine), 10))
	}
	return result
}

func int64ToConfine(confines ...int64) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(confine, 10))
	}
	return result
}

func uintToConfine(confines ...uint) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func uint8ToConfine(confines ...uint8) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func uint16ToConfine(confines ...uint16) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func uint32ToConfine(confines ...uint32) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func uint64ToConfine(confines ...uint64) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(confine, 10))
	}
	return result
}

func float32ToConfine(confines ...float32) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatFloat(float64(confine), 'f', -1, 32))
	}
	return result
}

func float64ToConfine(confines ...float64) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatFloat(confine, 'f', -1, 32))
	}
	return result
}
