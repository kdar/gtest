package gtest

import "testing"

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
	So("a", ShouldEqual, "b").Else(func(m string) {
		called1 = true
	})

	if !called1 {
		t.Error("expected callback1 to be called")
	}

	called2 := false
	So("a", ShouldEqual, "a").Else(func(m string) {
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
	So("a", ShouldEqual, "b").ElseFatal(mt)
	if !called1 {
		t.Error("expected Fatal to be called")
	}

	called2 := false
	mt.OnFatal = func() {
		called2 = true
	}
	So("a", ShouldEqual, "a").ElseFatal(mt)
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

		So("a", ShouldEqual, "b").ElseFatal()
	}()
	if !didPanic {
		t.Error("expected ElseFatal to panic when not passed testing.TB")
	}
}

func TestElseError(t *testing.T) {
	called1 := false
	mt := &mockT{}
	mt.OnError = func() {
		called1 = true
	}
	So("a", ShouldEqual, "b").ElseError(mt)
	if !called1 {
		t.Error("expected Fatal to be called")
	}

	called2 := false
	mt.OnError = func() {
		called2 = true
	}
	So("a", ShouldEqual, "a").ElseError(mt)
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

		So("a", ShouldEqual, "b").ElseError()
	}()
	if !didPanic {
		t.Error("expected ElseError to panic when not passed testing.TB")
	}
}
