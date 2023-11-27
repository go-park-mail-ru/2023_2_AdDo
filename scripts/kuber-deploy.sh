#!/bin/bash

if [[ -z "${KUBECONFIG}" ]]; then
    echo "KUBECONFIG is not set!"
    exit 1
fi

ENTRYPOINT_MANIFEST=deployments/kubernetes/entrypoint/entrypoint.yml
kubectl -f $ENTRYPOINT_MANIFEST delete
kubectl -f $ENTRYPOINT_MANIFEST apply
if [ $? -ne "0" ]; then
    echo "Error apply entrypoint.yml manifest"
    exit 1
fi

for manifest in deployments/kubernetes/micros/*; do
    kubectl -f $manifest delete
    kubectl -f $manifest apply
    if [ $? -ne "0" ]; then
        echo "Error apply $manifest manifest"
        exit 1
    fi
done
