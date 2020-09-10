# Overview

Go's defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns.

## Defer

- `defer` keyword
- `defer <funcCall>`
- stacked execution - [FILO](https://tour.golang.org/flowcontrol/13)
- Clean up with defer
- Args will be evaluated before deferring
- defer executes even in case of panic
