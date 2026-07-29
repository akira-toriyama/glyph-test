// Package probe is TEMPORARY live ammunition for the fleet lint-floor
// rehearsal (t-5xyp): it carries one G204 pattern and one G306 pattern so the
// go-ci `--enable` floor can be observed firing — and being scoped by repo
// settings — on real CI. Delete with the probe workflow after the rehearsal.
package probe

import (
	"os"
	"os/exec"
)

// RunArgv executes argv (argv[0] is the program, the rest its arguments) —
// the same deliberate G204 pattern rundiff's runner carries.
func RunArgv(argv []string) error {
	return exec.Command(argv[0], argv[1:]...).Run()
}

// WriteNote writes content to path with 0644 permissions — a deliberate G306
// pattern (G306 is OUTSIDE the G204/G304 scope the fleet convention sets, so
// it distinguishes "gosec ran fully" from "repo scoping applied").
func WriteNote(path string, content []byte) error {
	return os.WriteFile(path, content, 0o644)
}
