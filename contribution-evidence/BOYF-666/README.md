# Duplicate-mark feedback — running UI evidence

Raw, unedited browser screenshots from the registered Web issue page, its status action, ModalRegistry, shared picker, normal useUpdateIssue, application createQueryClient and ApiClient.

- `before-silent409.png`: base 73975653323750d04c27047ea9c986ef06bdf6a1; the delayed target-became-duplicate409 rolls back to Todo without the recovery message.
- `after-error409.png`: reviewed product head ffed7618d6d352470d712f021839b73ea463c0de; the same action shows the FIX-2 recovery instruction exactly once after picker unmount.
- `after-success.png`: that same reviewed head; cancelled state, duplicate relation, and one complete success toast.

Only synthetic actor/workspace/FIX-1..3 and a task-owned ephemeral loopback HTTP fixture were used. Browser requests were restricted to that origin, with only the source-issue PUT and automatic inbox-read POST permitted. No real business data, backend, authentication, WebSocket, Electron/native or full Next application is represented. Desktop shares this registered component by source inspection only.

Independent verification: baseline normal business control passes; baseline missing-error target fails as expected; all five HEAD browser scenarios pass; three adjacent Vitest suites pass 13 tests; views TypeScript 7.0.2/6.0.3, changed-file ESLint and UI guards pass. AI agents authored the code and performed the independent review. These images are direct running-browser output, not generated or edited images.

This materials-only branch starts at the reviewed base and does not alter the product head or add a product PR.
