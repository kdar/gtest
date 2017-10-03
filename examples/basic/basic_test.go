package basic

import (
	"testing"

	"github.com/kdar/gtest"
	"github.com/smartystreets/assertions"
)

func TestAdd(t *testing.T) {
	g := gtest.New(t)

	answer := Add(5, 6)
	// Error using t.Error if they do not equal.
	g.So(answer, assertions.ShouldEqual, 11).ElseError()

	answer2 := Add(-5, -6)
	// Call our function if they do not equal.
	g.So(answer2, assertions.ShouldEqual, -11).ElseErrorf("%s", gtest.MSG)

	answer3 := Add(0, 0)
	// Call t.Fatal if they do not equal.
	g.So(answer3, assertions.ShouldEqual, 0).ElseFatal()
}

func TestFailAdd(t *testing.T) {
	g := gtest.New(t)

	answer := FailAdd(5, 6)
	// Error using t.Error if they do not equal.
	g.So(answer, assertions.ShouldEqual, 11).ElseError()

	answer2 := FailAdd(-5, -6)
	// Call our function if they do not equal.
	g.So(answer2, assertions.ShouldEqual, -11).ElseErrorf("extra error info:\n%s\n", gtest.MSG)

	answer3 := FailAdd(0, 0)
	// Call t.Fatal if they do not equal.
	g.So(answer3, assertions.ShouldEqual, 0).ElseFatal()
}
