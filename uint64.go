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

type checkUint64 struct {
	ctx *Context
	u   *uint64
}

func (c *checkUint64) success() msg[*checkUint64] {
	return msg[*checkUint64]{t: c}
}

func (c *checkUint64) success_() msg_[*checkUint64] {
	return msg_[*checkUint64]{msg: c.success()}
}

func (c *checkUint64) fail(confines ...[]string) msg[*checkUint64] {
	c.ctx.wronged = true
	for _, cs := range confines {
		c.ctx.confines = append(c.ctx.confines, cs...)
	}
	return msg[*checkUint64]{ctx: c.ctx, t: c}
}

func (c *checkUint64) fail_(k defaultMsgKey, confines ...[]string) msg_[*checkUint64] {
	return msg_[*checkUint64]{msg: c.fail(confines...), k: k}
}

func (c *checkUint64) NotNil() msg_[*checkUint64] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		return c.fail_(default_msg_uint64_notnil)
	}
	return c.success_()
}

func (c *checkUint64) Min(min uint64) msg_[*checkUint64] {
	return c.Min_(min, false)
}

func (c *checkUint64) Min_(min uint64, omitNil bool) msg_[*checkUint64] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint64_min, c.ctx.uint64Confines(min))
		}
	} else if *c.u < min {
		return c.fail_(default_msg_uint64_min, c.ctx.uint64Confines(min))
	}
	return c.success_()
}

func (c *checkUint64) Max(max uint64) msg_[*checkUint64] {
	return c.Max_(max, false)
}

func (c *checkUint64) Max_(max uint64, omitNil bool) msg_[*checkUint64] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint64_max, c.ctx.uint64Confines(max))
		}
	} else if *c.u > max {
		return c.fail_(default_msg_uint64_max, c.ctx.uint64Confines(max))
	}
	return c.success_()
}

func (c *checkUint64) Range(min, max uint64) msg_[*checkUint64] {
	return c.Range_(min, max, false)
}

func (c *checkUint64) Range_(min, max uint64, omitNil bool) msg_[*checkUint64] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint64_range, c.ctx.uint64Confines(min, max))
		}
	} else if *c.u < min {
		return c.fail_(default_msg_uint64_range, c.ctx.uint64Confines(min, max))
	} else if *c.u > max {
		return c.fail_(default_msg_uint64_range, c.ctx.uint64Confines(min, max))
	}
	return c.success_()
}

func (c *checkUint64) Gt(min uint64) msg_[*checkUint64] {
	return c.Gt_(min, false)
}

func (c *checkUint64) Gt_(min uint64, omitNil bool) msg_[*checkUint64] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint64_gt, c.ctx.uint64Confines(min))
		}
	} else if *c.u <= min {
		return c.fail_(default_msg_uint64_gt, c.ctx.uint64Confines(min))
	}
	return c.success_()
}

func (c *checkUint64) Lt(max uint64) msg_[*checkUint64] {
	return c.Lt_(max, false)
}

func (c *checkUint64) Lt_(max uint64, omitNil bool) msg_[*checkUint64] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint64_lt, c.ctx.uint64Confines(max))
		}
	} else if *c.u >= max {
		return c.fail_(default_msg_uint64_lt, c.ctx.uint64Confines(max))
	}
	return c.success_()
}

func (c *checkUint64) Within(min, max uint64) msg_[*checkUint64] {
	return c.Within_(min, max, false)
}

func (c *checkUint64) Within_(min, max uint64, omitNil bool) msg_[*checkUint64] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint64_within, c.ctx.uint64Confines(min, max))
		}
	} else if *c.u <= min {
		return c.fail_(default_msg_uint64_within, c.ctx.uint64Confines(min, max))
	} else if *c.u >= max {
		return c.fail_(default_msg_uint64_within, c.ctx.uint64Confines(min, max))
	}
	return c.success_()
}

func (c *checkUint64) Options(options []uint64) msg_[*checkUint64] {
	return c.Options_(options, false)
}

func (c *checkUint64) Options_(options []uint64, omitNil bool) msg_[*checkUint64] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint64_options, c.ctx.uint64Confines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.u == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_uint64_options, c.ctx.uint64Confines(options...))
		}
	}
	return c.success_()
}

func (c *checkUint64) Custom(custom func(u uint64) bool) msg[*checkUint64] {
	return c.Custom_(custom, false)
}

func (c *checkUint64) Custom_(custom func(u uint64) bool, omitNil bool) msg[*checkUint64] {
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
