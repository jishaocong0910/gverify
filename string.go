// Copyright 2024 jishaocong0910/@163.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vfy

import (
	"regexp"
	"unicode"
	"unicode/utf8"
)

type checkString struct {
	*Context
	s *string
}

func (c *checkString) success() msg[*checkString] {
	return msg[*checkString]{t: c}
}

func (c *checkString) success_() msg_[*checkString] {
	return msg_[*checkString]{msg: c.success()}
}

func (c *checkString) fail(confines ...[]string) msg[*checkString] {
	c.wronged = true
	for _, cs := range confines {
		c.confines = append(c.confines, cs...)
	}
	return msg[*checkString]{ctx: c.Context, t: c}
}

func (c *checkString) fail_(k defaultMsgKey, confines ...[]string) msg_[*checkString] {
	return msg_[*checkString]{msg: c.fail(confines...), k: k}
}

func (c *checkString) NotNil() msg_[*checkString] {
	if c.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		return c.fail_(default_msg_string_notnil)
	}
	return c.success_()
}

func (c *checkString) NotBlank() msg_[*checkString] {
	return c.NotBlank_(false)
}

func (c *checkString) NotBlank_(omitNil bool) msg_[*checkString] {
	if c.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_string_notblank)
		}
	} else {
		blank := true
		for _, r := range *c.s {
			if !unicode.IsSpace(r) {
				blank = false
				break
			}
		}
		if blank {
			return c.fail_(default_msg_string_notblank)
		}
	}
	return c.success_()
}

func (c *checkString) Length(l int) msg_[*checkString] {
	return c.Length_(l, false)
}

func (c *checkString) Length_(l int, omitNil bool) msg_[*checkString] {
	if c.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_string_length, intToConfines(l))
		}
	} else {
		if utf8.RuneCountInString(*c.s) != l {
			return c.fail_(default_msg_string_length, intToConfines(l))
		}
	}
	return c.success_()
}

func (c *checkString) Regex(r *regexp.Regexp) msg_[*checkString] {
	return c.Regex_(r, false)
}

func (c *checkString) Regex_(r *regexp.Regexp, omitNil bool) msg_[*checkString] {
	if c.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_string_regex)
		}
	} else if r != nil {
		if !r.MatchString(*c.s) {
			return c.fail_(default_msg_string_regex)
		}
	}
	return c.success_()
}

func (c *checkString) Min(min int) msg_[*checkString] {
	return c.Min_(min, false)
}

func (c *checkString) Min_(min int, omitNil bool) msg_[*checkString] {
	if c.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_string_min, intToConfines(min))
		}
	} else if utf8.RuneCountInString(*c.s) < min {
		return c.fail_(default_msg_string_min, intToConfines(min))
	}
	return c.success_()
}

func (c *checkString) Max(max int) msg_[*checkString] {
	return c.Max_(max, false)
}

func (c *checkString) Max_(max int, omitNil bool) msg_[*checkString] {
	if c.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_string_max, intToConfines(max))
		}
	} else if utf8.RuneCountInString(*c.s) > max {
		return c.fail_(default_msg_string_max, intToConfines(max))
	}
	return c.success_()
}

func (c *checkString) Range(min, max int) msg_[*checkString] {
	return c.Range_(min, max, false)
}

func (c *checkString) Range_(min, max int, omitNil bool) msg_[*checkString] {
	if c.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_string_range, intToConfines(min, max))
		}
	} else if utf8.RuneCountInString(*c.s) < min {
		return c.fail_(default_msg_string_range, intToConfines(min, max))
	} else if utf8.RuneCountInString(*c.s) > max {
		return c.fail_(default_msg_string_range, intToConfines(min, max))
	}
	return c.success_()
}

func (c *checkString) Gt(min int) msg_[*checkString] {
	return c.Gt_(min, false)
}

func (c *checkString) Gt_(min int, omitNil bool) msg_[*checkString] {
	if c.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_string_gt, intToConfines(min))
		}
	} else if utf8.RuneCountInString(*c.s) <= min {
		return c.fail_(default_msg_string_gt, intToConfines(min))
	}
	return c.success_()
}

func (c *checkString) Lt(max int) msg_[*checkString] {
	return c.Lt_(max, false)
}

func (c *checkString) Lt_(max int, omitNil bool) msg_[*checkString] {
	if c.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_string_lt, intToConfines(max))
		}
	} else if utf8.RuneCountInString(*c.s) >= max {
		return c.fail_(default_msg_string_lt, intToConfines(max))
	}
	return c.success_()
}

func (c *checkString) Within(min, max int) msg_[*checkString] {
	return c.Within_(min, max, false)
}

func (c *checkString) Within_(min, max int, omitNil bool) msg_[*checkString] {
	if c.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_string_within, intToConfines(min, max))
		}
	} else if utf8.RuneCountInString(*c.s) <= min {
		return c.fail_(default_msg_string_within, intToConfines(min, max))
	} else if utf8.RuneCountInString(*c.s) >= max {
		return c.fail_(default_msg_string_within, intToConfines(min, max))
	}
	return c.success_()
}

func (c *checkString) Options(options []string) msg_[*checkString] {
	return c.Options_(options, false)
}

func (c *checkString) Options_(options []string, omitNil bool) msg_[*checkString] {
	if c.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_string_options, options)
		}
	} else {
		match := false
		for _, o := range options {
			if *c.s == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_string_options, options)
		}
	}
	return c.success_()
}

func (c *checkString) Custom(custom func(i string) bool) msg[*checkString] {
	return c.Custom_(custom, false)
}

func (c *checkString) Custom_(custom func(i string) bool, omitNil bool) msg[*checkString] {
	if c.interrupt() {
		return c.success()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail()
		}
	} else {
		ok := custom(*c.s)
		if !ok {
			return c.fail()
		}
	}
	return c.success()
}
