## Desktop closed-tab recovery review media

These original PNGs were captured by the independent reviewer with Playwright Chromium at 1152 × 640. The workspace and issue are synthetic. The tab strip, Base UI context menu, tab store, navigation-input hook and CSS use production source. Resource icons/titles and the destination content frame are fixture presentation. The screenshots cover the renderer interaction.

The base uses `61e7e8f966f53478f3efe68ec50c5a5fdd00a51f`. The reviewed product uses `3aa2b88a69d17ad1f672ec5dd83a787e9c2ef74f`. For the base capture, all changed production files were replaced with exact base blobs and restored afterward. The product checkout remained at the reviewed tree after verification.

- `base-menu.png` shows the actual base menu after closing the synthetic issue tab.
- `head-menu.png` shows the recovery action under the same conditions at the reviewed head.
- `head-reopened.png` shows the recovered destination at its original strip position with full query and hash.

This evidence branch contains only these images and this README on the exact base. Its commit is separate from the reviewed product commit. Browser and jsdom checks use fixture API/preload/storage. Physical OS shortcuts, full Electron/backend integration and server permission enforcement were outside this verification.

| File | Bytes | SHA-256 |
| --- | --- | --- |
| base-menu.png | 26382 | `ccec3c48fd814b2ef31d568c2836e65eae86bc3178f3ef45c607d277ff710d2a` |
| head-menu.png | 30302 | `7ff2b9997824416d0e7322c157ce2fbc2b7bb66f0fb0ed563e1ffc41b5f7c9ef` |
| head-reopened.png | 29456 | `80d7174150c915e17d005298013fc77eb240ba1ff2e404b67e1a4cec8dd2e74e` |
