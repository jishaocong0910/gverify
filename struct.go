package vfy

type checkStruct[V Verifiable] struct {
	vc *VContext
	s  *V
}

func (c *checkStruct[V]) Required(opts ...ruleOption) *checkStruct[V] {
	checkPredicate[int, V](c.vc, c.s, opts, msgBuildFuncRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

func (c *checkStruct[V]) Custom(successIfNil bool, custom func(s V) bool, opts ...ruleOption) *checkStruct[V] {
	checkPredicate[int, V](c.vc, c.s, opts, msgBuildFuncDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(*c.s)
	})
	return c
}

func (c *checkStruct[V]) Dive() {
	if c.s != nil {
		f := c.vc.copyFieldInfo()
		c.vc.diveStruct()
		(*c.s).Checklist(c.vc)
		c.vc.fieldInfo = f
	}
}
