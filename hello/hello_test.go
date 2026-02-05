package main

import "testing"

// Requirements for tests
// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T

// To print docs offline for some package :
// $ go doc fmt

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// To tell the test suite that this is a helper method so that the error line reported as point of failure
	// is where func is called & not in the helper
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
