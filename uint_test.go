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

func TestCheckUint_NotNil(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(1)), "").NotNil().Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint(c, (*uint)(nil), "").NotNil().Msg("test fail by nil")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)

		vfy.Uint(c, (*uint)(nil), "").NotNil().Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, (*uint)(nil), "param").NotNil().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Uint().NotNil(func(ctx *vfy.Context) string {
			return "uint NotNil default setMsg"
		})
		vfy.Uint(c, (*uint)(nil), "param").NotNil().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint NotNil default setMsg", msg)
	}
}

func TestCheckUint_Min(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(10)), "").Min(10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint(c, (*uint)(nil), "param").Min(10).Msg("%s must not be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be less than 10", msg)

		vfy.Uint(c, (*uint)(nil), "").Min(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(9)), "param").Min(10).Msg("%s must not be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(9)), "param").Min(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Uint().Min(func(ctx *vfy.Context) string {
			return "uint Min default setMsg"
		})
		vfy.Uint(c, ptr(uint(9)), "param").Min(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint Min default setMsg", msg)
	}
}

func TestCheckUint_Max(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(10)), "").Max(10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint(c, (*uint)(nil), "param").Max(10).Msg("%s must not be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be greater than 10", msg)

		vfy.Uint(c, (*uint)(nil), "").Max(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(11)), "param").Max(10).Msg("%s must not be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(11)), "param").Max(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Uint().Max(func(ctx *vfy.Context) string {
			return "uint Max default setMsg"
		})
		vfy.Uint(c, ptr(uint(11)), "param").Max(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint Max default setMsg", msg)
	}
}

func TestCheckUint_Range(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(10)), "").Range(5, 10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint(c, (*uint)(nil), "param").Range(5, 10).Msg("%s must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)

		vfy.Uint(c, (*uint)(nil), "").Range(5, 10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(4)), "param").Range(5, 10).Msg("%s must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(11)), "param").Range(5, 10).Msg("%s must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(11)), "param").Range(5, 10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 5 to 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Uint().Range(func(ctx *vfy.Context) string {
			return "uint Range default setMsg"
		})
		vfy.Uint(c, ptr(uint(11)), "param").Range(5, 10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint Range default setMsg", msg)
	}
}

func TestCheckUint_Gt(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		ok, msg, _ := vfy.GetResult(c)
		vfy.Uint(c, ptr(uint(11)), "param").Gt(10).Msg("test success")
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint(c, (*uint)(nil), "param").Gt(10).Msg("%s must be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 10", msg)

		vfy.Uint(c, (*uint)(nil), "param").Gt(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(10)), "param").Gt(10).Msg("%s must be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(10)), "param").Gt(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Uint().Gt(func(ctx *vfy.Context) string {
			return "uint Gt default setMsg"
		})
		vfy.Uint(c, ptr(uint(10)), "param").Gt(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint Gt default setMsg", msg)
	}
}

func TestCheckUint_Lt(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(9)), "param").Lt(10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint(c, (*uint)(nil), "param").Lt(10).Msg("%s must be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be less than 10", msg)

		vfy.Uint(c, (*uint)(nil), "param").Lt(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(11)), "param").Lt(10).Msg("%s must be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(11)), "param").Lt(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Uint().Lt(func(ctx *vfy.Context) string {
			return "uint Lt default setMsg"
		})
		vfy.Uint(c, ptr(uint(11)), "param").Lt(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint Lt default setMsg", msg)
	}
}

func TestCheckUint_Within(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(9)), "param").Within(5, 10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint(c, (*uint)(nil), "param").Within(5, 10).Msg("%s must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)

		vfy.Uint(c, (*uint)(nil), "param").Within(5, 10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(5)), "param").Within(5, 10).Msg("%s must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(10)), "param").Within(5, 10).Msg("%s must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(10)), "param").Within(5, 10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 5 and less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Uint().Within(func(ctx *vfy.Context) string {
			return "uint Within default setMsg"
		})
		vfy.Uint(c, ptr(uint(10)), "param").Within(5, 10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint Within default setMsg", msg)
	}
}

func TestCheckUint_Options(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(1)), "").Options([]uint{1, 2, 3}).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint(c, (*uint)(nil), "param").Options([]uint{1, 2, 3}).Msg("%s must be %s", c.FieldName(), c.Confines())
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 1, 2 or 3", msg)

		vfy.Uint(c, (*uint)(nil), "").Options([]uint{1, 2, 3}).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 1, 2 or 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(4)), "param").Options([]uint{1, 2, 3}).Msg("%s must be %s", c.FieldName(), c.Confines())
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 1, 2 or 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(4)), "param").Options([]uint{1, 2, 3}).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 1, 2 or 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Uint().Options(func(ctx *vfy.Context) string {
			return "uint Options default setMsg"
		})
		vfy.Uint(c, ptr(uint(4)), "param").Options([]uint{1, 2, 3}).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("uint Options default setMsg", msg)
	}
}

func TestCheckUint_Custom(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(1)), "").Custom(func(i uint) bool {
			return true
		}).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Uint(c, (*uint)(nil), "").Custom(nil).Msg("test fail by nil")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)

		vfy.Uint(c, (*uint)(nil), "").Custom(nil).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Uint(c, ptr(uint(2)), "").Custom(func(i uint) bool {
			r.Equal(uint(2), i)
			return false
		}).Msg("test fail by custom")
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by custom", msg)
	}
}
