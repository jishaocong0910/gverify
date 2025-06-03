package vfy

import "strconv"

type checkSlice[T any] struct {
	vc *VContext
	s  []T
}

func (c *checkSlice[T]) Required(opts ...ruleOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

func (c *checkSlice[T]) NotEmpty(opts ...ruleOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncNotEmpty, nil, func() bool {
		return false
	}, func() bool {
		return len(c.s) != 0
	})
	return c
}

func (c *checkSlice[T]) Length(l int, opts ...ruleOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncLength, func() []string {
		return intToStr(l)
	}, func() bool {
		return l == 0
	}, func() bool {
		return len(c.s) == l
	})
	return c
}

func (c *checkSlice[T]) Min(min int, opts ...ruleOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncLengthMin, func() []string {
		return intToStr(min)
	}, func() bool {
		return min <= 0
	}, func() bool {
		return len(c.s) >= min
	})
	return c
}

func (c *checkSlice[T]) Max(max int, opts ...ruleOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncLengthMax, func() []string {
		return intToStr(max)
	}, func() bool {
		return max >= 0
	}, func() bool {
		return len(c.s) <= max
	})
	return c
}

func (c *checkSlice[T]) Range(min, max int, opts ...ruleOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncLengthRange, func() []string {
		return intToStr(min, max)
	}, func() bool {
		return min <= 0 && max >= 0
	}, func() bool {
		l := len(c.s)
		return l >= min && l <= max
	})
	return c
}

func (c *checkSlice[T]) Gt(min int, opts ...ruleOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncLengthGt, func() []string {
		return intToStr(min)
	}, func() bool {
		return min < 0
	}, func() bool {
		return len(c.s) > min
	})
	return c
}

func (c *checkSlice[T]) Lt(max int, opts ...ruleOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncLengthLt, func() []string {
		return intToStr(max)
	}, func() bool {
		return max > 0
	}, func() bool {
		return len(c.s) < max
	})
	return c
}

func (c *checkSlice[T]) Within(min, max int, opts ...ruleOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncLengthWithin, func() []string {
		return intToStr(min, max)
	}, func() bool {
		return min < 0 && max > 0
	}, func() bool {
		l := len(c.s)
		return l > min && l < max
	})
	return c
}

func (c *checkSlice[T]) Custom(successIfNil bool, custom func(s []T) bool, opts ...ruleOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(c.s)
	})
	return c
}

func (c *checkSlice[T]) Dive(f func(t T)) {
	if c.vc.interrupt() {
		return
	}
	if c.s != nil && f != nil {
		fi := c.vc.copyFieldInfo()
		c.vc.diveSliceMap()
		for i, t := range c.s {
			if c.vc.interrupt() {
				break
			}
			fi := c.vc.copyFieldInfo()
			c.vc.beforeCheckElem("[" + strconv.Itoa(i) + "]")
			f(t)
			c.vc.fieldInfo = fi
		}
		c.vc.fieldInfo = fi
	}
}
