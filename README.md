# Proxmox BPG Provider

`provider-proxmox-bpg` is a [Crossplane](https://crossplane.io/) provider built using
[Upjet](https://github.com/crossplane/upjet). It exposes XRM-conformant managed
resources for the [Proxmox Virtual Environment](https://www.proxmox.com/) API. This provider it is already compatible with crossplane v2.

## 🚀 Release Automation

This repository automatically publishes a new release every time the upstream Terraform provider releases a new version:
👉 **Original Terraform Provider:** [bpg/terraform-provider-proxmox](https://github.com/bpg/terraform-provider-proxmox)

The automation works through a fully-integrated release pipeline:

1. **Version Detection** – Handled by our configured [Renovate Bot](https://github.com/valkiriaaquatica/provider-proxmox-bpg/blob/main/.github/renovate.json), which tracks new upstream releases.
2. **Version Preparation** – When Renovate opens a PR, the custom [Prepare New Version GitHub Action](https://github.com/valkiriaaquatica/provider-proxmox-bpg/blob/main/.github/workflows/prepare-new-version.yaml) adjusts the module and prepares the next tag.
3. **Automated Publishing** – Once the PR is merged, [Release Please](https://github.com/valkiriaaquatica/provider-proxmox-bpg/blob/main/.github/workflows/release-please.yml) generates changelogs and publishes the release automaticaly

Thanks to this pipeline, **this repository always stays up-to-date with the latest upstream provider version** — no manual intervention required, **JUST** when a new CRD (resource) is added (that will be future automated).

---

## 🔧 For Upjet / Crossplane Developers

If you’re building a Crossplane provider using **Upjet**, feel free to reuse or adapt this release configuration to automate your own provider’s lifecycle.
Having an auto-syncing provider dramatically simplifies maintenance and ensures your Crossplane collection stays aligned with upstream changes.


## Installation (make sure you have Crossplane before installed in your cluster)
- Using [up](https://docs.upbound.io/reference/cli/):
  Install the provider by using the following command after changing the image tag
  to the [latest release](https://marketplace.upbound.io/providers/upbound/provider-proxmox-bpg):
  ```
  up ctp provider install xpkg.upbound.io/valkiriaaquaticamendi/provider-proxmox-bpg:v1.0.0
  ```
- Declarative installation
  ```
  cat <<EOF | kubectl apply -f -
  apiVersion: pkg.crossplane.io/v1
  kind: Provider
  metadata:
    name: valkiriaaquaticamendi-provider-proxmox-bpg
  spec:
    package: xpkg.upbound.io/valkiriaaquaticamendi/provider-proxmox-bpg:v1.0.0
  EOF
  ```
  or
  ```
  kubectl apply -f examples/install.yaml
  ```
  Now create the seecret with your Proxmox credentials, filling the secret and apply it
  ```
  vi examples/providerconfig/secret.yaml.tmpl
  kubectl apply -f examples/providerconfig/secret.yaml.tmpl
  ```
  Then create the Provider configuration using that secret
  ```
  kubectl apply -f examples/providerconfig/providerconfig.yaml
  ```

  In the folder examples/ and examples-generated/ you can have multiple examples to quick create. If you have any interesting example to add, feel free to contribute. examples/ folder is based on more testes examples while the examples-generated/ wrap the examples from Terraform docs  into Yamls.


## Developing
1. Run the generator

   ```bash
   make generate
   ```
2. Run the image  against an existant Kubernetes cluster that has Crossplane already installed
  Install the CRDs in the cluster:
    ```
    kubectl apply -f package/crds/
    ```
3. Run the image  against an existant Kubernetes cluster that has Crossplane already installed
    and test it works well
   ```bash
   make run
   ```
4. Run the tests
    and test it works well
   ```bash
   make test
   ```
4. Run the local docker build image
    and test it works well
   ```bash
   make build
   ```
---

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/valkiriaaquatica/provider-proxmox-bpg/issues).

## TODO
A action that can be maanually and launched once every day that checks if a new resource is created, as example of that is this: https://github.com/bpg/terraform-provider-proxmox/pull/2502. Note with that example and others that changes to trigger are: docs/resources (if it is documented :) 