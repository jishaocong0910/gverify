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

// 校验结构体
type checkStruct[V Verifiable] struct {
	vc *VContext
	s  *V
}

// Required 限制不能为nil
func (c *checkStruct[V]) Required(opts ...RuleOption) *checkStruct[V] {
	predicate[int, V](c.vc, c.s, opts, msgBuildFuncRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

// Custom 自定义校验
func (c *checkStruct[V]) Custom(successIfNil bool, custom func(s V) bool, opts ...RuleOption) *checkStruct[V] {
	predicate[int, V](c.vc, c.s, opts, msgBuildFuncDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(*c.s)
	})
	return c
}

// Dive 下沉校验结构体的字段
func (c *checkStruct[V]) Dive() {
	if c.s != nil {
		fi := c.vc.fieldInfo
		c.vc.beforeDiveStruct()
		(*c.s).Checklist(c.vc)
		c.vc.fieldInfo = fi
	}
}
