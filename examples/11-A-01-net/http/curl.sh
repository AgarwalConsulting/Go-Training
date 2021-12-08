#!/usr/bin/env bash

printf "GET /employees"
curl -i "localhost:8000/employees"

printf "POST /employees"
curl -X POST -i "localhost:8000/employees" --data '{"name": "Nikhil", "speciality": "SRE", "project": 1012}'
