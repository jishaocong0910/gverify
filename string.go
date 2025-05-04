package vfy

import (
	"regexp"
	"unicode"
)

type checkString struct {
	vc *VContext
	s  *string
}

func (c *checkString) Required(opts ...ItemOption) *checkString {
	checkRequired[int, string](c.vc, c.s, opts)
	return c
}

func (c *checkString) NotBlank(opts ...ItemOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msg_key_not_blank, nil, func() bool {
		for _, r := range *c.s {
			if !unicode.IsSpace(r) {
				return true
			}
		}
		return false
	})
	return c
}

func (c *checkString) Regex(r *regexp.Regexp, opts ...ItemOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msg_key_regex, nil, func() bool {
		return r.MatchString(*c.s)
	})
	return c
}

func (c *checkString) Length(l int, opts ...ItemOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msg_key_length, func() []string {
		return intToConfine(l)
	}, func() bool {
		return len(*c.s) == l
	})
	return c
}

func (c *checkString) Min(min int, opts ...ItemOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msg_key_length_min, func() []string {
		return intToConfine(min)
	}, func() bool {
		return len(*c.s) >= min
	})
	return c
}

func (c *checkString) Max(max int, opts ...ItemOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msg_key_length_max, func() []string {
		return intToConfine(max)
	}, func() bool {
		return len(*c.s) <= max
	})
	return c
}

func (c *checkString) Range(min, max int, opts ...ItemOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msg_key_length_range, func() []string {
		return intToConfine(min, max)
	}, func() bool {
		l := len(*c.s)
		return l >= min && l <= max
	})
	return c
}

func (c *checkString) Gt(min int, opts ...ItemOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msg_key_length_gt, func() []string {
		return intToConfine(min)
	}, func() bool {
		return len(*c.s) > min
	})
	return c
}

func (c *checkString) Lt(max int, opts ...ItemOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msg_key_length_lt, func() []string {
		return intToConfine(max)
	}, func() bool {
		return len(*c.s) < max
	})
	return c
}

func (c *checkString) Within(min, max int, opts ...ItemOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msg_key_length_within, func() []string {
		return intToConfine(min, max)
	}, func() bool {
		l := len(*c.s)
		return l > min && l < max
	})
	return c
}

func (c *checkString) Enum(enums []string, opts ...ItemOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msg_key_enum, func() []string {
		return enums
	}, func() bool {
		for _, o := range enums {
			if *c.s == o {
				return true
			}
		}
		return false
	})
	return c
}

func (c *checkString) Custom(custom func(s string) bool, opts ...ItemOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msg_key_default, nil, func() bool {
		return custom(*c.s)
	})
	return c
}
