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

type demoStruct struct {
}

func (d demoStruct) Checklist(vc *vfy.VContext) {
	vc.Deadline()
	vc.Done()
	vc.Err()
	vc.Value("")
}

func TestContext(t *testing.T) {
	r := require.New(t)
	vc := &mockContext{}
	vfy.Check(vc, (*demoStruct)(nil))
	r.True(vc.invokeDeadline)
	r.True(vc.invokeDone)
	r.True(vc.invokeErr)
	r.True(vc.invokeValue)
}
