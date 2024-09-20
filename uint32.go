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

type checkUint32 struct {
	*Context
	u *uint32
}

func (c *checkUint32) success() msg[*checkUint32] {
	return msg[*checkUint32]{t: c}
}

func (c *checkUint32) success_() msg_[*checkUint32] {
	return msg_[*checkUint32]{msg: c.success()}
}

func (c *checkUint32) fail(confines ...[]string) msg[*checkUint32] {
	c.wronged = true
	for _, cs := range confines {
		c.confines = append(c.confines, cs...)
	}
	return msg[*checkUint32]{ctx: c.Context, t: c}
}

func (c *checkUint32) fail_(k defaultMsgKey, confines ...[]string) msg_[*checkUint32] {
	return msg_[*checkUint32]{msg: c.fail(confines...), k: k}
}

func (c *checkUint32) NotNil() msg_[*checkUint32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		return c.fail_(default_msg_uint32_notnil)
	}
	return c.success_()
}

func (c *checkUint32) Min(min uint32) msg_[*checkUint32] {
	return c.Min_(min, false)
}

func (c *checkUint32) Min_(min uint32, omitNil bool) msg_[*checkUint32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint32_min, uint32ToConfines(min))
		}
	} else if *c.u < min {
		return c.fail_(default_msg_uint32_min, uint32ToConfines(min))
	}
	return c.success_()
}

func (c *checkUint32) Max(max uint32) msg_[*checkUint32] {
	return c.Max_(max, false)
}

func (c *checkUint32) Max_(max uint32, omitNil bool) msg_[*checkUint32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint32_max, uint32ToConfines(max))
		}
	} else if *c.u > max {
		return c.fail_(default_msg_uint32_max, uint32ToConfines(max))
	}
	return c.success_()
}

func (c *checkUint32) Range(min, max uint32) msg_[*checkUint32] {
	return c.Range_(min, max, false)
}

func (c *checkUint32) Range_(min, max uint32, omitNil bool) msg_[*checkUint32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint32_range, uint32ToConfines(min, max))
		}
	} else if *c.u < min {
		return c.fail_(default_msg_uint32_range, uint32ToConfines(min, max))
	} else if *c.u > max {
		return c.fail_(default_msg_uint32_range, uint32ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkUint32) Gt(min uint32) msg_[*checkUint32] {
	return c.Gt_(min, false)
}

func (c *checkUint32) Gt_(min uint32, omitNil bool) msg_[*checkUint32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint32_gt, uint32ToConfines(min))
		}
	} else if *c.u <= min {
		return c.fail_(default_msg_uint32_gt, uint32ToConfines(min))
	}
	return c.success_()
}

func (c *checkUint32) Lt(max uint32) msg_[*checkUint32] {
	return c.Lt_(max, false)
}

func (c *checkUint32) Lt_(max uint32, omitNil bool) msg_[*checkUint32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint32_lt, uint32ToConfines(max))
		}
	} else if *c.u >= max {
		return c.fail_(default_msg_uint32_lt, uint32ToConfines(max))
	}
	return c.success_()
}

func (c *checkUint32) Within(min, max uint32) msg_[*checkUint32] {
	return c.Within_(min, max, false)
}

func (c *checkUint32) Within_(min, max uint32, omitNil bool) msg_[*checkUint32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint32_within, uint32ToConfines(min, max))
		}
	} else if *c.u <= min {
		return c.fail_(default_msg_uint32_within, uint32ToConfines(min, max))
	} else if *c.u >= max {
		return c.fail_(default_msg_uint32_within, uint32ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkUint32) Options(options []uint32) msg_[*checkUint32] {
	return c.Options_(options, false)
}

func (c *checkUint32) Options_(options []uint32, omitNil bool) msg_[*checkUint32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint32_options, uint32ToConfines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.u == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_uint32_options, uint32ToConfines(options...))
		}
	}
	return c.success_()
}

func (c *checkUint32) Custom(custom func(u uint32) bool) msg[*checkUint32] {
	return c.Custom_(custom, false)
}

func (c *checkUint32) Custom_(custom func(u uint32) bool, omitNil bool) msg[*checkUint32] {
	if c.interrupt() {
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
