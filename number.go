package vfy

type number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type numberToStrFunc[T number] func(t ...T) []string

type checkNumber[T number] struct {
	vc   *VContext
	t    *T
	ntsf numberToStrFunc[T]
}

func (c *checkNumber[T]) Required(opts ...ItemOption) *checkNumber[T] {
	checkRequired[int, T](c.vc, c.t, opts)
	return c
}

func (c *checkNumber[T]) Min(min T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msgBuildFuncMin, func() []string {
		return c.ntsf(min)
	}, func() bool {
		return *c.t >= min
	})
	return c
}

func (c *checkNumber[T]) Max(max T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msgBuildFuncMax, func() []string {
		return c.ntsf(max)
	}, func() bool {
		return *c.t <= max
	})
	return c
}

func (c *checkNumber[T]) Range(min, max T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msgBuildFuncRange, func() []string {
		return c.ntsf(min, max)
	}, func() bool {
		return *c.t >= min && *c.t <= max
	})
	return c
}

func (c *checkNumber[T]) Gt(min T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msgBuildFuncGt, func() []string {
		return c.ntsf(min)
	}, func() bool {
		return *c.t > min
	})
	return c
}

func (c *checkNumber[T]) Lt(max T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msgBuildFuncLt, func() []string {
		return c.ntsf(max)
	}, func() bool {
		return *c.t < max
	})
	return c
}

func (c *checkNumber[T]) Within(min, max T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msgBuildFuncWithin, func() []string {
		return c.ntsf(min, max)
	}, func() bool {
		return *c.t > min && *c.t < max
	})
	return c
}

func (c *checkNumber[T]) Enum(enums []T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msgBuildFuncEnum, func() []string {
		return c.ntsf(enums...)
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
	checkPredicate[int, T](c.vc, c.t, opts, msgBuildFuncDefault, nil, func() bool {
		return custom(*c.t)
	})
	return c
}
