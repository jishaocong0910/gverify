/*
Copyright 2024-present jishaocong0910

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vfy

// 校验内嵌结构体
type checkEmbed[V Verifiable] struct {
	vc *VContext
	s  *V
}

// 校验下沉内嵌结构体中的字段。内嵌结构体没有其他规则方法，也没有选项，所以改方法不导出，创建 checkEmbed 时直接调用
func (c *checkEmbed[V]) dive() {
	if c.s != nil {
		(*c.s).Checklist(c.vc)
	}
}
