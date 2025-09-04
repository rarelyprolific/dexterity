package main

import (
	"testing"
)

func TestFormatError(t *testing.T) {
	got := formatError("dummy invalid value", "dummy description")
	want := "ERROR! The value [dummy invalid value] is not a valid dummy description!"

	if got != want {
		t.Errorf("formatError(): got %q, want %q", got, want)
	}
}
