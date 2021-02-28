#!/usr/bin/env bash

set -e
set -o pipefail

helm repo add stable https://charts.helm.sh/stable
# Get the latest chart list
helm repo update

helm install prometheus -f ./prometheus-values.yaml stable/prometheus --version 11.4.0

