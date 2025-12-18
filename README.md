# Stripe Provider

`provider-stripe` is a [Crossplane](https://crossplane.io/) provider built with [Upjet](https://github.com/crossplane/upjet). It exposes XRM-conformant managed resources for the Stripe API.

## 🚀 Release Automation

This repository automatically publishes a new release when the upstream Terraform provider releases a new version:
👉 **Original Terraform Provider:** [lukasaron/terraform-provider-stripe](https://github.com/lukasaron/terraform-provider-stripe)

The automation pipeline works as follows:

1. **Version Detection** – Handled by our Renovate configuration, which tracks new upstream releases.
2. **Version Preparation** – When Renovate opens a PR, our GitHub Actions workflow prepares the next version/tag.
3. **Automated Publishing** – Once the PR is merged, Release Please generates the changelog and publishes the release.

Thanks to this pipeline, the provider stays aligned with upstream without manual intervention.

## Installation

Make sure Crossplane is installed in your cluster.

- Using `up` (CLI):
  ```bash
  up ctp provider install xpkg.upbound.io/valkiriaaquaticamendi/provider-stripe:v0.1.0
  ```
- Declarative:
  ```bash
  cat <<EOF | kubectl apply -f -
  apiVersion: pkg.crossplane.io/v1
  kind: Provider
  metadata:
    name: valkiriaaquaticamendi-provider-stripe
  spec:
    package: xpkg.upbound.io/valkiriaaquaticamendi/provider-stripe:v0.1.0
  EOF
  ```
  or
  ```bash
  kubectl apply -f examples/install.yaml
  ```

Create the secret with your Stripe API key:
```bash
vi examples/providerconfig/secret.yaml.tmpl
kubectl apply -f examples/providerconfig/secret.yaml.tmpl
```

Create the ProviderConfig using that secret:
```bash
kubectl apply -f examples/providerconfig/providerconfig.yaml
```

In `examples/` and `examples-generated/` you will find reference manifests for both cluster-scoped and namespaced resources. The `test/` folder contains quick validation samples (e.g., `coupon`, `file`, `meter`, `portal_configuration`, `product`, `product_feature`, `promotion_code`, `shipping_rate`, `tax_rate`, `webhook_endpoint`, etc.).

## Development

- Generate code:
  ```bash
  make generate
  ```
- Install CRDs into the cluster:
  ```bash
  kubectl apply -f package/crds/
  ```
- Run the controller against a cluster with Crossplane:
  ```bash
  make run
  ```
- Tests:
  ```bash
  make test
  ```
- Build image:
  ```bash
  make build
  ```

## Report a Bug

For bugs, improvements, or new feature requests, please open an [issue](https://github.com/valkiriaaquatica/provider-stripe/issues).
