# Proxmox BPG Provider

`provider-proxmox-bpg` is a [Crossplane](https://crossplane.io/) provider built using
[Upjet](https://github.com/crossplane/upjet). It exposes XRM-conformant managed
resources for the [Proxmox Virtual Environment](https://www.proxmox.com/) API.

## Installation (make sure you have Crossplane before installed in your cluster)
- Using [up](https://docs.upbound.io/reference/cli/):
  Install the provider by using the following command after changing the image tag
  to the [latest release](https://marketplace.upbound.io/providers/upbound/provider-proxmox-bpg):
  ```
  up ctp provider install xpkg.upbound.io/valkiriaaquaticamendi/provider-proxmox-bpg:v0.1.0
  ```
- Declarative installation
  ```
  cat <<EOF | kubectl apply -f -
  apiVersion: pkg.crossplane.io/v1
  kind: Provider
  metadata:
    name: valkiriaaquaticamendi-provider-proxmox-bpg
  spec:
    package: xpkg.upbound.io/valkiriaaquaticamendi/provider-proxmox-bpg:v0.1.0
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
   # take a look to your locak .work/ folder :)
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

It should now work correctly without `description`/`page_title` errors and others.

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

## Known Issues Without Solution (4 the moment)
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

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/valkiriaaquatica/provider-proxmox-bpg/issues).
