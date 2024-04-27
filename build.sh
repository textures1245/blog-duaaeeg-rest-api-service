#!/bin/bash

# Download the dependencies
go mod download 

# Prefetch the binaries
go run github.com/steebchen/prisma-client-go prefetch

# Generate the Prisma Client Go client
go run github.com/steebchen/prisma-client-go generate

# Build the binary with all dependencies
go build -o  ./app ./api

# Run the binary
./app