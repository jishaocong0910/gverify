package vfy_test

import (
	"testing"

	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

func TestCheckNumber_Required(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, (*int)(nil), "param").Required()
	}, "param is required")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Required()
	})
}

func TestCheckNumber_Min(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Min(6)
	}, "param must not be less than 6")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Min(5)
	})
}

func TestCheckNumber_Max(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Max(4)
	}, "param must not be greater than 4")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Max(5)
	})
}

func TestCheckNumber_Range(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Range(1, 4)
	}, "param must not be 1 to 4")
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Range(6, 9)
	}, "param must not be 6 to 9")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Range(1, 9)
	})
}

func TestCheckNumber_Gt(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Gt(5)
	}, "param must be greater than 5")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Gt(4)
	})
}

func TestCheckNumber_Lt(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Lt(5)
	}, "param must be less than 5")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Lt(6)
	})
}

func TestCheckNumber_Within(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Within(1, 5)
	}, "param must be greater than 1 and less than 5")
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Within(5, 9)
	}, "param must be greater than 5 and less than 9")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Within(1, 9)
	})
}

func TestCheckNumber_Enum(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Enum([]int{1, 2, 3})
	}, "param must be 1, 2 or 3")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Enum([]int{3, 4, 5})
	})
}

func TestCheckNumber_Custom(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Custom(func(t int) bool {
			return false
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Custom(func(t int) bool {
			return true
		})
	})
}
