package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessge := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jun", "french")
		want := frHelloPrefix + "Jun"

		assertCorrectMessge(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessge(t, got, want)
	})
}
