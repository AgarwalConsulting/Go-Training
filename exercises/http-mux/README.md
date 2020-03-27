# Biblioteca

I am running a small library for the community in these difficult times. I need to manage the books in the library using the book management system. I have implemented some of the APIs but haven't had a chance to make it comprehensive. Can you help me?

## Install dependencies

To start off, using mod:

```bash
go mod init {}
go mod download
```

## Run

```bash
go run .
```

This starts a server on port `9000`.

Once you run the application it automatically installs the dependencies, updates go.mod and creates go.sum file.  Once the server is running, you will have to test the server using ./test-with-go.sh.

## Instructions

### Implementation details

Map the CRUD (Create, Read, Update & Destroy) operations to the appropriate end-point based on the table below:

| CRUD  |   Method |     Path     |    Params  |      Body                                        |  Response             |
|-------|----------|--------------|------------|--------------------------------------------------|-----------------------|
|  R    |   GET    |  /books      |  -         |    -                                             | [ {ID: , Title: ,} ]  |
|  R    |   GET    |  /books/{id} |  -         |    -                                             | {ID: , Title: ,}      |
| C     |   POST   |  /books      |  -         |   { Title: , Author: , }                         | {ID: , Title: ,}      |
|   U   |   PUT    |  /books/{id} |  -         |   { ID: , Title: , Author: , }                   | {ID: , Title: ,}      |
|   U   |   PATCH  |  /books/{id} |  -         |   {only send attributes which are to be updated} | {ID: , Title: ,}      |
|    D  |   DELETE |  /books/{id} |  -         |    -                                             | {ID: , Title: ,}      |

### Test

```bash
./test-with-curl.sh
```
