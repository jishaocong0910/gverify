package vfy_test

import (
	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

func testSuccess(r *require.Assertions, f func(vc *vfy.VContext)) {
	vc := vfy.NewDefaultContext()
	f(vc)
	code, _, _ := vfy.GetResult(vc)
	r.Equal(vfy.SUCCESS, code)
}

func testFail(r *require.Assertions, f func(vc *vfy.VContext), expectedMsg string) {
	vc := vfy.NewDefaultContext()
	f(vc)
	code, msg, _ := vfy.GetResult(vc)
	r.Equal(vfy.FAIL, code)
	r.Equal(expectedMsg, msg)
}
