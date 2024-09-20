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

type checkInt32 struct {
	*Context
	i *int32
}

func (c *checkInt32) success() msg[*checkInt32] {
	return msg[*checkInt32]{t: c}
}

func (c *checkInt32) success_() msg_[*checkInt32] {
	return msg_[*checkInt32]{msg: c.success()}
}

func (c *checkInt32) fail(confines ...[]string) msg[*checkInt32] {
	c.wronged = true
	for _, cs := range confines {
		c.confines = append(c.confines, cs...)
	}
	return msg[*checkInt32]{ctx: c.Context, t: c}
}

func (c *checkInt32) fail_(k defaultMsgKey, confines ...[]string) msg_[*checkInt32] {
	return msg_[*checkInt32]{msg: c.fail(confines...), k: k}
}

func (c *checkInt32) NotNil() msg_[*checkInt32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		return c.fail_(default_msg_int32_notnil)
	}
	return c.success_()
}

func (c *checkInt32) Min(min int32) msg_[*checkInt32] {
	return c.Min_(min, false)
}

func (c *checkInt32) Min_(min int32, omitNil bool) msg_[*checkInt32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int32_min, int32ToConfines(min))
		}
	} else if *c.i < min {
		return c.fail_(default_msg_int32_min, int32ToConfines(min))
	}
	return c.success_()
}

func (c *checkInt32) Max(max int32) msg_[*checkInt32] {
	return c.Max_(max, false)
}

func (c *checkInt32) Max_(max int32, omitNil bool) msg_[*checkInt32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int32_max, int32ToConfines(max))
		}
	} else if *c.i > max {
		return c.fail_(default_msg_int32_max, int32ToConfines(max))
	}
	return c.success_()
}

func (c *checkInt32) Range(min, max int32) msg_[*checkInt32] {
	return c.Range_(min, max, false)
}

func (c *checkInt32) Range_(min, max int32, omitNil bool) msg_[*checkInt32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int32_range, int32ToConfines(min, max))
		}
	} else if *c.i < min {
		return c.fail_(default_msg_int32_range, int32ToConfines(min, max))
	} else if *c.i > max {
		return c.fail_(default_msg_int32_range, int32ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkInt32) Gt(min int32) msg_[*checkInt32] {
	return c.Gt_(min, false)
}

func (c *checkInt32) Gt_(min int32, omitNil bool) msg_[*checkInt32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int32_gt, int32ToConfines(min))
		}
	} else if *c.i <= min {
		return c.fail_(default_msg_int32_gt, int32ToConfines(min))
	}
	return c.success_()
}

func (c *checkInt32) Lt(max int32) msg_[*checkInt32] {
	return c.Lt_(max, false)
}

func (c *checkInt32) Lt_(max int32, omitNil bool) msg_[*checkInt32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int32_lt, int32ToConfines(max))
		}
	} else if *c.i >= max {
		return c.fail_(default_msg_int32_lt, int32ToConfines(max))
	}
	return c.success_()
}

func (c *checkInt32) Within(min, max int32) msg_[*checkInt32] {
	return c.Within_(min, max, false)
}

func (c *checkInt32) Within_(min, max int32, omitNil bool) msg_[*checkInt32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int32_within, int32ToConfines(min, max))
		}
	} else if *c.i <= min {
		return c.fail_(default_msg_int32_within, int32ToConfines(min, max))
	} else if *c.i >= max {
		return c.fail_(default_msg_int32_within, int32ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkInt32) Options(options []int32) msg_[*checkInt32] {
	return c.Options_(options, false)
}

func (c *checkInt32) Options_(options []int32, omitNil bool) msg_[*checkInt32] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int32_options, int32ToConfines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.i == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_int32_options, int32ToConfines(options...))
		}
	}
	return c.success_()
}

func (c *checkInt32) Custom(custom func(i int32) bool) msg[*checkInt32] {
	return c.Custom_(custom, false)
}

func (c *checkInt32) Custom_(custom func(i int32) bool, omitNil bool) msg[*checkInt32] {
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
