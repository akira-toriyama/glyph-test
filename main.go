// glyph-test is the fleet's live-fire harness. This program exists so the repo carries a
// go.mod with one real dependency and a flake.nix with a buildGoModule
// vendorHash: the live ammunition for fleet-repin (akira-toriyama/.github),
// whose dependabot gomod bump of golang.org/x/term is the pull it re-pins.
package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	fmt.Println("glyph-test: stdout is a terminal:", term.IsTerminal(int(os.Stdout.Fd())))
}
