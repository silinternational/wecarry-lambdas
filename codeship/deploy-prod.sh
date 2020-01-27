#!/usr/bin/env bash

# Exit script with error if any step fails.
set -e

# Build binaries
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
"$DIR"/build.sh

# Export env vars
export SERVICE_INTEGRATION_URL="${PROD_SERVICE_INTEGRATION_URL}"
export SERVICE_INTEGRATION_TOKEN="${PROD_SERVICE_INTEGRATION_TOKEN}"

cd api/admin
serverless deploy -v --stage prod

cd ../agent
serverless deploy -v --stage prod
