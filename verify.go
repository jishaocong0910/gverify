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

func Check[V Verifiable](ctx context.Context, s V, opts ...structOption) (code string, msg string, msgs []string) {
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

func Bool(vc *VContext, b *bool, fieldName string, opts ...fieldOption) *checkBool {
	return &checkBool{vc: vc.beforeCheckField(fieldName, opts), b: b}
}

func Byte(vc *VContext, b *byte, fieldName string, opts ...fieldOption) *checkNumber[byte] {
	return &checkNumber[byte]{vc: vc.beforeCheckField(fieldName, opts), n: b, ntsf: uint8ToStr}
}

func Int(vc *VContext, i *int, fieldName string, opts ...fieldOption) *checkNumber[int] {
	return &checkNumber[int]{vc: vc.beforeCheckField(fieldName, opts), n: i, ntsf: intToStr}
}

func Int8(vc *VContext, i *int8, fieldName string, opts ...fieldOption) *checkNumber[int8] {
	return &checkNumber[int8]{vc: vc.beforeCheckField(fieldName, opts), n: i, ntsf: int8ToStr}
}

func Int16(vc *VContext, i *int16, fieldName string, opts ...fieldOption) *checkNumber[int16] {
	return &checkNumber[int16]{vc: vc.beforeCheckField(fieldName, opts), n: i, ntsf: int16ToStr}
}

func Int32(vc *VContext, i *int32, fieldName string, opts ...fieldOption) *checkNumber[int32] {
	return &checkNumber[int32]{vc: vc.beforeCheckField(fieldName, opts), n: i, ntsf: int32ToStr}
}

func Int64(vc *VContext, i *int64, fieldName string, opts ...fieldOption) *checkNumber[int64] {
	return &checkNumber[int64]{vc: vc.beforeCheckField(fieldName, opts), n: i, ntsf: int64ToStr}
}

func Uint(vc *VContext, u *uint, fieldName string, opts ...fieldOption) *checkNumber[uint] {
	return &checkNumber[uint]{vc: vc.beforeCheckField(fieldName, opts), n: u, ntsf: uintToStr}
}

func Uint8(vc *VContext, u *uint8, fieldName string, opts ...fieldOption) *checkNumber[uint8] {
	return &checkNumber[uint8]{vc: vc.beforeCheckField(fieldName, opts), n: u, ntsf: uint8ToStr}
}

func Uint16(vc *VContext, u *uint16, fieldName string, opts ...fieldOption) *checkNumber[uint16] {
	return &checkNumber[uint16]{vc: vc.beforeCheckField(fieldName, opts), n: u, ntsf: uint16ToStr}
}

func Uint32(vc *VContext, u *uint32, fieldName string, opts ...fieldOption) *checkNumber[uint32] {
	return &checkNumber[uint32]{vc: vc.beforeCheckField(fieldName, opts), n: u, ntsf: uint32ToStr}
}

func Uint64(vc *VContext, u *uint64, fieldName string, opts ...fieldOption) *checkNumber[uint64] {
	return &checkNumber[uint64]{vc: vc.beforeCheckField(fieldName, opts), n: u, ntsf: uint64ToStr}
}

func Float32(vc *VContext, f *float32, fieldName string, opts ...fieldOption) *checkNumber[float32] {
	return &checkNumber[float32]{vc: vc.beforeCheckField(fieldName, opts), n: f, ntsf: float32ToStr}
}

func Float64(vc *VContext, f *float64, fieldName string, opts ...fieldOption) *checkNumber[float64] {
	return &checkNumber[float64]{vc: vc.beforeCheckField(fieldName, opts), n: f, ntsf: float64ToStr}
}

func String(vc *VContext, s *string, fieldName string, opts ...fieldOption) *checkString {
	return &checkString{vc: vc.beforeCheckField(fieldName, opts), s: s}
}

func Struct[V Verifiable](vc *VContext, s *V, fieldName string, opts ...fieldOption) *checkStruct[V] {
	return &checkStruct[V]{vc: vc.beforeCheckField(fieldName, opts), s: s}
}

func Slice[T any](vc *VContext, s []T, fieldName string, opts ...fieldOption) *checkSlice[T] {
	return &checkSlice[T]{vc: vc.beforeCheckField(fieldName, opts), s: s}
}

func Map[K comparable, V any](vc *VContext, m map[K]V, fieldName string, opts ...fieldOption) *checkMap[K, V] {
	return &checkMap[K, V]{vc: vc.beforeCheckField(fieldName, opts), m: m}
}

func Any[T any](vc *VContext, a *T, fieldName string, opts ...fieldOption) *checkAny[T] {
	return &checkAny[T]{vc: vc.beforeCheckField(fieldName, opts), a: a}
}

func checkPredicate[C comparable, A any, T *A | []A | map[C]A](vc *VContext, t T, opts []ruleOption, mbf msgBuildFunc, confineFunc func() []string, predicateNil func() bool, predicate func() bool) {
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
	} else if predicate != nil {
		fail = !predicate()
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
