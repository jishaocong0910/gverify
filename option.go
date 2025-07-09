package vfy

type ruleOption func(*FieldInfo)
type fieldOption func(*FieldInfo)
type structOption func(vc *VContext)

func Code(code string) ruleOption {
	return func(o *FieldInfo) {
		o.code = code
	}
}

func Msg(mbf msgBuildFunc) ruleOption {
	return func(o *FieldInfo) {
		o.mbf = mbf
	}
}

func Omittable() fieldOption {
	return func(o *FieldInfo) {
		o.omittable = true
	}
}

func All() structOption {
	return func(vc *VContext) {
		vc.all = true
	}
}
