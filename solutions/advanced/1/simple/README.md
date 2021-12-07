# Biblioteca

I am running a small library for the community in these difficult times. I need to manage the books in the library using the book management system. I have implemented some of the APIs but haven't had a chance to make it comprehensive. Can you help me?

Notes for HTTP: [here](https://github.com/AgarwalConsulting/Go-Training/tree/master/examples/11-A-01-net/http).

## Install dependencies

To start off, using mod:

```bash
go mod init <your-module-name>
```

## Run

```bash
go run .
```

This starts a server on port `9000`.

Once you run the application it automatically installs the dependencies, updates go.mod and creates go.sum file. Once the server is running, you will have to test the server using `./test-with-curl.sh`.

## Using modules

```bash
go mod download
```

## Instructions

### Implementation details

Map the CRUD (Create, Read, Update & Destroy) operations to the appropriate end-point based on the table below:

| CRUD  | Action   |   Method |     Path     |    Params  |      Body                                        |  Response             |
|-------|----------|----------|--------------|------------|--------------------------------------------------|-----------------------|
|  R    | Index    |   GET    |  /books      |  -         |    -                                             | [ {ID: , Title: ,} ]  |
|  R    | Show     |   GET    |  /books/{id} |  -         |    -                                             | {ID: , Title: ,}      |
| C     | Create   |   POST   |  /books      |  -         |   { Title: , Author: , }                         | {ID: , Title: ,}      |
|   U   | Update   |   PUT    |  /books/{id} |  -         |   { ID: , Title: , Author: , }                   | {ID: , Title: ,}      |
|   U   | Update   |   PATCH  |  /books/{id} |  -         |   {only send attributes which are to be updated} | {ID: , Title: ,}      |
|    D  | Delete   |  DELETE |  /books/{id} |  -         |    -                                             | {ID: , Title: ,}      |

### Test

```bash
./test-with-curl.sh
```

### Questions and Hints

- Look at `encoding/json` for decoding req.Body
- Is the API json compliant? If not, why?
  - What would it take to turn it in to `xml`?
  - Can you support both?
- Can you make the field names in json `camelCase`?
  - What happens when you change the attribute name of books to start with lower case?
- Handle edge cases too!
  - What happens when a user tries to delete a non-existent book?
  - Don't accept any price <5. See: [Validator](https://github.com/go-playground/validator)
- What happens when you make concurrent requests?
