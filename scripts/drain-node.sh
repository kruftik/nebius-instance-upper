#!/bin/bash

set -eu -o pipefail

if [[ "$#" -eq 0 ]]; then
  echo "[INFO] Running drain-node-script for ${NODE_NAME} at ${NODE_IP}..."

  kubectl drain --ignore-daemonsets --delete-emptydir-data "${NODE_NAME}"
elif [[ "${1}" == "revert" ]]; then
  echo "[INFO] Running REVERT drain-node-script for ${NODE_NAME} at ${NODE_IP}..."

  kubectl uncordon "${NODE_NAME}"
else
  echo "[WARN] unknown sub-action: ${1}"
fi

