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

type checkUint struct {
	ctx *Context
	u   *uint
}

func (c *checkUint) success() setMsg[*checkUint] {
	return setMsg[*checkUint]{t: c}
}

func (c *checkUint) success_() setMsgOrDefault[*checkUint] {
	return setMsgOrDefault[*checkUint]{setMsg: c.success()}
}

func (c *checkUint) fail(confines ...[]string) setMsg[*checkUint] {
	c.ctx.wronged = true
	for _, cs := range confines {
		c.ctx.confines = append(c.ctx.confines, cs...)
	}
	return setMsg[*checkUint]{ctx: c.ctx, t: c}
}

func (c *checkUint) fail_(k defaultMsgKey, confines ...[]string) setMsgOrDefault[*checkUint] {
	return setMsgOrDefault[*checkUint]{setMsg: c.fail(confines...), k: k}
}

func (c *checkUint) NotNil() setMsgOrDefault[*checkUint] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		return c.fail_(default_msg_uint_notnil)
	}
	return c.success_()
}

func (c *checkUint) Min(min uint) setMsgOrDefault[*checkUint] {
	return c.Min_(min, false)
}

func (c *checkUint) Min_(min uint, omitNil bool) setMsgOrDefault[*checkUint] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint_min, uintToConfines(min))
		}
	} else if *c.u < min {
		return c.fail_(default_msg_uint_min, uintToConfines(min))
	}
	return c.success_()
}

func (c *checkUint) Max(max uint) setMsgOrDefault[*checkUint] {
	return c.Max_(max, false)
}

func (c *checkUint) Max_(max uint, omitNil bool) setMsgOrDefault[*checkUint] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint_max, uintToConfines(max))
		}
	} else if *c.u > max {
		return c.fail_(default_msg_uint_max, uintToConfines(max))
	}
	return c.success_()
}

func (c *checkUint) Range(min, max uint) setMsgOrDefault[*checkUint] {
	return c.Range_(min, max, false)
}

func (c *checkUint) Range_(min, max uint, omitNil bool) setMsgOrDefault[*checkUint] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint_range, uintToConfines(min, max))
		}
	} else if *c.u < min {
		return c.fail_(default_msg_uint_range, uintToConfines(min, max))
	} else if *c.u > max {
		return c.fail_(default_msg_uint_range, uintToConfines(min, max))
	}
	return c.success_()
}

func (c *checkUint) Gt(min uint) setMsgOrDefault[*checkUint] {
	return c.Gt_(min, false)
}

func (c *checkUint) Gt_(min uint, omitNil bool) setMsgOrDefault[*checkUint] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint_gt, uintToConfines(min))
		}
	} else if *c.u <= min {
		return c.fail_(default_msg_uint_gt, uintToConfines(min))
	}
	return c.success_()
}

func (c *checkUint) Lt(max uint) setMsgOrDefault[*checkUint] {
	return c.Lt_(max, false)
}

func (c *checkUint) Lt_(max uint, omitNil bool) setMsgOrDefault[*checkUint] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint_lt, uintToConfines(max))
		}
	} else if *c.u >= max {
		return c.fail_(default_msg_uint_lt, uintToConfines(max))
	}
	return c.success_()
}

func (c *checkUint) Within(min, max uint) setMsgOrDefault[*checkUint] {
	return c.Within_(min, max, false)
}

func (c *checkUint) Within_(min, max uint, omitNil bool) setMsgOrDefault[*checkUint] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint_within, uintToConfines(min, max))
		}
	} else if *c.u <= min {
		return c.fail_(default_msg_uint_within, uintToConfines(min, max))
	} else if *c.u >= max {
		return c.fail_(default_msg_uint_within, uintToConfines(min, max))
	}
	return c.success_()
}

func (c *checkUint) Options(options []uint) setMsgOrDefault[*checkUint] {
	return c.Options_(options, false)
}

func (c *checkUint) Options_(options []uint, omitNil bool) setMsgOrDefault[*checkUint] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint_options, uintToConfines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.u == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_uint_options, uintToConfines(options...))
		}
	}
	return c.success_()
}

func (c *checkUint) Custom(custom func(u uint) bool) setMsg[*checkUint] {
	return c.Custom_(custom, false)
}

func (c *checkUint) Custom_(custom func(u uint) bool, omitNil bool) setMsg[*checkUint] {
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
