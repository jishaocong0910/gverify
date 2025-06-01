package vfy_test

import (
	"testing"

	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

func TestCheckSlice_Required(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, ([]int)(nil), "param").Required()
	}, "param is required")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{}, "param").Required()
	})
}

func TestCheckSlice_NotEmpty(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{}, "param").NotEmpty()
	}, "param must not be empty")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1}, "param").NotEmpty()
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").NotEmpty()
	}, "param must not be empty")
}

func TestCheckSlice_Length(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1}, "param").Length(2)
	}, "param's length must be 2")
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1, 2, 3}, "param").Length(2)
	}, "param's length must be 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1, 2}, "param").Length(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Length(2)
	}, "param's length must be 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Length(0)
	})
}

func TestCheckSlice_Min(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1}, "param").Min(2)
	}, "param's length must not be less than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1, 2}, "param").Min(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Min(1)
	}, "param's length must not be less than 1")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Min(0)
	})
}

func TestCheckSlice_Max(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1, 2, 3}, "param").Max(2)
	}, "param's length must not be greater than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1, 2}, "param").Max(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Max(-1)
	}, "param's length must not be greater than -1")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Max(0)
	})
}

func TestCheckSlice_Range(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1}, "param").Range(2, 3)
	}, "param's length must be 2 to 3")
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1, 2, 3, 4}, "param").Range(2, 3)
	}, "param's length must be 2 to 3")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1, 2, 3}, "param").Range(2, 3)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Range(1, 3)
	}, "param's length must be 1 to 3")
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Range(-3, -1)
	}, "param's length must be -3 to -1")
}

func TestCheckSlice_Gt(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1, 2}, "param").Gt(2)
	}, "param's length must be greater than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1, 2, 3}, "param").Gt(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Gt(0)
	}, "param's length must be greater than 0")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Gt(-1)
	})
}

func TestCheckSlice_Lt(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1, 2}, "param").Lt(2)
	}, "param's length must be less than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1}, "param").Lt(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Lt(0)
	}, "param's length must be less than 0")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Lt(1)
	})
}

func TestCheckSlice_Within(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1}, "param").Within(2, 5)
	}, "param's length must be greater than 2 and less than 5")
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1, 2, 3, 4, 5, 6}, "param").Within(2, 5)
	}, "param's length must be greater than 2 and less than 5")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{1, 2, 3}, "param").Within(2, 5)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Within(0, 5)
	}, "param's length must be greater than 0 and less than 5")
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Within(-5, 0)
	}, "param's length must be greater than -5 and less than 0")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Within(-1, 5)
	})
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Within(-5, 1)
	})
}

func TestCheckSlice_Custom(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{}, "param").Custom(true, func(t []int) bool {
			return false
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int{}, "param").Custom(false, func(t []int) bool {
			return true
		})
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Custom(false, func(t []int) bool {
			return true
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Slice(vc, []int(nil), "param").Custom(true, func(t []int) bool {
			return false
		})
	})
}

func TestCheckSlice_Interrupt(t *testing.T) {
	r := require.New(t)
	{
		vc := vfy.NewDefaultContext()
		vfy.SetAll(vc)
		vfy.Slice(vc, []int{8, 3}, "param").Dive(func(t int) {
			vfy.Int(vc, &t, "").Lt(7).Gt(4)
		})
		_, _, msgs := vfy.GetResult(vc)
		r.Len(msgs, 2)
		r.Equal("param[0] must be less than 7", msgs[0])
		r.Equal("param[1] must be greater than 4", msgs[1])
	}
	{
		vc := vfy.NewDefaultContext()
		vfy.Slice(vc, []int{8, 3}, "param").Dive(func(t int) {
			vfy.Int(vc, &t, "").Lt(7).Gt(4)
		})
		_, _, msgs := vfy.GetResult(vc)
		r.Len(msgs, 1)
		r.Equal("param[0] must be less than 7", msgs[0])
	}
}
