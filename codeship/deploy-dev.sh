#!/usr/bin/env bash

# Exit script with error if any step fails.
set -e

set -x

# Build binaries
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
"$DIR"/build.sh

# Export env vars
export SERVICE_INTEGRATION_URL="${DEV_SERVICE_INTEGRATION_URL}"
export SERVICE_INTEGRATION_TOKEN="${DEV_SERVICE_INTEGRATION_TOKEN}"

cd cron/maintenance
serverless deploy -v --stage dev
