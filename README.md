# Helm Chart for http-echo-headers

The code is provided as-is with no warranties.

## Usage
[Helm](https://helm.sh) must be installed to use the charts. Please refer to Helm's [documentation](https://helm.sh/docs/) to get started.

Once Helm is set up properly, add the repo as follows:
```shell
helm repo add echo-headers https://abhinav1107.github.io/http-echo-headers
helm repo update
```

You can then run `helm search repo echo-headers` to see the charts.

## Deploy http-echo-headers to your cluster

### Deploy with default config
```shell
helm upgrade --install echo-headers echo-headers/echo-headers
```
