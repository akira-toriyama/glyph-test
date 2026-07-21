package greet

import "testing"

func TestGreetNamesTheSubject(t *testing.T) {
	if got := Greet("ada"); got != "hi ada" {
		t.Fatalf("Greet(\"ada\") = %q, want %q", got, "hi ada")
	}
}
