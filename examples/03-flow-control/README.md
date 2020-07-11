# Overview

## For

- normal `for <init>; <condition>; <post> {`
- while `for <condition> {`
- forever `for {`
  - can control execution using `break` & `continue`

## if

- with [statements](https://tour.golang.org/flowcontrol/6)
- `if <statement>; <condition> {`
  - `var`s declared within statement are available only within `if / else if / else` block

## [Switch statements](https://gobyexample.com/switch)

- no break
- statements are optional
- top-down execution - where execution stops after a case succeeds
- normal `switch <value> { case <match>`
- takes statement `switch <statement>; <value> { case <match>`
- conditional `switch { case <bool-statement>: `
  - `if-else-if`

## [Ranges](https://gobyexample.com/range)

- `for` along with `range`
  - range returns `index` & `element`
  - `for index, element := range "abc" {`
  - `for <index>, <value> := range <var> {`
- For maps:
  - `for <key>, <value> := range <map-type> {`
