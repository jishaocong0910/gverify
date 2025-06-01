package vfy

type checkBool struct {
	vc *VContext
	b  *bool
}

func (c *checkBool) Required(opts ...checkOption) *checkBool {
	checkPredicate[int, bool](c.vc, c.b, opts, msgBuildFuncRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

func (c *checkBool) Custom(successIfNil bool, custom func(b bool) bool, opts ...checkOption) *checkBool {
	checkPredicate[int, bool](c.vc, c.b, opts, msgBuildFuncDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(*c.b)
	})
	return c
}
