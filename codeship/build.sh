#!/usr/bin/env bash

# Exit script with error if any step fails.
set -e

# Echo out all commands for monitoring progress
set -x

# Build all the things
go build -ldflags="-s -w" -o bin/maintenance  cron/maintenance/main.go
go build -ldflags="-s -w" -o bin/outdatedrequests  cron/outdatedrequests/main.go
