
# gtest [![GoDoc](https://godoc.org/github.com/kdar/gtest?status.png)](http://godoc.org/github.com/kdar/gtest) 

    import "github.com/kdar/gtest"

gtest is a testing package revolving around the amazing
GoConvey assertions. It's for those of you who like its assertions
but do not like BDD style testing.

## Examples

Do a test and call t.Fatal if it fails

	So(a, ShouldResemble, b).ElseFatal(t)

Do a test and call t.Error if it fails

	So(a, ShouldResemble, b).ElseError(t)

Do a test and call your own function if it fails

	So(a, ShouldResemble, b).Else(func(msg string) {
		t.Fatalf("\nfailed at index %d\n%s", i, msg)
	})

Make your own test object so you don't need to pass `t` every time.

	g := NewTest(t)
	g.So(a, ShouldResemble, b).ElseError()
