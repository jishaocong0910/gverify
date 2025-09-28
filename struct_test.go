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

type structTest struct {
	field *string
}

func (t structTest) Checklist(vc *vfy.VContext) {
	vfy.String(vc, t.field, "field").Required()
}

func TestCheckStruct_Required(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Struct(vc, (*structTest)(nil), "param").Required()
	}, "param is required")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Struct(vc, &structTest{}, "param").Required()
	})
}

func TestCheckStruct_Custom(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Struct(vc, &structTest{}, "param").Custom(true, func(t structTest) bool {
			return false
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Struct(vc, &structTest{}, "param").Custom(false, func(t structTest) bool {
			return true
		})
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Struct(vc, (*structTest)(nil), "param").Custom(false, func(t structTest) bool {
			return true
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Struct(vc, (*structTest)(nil), "param").Custom(true, func(t structTest) bool {
			return false
		})
	})
}
