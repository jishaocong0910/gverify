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

type checkInt struct {
	*Context
	i *int
}

func (c *checkInt) success() msg[*checkInt] {
	return msg[*checkInt]{t: c}
}

func (c *checkInt) success_() msg_[*checkInt] {
	return msg_[*checkInt]{msg: c.success()}
}

func (c *checkInt) fail(confines ...[]string) msg[*checkInt] {
	c.wronged = true
	for _, cs := range confines {
		c.confines = append(c.confines, cs...)
	}
	return msg[*checkInt]{ctx: c.Context, t: c}
}

func (c *checkInt) fail_(k defaultMsgKey, confines ...[]string) msg_[*checkInt] {
	return msg_[*checkInt]{msg: c.fail(confines...), k: k}
}

func (c *checkInt) NotNil() msg_[*checkInt] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		return c.fail_(default_msg_int_notnil)
	}
	return c.success_()
}

func (c *checkInt) Min(min int) msg_[*checkInt] {
	return c.Min_(min, false)
}

func (c *checkInt) Min_(min int, omitNil bool) msg_[*checkInt] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int_min, intToConfines(min))
		}
	} else if *c.i < min {
		return c.fail_(default_msg_int_min, intToConfines(min))
	}
	return c.success_()
}

func (c *checkInt) Max(max int) msg_[*checkInt] {
	return c.Max_(max, false)
}

func (c *checkInt) Max_(max int, omitNil bool) msg_[*checkInt] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int_max, intToConfines(max))
		}
	} else if *c.i > max {
		return c.fail_(default_msg_int_max, intToConfines(max))
	}
	return c.success_()
}

func (c *checkInt) Range(min, max int) msg_[*checkInt] {
	return c.Range_(min, max, false)
}

func (c *checkInt) Range_(min, max int, omitNil bool) msg_[*checkInt] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int_range, intToConfines(min, max))
		}
	} else if *c.i < min {
		return c.fail_(default_msg_int_range, intToConfines(min, max))
	} else if *c.i > max {
		return c.fail_(default_msg_int_range, intToConfines(min, max))
	}
	return c.success_()
}

func (c *checkInt) Gt(min int) msg_[*checkInt] {
	return c.Gt_(min, false)
}

func (c *checkInt) Gt_(min int, omitNil bool) msg_[*checkInt] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int_gt, intToConfines(min))
		}
	} else if *c.i <= min {
		return c.fail_(default_msg_int_gt, intToConfines(min))
	}
	return c.success_()
}

func (c *checkInt) Lt(max int) msg_[*checkInt] {
	return c.Lt_(max, false)
}

func (c *checkInt) Lt_(max int, omitNil bool) msg_[*checkInt] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int_lt, intToConfines(max))
		}
	} else if *c.i >= max {
		return c.fail_(default_msg_int_lt, intToConfines(max))
	}
	return c.success_()
}

func (c *checkInt) Within(min, max int) msg_[*checkInt] {
	return c.Within_(min, max, false)
}

func (c *checkInt) Within_(min, max int, omitNil bool) msg_[*checkInt] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int_within, intToConfines(min, max))
		}
	} else if *c.i <= min {
		return c.fail_(default_msg_int_within, intToConfines(min, max))
	} else if *c.i >= max {
		return c.fail_(default_msg_int_within, intToConfines(min, max))
	}
	return c.success_()
}

func (c *checkInt) Options(options []int) msg_[*checkInt] {
	return c.Options_(options, false)
}

func (c *checkInt) Options_(options []int, omitNil bool) msg_[*checkInt] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int_options, intToConfines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.i == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_int_options, intToConfines(options...))
		}
	}
	return c.success_()
}

func (c *checkInt) Custom(custom func(i int) bool) msg[*checkInt] {
	return c.Custom_(custom, false)
}

func (c *checkInt) Custom_(custom func(i int) bool, omitNil bool) msg[*checkInt] {
	if c.interrupt() {
		return c.success()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail()
		}
	} else {
		ok := custom(*c.i)
		if !ok {
			return c.fail()
		}
	}
	return c.success()
}
