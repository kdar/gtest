Several styles to consider:

## Calling your own t.Error/t.Fatal

### Original

```
if ok, msg := So(a, ShouldResemble, b); !ok {
	t.Error("\n" + msg)
}
if ok, msg := So(a, ShouldResemble, b); !ok {
	t.Fatal("\n" + msg)
}
```

### Call back function so you can call t.Error or t.Fatal

```
So(a, ShouldResemble, b).Else(func(m string) {
 t.Error(m)
})
So(a, ShouldResemble, b).Else(func(m string) {
 t.Fatal(m)
})
```

## Creating an object and having it call t.Error/t.Fatal itself.

### Example 1

```
g = New(t)
g.Check(a, ShouldResemble, b) // calls t.Error
g.Assert(a, ShouldResemble, b) // calls t.Fatal
```

### Example 1 with custom messages

```
g = New(t)
g.Msg("%d. failed because blah").Check(a, ShouldResemble, b) // calls t.Error
g.Msg("%d. failed because blah").Assert(a, ShouldResemble, b) // calls t.Fatal
```

## Example 2

```
g = New(t)
g.So(a, ShouldResemble, b) // calls t.Error
g.So(a, MustResemble, b) // calls t.Fatal
```

## Example 2 with custom messages

```
g = New(t)
g.Msg("%d. failed because blah").So(a, ShouldResemble, b) // calls t.Error
g.Msg("%d. failed because blah").So(a, MustResemble, b) // calls t.Fatal
```
