// +build ignore

package goctest

import (
	"fmt"
	"runtime"
	"strings"

	//"sync"
	"testing"

	"github.com/smartystreets/assertions"
)

func New(t *testing.T) *goctest {
	return &goctest{t: t}
}

type goctest struct {
	t *testing.T
}

func (g *goctest) Assert(actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) {
	if ok, message := assertions.So(actual, assert, expected...); !ok {
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

		g.t.Fatalf("\n%s:%d\n%s", file, line, message)
	}
}

func (g *goctest) Check(actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) bool {
	ok, message := assertions.So(actual, assert, expected...)

	if !ok {
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

		g.t.Errorf("\r%s\r\t%s:%d\n%s", getWhitespaceString(), file, line, message)
	}

	return ok
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
