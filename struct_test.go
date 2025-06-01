package vfy_test

import (
	"testing"

	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

type tmp struct {
	name string
}

func (d tmp) Checklist(vc *vfy.VContext) {
}

func TestCheckStruct_Required(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Struct(vc, (*tmp)(nil), "param").Required()
	}, "param is required")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Struct(vc, &tmp{}, "param").Required()
	})
}

func TestCheckStruct_Custom(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Struct(vc, &tmp{}, "param").Custom(true, func(t tmp) bool {
			return false
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Struct(vc, &tmp{}, "param").Custom(false, func(t tmp) bool {
			return true
		})
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Struct(vc, (*tmp)(nil), "param").Custom(false, func(t tmp) bool {
			return true
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Struct(vc, (*tmp)(nil), "param").Custom(true, func(t tmp) bool {
			return false
		})
	})
}
