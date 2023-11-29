#!/bin/bash

if [[ -z "${KUBECONFIG}" ]]; then
    echo "KUBECONFIG is not set!"
    exit 1
fi

desired_tag=$1
entrypoint_manifest=deployments/kubernetes/entrypoint/entrypoint.yml
if [[ ! -z "$desired_tag" ]]; then
    image=$(grep -oP 'image: \K.*' $entrypoint_manifest)
    sed -i.bak "s|$image|$image:$desired_tag|" $entrypoint_manifest
    for manifest in deployments/kubernetes/micros/*; do
        image=$(grep -oP 'image: \K.*' $manifest)
        sed -i.bak "s|$image|$image:$desired_tag|" $manifest
    done
fi

kubectl -f $entrypoint_manifest delete
kubectl -f $entrypoint_manifest apply
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

if [[ ! -z "$desired_tag" ]]; then
    mv $entrypoint_manifest.bak $entrypoint_manifest
    for manifest in deployments/kubernetes/micros/*.yml; do
        mv $manifest.bak $manifest
    done
fi