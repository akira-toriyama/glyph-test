package greet

import "testing"

func TestGreetNamesTheSubject(t *testing.T) {
	if got := Greet("ada"); got != "hi ada" {
		t.Fatalf("Greet(\"ada\") = %q, want %q", got, "hi ada")
	}
}

func TestGreetTrimsSurroundingWhitespace(t *testing.T) {
	if got := Greet("  ada\t"); got != "hi ada" {
		t.Fatalf("Greet(%q) = %q, want %q", "  ada\t", got, "hi ada")
	}
}
