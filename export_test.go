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

func CheckPredicate[T any](vc *VContext, t *T, opts []checkOption, confineFunc func() []string, predicateNil func() bool, predicate func() bool) {
	checkPredicate[int, T](vc, t, opts, msgBuildFuncDefault, confineFunc, predicateNil, predicate)
}

type Number = number

type CheckOption = checkOption

func IntToConfine(confines ...int) []string {
	return intToStr(confines...)
}

func Int8ToConfine(confines ...int8) []string {
	return int8ToStr(confines...)
}

func Int16ToConfine(confines ...int16) []string {
	return int16ToStr(confines...)
}

func Int32ToConfine(confines ...int32) []string {
	return int32ToStr(confines...)
}

func Int64ToConfine(confines ...int64) []string {
	return int64ToStr(confines...)
}

func UintToConfine(confines ...uint) []string {
	return uintToStr(confines...)
}

func Uint8ToConfine(confines ...uint8) []string {
	return uint8ToStr(confines...)
}

func Uint16ToConfine(confines ...uint16) []string {
	return uint16ToStr(confines...)
}

func Uint32ToConfine(confines ...uint32) []string {
	return uint32ToStr(confines...)
}

func Uint64ToConfine(confines ...uint64) []string {
	return uint64ToStr(confines...)
}

func Float32ToConfine(confines ...float32) []string {
	return float32ToStr(confines...)
}

func Float64ToConfine(confines ...float64) []string {
	return float64ToStr(confines...)
}
