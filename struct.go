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

type checkStruct[T Verifiable] struct {
	ctx *Context
	t   *T
}

func (c *checkStruct[T]) success() msg[*checkStruct[T]] {
	return msg[*checkStruct[T]]{t: c}
}

func (c *checkStruct[T]) success_() msg_[*checkStruct[T]] {
	return msg_[*checkStruct[T]]{msg: c.success()}
}

func (c *checkStruct[T]) fail() msg[*checkStruct[T]] {
	c.ctx.wronged = true
	return msg[*checkStruct[T]]{ctx: c.ctx, t: c}
}

func (c *checkStruct[T]) fail_(k defaultMsgKey) msg_[*checkStruct[T]] {
	return msg_[*checkStruct[T]]{msg: c.fail(), k: k}
}

func (c *checkStruct[T]) NotNil() msg_[*checkStruct[T]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.t == nil {
		return c.fail_(default_msg_struct_notnil)
	}
	return c.success_()
}

func (c *checkStruct[T]) Dive() {
	if c.t != nil {
		c.ctx.diveStruct(c.ctx.currentFieldName, func() {
			(*c.t).Checklist(c.ctx)
		})
	}
}
