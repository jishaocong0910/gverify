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

package vfy_test

import (
	"testing"

	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

func TestCheckUint16_NotNil(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(1)), "").NotNil().Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint16(c, (*uint16)(nil), "").NotNil().Msg("test fail by nil")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)

		vfy.Uint16(c, (*uint16)(nil), "").NotNil().Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Uint16().NotNil(func(ctx *vfy.Context) string {
			return "uint16 NotNil default msg"
		})
		vfy.Uint16(c, (*uint16)(nil), "").NotNil().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint16 NotNil default msg", msg)
	}
}

func TestCheckUint16_Min(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(10)), "").Min(10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint16(c, (*uint16)(nil), "param").Min(10).Msg("%s must not be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be less than 10", msg)

		vfy.Uint16(c, (*uint16)(nil), "").Min(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(9)), "param").Min(10).Msg("%s must not be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Uint16().Min(func(ctx *vfy.Context) string {
			return "uint16 Min default msg"
		})
		vfy.Uint16(c, ptr(uint16(9)), "param").Min(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint16 Min default msg", msg)
	}
}

func TestCheckUint16_Max(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(10)), "").Max(10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint16(c, (*uint16)(nil), "param").Max(10).Msg("%s must not be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be greater than 10", msg)

		vfy.Uint16(c, (*uint16)(nil), "").Max(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(11)), "param").Max(10).Msg("%s must not be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Uint16().Max(func(ctx *vfy.Context) string {
			return "uint16 Max default msg"
		})
		vfy.Uint16(c, ptr(uint16(11)), "param").Max(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint16 Max default msg", msg)
	}
}

func TestCheckUint16_Range(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(10)), "").Range(5, 10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint16(c, (*uint16)(nil), "param").Range(5, 10).Msg("%s must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)

		vfy.Uint16(c, (*uint16)(nil), "").Range(5, 10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(4)), "param").Range(5, 10).Msg("%s must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(11)), "param").Range(5, 10).Msg("%s must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Uint16().Range(func(ctx *vfy.Context) string {
			return "uint16 Range default msg"
		})
		vfy.Uint16(c, ptr(uint16(11)), "param").Range(5, 10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint16 Range default msg", msg)
	}
}

func TestCheckUint16_Gt(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		ok, msg, _ := vfy.GetResult(c)
		vfy.Uint16(c, ptr(uint16(11)), "param").Gt(10).Msg("test success")
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint16(c, (*uint16)(nil), "param").Gt(10).Msg("%s must be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 10", msg)

		vfy.Uint16(c, (*uint16)(nil), "param").Gt(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(10)), "param").Gt(10).Msg("%s must be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Uint16().Gt(func(ctx *vfy.Context) string {
			return "uint16 Gt default msg"
		})
		vfy.Uint16(c, ptr(uint16(10)), "param").Gt(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint16 Gt default msg", msg)
	}
}

func TestCheckUint16_Lt(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(9)), "param").Lt(10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint16(c, (*uint16)(nil), "param").Lt(10).Msg("%s must be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be less than 10", msg)

		vfy.Uint16(c, (*uint16)(nil), "param").Lt(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(11)), "param").Lt(10).Msg("%s must be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Uint16().Lt(func(ctx *vfy.Context) string {
			return "uint16 Lt default msg"
		})
		vfy.Uint16(c, ptr(uint16(11)), "param").Lt(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint16 Lt default msg", msg)
	}
}

func TestCheckUint16_Within(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(9)), "param").Within(5, 10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint16(c, (*uint16)(nil), "param").Within(5, 10).Msg("%s must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)

		vfy.Uint16(c, (*uint16)(nil), "param").Within(5, 10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(5)), "param").Within(5, 10).Msg("%s must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(10)), "param").Within(5, 10).Msg("%s must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Uint16().Within(func(ctx *vfy.Context) string {
			return "uint16 Within default msg"
		})
		vfy.Uint16(c, ptr(uint16(10)), "param").Within(5, 10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint16 Within default msg", msg)
	}
}

func TestCheckUint16_Options(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(1)), "").Options([]uint16{1, 2, 3}).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint16(c, (*uint16)(nil), "param").Options([]uint16{1, 2, 3}).Msg("%s must be %s", c.FieldName(), c.Confines())
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 1, 2 or 3", msg)

		vfy.Uint16(c, (*uint16)(nil), "").Options([]uint16{1, 2, 3}).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 1, 2 or 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(4)), "param").Options([]uint16{1, 2, 3}).Msg("%s must be %s", c.FieldName(), c.Confines())
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 1, 2 or 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Uint16().Options(func(ctx *vfy.Context) string {
			return "uint16 Options default msg"
		})
		vfy.Uint16(c, ptr(uint16(4)), "param").Options([]uint16{1, 2, 3}).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint16 Options default msg", msg)
	}
}

func TestCheckUint16_Custom(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(1)), "").Custom(func(i uint16) bool {
			return true
		}).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint16(c, (*uint16)(nil), "").Custom(nil).Msg("test fail by nil")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)

		vfy.Uint16(c, (*uint16)(nil), "").Custom(nil).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint16(c, ptr(uint16(2)), "").Custom(func(i uint16) bool {
			r.Equal(uint16(2), i)
			return false
		}).Msg("test fail by custom")
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by custom", msg)
	}
}
