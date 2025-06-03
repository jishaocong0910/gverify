package vfy

type checkMap[K comparable, V any] struct {
	vc *VContext
	m  map[K]V
}

func (c *checkMap[K, V]) Required(opts ...ruleOption) *checkMap[K, V] {
	checkPredicate[K, V](c.vc, c.m, opts, msgBuildFuncRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

func (c *checkMap[K, V]) NotEmpty(opts ...ruleOption) *checkMap[K, V] {
	checkPredicate[K, V](c.vc, c.m, opts, msgBuildFuncNotEmpty, nil, func() bool {
		return false
	}, func() bool {
		return len(c.m) > 0
	})
	return c
}

func (c *checkMap[K, V]) Length(l int, opts ...ruleOption) *checkMap[K, V] {
	checkPredicate[K, V](c.vc, c.m, opts, msgBuildFuncLength, func() []string {
		return intToStr(l)
	}, func() bool {
		return l == 0
	}, func() bool {
		return len(c.m) == l
	})
	return c
}

func (c *checkMap[K, V]) Min(min int, opts ...ruleOption) *checkMap[K, V] {
	checkPredicate[K, V](c.vc, c.m, opts, msgBuildFuncLengthMin, func() []string {
		return intToStr(min)
	}, func() bool {
		return min <= 0
	}, func() bool {
		return len(c.m) >= min
	})
	return c
}

func (c *checkMap[K, V]) Max(max int, opts ...ruleOption) *checkMap[K, V] {
	checkPredicate[K, V](c.vc, c.m, opts, msgBuildFuncLengthMax, func() []string {
		return intToStr(max)
	}, func() bool {
		return max >= 0
	}, func() bool {
		return len(c.m) <= max
	})
	return c
}

func (c *checkMap[K, V]) Range(min, max int, opts ...ruleOption) *checkMap[K, V] {
	checkPredicate[K, V](c.vc, c.m, opts, msgBuildFuncLengthRange, func() []string {
		return intToStr(min, max)
	}, func() bool {
		return min <= 0 && max >= 0
	}, func() bool {
		l := len(c.m)
		return l >= min && l <= max
	})
	return c
}

func (c *checkMap[K, V]) Gt(min int, opts ...ruleOption) *checkMap[K, V] {
	checkPredicate[K, V](c.vc, c.m, opts, msgBuildFuncLengthGt, func() []string {
		return intToStr(min)
	}, func() bool {
		return min < 0
	}, func() bool {
		return len(c.m) > min
	})
	return c
}

func (c *checkMap[K, V]) Lt(max int, opts ...ruleOption) *checkMap[K, V] {
	checkPredicate[K, V](c.vc, c.m, opts, msgBuildFuncLengthLt, func() []string {
		return intToStr(max)
	}, func() bool {
		return max > 0
	}, func() bool {
		return len(c.m) < max
	})
	return c
}

func (c *checkMap[K, V]) Within(min, max int, opts ...ruleOption) *checkMap[K, V] {
	checkPredicate[K, V](c.vc, c.m, opts, msgBuildFuncLengthWithin, func() []string {
		return intToStr(min, max)
	}, func() bool {
		return min < 0 && max > 0
	}, func() bool {
		l := len(c.m)
		return l > min && l < max
	})
	return c
}

func (c *checkMap[K, V]) Custom(successIfNil bool, custom func(m map[K]V) bool, opts ...ruleOption) *checkMap[K, V] {
	checkPredicate[K, V](c.vc, c.m, opts, msgBuildFuncDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(c.m)
	})
	return c
}

func (c *checkMap[K, V]) Dive(key func(k K), value func(v V)) {
	if c.vc.interrupt() {
		return
	}
	if c.m != nil {
		f := c.vc.copyFieldInfo()
		c.vc.diveSliceMap()
		for k, v := range c.m {
			if key != nil {
				if c.vc.interrupt() {
					break
				}
				f := c.vc.copyFieldInfo()
				c.vc.beforeCheckElem("$key")
				key(k)
				c.vc.fieldInfo = f
			}
			if value != nil {
				if c.vc.interrupt() {
					break
				}
				f := c.vc.copyFieldInfo()
				c.vc.beforeCheckElem("$value")
				value(v)
				c.vc.fieldInfo = f
			}
		}
		c.vc.fieldInfo = f
	}
}
