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

// StructOption 校验结构体的选项
type StructOption func(vc *VContext)

// FieldOption 校验字段的选项
type FieldOption func(*FieldInfo)

// RuleOption 校验规则方法的选项
type RuleOption func(*FieldInfo)

// All 是否校验所有规则方法（将输出所有错误消息）
func All() StructOption {
	return func(vc *VContext) {
		vc.all = true
	}
}

// Omitempty 字段值为nil时不校验
func Omitempty() FieldOption {
	return func(o *FieldInfo) {
		o.omitempty = true
	}
}

// Amend 修改字段值，会影响原结构体。
//
// 无效场景：
//  1. 字段值为nil
//  3. 字段类型函数 [Slice]、[Map]、[Embed]
//  2. 泛型与字段类型不一致
func Amend[T any](amend func(t T) T) FieldOption {
	return func(o *FieldInfo) {
		o.amend = func(a any) any {
			if t, ok := a.(T); ok {
				return amend(t)
			}
			return a
		}
	}
}

// Code 自定义错误码，只有不使用 [All] 选项时有效果。
func Code(code string) RuleOption {
	return func(o *FieldInfo) {
		o.code = code
	}
}

// Msg 自定义错误消息
func Msg(mbf msgBuildFunc) RuleOption {
	return func(o *FieldInfo) {
		o.mbf = mbf
	}
}
