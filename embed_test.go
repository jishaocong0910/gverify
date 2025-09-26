package vfy_test

import (
	"testing"

	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

func TestCheckEmbed_Required(t *testing.T) {
	r := require.New(t)
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Embed(vc, (*tmp)(nil))
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Struct(vc, &tmp{}, "param").Required()
	}, "")
}
