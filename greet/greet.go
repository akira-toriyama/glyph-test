// Package greet is the smallest possible thing a test can be wrong about.
//
// glyph-test is where the fleet's CI machinery gets fired with live ammunition
// instead of being reasoned about. The go-bite gate needs a Go package whose
// behaviour can be broken and fixed on demand; this is it. Nothing ships from
// here.
package greet

// Greet addresses someone by name.
func Greet(name string) string {
	return "hi " + name
}
