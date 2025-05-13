package vfy

type number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type numberToStrFunc[T number] func(t ...T) []string

type checkNumber[T number] struct {
	vc   *VContext
	n    *T
	ntsf numberToStrFunc[T]
}

func (c *checkNumber[T]) Required(opts ...ItemOption) *checkNumber[T] {
	checkRequired[int, T](c.vc, c.n, opts)
	return c
}

func (c *checkNumber[T]) Min(min T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncMin, func() []string {
		return c.ntsf(min)
	}, func() bool {
		return *c.n >= min
	})
	return c
}

func (c *checkNumber[T]) Max(max T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncMax, func() []string {
		return c.ntsf(max)
	}, func() bool {
		return *c.n <= max
	})
	return c
}

func (c *checkNumber[T]) Range(min, max T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncRange, func() []string {
		return c.ntsf(min, max)
	}, func() bool {
		return *c.n >= min && *c.n <= max
	})
	return c
}

func (c *checkNumber[T]) Gt(min T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncGt, func() []string {
		return c.ntsf(min)
	}, func() bool {
		return *c.n > min
	})
	return c
}

func (c *checkNumber[T]) Lt(max T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncLt, func() []string {
		return c.ntsf(max)
	}, func() bool {
		return *c.n < max
	})
	return c
}

func (c *checkNumber[T]) Within(min, max T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncWithin, func() []string {
		return c.ntsf(min, max)
	}, func() bool {
		return *c.n > min && *c.n < max
	})
	return c
}

func (c *checkNumber[T]) Enum(enums []T, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncEnum, func() []string {
		return c.ntsf(enums...)
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

func (c *checkNumber[T]) Custom(custom func(n T) bool, opts ...ItemOption) *checkNumber[T] {
	checkPredicate[int, T](c.vc, c.n, opts, msgBuildFuncDefault, nil, func() bool {
		return custom(*c.n)
	})
	return c
}
