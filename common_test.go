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
	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

func testSuccess(r *require.Assertions, f func(vc *vfy.VContext)) {
	vc := vfy.NewDefaultContext()
	f(vc)
	code, _, _ := vfy.GetResult(vc)
	r.Equal(vfy.SUCCESS, code)
}

func testFail(r *require.Assertions, f func(vc *vfy.VContext), expectedMsg string) {
	vc := vfy.NewDefaultContext()
	f(vc)
	code, msg, _ := vfy.GetResult(vc)
	r.Equal(vfy.FAIL, code)
	r.Equal(expectedMsg, msg)
}
