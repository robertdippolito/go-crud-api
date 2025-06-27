# k8s-api

This repository contains a minimal Go API and the Helm chart used to deploy it.

## Helm charts

The `_infra` directory contains all Helm related files. The existing workflow
installs the `k8s-api` chart located under `_infra/k8s-api-chart`.

### Ingress NGINX

To expose the API through an ingress controller with metrics enabled, a values
file is provided at `_infra/ingress-nginx-values.yaml`. This file configures the
`ingress-nginx` chart with Prometheus scraping and a `ServiceMonitor`.

Workflows deploy the ingress controller using this file, but you can also run it
locally:

```bash
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm upgrade --install ingress-nginx ingress-nginx/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --version 4.7.1 \
  --values _infra/ingress-nginx-values.yaml
```
