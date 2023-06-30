package main

import "testing"

func TestHelloMsg(t *testing.T) {
	t.Run("Saying helllo to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMsg(t, got, want)
	})

	t.Run("Empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMsg(t, got, want)
	})
}

func assertCorrectMsg(t testing.TB, got, want, string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}