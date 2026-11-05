# Calendar keyboard behavior evidence

These four unedited PNGs were originally captured by the BOYF-504 development step using Chromium 148 and synthetic dates in Asia/Shanghai. They show real production Calendar components without application CSS; they are behavior evidence, not full application styling or backend/Electron/Safari/accessibility-technology testing.

Before: base 98846ce564a6a7a0d26cb4f071263700003d3041. After: original Calendar commit cb6a5debe9d1debb4bd9d5329e97285e3d9d6e6a. Product candidate aa7ff5efece7f1969ca3b75656fb4c571a4196d6 keeps every original Calendar, test, DateOnlyPicker and project wrapper byte unchanged; only the two reviewed guard scripts were appended. These are original captures, not new-head screenshots. The independent reviewer replayed the strict base/head keyboard contract and 100 navigation cases on aa7ff5e.

At 2026-11-05, Tab into Calendar, press ArrowRight, then Enter. Before, DOM focus remains on 5 and commits 2026-11-05. After, actual focus moves to 6 and commits 2026-11-06. The outputs below the calendar contain only these synthetic dates and local calls.

Source archive: BOYF-504-calendar-focus-evidence.zip, SHA256 932b9bf66b4bf18706883593bd5300c08d28348ed5d48e9ba81275ea5106ce3d. Original 117-entry manifest and the later 237-entry integration manifest were independently verified.

| File | Original archive path | SHA256 |
|---|---|---|
| expanded-base-Asia-Shanghai-focus.png | evidence/expanded-base/Asia-Shanghai-focus.png | 40b0e5386b47ee0a31e1e09c3381f7d3f7292f91c94851c1f173d9f9d9d282c5 |
| expanded-base-Asia-Shanghai-submitted.png | evidence/expanded-base/Asia-Shanghai-submitted.png | 35b1da92d723224d8891c2e7e13990889f665e01a894718a9129d297af9d4b3b |
| expanded-head-Asia-Shanghai-focus.png | evidence/expanded-head/Asia-Shanghai-focus.png | e72be79d047ae85f45284d384308a424d51b30cada8ee3d74f1d62041161841c |
| expanded-head-Asia-Shanghai-submitted.png | evidence/expanded-head/Asia-Shanghai-submitted.png | c6e95ddca505d0652d80a1cbbf0c37b958a1004ae37ee160f65b762938b2a934 |
