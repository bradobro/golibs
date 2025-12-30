package gu

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

// testingT is an interface that matches the methods used by our test helpers
type testingT interface {
	Helper()
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

// True checks if a condition is true and reports an error if it is false.
// It includes an optional hint for additional context.
func True(t testingT, got bool, hint string) {
	t.Helper()
	assertTrue(t, got, hinted("got false; want true", hint))
}

// False checks if a condition is false and reports an error if it is true.
// It includes an optional hint for additional context.
func False(t testingT, got bool, hint string) {
	t.Helper()
	assertTrue(t, !got, hinted("got true; want false", hint))
}

// Equal checks if two values are equal and reports an error if they are not.
// It includes an optional hint for additional context.
func Equal[T any](t testingT, got T, want T, hint string) {
	// This simpler Equal works most of the time, but doesn't handle all edge cases
	// func Equal[T comparable](t testingT, got T, want T, hint string) {
	// 	t.Helper()
	// 	assertTrue(t, got == want, hinted(fmt.Sprintf("got %v; want %v", got, want), hint))
	// }
	t.Helper()
	msg := hinted(fmt.Sprintf("got %v; want %v", got, want), hint)
	assertTrue(t, isEqual(got, want), msg)
}

// NotEqual checks if two values are not equal and reports an error if they are equal.
// It includes an optional hint for additional context.
func NotEqual[T any](t testingT, got T, want T, hint string) {
	t.Helper()
	msg := hinted(fmt.Sprintf("got %v; want something different", got), hint)
	assertTrue(t, !isEqual(got, want), msg)
}

// Nil checks if a value is nil and reports an error if it is not.
// It includes an optional hint for additional context.
func Nil(t testingT, got any, hint string) {
	t.Helper()
	msg := hinted(fmt.Sprintf("got %v; want nil", got), hint)
	assertTrue(t, !isNil(got), msg)
}

// NotNil checks if a value is not nil and reports an error if it is nil.
// It includes an optional hint for additional context.
func NotNil(t testingT, got any, hint string) {
	t.Helper()
	msg := hinted("got nil; want non-nil", hint)
	assertTrue(t, isNil(got), msg)
}

// ErrorIs checks if an error is of a specific type and reports an error if it is not.
// It includes an optional hint for additional context.
func ErrorIs(t testingT, got, want error, hint string) {
	t.Helper()
	msg := hinted(fmt.Sprintf("got %v; want something different", got), hint)
	assertTrue(t, errors.Is(got, want), msg)
}

// ErrorAs checks if an error can be assigned to a target type and reports an error if it cannot.
// It includes an optional hint for additional context.
func ErrorAs(t testingT, got error, target any, hint string) {
	t.Helper()
	msg := hinted(fmt.Sprintf("got %v; want assignable to: %T", got, target), hint)
	assertTrue(t, got != nil, msg)
	if got == nil {
		t.Fatal(msg)
	} else {
		assertTrue(t, errors.As(got, target), msg)
	}
}

// MatchesRegexp checks if a string matches a regular expression pattern and reports an error if it does not.
// It includes an optional hint for additional context.
func MatchesRegexp(t testingT, got, pattern, hint string) {
	t.Helper()
	msg := fmt.Sprintf("got %v; doesn't match %q", got, pattern)
	matched, err := regexp.MatchString(pattern, got)
	if err != nil {
		t.Fatalf("unable to parse regexp pattern %s: %s", pattern, err.Error())
	} else {
		assertTrue(t, matched, msg)
	}
}

// isEqual compares two values for equality, handling nil values and custom equality methods.
func isEqual[T any](got, want T) bool {
	if isNil(got) && isNil(want) {
		return true
	}
	if equalable, ok := any(got).(interface{ Equal(T) bool }); ok {
		return equalable.Equal(want)
	}
	return reflect.DeepEqual(got, want)
}

// isNil checks if a value is nil, considering various types like pointers, slices, and maps.
func isNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return rv.IsNil()
	}
	return false
}

// hinted adds additional information (a "hint") to the failure message
// if provided
func hinted(message, hint string) string {
	if hint == "" {
		return message
	}
	return hint + ": " + message
}

// assertTrue checks if a condition is true, failing the test if not
func assertTrue(t testingT, got bool, message string) {
	t.Helper()
	if !got {
		t.Fatal(message)
	}
}
