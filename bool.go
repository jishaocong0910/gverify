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

type checkBool struct {
	ctx *Context
	b   *bool
}

func (c *checkBool) success() setMsg[*checkBool] {
	return setMsg[*checkBool]{t: c}
}

func (c *checkBool) success_() setMsgOrDefault[*checkBool] {
	return setMsgOrDefault[*checkBool]{setMsg: c.success()}
}

func (c *checkBool) fail() setMsg[*checkBool] {
	c.ctx.wronged = true
	return setMsg[*checkBool]{ctx: c.ctx, t: c}
}

func (c *checkBool) fail_(k defaultMsgKey) setMsgOrDefault[*checkBool] {
	return setMsgOrDefault[*checkBool]{setMsg: c.fail(), k: k}
}

func (c *checkBool) NotNil() setMsgOrDefault[*checkBool] {
	if c.ctx.interrupt() {
		return c.success_()
	}
	if c.b == nil {
		return c.fail_(default_msg_bool_notnil)
	}
	return c.success_()
}

func (c *checkBool) Custom(custom func(b bool) bool) setMsg[*checkBool] {
	return c.Custom_(custom, false)
}

func (c *checkBool) Custom_(custom func(b bool) bool, omitNil bool) setMsg[*checkBool] {
	if c.ctx.interrupt() {
		return c.success()
	}
	if c.b == nil {
		if !omitNil {
			return c.fail()
		}
	} else {
		ok := custom(*c.b)
		if !ok {
			return c.fail()
		}
	}
	return c.success()
}
