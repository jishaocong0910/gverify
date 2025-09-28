package vfy_test

import (
	"testing"

	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

type embedTest struct {
	field *string
}

func (e embedTest) Checklist(vc *vfy.VContext) {
	vfy.String(vc, e.field, "field").Required()
}

func TestCheckEmbed_Required(t *testing.T) {
	r := require.New(t)
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Embed(vc, (*embedTest)(nil))
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Embed(vc, &embedTest{})
	}, "field is required")
}
