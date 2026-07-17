# glyph-test

Throwaway end-to-end demo of [glyph](https://github.com/akira-toriyama/glyph) —
the gitmoji-driven commit-lint / semver / release-notes engine.

This repo exists only to watch the whole loop run for real:

1. **PR-time lint** — `commit-lint.yml` calls glyph's reusable `lint.yml`;
   a malformed commit fails the check with a SHA-named annotation.
2. **Squash-safe release** — every push to `main` runs `glyph release`,
   which resolves each squash commit back to its merged PR and classifies
   the PR's individual (pre-squash) commits, then upserts ONE rolling
   DRAFT release (tag = next version, body = grouped notes).
3. **Human publish** — no git tag exists until the draft is published in
   the UI (or via `gh release edit --draft=false`); after publishing, the
   next merge rolls a fresh draft above the published floor.

The "app" here is [haiku.md](haiku.md). No binaries are built — the
release workflow runs glyph's verdict + draft path only.

---（和訳）

glyph（gitmoji 駆動の commit-lint / semver / release-notes エンジン）の
端から端までを実際に動かして見るための使い捨てデモ repo。PR-time lint →
squash-safe な rolling draft release → 人間の Publish、の一巡を回す。
