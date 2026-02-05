package main

import "testing"

func TestArraySum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := ArraySum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %v want %v given %v", got, want, numbers)
	}
}
