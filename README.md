# testhelper

[![pkg.go.dev][gopkg-badge]][gopkg]

`testhelper` finds a package function which is not a test function and receives a value of `*testing.T` as a parameter but it does not call `(*testing.T).Helper`.

```go
// OK
func helper1(t *testing.T) {
	t.Helper()
}

// NG
func helper2(t *testing.T) {
	// without t.Helper
}
```

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gostaticanalysis/testhelper
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gostaticanalysis/testhelper?status.svg
