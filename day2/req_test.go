package main

import "testing"

func TetReq(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Say 'Hello, Wolrd' when an empty string is supplied",
		func(t *testing.T) {
			got := Hello("")
			want := "Hello , World"

			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hello, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
