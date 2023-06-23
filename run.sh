#!/bin/bash

# Building container
docker build -t go-cipher-api .

# Running container
docker run -p 3000:3000 go-cipher-api
