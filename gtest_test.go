package gtest

import (
	"testing"

	"github.com/smartystreets/assertions/should"
)

type mockT struct {
	*testing.T
	OnError func()
	OnFatal func()
}

func (t *mockT) Error(args ...interface{}) {
	if t.OnError != nil {
		t.OnError()
	}
}
func (t *mockT) Errorf(format string, args ...interface{}) {
	if t.OnError != nil {
		t.OnError()
	}
}
func (t *mockT) Fatal(args ...interface{}) {
	if t.OnFatal != nil {
		t.OnFatal()
	}
}
func (t *mockT) Fatalf(format string, args ...interface{}) {
	if t.OnFatal != nil {
		t.OnFatal()
	}
}

func TestElse(t *testing.T) {
	called1 := false
	So("a", should.Equal, "b").Else(func(m string) {
		called1 = true
	})

	if !called1 {
		t.Error("expected callback1 to be called")
	}

	called2 := false
	So("a", should.Equal, "a").Else(func(m string) {
		called2 = true
	})

	if called2 {
		t.Error("expected callback2 not to be called")
	}
}

func TestNewElse(t *testing.T) {
	a := New(t)

	called1 := false
	a.So("a", should.Equal, "b").Else(func(m string) {
		called1 = true
	})

	if !called1 {
		t.Error("expected callback1 to be called")
	}

	called2 := false
	a.So("a", should.Equal, "a").Else(func(m string) {
		called2 = true
	})

	if called2 {
		t.Error("expected callback2 not to be called")
	}
}

func TestElseFatal(t *testing.T) {
	called1 := false
	mt := &mockT{}
	mt.OnFatal = func() {
		called1 = true
	}
	So("a", should.Equal, "b").ElseFatal(mt)
	if !called1 {
		t.Error("expected Fatal to be called")
	}

	called2 := false
	mt.OnFatal = func() {
		called2 = true
	}
	So("a", should.Equal, "a").ElseFatal(mt)
	if called2 {
		t.Error("expected Fatal not to be called")
	}

	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()

		So("a", should.Equal, "b").ElseFatal()
	}()
	if !didPanic {
		t.Error("expected ElseFatal to panic when not passed testing.TB")
	}
}

func TestNewElseFatal(t *testing.T) {
	a := New(t)

	called1 := false
	mt := &mockT{}
	mt.OnFatal = func() {
		called1 = true
	}
	a.So("a", should.Equal, "b").ElseFatal(mt)
	if !called1 {
		t.Error("expected Fatal to be called")
	}

	called2 := false
	mt.OnFatal = func() {
		called2 = true
	}
	a.So("a", should.Equal, "a").ElseFatal(mt)
	if called2 {
		t.Error("expected Fatal not to be called")
	}
}

func TestElseError(t *testing.T) {
	called1 := false
	mt := &mockT{}
	mt.OnError = func() {
		called1 = true
	}
	So("a", should.Equal, "b").ElseError(mt)
	if !called1 {
		t.Error("expected Fatal to be called")
	}

	called2 := false
	mt.OnError = func() {
		called2 = true
	}
	So("a", should.Equal, "a").ElseError(mt)
	if called2 {
		t.Error("expected Fatal not to be called")
	}

	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()

		So("a", should.Equal, "b").ElseError()
	}()
	if !didPanic {
		t.Error("expected ElseError to panic when not passed testing.TB")
	}
}

func TestNewElseError(t *testing.T) {
	a := New(t)

	called1 := false
	mt := &mockT{}
	mt.OnError = func() {
		called1 = true
	}
	a.So("a", should.Equal, "b").ElseError(mt)
	if !called1 {
		t.Error("expected Fatal to be called")
	}

	called2 := false
	mt.OnError = func() {
		called2 = true
	}
	a.So("a", should.Equal, "a").ElseError(mt)
	if called2 {
		t.Error("expected Fatal not to be called")
	}
}
