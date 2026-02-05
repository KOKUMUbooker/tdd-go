package main

import "testing"

func TestAdder(t *testing.T) {
	got := Adder(1,2)
	want := 3

	if got != want {
		t.Errorf("got : %q want : %q",got,want)
	}
}