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

type checkStruct[V Verifiable] struct {
	ctx *Context
	v   *V
}

func (c *checkStruct[V]) success() setMsg[*checkStruct[V]] {
	return setMsg[*checkStruct[V]]{t: c}
}

func (c *checkStruct[V]) success_() setMsgOrDefault[*checkStruct[V]] {
	return setMsgOrDefault[*checkStruct[V]]{setMsg: c.success()}
}

func (c *checkStruct[V]) fail() setMsg[*checkStruct[V]] {
	c.ctx.wronged = true
	return setMsg[*checkStruct[V]]{ctx: c.ctx, t: c}
}

func (c *checkStruct[V]) fail_(k defaultMsgKey) setMsgOrDefault[*checkStruct[V]] {
	return setMsgOrDefault[*checkStruct[V]]{setMsg: c.fail(), k: k}
}

func (c *checkStruct[V]) NotNil() setMsgOrDefault[*checkStruct[V]] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.v == nil {
		return c.fail_(default_msg_struct_notnil)
	}
	return c.success_()
}

func (c *checkStruct[V]) Dive() {
	if c.v != nil {
		s := c.ctx.savepoint()
		c.ctx.beforeDive(dive_struct, "", ".", 0)
		(*c.v).Checklist(c.ctx)
		c.ctx.afterDive(s)
	}
}
