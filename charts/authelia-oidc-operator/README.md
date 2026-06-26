### Testing the chart output
Make the deployment test yaml:
```bash
IMG="ghcr.io/milas/authelia-oidc-operator:latest" make test-deploy
```

Render the chart to a file:
```bash
cd /charts/authelia-oidc-operator
helm dependency update
helm template authelia-oidc-operator . --namespace security -f values.yaml > authelia-oidc-operator.yaml
```

Compare the chart output file `authelia-oidc-operator.yaml` with the `test-deployment.yaml` file.
