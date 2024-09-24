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

func (c *checkInt8) success() setMsg[*checkInt8] {
	return setMsg[*checkInt8]{t: c}
}

func (c *checkInt8) success_() setMsgOrDefault[*checkInt8] {
	return setMsgOrDefault[*checkInt8]{setMsg: c.success()}
}

func (c *checkInt8) fail(confines ...[]string) setMsg[*checkInt8] {
	c.ctx.wronged = true
	for _, cs := range confines {
		c.ctx.confines = append(c.ctx.confines, cs...)
	}
	return setMsg[*checkInt8]{ctx: c.ctx, t: c}
}

func (c *checkInt8) fail_(k defaultMsgKey, confines ...[]string) setMsgOrDefault[*checkInt8] {
	return setMsgOrDefault[*checkInt8]{setMsg: c.fail(confines...), k: k}
}

func (c *checkInt8) NotNil() setMsgOrDefault[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		return c.fail_(default_msg_int8_notnil)
	}
	return c.success_()
}

func (c *checkInt8) Min(min int8) setMsgOrDefault[*checkInt8] {
	return c.Min_(min, false)
}

func (c *checkInt8) Min_(min int8, omitNil bool) setMsgOrDefault[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int8_min, int8ToConfines(min))
		}
	} else if *c.i < min {
		return c.fail_(default_msg_int8_min, int8ToConfines(min))
	}
	return c.success_()
}

func (c *checkInt8) Max(max int8) setMsgOrDefault[*checkInt8] {
	return c.Max_(max, false)
}

func (c *checkInt8) Max_(max int8, omitNil bool) setMsgOrDefault[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int8_max, int8ToConfines(max))
		}
	} else if *c.i > max {
		return c.fail_(default_msg_int8_max, int8ToConfines(max))
	}
	return c.success_()
}

func (c *checkInt8) Range(min, max int8) setMsgOrDefault[*checkInt8] {
	return c.Range_(min, max, false)
}

func (c *checkInt8) Range_(min, max int8, omitNil bool) setMsgOrDefault[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int8_range, int8ToConfines(min, max))
		}
	} else if *c.i < min {
		return c.fail_(default_msg_int8_range, int8ToConfines(min, max))
	} else if *c.i > max {
		return c.fail_(default_msg_int8_range, int8ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkInt8) Gt(min int8) setMsgOrDefault[*checkInt8] {
	return c.Gt_(min, false)
}

func (c *checkInt8) Gt_(min int8, omitNil bool) setMsgOrDefault[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int8_gt, int8ToConfines(min))
		}
	} else if *c.i <= min {
		return c.fail_(default_msg_int8_gt, int8ToConfines(min))
	}
	return c.success_()
}

func (c *checkInt8) Lt(max int8) setMsgOrDefault[*checkInt8] {
	return c.Lt_(max, false)
}

func (c *checkInt8) Lt_(max int8, omitNil bool) setMsgOrDefault[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int8_lt, int8ToConfines(max))
		}
	} else if *c.i >= max {
		return c.fail_(default_msg_int8_lt, int8ToConfines(max))
	}
	return c.success_()
}

func (c *checkInt8) Within(min, max int8) setMsgOrDefault[*checkInt8] {
	return c.Within_(min, max, false)
}

func (c *checkInt8) Within_(min, max int8, omitNil bool) setMsgOrDefault[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int8_within, int8ToConfines(min, max))
		}
	} else if *c.i <= min {
		return c.fail_(default_msg_int8_within, int8ToConfines(min, max))
	} else if *c.i >= max {
		return c.fail_(default_msg_int8_within, int8ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkInt8) Options(options []int8) setMsgOrDefault[*checkInt8] {
	return c.Options_(options, false)
}

func (c *checkInt8) Options_(options []int8, omitNil bool) setMsgOrDefault[*checkInt8] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int8_options, int8ToConfines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.i == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_int8_options, int8ToConfines(options...))
		}
	}
	return c.success_()
}

func (c *checkInt8) Custom(custom func(i int8) bool) setMsg[*checkInt8] {
	return c.Custom_(custom, false)
}

func (c *checkInt8) Custom_(custom func(i int8) bool, omitNil bool) setMsg[*checkInt8] {
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
