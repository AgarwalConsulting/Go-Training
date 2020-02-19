#!/usr/bin/env bash

# GET /books
printf "Testing... GET /books: \n"
curl 'http://localhost:9000/books'

# POST
printf "\n\nTesting... POST /books: \n"
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"Title": "A Song of Ice and Fire", "Author": "George R.R. Martin", "ISBN": "1239000890", "Description": "The Game of Thrones book",
  "Price": 99.99}' \
  http://localhost:9000/books

# GET /books/:id
printf "\n\nTesting... GET /books/1: \n"
curl 'http://localhost:9000/books/1'

# PUT
printf "\n\nTesting... PUT /books/1: \n"
curl --header "Content-Type: application/json" \
  --request PUT \
  --data '{"title": "A Song of Ice and Fire", "Author": "George R.R. Martin", "ISBN": "0000000001", "Description": "The Game of Thrones book",
  "Price": 99.99}' \
  http://localhost:9000/books/1

# PATCH
printf "\n\nTesting... PATCH /books/1: \n"
curl --header "Content-Type: application/json" \
  --request PATCH \
  --data '{"ISBN": "0000000001"}' \
  http://localhost:9000/books/1

# DELETE
printf "\n\nTesting... DELETE /books/1: \n"
curl --header "Content-Type: application/json" \
  --request DELETE \
  http://localhost:9000/books/1

printf "\nCompleted!\n"
