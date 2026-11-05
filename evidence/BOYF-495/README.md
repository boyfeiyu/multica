# TimeInput modifier shortcut behavior

These two unmodified PNGs show the actual production TimeInput in a minimal, CSS-free controlled React harness. All data is synthetic (09:30). They demonstrate keyboard behavior, not the full application layout or styling.

- Before: base `98846ce564a6a7a0d26cb4f071263700003d3041`, Command+1 changes 09:30 to 01:30 and emits onChange.
- After: reviewed product commit `b0937ec489571f7e24a95cf8f4563ca66a94f945`, the same input leaves 09:30 unchanged and emits no onChange.
- Captured with Playwright 1.58.2 / Chromium 148.0.7778.96 on macOS. Independent review replayed the same production component and observed identical behavior for Control+1, Meta+1 and Alt+1, with defaultPrevented false after the fix and no page errors.
- This is browser automation in a synthetic component harness. It does not establish OS shortcut integration, Electron, Safari, native AltGraph/IME, backend permissions or full-application E2E coverage.

| Original image | SHA256 | Bytes |
| --- | --- | --- |
| base.png | e497a3089c099def18cca700cf3a0ee17d5d247b815dca3473d6be283a51205d | 26198 |
| head.png | 70f7ba127611a42389a6080b2ba16a66e04133662b6486c01c43786c7729cf60 | 25875 |

The files contain no user account, production schedule or credentials. This evidence branch is separate from the product change; no evidence files are included in the product PR diff.
