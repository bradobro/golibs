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

// TestAssertionsTableDriven tests all assertion functions using a table-driven approach.
func TestAssertionsTableDriven(t *testing.T) {
	tests := []struct {
		testName   string
		assertFunc func(t testingT, args ...interface{})
		args       []interface{}
		expectFail bool
		expectMsg  string
	}{
		// True tests
		{
			testName: "True passes",
			assertFunc: func(t testingT, args ...interface{}) {
				True(t, args[0].(bool), args[1].(string))
			},
			args:       []interface{}{true, ""},
			expectFail: false,
		},
		{
			testName: "True fails",
			assertFunc: func(t testingT, args ...interface{}) {
				True(t, args[0].(bool), args[1].(string))
			},
			args:       []interface{}{false, "custom message"},
			expectFail: true,
			expectMsg:  "custom message: got false; want true",
		},
		// False tests
		{
			testName: "False passes",
			assertFunc: func(t testingT, args ...interface{}) {
				False(t, args[0].(bool), args[1].(string))
			},
			args:       []interface{}{false, ""},
			expectFail: false,
		},
		{
			testName: "False fails",
			assertFunc: func(t testingT, args ...interface{}) {
				False(t, args[0].(bool), args[1].(string))
			},
			args:       []interface{}{true, "custom message"},
			expectFail: true,
			expectMsg:  "custom message: got true; want false",
		},
		// Equal tests
		{
			testName: "Equal passes",
			assertFunc: func(t testingT, args ...interface{}) {
				Equal(t, args[0], args[1], args[2].(string))
			},
			args:       []interface{}{42, 42, ""},
			expectFail: false,
		},
		{
			testName: "Equal fails",
			assertFunc: func(t testingT, args ...interface{}) {
				Equal(t, args[0], args[1], args[2].(string))
			},
			args:       []interface{}{42, 43, "custom message"},
			expectFail: true,
			expectMsg:  "custom message: got 42; want 43",
		},
		// NotEqual tests
		{
			testName: "NotEqual passes",
			assertFunc: func(t testingT, args ...interface{}) {
				NotEqual(t, args[0], args[1], args[2].(string))
			},
			args:       []interface{}{42, 43, ""},
			expectFail: false,
		},
		{
			testName: "NotEqual fails",
			assertFunc: func(t testingT, args ...interface{}) {
				NotEqual(t, args[0], args[1], args[2].(string))
			},
			args:       []interface{}{42, 42, "custom message"},
			expectFail: true,
			expectMsg:  "custom message: got 42; want something different",
		},
		// Nil tests
		{
			testName: "Nil passes",
			assertFunc: func(t testingT, args ...interface{}) {
				Nil(t, args[0], args[1].(string))
			},
			args:       []interface{}{nil, ""},
			expectFail: false,
		},
		{
			testName: "Nil fails",
			assertFunc: func(t testingT, args ...interface{}) {
				Nil(t, args[0], args[1].(string))
			},
			args:       []interface{}{42, "custom message"},
			expectFail: true,
			expectMsg:  "custom message: got 42; want nil",
		},
		// NotNil tests
		{
			testName: "NotNil passes",
			assertFunc: func(t testingT, args ...interface{}) {
				NotNil(t, args[0], args[1].(string))
			},
			args:       []interface{}{42, ""},
			expectFail: false,
		},
		{
			testName: "NotNil fails",
			assertFunc: func(t testingT, args ...interface{}) {
				NotNil(t, args[0], args[1].(string))
			},
			args:       []interface{}{nil, "custom message"},
			expectFail: true,
			expectMsg:  "custom message: got nil; want non-nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fake := &fakeT{TB: t}
			tt.assertFunc(fake, tt.args...)

			if tt.expectFail && !fake.fatalCalled {
				t.Errorf("%s: expected Fatal to be called", tt.testName)
			}
			if !tt.expectFail && fake.fatalCalled {
				t.Errorf("%s: did not expect Fatal to be called", tt.testName)
			}
			if tt.expectFail && !strings.Contains(fake.fatalMessage, tt.expectMsg) {
				t.Errorf("%s: expected error message to contain '%s', got: '%s'", tt.testName, tt.expectMsg, fake.fatalMessage)
			}
		})
	}
}
