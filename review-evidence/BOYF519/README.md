# Clipboard selection regression evidence

Original synthetic-data Chromium screenshots from the BOYF-520 development fixture, inspected and hash-verified independently in BOYF-521. This branch carries review material only; the product change remains commit 44844845ea8b94c8e1b039f081f04ba063819563.

Base: 61e7e8f966f53478f3efe68ec50c5a5fdd00a51f. Head: 44844845ea8b94c8e1b039f081f04ba063819563.

The fixture mounts the actual React CodeBlockView through a real Tiptap editor and ReactNodeViewRenderer, using the original code.css plus minimal layout. Unused Mermaid/HTML preview and translation dependencies use documented fixtures. The code sample and form values are synthetic. Copy buttons and their handlers are production components. Clipboard/execCommand are test doubles; browser selection, focus and keyboard input are real.

`*-after-copy.png` shows the same forward selection after clicking Copy code. Base loses it; head preserves it. `*-after-typing.png` shows focusing the editor and typing Z: head replaces the selected range with Z (`syZitable text`), while the base fixture does not replace that range. These are isolated component fixtures, not screenshots of a complete product route, Electron, native OS clipboard, or permission/backend E2E.

The product source is distributed under the Multica License included in the repository. This material branch has no product changes or separate pull request.

## Original PNG SHA256

- `base-after-copy.png`: `870f07de7dcebcce2c0c25c31879ec9256752a6e83b81046ce71f4108d92cfb2` (15979 bytes)
- `base-after-typing.png`: `35b123b3b9f31f0dba7b2faefc2ef7d7b9a7cdfb3aa32f8dc5afa7810bf3425e` (16059 bytes)
- `head-after-copy.png`: `8e1978dde39c26fe9679ffe24f1fa810bfcff213c5bfc625a12303f7239948e1` (16329 bytes)
- `head-after-typing.png`: `265b7f4ac31908ed9ce1d1f6f44012f90832a1801dd89d22261eef12c1b1605c` (15132 bytes)
