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

// 校验任意类型
type checkAny[T any] struct {
	vc *VContext
	a  *T
}

// Required 必填
func (c *checkAny[T]) Required(opts ...RuleOption) *checkAny[T] {
	predicate[int, T](c.vc, c.a, opts, msgBuildFuncRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

// Custom 自定义校验
func (c *checkAny[T]) Custom(successIfNil bool, custom func(a T) bool, opts ...RuleOption) *checkAny[T] {
	predicate[int, T](c.vc, c.a, opts, msgBuildFuncDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(*c.a)
	})
	return c
}
