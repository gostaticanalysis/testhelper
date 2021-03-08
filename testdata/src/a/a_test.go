package a_test

import "testing"

func f6(t *testing.T) { // want ".* is a test helper but it does not call t.Helper"
}

// OK
func f7(t *testing.T) {
	t.Helper()
}

// OK
var f8 = func(t *testing.T) {}

// OK
func TestF9(t *testing.T) {
	// TestF8 is a test function
}

// OK
func f10() {
}

// no require checking: go vet reports an error for this function
//func Test(t *testing.T, _ interface{}) {}
