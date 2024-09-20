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
	"strconv"
	"testing"

	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

func TestCheckSlice_NotNil(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello", "world"}, "param").NotNil().Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").NotNil().Msg("%s must not be nil", c.FieldName())
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be nil", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").NotNil().Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Slices().NotNil(func(ctx *vfy.Context) string {
			return "slices NotNil default msg"
		})
		vfy.Slices[string](c, ([]string)(nil), "param").NotNil().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("slices NotNil default msg", msg)
	}
}

func TestCheckSlice_NotEmpty(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello", "world"}, "param").NotEmpty().Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").NotEmpty().Msg("%s must not be empty", c.FieldName())
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be empty", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").NotEmpty().Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be empty", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{}, "param").NotEmpty().Msg("%s must not be empty", c.FieldName())
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be empty", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Slices().NotEmpty(func(ctx *vfy.Context) string {
			return "slices NotEmpty default msg"
		})
		vfy.Slices[string](c, []string{}, "param").NotEmpty().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("slices NotEmpty default msg", msg)
	}
}

func TestCheckSlice_Length(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello", "world"}, "param").Length(2).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").Length(2).Msg("%s's length must be %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be 2", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").Length(2).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{}, "param").Length(2).Msg("%s's length must be %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Slices().Length(func(ctx *vfy.Context) string {
			return "slices Length default msg"
		})
		vfy.Slices[string](c, []string{}, "param").Length(2).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("slices Length default msg", msg)
	}
}

func TestCheckSlice_Min(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello", "world"}, "param").Min(2).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").Min(2).Msg("%s's length must not be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be less than 2", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").Min(2).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be less than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello"}, "param").Min(2).Msg("%s's length must not be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be less than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Slices().Min(func(ctx *vfy.Context) string {
			return "slices Min default msg"
		})
		vfy.Slices[string](c, []string{"hello"}, "param").Min(2).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("slices Min default msg", msg)
	}
}

func TestCheckSlice_Max(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello", "world"}, "param").Max(2).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").Max(2).Msg("%s's length must not be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be greater than 2", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").Max(2).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be greater than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello", "world", "!"}, "param").Max(2).Msg("%s's length must not be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be greater than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Slices().Max(func(ctx *vfy.Context) string {
			return "slices Max default msg"
		})
		vfy.Slices[string](c, []string{"hello", "world", "!"}, "param").Max(2).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("slices Max default msg", msg)
	}
}

func TestCheckSlice_Range(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello", "world"}, "param").Range(1, 2).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").Range(1, 2).Msg("%s's length must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must between 1 and 2", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").Range(1, 2).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must between 1 and 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{}, "param").Range(1, 2).Msg("%s's length must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must between 1 and 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello", "world", "!"}, "param").Range(1, 2).Msg("%s's length must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must between 1 and 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Slices().Range(func(ctx *vfy.Context) string {
			return "slices Range default msg"
		})
		vfy.Slices[string](c, []string{"hello", "world", "!"}, "param").Range(1, 2).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("slices Range default msg", msg)
	}
}

func TestCheckSlice_Gt(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello", "world", "!"}, "param").Gt(2).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").Gt(2).Msg("%s's length must be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be greater than 2", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").Gt(2).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be greater than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello"}, "param").Gt(2).Msg("%s's length must be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be greater than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Slices().Gt(func(ctx *vfy.Context) string {
			return "slices Gt default msg"
		})
		vfy.Slices[string](c, []string{"hello"}, "param").Gt(2).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("slices Gt default msg", msg)
	}
}

func TestCheckSlice_Lt(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello"}, "param").Lt(2).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").Lt(2).Msg("%s's length must be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be less than 2", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").Lt(2).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be less than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Slices().Lt(func(ctx *vfy.Context) string {
			return "slices Lt default msg"
		})
		vfy.Slices[string](c, []string{"hello", "world", "!"}, "param").Lt(2).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("slices Lt default msg", msg)
	}
}

