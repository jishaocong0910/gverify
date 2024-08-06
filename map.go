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

type checkMap[K comparable, V any] struct {
	ctx *Context
	m   map[K]V
}

func (c *checkMap[K, V]) success() msg[*checkMap[K, V]] {
	return msg[*checkMap[K, V]]{t: c}
}

func (c *checkMap[K, V]) success_() msg_[*checkMap[K, V]] {
	return msg_[*checkMap[K, V]]{msg: c.success()}
}

func (c *checkMap[K, V]) fail(confines ...[]string) msg[*checkMap[K, V]] {
	c.ctx.wronged = true
	for _, cs := range confines {
		c.ctx.confines = append(c.ctx.confines, cs...)
	}
	return msg[*checkMap[K, V]]{ctx: c.ctx, t: c}
}

func (c *checkMap[K, V]) fail_(k defaultMsgKey, confines ...[]string) msg_[*checkMap[K, V]] {
	return msg_[*checkMap[K, V]]{msg: c.fail(confines...), k: k}
}

func (c *checkMap[K, V]) NotNil() msg_[*checkMap[K, V]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.m == nil {
		return c.fail_(default_msg_map_notnil)
	}
	return c.success_()
}

func (c *checkMap[K, V]) NotEmpty() msg_[*checkMap[K, V]] {
	return c.NotEmpty_(false)
}

func (c *checkMap[K, V]) NotEmpty_(omitNil bool) msg_[*checkMap[K, V]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.m == nil {
		if !omitNil {
			return c.fail_(default_msg_map_notempty)
		}
	} else if len(c.m) == 0 {
		return c.fail_(default_msg_map_notempty)
	}
	return c.success_()
}

func (c *checkMap[K, V]) Length(l int) msg_[*checkMap[K, V]] {
	return c.Length_(l, false)
}

func (c *checkMap[K, V]) Length_(l int, omitNil bool) msg_[*checkMap[K, V]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.m == nil {
		if !omitNil {
			return c.fail_(default_msg_map_length, c.ctx.intConfines(l))
		}
	} else {
		if len(c.m) != l {
			return c.fail_(default_msg_map_length, c.ctx.intConfines(l))
		}
	}
	return c.success_()
}

func (c *checkMap[K, V]) Min(min int) msg_[*checkMap[K, V]] {
	return c.Min_(min, false)
}

func (c *checkMap[K, V]) Min_(min int, omitNil bool) msg_[*checkMap[K, V]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.m == nil {
		if !omitNil {
			return c.fail_(default_msg_map_min, c.ctx.intConfines(min))
		}
	} else if len(c.m) < min {
		return c.fail_(default_msg_map_min, c.ctx.intConfines(min))
	}
	return c.success_()
}

func (c *checkMap[K, V]) Max(max int) msg_[*checkMap[K, V]] {
	return c.Max_(max, false)
}

func (c *checkMap[K, V]) Max_(max int, omitNil bool) msg_[*checkMap[K, V]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.m == nil {
		if !omitNil {
			return c.fail_(default_msg_map_max, c.ctx.intConfines(max))
		}
	} else if len(c.m) > max {
		return c.fail_(default_msg_map_max, c.ctx.intConfines(max))
	}
	return c.success_()
}

func (c *checkMap[K, V]) Range(min, max int) msg_[*checkMap[K, V]] {
	return c.Range_(min, max, false)
}

func (c *checkMap[K, V]) Range_(min, max int, omitNil bool) msg_[*checkMap[K, V]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.m == nil {
		if !omitNil {
			return c.fail_(default_msg_map_range, c.ctx.intConfines(min, max))
		}
	} else if len(c.m) < min {
		return c.fail_(default_msg_map_range, c.ctx.intConfines(min, max))
	} else if len(c.m) > max {
		return c.fail_(default_msg_map_range, c.ctx.intConfines(min, max))
	}
	return c.success_()
}

func (c *checkMap[k, v]) Gt(min int) msg_[*checkMap[k, v]] {
	return c.Gt_(min, false)
}

func (c *checkMap[k, v]) Gt_(min int, omitNil bool) msg_[*checkMap[k, v]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.m == nil {
		if !omitNil {
			return c.fail_(default_msg_map_gt, c.ctx.intConfines(min))
		}
	} else if len(c.m) <= min {
		return c.fail_(default_msg_map_gt, c.ctx.intConfines(min))
	}
	return c.success_()
}

func (c *checkMap[k, v]) Lt(max int) msg_[*checkMap[k, v]] {
	return c.Lt_(max, false)
}

func (c *checkMap[k, v]) Lt_(max int, omitNil bool) msg_[*checkMap[k, v]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.m == nil {
		if !omitNil {
			return c.fail_(default_msg_map_lt, c.ctx.intConfines(max))
		}
	} else if len(c.m) >= max {
		return c.fail_(default_msg_map_lt, c.ctx.intConfines(max))
	}
	return c.success_()
}

func (c *checkMap[k, v]) Within(min, max int) msg_[*checkMap[k, v]] {
	return c.Within_(min, max, false)
}

func (c *checkMap[k, v]) Within_(min, max int, omitNil bool) msg_[*checkMap[k, v]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.m == nil {
		if !omitNil {
			return c.fail_(default_msg_map_within, c.ctx.intConfines(min, max))
		}
	} else if len(c.m) <= min {
		return c.fail_(default_msg_map_within, c.ctx.intConfines(min, max))
	} else if len(c.m) >= max {
		return c.fail_(default_msg_map_within, c.ctx.intConfines(min, max))
	}
	return c.success_()
}

func (c *checkMap[K, V]) Dive(key func(k K), value func(v V)) {
	if c.m != nil {
		for k, v := range c.m {
			if key != nil {
				c.ctx.diveElem(c.ctx.currentFieldName+"$key", func() {
					key(k)
				})
			}
			if value != nil {
				c.ctx.diveElem(c.ctx.currentFieldName+"$value", func() {
					value(v)
				})
			}
		}
	}
}
