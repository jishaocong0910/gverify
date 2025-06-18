package vfy

type checkEmbed[V Verifiable] struct {
	vc *VContext
	s  *V
}

func (c *checkEmbed[V]) Dive() {
	if c.s != nil {
		(*c.s).Checklist(c.vc)
	}
}
