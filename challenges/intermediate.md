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

### Extra Challenge

- Make the `filter` function accept variadic number of predicates
- Make the `filter` function "generic"

*Hint*: You can use the [`strings`](https://pkg.go.dev/strings) package

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

## 2. [stringer challenge](https://tour.golang.org/methods/18)

## 3. Reading from a JSON/XML file

Given the following contents in a file: [`employee.json`](https://github.com/AgarwalConsulting/Go-Training/tree/master/exercises/basic/reading/employee.json)

```json
{"Name": "Gaurav", "Designation": "Director of Engineering", "Department": "LnD", "ProjectID": 1001}
```

Read these contents into a Go DS of your choice. Look at `encoding/json` package.

## Steps

- Create a `/tmp/employee.json` (or `%TEMP%\employee.json` in windows) in your system with the above contents before you begin.

- In your Go program,

  - Open the file

  - Read the contents

  - Print only the Name of the employee

### Extra Challenge

- Read from [`employees.json`](https://github.com/AgarwalConsulting/Go-Training/tree/master/exercises/basic/reading/employees.json) and print all the unique project IDs.

## 4. [rot13Reader](https://tour.golang.org/methods/23)

<details>
  <summary>Hint</summary>

  ```golang
  package main

  import (
    "io"
    "os"
    "strings"
  )

  type rot13Reader struct {
    r io.Reader
  }

  func (reader *rot13Reader) Read(p []byte) (n int, err error) {
    n, err = reader.r.Read(p)

    // Perform the rot13 substitution

    return
  }

  func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
  }
  ```
</details>

## 5. Exercism - [PaasIO](https://github.com/AgarwalConsulting/Go-Training/tree/master/exercises/exercism/paasio)

## 6. Write your own http server using standard library

Implement a simple http server.

- Send your name from the client and process it on the server
- Switch the data types for the payload, see how they behave
- Write few more handlers and call them on the server side
- Experiment with `encoding/xml`
