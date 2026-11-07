# Title baseline interaction evidence

These unmodified Chromium screenshots show the real production IssueDetail and TitleEditor/Tiptap running with normal data hooks and ApiClient against a synthetic loopback HTTP fixture. The banner identifies the synthetic data. They contain no real user issue, workspace data, credentials, or logs. This is component/browser evidence, not full application/authentication/native E2E or real WebSocket delivery.

- before.png: original main e437b11a718d6afce9cd800ba5b8554f3691a133. A clean editor adopted Remote title, but the next edit falsely conflicts because its title_base is still Original title.
- after.png: reviewed product 38b18320bdbf0570209e032753ecdc9d102de1dc. The identical interaction succeeds using the adopted Remote title baseline.
- aba-before.png: prior implementation 09eb3b6977eea732f5e6bdc9fd001381c8f76fee, captured by the developer's repair reproduction. A delayed own response overwrites the baseline after newer text returns A→B→A, causing a false conflict.
- aba-after.png: reviewed product 38b18320bdbf0570209e032753ecdc9d102de1dc, independently captured during review. The delayed response no longer replaces the newer adopted baseline.

The ordinary product PR contains only the product fix and integration tests. This separate evidence branch is based on the original main and contributes no product changes.

SHA256 of original images:

- before.png: `496d1c7b58df031509bce2b4a3f08b8f4ad8318f9b7e1cbba017eaced0614599`
- after.png: `c01f6fbe6056876ed32368c1fa2640149a76c88f7c03344ed2d482e0ee2824fb`
- aba-before.png: `db03ac989053638232c75a2e4998dd7046fae01352ea9e5c3cc82de9dc2fdb64`
- aba-after.png: `d8aa56713756d2897bd62f20e4ddad7992dc22dcb29f929feeab26cf79a87ae4`
