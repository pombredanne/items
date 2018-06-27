#!/bin/sh
protoc -I . --go_out=plugins=grpc:. ./data_sync.proto

