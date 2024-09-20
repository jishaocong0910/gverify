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

type checkInt64 struct {
	*Context
	i *int64
}

func (c *checkInt64) success() msg[*checkInt64] {
	return msg[*checkInt64]{t: c}
}

func (c *checkInt64) success_() msg_[*checkInt64] {
	return msg_[*checkInt64]{msg: c.success()}
}

func (c *checkInt64) fail(confines ...[]string) msg[*checkInt64] {
	c.wronged = true
	for _, cs := range confines {
		c.confines = append(c.confines, cs...)
	}
	return msg[*checkInt64]{ctx: c.Context, t: c}
}

func (c *checkInt64) fail_(k defaultMsgKey, confines ...[]string) msg_[*checkInt64] {
	return msg_[*checkInt64]{msg: c.fail(confines...), k: k}
}

func (c *checkInt64) NotNil() msg_[*checkInt64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		return c.fail_(default_msg_int64_notnil)
	}
	return c.success_()
}

func (c *checkInt64) Min(min int64) msg_[*checkInt64] {
	return c.Min_(min, false)
}

func (c *checkInt64) Min_(min int64, omitNil bool) msg_[*checkInt64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int64_min, int64ToConfines(min))
		}
	} else if *c.i < min {
		return c.fail_(default_msg_int64_min, int64ToConfines(min))
	}
	return c.success_()
}

func (c *checkInt64) Max(max int64) msg_[*checkInt64] {
	return c.Max_(max, false)
}

func (c *checkInt64) Max_(max int64, omitNil bool) msg_[*checkInt64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int64_max, int64ToConfines(max))
		}
	} else if *c.i > max {
		return c.fail_(default_msg_int64_max, int64ToConfines(max))
	}
	return c.success_()
}

func (c *checkInt64) Range(min, max int64) msg_[*checkInt64] {
	return c.Range_(min, max, false)
}

func (c *checkInt64) Range_(min, max int64, omitNil bool) msg_[*checkInt64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int64_range, int64ToConfines(min, max))
		}
	} else if *c.i < min {
		return c.fail_(default_msg_int64_range, int64ToConfines(min, max))
	} else if *c.i > max {
		return c.fail_(default_msg_int64_range, int64ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkInt64) Gt(min int64) msg_[*checkInt64] {
	return c.Gt_(min, false)
}

func (c *checkInt64) Gt_(min int64, omitNil bool) msg_[*checkInt64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int64_gt, int64ToConfines(min))
		}
	} else if *c.i <= min {
		return c.fail_(default_msg_int64_gt, int64ToConfines(min))
	}
	return c.success_()
}

func (c *checkInt64) Lt(max int64) msg_[*checkInt64] {
	return c.Lt_(max, false)
}

func (c *checkInt64) Lt_(max int64, omitNil bool) msg_[*checkInt64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int64_lt, int64ToConfines(max))
		}
	} else if *c.i >= max {
		return c.fail_(default_msg_int64_lt, int64ToConfines(max))
	}
	return c.success_()
}

func (c *checkInt64) Within(min, max int64) msg_[*checkInt64] {
	return c.Within_(min, max, false)
}

func (c *checkInt64) Within_(min, max int64, omitNil bool) msg_[*checkInt64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int64_within, int64ToConfines(min, max))
		}
	} else if *c.i <= min {
		return c.fail_(default_msg_int64_within, int64ToConfines(min, max))
	} else if *c.i >= max {
		return c.fail_(default_msg_int64_within, int64ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkInt64) Options(options []int64) msg_[*checkInt64] {
	return c.Options_(options, false)
}

func (c *checkInt64) Options_(options []int64, omitNil bool) msg_[*checkInt64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.i == nil {
		if !omitNil {
			return c.fail_(default_msg_int64_options, int64ToConfines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.i == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_int64_options, int64ToConfines(options...))
		}
	}
	return c.success_()
}

func (c *checkInt64) Custom(custom func(i int64) bool) msg[*checkInt64] {
	return c.Custom_(custom, false)
}

func (c *checkInt64) Custom_(custom func(i int64) bool, omitNil bool) msg[*checkInt64] {
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
