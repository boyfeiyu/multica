Original invitation recovery screenshots from BOYF-672 independent review.

Product BASE: 73975653323750d04c27047ea9c986ef06bdf6a1
Product reviewed HEAD: 8f87856307b9d09f15eae0a5f3e8dc404e7de6ff

All four 1280x900 PNGs are unedited browser captures of running exact-source components, captured and inspected by the independent reviewer. Every person, workspace and invitation is synthetic. No real business service was used.

- base-stale-page.png: successful cached old pending row remains after a remote fixture list change, real31.1s wait and trusted focus/online recovery; one account invitation GET.
- head-recovered-page.png: eligible focus fetch replaces the old row with the new pending row; two GETs total, subsequent fresh focus/online adds no reads.
- head-normal-page.png: normal warm list before the remote change; zero-selection Skip still navigates to onboarding without writes. The old fixture label deliberately says ended; the remote conclusion occurs only later in target probes.
- head-recovered-rail.png: eligible offline-to-online recovery in real shared RailWorkspaces; two GETs. Existing narrow menu text uses ellipsis.

The Web InvitationsRoutePage and shared InvitationsPage run with synthetic auth/navigation providers and normal createQueryClient/QueryClientProvider/query options/ApiClient against an own random-port127.0.0.1 GET-only fixture. RailWorkspaces runs in a synthetic workspace/provider container; normal AppRail registration was statically checked. Product CSS is compiled from the same checkout; the harness uses local font fallbacks. No full Next/Electron/native/auth/WS/backend E2E or browser invitation mutation is claimed.

This independent evidence branch has no product changes and no extra PR. The product PR must point only to the reviewed HEAD above.

SHA256 and Git blob OID inventory:

[
  {
    "filename": "base-stale-page.png",
    "source_run": "base-target-page",
    "sha256": "430e5f10d586966dad8a594a58f0261252ee62b481b8015bc60152da118253ef",
    "git_blob_oid": "a71ed6b106d53ace21b56e855f87414132c28665"
  },
  {
    "filename": "head-recovered-page.png",
    "source_run": "head-focus-page",
    "sha256": "a95fc678e693e8eb35c09de3d4ece3a24bac50aca6608d2913a63c0f13169c09",
    "git_blob_oid": "9126d976012dd48a0cdacf7d347df642a955cb66"
  },
  {
    "filename": "head-normal-page.png",
    "source_run": "head-normal",
    "sha256": "430e5f10d586966dad8a594a58f0261252ee62b481b8015bc60152da118253ef",
    "git_blob_oid": "a71ed6b106d53ace21b56e855f87414132c28665"
  },
  {
    "filename": "head-recovered-rail.png",
    "source_run": "head-online-rail",
    "sha256": "2d87a1e7f7e40748f68f1b44f95667f1ce1b91ef73219a5251c62ea1b29afc92",
    "git_blob_oid": "390cce61d0c926b5e1466d40311937f171557fcb"
  }
]
