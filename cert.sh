#!/usr/bin/env bash

openssl req \
  -newkey rsa:2048 \
  -x509 \
  -nodes \
  -keyout server.key \
  -new \
  -out server.pem \
  -subj /CN=localhost \
  -extensions san \
  -config \
  <(echo "[req]";
    echo distinguished_name=req;
    echo "[san]";
    echo subjectAltName=DNS:localhost,IP:10.0.0.1
   ) \
  -sha256 \
  -days 3650

#openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.pem -days 3650 -nodes -subj '/CN=localhost'
