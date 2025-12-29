package gu

// copied from blog: https://www.alexedwards.net/blog/the-9-go-test-assertions-i-use

// // Equal checks if two values are equal and reports an error if they are not.
// func Equal[T any](t *testing.T, got, want T) {
// 	t.Helper()
// 	if !isEqual(got, want) {
// 		t.Errorf("got: %v; want: %v", got, want)
// 	}
// }

// // NotEqual checks if two values are not equal and reports an error if they are equal.
// func NotEqual[T any](t *testing.T, got, want T) {
// 	t.Helper()
// 	if isEqual(got, want) {
// 		t.Errorf("got: %v; expected values to be different", got)
// 	}
// }

// // True checks if a boolean value is true and reports an error if it is false.
// func True(t *testing.T, got bool) {
// 	t.Helper()
// 	if !got {
// 		t.Errorf("got: false; want: true")
// 	}
// }

// // False checks if a boolean value is false and reports an error if it is true.
// func False(t *testing.T, got bool) {
// 	t.Helper()
// 	if got {
// 		t.Errorf("got: true; want: false")
// 	}
// }

// // Nil checks if a value is nil and reports an error if it is not.
// func Nil(t *testing.T, got any) {
// 	t.Helper()
// 	if !isNil(got) {
// 		t.Errorf("got: %v; want: nil", got)
// 	}
// }

// // NotNil checks if a value is not nil and reports an error if it is nil.
// func NotNil(t *testing.T, got any) {
// 	t.Helper()
// 	if isNil(got) {
// 		t.Errorf("got: nil; want: non-nil")
// 	}
// // }

// // ErrorIs checks if an error is of a specific type and reports an error if it is not.
// func ErrorIs(t *testing.T, got, want error) {
// 	t.Helper()
// 	if !errors.Is(got, want) {
// 		t.Errorf("got: %v; want: %v", got, want)
// 	}
// }

// // ErrorAs checks if an error can be assigned to a target type and reports an error if it cannot.
// func ErrorAs(t *testing.T, got error, target any) {
// 	t.Helper()
// 	if got == nil {
// 		t.Errorf("got: nil; want assignable to: %T", target)
// 		return
// 	}
// 	if !errors.As(got, target) {
// 		t.Errorf("got: %v; want assignable to: %T", got, target)
// 	}
// }

// // MatchesRegexp checks if a string matches a regular expression pattern and reports an error if it does not.
// func MatchesRegexp(t *testing.T, got, pattern string) {
// 	t.Helper()
// 	matched, err := regexp.MatchString(pattern, got)
// 	if err != nil {
// 		t.Fatalf("unable to parse regexp pattern %s: %s", pattern, err.Error())
// 		return
// 	}
// 	if !matched {
// 		t.Errorf("got: %q; want to match %q", got, pattern)
// 	}
// }

// func isEqual[T any](got, want T) bool {
// 	if isNil(got) && isNil(want) {
// 		return true
// 	}
// 	if equalable, ok := any(got).(interface{ Equal(T) bool }); ok {
// 		return equalable.Equal(want)
// 	}
// 	return reflect.DeepEqual(got, want)
// }

// func isNil(v any) bool {
// 	if v == nil {
// 		return true
// 	}
// 	rv := reflect.ValueOf(v)
// 	switch rv.Kind() {
// 	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
// 		return rv.IsNil()
// 	}
// 	return false
// }
