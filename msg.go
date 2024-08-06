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

type msg[T any] struct {
	ctx *Context
	t   T
}

func (m msg[T]) Msg(msg string, args ...any) T {
	if m.ctx != nil {
		m.ctx.addMsg(msg, args...)
	}
	return m.t
}

type msg_[T any] struct {
	msg[T]
	k defaultMsgKey
}

func (m msg_[T]) DefaultMsg() T {
	if m.ctx != nil {
		var dm string
		if f, ok := defaultMsgs[m.k]; ok {
			dm = f(m.ctx)
		}
		m.ctx.addMsg(dm)
	}
	return m.t
}

type defaultMsgKey int

const (
	default_msg_bool_notnil defaultMsgKey = iota + 1
	default_msg_byte_notnil
	default_msg_byte_min
	default_msg_byte_max
	default_msg_byte_range
	default_msg_byte_gt
	default_msg_byte_lt
	default_msg_byte_within
	default_msg_byte_options
	default_msg_int_notnil
	default_msg_int_min
	default_msg_int_max
	default_msg_int_range
	default_msg_int_gt
	default_msg_int_lt
	default_msg_int_within
	default_msg_int_options
	default_msg_int8_notnil
	default_msg_int8_min
	default_msg_int8_max
	default_msg_int8_range
	default_msg_int8_gt
	default_msg_int8_lt
	default_msg_int8_within
	default_msg_int8_options
	default_msg_int16_notnil
	default_msg_int16_min
	default_msg_int16_max
	default_msg_int16_range
	default_msg_int16_gt
	default_msg_int16_lt
	default_msg_int16_within
	default_msg_int16_options
	default_msg_int32_notnil
	default_msg_int32_min
	default_msg_int32_max
	default_msg_int32_range
	default_msg_int32_gt
	default_msg_int32_lt
	default_msg_int32_within
	default_msg_int32_options
	default_msg_int64_notnil
	default_msg_int64_min
	default_msg_int64_max
	default_msg_int64_range
	default_msg_int64_gt
	default_msg_int64_lt
	default_msg_int64_within
	default_msg_int64_options
	default_msg_uint_notnil
	default_msg_uint_min
	default_msg_uint_max
	default_msg_uint_range
	default_msg_uint_gt
	default_msg_uint_lt
	default_msg_uint_within
	default_msg_uint_options
	default_msg_uint8_notnil
	default_msg_uint8_min
	default_msg_uint8_max
	default_msg_uint8_range
	default_msg_uint8_gt
	default_msg_uint8_lt
	default_msg_uint8_within
	default_msg_uint8_options
	default_msg_uint16_notnil
	default_msg_uint16_min
	default_msg_uint16_max
	default_msg_uint16_range
	default_msg_uint16_gt
	default_msg_uint16_lt
	default_msg_uint16_within
	default_msg_uint16_options
	default_msg_uint32_notnil
	default_msg_uint32_min
	default_msg_uint32_max
	default_msg_uint32_range
	default_msg_uint32_gt
	default_msg_uint32_lt
	default_msg_uint32_within
	default_msg_uint32_options
	default_msg_uint64_notnil
	default_msg_uint64_min
	default_msg_uint64_max
	default_msg_uint64_range
	default_msg_uint64_gt
	default_msg_uint64_lt
	default_msg_uint64_within
	default_msg_uint64_options
	default_msg_float32_notnil
	default_msg_float32_min
	default_msg_float32_max
	default_msg_float32_range
	default_msg_float32_gt
	default_msg_float32_lt
	default_msg_float32_within
	default_msg_float32_options
	default_msg_float64_notnil
	default_msg_float64_min
	default_msg_float64_max
	default_msg_float64_range
	default_msg_float64_gt
	default_msg_float64_lt
	default_msg_float64_within
	default_msg_float64_options
	default_msg_string_notnil
	default_msg_string_notblank
	default_msg_string_length
	default_msg_string_regex
	default_msg_string_min
	default_msg_string_max
	default_msg_string_range
	default_msg_string_gt
	default_msg_string_lt
	default_msg_string_within
	default_msg_string_options
	default_msg_slices_notnil
	default_msg_slices_notempty
	default_msg_slices_length
	default_msg_slices_min
	default_msg_slices_max
	default_msg_slices_range
	default_msg_slices_gt
	default_msg_slices_lt
	default_msg_slices_within
	default_msg_map_notnil
	default_msg_map_notempty
	default_msg_map_length
	default_msg_map_min
	default_msg_map_max
	default_msg_map_range
	default_msg_map_gt
	default_msg_map_lt
	default_msg_map_within
	default_msg_struct_notnil
	default_msg_any_notnil
)

