package vfy

type number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type numberToConfineFunc[T number] func(t ...T) []string

type checkNumber[T number] struct {
	vc   *VContext
	t    *T
	ntcf numberToConfineFunc[T]
}

func (c *checkNumber[T]) Required(opts ...ItemOption) *checkNumber[T] {
	checkRequired[int, T](c.vc, c.t, opts)
	return c
}

func (c *checkNumber[T]) Min(min T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msg_key_min, func() []string {
		return c.ntcf(min)
	}, func() bool {
		return *c.t >= min
	})
	return c
}

func (c *checkNumber[T]) Max(max T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msg_key_max, func() []string {
		return c.ntcf(max)
	}, func() bool {
		return *c.t <= max
	})
	return c
}

func (c *checkNumber[T]) Range(min, max T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msg_key_range, func() []string {
		return c.ntcf(min, max)
	}, func() bool {
		return *c.t >= min && *c.t <= max
	})
	return c
}

func (c *checkNumber[T]) Gt(min T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msg_key_gt, func() []string {
		return c.ntcf(min)
	}, func() bool {
		return *c.t > min
	})
	return c
}

func (c *checkNumber[T]) Lt(max T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msg_key_lt, func() []string {
		return c.ntcf(max)
	}, func() bool {
		return *c.t < max
	})
	return c
}

func (c *checkNumber[T]) Within(min, max T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msg_key_within, func() []string {
		return c.ntcf(min, max)
	}, func() bool {
		return *c.t > min && *c.t < max
	})
	return c
}

func (c *checkNumber[T]) Enum(enums []T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msg_key_enum, func() []string {
		return c.ntcf(enums...)
	}, func() bool {
		for _, e := range enums {
			if *c.t == e {
				return true
			}
		}
		return false
	})
	return c
}

func (c *checkNumber[T]) Custom(custom func(t T) bool, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msg_key_default, nil, func() bool {
		return custom(*c.t)
	})
	return c
}
