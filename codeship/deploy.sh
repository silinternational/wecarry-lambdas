#!/usr/bin/env bash

# Exit script with error if any step fails.
set -e

# Build binaries
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
"$DIR"/build.sh

# Export env vars
if [ "$1" = "prod" ]; then

  export AWS_ACCESS_KEY_ID="${PROD_AWS_ACCESS_KEY_ID}"
  export AWS_SECRET_ACCESS_KEY="${PROD_AWS_SECRET_ACCESS_KEY}"
  export SERVICE_INTEGRATION_URL="${PROD_SERVICE_INTEGRATION_URL}"
  export SERVICE_INTEGRATION_TOKEN="${PROD_SERVICE_INTEGRATION_TOKEN}"

elif [ "$1" = "stg" ]; then

  export AWS_ACCESS_KEY_ID="${STG_AWS_ACCESS_KEY_ID}"
  export AWS_SECRET_ACCESS_KEY="${STG_AWS_SECRET_ACCESS_KEY}"
  export SERVICE_INTEGRATION_URL="${STG_SERVICE_INTEGRATION_URL}"
  export SERVICE_INTEGRATION_TOKEN="${STG_SERVICE_INTEGRATION_TOKEN}"

else

  echo "invalid stage $1"
  exit 1

fi

echo "Deploying stage $1..."
$HOME/.serverless/bin/serverless deploy --verbose --stage "$1"
