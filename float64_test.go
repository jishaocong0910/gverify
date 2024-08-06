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

func TestCheckFloat64_NotNil(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(1)), "").NotNil().Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Float64(c, (*float64)(nil), "").NotNil().Msg("test fail by nil")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)

		vfy.Float64(c, (*float64)(nil), "").NotNil().Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Float64().NotNil(func(ctx *vfy.Context) string {
			return "float64 NotNil default msg"
		})
		vfy.Float64(c, (*float64)(nil), "").NotNil().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("float64 NotNil default msg", msg)
	}
}

func TestCheckFloat64_Min(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(10)), "").Min(10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Float64(c, (*float64)(nil), "param").Min(10).Msg("%s must not be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be less than 10", msg)

		vfy.Float64(c, (*float64)(nil), "").Min(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(9)), "param").Min(10).Msg("%s must not be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Float64().Min(func(ctx *vfy.Context) string {
			return "float64 Min default msg"
		})
		vfy.Float64(c, ptr(float64(9)), "param").Min(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("float64 Min default msg", msg)
	}
}

func TestCheckFloat64_Max(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(10)), "").Max(10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Float64(c, (*float64)(nil), "param").Max(10).Msg("%s must not be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be greater than 10", msg)

		vfy.Float64(c, (*float64)(nil), "").Max(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(11)), "param").Max(10).Msg("%s must not be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Float64().Max(func(ctx *vfy.Context) string {
			return "float64 Max default msg"
		})
		vfy.Float64(c, ptr(float64(11)), "param").Max(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("float64 Max default msg", msg)
	}
}

func TestCheckFloat64_Range(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(10)), "").Range(5, 10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Float64(c, (*float64)(nil), "param").Range(5, 10).Msg("%s must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)

		vfy.Float64(c, (*float64)(nil), "").Range(5, 10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(4)), "param").Range(5, 10).Msg("%s must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(11)), "param").Range(5, 10).Msg("%s must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must between 5 and 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Float64().Range(func(ctx *vfy.Context) string {
			return "float Range default msg"
		})
		vfy.Float64(c, ptr(float64(11)), "param").Range(5, 10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("float Range default msg", msg)
	}
}

func TestCheckFloat64_Gt(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		ok, msg, _ := vfy.GetResult(c)
		vfy.Float64(c, ptr(float64(11)), "param").Gt(10).Msg("test success")
		r.True(ok)
		r.Equal("", msg)

		vfy.Float64(c, (*float64)(nil), "param").Gt(10).Msg("%s must be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 10", msg)

		vfy.Float64(c, (*float64)(nil), "param").Gt(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(10)), "param").Gt(10).Msg("%s must be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be greater than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Float64().Gt(func(ctx *vfy.Context) string {
			return "float64 Gt default msg"
		})
		vfy.Float64(c, ptr(float64(10)), "param").Gt(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("float64 Gt default msg", msg)
	}
}

func TestCheckFloat64_Lt(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(9)), "param").Lt(10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Float64(c, (*float64)(nil), "param").Lt(10).Msg("%s must be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be less than 10", msg)

		vfy.Float64(c, (*float64)(nil), "param").Lt(10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(11)), "param").Lt(10).Msg("%s must be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be less than 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Float64().Lt(func(ctx *vfy.Context) string {
			return "float64 Lt default msg"
		})
		vfy.Float64(c, ptr(float64(11)), "param").Lt(10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("float64 Lt default msg", msg)
	}
}

func TestCheckFloat64_Within(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(9)), "param").Within(5, 10).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Float64(c, (*float64)(nil), "param").Within(5, 10).Msg("%s must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)

		vfy.Float64(c, (*float64)(nil), "param").Within(5, 10).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(5)), "param").Within(5, 10).Msg("%s must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(10)), "param").Within(5, 10).Msg("%s must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be > 5 and < 10", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Float64().Within(func(ctx *vfy.Context) string {
			return "float64 Within default msg"
		})
		vfy.Float64(c, ptr(float64(10)), "param").Within(5, 10).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("float64 Within default msg", msg)
	}
}

func TestCheckFloat64_Options(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(1)), "").Options([]float64{1, 2, 3}).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Float64(c, (*float64)(nil), "param").Options([]float64{1, 2, 3}).Msg("%s must be %s", c.FieldName(), c.Confines())
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 1, 2 or 3", msg)

		vfy.Float64(c, (*float64)(nil), "").Options([]float64{1, 2, 3}).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 1, 2 or 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(4)), "param").Options([]float64{1, 2, 3}).Msg("%s must be %s", c.FieldName(), c.Confines())
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be 1, 2 or 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Float64().Options(func(ctx *vfy.Context) string {
			return "float64 Options default msg"
		})
		vfy.Float64(c, ptr(float64(4)), "param").Options([]float64{1, 2, 3}).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("float64 Options default msg", msg)
	}
}

func TestCheckFloat64_Custom(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(1)), "").Custom(func(i float64) bool {
			return true
		}).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Float64(c, (*float64)(nil), "").Custom(nil).Msg("test fail by nil")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)

		vfy.Float64(c, (*float64)(nil), "").Custom(nil).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Float64(c, ptr(float64(2)), "").Custom(func(i float64) bool {
			r.Equal(float64(2), i)
			return false
		}).Msg("test fail by custom")
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by custom", msg)
	}
}
