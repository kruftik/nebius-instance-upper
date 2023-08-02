#!/bin/bash

set -eu -o pipefail

echo "[INFO] Running pre-stop-script for ${NODE_NAME} at ${NODE_IP}..."

SLEEP_TIME=5

while true; do
  NODE_STATUS=$(kubectl get nodes "${NODE_NAME}" -ojson | jq -r '.status.conditions[] | select(.type == "Ready") | .status')

  if [[ "${NODE_STATUS}" == "True" ]]; then
    break
  fi

  sleep ${SLEEP_TIME}
done
