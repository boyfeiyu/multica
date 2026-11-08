# Run log links after upstream integration

Product revision: cc76fcae8940e2a8dbd48f013c08009131e72917.
Integrated upstream main: 57fa7580469a9b3dae5ebe99629107b7dd6438e7.

These original PNGs were captured and visually inspected during independent
BOYF-646 review using the current checkout's compiled CSS and actual registered
web and desktop renderer pages. All displayed data is synthetic. The writer
has the upstream Stop run control next to the log and copy links; the reader
can open and copy logs with no Stop control. The cold terminal run is outside
the latest 20 rows. Missing and denied runs share the opaque unavailable state.

HTTP was limited to exact GET paths on an ephemeral foreground loopback server.
Next router history hooks, desktop useParams and preload runtimeConfig, and
workspace directory data were fixture boundaries. Actual shared components,
API client, query hooks, editor/transcript and navigation adapters ran. This
is not a full Next server, backend auth, WebSocket, native Electron or Expo E2E.
No real provider, task, daemon, database or user profile was used. Stop was not
clicked against a backend; its request is covered by a synthetic API unit test.

The product passed 472 focused unit/component tests and 124 browser cases.
Eight repeated screenshot controls are additional captures, not unique cases.
The original historical screenshot commit remains unchanged:
734bf461d1284752eea078fd0bb24fa7184c1524.

## Original PNG hashes

- web-stop-log-writer.png: SHA256 dc6a569160acffc57b32fde99000ba9a61c90c0880a17472d097c0aed49533f6
- desktop-stop-log-reader.png: SHA256 d98987855cfaefef6c0156b5b6ea6bdd14489001e6b0bd88409a73ae3ae141f5
- web-cold-terminal.png: SHA256 594dce8e38713addb337ad5d065f96a98c0385cb70b02a61d3b6b144d679aa92
- desktop-opaque-404.png: SHA256 ff6ab8e5814bdbb83b8aadfe6b9387af3825cbedb5246e5a222ac01b48ef1734
