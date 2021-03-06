#!/usr/bin/env bash

# Exit script with error if any step fails.
set -e

# Build binaries
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
"$DIR"/build.sh

# Export env vars
export AWS_ACCESS_KEY_ID="${PROD_AWS_ACCESS_KEY_ID}"
export AWS_SECRET_ACCESS_KEY="${PROD_AWS_SECRET_ACCESS_KEY}"
export SERVICE_INTEGRATION_URL="${PROD_SERVICE_INTEGRATION_URL}"
export SERVICE_INTEGRATION_TOKEN="${PROD_SERVICE_INTEGRATION_TOKEN}"

cd cron/maintenance
serverless deploy -v --stage prod
