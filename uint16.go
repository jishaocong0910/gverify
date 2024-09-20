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

type checkUint16 struct {
	*Context
	u *uint16
}

func (c *checkUint16) success() msg[*checkUint16] {
	return msg[*checkUint16]{t: c}
}

func (c *checkUint16) success_() msg_[*checkUint16] {
	return msg_[*checkUint16]{msg: c.success()}
}

func (c *checkUint16) fail(confines ...[]string) msg[*checkUint16] {
	c.wronged = true
	for _, cs := range confines {
		c.confines = append(c.confines, cs...)
	}
	return msg[*checkUint16]{ctx: c.Context, t: c}
}

func (c *checkUint16) fail_(k defaultMsgKey, confines ...[]string) msg_[*checkUint16] {
	return msg_[*checkUint16]{msg: c.fail(confines...), k: k}
}

func (c *checkUint16) NotNil() msg_[*checkUint16] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		return c.fail_(default_msg_uint16_notnil)
	}
	return c.success_()
}

func (c *checkUint16) Min(min uint16) msg_[*checkUint16] {
	return c.Min_(min, false)
}

func (c *checkUint16) Min_(min uint16, omitNil bool) msg_[*checkUint16] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint16_min, uint16ToConfines(min))
		}
	} else if *c.u < min {
		return c.fail_(default_msg_uint16_min, uint16ToConfines(min))
	}
	return c.success_()
}

func (c *checkUint16) Max(max uint16) msg_[*checkUint16] {
	return c.Max_(max, false)
}

func (c *checkUint16) Max_(max uint16, omitNil bool) msg_[*checkUint16] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint16_max, uint16ToConfines(max))
		}
	} else if *c.u > max {
		return c.fail_(default_msg_uint16_max, uint16ToConfines(max))
	}
	return c.success_()
}

func (c *checkUint16) Range(min, max uint16) msg_[*checkUint16] {
	return c.Range_(min, max, false)
}

func (c *checkUint16) Range_(min, max uint16, omitNil bool) msg_[*checkUint16] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint16_range, uint16ToConfines(min, max))
		}
	} else if *c.u < min {
		return c.fail_(default_msg_uint16_range, uint16ToConfines(min, max))
	} else if *c.u > max {
		return c.fail_(default_msg_uint16_range, uint16ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkUint16) Gt(min uint16) msg_[*checkUint16] {
	return c.Gt_(min, false)
}

func (c *checkUint16) Gt_(min uint16, omitNil bool) msg_[*checkUint16] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint16_gt, uint16ToConfines(min))
		}
	} else if *c.u <= min {
		return c.fail_(default_msg_uint16_gt, uint16ToConfines(min))
	}
	return c.success_()
}

func (c *checkUint16) Lt(max uint16) msg_[*checkUint16] {
	return c.Lt_(max, false)
}

func (c *checkUint16) Lt_(max uint16, omitNil bool) msg_[*checkUint16] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint16_lt, uint16ToConfines(max))
		}
	} else if *c.u >= max {
		return c.fail_(default_msg_uint16_lt, uint16ToConfines(max))
	}
	return c.success_()
}

func (c *checkUint16) Within(min, max uint16) msg_[*checkUint16] {
	return c.Within_(min, max, false)
}

func (c *checkUint16) Within_(min, max uint16, omitNil bool) msg_[*checkUint16] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint16_within, uint16ToConfines(min, max))
		}
	} else if *c.u <= min {
		return c.fail_(default_msg_uint16_within, uint16ToConfines(min, max))
	} else if *c.u >= max {
		return c.fail_(default_msg_uint16_within, uint16ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkUint16) Options(options []uint16) msg_[*checkUint16] {
	return c.Options_(options, false)
}

func (c *checkUint16) Options_(options []uint16, omitNil bool) msg_[*checkUint16] {
	if c.interrupt() {
		return c.success_()
	}
	if c.u == nil {
		if !omitNil {
			return c.fail_(default_msg_uint16_options, uint16ToConfines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.u == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_uint16_options, uint16ToConfines(options...))
		}
	}
	return c.success_()
}

func (c *checkUint16) Custom(custom func(u uint16) bool) msg[*checkUint16] {
	return c.Custom_(custom, false)
}

func (c *checkUint16) Custom_(custom func(u uint16) bool, omitNil bool) msg[*checkUint16] {
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
