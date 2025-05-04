package vfy_test

import (
	"regexp"
	"testing"

	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

func TestCheckString_Required(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Required()
	}, "param is required")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("a"), "param").Required()
	})
}

func TestCheckString_NotBlank(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr(""), "param").NotBlank()
	}, "param must not be blank")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("a"), "param").NotBlank()
	})
}

func TestCheckString_Regex(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("aaa"), "param").Regex(regexp.MustCompile(`\d+`))
	}, "param's format is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("123"), "param").Regex(regexp.MustCompile(`\d+`))
	})
}

func TestCheckString_Custom(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("a"), "param").Custom(func(t string) bool {
			return false
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("a"), "param").Custom(func(t string) bool {
			return true
		})
	})
}

func TestCheckString_Length(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("a"), "param").Length(2)
	}, "param's length must be 2")
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("aaa"), "param").Length(2)
	}, "param's length must be 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("aa"), "param").Length(2)
	})
}

func TestCheckString_Min(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("a"), "param").Min(2)
	}, "param's length must not be less than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("aa"), "param").Min(2)
	})
}

func TestCheckString_Max(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("aaa"), "param").Max(2)
	}, "param's length must not be greater than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("aa"), "param").Max(2)
	})
}

func TestCheckString_Range(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("a"), "param").Range(2, 3)
	}, "param's length must be 2 to 3")
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("aaaa"), "param").Range(2, 3)
	}, "param's length must be 2 to 3")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("aaa"), "param").Range(2, 3)
	})
}

func TestCheckString_Gt(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("aa"), "param").Gt(2)
	}, "param's length must be greater than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("aaa"), "param").Gt(2)
	})
}

func TestCheckString_Lt(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("aa"), "param").Lt(2)
	}, "param's length must be less than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("a"), "param").Lt(2)
	})
}

func TestCheckString_Within(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("a"), "param").Within(2, 5)
	}, "param's length must be greater than 2 and less than 5")
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("aaaaaa"), "param").Within(2, 5)
	}, "param's length must be greater than 2 and less than 5")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("aaa"), "param").Within(2, 5)
	})
}

func TestCheckString_Enum(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("d"), "param").Enum([]string{"a", "fieldInfo", "c"})
	}, "param must be a, fieldInfo or c")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("c"), "param").Enum([]string{"a", "fieldInfo", "c"})
	})
}
