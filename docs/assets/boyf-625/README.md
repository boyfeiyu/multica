# Pending description save evidence

Original, unedited 1440 × 1000 screenshots from an independent Chromium run of the production IssueDetailRoute → Shift+Click issue mention → IssuePeekHost → IssueDetail/ContentEditor/Tiptap components.

The loopback fixture uses synthetic issues, accounts and text. It keeps the normal QueryClient, hooks and ApiClient and substitutes only the HTTP service and app navigation shell. These are client interaction tests, not live backend, permission, native, media persistence or WebSocket E2E tests. The visible BOYF-626 marker identifies the original fixture family.

- `before-d1-only.png`: base 4239ec11eaa67541c28bb20e836768263e363286. D1 is held, D2 passes the normal 1500 ms debounce into the component queue, then the real Peek closes and unmounts. Releasing D1 leaves only D1 when reopened.
- `after-d2-saved.png`: reviewed dcbf3b98d5daf5a00a1711eac9c6a148fc079589. The same interaction sends D2 after D1 and reopens with D2.
- `new-session-b-d3.png`: reviewed dcbf3b98d5daf5a00a1711eac9c6a148fc079589. A logs out while D1 is held; B logs in and queues B D1/D2/D3. A's late completion does not resume A's unsent draft or remove B's queue; B D3 is retained.

This branch contains presentation evidence only. The product commit stays unchanged. Screenshots show synthetic data and have been visually inspected; no logs, archives, tokens or workspace metadata are published.
