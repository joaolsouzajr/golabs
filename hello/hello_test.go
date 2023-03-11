package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris"

		assertMessage(t, got, want)
	})

	t.Run("saying hello world when an empty strig is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"

		assertMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola Elodie"

		assertMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Nadir", "French")
		want := "Bonjour Nadir"

		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
