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

type checkInt8 struct {
	ctx *Context
	i   *int8
}

func (c *checkInt8) success() msg[*checkInt8] {
	return msg[*checkInt8]{t: c}
}

func (c *checkInt8) success_() msg_[*checkInt8] {
	return msg_[*checkInt8]{msg: c.success()}
}

func (c *checkInt8) fail(confines ...[]string) msg[*checkInt8] {
	c.ctx.wronged = true
	for _, cs := range confines {
		c.ctx.confines = append(c.ctx.confines, cs...)
	}
	return msg[*checkInt8]{ctx: c.ctx, t: c}
}

func (c *checkInt8) fail_(k defaultMsgKey, confines ...[]string) msg_[*checkInt8] {
	return msg_[*checkInt8]{msg: c.fail(confines...), k: k}
}

func (c *checkInt8) NotNil() msg_[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		return c.fail_(default_msg_int8_notnil)
	}
	return c.success_()
}

func (c *checkInt8) Min(min int8) msg_[*checkInt8] {
	return c.Min_(min, false)
}

func (c *checkInt8) Min_(min int8, omitNil bool) msg_[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int8_min, c.ctx.int8Confines(min))
		}
	} else if *c.i < min {
		return c.fail_(default_msg_int8_min, c.ctx.int8Confines(min))
	}
	return c.success_()
}

func (c *checkInt8) Max(max int8) msg_[*checkInt8] {
	return c.Max_(max, false)
}

func (c *checkInt8) Max_(max int8, omitNil bool) msg_[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int8_max, c.ctx.int8Confines(max))
		}
	} else if *c.i > max {
		return c.fail_(default_msg_int8_max, c.ctx.int8Confines(max))
	}
	return c.success_()
}

func (c *checkInt8) Range(min, max int8) msg_[*checkInt8] {
	return c.Range_(min, max, false)
}

func (c *checkInt8) Range_(min, max int8, omitNil bool) msg_[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int8_range, c.ctx.int8Confines(min, max))
		}
	} else if *c.i < min {
		return c.fail_(default_msg_int8_range, c.ctx.int8Confines(min, max))
	} else if *c.i > max {
		return c.fail_(default_msg_int8_range, c.ctx.int8Confines(min, max))
	}
	return c.success_()
}

func (c *checkInt8) Gt(min int8) msg_[*checkInt8] {
	return c.Gt_(min, false)
}

func (c *checkInt8) Gt_(min int8, omitNil bool) msg_[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int8_gt, c.ctx.int8Confines(min))
		}
	} else if *c.i <= min {
		return c.fail_(default_msg_int8_gt, c.ctx.int8Confines(min))
	}
	return c.success_()
}

func (c *checkInt8) Lt(max int8) msg_[*checkInt8] {
	return c.Lt_(max, false)
}

func (c *checkInt8) Lt_(max int8, omitNil bool) msg_[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int8_lt, c.ctx.int8Confines(max))
		}
	} else if *c.i >= max {
		return c.fail_(default_msg_int8_lt, c.ctx.int8Confines(max))
	}
	return c.success_()
}

func (c *checkInt8) Within(min, max int8) msg_[*checkInt8] {
	return c.Within_(min, max, false)
}

func (c *checkInt8) Within_(min, max int8, omitNil bool) msg_[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int8_within, c.ctx.int8Confines(min, max))
		}
	} else if *c.i <= min {
		return c.fail_(default_msg_int8_within, c.ctx.int8Confines(min, max))
	} else if *c.i >= max {
		return c.fail_(default_msg_int8_within, c.ctx.int8Confines(min, max))
	}
	return c.success_()
}

func (c *checkInt8) Options(options []int8) msg_[*checkInt8] {
	return c.Options_(options, false)
}

func (c *checkInt8) Options_(options []int8, omitNil bool) msg_[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int8_options, c.ctx.int8Confines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.i == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_int8_options, c.ctx.int8Confines(options...))
		}
	}
	return c.success_()
}

func (c *checkInt8) Custom(custom func(i int8) bool) msg[*checkInt8] {
	return c.Custom_(custom, false)
}

func (c *checkInt8) Custom_(custom func(i int8) bool, omitNil bool) msg[*checkInt8] {
	if c.ctx.interrupt() {
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
