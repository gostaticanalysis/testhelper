package a_test

import "testing"

func f5(t *testing.T) { // want ".* is a test helper but it does not call t.Helper"
}

// OK
func f6(t *testing.T) {
	t.Helper()
}

// OK
var f7 = func(t *testing.T) {}

// OK
func TestF8(t *testing.T) {
	// TestF8 is a test function
}

// no require checking: go vet reports an error for this function
//func Test(t *testing.T, _ interface{}) {}
