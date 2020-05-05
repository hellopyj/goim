#!/usr/bin/env bash
protoc -I ./ ./push.proto --go_out=plugins=grpc:./
