
# gtest [![GoDoc](https://godoc.org/github.com/kdar/gtest?status.png)](http://godoc.org/github.com/kdar/gtest)

    import "github.com/kdar/gtest"

gtest is a testing package revolving around the amazing
GoConvey assertions. It's for those of you who like its assertions
but do not like BDD style testing.

## Examples

Do a test and call t.Fatal if it fails

	g := NewTest(t)
	g.So(a, should.Resemble, b).ElseFatal()

Do a test and call t.Error if it fails

	g := NewTest(t)
	g.So(a, should.Resemble, b).ElseError()

Do a test and call t.Fatalf

	g := NewTest(t)
	g.So(a, should.Resemble, b).ElseFatalf("failed at index %d\n%s\n", i, gotest.MSG)
