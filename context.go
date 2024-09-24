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
	"context"
	"fmt"
	"strconv"
	"strings"
)

type diveType int

const (
	dive_struct diveType = iota + 1
	dive_slices
	dive_map
)

type Context struct {
	context.Context
	code            string
	all             bool
	wronged         bool
	msgs            []string
	confines        []string
	fieldNamePrefix string
	fieldName       string
	nameSeparator   string
	diveType        diveType
	index           int
}

func (c *Context) FieldName() string {
	return c.fieldName
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

func (c *Context) Index() int {
	i := -1
	if c.diveType == dive_slices {
		i = c.index
	}
	return i
}

func (c *Context) interrupt() bool {
	return c.wronged && !c.all
}

func (c *Context) reset(fieldName string) *Context {
	if c.diveType != dive_slices && c.diveType != dive_map || fieldName != "" {
		c.fieldName = c.fieldNamePrefix + c.nameSeparator + fieldName
	}
	c.confines = nil
	return c
}

func (c *Context) addMsg(msg string, args ...any) {
	c.msgs = append(c.msgs, fmt.Sprintf(msg, args...))
}

func (c *Context) savepoint() savepoint {
	return savepoint{
		fieldNamePrefix: c.fieldNamePrefix,
		fieldName:       c.fieldName,
		nameSeparator:   c.nameSeparator,
		diveType:        c.diveType,
		index:           c.index,
	}
}

func (c *Context) beforeDive(diveType diveType, defaultFieldName, nameSeparator string, index int) {
	c.fieldNamePrefix = c.fieldName
	c.nameSeparator = nameSeparator
	c.reset(defaultFieldName)
	c.diveType = diveType
	c.index = index
}

func (c *Context) afterDive(s savepoint) {
	c.fieldNamePrefix = s.fieldNamePrefix
	c.fieldName = s.fieldName
	c.nameSeparator = s.nameSeparator
	c.diveType = s.diveType
	c.index = s.index
}

type savepoint struct {
	fieldNamePrefix string
	fieldName       string
	nameSeparator   string
	diveType        diveType
	index           int
}

func byteToConfines(confines ...byte) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func intToConfines(confines ...int) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(int64(confine), 10))
	}
	return result
}

func int8ToConfines(confines ...int8) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(int64(confine), 10))
	}
	return result
}

func int16ToConfines(confines ...int16) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(int64(confine), 10))
	}
	return result
}

func int32ToConfines(confines ...int32) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(int64(confine), 10))
	}
	return result
}

func int64ToConfines(confines ...int64) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatInt(confine, 10))
	}
	return result
}

func uintToConfines(confines ...uint) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func uint8ToConfines(confines ...uint8) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func uint16ToConfines(confines ...uint16) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func uint32ToConfines(confines ...uint32) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func uint64ToConfines(confines ...uint64) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatUint(uint64(confine), 10))
	}
	return result
}

func float32ToConfines(confines ...float32) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatFloat(float64(confine), 'f', -1, 32))
	}
	return result
}

func float64ToConfines(confines ...float64) []string {
	result := make([]string, 0, len(confines))
	for _, confine := range confines {
		result = append(result, strconv.FormatFloat(confine, 'f', -1, 32))
	}
	return result
}
