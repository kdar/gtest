package gtest

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/smartystreets/assertions"
)

var (
	defaultTest = &test{}
)

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

// NewTest creates a new test object. This is not needed unless you
// want to pass in your own `t` at initialization.
func NewTest(t testing.TB) *test {
	return &test{t: t}
}

// So is a convenience function for running assertions on arbitrary arguments
// in any context, be it for testing or even application logging. It allows you
// to perform assertion-like behavior (and get nicely formatted messages detailing
// discrepancies) but without the program blowing up or panicing. All that is
// required is to import this package and call `So` with one of the assertions
// exported by this package as the second parameter.
// This function uses the default test object created by the package.
func So(actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) SoTest {
	return defaultTest.So(actual, assert, expected...)
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

// Else allows you to provide a function to be called when the test fails.
// The callback is called with one parameter with the error message.
func (t SoTest) Else(args ...interface{}) {
	if !t.ok {
		var f func(string)
		for _, v := range args {
			switch t := v.(type) {
			case func(string):
				f = t
			}
		}

		f(t.message)
	}
}

// ElseFatal is used to call t.Fatal when the test fails.
// This function will overwrite the default go file/line number
// using "\r". This is hacky and will show up in logs weird. Use
// `Else()` instead of you want to avoid this.
func (t SoTest) ElseFatal(args ...interface{}) {
	if !t.ok {
		var tt testing.TB = t.test.t
		for _, v := range args {
			switch t := v.(type) {
			case testing.TB:
				tt = t
			}
		}

		if tt == nil {
			panic("gtest: called ElseFatal but no testing.TB")
		}

		_, file, line, ok := runtime.Caller(1)
		if ok {
			if index := strings.LastIndex(file, "/"); index >= 0 {
				file = file[index+1:]
			} else if index = strings.LastIndex(file, "\\"); index >= 0 {
				file = file[index+1:]
			}
		} else {
			file = "???"
			line = 1
		}

		tt.Fatalf("\r%s\r\t%s:%d\n%s", getWhitespaceString(), file, line, t.message)
	}
}

// ElseError is used to call t.Error when the test fails.
// This function will overwrite the default go file/line number
// using "\r". This is hacky and will show up in logs weird. Use
// `Else()` instead of you want to avoid this.
func (t SoTest) ElseError(args ...interface{}) {
	if !t.ok {
		var tt testing.TB = t.test.t
		for _, v := range args {
			switch t := v.(type) {
			case testing.TB:
				tt = t
			}
		}

		if tt == nil {
			panic("gtest: called ElseError but no testing.TB")
		}

		_, file, line, ok := runtime.Caller(1)
		if ok {
			if index := strings.LastIndex(file, "/"); index >= 0 {
				file = file[index+1:]
			} else if index = strings.LastIndex(file, "\\"); index >= 0 {
				file = file[index+1:]
			}
		} else {
			file = "???"
			line = 1
		}

		tt.Errorf("\r%s\r\t%s:%d\n%s", getWhitespaceString(), file, line, t.message)
	}
}

// getWhitespaceString returns a string that is long enough to overwrite the default
// output from the go testing framework.
func getWhitespaceString() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	parts := strings.Split(file, "/")
	file = parts[len(parts)-1]

	return strings.Repeat(" ", len(fmt.Sprintf("%s:%d:      ", file, line)))
}
