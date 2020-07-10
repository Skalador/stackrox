#!/usr/bin/env bash

# Periodically check the cluster and SR deployment to debug test failures due to
# connectivity issues.

set +e

API_HOSTNAME=localhost
API_PORT=8000
if [[ "${LOAD_BALANCER}" == "lb" ]]; then
  API_HOSTNAME=$(./scripts/k8s/get-lb-ip.sh)
  API_PORT=443
fi
API_ENDPOINT="${API_HOSTNAME}:${API_PORT}"
METADATA_URL="https://${API_ENDPOINT}/v1/metadata"
echo "StackRox METADATA_URL is set to ${METADATA_URL}"

set -x

kubectl config get-contexts
context=$(kubectl config get-contexts | grep -E '^\*' | tr -s ' ' | cut -d ' ' -f 3)
echo "Pulled context: ${context}"
server=$(echo "${context}" | sed -En 's/(.*)\-(\w+)\-ci\-rox\-systems\-.*/\1.\2.ci.rox.systems/p')
echo "Pulled server: ${server}"

while true; do
  date
  if [[ ! -z "${server}" ]]; then
    host "${server}"
  fi
  kubectl get nodes
  ps axw | grep "port-forward"
  curl --silent --insecure --show-error "${METADATA_URL}" | jq
  sleep 60
done
