package gu

import (
	"strings"
	"testing"
)

// fakeT is a test double that implements the methods needed by assertTrue and assertEqual
// It captures calls without actually failing the test
type fakeT struct {
	testing.TB
	helperCalled bool
	fatalCalled  bool
	fatalMessage string
}

func (f *fakeT) Helper() {
	f.helperCalled = true
}

func (f *fakeT) Fatal(args ...interface{}) {
	f.fatalCalled = true
	if len(args) > 0 {
		f.fatalMessage = args[0].(string)
	}
}

func (f *fakeT) Fatalf(format string, args ...interface{}) {
	f.fatalCalled = true
	// Don't actually format to avoid complexity, just concatenate for testing purposes
	// This is a simplified version that's sufficient for our test verification
	if len(args) > 0 {
		// For our tests, we just need to know that both the prefix and suffix are present
		f.fatalMessage = format
		// Store args separately if needed for more detailed checks
	} else {
		f.fatalMessage = format
	}
}

// TestAssertTrueSuccess tests assertTrue with passing conditions
func TestAssertTrueSuccess(t *testing.T) {
	// Test with real t to ensure normal success path works
	assertTrue(t, true, "")
	assertTrue(t, true, "with message")
	assertTrue(t, 1 == 2-1, "")
}

// TestAssertTrueFailureEmptyMessage tests assertTrue failure with empty message
func TestAssertTrueFailureEmptyMessage(t *testing.T) {
	fake := &fakeT{TB: t}
	assertTrue(fake, false, "")

	if !fake.fatalCalled {
		t.Error("Expected Fatal to be called")
	}
	if !fake.helperCalled {
		t.Error("Expected Helper to be called")
	}
	if !strings.Contains(fake.fatalMessage, "expected true but got false") {
		t.Errorf("Expected error message to contain 'expected true but got false', got: %s", fake.fatalMessage)
	}
}

// TestAssertTrueFailureWithMessage tests assertTrue failure with custom message
func TestAssertTrueFailureWithMessage(t *testing.T) {
	fake := &fakeT{TB: t}
	assertTrue(fake, false, "custom message")

	if !fake.fatalCalled {
		t.Error("Expected Fatal to be called")
	}
	if !fake.helperCalled {
		t.Error("Expected Helper to be called")
	}
	// The format string will be "%s: expected true but got false"
	if !strings.Contains(fake.fatalMessage, "%s") {
		t.Errorf("Expected format string to contain '%%s', got: %s", fake.fatalMessage)
	}
	if !strings.Contains(fake.fatalMessage, "expected true but got false") {
		t.Errorf("Expected error message to contain 'expected true but got false', got: %s", fake.fatalMessage)
	}
}

// TestAssertEqualSuccess tests assertEqual with equal values
func TestAssertEqualSuccess(t *testing.T) {
	// Test with real t to ensure normal success path works
	assertEqual(t, 42, 42, "")
	assertEqual(t, "hello", "hello", "")
	assertEqual(t, true, true, "with message")
}

// TestAssertEqualFailureEmptyMessage tests assertEqual failure with empty message
func TestAssertEqualFailureEmptyMessage(t *testing.T) {
	fake := &fakeT{TB: t}
	assertEqual(fake, 1, 2, "")

	if !fake.fatalCalled {
		t.Error("Expected Fatal to be called")
	}
	if !fake.helperCalled {
		t.Error("Expected Helper to be called")
	}
	if !strings.Contains(fake.fatalMessage, "expected") || !strings.Contains(fake.fatalMessage, "but got") {
		t.Errorf("Expected error message with 'expected' and 'but got', got: %s", fake.fatalMessage)
	}
}

// TestAssertEqualFailureWithMessage tests assertEqual failure with custom message
func TestAssertEqualFailureWithMessage(t *testing.T) {
	fake := &fakeT{TB: t}
	assertEqual(fake, "foo", "bar", "values differ")

	if !fake.fatalCalled {
		t.Error("Expected Fatal to be called")
	}
	if !fake.helperCalled {
		t.Error("Expected Helper to be called")
	}
	// The format string will be "%s: expected %v but got %v"
	if !strings.Contains(fake.fatalMessage, "%s") {
		t.Errorf("Expected format string to contain '%%s', got: %s", fake.fatalMessage)
	}
	if !strings.Contains(fake.fatalMessage, "expected") || !strings.Contains(fake.fatalMessage, "but got") {
		t.Errorf("Expected error message with 'expected' and 'but got', got: %s", fake.fatalMessage)
	}
}
