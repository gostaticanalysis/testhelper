package a

import "testing"

func f1(t *testing.T) { // want ".* is a test helper but it does not call t.Helper"
}

// OK
func f2(t *testing.T) {
	t.Helper()
}

// OK
var f3 = func(t *testing.T) {
}

func TestF4(t *testing.T) { // want ".* is a test helper but it does not call t.Helper"
	// TestF4 is not a test function
}

// OK
func f5() {
}
