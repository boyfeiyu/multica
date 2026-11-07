# Draft label selection interaction evidence

These unmodified 1200 x 900 Chromium screenshots show the production CreateIssueDialog, ManualCreatePanel, LabelPicker and PropertyPicker, with normal providers, Query hooks and ApiClient. Only HTTP responses are synthetic on a temporary loopback server. Alpha, Beta, Gamma and the issue/workspace shown are fake data. No real user data, credentials, raw logs or workspace metadata are included.

Select Alpha, start inline creation of Gamma, then clear the search, remove Alpha and select Beta while the response waits. Complete creation:

- before.png: exact original base e437b11a718d6afce9cd800ba5b8554f3691a133 incorrectly restores Alpha and drops Beta. The UI and real browser local draft both contain Alpha, Gamma.
- after.png: independently reviewed product 5b0dac351ee1730b7d6b0416dff5aa48b109de90 retains Beta, Gamma in the UI and real browser local draft.

This is synthetic browser integration evidence, not live authentication, backend/database persistence, WebSocket/native integration or full application E2E. Unmount, reopen/new draft, account/workspace switching and cancellation behavior are outside the verified scope.

This branch starts at the original base and contains only these two images and this README. The product branch contains the fix and adjacent tests, with no evidence files.

SHA256 of the original PNG bytes:

- before.png: `4a29a31a35b5c5698f440afc617cce7891341cad64e3e34fb7ae93c146771a74`
- after.png: `c49156d5a3a9e6e73b2eef87eab81e4fb303dc0598bbe82676f67377b21f97fa`
