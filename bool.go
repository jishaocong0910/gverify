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

// 校验bool类型
type checkBool struct {
	vc *VContext
	b  *bool
}

// Required 必填
func (c *checkBool) Required(opts ...RuleOption) *checkBool {
	predicate[int, bool](c.vc, c.b, opts, msgBuildFuncRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

// Custom 自定义校验
func (c *checkBool) Custom(successIfNil bool, custom func(b bool) bool, opts ...RuleOption) *checkBool {
	predicate[int, bool](c.vc, c.b, opts, msgBuildFuncDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(*c.b)
	})
	return c
}
