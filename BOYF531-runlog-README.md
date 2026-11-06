# Run-only log link screenshots

Product revision: `2c49dae22f30777fbc1b9db3fef64df3e654e1e3`.

These six original, unedited PNGs were captured and visually inspected during independent review. All workspace names, run IDs, payloads and transcript text are synthetic. They exercise production shared page/transcript components, ApiClient and web/desktop navigation adapters in Chromium. HTTP, Next router hooks, workspace directory data and desktop preload are fixtures. This is not a full backend, Next server or native Electron session.

The terminal pair shows a cold link opening a completed run outside the latest list, including the existing read-only payload display. The paused pair shows an invalidated successful cache waiting offline without stale events or a false zero-step result. The retry pair shows a 503 manual read retry with immediate Loading feedback and a disabled action.

Before this change, the upstream UI could open a run transcript only from its row; a cold URL could not select a run. The original request and before screenshots are in https://github.com/multica-ai/multica/issues/10017 .

The evidence commit contains only these six PNGs and this README on top of base `61e7e8f966f53478f3efe68ec50c5a5fdd00a51f`; it is separate from the product branch.

| File | Bytes | SHA256 |
| --- | ---: | --- |
| BOYF531-web-terminal.png | 62792 | `04acaec2ba44d33d7785a135f8780717302e80b57eb9d1af9868663fbf0de58d` |
| BOYF531-desktop-terminal.png | 63630 | `fceeeb119f7589e2aa4eefbfcfb54f2f55dff16e95e1913bd6bc5a44e03e608e` |
| BOYF531-web-paused.png | 52263 | `8fc0581d52225096537e13f64b7dc133e12289c4f918ea1acdf58bd9949f1d13` |
| BOYF531-desktop-paused.png | 52620 | `2de54718cf929476c7496581f03a66f6645376b7c3fc9643e8151966fb2c1e64` |
| BOYF531-web-retry.png | 140075 | `e2636f6533ebb8d2ad7c0e3ef2b6827c910830dfdcf2dd319ca1f8ba76e722c0` |
| BOYF531-desktop-retry.png | 140463 | `15969eec9107c06060f6bf47824898fba5d66462407aa16f38df77b00611a4c3` |
