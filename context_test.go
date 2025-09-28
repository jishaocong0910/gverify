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
	"time"

	vfy "github.com/jishaocong0910/gverify"
	"github.com/stretchr/testify/require"
)

type mockContext struct {
	invokeDeadline,
	invokeDone,
	invokeErr,
	invokeValue bool
}

func (m *mockContext) Deadline() (deadline time.Time, ok bool) {
	m.invokeDeadline = true
	return time.Now(), false
}

func (m *mockContext) Done() <-chan struct{} {
	m.invokeDone = true
	return nil
}

func (m *mockContext) Err() error {
	m.invokeErr = true
	return nil
}

func (m *mockContext) Value(key any) any {
	m.invokeValue = true
	return nil
}

type contextTest struct {
}

func (r contextTest) Checklist(vc *vfy.VContext) {
	vc.Deadline()
	vc.Done()
	vc.Err()
	vc.Value("")
}

func TestContext(t *testing.T) {
	r := require.New(t)
	vc := &mockContext{}
	vfy.Check(vc, &contextTest{})
	r.True(vc.invokeDeadline)
	r.True(vc.invokeDone)
	r.True(vc.invokeErr)
	r.True(vc.invokeValue)
}
