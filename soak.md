# soak

A page that exists to be added in one release and taken away in the next.

The point is not the page. It is that `git revert` writes a subject no gitmoji
pattern claims — `Revert "..."` — and that glyph-test's `glyph.toml` gives that
shape a fixed `~` rather than refusing the range. A revert is a change, so it
gets a version; it does not cancel the release that shipped what it undoes,
because working out which commit a revert answers is exactly the kind of guess
glyph declines to make.
