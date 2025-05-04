package vfy

type checkStruct[V Verifiable] struct {
	vc *VContext
	s  *V
}

func (c *checkStruct[V]) Required(opts ...ItemOption) *checkStruct[V] {
	checkRequired[int, V](c.vc, c.s, opts)
	return c
}

func (c *checkStruct[V]) Custom(custom func(s V) bool, opts ...ItemOption) *checkStruct[V] {
	checkPredicate[int, V](c.vc, c.s, opts, msgBuildFuncDefault, nil, func() bool {
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
