#!/usr/bin/env bash

# Generate server certificates
openssl req -newkey rsa:2048 \
  -new -nodes -x509 \
  -days 3650 \
  -out cert.pem \
  -keyout key.pem \
  -subj "/C=IN/ST=TN/L=Chennai/O=Agarwal Consulting.io/OU=algogrit/CN=localhost"

# Generate client certificates
openssl req -newkey rsa:2048 \
  -new -nodes -x509 \
  -days 3650 \
  -out client_cert.pem \
  -keyout client_key.pem \
  -subj "/C=IN/ST=TN/L=Chennai/O=Agarwal Consulting.io/OU=algogrit/CN=localhost"
