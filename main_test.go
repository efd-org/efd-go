package main

import "testing"

func TestGetGreeting(t *testing.T) {
	if got := getGreeting(); got != "Welcome to EFD Go project!" {
		t.Fatalf("getGreeting() = %q, want %q", got, "Welcome to EFD Go project!")
	}
}
