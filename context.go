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

import (
	"fmt"
	"strconv"
	"strings"
)

type Context struct {
	all              bool
	wronged          bool
	msgs             []string
	parentFieldNames []string
	currentFieldName string
	isDiveElem       bool
	confines         []string
}

func (c *Context) FieldName() string {
	if len(c.parentFieldNames) == 0 {
		return c.currentFieldName
	}
	var builder strings.Builder
	for _, pn := range c.parentFieldNames {
		builder.WriteString(pn)
		builder.WriteString(".")
	}
	builder.WriteString(c.currentFieldName)
	return builder.String()
}

func (c *Context) Confine(i int) string {
	confine := ""
	if i < len(c.confines) {
		confine = c.confines[i]
	}
	return confine
}

func (c *Context) Confines() string {
	var builder strings.Builder
	l := len(c.confines)
	if l > 0 {
		builder.WriteString(c.confines[0])
	}
	for i := 1; i < l-1; i++ {
		builder.WriteString(", ")
		builder.WriteString(c.confines[i])
	}
	if l > 1 {
		builder.WriteString(" or ")
		builder.WriteString(c.confines[l-1])
	}
	return builder.String()
}

func (c *Context) interrupt() bool {
	return c.wronged && !c.all
}

func (c *Context) setCurrentFieldName(fieldName string) *Context {
	if !c.isDiveElem {
		c.currentFieldName = fieldName
	}
	c.confines = nil
	return c
}

func (c *Context) addMsg(msg string, args ...any) {
	c.msgs = append(c.msgs, fmt.Sprintf(msg, args...))
}

func (c *Context) diveElem(currentName string, f func()) {
	n := c.currentFieldName
	i := c.isDiveElem
	c.currentFieldName = currentName
	c.isDiveElem = true
	f()
	c.currentFieldName = n
	c.isDiveElem = i
}

func (c *Context) diveStruct(parentName string, f func()) {
	n := c.currentFieldName
	i := c.isDiveElem
	c.isDiveElem = false
	c.parentFieldNames = append(c.parentFieldNames, parentName)
	f()
	c.currentFieldName = n
	c.isDiveElem = i
	c.parentFieldNames = c.parentFieldNames[:len(c.parentFieldNames)-1]

}

func (c *Context) byteConfines(confines ...byte) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func (c *Context) intConfines(confines ...int) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(int64(confine), 10))
	}
	return result
}

func (c *Context) int8Confines(confines ...int8) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(int64(confine), 10))
	}
	return result
}

func (c *Context) int16Confines(confines ...int16) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(int64(confine), 10))
	}
	return result
}

func (c *Context) int32Confines(confines ...int32) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(int64(confine), 10))
	}
	return result
}

func (c *Context) int64Confines(confines ...int64) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(confine, 10))
	}
	return result
}

func (c *Context) uintConfines(confines ...uint) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func (c *Context) uint8Confines(confines ...uint8) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func (c *Context) uint16Confines(confines ...uint16) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func (c *Context) uint32Confines(confines ...uint32) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func (c *Context) uint64Confines(confines ...uint64) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func (c *Context) float32Confines(confines ...float32) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatFloat(float64(confine), 'f', -1, 32))
	}
	return result
}

func (c *Context) float64Confines(confines ...float64) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatFloat(confine, 'f', -1, 32))
	}
	return result
}
