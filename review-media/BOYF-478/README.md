# Inbox within Work: synthetic UI screenshots

These unchanged original PNGs were captured by the developer from Chromium 148.0.7778.96, using production shared navigation, InboxList and IssueDetail with an in-memory synthetic API. They contain synthetic Acme/Beta workspaces, TEST issues and example.invalid fixture identity, not user data.

Base: `f230f9f55bdfc4de68015473b07de329a5079f87`. Reviewed product head: `2c733aad65cd5e4dffc915218ed0d261441b04c6`. This branch only adds review media; it is not the product branch and does not alter the reviewed product head.

Before: `base-wide-initial.png` and `base-wide-selected.png` show Inbox as its own rail destination. After: `head-wide-initial.png` and `head-wide-selected.png` show the Work navigation, Inbox list and detail. Other images show navigation collapse, one Work hover preview, archived selection, the 900px automatic collapse, the 390px compact detail/sheet, and the actual desktop navigation adapter/tab-store fixture. Wide screenshots use 1440px.

The original file named compact-list-after-back.png was omitted: its pixels still show the detail, so it cannot demonstrate the completed Back transition. The independent reviewer separately reran the fixture and verified the completed transition using a hidden-detail and cleared-selection assertion. All included PNGs were byte/hash checked and visually inspected before publication.

Limits: AccountMenu, HelpLauncher and Team status are stubbed; non-Inbox destination content is simplified. No real backend, authentication, permissions, Electron/Expo, live comment scrolling or server-side issue creation was exercised. Project creation stops at the existing modal store. These screenshots are not full-stack E2E evidence. A separate independent run exercised the production components and adapter at the same viewports.

| File | Bytes | SHA256 |
| --- | ---: | --- |
| [base-wide-initial.png](base-wide-initial.png) | 61330 | `1b05d1c327e90770cbcbf428228d9fb933947f3bb49bdd7570787ec0d6f69170` |
| [base-wide-selected.png](base-wide-selected.png) | 79786 | `6e01b51788fc839793c37f674d4e4648b4457a647e1336b703e63de1e112a739` |
| [head-compact-selected.png](head-compact-selected.png) | 32843 | `68e760101d0abdf79f4bbd468138924769cb15b411d43ca2a90ec7e92b6e98d9` |
| [head-compact-work-sheet.png](head-compact-work-sheet.png) | 31611 | `f5601ff8a3d46bcf3622f39d653bbc9af02e21d81e14b905a6d166516a08c705` |
| [head-desktop-adapter-restored-inbox.png](head-desktop-adapter-restored-inbox.png) | 88946 | `3fe8e2d5656d74b92282fd4bf0e1d6dce767a1432294542e512eee40e0150c6b` |
| [head-desktop-rendered-inbox-transition.png](head-desktop-rendered-inbox-transition.png) | 72652 | `b2762c81b9cb29c6113072cdaec24a63f1ee4af694949934d5e7ee51465cd8e0` |
| [head-medium-auto-collapse.png](head-medium-auto-collapse.png) | 56079 | `98c1a83d3faba2534cbc66d7765593172e025f845cbaef8869d7ad5086187192` |
| [head-wide-archived-selected.png](head-wide-archived-selected.png) | 89876 | `f6a624f55122694096fc07103dd407585eb7c088c3a2d7a4a19f4c7a31bdeba2` |
| [head-wide-initial.png](head-wide-initial.png) | 73388 | `36ba973620bdab94ce243d9143dd0e0ea9914c1e6ed9e91c8e8278250b698363` |
| [head-wide-nav-collapsed.png](head-wide-nav-collapsed.png) | 79004 | `07141b7cb5289a9cb7a76a1d61cb29c201075f9297c8ebc35a6e2440ea6a1eb9` |
| [head-wide-selected.png](head-wide-selected.png) | 90335 | `9f03896e4600c19140d2192605b54eb6564f3a07e717ad6e5e4c000ce540162b` |
| [head-wide-work-hover.png](head-wide-work-hover.png) | 60060 | `ffbd4f4e9c4bed3f28aadfb7433f10fb3e7534d0342f9f21ba04d00de42ca7bd` |
