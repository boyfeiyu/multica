# Cold list first-page retry reproduction

These original, unedited screenshots use synthetic workspace, agent, and issue data in a loopback browser fixture. The real AgentWorkDialog, ActorIssuesPanel, IssueSurface, production query hooks, and ApiClient are mounted.

The baseline is `7c75b34a23fdb636c67c39e60d158e6eeeec6d3b`. The reviewed fix is `981ec5b8fdbfccb3bedb1671c0bca6609105caa2`.

- `base-failure.png` shows an authoritative Todo count of 1 while the first row request returns HTTP 503. The section displays “No issues”.
- `head-failure.png` shows the same failure with “Could not load. Retry”.
- `head-recovered.png` shows the returned issue after clicking Retry on the current query.
- `head-normal.png` shows the successful cold-load control.

The fixture allows read-only table rows/facets POST queries and necessary synthetic GET responses. It rejects other mutations and non-loopback traffic. QueryClient uses `retry: false` for a deterministic terminal-error observation. Production retry policy is unchanged. This evidence covers shared UI behavior in a browser fixture. Physical Desktop, native OS integration, backend authentication, and full Electron execution were outside the checks.
