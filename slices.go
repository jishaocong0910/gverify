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
	checkPredicate[int, T](c.vc, c.s, opts, msg_key_not_empty, nil, func() bool {
		return len(c.s) != 0
	})
	return c
}

func (c *checkSlice[T]) Length(l int, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msg_key_length, func() []string {
		return intToConfine(l)
	}, func() bool {
		return len(c.s) == l
	})
	return c
}

func (c *checkSlice[T]) Min(min int, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msg_key_length_min, func() []string {
		return intToConfine(min)
	}, func() bool {
		return len(c.s) >= min
	})
	return c
}

func (c *checkSlice[T]) Max(max int, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msg_key_length_max, func() []string {
		return intToConfine(max)
	}, func() bool {
		return len(c.s) <= max
	})
	return c
}

func (c *checkSlice[T]) Range(min, max int, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msg_key_length_range, func() []string {
		return intToConfine(min, max)
	}, func() bool {
		l := len(c.s)
		return l >= min && l <= max
	})
	return c
}

func (c *checkSlice[T]) Gt(min int, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msg_key_length_gt, func() []string {
		return intToConfine(min)
	}, func() bool {
		return len(c.s) > min
	})
	return c
}

func (c *checkSlice[T]) Lt(max int, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msg_key_length_lt, func() []string {
		return intToConfine(max)
	}, func() bool {
		return len(c.s) < max
	})
	return c
}

func (c *checkSlice[T]) Within(min, max int, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msg_key_length_within, func() []string {
		return intToConfine(min, max)
	}, func() bool {
		l := len(c.s)
		return l > min && l < max
	})
	return c
}

func (c *checkSlice[T]) Custom(custom func(b []T) bool, opts ...ItemOption) *checkSlice[T] {
	checkPredicate[int, T](c.vc, c.s, opts, msg_key_default, nil, func() bool {
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
