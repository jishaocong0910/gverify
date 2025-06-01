package vfy

type checkOption func(*FieldInfo)
type fieldOption func(*FieldInfo)
type structOption func(ctx *VContext)

func Code(code string) checkOption {
	return func(o *FieldInfo) {
		o.code = code
	}
}

func Msg(mbf msgBuildFunc) checkOption {
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
