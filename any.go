package vfy

type checkAny[T any] struct {
	vc *VContext
	t  *T
}

func (c *checkAny[T]) Required(opts ...ItemOption) *checkAny[T] {
	checkRequired[int, T](c.vc, c.t, opts)
	return c
}

func (c *checkAny[T]) Custom(custom func(t T) bool, opts ...ItemOption) *checkAny[T] {
	checkPredicate[int, T](c.vc, c.t, opts, msg_key_default, nil, func() bool {
		return custom(*c.t)
	})
	return c
}
