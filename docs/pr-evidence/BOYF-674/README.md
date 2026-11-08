Invitation acceptance: legacy pending invitation for an existing member

These unedited 1280×900 screenshots were captured by an independent review run from the registered Web InvitationsRoutePage and shared InvitationsPage, normal QueryClient/hooks/ApiClient, and a local product router with JWTAuth and real PostgreSQL transactions. Every actor, workspace, invitation, and JWT was synthetic in a disposable task-owned environment.

- Before (`73975653323750d04c27047ea9c986ef06bdf6a1`): normal new-member acceptance is 200; legacy pending invitation acceptance is 409 and rolls back, preserving the pending row.
- After (`cbd95169cec7f6a3e880a3d4fe969f919483276e`): legacy acceptance is 200; the original membership ID and current admin role are preserved; the invitation is accepted and disappears from the pending list.

The empty-list “Continue to setup” screen is the harness page state after acceptance, not evidence of navigation through the complete workspace shell. Session storage and navigation are harness adapters; this is not full login, Next.js, native, WebSocket, or cloud/billing E2E. The one-connection-pool matrix is separately tested at the handler with real PostgreSQL, not through this UI harness.

This branch holds review materials only. It does not modify the product change or open another pull request.
