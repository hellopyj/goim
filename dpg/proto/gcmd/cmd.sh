#!/usr/bin/env bash
protoc -I ./ ./cmd.proto --gogo_out=plugins=grpc:./