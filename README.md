
# gtest
    import "github.com/kdar/gtest"

GTest is a testing package revolving around the amazing
GoConvey assertions. It's for those of you who like its assertions
but do not like BDD style testing.

Examples:

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





## Variables
``` go
var (
    ShouldAlmostEqual            = assertions.ShouldAlmostEqual
    ShouldBeBetween              = assertions.ShouldBeBetween
    ShouldBeBetweenOrEqual       = assertions.ShouldBeBetweenOrEqual
    ShouldBeBlank                = assertions.ShouldBeBlank
    ShouldBeChronological        = assertions.ShouldBeChronological
    ShouldBeEmpty                = assertions.ShouldBeEmpty
    ShouldBeFalse                = assertions.ShouldBeFalse
    ShouldBeGreaterThan          = assertions.ShouldBeGreaterThan
    ShouldBeGreaterThanOrEqualTo = assertions.ShouldBeGreaterThanOrEqualTo
    ShouldBeIn                   = assertions.ShouldBeIn
    ShouldBeLessThan             = assertions.ShouldBeLessThan
    ShouldBeLessThanOrEqualTo    = assertions.ShouldBeLessThanOrEqualTo
    ShouldBeNil                  = assertions.ShouldBeNil
    ShouldBeTrue                 = assertions.ShouldBeTrue
    ShouldBeZeroValue            = assertions.ShouldBeZeroValue
    ShouldContain                = assertions.ShouldContain
    ShouldContainSubstring       = assertions.ShouldContainSubstring
    ShouldEndWith                = assertions.ShouldEndWith
    ShouldEqual                  = assertions.ShouldEqual
    ShouldHappenAfter            = assertions.ShouldHappenAfter
    ShouldHappenBefore           = assertions.ShouldHappenBefore
    ShouldHappenBetween          = assertions.ShouldHappenBetween
    ShouldHappenOnOrAfter        = assertions.ShouldHappenOnOrAfter
    ShouldHappenOnOrBefore       = assertions.ShouldHappenOnOrBefore
    ShouldHappenOnOrBetween      = assertions.ShouldHappenOnOrBetween
    ShouldHappenWithin           = assertions.ShouldHappenWithin
    ShouldHaveSameTypeAs         = assertions.ShouldHaveSameTypeAs
    ShouldImplement              = assertions.ShouldImplement
    ShouldNotAlmostEqual         = assertions.ShouldNotAlmostEqual
    ShouldNotBeBetween           = assertions.ShouldNotBeBetween
    ShouldNotBeBetweenOrEqual    = assertions.ShouldNotBeBetweenOrEqual
    ShouldNotBeBlank             = assertions.ShouldNotBeBlank
    ShouldNotBeEmpty             = assertions.ShouldNotBeEmpty
    ShouldNotBeIn                = assertions.ShouldNotBeIn
    ShouldNotBeNil               = assertions.ShouldNotBeNil
    ShouldNotContain             = assertions.ShouldNotContain
    ShouldNotContainSubstring    = assertions.ShouldNotContainSubstring
    ShouldNotEndWith             = assertions.ShouldNotEndWith
    ShouldNotEqual               = assertions.ShouldNotEqual
    ShouldNotHappenOnOrBetween   = assertions.ShouldNotHappenOnOrBetween
    ShouldNotHappenWithin        = assertions.ShouldNotHappenWithin
    ShouldNotHaveSameTypeAs      = assertions.ShouldNotHaveSameTypeAs
    ShouldNotImplement           = assertions.ShouldNotImplement
    ShouldNotPanic               = assertions.ShouldNotPanic
    ShouldNotPanicWith           = assertions.ShouldNotPanicWith
    ShouldNotPointTo             = assertions.ShouldNotPointTo
    ShouldNotResemble            = assertions.ShouldNotResemble
    ShouldNotStartWith           = assertions.ShouldNotStartWith
    ShouldPanic                  = assertions.ShouldPanic
    ShouldPanicWith              = assertions.ShouldPanicWith
    ShouldPointTo                = assertions.ShouldPointTo
    ShouldResemble               = assertions.ShouldResemble
    ShouldStartWith              = assertions.ShouldStartWith
)
```

## func NewTest
``` go
func NewTest(t testing.TB) *test
```
NewTest creates a new test object. This is not needed unless you
want to pass in your own `t` at initialization.



## type SoTest
``` go
type SoTest struct {
    // contains filtered or unexported fields
}
```








### func So
``` go
func So(actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) SoTest
```
So is a convenience function for running assertions on arbitrary arguments
in any context, be it for testing or even application logging. It allows you
to perform assertion-like behavior (and get nicely formatted messages detailing
discrepancies) but without the program blowing up or panicing. All that is
required is to import this package and call `So` with one of the assertions
exported by this package as the second parameter.
This function uses the default test object created by the package.




### func (SoTest) Else
``` go
func (t SoTest) Else(args ...interface{})
```
Else allows you to provide a function to be called when the test fails.
The callback is called with one parameter with the error message.



### func (SoTest) ElseError
``` go
func (t SoTest) ElseError(args ...interface{})
```
ElseError is used to call t.Error when the test fails.
This function will overwrite the default go file/line number
using "\r". This is hacky and will show up in logs weird. Use
`Else()` instead of you want to avoid this.



### func (SoTest) ElseFatal
``` go
func (t SoTest) ElseFatal(args ...interface{})
```
ElseFatal is used to call t.Fatal when the test fails.
This function will overwrite the default go file/line number
using "\r". This is hacky and will show up in logs weird. Use
`Else()` instead of you want to avoid this.









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)