package gtest

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/smartystreets/assertions"
)

type kind int

const MSG kind = 0

// var (
// 	defaultTest = &test{}
// )

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
// This function uses the default test object created by the package.
// func So(actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) SoTest {
// 	return defaultTest.So(actual, assert, expected...)
// }

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
// func (t SoTest) Else(args ...interface{}) {
// 	if !t.ok {
// 		var f func(string)
// 		for _, v := range args {
// 			switch t := v.(type) {
// 			case func(string):
// 				f = t
// 			}
// 		}

// 		_, file, line, _ := runtime.Caller(1)
// 		f(fmt.Sprintf("\n%s:%d\n%s", file, line, t.message))
// 	}
// }

// ElseFatal is used to call t.Fatal when the test fails.
func (t SoTest) ElseFatal() {
	if !t.ok {
		tt := t.test.t
		// for _, v := range args {
		// 	switch t := v.(type) {
		// 	case testing.TB:
		// 		tt = t
		// 	}
		// }

		if tt == nil {
			panic("gtest: called ElseFatal but no testing.TB")
		}

		//ci := CallerInfo()
		//tt.Fatalf("\r%s\r\t%s\n%s\n\n", getWhitespaceString(), ci[len(ci)-1], t.message)
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
		// for _, v := range args {
		// 	switch t := v.(type) {
		// 	case testing.TB:
		// 		tt = t
		// 	}
		// }

		if tt == nil {
			panic("gtest: called ElseError but no testing.TB")
		}

		_, file, line, _ := runtime.Caller(1)
		tt.Errorf("\n%s:%d\n%s", file, line, t.message)

		//ci := CallerInfo()
		//tt.Errorf("\r%s\r\t%s\n%s\n\n", getWhitespaceString(), ci[len(ci)-1], t.message)
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

// ElseFataf is just like ElseErrorf except it calls t.Fatalf.
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

// Stolen from: https://github.com/stretchr/testify
// Copyright (c) 2012 - 2013 Mat Ryer and Tyler Bunnell
func CallerInfo() []string {
	pc := uintptr(0)
	file := ""
	line := 0
	ok := false
	name := ""

	callers := []string{}
	for i := 0; ; i++ {
		pc, file, line, ok = runtime.Caller(i)
		if !ok {
			// The breaks below failed to terminate the loop, and we ran off the
			// end of the call stack.
			break
		}

		// This is a huge edge case, but it will panic if this is the case.
		// See https://github.com/stretchr/testify/issues/180
		if file == "<autogenerated>" {
			break
		}

		f := runtime.FuncForPC(pc)
		if f == nil {
			break
		}
		name = f.Name()

		// testing.tRunner is the standard library function that calls
		// tests. Subtests are called directly by tRunner, without going through
		// the Test/Benchmark/Example function that contains the t.Run calls, so
		// with subtests we should break when we hit tRunner, without adding it
		// to the list of callers.
		if name == "testing.tRunner" {
			break
		}

		parts := strings.Split(file, "/")
		file = parts[len(parts)-1]
		if len(parts) > 1 {
			dir := parts[len(parts)-2]
			if (dir != "assert" && dir != "mock" && dir != "require") || file == "mock_test.go" {
				callers = append(callers, fmt.Sprintf("%s:%d", file, line))
			}
		}

		// Drop the package
		segments := strings.Split(name, ".")
		name = segments[len(segments)-1]
		if isTest(name, "Test") ||
			isTest(name, "Benchmark") ||
			isTest(name, "Example") {
			break
		}
	}

	return callers
}

// isTest tells whether name looks like a test (or benchmark, according to prefix).
// It is a Test (say) if there is a character after Test that is not a lower-case letter.
// We don't want TesticularCancer.
// Stolen from the `go test` tool.
func isTest(name, prefix string) bool {
	if !strings.HasPrefix(name, prefix) {
		return false
	}
	if len(name) == len(prefix) { // "Test" is ok
		return true
	}
	rune, _ := utf8.DecodeRuneInString(name[len(prefix):])
	return !unicode.IsLower(rune)
}

// getWhitespaceString returns a string that is long enough to overwrite the default
// output from the go testing framework.
// Stolen from: https://github.com/stretchr/testify
// Copyright (c) 2012 - 2013 Mat Ryer and Tyler Bunnell
func getWhitespaceString() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	parts := strings.Split(file, "/")
	file = parts[len(parts)-1]

	return strings.Repeat(" ", len(fmt.Sprintf("%s:%d:      ", file, line)))
}
