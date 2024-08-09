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

func (d demoStruct) Checklist(ctx *vfy.Context) {
	ctx.Deadline()
	ctx.Done()
	ctx.Err()
	ctx.Value("")
}

func TestContext(t *testing.T) {
	r := require.New(t)
	ctx := &mockContext{}
	vfy.Check(ctx, &demoStruct{})
	r.True(ctx.invokeDeadline)
	r.True(ctx.invokeDone)
	r.True(ctx.invokeErr)
	r.True(ctx.invokeValue)
}
