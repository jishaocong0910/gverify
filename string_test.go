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
	"regexp"
	"testing"

	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

func TestCheckString_NotNil(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello"), "param").NotNil().Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.String(c, (*string)(nil), "param").NotNil().Msg("test fail by nil")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)

		vfy.String(c, (*string)(nil), "param").NotNil().Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, (*string)(nil), "param").NotNil().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().String().NotNil(func(ctx *vfy.Context) string {
			return "string NotNil default setMsg"
		})
		vfy.String(c, (*string)(nil), "param").NotNil().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("string NotNil default setMsg", msg)
	}
}

func TestCheckString_NotBlank(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello"), "param").NotBlank().Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.String(c, (*string)(nil), "param").NotBlank().Msg("%s must not be blank", c.FieldName())
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be blank", msg)

		vfy.String(c, (*string)(nil), "param").NotBlank().Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be blank", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr(""), "param").NotBlank().Msg("%s must not be blank", c.FieldName())
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be blank", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr(""), "param").NotBlank().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be blank", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().String().NotBlank(func(ctx *vfy.Context) string {
			return "string NotBlank default setMsg"
		})
		vfy.String(c, ptr(""), "param").NotBlank().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("string NotBlank default setMsg", msg)
	}
}

func TestCheckString_Length(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello"), "param").Length(5).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.String(c, (*string)(nil), "param").Length(5).Msg("%s's length must be %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be 5", msg)

		vfy.String(c, (*string)(nil), "param").Length(5).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hi"), "param").Length(5).Msg("%s's length must be %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hi"), "param").Length(5).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().String().Length(func(ctx *vfy.Context) string {
			return "string Length default setMsg"
		})
		vfy.String(c, ptr("hi"), "param").Length(5).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("string Length default setMsg", msg)
	}
}

func TestCheckString_Regex(t *testing.T) {
	r := require.New(t)
	reg := regexp.MustCompile(`^https://.+`)
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("https://google.com"), "param").Regex(reg).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.String(c, (*string)(nil), "param").Regex(reg).Msg("%s's format is illegal", c.FieldName())
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's format is illegal", msg)

		vfy.String(c, (*string)(nil), "param").Regex(reg).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's format is illegal", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("http://google.com"), "param").Regex(reg).Msg("%s's format is illegal", c.FieldName())
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's format is illegal", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("http://google.com"), "param").Regex(reg).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's format is illegal", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().String().Regex(func(ctx *vfy.Context) string {
			return "string Regex default setMsg"
		})
		vfy.String(c, ptr("http://google.com"), "param").Regex(reg).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("string Regex default setMsg", msg)
	}
}

func TestCheckString_Min(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello"), "param").Min(5).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.String(c, (*string)(nil), "param").Min(5).Msg("%s's length must not be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be less than 5", msg)

		vfy.String(c, (*string)(nil), "param").Min(5).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be less than 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hi"), "param").Min(5).Msg("%s's length must not be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be less than 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hi"), "param").Min(5).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be less than 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().String().Min(func(ctx *vfy.Context) string {
			return "string Min default setMsg"
		})
		vfy.String(c, ptr("hi"), "param").Min(5).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("string Min default setMsg", msg)
	}
}

func TestCheckString_Max(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello"), "param").Max(5).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.String(c, (*string)(nil), "param").Max(5).Msg("%s's length must not be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be greater than 5", msg)

		vfy.String(c, (*string)(nil), "param").Max(5).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be greater than 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello!"), "param").Max(5).Msg("%s's length must not be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be greater than 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello!"), "param").Max(5).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be greater than 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().String().Max(func(ctx *vfy.Context) string {
			return "string Max default setMsg"
		})
		vfy.String(c, ptr("hello!"), "param").Max(5).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("string Max default setMsg", msg)
	}
}

