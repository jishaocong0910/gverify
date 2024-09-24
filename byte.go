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

func (c *checkByte) success() setMsg[*checkByte] {
	return setMsg[*checkByte]{t: c}
}

func (c *checkByte) success_() setMsgOrDefault[*checkByte] {
	return setMsgOrDefault[*checkByte]{setMsg: c.success()}
}

func (c *checkByte) fail(confines ...[]string) setMsg[*checkByte] {
	c.ctx.wronged = true
	for _, cs := range confines {
		c.ctx.confines = append(c.ctx.confines, cs...)
	}
	return setMsg[*checkByte]{ctx: c.ctx, t: c}
}

func (c *checkByte) fail_(k defaultMsgKey, confines ...[]string) setMsgOrDefault[*checkByte] {
	return setMsgOrDefault[*checkByte]{setMsg: c.fail(confines...), k: k}
}

func (c *checkByte) NotNil() setMsgOrDefault[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		return c.fail_(default_msg_byte_notnil)
	}
	return c.success_()
}

func (c *checkByte) Min(min byte) setMsgOrDefault[*checkByte] {
	return c.Min_(min, false)
}

func (c *checkByte) Min_(min byte, omitNil bool) setMsgOrDefault[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail_(default_msg_byte_min, byteToConfines(min))
		}
	} else if *c.b < min {
		return c.fail_(default_msg_byte_min, byteToConfines(min))
	}
	return c.success_()
}

func (c *checkByte) Max(max byte) setMsgOrDefault[*checkByte] {
	return c.Max_(max, false)
}

func (c *checkByte) Max_(max byte, omitNil bool) setMsgOrDefault[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail_(default_msg_byte_max, byteToConfines(max))
		}
	} else if *c.b > max {
		return c.fail_(default_msg_byte_max, byteToConfines(max))
	}
	return c.success_()
}

func (c *checkByte) Range(min, max byte) setMsgOrDefault[*checkByte] {
	return c.Range_(min, max, false)
}

func (c *checkByte) Range_(min, max byte, omitNil bool) setMsgOrDefault[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail_(default_msg_byte_range, byteToConfines(min, max))
		}
	} else if *c.b < min {
		return c.fail_(default_msg_byte_range, byteToConfines(min, max))
	} else if *c.b > max {
		return c.fail_(default_msg_byte_range, byteToConfines(min, max))
	}
	return c.success_()
}

func (c *checkByte) Gt(min byte) setMsgOrDefault[*checkByte] {
	return c.Gt_(min, false)
}

func (c *checkByte) Gt_(min byte, omitNil bool) setMsgOrDefault[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail_(default_msg_byte_gt, byteToConfines(min))
		}
	} else if *c.b <= min {
		return c.fail_(default_msg_byte_gt, byteToConfines(min))
	}
	return c.success_()
}

func (c *checkByte) Lt(max byte) setMsgOrDefault[*checkByte] {
	return c.Lt_(max, false)
}

func (c *checkByte) Lt_(max byte, omitNil bool) setMsgOrDefault[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail_(default_msg_byte_lt, byteToConfines(max))
		}
	} else if *c.b >= max {
		return c.fail_(default_msg_byte_lt, byteToConfines(max))
	}
	return c.success_()
}

func (c *checkByte) Within(min, max byte) setMsgOrDefault[*checkByte] {
	return c.Within_(min, max, false)
}

func (c *checkByte) Within_(min, max byte, omitNil bool) setMsgOrDefault[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail_(default_msg_byte_within, byteToConfines(min, max))
		}
	} else if *c.b <= min {
		return c.fail_(default_msg_byte_within, byteToConfines(min, max))
	} else if *c.b >= max {
		return c.fail_(default_msg_byte_within, byteToConfines(min, max))
	}
	return c.success_()
}

func (c *checkByte) Options(options []byte) setMsgOrDefault[*checkByte] {
	return c.Options_(options, false)
}

func (c *checkByte) Options_(options []byte, omitNil bool) setMsgOrDefault[*checkByte] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail_(default_msg_byte_options, byteToConfines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.b == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_byte_options, byteToConfines(options...))
		}
	}
	return c.success_()
}

func (c *checkByte) Custom(custom func(b byte) bool) setMsg[*checkByte] {
	return c.Custom_(custom, false)
}

func (c *checkByte) Custom_(custom func(b byte) bool, omitNil bool) setMsg[*checkByte] {
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