var defaultMsgs = map[defaultMsgKey]func(ctx *Context) string{}

type defaultMsgBool struct {
}

func (d defaultMsgBool) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_bool_notnil] = f
}

type defaultMsgByte struct {
}

func (d defaultMsgByte) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_byte_notnil] = f
}

func (d defaultMsgByte) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_byte_min] = f
}

func (d defaultMsgByte) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_byte_max] = f
}

func (d defaultMsgByte) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_byte_range] = f
}

func (d defaultMsgByte) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_byte_gt] = f
}

func (d defaultMsgByte) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_byte_lt] = f
}

func (d defaultMsgByte) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_byte_within] = f
}

func (d defaultMsgByte) Options(f func(ctx *Context) string) {
	defaultMsgs[default_msg_byte_options] = f
}

type defaultMsgInt struct {
}

func (d defaultMsgInt) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int_notnil] = f
}

func (d defaultMsgInt) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int_min] = f
}

func (d defaultMsgInt) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int_max] = f
}

func (d defaultMsgInt) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int_range] = f
}

func (d defaultMsgInt) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int_gt] = f
}

func (d defaultMsgInt) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int_lt] = f
}

func (d defaultMsgInt) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int_within] = f
}

func (d defaultMsgInt) Options(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int_options] = f
}

type defaultMsgInt8 struct {
}

func (d defaultMsgInt8) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int8_notnil] = f
}

func (d defaultMsgInt8) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int8_min] = f
}

func (d defaultMsgInt8) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int8_max] = f
}

func (d defaultMsgInt8) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int8_range] = f
}

func (d defaultMsgInt8) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int8_gt] = f
}

func (d defaultMsgInt8) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int8_lt] = f
}

func (d defaultMsgInt8) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int8_within] = f
}

func (d defaultMsgInt8) Options(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int8_options] = f
}

type defaultMsgInt16 struct {
}

func (d defaultMsgInt16) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int16_notnil] = f
}

func (d defaultMsgInt16) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int16_min] = f
}

func (d defaultMsgInt16) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int16_max] = f
}

func (d defaultMsgInt16) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int16_range] = f
}

func (d defaultMsgInt16) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int16_gt] = f
}

func (d defaultMsgInt16) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int16_lt] = f
}

func (d defaultMsgInt16) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int16_within] = f
}

func (d defaultMsgInt16) Options(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int16_options] = f
}

type defaultMsgInt32 struct {
}

func (d defaultMsgInt32) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int32_notnil] = f
}

func (d defaultMsgInt32) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int32_min] = f
}

func (d defaultMsgInt32) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int32_max] = f
}

func (d defaultMsgInt32) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int32_range] = f
}

func (d defaultMsgInt32) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int32_gt] = f
}

func (d defaultMsgInt32) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int32_lt] = f
}

func (d defaultMsgInt32) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int32_within] = f
}

func (d defaultMsgInt32) Options(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int32_options] = f
}

type defaultMsgInt64 struct {
}

func (d defaultMsgInt64) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int64_notnil] = f
}

func (d defaultMsgInt64) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int64_min] = f
}

func (d defaultMsgInt64) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int64_max] = f
}

func (d defaultMsgInt64) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int64_range] = f
}

func (d defaultMsgInt64) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int64_gt] = f
}

func (d defaultMsgInt64) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int64_lt] = f
}

func (d defaultMsgInt64) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int64_within] = f
}

func (d defaultMsgInt64) Options(f func(ctx *Context) string) {
	defaultMsgs[default_msg_int64_options] = f
}

type defaultMsgUint struct {
}

func (d defaultMsgUint) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint_notnil] = f
}

func (d defaultMsgUint) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint_min] = f
}

func (d defaultMsgUint) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint_max] = f
}

func (d defaultMsgUint) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint_range] = f
}

func (d defaultMsgUint) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint_gt] = f
}

func (d defaultMsgUint) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint_lt] = f
}

func (d defaultMsgUint) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint_within] = f
}

func (d defaultMsgUint) Options(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint_options] = f
}

