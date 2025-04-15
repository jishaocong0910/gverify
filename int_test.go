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

func TestCheckInt_NotNil(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(1), "").NotNil().Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Int(c, (*int)(nil), "").NotNil().Msg("test fail by nil")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)

		vfy.Int(c, (*int)(nil), "").NotNil().Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, (*int)(nil), "param").NotNil().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Int().NotNil(func(ctx *vfy.Context) string {
			return "int NotNil default setMsg"
		})
		vfy.Int(c, (*int)(nil), "param").NotNil().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("int NotNil default setMsg", msg)
	}
}

func TestCheckInt_Min(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(10), "").Min(10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Int(c, (*int)(nil), "param").Min(10).Msg("%s must not be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be less than 10", msg)

		vfy.Int(c, (*int)(nil), "").Min(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(9), "param").Min(10).Msg("%s must not be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(9), "param").Min(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Int().Min(func(ctx *vfy.Context) string {
			return "int Min default setMsg"
		})
		vfy.Int(c, ptr(9), "param").Min(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("int Min default setMsg", msg)
	}
}

func TestCheckInt_Max(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(10), "").Max(10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Int(c, (*int)(nil), "param").Max(10).Msg("%s must not be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be greater than 10", msg)

		vfy.Int(c, (*int)(nil), "").Max(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(11), "param").Max(10).Msg("%s must not be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(11), "param").Max(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Int().Max(func(ctx *vfy.Context) string {
			return "int Max default setMsg"
		})
		vfy.Int(c, ptr(11), "param").Max(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("int Max default setMsg", msg)
	}
}

func TestCheckInt_Range(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(10), "").Range(5, 10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Int(c, (*int)(nil), "param").Range(5, 10).Msg("%s must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)

		vfy.Int(c, (*int)(nil), "").Range(5, 10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(4), "param").Range(5, 10).Msg("%s must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(11), "param").Range(5, 10).Msg("%s must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(11), "param").Range(5, 10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 5 to 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Int().Range(func(ctx *vfy.Context) string {
			return "int Range default setMsg"
		})
		vfy.Int(c, ptr(11), "param").Range(5, 10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("int Range default setMsg", msg)
	}
}

func TestCheckInt_Gt(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		ok, msg, _ := vfy.GetResult(c)
		vfy.Int(c, ptr(11), "param").Gt(10).Msg("test success")
		r.True(ok)
		r.Equal("", msg)

		vfy.Int(c, (*int)(nil), "param").Gt(10).Msg("%s must be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 10", msg)

		vfy.Int(c, (*int)(nil), "param").Gt(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(10), "param").Gt(10).Msg("%s must be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(10), "param").Gt(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Int().Gt(func(ctx *vfy.Context) string {
			return "int Gt default setMsg"
		})
		vfy.Int(c, ptr(10), "param").Gt(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("int Gt default setMsg", msg)
	}
}

func TestCheckInt_Lt(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(9), "param").Lt(10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Int(c, (*int)(nil), "param").Lt(10).Msg("%s must be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be less than 10", msg)

		vfy.Int(c, (*int)(nil), "param").Lt(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(11), "param").Lt(10).Msg("%s must be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(11), "param").Lt(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Int().Lt(func(ctx *vfy.Context) string {
			return "int Lt default setMsg"
		})
		vfy.Int(c, ptr(11), "param").Lt(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("int Lt default setMsg", msg)
	}
}

func TestCheckInt_Within(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(9), "param").Within(5, 10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Int(c, (*int)(nil), "param").Within(5, 10).Msg("%s must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)

		vfy.Int(c, (*int)(nil), "param").Within(5, 10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(5), "param").Within(5, 10).Msg("%s must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(10), "param").Within(5, 10).Msg("%s must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(10), "param").Within(5, 10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 5 and less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Int().Within(func(ctx *vfy.Context) string {
			return "int Within default setMsg"
		})
		vfy.Int(c, ptr(10), "param").Within(5, 10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("int Within default setMsg", msg)
	}
}

func TestCheckInt_Options(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(1), "").Options([]int{1, 2, 3}).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Int(c, (*int)(nil), "param").Options([]int{1, 2, 3}).Msg("%s must be %s", c.FieldName(), c.Confines())
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 1, 2 or 3", msg)

		vfy.Int(c, (*int)(nil), "").Options([]int{1, 2, 3}).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 1, 2 or 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(4), "param").Options([]int{1, 2, 3}).Msg("%s must be %s", c.FieldName(), c.Confines())
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 1, 2 or 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(4), "param").Options([]int{1, 2, 3}).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 1, 2 or 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().Int().Options(func(ctx *vfy.Context) string {
			return "int Options default setMsg"
		})
		vfy.Int(c, ptr(4), "param").Options([]int{1, 2, 3}).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("int Options default setMsg", msg)
	}
}

func TestCheckInt_Custom(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(1), "").Custom(func(i int) bool {
			return true
		}).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Int(c, (*int)(nil), "").Custom(nil).Msg("test fail by nil")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)

		vfy.Int(c, (*int)(nil), "").Custom(nil).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Int(c, ptr(2), "").Custom(func(i int) bool {
			r.Equal(2, i)
			return false
		}).Msg("test fail by custom")
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by custom", msg)
	}
}
