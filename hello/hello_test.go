package main

import "testing"

// Requirements for tests
// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T

// To print docs offline for some package :
// $ go doc fmt

func TestHello(t *testing.T) {
	got := Hello("John")
	want := "Hello, John"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
