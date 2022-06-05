package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Chris")
	want := "Hello, Chris!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
