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

type checkInt16 struct {
	ctx *Context
	i   *int16
}

func (c *checkInt16) success() msg[*checkInt16] {
	return msg[*checkInt16]{t: c}
}

func (c *checkInt16) success_() msg_[*checkInt16] {
	return msg_[*checkInt16]{msg: c.success()}
}

func (c *checkInt16) fail(confines ...[]string) msg[*checkInt16] {
	c.ctx.wronged = true
	for _, cs := range confines {
		c.ctx.confines = append(c.ctx.confines, cs...)
	}
	return msg[*checkInt16]{ctx: c.ctx, t: c}
}

func (c *checkInt16) fail_(k defaultMsgKey, confines ...[]string) msg_[*checkInt16] {
	return msg_[*checkInt16]{msg: c.fail(confines...), k: k}
}

func (c *checkInt16) NotNil() msg_[*checkInt16] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		return c.fail_(default_msg_int16_notnil)
	}
	return c.success_()
}

func (c *checkInt16) Min(min int16) msg_[*checkInt16] {
	return c.Min_(min, false)
}

func (c *checkInt16) Min_(min int16, omitNil bool) msg_[*checkInt16] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int16_min, c.ctx.int16Confines(min))
		}
	} else if *c.i < min {
		return c.fail_(default_msg_int16_min, c.ctx.int16Confines(min))
	}
	return c.success_()
}

func (c *checkInt16) Max(max int16) msg_[*checkInt16] {
	return c.Max_(max, false)
}

func (c *checkInt16) Max_(max int16, omitNil bool) msg_[*checkInt16] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int16_max, c.ctx.int16Confines(max))
		}
	} else if *c.i > max {
		return c.fail_(default_msg_int16_max, c.ctx.int16Confines(max))
	}
	return c.success_()
}

func (c *checkInt16) Range(min, max int16) msg_[*checkInt16] {
	return c.Range_(min, max, false)
}

func (c *checkInt16) Range_(min, max int16, omitNil bool) msg_[*checkInt16] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int16_range, c.ctx.int16Confines(min, max))
		}
	} else if *c.i < min {
		return c.fail_(default_msg_int16_range, c.ctx.int16Confines(min, max))
	} else if *c.i > max {
		return c.fail_(default_msg_int16_range, c.ctx.int16Confines(min, max))
	}
	return c.success_()
}

func (c *checkInt16) Gt(min int16) msg_[*checkInt16] {
	return c.Gt_(min, false)
}

func (c *checkInt16) Gt_(min int16, omitNil bool) msg_[*checkInt16] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int16_gt, c.ctx.int16Confines(min))
		}
	} else if *c.i <= min {
		return c.fail_(default_msg_int16_gt, c.ctx.int16Confines(min))
	}
	return c.success_()
}

func (c *checkInt16) Lt(max int16) msg_[*checkInt16] {
	return c.Lt_(max, false)
}

func (c *checkInt16) Lt_(max int16, omitNil bool) msg_[*checkInt16] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int16_lt, c.ctx.int16Confines(max))
		}
	} else if *c.i >= max {
		return c.fail_(default_msg_int16_lt, c.ctx.int16Confines(max))
	}
	return c.success_()
}

func (c *checkInt16) Within(min, max int16) msg_[*checkInt16] {
	return c.Within_(min, max, false)
}

func (c *checkInt16) Within_(min, max int16, omitNil bool) msg_[*checkInt16] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int16_within, c.ctx.int16Confines(min, max))
		}
	} else if *c.i <= min {
		return c.fail_(default_msg_int16_within, c.ctx.int16Confines(min, max))
	} else if *c.i >= max {
		return c.fail_(default_msg_int16_within, c.ctx.int16Confines(min, max))
	}
	return c.success_()
}

func (c *checkInt16) Options(options []int16) msg_[*checkInt16] {
	return c.Options_(options, false)
}

func (c *checkInt16) Options_(options []int16, omitNil bool) msg_[*checkInt16] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int16_options, c.ctx.int16Confines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.i == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_int16_options, c.ctx.int16Confines(options...))
		}
	}
	return c.success_()
}

func (c *checkInt16) Custom(custom func(i int16) bool) msg[*checkInt16] {
	return c.Custom_(custom, false)
}

func (c *checkInt16) Custom_(custom func(i int16) bool, omitNil bool) msg[*checkInt16] {
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
