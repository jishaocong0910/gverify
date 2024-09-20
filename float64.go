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

type checkFloat64 struct {
	*Context
	f *float64
}

func (c *checkFloat64) success() msg[*checkFloat64] {
	return msg[*checkFloat64]{t: c}
}

func (c *checkFloat64) success_() msg_[*checkFloat64] {
	return msg_[*checkFloat64]{msg: c.success()}
}

func (c *checkFloat64) fail(confines ...[]string) msg[*checkFloat64] {
	c.wronged = true
	for _, cs := range confines {
		c.confines = append(c.confines, cs...)
	}
	return msg[*checkFloat64]{ctx: c.Context, t: c}
}

func (c *checkFloat64) fail_(k defaultMsgKey, confines ...[]string) msg_[*checkFloat64] {
	return msg_[*checkFloat64]{msg: c.fail(confines...), k: k}
}

func (c *checkFloat64) NotNil() msg_[*checkFloat64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		return c.fail_(default_msg_float64_notnil)
	}
	return c.success_()
}

func (c *checkFloat64) Min(min float64) msg_[*checkFloat64] {
	return c.Min_(min, false)
}

func (c *checkFloat64) Min_(min float64, omitNil bool) msg_[*checkFloat64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail_(default_msg_float64_min, float64ToConfines(min))
		}
	} else if *c.f < min {
		return c.fail_(default_msg_float64_min, float64ToConfines(min))
	}
	return c.success_()
}

func (c *checkFloat64) Max(max float64) msg_[*checkFloat64] {
	return c.Max_(max, false)
}

func (c *checkFloat64) Max_(max float64, omitNil bool) msg_[*checkFloat64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail_(default_msg_float64_max, float64ToConfines(max))
		}
	} else if *c.f > max {
		return c.fail_(default_msg_float64_max, float64ToConfines(max))
	}
	return c.success_()
}

func (c *checkFloat64) Range(min, max float64) msg_[*checkFloat64] {
	return c.Range_(min, max, false)
}

func (c *checkFloat64) Range_(min, max float64, omitNil bool) msg_[*checkFloat64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail_(default_msg_float64_range, float64ToConfines(min, max))
		}
	} else if *c.f < min {
		return c.fail_(default_msg_float64_range, float64ToConfines(min, max))
	} else if *c.f > max {
		return c.fail_(default_msg_float64_range, float64ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkFloat64) Gt(min float64) msg_[*checkFloat64] {
	return c.Gt_(min, false)
}

func (c *checkFloat64) Gt_(min float64, omitNil bool) msg_[*checkFloat64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail_(default_msg_float64_gt, float64ToConfines(min))
		}
	} else if *c.f <= min {
		return c.fail_(default_msg_float64_gt, float64ToConfines(min))
	}
	return c.success_()
}

func (c *checkFloat64) Lt(max float64) msg_[*checkFloat64] {
	return c.Lt_(max, false)
}

func (c *checkFloat64) Lt_(max float64, omitNil bool) msg_[*checkFloat64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail_(default_msg_float64_lt, float64ToConfines(max))
		}
	} else if *c.f >= max {
		return c.fail_(default_msg_float64_lt, float64ToConfines(max))
	}
	return c.success_()
}

func (c *checkFloat64) Within(min, max float64) msg_[*checkFloat64] {
	return c.Within_(min, max, false)
}

func (c *checkFloat64) Within_(min, max float64, omitNil bool) msg_[*checkFloat64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail_(default_msg_float64_within, float64ToConfines(min, max))
		}
	} else if *c.f <= min {
		return c.fail_(default_msg_float64_within, float64ToConfines(min, max))
	} else if *c.f >= max {
		return c.fail_(default_msg_float64_within, float64ToConfines(min, max))
	}
	return c.success_()
}

func (c *checkFloat64) Options(options []float64) msg_[*checkFloat64] {
	return c.Options_(options, false)
}

func (c *checkFloat64) Options_(options []float64, omitNil bool) msg_[*checkFloat64] {
	if c.interrupt() {
		return c.success_()
	}
	if c.f == nil {
		if !omitNil {
			return c.fail_(default_msg_float64_options, float64ToConfines(options...))
		}
	} else {
		match := false
		for _, o := range options {
			if *c.f == o {
				match = true
			}
		}
		if !match {
			return c.fail_(default_msg_float64_options, float64ToConfines(options...))
		}
	}
	return c.success_()
}

func (c *checkFloat64) Custom(custom func(f float64) bool) msg[*checkFloat64] {
	return c.Custom_(custom, false)
}

func (c *checkFloat64) Custom_(custom func(f float64) bool, omitNil bool) msg[*checkFloat64] {
	if c.interrupt() {
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
