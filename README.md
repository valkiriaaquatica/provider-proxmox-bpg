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
  name: provider-proxmoxbpg
spec:
  package: crossplane/provider-proxmoxbpg:v0.1.0
EOF
```


Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/valkiriaaquatica/provider-proxmox-bpg)

## Developing

### (Optional) -> Intitializate the devbox environment
If you have devbox or want to work with it, it makes life easier for packages like go, do the following:
```console
cd devbox/
devbox install
devbox shell
```

###  Important --> Fix for `make generate`

When running `make generate`, the tool automatically pulls documentation files from the Terraform provider (`/docs`). However, a known issue — also seen in other providers like [`provider-confluent`](https://github.com/crossplane-contrib/provider-confluent?tab=readme-ov-file#getting-started), causes the process to fail if the generated Markdown files do not include both `page_title` and `description` in their front matter.

Example error:

```
../.work/bpg/proxmox/docs/resources/virtual_environment_acl.md: failed to find the prelude of the document using the xpath expressions: //text()[contains(., "description") and contains(., "page_title")]
```

---

### Steps to apply the fix

1. Clone the provider and prepare the local environment:

   ```bash
   make generate
   # This will pull the provider locally and likely return the above error
   ```

2. Make the fix script executable:

   ```bash
   chmod +x pre-make-generate.sh
   ```

3. Run the script to patch the Markdown files:

   ```bash
   ./pre-make-generate.sh
   ```

4. Re-run the generator:

   ```bash
   make generate
   ```

It should now work correctly without `description`/`page_title` errors.

---


Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```
Checkout sub-repositories:

```console
make submodules
```

Execute code generation:

```console
make generate
```


Run against a Kubernetes cluster:

Install the CRDs in the cluster:
```console
kubectl apply -f package/crds/
```

Run the image and keep this terminal to see the logs:
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


## TODO list

- Pending support due to schema error and not added (Upjet issue [#372](https://github.com/crossplane/upjet/issues/372)):
  - `proxmox_virtual_environment_cluster_options`
  - `proxmox_virtual_environment_hardware_mapping_dir`
  - `proxmox_virtual_environment_hardware_mapping_usb`
  - `proxmox_virtual_environment_hardware_mapping_pci`

- Resources not implementing `Get` properly: -> TRY MORE
  - `proxmox_virtual_environment_virtualenvironmentcertificate`
  - `proxmox_virtual_environment_virtualenvironmentdatastores`

- Token creation fails with error: Error reading user token: error retrieving user token
  - `proxmox_virtual_environment_user_token`

- Test upload functionality for resources using `"file"` fields

- Investigate `proxmox_virtual_environment_metrics_server`:
  - See implementation in [`metrics/server.go`](https://github.com/bpg/terraform-provider-proxmox/blob/main/proxmox/cluster/metrics/server.go)
  - Error during observe phase:  
    `error unmarshalling json with lists; observe failed: cannot run refresh: refresh failed`

- Test and adapt `Makefile` to use Terraform provider `v0.78.3`
  - Includes fixes for several broken resources
