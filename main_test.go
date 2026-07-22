package main

import "testing"

func TestGetGreeting(t *testing.T) {
	if got := getGreeting(); got != "Hello, World!" {
		t.Fatalf("getGreeting() = %q, want %q", got, "Hello, World!")
	}
}
