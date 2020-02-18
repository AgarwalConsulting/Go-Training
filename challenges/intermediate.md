# Intermediate

## 0. Implement your own `filter` function

  A filter function will take a slice & a predicate which evaluates to `bool`. The filter fn should return a slice back of same type.

Eg. Take a slice of string: `[]string{"Iron Man", "Batman", "Superman", "Spider-man", "Wonder Woman", "Iron Fist", "Daredevil", "Supergirl", "Flash"}`.

```golang
func filter(s []string, predicate func(string) bool) []string {

}
```

Some predicate examples...

- Display all heroes whose name's second character is a vowel (a, e, i, o u)
- Display all heroes whose name doesn't contain "man" in it

## 1. [Fibonacci closure](https://tour.golang.org/moretypes/26)

  Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

  ```golang
  package main

  import "fmt"

  // fibonacci is a function that returns
  // a function that returns an int.
  func fibonacci() func() int {
  }

  func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
      fmt.Println(f())
    }
  }
  ```

## 2. [rot13Reader](https://tour.golang.org/methods/23)

## 3. Implement linked list

## 4. Exercism - PaasIO

## 5. Write your own http server using standard library

Implement a simple http server.

- Send your name from the client and process it on the server
- Switch the data types for the payload, see how they behave
- Write few more handlers and call them on the server side
- Experiment with `encoding/xml`
