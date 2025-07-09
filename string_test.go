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
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").NotBlank()
	}, "param must not be blank")
}

func TestCheckString_Regex(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("aaa"), "param").Regex(regexp.MustCompile(`\d+`))
	}, "param's format is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("123"), "param").Regex(regexp.MustCompile(`\d+`))
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Regex(regexp.MustCompile(`\d+`))
	}, "param's format is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Regex(regexp.MustCompile(`\d*`))
	})
}

func TestCheckString_Custom(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("a"), "param").Custom(true, func(t string) bool {
			return false
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("a"), "param").Custom(false, func(t string) bool {
			return true
		})
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Custom(false, func(t string) bool {
			return true
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Custom(true, func(t string) bool {
			return false
		})
	})
}

func TestCheckString_Length(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中"), "param").Length(2)
	}, "param's length must be 2")
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中文中"), "param").Length(2)
	}, "param's length must be 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中文"), "param").Length(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Length(1)
	}, "param's length must be 1")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Length(0)
	})
}

func TestCheckString_Min(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中"), "param").Min(2)
	}, "param's length must not be less than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中文"), "param").Min(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Min(1)
	}, "param's length must not be less than 1")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Min(0)
	})
}

func TestCheckString_Max(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中文中"), "param").Max(2)
	}, "param's length must not be greater than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中文"), "param").Max(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Max(-1)
	}, "param's length must not be greater than -1")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Max(0)
	})
}

func TestCheckString_Range(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中"), "param").Range(2, 3)
	}, "param's length must be 2 to 3")
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中文中文"), "param").Range(2, 3)
	}, "param's length must be 2 to 3")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中文中"), "param").Range(2, 3)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Range(1, 3)
	}, "param's length must be 1 to 3")
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Range(-3, -1)
	}, "param's length must be -3 to -1")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Range(0, 3)
	})
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Range(-3, 0)
	})
}

func TestCheckString_Gt(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中文"), "param").Gt(2)
	}, "param's length must be greater than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中文中"), "param").Gt(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Gt(0)
	}, "param's length must be greater than 0")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Gt(-1)
	})
}

func TestCheckString_Lt(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中文"), "param").Lt(2)
	}, "param's length must be less than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中"), "param").Lt(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Lt(0)
	}, "param's length must be less than 0")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Lt(1)
	})
}

func TestCheckString_Within(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中"), "param").Within(2, 5)
	}, "param's length must be greater than 2 and less than 5")
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中文中文中文"), "param").Within(2, 5)
	}, "param's length must be greater than 2 and less than 5")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("中文中"), "param").Within(2, 5)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Within(0, 5)
	}, "param's length must be greater than 0 and less than 5")
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Within(-5, 0)
	}, "param's length must be greater than -5 and less than 0")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Within(-1, 5)
	})
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Within(-5, 1)
	})
}

func TestCheckString_Enum(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("d"), "param").Enum([]string{"a", "b", "c"})
	}, `param must be "a", "b" or "c"`)
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.String(vc, ptr("c"), "param").Enum([]string{"a", "b", "c"})
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.String(vc, (*string)(nil), "param").Enum([]string{"a", "b", "c"})
	}, `param must be "a", "b" or "c"`)
}
