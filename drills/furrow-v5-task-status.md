# furrow v5 task-status live fire

2026-08-31: canary drill for the furrow v4.0.0 -> v5.0.0 rollout (hub rollout ledger, t-satm).
This PR merge must drive sync-task-status.yml@v5.0.0 (delivered by the canary fleet-sync
apply, run 33398944811) and write the shared board via the SetStatus footer.
