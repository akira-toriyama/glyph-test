// Package greet is the smallest possible thing a test can be wrong about.
//
// glyph-test is where the fleet's CI machinery gets fired with live ammunition
// instead of being reasoned about. The go-bite gate needs a Go package whose
// behaviour can be broken and fixed on demand; this is it. Nothing ships from
// here.
package greet

import "strings"

// Greet addresses someone by name, ignoring whatever whitespace the caller
// happened to pass.
func Greet(name string) string {
	return "hi " + strings.TrimSpace(name)
}

// Shout greets in capitals. It is the shipping change this pull request makes;
// the test beside it deliberately pins something else.
func Shout(name string) string {
	return strings.ToUpper(Greet(name))
}
