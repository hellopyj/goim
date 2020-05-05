#!/usr/bin/env bash
protoc -I ./ ./api.proto --proto_path=/Users/hellopyj/Dev/go/pkg/mod --proto_path=/Users/hellopyj/Src/Go/goim --gofast_out=plugins=grpc:./
