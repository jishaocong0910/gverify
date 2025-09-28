/*
Copyright 2024-present jishaocong0910

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vfy_test

import (
	"testing"

	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

func TestCheckMap_Required(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, (map[string]int)(nil), "param").Required()
	}, "param is required")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{}, "param").Required()
	})
}

func TestCheckMap_NotEmpty(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{}, "param").NotEmpty()
	}, "param must not be empty")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1}, "param").NotEmpty()
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, (map[string]int)(nil), "param").NotEmpty()
	}, "param must not be empty")
}

func TestCheckMap_Length(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1}, "param").Length(2)
	}, "param's length must be 2")
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1, "b": 2, "c": 3}, "param").Length(2)
	}, "param's length must be 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1, "b": 2}, "param").Length(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Length(2)
	}, "param's length must be 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Length(0)
	})
}

func TestCheckMap_Min(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1}, "param").Min(2)
	}, "param's length must not be less than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1, "b": 2}, "param").Min(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Min(1)
	}, "param's length must not be less than 1")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Min(0)
	})
}

func TestCheckMap_Max(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1, "b": 2, "c": 3}, "param").Max(2)
	}, "param's length must not be greater than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1, "b": 2}, "param").Max(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Max(-1)
	}, "param's length must not be greater than -1")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Max(0)
	})
}

func TestCheckMap_Range(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1}, "param").Range(2, 3)
	}, "param's length must be 2 to 3")
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, "param").Range(2, 3)
	}, "param's length must be 2 to 3")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1, "b": 2, "c": 3}, "param").Range(2, 3)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Range(1, 3)
	}, "param's length must be 1 to 3")
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Range(-3, -1)
	}, "param's length must be -3 to -1")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Range(0, 3)
	})
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Range(-1, 0)
	})
}

func TestCheckMap_Gt(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1, "b": 2}, "param").Gt(2)
	}, "param's length must be greater than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1, "b": 2, "c": 3}, "param").Gt(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Gt(0)
	}, "param's length must be greater than 0")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Gt(-1)
	})
}

func TestCheckMap_Lt(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1, "b": 2, "c": 3}, "param").Lt(2)
	}, "param's length must be less than 2")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1}, "param").Lt(2)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Lt(0)
	}, "param's length must be less than 0")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Lt(1)
	})
}

func TestCheckMap_Within(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1}, "param").Within(2, 5)
	}, "param's length must be greater than 2 and less than 5")
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6}, "param").Within(2, 5)
	}, "param's length must be greater than 2 and less than 5")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{"a": 1, "b": 2, "c": 3}, "param").Within(2, 5)
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Within(0, 5)
	}, "param's length must be greater than 0 and less than 5")
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Within(-5, 0)
	}, "param's length must be greater than -5 and less than 0")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Within(-1, 5)
	})
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Within(-5, 1)
	})
}

func TestCheckMap_Custom(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{}, "param").Custom(true, func(t map[string]int) bool {
			return false
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int{}, "param").Custom(false, func(t map[string]int) bool {
			return true
		})
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Custom(false, func(t map[string]int) bool {
			return true
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Map(vc, map[string]int(nil), "param").Custom(true, func(t map[string]int) bool {
			return true
		})
	})
}

func TestCheckMap_Interrupt(t *testing.T) {
	r := require.New(t)
	{
		vc := vfy.NewDefaultContext()
		vfy.SetAll(vc)
		vfy.Map(vc, map[string]int{"aaa": 5, "bbb": 5, "a": 11, "b": 11}, "param").Dive(func(k string) {
			vfy.String(vc, &k, "").Lt(3)
		}, func(v int) {
			vfy.Int(vc, &v, "").Lt(10)
		})
		_, _, msgs := vfy.GetResult(vc)
		r.Len(msgs, 4)
	}
	{
		vc := vfy.NewDefaultContext()
		vfy.Map(vc, map[string]int{"aaa": 5, "bbb": 5}, "param").Dive(func(k string) {
			vfy.String(vc, &k, "").Lt(3)
		}, func(v int) {
			vfy.Int(vc, &v, "").Lt(10)
		})
		_, _, msgs := vfy.GetResult(vc)
		r.Len(msgs, 1)
	}
	{
		vc := vfy.NewDefaultContext()
		vfy.Map(vc, map[string]int{"a": 11, "b": 11}, "param").Dive(func(k string) {
			vfy.String(vc, &k, "").Lt(3)
		}, func(v int) {
			vfy.Int(vc, &v, "").Lt(10)
		})
		_, _, msgs := vfy.GetResult(vc)
		r.Len(msgs, 1)
	}
	{
		vc := vfy.NewDefaultContext()
		vfy.SetHasWrong(vc)
		vfy.Map(vc, map[string]int{"a": 11, "b": 11}, "param").Dive(func(k string) {
			vfy.String(vc, &k, "").Lt(3)
		}, func(v int) {
			vfy.Int(vc, &v, "").Lt(10)
		})
		_, _, msgs := vfy.GetResult(vc)
		r.Nil(msgs)
	}
}
