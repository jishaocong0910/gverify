package vfy_test

import (
	"testing"

	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

func TestCheckAny_Required(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Any(vc, (*int)(nil), "param").Required()
	}, "param is required")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Any(vc, ptr(5), "param").Required()
	})
}

func TestCheckAny_Custom(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Any(vc, ptr(5), "param").Custom(func(t int) bool {
			return false
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Any(vc, ptr(5), "param").Custom(func(t int) bool {
			return true
		})
	})
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Any(vc, (*int)(nil), "param", vfy.Omittable()).Custom(func(t int) bool {
			return false
		})
	})
}
