# glyph v4.3.0 candidate live fire

2026-10-04: e2e-v2 at glyph `main` is the live-ammunition stage before the v4.3.0 tag
(glyph rollout runbook, step 0). Two states of this sandbox kept it from answering:

- The rolling draft v2.0.1 stayed unpublished after #94 retired the v1-acceptance window, so
  `refusals` arm (a) walked `below:v2.0.0` = v1.0.0 through sigil-less history and wedged at
  exit 3 before it could reach the published floor (run 37203074410; the shipped v4.2.0 binary
  answers the same, so the red was this sandbox's). Publishing v2.0.1 put the walk back on
  history the grammar reads: `below:v2.0.1` answers exit 4 with both binaries.
- With v2.0.1 published and nothing releasable after it, `release --dry-run --since-tag=auto`
  answers none (exit 1), and the arms that read live main want a release (run 37203843011:
  `walk` live-release and `refusals` arm (c)).

This commit is the releasable change past v2.0.1 those arms read.

- live check: the fleet pin at glyph v4.3.0 refuses a sigil-less subject and passes the reworded one (t-s3e6)
