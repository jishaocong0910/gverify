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

import "strconv"

type checkSlices[T any] struct {
	ctx *Context
	s   []T
}

func (c *checkSlices[T]) success() setMsg[*checkSlices[T]] {
	return setMsg[*checkSlices[T]]{t: c}
}

func (c *checkSlices[T]) success_() setMsgOrDefault[*checkSlices[T]] {
	return setMsgOrDefault[*checkSlices[T]]{setMsg: c.success()}
}

func (c *checkSlices[T]) fail(confines ...[]string) setMsg[*checkSlices[T]] {
	c.ctx.wronged = true
	for _, cs := range confines {
		c.ctx.confines = append(c.ctx.confines, cs...)
	}
	return setMsg[*checkSlices[T]]{ctx: c.ctx, t: c}
}

func (c *checkSlices[T]) fail_(k defaultMsgKey, confines ...[]string) setMsgOrDefault[*checkSlices[T]] {
	return setMsgOrDefault[*checkSlices[T]]{setMsg: c.fail(confines...), k: k}
}

func (c *checkSlices[T]) NotNil() setMsgOrDefault[*checkSlices[T]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		return c.fail_(default_msg_slices_notnil)
	}
	return c.success_()
}

func (c *checkSlices[T]) NotEmpty() setMsgOrDefault[*checkSlices[T]] {
	return c.NotEmpty_(false)
}

func (c *checkSlices[T]) NotEmpty_(omitNil bool) setMsgOrDefault[*checkSlices[T]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_slices_notempty)
		}
	} else if len(c.s) == 0 {
		return c.fail_(default_msg_slices_notempty)
	}
	return c.success_()
}

func (c *checkSlices[T]) Length(l int) setMsgOrDefault[*checkSlices[T]] {
	return c.Length_(l, false)
}

func (c *checkSlices[T]) Length_(l int, omitNil bool) setMsgOrDefault[*checkSlices[T]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_slices_length, intToConfines(l))
		}
	} else {
		if len(c.s) != l {
			return c.fail_(default_msg_slices_length, intToConfines(l))
		}
	}
	return c.success_()
}

func (c *checkSlices[T]) Min(min int) setMsgOrDefault[*checkSlices[T]] {
	return c.Min_(min, false)
}

func (c *checkSlices[T]) Min_(min int, omitNil bool) setMsgOrDefault[*checkSlices[T]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_slices_min, intToConfines(min))
		}
	} else if len(c.s) < min {
		return c.fail_(default_msg_slices_min, intToConfines(min))
	}
	return c.success_()
}

func (c *checkSlices[T]) Max(max int) setMsgOrDefault[*checkSlices[T]] {
	return c.Max_(max, false)
}

func (c *checkSlices[T]) Max_(max int, omitNil bool) setMsgOrDefault[*checkSlices[T]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_slices_max, intToConfines(max))
		}
	} else if len(c.s) > max {
		return c.fail_(default_msg_slices_max, intToConfines(max))
	}
	return c.success_()
}

func (c *checkSlices[T]) Range(min, max int) setMsgOrDefault[*checkSlices[T]] {
	return c.Range_(min, max, false)
}

func (c *checkSlices[T]) Range_(min, max int, omitNil bool) setMsgOrDefault[*checkSlices[T]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_slices_range, intToConfines(min, max))
		}
	} else if len(c.s) < min {
		return c.fail_(default_msg_slices_range, intToConfines(min, max))
	} else if len(c.s) > max {
		return c.fail_(default_msg_slices_range, intToConfines(min, max))
	}
	return c.success_()
}

func (c *checkSlices[t]) Gt(min int) setMsgOrDefault[*checkSlices[t]] {
	return c.Gt_(min, false)
}

func (c *checkSlices[t]) Gt_(min int, omitNil bool) setMsgOrDefault[*checkSlices[t]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_slices_gt, intToConfines(min))
		}
	} else if len(c.s) <= min {
		return c.fail_(default_msg_slices_gt, intToConfines(min))
	}
	return c.success_()
}

func (c *checkSlices[t]) Lt(max int) setMsgOrDefault[*checkSlices[t]] {
	return c.Lt_(max, false)
}

func (c *checkSlices[t]) Lt_(max int, omitNil bool) setMsgOrDefault[*checkSlices[t]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_slices_lt, intToConfines(max))
		}
	} else if len(c.s) >= max {
		return c.fail_(default_msg_slices_lt, intToConfines(max))
	}
	return c.success_()
}

func (c *checkSlices[t]) Within(min, max int) setMsgOrDefault[*checkSlices[t]] {
	return c.Within_(min, max, false)
}

func (c *checkSlices[t]) Within_(min, max int, omitNil bool) setMsgOrDefault[*checkSlices[t]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.s == nil {
		if !omitNil {
			return c.fail_(default_msg_slices_within, intToConfines(min, max))
		}
	} else if len(c.s) <= min {
		return c.fail_(default_msg_slices_within, intToConfines(min, max))
	} else if len(c.s) >= max {
		return c.fail_(default_msg_slices_within, intToConfines(min, max))
	}
	return c.success_()
}

func (c *checkSlices[T]) Dive(f func(e T)) {
	if c.s != nil && f != nil {
		s := c.ctx.savepoint()
		for i, t := range c.s {
			c.ctx.beforeDive(dive_slices, "["+strconv.Itoa(i)+"]", "", i)
			f(t)
			c.ctx.afterDive(s)
			if c.ctx.interrupt() {
				break
			}
		}
	}
}
