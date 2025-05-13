package vfy

type checkAny[T any] struct {
	vc *VContext
	a  *T
}

func (c *checkAny[T]) Required(opts ...ItemOption) *checkAny[T] {
	checkRequired[int, T](c.vc, c.a, opts)
	return c
}

func (c *checkAny[T]) Custom(custom func(a T) bool, opts ...ItemOption) *checkAny[T] {
	checkPredicate[int, T](c.vc, c.a, opts, msgBuildFuncDefault, nil, func() bool {
		return custom(*c.a)
	})
	return c
}
