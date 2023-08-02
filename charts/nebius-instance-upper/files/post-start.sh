#!/bin/bash

set -eu -o pipefail

echo "[INFO] Running post-start-script for ${NODE_NAME} at ${NODE_IP}..."

SLEEP_TIME=5
DT=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

while true; do
  NODE_STATUS=$(kubectl get nodes "${NODE_NAME}" -ojson | jq -r '.status.conditions[] | select(.type == "Ready" and .lastHeartbeatTime >= "'${DT}'") | .status')

  if [[ "${NODE_STATUS}" == "True" ]]; then
    break
  fi

  sleep ${SLEEP_TIME}
done
