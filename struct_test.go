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

type Video struct {
	Title    string
	Author   string
	Tags     []string
	Duration *int64
	Snippet  *Snippet
}

func (v Video) Checklist(ctx *vfy.Context) {
	vfy.String(ctx, &v.Title, "title").NotBlank().Msg(ctx.FieldName() + " must not be blank")
	vfy.String(ctx, &v.Author, "author").NotBlank().Msg(ctx.FieldName() + " must not be blank")
	vfy.Slices(ctx, v.Tags, "tags").Dive(func(t string) {
		vfy.String(ctx, &t, "").Max(4).Msg(ctx.FieldName() + "'s length must less than 4")
	})
	vfy.Int64(ctx, v.Duration, "duration").Range(int64(60), int64(120)).Msg(ctx.FieldName() + " must between 60 and 120 second")
	vfy.Struct(ctx, v.Snippet, "snippet").Dive()
}

type Snippet struct {
	Caption     string
	Description string
	Thumbnails  []*Thumbnail
}

func (s Snippet) Checklist(ctx *vfy.Context) {
	vfy.String(ctx, &s.Caption, "caption").NotBlank().Msg(ctx.FieldName() + " must not be blank")
	vfy.String(ctx, &s.Description, "description").Min(6).Msg(ctx.FieldName() + "'s length must greater than 6")
	vfy.Slices(ctx, s.Thumbnails, "thumbnails").Dive(func(t *Thumbnail) {
		vfy.Struct(ctx, t, "").Dive()
	})
}

type Thumbnail struct {
	Url  *string
	Size *Size
}

func (t Thumbnail) Checklist(ctx *vfy.Context) {
	vfy.String(ctx, t.Url, "url").
		NotBlank().Msg(ctx.FieldName() + " must not be blank").
		Max(20).Msg(ctx.FieldName() + " length must be less than 20")
	vfy.Struct(ctx, t.Size, "size").NotNil().Msg(ctx.FieldName() + " must not be nil").Dive()
}

type Size struct {
	Width  int32
	Height int32
}

func (s Size) Checklist(ctx *vfy.Context) {
	vfy.Int32(ctx, &s.Width, "width").Max(2000).Msg(ctx.FieldName() + " must less than 2000")
	vfy.Int32(ctx, &s.Height, "height").Max(1500).Msg(ctx.FieldName() + " must less than 1500")
}

func TestCheckStruct_NotNil(t *testing.T) {
	v := Video{}
	r := require.New(t)
	{
		c := vfy.NewDefaultContext()
		vfy.Struct(c, &v, "").NotNil().Msg("test success")
		ok, msg, _ := vfy.GetResult(c)
		r.True(ok)
		r.Equal("", msg)

		var m *Video
		vfy.Struct(c, m, "").NotNil().Msg("test fail by nil")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)

		vfy.Struct(c, (*Video)(nil), "").NotNil().Msg("test already fail")
		ok, msg, _ = vfy.GetResult(c)
		r.False(ok)
		r.Equal("test fail by nil", msg)
	}
	{
		c := vfy.NewDefaultContext()
		vfy.DefaultMsg().Struct().NotNil(func(ctx *vfy.Context) string {
			return "struct NotNil default msg"
		})
		vfy.Struct(c, (*Video)(nil), "").NotNil().DefaultMsg()
		ok, msg, _ := vfy.GetResult(c)
		r.False(ok)
		r.Equal("struct NotNil default msg", msg)
	}
}

func TestCheckStruct_Dive(t *testing.T) {
	r := require.New(t)
	v := &Video{
		Tags: []string{"1001", "1002", "10003"},
		Snippet: &Snippet{
			Description: "abcd",
			Thumbnails: []*Thumbnail{
				{},
				{
					Url:  ptr("012345678901234567890123456789"),
					Size: &Size{int32(3000), int32(2000)},
				},
			},
		},
	}
	c := vfy.NewCheckAllContext()
	v.Checklist(c)
	ok, _, msgs := vfy.GetResult(c)
	r.False(ok)
	r.Equal(12, len(msgs))
	r.Equal("title must not be blank", msgs[0])
	r.Equal("author must not be blank", msgs[1])
	r.Equal("tags[2]'s length must less than 4", msgs[2])
	r.Equal("duration must between 60 and 120 second", msgs[3])
	r.Equal("snippet.caption must not be blank", msgs[4])
	r.Equal("snippet.description's length must greater than 6", msgs[5])
	r.Equal("snippet.thumbnails[0].url must not be blank", msgs[6])
	r.Equal("snippet.thumbnails[0].url length must be less than 20", msgs[7])
	r.Equal("snippet.thumbnails[0].size must not be nil", msgs[8])
	r.Equal("snippet.thumbnails[1].url length must be less than 20", msgs[9])
	r.Equal("snippet.thumbnails[1].size.width must less than 2000", msgs[10])
	r.Equal("snippet.thumbnails[1].size.height must less than 1500", msgs[11])
}
