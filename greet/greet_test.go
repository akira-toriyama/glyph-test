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

// TestGreetKeepsInnerWhitespace records what Greet does with whitespace INSIDE a
// name today: nothing. Whether it should collapse it is undecided, so this is a
// record of the current answer rather than an assertion that the answer is right.
//
// bite-exempt: pins today's answer on inner whitespace, which no change here fixes
func TestGreetKeepsInnerWhitespace(t *testing.T) {
	if got := Greet("ada  lovelace"); got != "hi ada  lovelace" {
		t.Fatalf("Greet(%q) = %q, want the inner whitespace kept", "ada  lovelace", got)
	}
}
