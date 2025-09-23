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

func TestCheckAny_Required(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Any(vc, (*int)(nil), "param").Required()
	}, "param is required")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Any(vc, ptr(5), "param").Required()
	})
}

func TestCheckAny_Custom(t *testing.T) {
	r := require.New(t)
	testFail(r, func(vc *vfy.VContext) {
		vfy.Any(vc, ptr(5), "param").Custom(true, func(t int) bool {
			return false
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Any(vc, ptr(5), "param").Custom(false, func(t int) bool {
			return true
		})
	})
	testFail(r, func(vc *vfy.VContext) {
		vfy.Any(vc, (*any)(nil), "param").Custom(false, func(t any) bool {
			return true
		})
	}, "param is illegal")
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Any(vc, (*any)(nil), "param").Custom(true, func(t any) bool {
			return false
		})
	})
	testSuccess(r, func(vc *vfy.VContext) {
		vfy.Any(vc, (*any)(nil), "param", vfy.Omittable()).Custom(false, func(t any) bool {
			return false
		})
	})
}
