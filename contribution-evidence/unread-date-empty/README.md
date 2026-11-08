Original screenshots from a running browser integration with synthetic data.

Before: base 57fa7580469a9b3dae5ebe99629107b7dd6438e7.
After: independently reviewed product commit ac2a4e50c85b349ee2a7db17a60832ab273db387.

The registered web Page renders IssuesPage, IssueSurface, the filter menu and normal hooks/QueryClient/ApiClient. Only Next's URL reader is adapted. Synthetic loopback responses contain one unread issue outside the selected Created/Updated interval. Worker/SharedWorker are unavailable, using the product's supported fallback. This is not full Next/server/auth/WebSocket/native E2E.

before-created.png: date-only Unread wrongly says No unread issues with no Clear filters.
after-created.png and after-updated.png: date-only filtered-empty copy and Clear filters.
after-clear.png: the known unread issue returns after clearing date plus priority, with Unread retained.

These four 1440x1000 PNGs are untouched browser screenshots, individually visually reviewed. They contain only synthetic issue data. No raw logs, API responses or credentials are published here.

SHA256:
03e4664a918afa02e7f55e81798e2f47460fd9a090712a6dac3647d694fa3b39  before-created.png
6b824986ae885316e25aac1e282856efd3f2d74719fcb0d7508a227740948c33  after-created.png
320954b6f9ce2c1d7f1a2e84baef771ed9a2ba17fdf069e83f12752360aacb7c  after-updated.png
fd21c24a301fe2260aabc6c7af3cc6ca593fd6bba54d1654d282a68275224979  after-clear.png
