# Biblioteca

To start off, using mod, install dependencies:

```bash
go mod init {}
go mod download
```

## Run

```bash
go run .
```

This starts a server on port `9000`.

## Instructions

### Implementation details

```
Method     Path         Params        Body
-------------------------------------------
GET       /books        -             -
GET       /books/{id}   -             -
POST      /books        -            {...}
PUT       /books/{id}   -            {...}
PATCH     /books/{id}   -            {only send attributes which are to be updated}
DELETE    /books/{id}   -             -
```

### Test

```bash
./test-with-curl.sh
```
