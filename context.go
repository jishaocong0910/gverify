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

func (c *VContext) interrupt() bool {
	return c.hasWrong && !c.all
}

func (c *VContext) diveStruct() {
	c.fieldInfo = &FieldInfo{
		fieldNamePrefix: c.fieldInfo.fieldName + ".",
	}
}

func (c *VContext) diveSliceMap() {
	c.fieldInfo = &FieldInfo{
		fieldNamePrefix: c.fieldInfo.fieldName,
	}
}

func (c *VContext) beforeCheckElem(elemName string) {
	c.fieldInfo = &FieldInfo{
		fieldNamePrefix: c.fieldInfo.fieldNamePrefix,
		omittable:       c.fieldInfo.omittable,
		elemName:        elemName,
	}
}

func (c *VContext) beforeCheckField(fieldName string, opts []fieldOption) *VContext {
	if c.fieldInfo.elemName != "" {
		fieldName = c.fieldInfo.elemName
	}
	c.fieldInfo = &FieldInfo{
		fieldNamePrefix: c.fieldInfo.fieldNamePrefix,
		omittable:       c.fieldInfo.omittable,
	}
	c.fieldInfo.fieldName = c.fieldInfo.fieldNamePrefix + fieldName
	for _, o := range opts {
		o(c.fieldInfo)
	}
	return c
}

func (c *VContext) fail(mbf msgBuildFunc, confines []string) {
	c.hasWrong = true
	if !c.all {
		c.code = c.fieldInfo.code
	}
	c.fieldInfo.confines = confines
	var msg string
	if c.fieldInfo.mbf == nil {
		if mbf != nil {
			mbf(c.fieldInfo)
		}
	} else {
		c.fieldInfo.mbf(c.fieldInfo)
	}
	msg = c.fieldInfo.msg
	c.msgs = append(c.msgs, msg)
}

func (c *VContext) copyFieldInfo() *FieldInfo {
	f := *c.fieldInfo
	f.mbf = nil
	f.confines = nil
	return &f
}

type FieldInfo struct {
	fieldNamePrefix string
	omittable       bool
	fieldName       string
	code            string
	mbf             msgBuildFunc
	elemName        string
	confines        []string
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
