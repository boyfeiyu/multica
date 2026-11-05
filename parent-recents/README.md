# Parent issue recent suggestions: synthetic browser evidence

For BOYF-466 / upstream #9984, product commit `476602ea01b5016fb0d6a62311127aa4a5ffc860`, base `3c5c8b245359b9b8e913a47f97211f2a3426f149`.

These eight original PNGs were captured by the developer using actual shared production components, ApiClient, styles and editor in Chromium 148. The independent reviewer inspected all eight and reran the same browser fixture against base and head. All HTTP responses are synthetic local test data. This is not native Electron, production permissions or backend E2E evidence. The synthetic-data label is preserved in every image.

The four before/after images show both existing parent pickers with an empty query. Additional images show keyboard selection, create-parent selection, the pending create request, and a successful own-created issue appearing first after a new detail read. `head-create-submitted.png` captures the pending state; it is not by itself proof of successful persistence.

The product PR contains eight source/test/export files only. This separate evidence commit contains exactly these PNGs and this README relative to base; it is not another contribution PR.

## SHA256 of original PNG bytes

`9e1e5052d585a5a16a60ae37bd974338c7326677f3256d9fc77ca0391f42f2e2`  base-create-parent.png  
`3dc04e12502e00f621b7b1baaca69ec4dda9f4aa57cde96651ff54d45d8ca8c3`  base-set-parent.png  
`5ef8ce1f07d6721892cfb7a25fa88f2ab74f2f7b863c15ad43d3d581b0835155`  head-create-parent-selected.png  
`8d85825ce4482adffd1c60cf97888e3160c77bb681ff7fd49d1a5811c7beb534`  head-create-parent.png  
`b63843c38ec6aadb3a65ab119c77b9e90e469042aca20f680cbfeb647ae621c5`  head-create-submitted.png  
`e3c519907af3da31e81a636ba3c630e90e2b694f866cfa8d86ebc6453e391a2b`  head-own-created-recent.png  
`6c3cb9824e8177e6f9b8b7faf30ca9d4fa1fc8968212173ebf9abe7b0e6c6a4a`  head-set-parent-selected.png  
`59aa9729571f840127321c1a3d99c6d18f658b9ccd02f66a3b552d1dd19fb8a4`  head-set-parent.png  
