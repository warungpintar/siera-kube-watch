# siera-kube-watch
**siera kubewatch** is a Kubernetes events watcher that aims to publish incident (unexpected event) as a notification through webhooks.

Supported webhooks:
- slack
- workplace chat
- webhook

# Build and Run

## Install Dependencies
```
$ go mod download
```

## Build
```
$ go build
```

## Run
```
$ ./siera-kube-watch
```

# Configuration
Copy the `config/config.example.yaml` as `.config.yaml` and setup your own configuration as needed. 
