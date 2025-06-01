package vfy

type checkAny[T any] struct {
	vc *VContext
	a  *T
}

func (c *checkAny[T]) Required(opts ...checkOption) *checkAny[T] {
	checkPredicate[int, T](c.vc, c.a, opts, msgBuildFuncRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

func (c *checkAny[T]) Custom(successIfNil bool, custom func(a T) bool, opts ...checkOption) *checkAny[T] {
	checkPredicate[int, T](c.vc, c.a, opts, msgBuildFuncDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(*c.a)
	})
	return c
}
