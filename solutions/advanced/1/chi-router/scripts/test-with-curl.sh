#!/usr/bin/env bash

# GET /books
printf "Testing... GET /books: \n"
curl -i 'http://localhost:8000/books'

# POST
printf "\n\nTesting... POST /books: \n"
curl -i --header "Content-Type: application/json" \
  --request POST \
  --data '{
    "title": "A Song of Ice and Fire",
    "writer": "George R.R. Martin",
    "ISBNCode": "1238000890",
    "synopsis": "The Game of Thrones book",
    "price": 99.99
  }' \
  http://localhost:8000/books

# GET /books/:id
printf "\n\nTesting... GET /books/1: \n"
curl -i 'http://localhost:8000/books/1'

# PUT
printf "\n\nTesting... PUT /books/1: \n"
curl -i --header "Content-Type: application/json" \
  --request PUT \
  --data '{
    "title": "A Song of Ice and Fire",
    "writer": "George R.R. Martin",
    "ISBNCode": "0000000001",
    "synopsis": "The Game of Thrones book",
    "price": 99.99
  }' \
  http://localhost:8000/books/3

# PATCH
printf "\n\nTesting... PATCH /books/1: \n"
curl -i --header "Content-Type: application/json" \
  --request PATCH \
  --data '{"synopsis": "Winter is coming..."}' \
  http://localhost:8000/books/3

# DELETE
printf "\n\nTesting... DELETE /books/1: \n"
curl -i --header "Content-Type: application/json" \
  --request DELETE \
  http://localhost:8000/books/1

printf "\nCompleted!\n"
