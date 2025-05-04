package vfy

type ItemOption func(*FieldInfo)
type FieldOption func(*FieldInfo)
type StructOption func(ctx *VContext)

func Code(code string) ItemOption {
	return func(o *FieldInfo) {
		o.code = code
	}
}

func Msg(mbf msgBuildFunc) ItemOption {
	return func(o *FieldInfo) {
		o.mbf = mbf
	}
}

func Omittable() FieldOption {
	return func(o *FieldInfo) {
		o.omittable = true
	}
}

func All() StructOption {
	return func(vc *VContext) {
		vc.all = true
	}
}
