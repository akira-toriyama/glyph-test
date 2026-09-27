# automerge probe (projects e-xj9m)

Scratch file the auto-merge probe pulls touch. One line per pull, appended by the pull itself.

- token arm: pull whose auto-merge was armed by GITHUB_TOKEN (automerge-probe.yml)
- user arm: pull whose auto-merge was armed by the user through `gh pr merge --auto` (the FLEET_SYNC_PAT actor)
- no-auto-merge arm: pull armed while the repository had allow_auto_merge=false, to read the API's refusal
