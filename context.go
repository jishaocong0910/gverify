package vfy

import (
	"context"
	"fmt"
	"strings"
)

type VContext struct {
	context.Context
	code      string
	all       bool
	hasWrong  bool
	msgs      []string
	fieldInfo *FieldInfo
}

func (vc *VContext) interrupt() bool {
	return vc.hasWrong && !vc.all
}

func (vc *VContext) beforeDiveStruct() {
	vc.fieldInfo = &FieldInfo{
		fieldNamePrefix: vc.fieldInfo.fieldName + ".",
	}
}

func (vc *VContext) beforeDiveSliceMap(elemName string) {
	vc.fieldInfo = &FieldInfo{
		fieldName:   vc.fieldInfo.fieldName + elemName,
		hasElemName: true,
	}
}

func (vc *VContext) beforeCheckField(fieldName string, opts []fieldOption) *VContext {
	if vc.fieldInfo.hasElemName {
		vc.fieldInfo = &FieldInfo{
			fieldName:   vc.fieldInfo.fieldName,
			hasElemName: true,
		}
	} else {
		vc.fieldInfo = &FieldInfo{
			fieldNamePrefix: vc.fieldInfo.fieldNamePrefix,
			fieldName:       vc.fieldInfo.fieldNamePrefix + fieldName,
		}
	}
	for _, o := range opts {
		o(vc.fieldInfo)
	}
	return vc
}

func (vc *VContext) beforeCheckEmbed() *VContext {
	vc.fieldInfo = &FieldInfo{
		fieldNamePrefix: vc.fieldInfo.fieldNamePrefix,
	}
	return vc
}

func (vc *VContext) fail(mbf msgBuildFunc, confines []string) {
	vc.hasWrong = true
	if !vc.all {
		vc.code = vc.fieldInfo.code
	}
	vc.fieldInfo.confines = confines
	var msg string
	if vc.fieldInfo.mbf == nil {
		if mbf != nil {
			mbf(vc.fieldInfo)
		}
	} else {
		vc.fieldInfo.mbf(vc.fieldInfo)
	}
	msg = vc.fieldInfo.msg
	vc.msgs = append(vc.msgs, msg)
}

type FieldInfo struct {
	fieldNamePrefix string
	hasElemName     bool
	omittable       bool
	fieldName       string
	confines        []string
	mbf             msgBuildFunc
	code            string
	msg             string
}

func (f *FieldInfo) FieldName() string {
	return f.fieldName
}

func (f *FieldInfo) Confine(i int) string {
	confine := ""
	if i < len(f.confines) {
		confine = f.confines[i]
	}
	return confine
}

func (f *FieldInfo) Confines() string {
	var builder strings.Builder
	l := len(f.confines)
	if l > 0 {
		builder.WriteString(f.confines[0])
	}
	for i := 1; i < l-1; i++ {
		builder.WriteString(", ")
		builder.WriteString(f.confines[i])
	}
	if l > 1 {
		builder.WriteString(" or ")
		builder.WriteString(f.confines[l-1])
	}
	return builder.String()
}

func (f *FieldInfo) Msg(msg string, args ...any) {
	f.msg = fmt.Sprintf(msg, args...)
}