func TestCheckSlice_Within(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello", "world"}, "param").Within(1, 3).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").Within(1, 3).Msg("%s's length must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be > 1 and < 3", msg)

		vfy.Slices[string](c, ([]string)(nil), "param").Within(1, 3).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be > 1 and < 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{}, "param").Within(1, 3).Msg("%s's length must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be > 1 and < 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello", "world", "!"}, "param").Within(1, 3).Msg("%s's length must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be > 1 and < 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Slices().Within(func(ctx *vfy.Context) string {
			return "slices Within default msg"
		})
		vfy.Slices[string](c, []string{"hello", "world", "!"}, "param").Within(1, 3).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("slices Within default msg", msg)
	}
}

func TestCheckSlice_Dive(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello", "world"}, "param").Dive(func(t string) {
			vfy.String(c, &t, "").Length(5).Msg("%s's length must be %s", c.FieldName(), c.Confine(0))
		})
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetChecklistFalse(c)
		vfy.Slices[string](c, []string{"hello", "world", "!"}, "param").Dive(nil)
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Slices[string](c, []string{"hello", "world", "!"}, "param").Dive(func(t string) {
			vfy.String(c, &t, "").Length(5).Msg("%s's length must be %s", c.FieldName(), c.Confine(0))
		})
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param[2]'s length must be 5", msg)
	}
	{
		m := MyStruct{
			MySlice: []string{"", ""},
		}
		c := vfy.NewCheckAllContext()
		vfy.Struct(c, &m, "myStruct").Dive()
		ok, _, msgs := vfy.GetResult(c)
		r.False(ok)
		r.Len(msgs, 2)
		r.Equal("myStruct.mySlice[0] must not be blank", msgs[0])
		r.Equal("myStruct.mySlice[1] must not be blank", msgs[1])
	}
	{
		m := MyStruct2{
			MySlice: []string{"", ""},
		}
		c := vfy.NewCheckAllContext()
		vfy.Struct(c, &m, "myStruct").Dive()
		ok, _, msgs := vfy.GetResult(c)
		r.False(ok)
		r.Len(msgs, 2)
		r.Equal("myStruct.mySlice[1] must not be blank", msgs[0])
		r.Equal("myStruct.mySlice[2] must not be blank", msgs[1])
	}
	{
		c := vfy.NewCheckAllContext()
		vfy.Slices[MyStruct3](c, []MyStruct3{{}}, "param").Dive(func(t MyStruct3) {
			vfy.Struct(c, &t, "").Dive()
		})
		ok, _, msgs := vfy.GetResult(c)
		r.False(ok)
		r.Len(msgs, 1)
		r.Equal("param[0].myField must not be blank", msgs[0])
	}
	{
		c := vfy.NewCheckAllContext()
		vfy.Slices[[]MyStruct3](c, [][]MyStruct3{{{MyField: "a"}, {}}}, "param").Dive(func(t []MyStruct3) {
			vfy.Slices(c, t, "").Dive(func(t MyStruct3) {
				vfy.Struct(c, &t, "").Dive()
			})
		})
		ok, _, msgs := vfy.GetResult(c)
		r.False(ok)
		r.Len(msgs, 1)
		r.Equal("param[0][1].myField must not be blank", msgs[0])
	}
}

type MyStruct struct {
	MySlice []string
}

func (m MyStruct) Checklist(ctx *vfy.Context) {
	vfy.Slices(ctx, m.MySlice, "mySlice").Dive(func(e string) {
		vfy.String(ctx, &e, "").NotBlank().Msg("%s must not be blank", ctx.FieldName())
	})
}

type MyStruct2 struct {
	MySlice []string
}

func (m MyStruct2) Checklist(ctx *vfy.Context) {
	vfy.Slices(ctx, m.MySlice, "mySlice").Dive(func(e string) {
		vfy.String(ctx, &e, "["+strconv.Itoa(ctx.Index()+1)+"]").NotBlank().Msg("%s must not be blank", ctx.FieldName())
	})
}

type MyStruct3 struct {
	MyField string
}

func (s MyStruct3) Checklist(ctx *vfy.Context) {
	vfy.String(ctx, &s.MyField, "myField").NotBlank().Msg("%s must not be blank", ctx.FieldName())
}
