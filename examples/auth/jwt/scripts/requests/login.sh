#!/usr/bin/env bash

curl -i -k \
 --data '{
    "username": "algogrit",
    "password": "drowssap"
  }' \
  http://localhost:8000/login
