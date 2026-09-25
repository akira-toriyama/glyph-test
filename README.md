# glyph-test

Permanent live-fire harness for [glyph](https://github.com/akira-toriyama/glyph) —
the sigil-driven commit-lint / semver / release-notes engine.

This repository exists to watch the whole loop run for real:

1. **PR-time lint** — `commit-lint.yml` calls glyph's reusable `lint.yml`;
   a malformed commit fails the check with a SHA-named annotation.
2. **Squash-safe release** — every push to `main` runs `glyph release`,
   which resolves each squash commit back to its merged PR and classifies
   the PR's individual (pre-squash) commits, then upserts ONE rolling
   DRAFT release (tag = next version, body = grouped notes).
3. **Human publish** — no git tag exists until the draft is published in
   the UI (or via `gh release edit --draft=false`); after publishing, the
   next merge rolls a fresh draft above the published floor.

Real pull requests, releases and tags here are fair game. History and tags
become frozen coordinates once an end-to-end workflow (`e2e-v2.yml`,
`livefire.yml`) references them, and are never rewritten.

The "app" here is [haiku.md](haiku.md). No binaries are built — the
release workflow runs glyph's verdict + draft path only.

## How the collection grows

One PR per season. The `:sparkles:` commits drive minor bumps; typo fixes
(`:pencil2:`) drive patches; docs like this line drive nothing.
