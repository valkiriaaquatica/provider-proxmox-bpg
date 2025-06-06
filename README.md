# Proxmox BPG Provider

`provider-proxmox-bpg` is a [Crossplane](https://crossplane.io/) provider built using
[Upjet](https://github.com/crossplane/upjet). It exposes XRM-conformant managed
resources for the [Proxmox Virtual Environment](https://www.proxmox.com/) API.
## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/upbound/provider-proxmox-bpg):
```
up ctp provider install provider-proxmox-bpg:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: upjet-provider-template
spec:
  package: upbound/upjet-provider-template:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/valkiriaaquatica/provider-proxmox-bpg)

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/valkiriaaquatica/provider-proxmox-bpg/issues).
