#!/usr/bin/env bash

set -ex


helm repo add grafana https://grafana.github.io/helm-charts

# Create dashboard configmap from local files
kubectl create configmap grafana-dashboards \
    --from-file=./dashboards/coredns.json
# Annotate configmap so grafana chart knows where to look for them
kubectl patch configmap grafana-dashboards --patch \
    '{"metadata":{"labels":{"grafana_dashboard": "true"}}}'

helm install grafana -f ./grafana-values.yaml grafana/grafana --version 6.4.4
