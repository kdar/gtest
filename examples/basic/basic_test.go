package basic

import (
	"testing"

	"github.com/kdar/gtest"
	"github.com/smartystreets/assertions"
)

func TestAdd(t *testing.T) {
	answer := Add(5, 6)
	// Error using t.Error if they do not equal.
	gtest.So(answer, assertions.ShouldEqual, 11).ElseError(t)

	answer2 := Add(-5, -6)
	// Call our function if they do not equal.
	gtest.So(answer2, assertions.ShouldEqual, -11).Else(func(m string) {
		t.Error(m)
	})

	answer3 := Add(0, 0)
	// Call t.Fatal if they do not equal.
	gtest.So(answer3, assertions.ShouldEqual, 0).ElseFatal(t)
}

func TestFailAdd(t *testing.T) {
	answer := FailAdd(5, 6)
	// Error using t.Error if they do not equal.
	gtest.So(answer, assertions.ShouldEqual, 11).ElseError(t)

	answer2 := FailAdd(-5, -6)
	// Call our function if they do not equal.
	gtest.So(answer2, assertions.ShouldEqual, -11).Else(func(m string) {
		t.Errorf("\nextra error info: %s\n", m)
	})

	//gtest.So(answer2, assertions.ShouldEqual, -11).ElseFatalf(t, "\nextra error info: %s\n", gtest.MSG)

	answer3 := FailAdd(0, 0)
	// Call t.Fatal if they do not equal.
	gtest.So(answer3, assertions.ShouldEqual, 0).ElseFatal(t)
}
