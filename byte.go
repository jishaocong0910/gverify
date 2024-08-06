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

type checkByte struct {
	ctx *Context
	b   *byte
}

func (c *checkByte) success() msg[*checkByte] {
	return msg[*checkByte]{t: c}
}

func (c *checkByte) success_() msg_[*checkByte] {
	return msg_[*checkByte]{msg: c.success()}
}

func (c *checkByte) fail(confines ...[]string) msg[*checkByte] {
	c.ctx.wronged = true
	for _, cs := range confines {
		c.ctx.confines = append(c.ctx.confines, cs...)
	}
	return msg[*checkByte]{ctx: c.ctx, t: c}
}

func (c *checkByte) fail_(k defaultMsgKey, confines ...[]string) msg_[*checkByte] {
	return msg_[*checkByte]{msg: c.fail(confines...), k: k}
}

func (c *checkByte) NotNil() msg_[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		return c.fail_(default_msg_byte_notnil)
	}
	return c.success_()
}

func (c *checkByte) Min(min byte) msg_[*checkByte] {
	return c.Min_(min, false)
}

func (c *checkByte) Min_(min byte, omitNil bool) msg_[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail_(default_msg_byte_min, c.ctx.byteConfines(min))
		}
	} else if *c.b < min {
		return c.fail_(default_msg_byte_min, c.ctx.byteConfines(min))
	}
	return c.success_()
}

func (c *checkByte) Max(max byte) msg_[*checkByte] {
	return c.Max_(max, false)
}

func (c *checkByte) Max_(max byte, omitNil bool) msg_[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail_(default_msg_byte_max, c.ctx.byteConfines(max))
		}
	} else if *c.b > max {
		return c.fail_(default_msg_byte_max, c.ctx.byteConfines(max))
	}
	return c.success_()
}

func (c *checkByte) Range(min, max byte) msg_[*checkByte] {
	return c.Range_(min, max, false)
}

func (c *checkByte) Range_(min, max byte, omitNil bool) msg_[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail_(default_msg_byte_range, c.ctx.byteConfines(min, max))
		}
	} else if *c.b < min {
		return c.fail_(default_msg_byte_range, c.ctx.byteConfines(min, max))
	} else if *c.b > max {
		return c.fail_(default_msg_byte_range, c.ctx.byteConfines(min, max))
	}
	return c.success_()
}

func (c *checkByte) Gt(min byte) msg_[*checkByte] {
	return c.Gt_(min, false)
}

func (c *checkByte) Gt_(min byte, omitNil bool) msg_[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail_(default_msg_byte_gt, c.ctx.byteConfines(min))
		}
	} else if *c.b <= min {
		return c.fail_(default_msg_byte_gt, c.ctx.byteConfines(min))
	}
	return c.success_()
}

func (c *checkByte) Lt(max byte) msg_[*checkByte] {
	return c.Lt_(max, false)
}

func (c *checkByte) Lt_(max byte, omitNil bool) msg_[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail_(default_msg_byte_lt, c.ctx.byteConfines(max))
		}
	} else if *c.b >= max {
		return c.fail_(default_msg_byte_lt, c.ctx.byteConfines(max))
	}
	return c.success_()
}

func (c *checkByte) Within(min, max byte) msg_[*checkByte] {
	return c.Within_(min, max, false)
}

func (c *checkByte) Within_(min, max byte, omitNil bool) msg_[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail_(default_msg_byte_within, c.ctx.byteConfines(min, max))
		}
	} else if *c.b <= min {
		return c.fail_(default_msg_byte_within, c.ctx.byteConfines(min, max))
	} else if *c.b >= max {
		return c.fail_(default_msg_byte_within, c.ctx.byteConfines(min, max))
	}
	return c.success_()
}

func (c *checkByte) Options(options []byte) msg_[*checkByte] {
	return c.Options_(options, false)
}

func (c *checkByte) Options_(options []byte, omitNil bool) msg_[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail_(default_msg_byte_options, c.ctx.byteConfines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.b == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_byte_options, c.ctx.byteConfines(options...))
		}
	}
	return c.success_()
}

func (c *checkByte) Custom(custom func(b byte) bool) msg[*checkByte] {
	return c.Custom_(custom, false)
}

func (c *checkByte) Custom_(custom func(b byte) bool, omitNil bool) msg[*checkByte] {
	if c.ctx.interrupt() {
		return c.success()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail()
		}
	} else {
		ok := custom(*c.b)
		if !ok {
			return c.fail()
		}
	}
	return c.success()
}
