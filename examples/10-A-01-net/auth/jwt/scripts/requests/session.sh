#!/usr/bin/env bash

curl -i http://localhost:8000/currentUser \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiRyBBIiwidXNlcklEIjoxfQ.KmGH7ltXYETggTZK8aBxAwDLFVxdCpb9GPymof3jVvA' \
--header 'Content-Type: application/json'
