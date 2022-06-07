package main

import "testing"

func TestGreet(t *testing.T) {
	assertGreeting := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Greet("Chris", "")
		want := "Hello, Chris!"

		assertGreeting(t, got, want)
	})
	t.Run("say 'Hello, World! to strangers", func(t *testing.T) {
		got := Greet("", "")
		want := "Hello, World!"

		assertGreeting(t, got, want)
	})
	t.Run("say hello in spanish", func(t *testing.T) {
		got := Greet("Elodie", "Spanish")
		want := "Hola, Elodie!"

		assertGreeting(t, got, want)
	})
}
