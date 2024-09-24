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

type checkFloat32 struct {
	ctx *Context
	f   *float32
}

func (c *checkFloat32) success() setMsg[*checkFloat32] {
	return setMsg[*checkFloat32]{t: c}
}

func (c *checkFloat32) success_() setMsgOrDefault[*checkFloat32] {
	return setMsgOrDefault[*checkFloat32]{setMsg: c.success()}
}

func (c *checkFloat32) fail(confines ...[]string) setMsg[*checkFloat32] {
	c.ctx.wronged = true
	for _, cs := range confines {
		c.ctx.confines = append(c.ctx.confines, cs...)
	}
	return setMsg[*checkFloat32]{ctx: c.ctx, t: c}
}

func (c *checkFloat32) fail_(k defaultMsgKey, confines ...[]string) setMsgOrDefault[*checkFloat32] {
	return setMsgOrDefault[*checkFloat32]{setMsg: c.fail(confines...), k: k}
}

func (c *checkFloat32) NotNil() setMsgOrDefault[*checkFloat32] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		return c.fail_(default_msg_float32_notnil)
	}
	return c.success_()
}

func (c *checkFloat32) Min(min float32) setMsgOrDefault[*checkFloat32] {
	return c.Min_(min, false)
}

func (c *checkFloat32) Min_(min float32, omitNil bool) setMsgOrDefault[*checkFloat32] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail_(default_msg_float32_min, float32ToConfines(min))
		}
	} else if *c.f < min {
		return c.fail_(default_msg_float32_min, float32ToConfines(min))
	}
	return c.success_()
}

func (c *checkFloat32) Max(max float32) setMsgOrDefault[*checkFloat32] {
	return c.Max_(max, false)
}

func (c *checkFloat32) Max_(max float32, omitNil bool) setMsgOrDefault[*checkFloat32] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail_(default_msg_float32_max, float32ToConfines(max))
		}
	} else if *c.f > max {
		return c.fail_(default_msg_float32_max, float32ToConfines(max))
	}
	return c.success_()
}

func (c *checkFloat32) Range(min, max float32) setMsgOrDefault[*checkFloat32] {
	return c.Range_(min, max, false)
}

func (c *checkFloat32) Range_(min, max float32, omitNil bool) setMsgOrDefault[*checkFloat32] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail_(default_msg_float32_range, float32ToConfines(min, max))
		}
	} else if *c.f < min {
		return c.fail_(default_msg_float32_range, float32ToConfines(min, max))
	} else if *c.f > max {
		return c.fail_(default_msg_float32_range, float32ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkFloat32) Gt(min float32) setMsgOrDefault[*checkFloat32] {
	return c.Gt_(min, false)
}

func (c *checkFloat32) Gt_(min float32, omitNil bool) setMsgOrDefault[*checkFloat32] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail_(default_msg_float32_gt, float32ToConfines(min))
		}
	} else if *c.f <= min {
		return c.fail_(default_msg_float32_gt, float32ToConfines(min))
	}
	return c.success_()
}

func (c *checkFloat32) Lt(max float32) setMsgOrDefault[*checkFloat32] {
	return c.Lt_(max, false)
}

func (c *checkFloat32) Lt_(max float32, omitNil bool) setMsgOrDefault[*checkFloat32] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail_(default_msg_float32_lt, float32ToConfines(max))
		}
	} else if *c.f >= max {
		return c.fail_(default_msg_float32_lt, float32ToConfines(max))
	}
	return c.success_()
}

func (c *checkFloat32) Within(min, max float32) setMsgOrDefault[*checkFloat32] {
	return c.Within_(min, max, false)
}

func (c *checkFloat32) Within_(min, max float32, omitNil bool) setMsgOrDefault[*checkFloat32] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail_(default_msg_float32_within, float32ToConfines(min, max))
		}
	} else if *c.f <= min {
		return c.fail_(default_msg_float32_within, float32ToConfines(min, max))
	} else if *c.f >= max {
		return c.fail_(default_msg_float32_within, float32ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkFloat32) Options(options []float32) setMsgOrDefault[*checkFloat32] {
	return c.Options_(options, false)
}

func (c *checkFloat32) Options_(options []float32, omitNil bool) setMsgOrDefault[*checkFloat32] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail_(default_msg_float32_options, float32ToConfines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.f == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_float32_options, float32ToConfines(options...))
		}
	}
	return c.success_()
}

func (c *checkFloat32) Custom(custom func(f float32) bool) setMsg[*checkFloat32] {
	return c.Custom_(custom, false)
}

func (c *checkFloat32) Custom_(custom func(f float32) bool, omitNil bool) setMsg[*checkFloat32] {
	if c.ctx.interrupt() {
		return c.success()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail()
		}
	} else {
		ok := custom(*c.f)
		if !ok {
			return c.fail()
		}
	}
	return c.success()
}
