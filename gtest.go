package gtest

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"testing"

	"github.com/smartystreets/assertions"
)

var (
	defaultTest = &test{}
)

type sotest struct {
	actual   interface{}
	assert   func(actual interface{}, expected ...interface{}) string
	expected []interface{}
	test     *test
}

type test struct {
	Prefix string
	t      *testing.T
}

// NewTest creates a new test object. This is not needed unless you
// want to pass in your own `t` at initialization.
func NewTest(t *testing.T) *test {
	return &test{t: t}
}

// So is a convenience function for running assertions on arbitrary arguments
// in any context, be it for testing or even application logging. It allows you
// to perform assertion-like behavior (and get nicely formatted messages detailing
// discrepancies) but without the program blowing up or panicing. All that is
// required is to import this package and call `So` with one of the assertions
// exported by this package as the second parameter.
// This function uses the default test object created by the package.
func So(actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) sotest {
	return defaultTest.So(actual, assert, expected...)
}

// So is a convenience function for running assertions on arbitrary arguments
// in any context, be it for testing or even application logging. It allows you
// to perform assertion-like behavior (and get nicely formatted messages detailing
// discrepancies) but without the program blowing up or panicing. All that is
// required is to import this package and call `So` with one of the assertions
// exported by this package as the second parameter.
func (t *test) So(actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) sotest {
	return sotest{
		actual:   actual,
		assert:   assert,
		expected: expected,
		test:     t,
	}
}

// Else allows you to provide a function to be called when the test fails.
// The callback is called with one parameter with the error message.
func (t sotest) Else(f func(string)) {
	ok, message := assertions.So(t.actual, t.assert, t.expected...)
	if !ok {
		f(message)
	}
}

// ElseFatal is called when the test fails, and will call t.Fatal.
// This function will overwrite the default go file/line number
// using "\r". This is hacky and will show up in logs weird. Use
// `Else()` instead of you want to avoid this.
func (t sotest) ElseFatal(tt ...*testing.T) {
	assertok, message := assertions.So(t.actual, t.assert, t.expected...)
	if !assertok {
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

		testingt := t.test.t
		if testingt == nil {
			if len(tt) == 0 || tt[0] == nil {
				log.Fatal("gtest: called ElseFatal but no *testing.T")
			}
			testingt = tt[0]
		}

		testingt.Fatalf("\r%s\r\t%s:%d\n%s", getWhitespaceString(), file, line, message)
	}
}

// ElseError is called when the test fails, and will call t.Error.
// This function will overwrite the default go file/line number
// using "\r". This is hacky and will show up in logs weird. Use
// `Else()` instead of you want to avoid this.
func (t sotest) ElseError(tt ...*testing.T) {
	assertok, message := assertions.So(t.actual, t.assert, t.expected...)
	if !assertok {
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

		testingt := t.test.t
		if testingt == nil {
			if len(tt) == 0 || tt[0] == nil {
				log.Fatal("gtest: called ElseError but no *testing.T")
			}
			testingt = tt[0]
		}

		testingt.Errorf("\r%s\r\t%s:%d\n%s", getWhitespaceString(), file, line, message)
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
