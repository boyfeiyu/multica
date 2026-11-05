# Sub-issues report: synthetic UI evidence

These four unmodified Chromium screenshots were captured during BOYF-462's independent review at product commit `c11feffa952084ff3a2fe04cd9c4a3d7daa69610`, compared against base `3c5c8b245359b9b8e913a47f97211f2a3426f149`. They use **Synthetic local test data** and the production shared Web/Desktop components, API client and zod schema with synthetic local HTTP responses. They are not native Electron or production permission end-to-end tests. No real customer data or real model execution is involved.

The follow-up product commit `e5c4271f2aa1b0e33fa0ea27b7066965d55a9a9f` changes only seven lines of Korean/Traditional Chinese copy in three files. All runtime code, tests and the English UI shown here are byte-identical to the independently exercised version. These are reused screenshots, not a claim that the browser scenario was rerun on the follow-up commit. The separate PostgreSQL scheduler/API validation is not supplied by this browser fixture.

The product PR adds an explicit opt-in wait for all selected direct sub-issues to reach exact `in_review` or a terminal `done`/`closed` category. Cancellation ends the wait without implying success or review approval. Default terminal-only waits and system stage barriers stay terminal-only. The new condition does not approve work, assign agents or advance stages, and retains 30-second polling/count-fingerprint limitations.

| Image | Meaning | SHA256 |
| --- | --- | --- |
| [base-menu.png](base-menu.png) | Before: condition menu at base | `673afd661a07ecc0520464dc8d466b20e4ec6519381b50507fa0bec0112d8e22` |
| [head-menu.png](head-menu.png) | After: new opt-in menu choice | `82d08c459ac55f5c7d30b7fb6f45f79760bfb07ffaae0b407d754fd900a0a66a` |
| [head-stage-form.png](head-stage-form.png) | After: stage 2 wait form | `d2bcf7a4d459e04dcb5fab80f59ed9ba366b9d3c4847ef661379067620d3c372` |
| [head-saved-rule.png](head-saved-rule.png) | After: saved rule read back through the API client | `affe6d1301e253ee55deab32ee125220616060352559adabaa576ed6c88ed961` |
