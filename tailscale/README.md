# Tailscale Daemonset

This directory contains a dockerfile for building an alpine image with Tailscale
installed.


## Instructions

```bash
kubectl apply -f daemonset.yaml
for po in $(k get po -l app=tailscale -o json | jq -r .items[].metadata.name ); do
    kubectl exec -it $po -- /usr/sbin/tailscale up
    # open the login link printed out
done
```
