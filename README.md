# siera-kube-watch
**siera kubewatch** is a Kubernetes events watcher that aims to publish incident (unexpected event) as a notification through webhooks.

Supported webhooks:
- slack
- telegram
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
Copy the `config.example.yaml` as `config.yaml` and setup your own configuration as needed. 

# Helm

We provide a helm chart for easy installation https://github.com/warungpintar/charts/tree/master/warpin/siera-kube-watch

## Result workplace chat
![GitHub Logo](docs/result-workchat.png)

## Result slack
![GitHub Logo](docs/result-slack.png)

## Result telegram
![GitHub Logo](docs/result-telegram.png)

docker image size  35.2MB