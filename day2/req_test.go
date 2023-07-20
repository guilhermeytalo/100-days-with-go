package main

import "testing"

func TestReq(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := HelloReq("Chris", "")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Say 'Hello, Wolrd' when an empty string is supplied",
		func(t *testing.T) {
			got := HelloReq("", "")
			want := "Hello, World"

			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})

	t.Run("in Spanish", func(t *testing.T) {
		got := HelloReq("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
