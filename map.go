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

// 校验map类型
type checkMap[K comparable, V any] struct {
	vc *VContext
	m  map[K]V
}

// Required 限制不能为nil
func (c *checkMap[K, V]) Required(opts ...RuleOption) *checkMap[K, V] {
	predicate[K, V](c.vc, c.m, opts, msgBuildFuncRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

// NotEmpty 限制不能为空
func (c *checkMap[K, V]) NotEmpty(opts ...RuleOption) *checkMap[K, V] {
	predicate[K, V](c.vc, c.m, opts, msgBuildFuncNotEmpty, nil, func() bool {
		return false
	}, func() bool {
		return len(c.m) > 0
	})
	return c
}

// Length 限制长度必须为指定值
func (c *checkMap[K, V]) Length(l int, opts ...RuleOption) *checkMap[K, V] {
	predicate[K, V](c.vc, c.m, opts, msgBuildFuncLength, func() []string {
		return intToStr(l)
	}, func() bool {
		return l == 0
	}, func() bool {
		return len(c.m) == l
	})
	return c
}

// Min 限制最小长度
func (c *checkMap[K, V]) Min(min int, opts ...RuleOption) *checkMap[K, V] {
	predicate[K, V](c.vc, c.m, opts, msgBuildFuncLengthMin, func() []string {
		return intToStr(min)
	}, func() bool {
		return min <= 0
	}, func() bool {
		return len(c.m) >= min
	})
	return c
}

// Max 限制最大长度
func (c *checkMap[K, V]) Max(max int, opts ...RuleOption) *checkMap[K, V] {
	predicate[K, V](c.vc, c.m, opts, msgBuildFuncLengthMax, func() []string {
		return intToStr(max)
	}, func() bool {
		return max >= 0
	}, func() bool {
		return len(c.m) <= max
	})
	return c
}

// Range 限制长度范围（包含边界）
func (c *checkMap[K, V]) Range(min, max int, opts ...RuleOption) *checkMap[K, V] {
	predicate[K, V](c.vc, c.m, opts, msgBuildFuncLengthRange, func() []string {
		return intToStr(min, max)
	}, func() bool {
		return min <= 0 && max >= 0
	}, func() bool {
		l := len(c.m)
		return l >= min && l <= max
	})
	return c
}

// Gt 限制长度必须大于指定值
func (c *checkMap[K, V]) Gt(min int, opts ...RuleOption) *checkMap[K, V] {
	predicate[K, V](c.vc, c.m, opts, msgBuildFuncLengthGt, func() []string {
		return intToStr(min)
	}, func() bool {
		return min < 0
	}, func() bool {
		return len(c.m) > min
	})
	return c
}

// Lt 限制长度必须小于指定值
func (c *checkMap[K, V]) Lt(max int, opts ...RuleOption) *checkMap[K, V] {
	predicate[K, V](c.vc, c.m, opts, msgBuildFuncLengthLt, func() []string {
		return intToStr(max)
	}, func() bool {
		return max > 0
	}, func() bool {
		return len(c.m) < max
	})
	return c
}

// Within 限制长度范围（不包含边界）
func (c *checkMap[K, V]) Within(min, max int, opts ...RuleOption) *checkMap[K, V] {
	predicate[K, V](c.vc, c.m, opts, msgBuildFuncLengthWithin, func() []string {
		return intToStr(min, max)
	}, func() bool {
		return min < 0 && max > 0
	}, func() bool {
		l := len(c.m)
		return l > min && l < max
	})
	return c
}

// Custom 自定义校验
func (c *checkMap[K, V]) Custom(successIfNil bool, custom func(m map[K]V) bool, opts ...RuleOption) *checkMap[K, V] {
	predicate[K, V](c.vc, c.m, opts, msgBuildFuncDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(c.m)
	})
	return c
}

// Dive 下沉校验key和value
func (c *checkMap[K, V]) Dive(key func(k K), value func(v V)) {
	if c.vc.interrupt() {
		return
	}
	if c.m != nil {
		for k, v := range c.m {
			if key != nil {
				if c.vc.interrupt() {
					break
				}
				fi := c.vc.fieldInfo
				c.vc.beforeDiveSliceMap("$key")
				key(k)
				c.vc.fieldInfo = fi
			}
			if value != nil {
				if c.vc.interrupt() {
					break
				}
				fi := c.vc.fieldInfo
				c.vc.beforeDiveSliceMap("$value")
				value(v)
				c.vc.fieldInfo = fi
			}
		}
	}
}
