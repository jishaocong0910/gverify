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

func ptr[T any](t T) *T {
	return &t
}

type Image struct {
	Url string
}

func (i Image) Checklist(c *vfy.Context) {
	vfy.String(c, &i.Url, "url").NotBlank().Msg("%s must not be blank", c.FieldName())
}

func TestCheck(t *testing.T) {
	r := require.New(t)
	{
		i := &Image{}
		ok, msg := vfy.Check(i)
		r.False(ok)
		r.Equal("url must not be blank", msg)
	}
	{
		i := Image{}
		ok, msg := vfy.Check(i)
		r.False(ok)
		r.Equal("url must not be blank", msg)
	}
	{
		var i *Image
		ok, msg := vfy.Check(i)
		r.False(ok)
		r.Equal("url must not be blank", msg)
	}
}
