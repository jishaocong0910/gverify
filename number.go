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

type number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type numberToStrFunc[T number] func(t ...T) []string

// 校验数字类型
type checkNumber[T number] struct {
	vc   *VContext
	n    *T
	ntsf numberToStrFunc[T]
}

// Required 限制不能为nil
func (c *checkNumber[T]) Required(opts ...RuleOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

// Min 限制最小值
func (c *checkNumber[T]) Min(min T, opts ...RuleOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncMin, func() []string {
		return c.ntsf(min)
	}, func() bool {
		return false
	}, func() bool {
		return *c.n >= min
	})
	return c
}

// Max 限制最大值
func (c *checkNumber[T]) Max(max T, opts ...RuleOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncMax, func() []string {
		return c.ntsf(max)
	}, func() bool {
		return true
	}, func() bool {
		return *c.n <= max
	})
	return c
}

// Range 限制值范围（包含边界）
func (c *checkNumber[T]) Range(min, max T, opts ...RuleOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncRange, func() []string {
		return c.ntsf(min, max)
	}, func() bool {
		return false
	}, func() bool {
		return *c.n >= min && *c.n <= max
	})
	return c
}

// Gt 限制必须大于指定值
func (c *checkNumber[T]) Gt(min T, opts ...RuleOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncGt, func() []string {
		return c.ntsf(min)
	}, func() bool {
		return false
	}, func() bool {
		return *c.n > min
	})
	return c
}

// Lt 限制必须小于指定值
func (c *checkNumber[T]) Lt(max T, opts ...RuleOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncLt, func() []string {
		return c.ntsf(max)
	}, func() bool {
		return true
	}, func() bool {
		return *c.n < max
	})
	return c
}

// Within 限制值范围（不含边界）
func (c *checkNumber[T]) Within(min, max T, opts ...RuleOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncWithin, func() []string {
		return c.ntsf(min, max)
	}, func() bool {
		return false
	}, func() bool {
		return *c.n > min && *c.n < max
	})
	return c
}

// Enum 限制值必须在指定枚举值中
func (c *checkNumber[T]) Enum(enums []T, opts ...RuleOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncEnum, func() []string {
		return c.ntsf(enums...)
	}, func() bool {
		return false
	}, func() bool {
		for _, e := range enums {
			if *c.n == e {
				return true
			}
		}
		return false
	})
	return c
}

// Custom 自定义校验
func (c *checkNumber[T]) Custom(successIfNil bool, custom func(n T) bool, opts ...RuleOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncDefault, nil,
		func() bool {
			return successIfNil
		}, func() bool {
			return custom(*c.n)
		})
	return c
}
