# Synthetic issue-load failure evidence

These PNGs support BOYF-471's primary issue read error-feedback change. They contain synthetic issue data only, using the real shared production React components and ApiClient with local HTTP fixtures in headless Chromium. They were captured by the development agent; independent review inspected the material and separately reran the production-component HTTP tests in Vitest/jsdom. They are not independent browser reruns, native-desktop checks, real permission tests, or an application-wide E2E result.

- Product base: `f230f9f55bdfc4de68015473b07de329a5079f87`.
- Final reviewed product candidate: `5f8e2bd3808b08ae19927f54ba66efd520c0902b`.
- Original 13 PNGs: captured during the original `01b5423a631bf25e09ea9de966c5d10ffc57c2a3` development; reused byte-for-byte. That earlier commit failed independent review for temporary 408/429 classification and is not approved on its own. Its 503/network and 403/404 surfaces are unchanged by the additive repair.
- Four new 408/429 PNGs: captured against final `5f8e2bd` with keyboard Retry followed by synthetic HTTP 200 recovery.
- Recovery images include an unrelated Wakeups loading error. Only primary issue content recovery is demonstrated.
- The original permanently-pending issue's cause remains unknown. This demonstrates rejected-read feedback only, not an unsettled request, offline pause or staged-child fix.

This branch contains public review material only and has no separate product PR or contribution credit. PNGs are unedited originals. Fixture provenance and full test logs are delivered on the Multica review issue.

| File | Bytes | SHA256 |
| --- | ---: | --- |
| base-identifier-403.png | 20321 | `3b21a83bccecaf59aa70d814a3c7446b50d5dc94367a423f19c532d03f72b7ef` |
| base-identifier-404.png | 20321 | `3b21a83bccecaf59aa70d814a3c7446b50d5dc94367a423f19c532d03f72b7ef` |
| base-identifier-503.png | 20321 | `3b21a83bccecaf59aa70d814a3c7446b50d5dc94367a423f19c532d03f72b7ef` |
| base-inbox-503.png | 23291 | `7198d216b3b69c7122ffc853ba7838bdef89a510be26729008efc1d64846e71e` |
| base-peek-503.png | 23500 | `560fc5bd0ef4d016b80390aa5a758f345fd13b89b8da84e02085ec8dd7f4c10b` |
| base-uuid-network.png | 20321 | `3b21a83bccecaf59aa70d814a3c7446b50d5dc94367a423f19c532d03f72b7ef` |
| head-identifier-403.png | 20185 | `c7cae197c4bb34e29f2e13ad26cd3b493787d496490bbe4c0f5f4b282e773087` |
| head-identifier-404.png | 20185 | `c7cae197c4bb34e29f2e13ad26cd3b493787d496490bbe4c0f5f4b282e773087` |
| head-identifier-503.png | 15441 | `ce96209cdda03a3b716a316c795b2cf248951e6c9565e37ab486c9b22aebe061` |
| head-inbox-503.png | 18693 | `773b37dc7a11539f0c59695bb40ca3c0557275a6c33c5433ce6edb1fb486f0ef` |
| head-peek-503.png | 18585 | `611f8a9d39f9f60f96e812ffd781f6746168543e298b65ba4faf9cb15446723d` |
| head-retry-recovered.png | 42369 | `984a83983d7ecff7acec944064cdbe212ad9f86c97941984972f87c96d58ad20` |
| head-uuid-network.png | 15441 | `ce96209cdda03a3b716a316c795b2cf248951e6c9565e37ab486c9b22aebe061` |
| head-408-retry-recovered.png | 42369 | `984a83983d7ecff7acec944064cdbe212ad9f86c97941984972f87c96d58ad20` |
| head-429-retry-recovered.png | 42369 | `984a83983d7ecff7acec944064cdbe212ad9f86c97941984972f87c96d58ad20` |
| head-identifier-408.png | 15441 | `ce96209cdda03a3b716a316c795b2cf248951e6c9565e37ab486c9b22aebe061` |
| head-identifier-429.png | 15441 | `ce96209cdda03a3b716a316c795b2cf248951e6c9565e37ab486c9b22aebe061` |
