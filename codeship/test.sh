#!/usr/bin/env bash

# https://github.com/securego/gosec#available-rules
# G104 ignore errors not checked
gosec -exclude=G104 -quiet ./...

go test ./... -v
