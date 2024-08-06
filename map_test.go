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
	"strings"
	"testing"

	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

func TestCheckMap_NotNil(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2}, "param").NotNil().Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Map[string, int](c, nil, "param").NotNil().Msg("test fail by nil")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)

		vfy.Map[string, int](c, nil, "param").NotNil().Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Map().NotNil(func(ctx *vfy.Context) string {
			return "map NotNil default msg"
		})
		vfy.Map[string, int](c, nil, "param").NotNil().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("map NotNil default msg", msg)
	}
}

func TestCheckMap_NotEmpty(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2}, "param").NotEmpty().Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Map[string, int](c, nil, "param").NotEmpty().Msg("%s must not be empty", c.FieldName())
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be empty", msg)

		vfy.Map[string, int](c, nil, "param").NotEmpty().Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be empty", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{}, "param").NotEmpty().Msg("%s must not be empty", c.FieldName())
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param must not be empty", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Map().NotEmpty(func(ctx *vfy.Context) string {
			return "map NotEmpty default msg"
		})
		vfy.Map[string, int](c, map[string]int{}, "param").NotEmpty().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("map NotEmpty default msg", msg)
	}
}

func TestCheckMap_Length(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2}, "param").Length(2).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Map[string, int](c, nil, "param").Length(2).Msg("test fail by nil")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)

		vfy.Map[string, int](c, nil, "param").Length(2).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{}, "param").Length(2).Msg("%s's length must be %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Map().Length(func(ctx *vfy.Context) string {
			return "map Length default msg"
		})
		vfy.Map[string, int](c, map[string]int{}, "param").Length(2).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("map Length default msg", msg)
	}
}

func TestCheckMap_Min(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2}, "param").Min(2).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Map[string, int](c, nil, "param").Min(2).Msg("%s's length must not be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be less than 2", msg)

		vfy.Map[string, int](c, nil, "param").Min(2).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be less than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1}, "param").Min(2).Msg("%s's length must not be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be less than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Map().Min(func(ctx *vfy.Context) string {
			return "map Min default msg"
		})
		vfy.Map[string, int](c, map[string]int{"key1": 1}, "param").Min(2).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("map Min default msg", msg)
	}
}

func TestCheckMap_Max(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2}, "param").Max(2).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Map[string, int](c, nil, "param").Max(2).Msg("%s's length must not be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be greater than 2", msg)

		vfy.Map[string, int](c, nil, "param").Max(2).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be greater than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2, "C": 67}, "param").Max(2).Msg("%s's length must not be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must not be greater than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Map().Max(func(ctx *vfy.Context) string {
			return "map Max default msg"
		})
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2, "C": 67}, "param").Max(2).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("map Max default msg", msg)
	}
}

func TestCheckMap_Range(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2}, "param").Range(1, 2).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Map[string, int](c, nil, "param").Range(1, 2).Msg("%s' length must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param' length must between 1 and 2", msg)

		vfy.Map[string, int](c, nil, "param").Range(1, 2).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param' length must between 1 and 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{}, "param").Range(1, 2).Msg("%s' length must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param' length must between 1 and 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2, "C": 67}, "param").Range(1, 2).Msg("%s' length must between %s and %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param' length must between 1 and 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Map().Range(func(ctx *vfy.Context) string {
			return "map Range default msg"
		})
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2, "C": 67}, "param").Range(1, 2).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("map Range default msg", msg)
	}
}

func TestCheckMap_Gt(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2, "key3": 3}, "param").Gt(2).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Map[string, int](c, nil, "param").Gt(2).Msg("%s's length must be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be greater than 2", msg)

		vfy.Map[string, int](c, nil, "param").Gt(2).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be greater than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1}, "param").Gt(2).Msg("%s's length must be greater than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be greater than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Map().Gt(func(ctx *vfy.Context) string {
			return "map Gt default msg"
		})
		vfy.Map[string, int](c, map[string]int{"key1": 1}, "param").Gt(2).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("map Gt default msg", msg)
	}
}

func TestCheckMap_Lt(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1}, "param").Lt(2).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Map[string, int](c, nil, "param").Lt(2).Msg("%s's length must be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be less than 2", msg)

		vfy.Map[string, int](c, nil, "param").Lt(2).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be less than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2, "C": 67}, "param").Lt(2).Msg("%s's length must be less than %s", c.FieldName(), c.Confine(0))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param's length must be less than 2", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Map().Lt(func(ctx *vfy.Context) string {
			return "map Lt default msg"
		})
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2, "C": 67}, "param").Lt(2).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("map Lt default msg", msg)
	}
}

func TestCheckMap_Within(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2}, "param").Within(1, 3).Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		vfy.Map[string, int](c, nil, "param").Within(1, 3).Msg("%s' length must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param' length must be > 1 and < 3", msg)

		vfy.Map[string, int](c, nil, "param").Within(1, 3).Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("param' length must be > 1 and < 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{}, "param").Within(1, 3).Msg("%s' length must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param' length must be > 1 and < 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2, "C": 67}, "param").Within(1, 3).Msg("%s' length must be > %s and < %s", c.FieldName(), c.Confine(0), c.Confine(1))
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param' length must be > 1 and < 3", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Map().Within(func(ctx *vfy.Context) string {
			return "map Within default msg"
		})
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2, "C": 67}, "param").Within(1, 3).DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("map Within default msg", msg)
	}
}

func TestCheckMap_Dive(t *testing.T) {
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2}, "map").Dive(func(k string) {
			vfy.String(c, &k, "").Custom(func(k string) bool {
				return strings.Index(k, "key") == 0
			}).Msg(c.FieldName() + " must start with 'key'")
		}, func(v int) {
			vfy.Int(c, &v, "").Range(1, 99).Msg(c.FieldName() + " must between 1 to 99")
		})
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.SetChecklistFalse(c)
		vfy.Map[string, int](c, map[string]int{"key1": 1, "key2": 2}, "map").Dive(nil, nil)
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"k1": 1, "key2": 2}, "param").Dive(func(k string) {
			vfy.String(c, &k, "").Custom(func(k string) bool {
				return strings.Index(k, "key") == 0
			}).Msg(c.FieldName() + " must start with 'key'")
		}, func(v int) {
			vfy.Int(c, &v, "").Range(1, 99).Msg(c.FieldName() + " must between 1 to 99")
		})
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param$key must start with 'key'", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.Map[string, int](c, map[string]int{"key1": 0, "key2": 2}, "param").Dive(func(k string) {
			vfy.String(c, &k, "").Custom(func(k string) bool {
				return strings.Index(k, "key") == 0
			}).Msg(c.FieldName() + " must start with 'key'")
		}, func(v int) {
			vfy.Int(c, &v, "").Range(1, 99).Msg(c.FieldName() + " must between 1 to 99")
		})
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("param$value must between 1 to 99", msg)
	}
}
