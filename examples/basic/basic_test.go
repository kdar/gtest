package basic

import (
	"testing"

	. "github.com/kdar/gtest"
)

func TestAdd(t *testing.T) {
	answer := Add(5, 6)
	// Error using t.Error if they do not equal.
	So(answer, ShouldEqual, 11).ElseError(t)

	answer2 := Add(-5, -6)
	// Call our function if they do not equal.
	So(answer2, ShouldEqual, -11).Else(func(m string) {
		t.Error(m)
	})

	answer3 := Add(0, 0)
	// Call t.Fatal if they do not equal.
	So(answer3, ShouldEqual, 0).ElseFatal(t)
}
