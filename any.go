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

type checkAny[T any] struct {
	ctx *Context
	t   *T
}

func (c *checkAny[T]) success() setMsg[*checkAny[T]] {
	return setMsg[*checkAny[T]]{t: c}
}

func (c *checkAny[T]) success_() setMsgOrDefault[*checkAny[T]] {
	return setMsgOrDefault[*checkAny[T]]{setMsg: c.success()}
}

func (c *checkAny[T]) fail() setMsg[*checkAny[T]] {
	c.ctx.wronged = true
	return setMsg[*checkAny[T]]{ctx: c.ctx, t: c}
}

func (c *checkAny[T]) fail_(k defaultMsgKey) setMsgOrDefault[*checkAny[T]] {
	return setMsgOrDefault[*checkAny[T]]{setMsg: c.fail(), k: k}
}

func (c *checkAny[T]) NotNil() setMsgOrDefault[*checkAny[T]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.t == nil {
		return c.fail_(default_msg_any_notnil)
	}
	return c.success_()
}

func (c *checkAny[T]) Custom(custom func(t T) bool) setMsg[*checkAny[T]] {
	return c.Custom_(custom, false)
}

func (c *checkAny[T]) Custom_(custom func(t T) bool, omitNil bool) setMsg[*checkAny[T]] {
	if c.ctx.interrupt() {
		return c.success()
	}
	if c.t == nil {
		if !omitNil {
			return c.fail()
		}
	} else {
		ok := custom(*c.t)
		if !ok {
			return c.fail()
		}
	}
	return c.success()
}
