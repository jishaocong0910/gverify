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
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, (*int)(nil), "param").Min(0)
	}, "param must not be less than 0")
}

func TestCheckNumber_Max(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Max(4)
	}, "param must not be greater than 4")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Max(5)
	})
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, (*int)(nil), "param").Max(0)
	})
}

func TestCheckNumber_Range(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Range(1, 4)
	}, "param must be 1 to 4")
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Range(6, 9)
	}, "param must be 6 to 9")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Range(1, 9)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, (*int)(nil), "param").Range(0, 9)
	}, "param must be 0 to 9")
}

func TestCheckNumber_Gt(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Gt(5)
	}, "param must be greater than 5")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Gt(4)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, (*int)(nil), "param").Gt(0)
	}, "param must be greater than 0")
}

func TestCheckNumber_Lt(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Lt(5)
	}, "param must be less than 5")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Lt(6)
	})
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, (*int)(nil), "param").Lt(0)
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
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, (*int)(nil), "param").Within(-5, 5)
	}, "param must be greater than -5 and less than 5")
}

func TestCheckNumber_Enum(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Enum([]int{1, 2, 3})
	}, "param must be 1, 2 or 3")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Enum([]int{3, 4, 5})
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, (*int)(nil), "param").Enum([]int{1, 2, 3})
	}, "param must be 1, 2 or 3")
}

func TestCheckNumber_Custom(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Custom(true, func(t int) bool {
			return false
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, ptr(5), "param").Custom(false, func(t int) bool {
			return true
		})
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Int(vc, (*int)(nil), "param").Custom(false, func(t int) bool {
			return true
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Int(vc, (*int)(nil), "param").Custom(true, func(t int) bool {
			return false
		})
	})
}
