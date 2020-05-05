#!/usr/bin/env bash
protoc -I ./ ./api.proto --proto_path=/Users/hellopyj/Dev/go/pkg/mod  --gofast_out=plugins=grpc:./
