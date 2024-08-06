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

func NewDefaultContext() *Context {
	return &Context{all: false}
}

func NewCheckAllContext() *Context {
	return &Context{all: true}
}

func GetResult(ctx *Context) (bool, string, []string) {
	msg := ""
	if len(ctx.msgs) > 0 {
		msg = ctx.msgs[0]
	}
	return !ctx.wronged, msg, ctx.msgs
}

func SetChecklistFalse(c *Context) {
	c.wronged = true
}
