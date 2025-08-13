package vfy

type StructOption func(vc *VContext)
type FieldOption func(*FieldInfo)
type RuleOption func(*FieldInfo)

func All() StructOption {
	return func(vc *VContext) {
		vc.all = true
	}
}

func Omittable() FieldOption {
	return func(o *FieldInfo) {
		o.omittable = true
	}
}

func Code(code string) RuleOption {
	return func(o *FieldInfo) {
		o.code = code
	}
}

func Msg(mbf msgBuildFunc) RuleOption {
	return func(o *FieldInfo) {
		o.mbf = mbf
	}
}