func TestCheckString_Range(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello"), "param").Range(3, 5).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.String(c, (*string)(nil), "param").Range(3, 5).Msg("%s's length must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must between 3 and 5", msg)

		vfy.String(c, (*string)(nil), "param").Range(3, 5).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must between 3 and 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hi"), "param").Range(3, 5).Msg("%s's length must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must between 3 and 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello!"), "param").Range(3, 5).Msg("%s's length must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must between 3 and 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello!"), "param").Range(3, 5).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be 3 to 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().String().Range(func(ctx *vfy.Context) string {
			return "string Range default setMsg"
		})
		vfy.String(c, ptr("hello!"), "param").Range(3, 5).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("string Range default setMsg", msg)
	}
}

func TestCheckString_Gt(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello!"), "param").Gt(5).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.String(c, (*string)(nil), "param").Gt(5).Msg("%s's length must be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be greater than 5", msg)

		vfy.String(c, (*string)(nil), "param").Gt(5).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be greater than 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hi"), "param").Gt(5).Msg("%s's length must be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be greater than 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hi"), "param").Gt(5).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be greater than 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().String().Gt(func(ctx *vfy.Context) string {
			return "string Gt default setMsg"
		})
		vfy.String(c, ptr("hi"), "param").Gt(5).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("string Gt default setMsg", msg)
	}
}

func TestCheckString_Lt(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hi"), "param").Lt(5).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.String(c, (*string)(nil), "param").Lt(5).Msg("%s's length must be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be less than 5", msg)

		vfy.String(c, (*string)(nil), "param").Lt(5).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be less than 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello!"), "param").Lt(5).Msg("%s's length must be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be less than 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello!"), "param").Lt(5).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be less than 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().String().Lt(func(ctx *vfy.Context) string {
			return "string Lt default setMsg"
		})
		vfy.String(c, ptr("hello!"), "param").Lt(5).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("string Lt default setMsg", msg)
	}
}

func TestCheckString_Within(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("abcd"), "param").Within(3, 5).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.String(c, (*string)(nil), "param").Within(3, 5).Msg("%s's length must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be > 3 and < 5", msg)

		vfy.String(c, (*string)(nil), "param").Within(3, 5).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be > 3 and < 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hi"), "param").Within(3, 5).Msg("%s's length must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be > 3 and < 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello"), "param").Within(3, 5).Msg("%s's length must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be > 3 and < 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("hello"), "param").Within(3, 5).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be greater than 3 and less than 5", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().String().Within(func(ctx *vfy.Context) string {
			return "string Within default setMsg"
		})
		vfy.String(c, ptr("hello"), "param").Within(3, 5).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("string Within default setMsg", msg)
	}
}

func TestCheckString_Options(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("b"), "param").Options([]string{"a", "b", "c"}).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.String(c, (*string)(nil), "param").Options([]string{"a", "b", "c"}).Msg("%s must be %s", c.FieldName(), c.Confines())
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be a, b or c", msg)

		vfy.String(c, (*string)(nil), "param").Options(nil).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be a, b or c", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("d"), "param").Options([]string{"a", "b", "c"}).Msg("%s must be %s", c.FieldName(), c.Confines())
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be a, b or c", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("d"), "param").Options([]string{"a", "b", "c"}).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must be a, b or c", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetDefaultMsg().String().Options(func(ctx *vfy.Context) string {
			return "string Options default setMsg"
		})
		vfy.String(c, ptr("d"), "param").Options([]string{"a", "b", "c"}).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("string Options default setMsg", msg)
	}
}

func TestCheckString_Custom(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("str"), "param").Custom(func(s string) bool {
			return true
		}).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.String(c, (*string)(nil), "param").Custom(nil).Msg("test fail by nil")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)

		vfy.String(c, (*string)(nil), "param").Custom(nil).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.String(c, ptr("str"), "param").Custom(func(s string) bool {
			r.Equal("str", s)
			return false
		}).Msg("test fail by custom")
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by custom", msg)
	}
}
