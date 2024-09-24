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

type setMsg[T any] struct {
	ctx *Context
	t   T
}

func (s setMsg[T]) Msg(msg string, args ...any) T {
	return s.Msg_("", msg, args...)
}

func (s setMsg[T]) Msg_(code, msg string, args ...any) T {
	if s.ctx != nil {
		if !s.ctx.all {
			s.ctx.code = code
		}
		s.ctx.addMsg(msg, args...)
	}
	return s.t
}

type setMsgOrDefault[T any] struct {
	setMsg[T]
	k defaultMsgKey
}

func (s setMsgOrDefault[T]) DefaultMsg() T {
	return s.DefaultMsg_("")
}

func (s setMsgOrDefault[T]) DefaultMsg_(code string) T {
	if s.ctx != nil {
		var m string
		if f, ok := defaultMsgs[s.k]; ok {
			m = f(s.ctx)
		}
		s.Msg_(code, m)
	}
	return s.t
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

var defaultMsgs = map[defaultMsgKey]func(*Context) string{}

type setDefaultMsg struct {
}

func (s setDefaultMsg) Bool() setDefaultMsgBool {
	return setDefaultMsgBool{}
}

func (s setDefaultMsg) Byte() setDefaultMsgByte {
	return setDefaultMsgByte{}
}

func (s setDefaultMsg) Int() setDefaultMsgInt {
	return setDefaultMsgInt{}
}

func (s setDefaultMsg) Int8() setDefaultMsgInt8 {
	return setDefaultMsgInt8{}
}

func (s setDefaultMsg) Int16() setDefaultMsgInt16 {
	return setDefaultMsgInt16{}
}

func (s setDefaultMsg) Int32() setDefaultMsgInt32 {
	return setDefaultMsgInt32{}
}

func (s setDefaultMsg) Int64() setDefaultMsgInt64 {
	return setDefaultMsgInt64{}
}

func (s setDefaultMsg) Uint() setDefaultMsgUint {
	return setDefaultMsgUint{}
}

func (s setDefaultMsg) Uint8() setDefaultMsgUint8 {
	return setDefaultMsgUint8{}
}

func (s setDefaultMsg) Uint16() setDefaultMsgUint16 {
	return setDefaultMsgUint16{}
}

func (s setDefaultMsg) Uint32() setDefaultMsgUint32 {
	return setDefaultMsgUint32{}
}

func (s setDefaultMsg) Uint64() setDefaultMsgUint64 {
	return setDefaultMsgUint64{}
}

func (s setDefaultMsg) Float32() setDefaultMsgFloat32 {
	return setDefaultMsgFloat32{}
}

func (s setDefaultMsg) Float64() setDefaultMsgFloat64 {
	return setDefaultMsgFloat64{}
}

func (s setDefaultMsg) String() setDefaultMsgString {
	return setDefaultMsgString{}
}

func (s setDefaultMsg) Slices() setDefaultMsgSlices {
	return setDefaultMsgSlices{}
}

func (s setDefaultMsg) Map() setDefaultMsgMap {
	return setDefaultMsgMap{}
}

func (s setDefaultMsg) Struct() setDefaultMsgStruct {
	return setDefaultMsgStruct{}
}

func (s setDefaultMsg) Any() setDefaultMsgAny {
	return setDefaultMsgAny{}
}

type setDefaultMsgBool struct {
}

func (s setDefaultMsgBool) NotNil(f func(ctx *Context) string) setDefaultMsgBool {
	defaultMsgs[default_msg_bool_notnil] = f
	return s
}

type setDefaultMsgByte struct {
}

func (s setDefaultMsgByte) NotNil(f func(ctx *Context) string) setDefaultMsgByte {
	defaultMsgs[default_msg_byte_notnil] = f
	return s
}

func (s setDefaultMsgByte) Min(f func(ctx *Context) string) setDefaultMsgByte {
	defaultMsgs[default_msg_byte_min] = f
	return s
}

func (s setDefaultMsgByte) Max(f func(ctx *Context) string) setDefaultMsgByte {
	defaultMsgs[default_msg_byte_max] = f
	return s
}

func (s setDefaultMsgByte) Range(f func(ctx *Context) string) setDefaultMsgByte {
	defaultMsgs[default_msg_byte_range] = f
	return s
}

func (s setDefaultMsgByte) Gt(f func(ctx *Context) string) setDefaultMsgByte {
	defaultMsgs[default_msg_byte_gt] = f
	return s
}

func (s setDefaultMsgByte) Lt(f func(ctx *Context) string) setDefaultMsgByte {
	defaultMsgs[default_msg_byte_lt] = f
	return s
}

func (s setDefaultMsgByte) Within(f func(ctx *Context) string) setDefaultMsgByte {
	defaultMsgs[default_msg_byte_within] = f
	return s
}

func (s setDefaultMsgByte) Options(f func(ctx *Context) string) setDefaultMsgByte {
	defaultMsgs[default_msg_byte_options] = f
	return s
}

type setDefaultMsgInt struct {
}

func (s setDefaultMsgInt) NotNil(f func(ctx *Context) string) setDefaultMsgInt {
	defaultMsgs[default_msg_int_notnil] = f
	return s
}

func (s setDefaultMsgInt) Min(f func(ctx *Context) string) setDefaultMsgInt {
	defaultMsgs[default_msg_int_min] = f
	return s
}

func (s setDefaultMsgInt) Max(f func(ctx *Context) string) setDefaultMsgInt {
	defaultMsgs[default_msg_int_max] = f
	return s
}

func (s setDefaultMsgInt) Range(f func(ctx *Context) string) setDefaultMsgInt {
	defaultMsgs[default_msg_int_range] = f
	return s
}

func (s setDefaultMsgInt) Gt(f func(ctx *Context) string) setDefaultMsgInt {
	defaultMsgs[default_msg_int_gt] = f
	return s
}

func (s setDefaultMsgInt) Lt(f func(ctx *Context) string) setDefaultMsgInt {
	defaultMsgs[default_msg_int_lt] = f
	return s
}

func (s setDefaultMsgInt) Within(f func(ctx *Context) string) setDefaultMsgInt {
	defaultMsgs[default_msg_int_within] = f
	return s
}

func (s setDefaultMsgInt) Options(f func(ctx *Context) string) setDefaultMsgInt {
	defaultMsgs[default_msg_int_options] = f
	return s
}

type setDefaultMsgInt8 struct {
}

func (s setDefaultMsgInt8) NotNil(f func(ctx *Context) string) setDefaultMsgInt8 {
	defaultMsgs[default_msg_int8_notnil] = f
	return s
}

func (s setDefaultMsgInt8) Min(f func(ctx *Context) string) setDefaultMsgInt8 {
	defaultMsgs[default_msg_int8_min] = f
	return s
}

func (s setDefaultMsgInt8) Max(f func(ctx *Context) string) setDefaultMsgInt8 {
	defaultMsgs[default_msg_int8_max] = f
	return s
}

func (s setDefaultMsgInt8) Range(f func(ctx *Context) string) setDefaultMsgInt8 {
	defaultMsgs[default_msg_int8_range] = f
	return s
}

func (s setDefaultMsgInt8) Gt(f func(ctx *Context) string) setDefaultMsgInt8 {
	defaultMsgs[default_msg_int8_gt] = f
	return s
}

func (s setDefaultMsgInt8) Lt(f func(ctx *Context) string) setDefaultMsgInt8 {
	defaultMsgs[default_msg_int8_lt] = f
	return s
}

func (s setDefaultMsgInt8) Within(f func(ctx *Context) string) setDefaultMsgInt8 {
	defaultMsgs[default_msg_int8_within] = f
	return s
}

func (s setDefaultMsgInt8) Options(f func(ctx *Context) string) setDefaultMsgInt8 {
	defaultMsgs[default_msg_int8_options] = f
	return s
}

type setDefaultMsgInt16 struct {
}

func (s setDefaultMsgInt16) NotNil(f func(ctx *Context) string) setDefaultMsgInt16 {
	defaultMsgs[default_msg_int16_notnil] = f
	return s
}

func (s setDefaultMsgInt16) Min(f func(ctx *Context) string) setDefaultMsgInt16 {
	defaultMsgs[default_msg_int16_min] = f
	return s
}

func (s setDefaultMsgInt16) Max(f func(ctx *Context) string) setDefaultMsgInt16 {
	defaultMsgs[default_msg_int16_max] = f
	return s
}

func (s setDefaultMsgInt16) Range(f func(ctx *Context) string) setDefaultMsgInt16 {
	defaultMsgs[default_msg_int16_range] = f
	return s
}

func (s setDefaultMsgInt16) Gt(f func(ctx *Context) string) setDefaultMsgInt16 {
	defaultMsgs[default_msg_int16_gt] = f
	return s
}

func (s setDefaultMsgInt16) Lt(f func(ctx *Context) string) setDefaultMsgInt16 {
	defaultMsgs[default_msg_int16_lt] = f
	return s
}

func (s setDefaultMsgInt16) Within(f func(ctx *Context) string) setDefaultMsgInt16 {
	defaultMsgs[default_msg_int16_within] = f
	return s
}

func (s setDefaultMsgInt16) Options(f func(ctx *Context) string) setDefaultMsgInt16 {
	defaultMsgs[default_msg_int16_options] = f
	return s
}

type setDefaultMsgInt32 struct {
}

func (s setDefaultMsgInt32) NotNil(f func(ctx *Context) string) setDefaultMsgInt32 {
	defaultMsgs[default_msg_int32_notnil] = f
	return s
}

func (s setDefaultMsgInt32) Min(f func(ctx *Context) string) setDefaultMsgInt32 {
	defaultMsgs[default_msg_int32_min] = f
	return s
}

func (s setDefaultMsgInt32) Max(f func(ctx *Context) string) setDefaultMsgInt32 {
	defaultMsgs[default_msg_int32_max] = f
	return s
}

func (s setDefaultMsgInt32) Range(f func(ctx *Context) string) setDefaultMsgInt32 {
	defaultMsgs[default_msg_int32_range] = f
	return s
}

func (s setDefaultMsgInt32) Gt(f func(ctx *Context) string) setDefaultMsgInt32 {
	defaultMsgs[default_msg_int32_gt] = f
	return s
}

func (s setDefaultMsgInt32) Lt(f func(ctx *Context) string) setDefaultMsgInt32 {
	defaultMsgs[default_msg_int32_lt] = f
	return s
}

func (s setDefaultMsgInt32) Within(f func(ctx *Context) string) setDefaultMsgInt32 {
	defaultMsgs[default_msg_int32_within] = f
	return s
}

func (s setDefaultMsgInt32) Options(f func(ctx *Context) string) setDefaultMsgInt32 {
	defaultMsgs[default_msg_int32_options] = f
	return s
}

type setDefaultMsgInt64 struct {
}

func (s setDefaultMsgInt64) NotNil(f func(ctx *Context) string) setDefaultMsgInt64 {
	defaultMsgs[default_msg_int64_notnil] = f
	return s
}

func (s setDefaultMsgInt64) Min(f func(ctx *Context) string) setDefaultMsgInt64 {
	defaultMsgs[default_msg_int64_min] = f
	return s
}

func (s setDefaultMsgInt64) Max(f func(ctx *Context) string) setDefaultMsgInt64 {
	defaultMsgs[default_msg_int64_max] = f
	return s
}

func (s setDefaultMsgInt64) Range(f func(ctx *Context) string) setDefaultMsgInt64 {
	defaultMsgs[default_msg_int64_range] = f
	return s
}

func (s setDefaultMsgInt64) Gt(f func(ctx *Context) string) setDefaultMsgInt64 {
	defaultMsgs[default_msg_int64_gt] = f
	return s
}

func (s setDefaultMsgInt64) Lt(f func(ctx *Context) string) setDefaultMsgInt64 {
	defaultMsgs[default_msg_int64_lt] = f
	return s
}

func (s setDefaultMsgInt64) Within(f func(ctx *Context) string) setDefaultMsgInt64 {
	defaultMsgs[default_msg_int64_within] = f
	return s
}

func (s setDefaultMsgInt64) Options(f func(ctx *Context) string) setDefaultMsgInt64 {
	defaultMsgs[default_msg_int64_options] = f
	return s
}

type setDefaultMsgUint struct {
}

func (s setDefaultMsgUint) NotNil(f func(ctx *Context) string) setDefaultMsgUint {
	defaultMsgs[default_msg_uint_notnil] = f
	return s
}

func (s setDefaultMsgUint) Min(f func(ctx *Context) string) setDefaultMsgUint {
	defaultMsgs[default_msg_uint_min] = f
	return s
}

func (s setDefaultMsgUint) Max(f func(ctx *Context) string) setDefaultMsgUint {
	defaultMsgs[default_msg_uint_max] = f
	return s
}

func (s setDefaultMsgUint) Range(f func(ctx *Context) string) setDefaultMsgUint {
	defaultMsgs[default_msg_uint_range] = f
	return s
}

func (s setDefaultMsgUint) Gt(f func(ctx *Context) string) setDefaultMsgUint {
	defaultMsgs[default_msg_uint_gt] = f
	return s
}

func (s setDefaultMsgUint) Lt(f func(ctx *Context) string) setDefaultMsgUint {
	defaultMsgs[default_msg_uint_lt] = f
	return s
}

func (s setDefaultMsgUint) Within(f func(ctx *Context) string) setDefaultMsgUint {
	defaultMsgs[default_msg_uint_within] = f
	return s
}

func (s setDefaultMsgUint) Options(f func(ctx *Context) string) setDefaultMsgUint {
	defaultMsgs[default_msg_uint_options] = f
	return s
}

type setDefaultMsgUint8 struct {
}

func (s setDefaultMsgUint8) NotNil(f func(ctx *Context) string) setDefaultMsgUint8 {
	defaultMsgs[default_msg_uint8_notnil] = f
	return s
}

func (s setDefaultMsgUint8) Min(f func(ctx *Context) string) setDefaultMsgUint8 {
	defaultMsgs[default_msg_uint8_min] = f
	return s
}

func (s setDefaultMsgUint8) Max(f func(ctx *Context) string) setDefaultMsgUint8 {
	defaultMsgs[default_msg_uint8_max] = f
	return s
}

func (s setDefaultMsgUint8) Range(f func(ctx *Context) string) setDefaultMsgUint8 {
	defaultMsgs[default_msg_uint8_range] = f
	return s
}

func (s setDefaultMsgUint8) Gt(f func(ctx *Context) string) setDefaultMsgUint8 {
	defaultMsgs[default_msg_uint8_gt] = f
	return s
}

func (s setDefaultMsgUint8) Lt(f func(ctx *Context) string) setDefaultMsgUint8 {
	defaultMsgs[default_msg_uint8_lt] = f
	return s
}

func (s setDefaultMsgUint8) Within(f func(ctx *Context) string) setDefaultMsgUint8 {
	defaultMsgs[default_msg_uint8_within] = f
	return s
}

func (s setDefaultMsgUint8) Options(f func(ctx *Context) string) setDefaultMsgUint8 {
	defaultMsgs[default_msg_uint8_options] = f
	return s
}

type setDefaultMsgUint16 struct {
}

func (s setDefaultMsgUint16) NotNil(f func(ctx *Context) string) setDefaultMsgUint16 {
	defaultMsgs[default_msg_uint16_notnil] = f
	return s
}

func (s setDefaultMsgUint16) Min(f func(ctx *Context) string) setDefaultMsgUint16 {
	defaultMsgs[default_msg_uint16_min] = f
	return s
}

func (s setDefaultMsgUint16) Max(f func(ctx *Context) string) setDefaultMsgUint16 {
	defaultMsgs[default_msg_uint16_max] = f
	return s
}

func (s setDefaultMsgUint16) Range(f func(ctx *Context) string) setDefaultMsgUint16 {
	defaultMsgs[default_msg_uint16_range] = f
	return s
}

func (s setDefaultMsgUint16) Gt(f func(ctx *Context) string) setDefaultMsgUint16 {
	defaultMsgs[default_msg_uint16_gt] = f
	return s
}

func (s setDefaultMsgUint16) Lt(f func(ctx *Context) string) setDefaultMsgUint16 {
	defaultMsgs[default_msg_uint16_lt] = f
	return s
}

func (s setDefaultMsgUint16) Within(f func(ctx *Context) string) setDefaultMsgUint16 {
	defaultMsgs[default_msg_uint16_within] = f
	return s
}

func (s setDefaultMsgUint16) Options(f func(ctx *Context) string) setDefaultMsgUint16 {
	defaultMsgs[default_msg_uint16_options] = f
	return s
}

type setDefaultMsgUint32 struct {
}

func (s setDefaultMsgUint32) NotNil(f func(ctx *Context) string) setDefaultMsgUint32 {
	defaultMsgs[default_msg_uint32_notnil] = f
	return s
}

func (s setDefaultMsgUint32) Min(f func(ctx *Context) string) setDefaultMsgUint32 {
	defaultMsgs[default_msg_uint32_min] = f
	return s
}

func (s setDefaultMsgUint32) Max(f func(ctx *Context) string) setDefaultMsgUint32 {
	defaultMsgs[default_msg_uint32_max] = f
	return s
}

func (s setDefaultMsgUint32) Range(f func(ctx *Context) string) setDefaultMsgUint32 {
	defaultMsgs[default_msg_uint32_range] = f
	return s
}

func (s setDefaultMsgUint32) Gt(f func(ctx *Context) string) setDefaultMsgUint32 {
	defaultMsgs[default_msg_uint32_gt] = f
	return s
}

func (s setDefaultMsgUint32) Lt(f func(ctx *Context) string) setDefaultMsgUint32 {
	defaultMsgs[default_msg_uint32_lt] = f
	return s
}

func (s setDefaultMsgUint32) Within(f func(ctx *Context) string) setDefaultMsgUint32 {
	defaultMsgs[default_msg_uint32_within] = f
	return s
}

func (s setDefaultMsgUint32) Options(f func(ctx *Context) string) setDefaultMsgUint32 {
	defaultMsgs[default_msg_uint32_options] = f
	return s
}

type setDefaultMsgUint64 struct {
}

func (s setDefaultMsgUint64) NotNil(f func(ctx *Context) string) setDefaultMsgUint64 {
	defaultMsgs[default_msg_uint64_notnil] = f
	return s
}

func (s setDefaultMsgUint64) Min(f func(ctx *Context) string) setDefaultMsgUint64 {
	defaultMsgs[default_msg_uint64_min] = f
	return s
}

func (s setDefaultMsgUint64) Max(f func(ctx *Context) string) setDefaultMsgUint64 {
	defaultMsgs[default_msg_uint64_max] = f
	return s
}

func (s setDefaultMsgUint64) Range(f func(ctx *Context) string) setDefaultMsgUint64 {
	defaultMsgs[default_msg_uint64_range] = f
	return s
}

func (s setDefaultMsgUint64) Gt(f func(ctx *Context) string) setDefaultMsgUint64 {
	defaultMsgs[default_msg_uint64_gt] = f
	return s
}

func (s setDefaultMsgUint64) Lt(f func(ctx *Context) string) setDefaultMsgUint64 {
	defaultMsgs[default_msg_uint64_lt] = f
	return s
}

func (s setDefaultMsgUint64) Within(f func(ctx *Context) string) setDefaultMsgUint64 {
	defaultMsgs[default_msg_uint64_within] = f
	return s
}

func (s setDefaultMsgUint64) Options(f func(ctx *Context) string) setDefaultMsgUint64 {
	defaultMsgs[default_msg_uint64_options] = f
	return s
}

type setDefaultMsgFloat32 struct {
}

func (s setDefaultMsgFloat32) NotNil(f func(ctx *Context) string) setDefaultMsgFloat32 {
	defaultMsgs[default_msg_float32_notnil] = f
	return s
}

func (s setDefaultMsgFloat32) Min(f func(ctx *Context) string) setDefaultMsgFloat32 {
	defaultMsgs[default_msg_float32_min] = f
	return s
}

func (s setDefaultMsgFloat32) Max(f func(ctx *Context) string) setDefaultMsgFloat32 {
	defaultMsgs[default_msg_float32_max] = f
	return s
}

func (s setDefaultMsgFloat32) Range(f func(ctx *Context) string) setDefaultMsgFloat32 {
	defaultMsgs[default_msg_float32_range] = f
	return s
}

func (s setDefaultMsgFloat32) Gt(f func(ctx *Context) string) setDefaultMsgFloat32 {
	defaultMsgs[default_msg_float32_gt] = f
	return s
}

func (s setDefaultMsgFloat32) Lt(f func(ctx *Context) string) setDefaultMsgFloat32 {
	defaultMsgs[default_msg_float32_lt] = f
	return s
}

func (s setDefaultMsgFloat32) Within(f func(ctx *Context) string) setDefaultMsgFloat32 {
	defaultMsgs[default_msg_float32_within] = f
	return s
}

func (s setDefaultMsgFloat32) Options(f func(ctx *Context) string) setDefaultMsgFloat32 {
	defaultMsgs[default_msg_float32_options] = f
	return s
}

type setDefaultMsgFloat64 struct {
}

func (s setDefaultMsgFloat64) NotNil(f func(ctx *Context) string) setDefaultMsgFloat64 {
	defaultMsgs[default_msg_float64_notnil] = f
	return s
}

func (s setDefaultMsgFloat64) Min(f func(ctx *Context) string) setDefaultMsgFloat64 {
	defaultMsgs[default_msg_float64_min] = f
	return s
}

func (s setDefaultMsgFloat64) Max(f func(ctx *Context) string) setDefaultMsgFloat64 {
	defaultMsgs[default_msg_float64_max] = f
	return s
}

func (s setDefaultMsgFloat64) Range(f func(ctx *Context) string) setDefaultMsgFloat64 {
	defaultMsgs[default_msg_float64_range] = f
	return s
}

func (s setDefaultMsgFloat64) Gt(f func(ctx *Context) string) setDefaultMsgFloat64 {
	defaultMsgs[default_msg_float64_gt] = f
	return s
}

func (s setDefaultMsgFloat64) Lt(f func(ctx *Context) string) setDefaultMsgFloat64 {
	defaultMsgs[default_msg_float64_lt] = f
	return s
}

func (s setDefaultMsgFloat64) Within(f func(ctx *Context) string) setDefaultMsgFloat64 {
	defaultMsgs[default_msg_float64_within] = f
	return s
}

func (s setDefaultMsgFloat64) Options(f func(ctx *Context) string) setDefaultMsgFloat64 {
	defaultMsgs[default_msg_float64_options] = f
	return s
}

type setDefaultMsgString struct {
}

func (s setDefaultMsgString) NotNil(f func(ctx *Context) string) setDefaultMsgString {
	defaultMsgs[default_msg_string_notnil] = f
	return s
}

func (s setDefaultMsgString) NotBlank(f func(ctx *Context) string) setDefaultMsgString {
	defaultMsgs[default_msg_string_notblank] = f
	return s
}

func (s setDefaultMsgString) Length(f func(ctx *Context) string) setDefaultMsgString {
	defaultMsgs[default_msg_string_length] = f
	return s
}

func (s setDefaultMsgString) Regex(f func(ctx *Context) string) setDefaultMsgString {
	defaultMsgs[default_msg_string_regex] = f
	return s
}

func (s setDefaultMsgString) Min(f func(ctx *Context) string) setDefaultMsgString {
	defaultMsgs[default_msg_string_min] = f
	return s
}

func (s setDefaultMsgString) Max(f func(ctx *Context) string) setDefaultMsgString {
	defaultMsgs[default_msg_string_max] = f
	return s
}

func (s setDefaultMsgString) Range(f func(ctx *Context) string) setDefaultMsgString {
	defaultMsgs[default_msg_string_range] = f
	return s
}

func (s setDefaultMsgString) Gt(f func(ctx *Context) string) setDefaultMsgString {
	defaultMsgs[default_msg_string_gt] = f
	return s
}

func (s setDefaultMsgString) Lt(f func(ctx *Context) string) setDefaultMsgString {
	defaultMsgs[default_msg_string_lt] = f
	return s
}

func (s setDefaultMsgString) Within(f func(ctx *Context) string) setDefaultMsgString {
	defaultMsgs[default_msg_string_within] = f
	return s
}

func (s setDefaultMsgString) Options(f func(ctx *Context) string) setDefaultMsgString {
	defaultMsgs[default_msg_string_options] = f
	return s
}

type setDefaultMsgSlices struct {
}

func (s setDefaultMsgSlices) NotNil(f func(ctx *Context) string) setDefaultMsgSlices {
	defaultMsgs[default_msg_slices_notnil] = f
	return s
}

func (s setDefaultMsgSlices) NotEmpty(f func(ctx *Context) string) setDefaultMsgSlices {
	defaultMsgs[default_msg_slices_notempty] = f
	return s
}

func (s setDefaultMsgSlices) Length(f func(ctx *Context) string) setDefaultMsgSlices {
	defaultMsgs[default_msg_slices_length] = f
	return s
}

func (s setDefaultMsgSlices) Min(f func(ctx *Context) string) setDefaultMsgSlices {
	defaultMsgs[default_msg_slices_min] = f
	return s
}

func (s setDefaultMsgSlices) Max(f func(ctx *Context) string) setDefaultMsgSlices {
	defaultMsgs[default_msg_slices_max] = f
	return s
}

func (s setDefaultMsgSlices) Range(f func(ctx *Context) string) setDefaultMsgSlices {
	defaultMsgs[default_msg_slices_range] = f
	return s
}

func (s setDefaultMsgSlices) Gt(f func(ctx *Context) string) setDefaultMsgSlices {
	defaultMsgs[default_msg_slices_gt] = f
	return s
}

func (s setDefaultMsgSlices) Lt(f func(ctx *Context) string) setDefaultMsgSlices {
	defaultMsgs[default_msg_slices_lt] = f
	return s
}

func (s setDefaultMsgSlices) Within(f func(ctx *Context) string) setDefaultMsgSlices {
	defaultMsgs[default_msg_slices_within] = f
	return s
}

type setDefaultMsgMap struct {
}

func (s setDefaultMsgMap) NotNil(f func(ctx *Context) string) setDefaultMsgMap {
	defaultMsgs[default_msg_map_notnil] = f
	return s
}

func (s setDefaultMsgMap) NotEmpty(f func(ctx *Context) string) setDefaultMsgMap {
	defaultMsgs[default_msg_map_notempty] = f
	return s
}

func (s setDefaultMsgMap) Length(f func(ctx *Context) string) setDefaultMsgMap {
	defaultMsgs[default_msg_map_length] = f
	return s
}

func (s setDefaultMsgMap) Min(f func(ctx *Context) string) setDefaultMsgMap {
	defaultMsgs[default_msg_map_min] = f
	return s
}

func (s setDefaultMsgMap) Max(f func(ctx *Context) string) setDefaultMsgMap {
	defaultMsgs[default_msg_map_max] = f
	return s
}

func (s setDefaultMsgMap) Range(f func(ctx *Context) string) setDefaultMsgMap {
	defaultMsgs[default_msg_map_range] = f
	return s
}

func (s setDefaultMsgMap) Gt(f func(ctx *Context) string) setDefaultMsgMap {
	defaultMsgs[default_msg_map_gt] = f
	return s
}

func (s setDefaultMsgMap) Lt(f func(ctx *Context) string) setDefaultMsgMap {
	defaultMsgs[default_msg_map_lt] = f
	return s
}

func (s setDefaultMsgMap) Within(f func(ctx *Context) string) setDefaultMsgMap {
	defaultMsgs[default_msg_map_within] = f
	return s
}

type setDefaultMsgStruct struct {
}

func (s setDefaultMsgStruct) NotNil(f func(ctx *Context) string) setDefaultMsgStruct {
	defaultMsgs[default_msg_struct_notnil] = f
	return s
}

type setDefaultMsgAny struct {
}

func (s setDefaultMsgAny) NotNil(f func(ctx *Context) string) setDefaultMsgAny {
	defaultMsgs[default_msg_any_notnil] = f
	return s
}

func SetDefaultMsg() setDefaultMsg {
	return setDefaultMsg{}
}
