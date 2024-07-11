package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying Hello to people", func(t *testing.T) {
		got := Hello("John")
		want := "Hello, John"

		if got != want {
			t.Errorf("got %q, wanted %q ", got, want)
		}
	})

	t.Run("saying 'Hello world' when empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q, wanted %q ", got, want)
		}
	})
}
