---
# https://vitepress.dev/reference/default-theme-home-page
---

# Contribute

## Project architecture

```
.
├── chart           # Where lies the Helm chart
├── docs            # This is what you are reading
├── k8s             # Raw manifests for dev purposes
├── server          # API and Kubeconfig Controller of the Kubebrowser
└── ui              # The UI
```

## Local setup

### Prerequisites

You’ll need:

- A local Kubernetes cluster running, we recommend [minikube](https://minikube.sigs.k8s.io/docs/).
- [kubectl](https://kubernetes.io/docs/reference/kubectl/) and [helm](https://helm.sh/docs/intro/install/) CLI.
- An OpenID provider.
  - Could be any provider, like a local [Keycloak](https://www.keycloak.org/securing-apps/oidc-layers) or even [Google](https://developers.google.com/identity/openid-connect/openid-connect).
- Golang, NodeJS with pnpm and Skaffold.
  - *Optional* We are using [Devbox](https://www.jetify.com/docs/devbox/) to easily create isolated shells for development. Simply run `devbox shell` after installing `devbox`.
  - Otherwise you can manually manually the right versions listed in the `devbox.json` file.
- *Optional* [pre-commit](https://pre-commit.com/) as its hooks are checked by the pipeline.
- *Optional* [direnv](https://direnv.net/).

### Run Kubebrowser locally

Set up your environment
1. Create a new OIDC application following your provider’s documentation (you can set the redirect address to be `http://localhost:8080`).
1. Copy the `.envrc_example` to a new file named `.envrc`.
1. Load your `client_id`, `client_secret` and `issuer_url` from your OIDC application within your `.envrc`.
1. Source your `.envrc` or run `direnv allow` if you have installed direnv.
1. Run `skaffold dev`.
1. Access http://localhost:8080.
