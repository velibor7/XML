#!/bin/bash

service_name=$1
service_full_name="${service_name}_service"

go mod tidy

go get \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

go mod tidy

export PATH="$PATH:$(go env GOPATH)/bin"


protoc -I ./$service_full_name --go_out ./$service_full_name --go_opt paths=source_relative --go-grpc_out ./$service_full_name --go-grpc_opt paths=source_relative  --grpc-gateway_out ./$service_full_name --grpc-gateway_opt paths=source_relative ./$service_full_name/$service_full_name.proto