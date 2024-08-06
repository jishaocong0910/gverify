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

type checkUint8 struct {
	ctx *Context
	u   *uint8
}

func (c *checkUint8) success() msg[*checkUint8] {
	return msg[*checkUint8]{t: c}
}

func (c *checkUint8) success_() msg_[*checkUint8] {
	return msg_[*checkUint8]{msg: c.success()}
}

func (c *checkUint8) fail(confines ...[]string) msg[*checkUint8] {
	c.ctx.wronged = true
	for _, cs := range confines {
		c.ctx.confines = append(c.ctx.confines, cs...)
	}
	return msg[*checkUint8]{ctx: c.ctx, t: c}
}

func (c *checkUint8) fail_(k defaultMsgKey, confines ...[]string) msg_[*checkUint8] {
	return msg_[*checkUint8]{msg: c.fail(confines...), k: k}
}

func (c *checkUint8) NotNil() msg_[*checkUint8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		return c.fail_(default_msg_uint8_notnil)
	}
	return c.success_()
}

func (c *checkUint8) Min(min uint8) msg_[*checkUint8] {
	return c.Min_(min, false)
}

func (c *checkUint8) Min_(min uint8, omitNil bool) msg_[*checkUint8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint8_min, c.ctx.uint8Confines(min))
		}
	} else if *c.u < min {
		return c.fail_(default_msg_uint8_min, c.ctx.uint8Confines(min))
	}
	return c.success_()
}

func (c *checkUint8) Max(max uint8) msg_[*checkUint8] {
	return c.Max_(max, false)
}

func (c *checkUint8) Max_(max uint8, omitNil bool) msg_[*checkUint8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint8_max, c.ctx.uint8Confines(max))
		}
	} else if *c.u > max {
		return c.fail_(default_msg_uint8_max, c.ctx.uint8Confines(max))
	}
	return c.success_()
}

func (c *checkUint8) Range(min, max uint8) msg_[*checkUint8] {
	return c.Range_(min, max, false)
}

func (c *checkUint8) Range_(min, max uint8, omitNil bool) msg_[*checkUint8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint8_range, c.ctx.uint8Confines(min, max))
		}
	} else if *c.u < min {
		return c.fail_(default_msg_uint8_range, c.ctx.uint8Confines(min, max))
	} else if *c.u > max {
		return c.fail_(default_msg_uint8_range, c.ctx.uint8Confines(min, max))
	}
	return c.success_()
}

func (c *checkUint8) Gt(min uint8) msg_[*checkUint8] {
	return c.Gt_(min, false)
}

func (c *checkUint8) Gt_(min uint8, omitNil bool) msg_[*checkUint8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint8_gt, c.ctx.uint8Confines(min))
		}
	} else if *c.u <= min {
		return c.fail_(default_msg_uint8_gt, c.ctx.uint8Confines(min))
	}
	return c.success_()
}

func (c *checkUint8) Lt(max uint8) msg_[*checkUint8] {
	return c.Lt_(max, false)
}

func (c *checkUint8) Lt_(max uint8, omitNil bool) msg_[*checkUint8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint8_lt, c.ctx.uint8Confines(max))
		}
	} else if *c.u >= max {
		return c.fail_(default_msg_uint8_lt, c.ctx.uint8Confines(max))
	}
	return c.success_()
}

func (c *checkUint8) Within(min, max uint8) msg_[*checkUint8] {
	return c.Within_(min, max, false)
}

func (c *checkUint8) Within_(min, max uint8, omitNil bool) msg_[*checkUint8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint8_within, c.ctx.uint8Confines(min, max))
		}
	} else if *c.u <= min {
		return c.fail_(default_msg_uint8_within, c.ctx.uint8Confines(min, max))
	} else if *c.u >= max {
		return c.fail_(default_msg_uint8_within, c.ctx.uint8Confines(min, max))
	}
	return c.success_()
}

func (c *checkUint8) Options(options []uint8) msg_[*checkUint8] {
	return c.Options_(options, false)
}

func (c *checkUint8) Options_(options []uint8, omitNil bool) msg_[*checkUint8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint8_options, c.ctx.uint8Confines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.u == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_uint8_options, c.ctx.uint8Confines(options...))
		}
	}
	return c.success_()
}

func (c *checkUint8) Custom(custom func(u uint8) bool) msg[*checkUint8] {
	return c.Custom_(custom, false)
}

func (c *checkUint8) Custom_(custom func(u uint8) bool, omitNil bool) msg[*checkUint8] {
	if c.ctx.interrupt() {
		return c.success()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail()
		}
	} else {
		ok := custom(*c.u)
		if !ok {
			return c.fail()
		}
	}
	return c.success()
}
