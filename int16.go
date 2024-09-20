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
	*Context
	i *int16
}

func (c *checkInt16) success() msg[*checkInt16] {
	return msg[*checkInt16]{t: c}
}

func (c *checkInt16) success_() msg_[*checkInt16] {
	return msg_[*checkInt16]{msg: c.success()}
}

func (c *checkInt16) fail(confines ...[]string) msg[*checkInt16] {
	c.wronged = true
	for _, cs := range confines {
		c.confines = append(c.confines, cs...)
	}
	return msg[*checkInt16]{ctx: c.Context, t: c}
}

func (c *checkInt16) fail_(k defaultMsgKey, confines ...[]string) msg_[*checkInt16] {
	return msg_[*checkInt16]{msg: c.fail(confines...), k: k}
}

func (c *checkInt16) NotNil() msg_[*checkInt16] {
	if c.interrupt() {
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
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int16_min, int16ToConfines(min))
		}
	} else if *c.i < min {
		return c.fail_(default_msg_int16_min, int16ToConfines(min))
	}
	return c.success_()
}

func (c *checkInt16) Max(max int16) msg_[*checkInt16] {
	return c.Max_(max, false)
}

func (c *checkInt16) Max_(max int16, omitNil bool) msg_[*checkInt16] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int16_max, int16ToConfines(max))
		}
	} else if *c.i > max {
		return c.fail_(default_msg_int16_max, int16ToConfines(max))
	}
	return c.success_()
}

func (c *checkInt16) Range(min, max int16) msg_[*checkInt16] {
	return c.Range_(min, max, false)
}

func (c *checkInt16) Range_(min, max int16, omitNil bool) msg_[*checkInt16] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int16_range, int16ToConfines(min, max))
		}
	} else if *c.i < min {
		return c.fail_(default_msg_int16_range, int16ToConfines(min, max))
	} else if *c.i > max {
		return c.fail_(default_msg_int16_range, int16ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkInt16) Gt(min int16) msg_[*checkInt16] {
	return c.Gt_(min, false)
}

func (c *checkInt16) Gt_(min int16, omitNil bool) msg_[*checkInt16] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int16_gt, int16ToConfines(min))
		}
	} else if *c.i <= min {
		return c.fail_(default_msg_int16_gt, int16ToConfines(min))
	}
	return c.success_()
}

func (c *checkInt16) Lt(max int16) msg_[*checkInt16] {
	return c.Lt_(max, false)
}

func (c *checkInt16) Lt_(max int16, omitNil bool) msg_[*checkInt16] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int16_lt, int16ToConfines(max))
		}
	} else if *c.i >= max {
		return c.fail_(default_msg_int16_lt, int16ToConfines(max))
	}
	return c.success_()
}

func (c *checkInt16) Within(min, max int16) msg_[*checkInt16] {
	return c.Within_(min, max, false)
}

func (c *checkInt16) Within_(min, max int16, omitNil bool) msg_[*checkInt16] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int16_within, int16ToConfines(min, max))
		}
	} else if *c.i <= min {
		return c.fail_(default_msg_int16_within, int16ToConfines(min, max))
	} else if *c.i >= max {
		return c.fail_(default_msg_int16_within, int16ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkInt16) Options(options []int16) msg_[*checkInt16] {
	return c.Options_(options, false)
}

func (c *checkInt16) Options_(options []int16, omitNil bool) msg_[*checkInt16] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int16_options, int16ToConfines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.i == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_int16_options, int16ToConfines(options...))
		}
	}
	return c.success_()
}

func (c *checkInt16) Custom(custom func(i int16) bool) msg[*checkInt16] {
	return c.Custom_(custom, false)
}

func (c *checkInt16) Custom_(custom func(i int16) bool, omitNil bool) msg[*checkInt16] {
	if c.interrupt() {
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
