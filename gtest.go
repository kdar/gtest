package gtest

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/smartystreets/assertions"
)

type kind int

const MSG kind = 0

type SoTest struct {
	// actual   interface{}
	// assert   func(actual interface{}, expected ...interface{}) string
	// expected []interface{}
	message string
	ok      bool
	test    *test
}

type test struct {
	//prefix string
	t testing.TB
}

// New creates a new test object. This is not needed unless you
// want to pass in your own `t` at initialization.
func New(t testing.TB) *test {
	return &test{t: t}
}

// So is a convenience function for running assertions on arbitrary arguments
// in any context, be it for testing or even application logging. It allows you
// to perform assertion-like behavior (and get nicely formatted messages detailing
// discrepancies) but without the program blowing up or panicing. All that is
// required is to import this package and call `So` with one of the assertions
// exported by this package as the second parameter.
func (t *test) So(actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) SoTest {
	ok, message := assertions.So(actual, assert, expected...)
	return SoTest{
		ok:      ok,
		message: message,
		test:    t,
	}
}

func (t *test) Assert(actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) bool {
	ok, message := assertions.So(actual, assert, expected...)

	if !ok {
		if t.t == nil {
			panic("gtest: called Require but no testing.TB")
		}

		_, file, line, _ := runtime.Caller(1)
		t.t.Errorf("\n%s:%d\n%s", file, line, message)
	}

	return ok
}

func (t *test) Require(actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) bool {
	ok, message := assertions.So(actual, assert, expected...)

	if !ok {
		if t.t == nil {
			panic("gtest: called Require but no testing.TB")
		}

		_, file, line, _ := runtime.Caller(1)
		t.t.Fatalf("\n%s:%d\n%s", file, line, message)
	}

	return ok
}

// ElseFatal is used to call t.Fatal when the test fails.
func (t SoTest) ElseFatal() {
	if !t.ok {
		tt := t.test.t

		if tt == nil {
			panic("gtest: called ElseFatal but no testing.TB")
		}

		_, file, line, _ := runtime.Caller(1)
		tt.Fatalf("\n%s:%d\n%s", file, line, t.message)
	}
}

// ElseError is used to call t.Error when the test fails.
// This function will overwrite the default go file/line number
// using "\r". This is hacky and will show up in logs weird. Use
// `Else()` instead of you want to avoid this.
func (t SoTest) ElseError() {
	if !t.ok {
		tt := t.test.t

		if tt == nil {
			panic("gtest: called ElseError but no testing.TB")
		}

		_, file, line, _ := runtime.Caller(1)
		tt.Errorf("\n%s:%d\n%s", file, line, t.message)
	}
}

// ElseErrorf is just like ElseError except it lets you format your
// message. Pass gtest.MSG to the formatter to dictate where it will
// be placed. E.g.
// g.So(answer, assertions.ShouldEqual, -11).ElseErrorf("%s", gtest.MSG)
func (t SoTest) ElseErrorf(format string, args ...interface{}) {
	if !t.ok {
		tt := t.test.t
		if tt == nil {
			panic("gtest: called ElseErrorf but no testing.TB")
		}

		for i := range args {
			switch args[i].(type) {
			case kind:
				args[i] = t.message
			}
		}

		_, file, line, _ := runtime.Caller(1)
		tt.Errorf("\n%s:%d\n%s", file, line, fmt.Sprintf(format, args...))
	}
}

// ElseFatalf is just like ElseErrorf except it calls t.Fatalf.
func (t SoTest) ElseFatalf(format string, args ...interface{}) {
	if !t.ok {
		tt := t.test.t
		if tt == nil {
			panic("gtest: called ElseFatalf but no testing.TB")
		}

		for i := range args {
			switch args[i].(type) {
			case kind:
				args[i] = t.message
			}
		}

		_, file, line, _ := runtime.Caller(1)
		tt.Fatalf("\n%s:%d\n%s", file, line, fmt.Sprintf(format, args...))
	}
}