type defaultMsgUint8 struct {
}

func (d defaultMsgUint8) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint8_notnil] = f
}

func (d defaultMsgUint8) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint8_min] = f
}

func (d defaultMsgUint8) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint8_max] = f
}

func (d defaultMsgUint8) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint8_range] = f
}

func (d defaultMsgUint8) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint8_gt] = f
}

func (d defaultMsgUint8) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint8_lt] = f
}

func (d defaultMsgUint8) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint8_within] = f
}

func (d defaultMsgUint8) Options(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint8_options] = f
}

type defaultMsgUint16 struct {
}

func (d defaultMsgUint16) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint16_notnil] = f
}

func (d defaultMsgUint16) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint16_min] = f
}

func (d defaultMsgUint16) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint16_max] = f
}

func (d defaultMsgUint16) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint16_range] = f
}

func (d defaultMsgUint16) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint16_gt] = f
}

func (d defaultMsgUint16) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint16_lt] = f
}

func (d defaultMsgUint16) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint16_within] = f
}

func (d defaultMsgUint16) Options(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint16_options] = f
}

type defaultMsgUint32 struct {
}

func (d defaultMsgUint32) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint32_notnil] = f
}

func (d defaultMsgUint32) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint32_min] = f
}

func (d defaultMsgUint32) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint32_max] = f
}

func (d defaultMsgUint32) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint32_range] = f
}

func (d defaultMsgUint32) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint32_gt] = f
}

func (d defaultMsgUint32) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint32_lt] = f
}

func (d defaultMsgUint32) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint32_within] = f
}

func (d defaultMsgUint32) Options(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint32_options] = f
}

type defaultMsgUint64 struct {
}

func (d defaultMsgUint64) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint64_notnil] = f
}

func (d defaultMsgUint64) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint64_min] = f
}

func (d defaultMsgUint64) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint64_max] = f
}

func (d defaultMsgUint64) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint64_range] = f
}

func (d defaultMsgUint64) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint64_gt] = f
}

func (d defaultMsgUint64) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint64_lt] = f
}

func (d defaultMsgUint64) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint64_within] = f
}

func (d defaultMsgUint64) Options(f func(ctx *Context) string) {
	defaultMsgs[default_msg_uint64_options] = f
}

type defaultMsgFloat32 struct {
}

func (d defaultMsgFloat32) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float32_notnil] = f
}

func (d defaultMsgFloat32) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float32_min] = f
}

func (d defaultMsgFloat32) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float32_max] = f
}

func (d defaultMsgFloat32) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float32_range] = f
}

func (d defaultMsgFloat32) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float32_gt] = f
}

func (d defaultMsgFloat32) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float32_lt] = f
}

func (d defaultMsgFloat32) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float32_within] = f
}

func (d defaultMsgFloat32) Options(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float32_options] = f
}

type defaultMsgFloat64 struct {
}

func (d defaultMsgFloat64) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float64_notnil] = f
}

func (d defaultMsgFloat64) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float64_min] = f
}

func (d defaultMsgFloat64) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float64_max] = f
}

func (d defaultMsgFloat64) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float64_range] = f
}

func (d defaultMsgFloat64) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float64_gt] = f
}

func (d defaultMsgFloat64) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float64_lt] = f
}

func (d defaultMsgFloat64) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float64_within] = f
}

func (d defaultMsgFloat64) Options(f func(ctx *Context) string) {
	defaultMsgs[default_msg_float64_options] = f
}

type defaultMsgString struct {
}

func (d defaultMsgString) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_string_notnil] = f
}

func (d defaultMsgString) NotBlank(f func(ctx *Context) string) {
	defaultMsgs[default_msg_string_notblank] = f
}

func (d defaultMsgString) Length(f func(ctx *Context) string) {
	defaultMsgs[default_msg_string_length] = f
}

func (d defaultMsgString) Regex(f func(ctx *Context) string) {
	defaultMsgs[default_msg_string_regex] = f
}

func (d defaultMsgString) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_string_min] = f
}

func (d defaultMsgString) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_string_max] = f
}

func (d defaultMsgString) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_string_range] = f
}

func (d defaultMsgString) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_string_gt] = f
}

func (d defaultMsgString) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_string_lt] = f
}

func (d defaultMsgString) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_string_within] = f
}

