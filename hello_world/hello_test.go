package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("")
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreeting(t *testing.T) {
	got := Hello("John Tester")
	want := "Hello, John Tester"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
