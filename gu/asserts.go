package gu

// testingT is an interface that matches the methods used by our test helpers
type testingT interface {
	Helper()
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

// assertTrue checks if a condition is true, failing the test if not
func assertTrue(t testingT, got bool, message string) {
	t.Helper()
	if !got {
		if message == "" {
			t.Fatal("expected true but got false")
		} else {
			t.Fatalf("%s: expected true but got false", message)
		}
	}
}

// assertEqual checks if two comparable values are equal
func assertEqual[T comparable](t testingT, got T, want T, message string) {
	t.Helper()
	if got != want {
		if message == "" {
			t.Fatalf("expected %v but got %v", want, got)
		} else {
			t.Fatalf("%s: expected %v but got %v", message, want, got)
		}
	}
}
