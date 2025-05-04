package vfy

type checkBool struct {
	vc *VContext
	b  *bool
}

func (c *checkBool) Required(opts ...ItemOption) *checkBool {
	checkRequired[int, bool](c.vc, c.b, opts)
	return c
}

func (c *checkBool) Custom(custom func(b bool) bool, opts ...ItemOption) *checkBool {
	checkPredicate[int, bool](c.vc, c.b, opts, msg_key_default, nil, func() bool {
		return custom(*c.b)
	})
	return c
}
