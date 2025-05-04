package vfy

func NewDefaultContext() *VContext {
	return &VContext{all: false, fieldInfo: &FieldInfo{}}
}

func GetResult(vc *VContext) (code string, msg string, msgs []string) {
	if !vc.hasWrong {
		code = SUCCESS
	} else if vc.code == "" {
		code = FAIL
	} else {
		code = vc.code
	}
	msgs = vc.msgs
	if len(msgs) > 0 {
		msg = msgs[0]
	}
	return
}

func SetHasWrong(vc *VContext) {
	vc.hasWrong = true
}

func SetAll(vc *VContext) {
	vc.all = true
}

func SetFieldName(vc *VContext, fieldName string) {
	vc.fieldInfo.fieldName = fieldName
}

func SetOmittable(vc *VContext) {
	Omittable()(vc.fieldInfo)
}

func CheckRequired[T any](vc *VContext, t *T, opts []ItemOption) {
	checkRequired[int, T](vc, t, opts)
}

func CheckPredicate[T any](vc *VContext, t *T, opts []ItemOption, k msgKey, confineFunc func() []string, predicate func() bool) {
	checkPredicate[int, T](vc, t, opts, k, confineFunc, predicate)
}

type Number = number

func IntToConfine(confines ...int) []string {
	return intToConfine(confines...)
}

func Int8ToConfine(confines ...int8) []string {
	return int8ToConfine(confines...)
}

func Int16ToConfine(confines ...int16) []string {
	return int16ToConfine(confines...)
}

func Int32ToConfine(confines ...int32) []string {
	return int32ToConfine(confines...)
}

func Int64ToConfine(confines ...int64) []string {
	return int64ToConfine(confines...)
}

func UintToConfine(confines ...uint) []string {
	return uintToConfine(confines...)
}

func Uint8ToConfine(confines ...uint8) []string {
	return uint8ToConfine(confines...)
}

func Uint16ToConfine(confines ...uint16) []string {
	return uint16ToConfine(confines...)
}

func Uint32ToConfine(confines ...uint32) []string {
	return uint32ToConfine(confines...)
}

func Uint64ToConfine(confines ...uint64) []string {
	return uint64ToConfine(confines...)
}

func Float32ToConfine(confines ...float32) []string {
	return float32ToConfine(confines...)
}

func Float64ToConfine(confines ...float64) []string {
	return float64ToConfine(confines...)
}
