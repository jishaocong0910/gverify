package vfy

import "strconv"

type checkSlice[T any] struct {
	vc *VContext
	s  []T
}

func (c *checkSlice[T]) Required(opts ...ItemOption) *checkSlice[T] {
	checkRequired[int, T](c.vc, c.s, opts)
	return c
}

func (c *checkSlice[T]) NotEmpty(opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncNotEmpty, nil, func() bool {
		return len(c.s) != 0
	})
	return c
}

func (c *checkSlice[T]) Length(l int, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncLength, func() []string {
		return intToStr(l)
	}, func() bool {
		return len(c.s) == l
	})
	return c
}

func (c *checkSlice[T]) Min(min int, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncLengthMin, func() []string {
		return intToStr(min)
	}, func() bool {
		return len(c.s) >= min
	})
	return c
}

func (c *checkSlice[T]) Max(max int, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncLengthMax, func() []string {
		return intToStr(max)
	}, func() bool {
		return len(c.s) <= max
	})
	return c
}

func (c *checkSlice[T]) Range(min, max int, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncLengthRange, func() []string {
		return intToStr(min, max)
	}, func() bool {
		l := len(c.s)
		return l >= min && l <= max
	})
	return c
}

func (c *checkSlice[T]) Gt(min int, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncLengthGt, func() []string {
		return intToStr(min)
	}, func() bool {
		return len(c.s) > min
	})
	return c
}

func (c *checkSlice[T]) Lt(max int, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncLengthLt, func() []string {
		return intToStr(max)
	}, func() bool {
		return len(c.s) < max
	})
	return c
}

func (c *checkSlice[T]) Within(min, max int, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncLengthWithin, func() []string {
		return intToStr(min, max)
	}, func() bool {
		l := len(c.s)
		return l > min && l < max
	})
	return c
}

func (c *checkSlice[T]) Custom(custom func(b []T) bool, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msgBuildFuncDefault, nil, func() bool {
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
