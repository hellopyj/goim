#!/usr/bin/env bash
protoc -I ./ ./handler.proto --go_out=plugins=grpc:./
