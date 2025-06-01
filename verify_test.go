package vfy_test

import (
	"testing"

	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

func ptr[T any](t T) *T {
	return &t
}

func TestCheckPredicate(t *testing.T) {
	r := require.New(t)
	{
		// 成功
		vc := vfy.NewDefaultContext()
		vfy.CheckPredicate(vc, ptr(5), nil, nil, func() bool {
			return false
		}, func() bool {
			return true
		})
		code, _, _ := vfy.GetResult(vc)
		r.Equal(vfy.SUCCESS, code)
	}
	{
		// 指定code和msg
		vc := vfy.NewDefaultContext()
		vfy.SetFieldName(vc, "param")
		vfy.CheckPredicate(vc, ptr(5), []vfy.CheckOption{vfy.Code("MY_CODE"), vfy.Msg(func(b *vfy.FieldInfo) {
			b.Msg("%s %s", b.FieldName(), b.Confines())
		})}, func() []string {
			return []string{"a", "b", "c"}
		}, func() bool {
			return true
		}, func() bool {
			return false
		})
		code, msg, _ := vfy.GetResult(vc)
		r.Equal("MY_CODE", code)
		r.Equal("param a, b or c", msg)
	}
	{
		// 默认code和msg
		vc := vfy.NewDefaultContext()
		vfy.SetFieldName(vc, "param")
		vfy.CheckPredicate(vc, ptr(5), nil, nil, func() bool {
			return true
		}, func() bool {
			return false
		})
		code, msg, _ := vfy.GetResult(vc)
		r.Equal(vfy.FAIL, code)
		r.Equal("param is illegal", msg)
	}
	{
		// nil导致错误
		vc := vfy.NewDefaultContext()
		vfy.SetFieldName(vc, "param")
		vfy.CheckPredicate(vc, (*int)(nil), nil, nil, func() bool {
			return false
		}, func() bool {
			return true
		})
		code, _, _ := vfy.GetResult(vc)
		r.Equal(vfy.FAIL, code)
	}
	{
		// 忽略nil
		vc := vfy.NewDefaultContext()
		vfy.SetOmittable(vc)
		vfy.CheckPredicate(vc, (*int)(nil), nil, nil, func() bool {
			return false
		}, func() bool {
			return false
		})
		code, _, _ := vfy.GetResult(vc)
		r.Equal(vfy.SUCCESS, code)
	}
	{
		// 中断
		vc := vfy.NewDefaultContext()
		vfy.SetHasWrong(vc)
		vfy.CheckPredicate(vc, (*int)(nil), []vfy.CheckOption{vfy.Code("MY_CODE")}, nil, func() bool {
			return false
		}, func() bool {
			return false
		})
		code, _, _ := vfy.GetResult(vc)
		r.Equal(vfy.FAIL, code)
	}
}

func testNumberToConfine[T vfy.Number](r *require.Assertions, ntsf func(t ...T) []string) {
	confines := ntsf(1, 20, 100)
	r.Len(confines, 3)
	r.Contains(confines, "1")
	r.Contains(confines, "20")
	r.Contains(confines, "100")
}

func TestNumberToConfine(t *testing.T) {
	r := require.New(t)
	testNumberToConfine(r, vfy.IntToConfine)
	testNumberToConfine(r, vfy.Int8ToConfine)
	testNumberToConfine(r, vfy.Int16ToConfine)
	testNumberToConfine(r, vfy.Int32ToConfine)
	testNumberToConfine(r, vfy.Int64ToConfine)
	testNumberToConfine(r, vfy.UintToConfine)
	testNumberToConfine(r, vfy.Uint8ToConfine)
	testNumberToConfine(r, vfy.Uint16ToConfine)
	testNumberToConfine(r, vfy.Uint32ToConfine)
	testNumberToConfine(r, vfy.Uint64ToConfine)
	testNumberToConfine(r, vfy.Float32ToConfine)
	testNumberToConfine(r, vfy.Float64ToConfine)
}

type user struct {
	name   string
	attach attach
}

func (u user) Checklist(vc *vfy.VContext) {
	vfy.String(vc, &u.name, "name").NotBlank()
	vfy.Struct(vc, &u.attach, "attach").Dive()
}

type attach struct {
	id      string
	images  []image
	score   map[string]int
	schools map[string]school
}

func (a attach) Checklist(vc *vfy.VContext) {
	vfy.String(vc, &a.id, "id").NotBlank().Length(10)
	vfy.Slice(vc, a.images, "images").Dive(func(i image) {
		vfy.Struct(vc, &i, "not work").Dive()
	})
	vfy.Map(vc, a.score, "score").Dive(nil, func(v int) {
		vfy.Int(vc, &v, "not work").Max(100)
	})
	vfy.Map(vc, a.schools, "schools").Dive(func(k string) {
		vfy.String(vc, &k, "not work").Max(4)
	}, func(v school) {
		vfy.Struct(vc, &v, "not work").Dive()
	})
}

type image struct {
	url string
}

func (i image) Checklist(vc *vfy.VContext) {
	vfy.String(vc, &i.url, "url").NotBlank()
}

type school struct {
	Name string
}

func (s school) Checklist(vc *vfy.VContext) {
	vfy.String(vc, &s.Name, "name").NotBlank()
}

func TestDive(t *testing.T) {
	r := require.New(t)
	u := user{
		attach: attach{
			images:  []image{{}, {}},
			score:   map[string]int{"math": 101},
			schools: map[string]school{"123456": {}},
		},
	}

	code, msg, msgs := vfy.Check(nil, u)
	r.Equal(vfy.FAIL, code)
	r.Equal("name must not be blank", msg)
	r.Len(msgs, 1)

	code, _, msgs = vfy.Check(nil, &u, vfy.All())
	r.Equal(vfy.FAIL, code)
	r.Len(msgs, 8)
	r.Equal("name must not be blank", msgs[0])
	r.Equal("attach.id must not be blank", msgs[1])
	r.Equal("attach.id's length must be 10", msgs[2])
	r.Equal("attach.images[0].url must not be blank", msgs[3])
	r.Equal("attach.images[1].url must not be blank", msgs[4])
	r.Equal("attach.score$value must not be greater than 100", msgs[5])
	r.Equal("attach.schools$key's length must not be greater than 4", msgs[6])
	r.Equal("attach.schools$value.name must not be blank", msgs[7])
}

func TestFieldVerifyFunc(t *testing.T) {
	r := require.New(t)
	vc := vfy.NewDefaultContext()
	r.NotNil(vfy.Bool(vc, ptr(true), "param"))
	r.NotNil(vfy.Byte(vc, ptr(byte(1)), "param"))
	r.NotNil(vfy.Int(vc, ptr(1), "param"))
	r.NotNil(vfy.Int8(vc, ptr(int8(1)), "param"))
	r.NotNil(vfy.Int16(vc, ptr(int16(1)), "param"))
	r.NotNil(vfy.Int32(vc, ptr(int32(1)), "param"))
	r.NotNil(vfy.Int64(vc, ptr(int64(1)), "param"))
	r.NotNil(vfy.Uint(vc, ptr(uint(1)), "param"))
	r.NotNil(vfy.Uint8(vc, ptr(uint8(1)), "param"))
	r.NotNil(vfy.Uint16(vc, ptr(uint16(1)), "param"))
	r.NotNil(vfy.Uint32(vc, ptr(uint32(1)), "param"))
	r.NotNil(vfy.Uint64(vc, ptr(uint64(1)), "param"))
	r.NotNil(vfy.Float32(vc, ptr(float32(0.1)), "param"))
	r.NotNil(vfy.Float64(vc, ptr(0.1), "param"))
}