func (d defaultMsgString) Options(f func(ctx *Context) string) {
	defaultMsgs[default_msg_string_options] = f
}

type defaultMsgSlices struct {
}

func (d defaultMsgSlices) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_slices_notnil] = f
}

func (d defaultMsgSlices) NotEmpty(f func(ctx *Context) string) {
	defaultMsgs[default_msg_slices_notempty] = f
}

func (d defaultMsgSlices) Length(f func(ctx *Context) string) {
	defaultMsgs[default_msg_slices_length] = f
}

func (d defaultMsgSlices) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_slices_min] = f
}

func (d defaultMsgSlices) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_slices_max] = f
}

func (d defaultMsgSlices) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_slices_range] = f
}

func (d defaultMsgSlices) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_slices_gt] = f
}

func (d defaultMsgSlices) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_slices_lt] = f
}

func (d defaultMsgSlices) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_slices_within] = f
}

type defaultMsgMap struct {
}

func (d defaultMsgMap) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_map_notnil] = f
}

func (d defaultMsgMap) NotEmpty(f func(ctx *Context) string) {
	defaultMsgs[default_msg_map_notempty] = f
}

func (d defaultMsgMap) Length(f func(ctx *Context) string) {
	defaultMsgs[default_msg_map_length] = f
}

func (d defaultMsgMap) Min(f func(ctx *Context) string) {
	defaultMsgs[default_msg_map_min] = f
}

func (d defaultMsgMap) Max(f func(ctx *Context) string) {
	defaultMsgs[default_msg_map_max] = f
}

func (d defaultMsgMap) Range(f func(ctx *Context) string) {
	defaultMsgs[default_msg_map_range] = f
}

func (d defaultMsgMap) Gt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_map_gt] = f
}

func (d defaultMsgMap) Lt(f func(ctx *Context) string) {
	defaultMsgs[default_msg_map_lt] = f
}

func (d defaultMsgMap) Within(f func(ctx *Context) string) {
	defaultMsgs[default_msg_map_within] = f
}

type defaultMsgStruct struct {
}

func (d defaultMsgStruct) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_struct_notnil] = f
}

type defaultMsgAny struct {
}

func (d defaultMsgAny) NotNil(f func(ctx *Context) string) {
	defaultMsgs[default_msg_any_notnil] = f
}

type defaultMsg struct {
}

func (d defaultMsg) Bool() defaultMsgBool {
	return defaultMsgBool{}
}

func (d defaultMsg) Byte() defaultMsgByte {
	return defaultMsgByte{}
}

func (d defaultMsg) Int() defaultMsgInt {
	return defaultMsgInt{}
}

func (d defaultMsg) Int8() defaultMsgInt8 {
	return defaultMsgInt8{}
}

func (d defaultMsg) Int16() defaultMsgInt16 {
	return defaultMsgInt16{}
}

func (d defaultMsg) Int32() defaultMsgInt32 {
	return defaultMsgInt32{}
}

func (d defaultMsg) Int64() defaultMsgInt64 {
	return defaultMsgInt64{}
}

func (d defaultMsg) Uint() defaultMsgUint {
	return defaultMsgUint{}
}

func (d defaultMsg) Uint8() defaultMsgUint8 {
	return defaultMsgUint8{}
}

func (d defaultMsg) Uint16() defaultMsgUint16 {
	return defaultMsgUint16{}
}

func (d defaultMsg) Uint32() defaultMsgUint32 {
	return defaultMsgUint32{}
}

func (d defaultMsg) Uint64() defaultMsgUint64 {
	return defaultMsgUint64{}
}

func (d defaultMsg) Float32() defaultMsgFloat32 {
	return defaultMsgFloat32{}
}

func (d defaultMsg) Float64() defaultMsgFloat64 {
	return defaultMsgFloat64{}
}

func (d defaultMsg) String() defaultMsgString {
	return defaultMsgString{}
}

func (d defaultMsg) Slices() defaultMsgSlices {
	return defaultMsgSlices{}
}

func (d defaultMsg) Map() defaultMsgMap {
	return defaultMsgMap{}
}

func (d defaultMsg) Struct() defaultMsgStruct {
	return defaultMsgStruct{}
}

func (d defaultMsg) Any() defaultMsgAny {
	return defaultMsgAny{}
}

func DefaultMsg() defaultMsg {
	return defaultMsg{}
}
