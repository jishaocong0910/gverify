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
	"fmt"
	"strings"
)

// VContext 验证过程的上下文
type VContext struct {
	context.Context
	// 错误码，一般情况只有“SUCCESS”和"FALSE"，使用自定义错误码选项 [Code] 时会有其他值。
	code string
	// 是否校验所有规则，使用选项 [All] 时为true
	all bool
	// 是否已有校验错误，在不校验全部规则时，通过此此值判断是否停止校验
	hasWrong bool
	// 错误消息，一般情况下只可能有一条，使用选项 [All] 时会有多条
	msgs []string
	// 当前校验字段的信息
	fieldInfo *FieldInfo
}

// 是否打断校验
func (vc *VContext) interrupt() bool {
	return vc.hasWrong && !vc.all
}

// 校验结构体（中的字段）的前置逻辑
func (vc *VContext) beforeDiveStruct() {
	vc.fieldInfo = &FieldInfo{
		fieldNamePrefix: vc.fieldInfo.fieldName + ".",
	}
}

// 校验切片元素的前置逻辑
func (vc *VContext) beforeDiveSliceMap(elemName string) {
	vc.fieldInfo = &FieldInfo{
		fieldName:   vc.fieldInfo.fieldName + elemName,
		hasElemName: true,
	}
}

// 校验字段的前置逻辑
func (vc *VContext) beforeCheckField(fieldName string, opts []FieldOption) *VContext {
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

// 校验内嵌结构体（中的字段）前置逻辑
func (vc *VContext) beforeCheckEmbed() *VContext {
	vc.fieldInfo = &FieldInfo{
		fieldNamePrefix: vc.fieldInfo.fieldNamePrefix,
	}
	return vc
}

// 校验失败处理
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

// FieldInfo 字段信息
type FieldInfo struct {
	// 字段名前缀
	fieldNamePrefix string
	// 是否为切片元素、map的key或value
	hasElemName bool
	// nil时忽略校验
	omitempty bool
	// 修字段值函数
	amend func(a any) any
	// 字段名
	fieldName string
	// 限制值的字符串形式，校验每个规则方法时会修改此字段
	confines []string
	// 错误消息函数，校验每个规则方法时会修改此字段
	mbf msgBuildFunc
	// 校验时的错误码，校验每个规则方法时会修改此字段
	code string
	// 校验失败的错误消息，校验每个规则方法时会修改此字段
	msg string
}

// FieldName 返回具有层级关系的字段名称。例如：title、author.name、categoryId[2].sort。
func (f *FieldInfo) FieldName() string {
	return f.fieldName
}

// Confine 返回规则方法的指定索引的限制值的字符串形式。例如：对于Max(10)，vc.Confine(0)返回10；对于Range(5, 15)，vc.Confine(0)返回5，vc.Confine(1)返回15。
func (f *FieldInfo) Confine(i int) string {
	confine := ""
	if i < len(f.confines) {
		confine = f.confines[i]
	}
	return confine
}

// Confines 返回规则方法的所有限制值的字符串形式，用","拼接，若数量超过两个，则最后一个用"or"拼接。例如：对于Enum("zh-cn", "en-US")，返回"zh-cn, en-US"；对于Enum("zh-cn", "en-US", "ja-JP")，返回"zh-cn, en-US or ja-JP"。
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

// Msg 自定义错误消息
func (f *FieldInfo) Msg(msg string, args ...any) {
	f.msg = fmt.Sprintf(msg, args...)
}
