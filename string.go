package vfy

import (
	"regexp"
	"unicode"
)

type checkString struct {
	vc *VContext
	s  *string
}

func (c *checkString) Required(opts ...checkOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msgBuildFuncRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

func (c *checkString) NotBlank(opts ...checkOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msgBuildFuncNotBlank, nil, func() bool {
		return false
	}, func() bool {
		for _, r := range *c.s {
			if !unicode.IsSpace(r) {
				return true
			}
		}
		return false
	})
	return c
}

func (c *checkString) Regex(r *regexp.Regexp, opts ...checkOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msgBuildFuncRegex, nil, func() bool {
		return r.MatchString("")
	}, func() bool {
		return r.MatchString(*c.s)
	})
	return c
}

func (c *checkString) Length(l int, opts ...checkOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msgBuildFuncLength, func() []string {
		return intToStr(l)
	}, func() bool {
		return l == 0
	}, func() bool {
		return len(*c.s) == l
	})
	return c
}

func (c *checkString) Min(min int, opts ...checkOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msgBuildFuncLengthMin, func() []string {
		return intToStr(min)
	}, func() bool {
		return min <= 0
	}, func() bool {
		return len(*c.s) >= min
	})
	return c
}

func (c *checkString) Max(max int, opts ...checkOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msgBuildFuncLengthMax, func() []string {
		return intToStr(max)
	}, func() bool {
		return max >= 0
	}, func() bool {
		return len(*c.s) <= max
	})
	return c
}

func (c *checkString) Range(min, max int, opts ...checkOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msgBuildFuncLengthRange, func() []string {
		return intToStr(min, max)
	}, func() bool {
		return min <= 0 && max >= 0
	}, func() bool {
		l := len(*c.s)
		return l >= min && l <= max
	})
	return c
}

func (c *checkString) Gt(min int, opts ...checkOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msgBuildFuncLengthGt, func() []string {
		return intToStr(min)
	}, func() bool {
		return min < 0
	}, func() bool {
		return len(*c.s) > min
	})
	return c
}

func (c *checkString) Lt(max int, opts ...checkOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msgBuildFuncLengthLt, func() []string {
		return intToStr(max)
	}, func() bool {
		return max > 0
	}, func() bool {
		return len(*c.s) < max
	})
	return c
}

func (c *checkString) Within(min, max int, opts ...checkOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msgBuildFuncLengthWithin, func() []string {
		return intToStr(min, max)
	}, func() bool {
		return min < 0 && max > 0
	}, func() bool {
		l := len(*c.s)
		return l > min && l < max
	})
	return c
}

func (c *checkString) Enum(enums []string, opts ...checkOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msgBuildFuncEnum, func() []string {
		var confines []string
		for _, e := range enums {
			confines = append(confines, "\""+e+"\"")
		}
		return confines
	}, func() bool {
		return false
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

func (c *checkString) Custom(successIfNil bool, custom func(s string) bool, opts ...checkOption) *checkString {
	checkPredicate[int, string](c.vc, c.s, opts, msgBuildFuncDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(*c.s)
	})
	return c
}
