package vfy

type msgBuildFunc func(f *FieldInfo)

type msgKey int

const (
	msg_key_default msgKey = iota
	msg_key_required
	msg_key_min
	msg_key_max
	msg_key_range
	msg_key_gt
	msg_key_lt
	msg_key_within
	msg_key_enum
	msg_key_not_blank
	msg_key_length
	msg_key_regex
	msg_key_length_min
	msg_key_length_max
	msg_key_length_range
	msg_key_length_gt
	msg_key_length_lt
	msg_key_length_within
	msg_key_not_empty
)

var defaultMsgBuildFuncs = map[msgKey]msgBuildFunc{
	msg_key_default: func(b *FieldInfo) {
		b.Msg("%s is illegal", b.fieldName)
	},
	msg_key_required: func(b *FieldInfo) {
		b.Msg("%s is required", b.fieldName)
	},
	msg_key_min: func(b *FieldInfo) {
		b.Msg("%s must not be less than %s", b.fieldName, b.Confine(0))
	},
	msg_key_max: func(b *FieldInfo) {
		b.Msg("%s must not be greater than %s", b.fieldName, b.Confine(0))
	},
	msg_key_range: func(b *FieldInfo) {
		b.Msg("%s must not be %s to %s", b.fieldName, b.Confine(0), b.Confine(1))
	},
	msg_key_gt: func(b *FieldInfo) {
		b.Msg("%s must be greater than %s", b.fieldName, b.Confine(0))
	},
	msg_key_lt: func(b *FieldInfo) {
		b.Msg("%s must be less than %s", b.fieldName, b.Confine(0))
	},
	msg_key_within: func(b *FieldInfo) {
		b.Msg("%s must be greater than %s and less than %s", b.fieldName, b.Confine(0), b.Confine(1))
	},
	msg_key_enum: func(b *FieldInfo) {
		b.Msg("%s must be %s", b.fieldName, b.Confines())
	},
	msg_key_not_blank: func(b *FieldInfo) {
		b.Msg("%s must not be blank", b.fieldName)
	},
	msg_key_not_empty: func(b *FieldInfo) {
		b.Msg("%s must not be empty", b.fieldName)
	},
	msg_key_regex: func(b *FieldInfo) {
		b.Msg("%s's format is illegal", b.fieldName)
	},
	msg_key_length: func(b *FieldInfo) {
		b.Msg("%s's length must be %s", b.fieldName, b.Confine(0))
	},
	msg_key_length_min: func(b *FieldInfo) {
		b.Msg("%s's length must not be less than %s", b.fieldName, b.Confine(0))
	},
	msg_key_length_max: func(b *FieldInfo) {
		b.Msg("%s's length must not be greater than %s", b.fieldName, b.Confine(0))
	},
	msg_key_length_range: func(b *FieldInfo) {
		b.Msg("%s's length must be %s to %s", b.fieldName, b.Confine(0), b.Confine(1))
	},
	msg_key_length_gt: func(b *FieldInfo) {
		b.Msg("%s's length must be greater than %s", b.fieldName, b.Confine(0))
	},
	msg_key_length_lt: func(b *FieldInfo) {
		b.Msg("%s's length must be less than %s", b.fieldName, b.Confine(0))
	},
	msg_key_length_within: func(b *FieldInfo) {
		b.Msg("%s's length must be greater than %s and less than %s", b.fieldName, b.Confine(0), b.Confine(1))
	},
}
