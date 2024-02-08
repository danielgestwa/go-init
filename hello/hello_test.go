package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Saying hello to Dan", func(t *testing.T) {
		got := Hello("Dan")
		want := "Hello Dan!"
		assertMessage(t, got, want)
	})

	t.Run("Saying hello to empty", func(t *testing.T) {
		got := Hello("")
		want := "Hello world!"
		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Error! %q is not %q", got, want)
	}
}
