
# Contribute

## Repository architecture

```
.
├── dev                       # 🧑‍💻 Raw manifests and values for dev purposes
├── docs                      # This is what you are reading
├── scripts                   # 🧑‍💻 Useful scripts
├── server                    # API and Kubeconfig controller
├── ui                        # UI
├── .envrc                    # 🧑‍💻 Loads environment variables and devbox, used by `direnv`
├── devbox.json               # 🧑‍💻 Devbox config
├── pnpm-workspace.yaml       # 🧑‍💻 Both UI and the docs are VueJS projects
├── .pre-commit-config.yaml   # 🧑‍💻 Precommit config
└── Tiltfile                  # 🧑‍💻 Tilt config
```

> Things noted with 🧑‍💻 are purely for development purposes.

## Local setup

Kubebrowser uses some tool for development purposes. Some are optional for minor changes, though I recommend to grab them all for an optimal developer experience.

- [Devbox](https://www.jetify.com/docs/devbox/) creates isolated shells for development. Simply run `devbox shell` after installing `devbox`.
  Devbox also installs every binary you need (golang, nodejs) at the same version for everyone working on the project.
- *Only for backend* [Tilt](https://tilt.dev/) eases development on Kubernetes.
- *Optional* [pre-commit](https://pre-commit.com/) asserts file formating upon commit.
- *Optional* [direnv](https://direnv.net/) loads dynamically content of `.envrc`.

## Contribute on the UI

The UI is a VueJS project. It can run and simulate API responses without the server or Kubernetes cluster to be running.

###  Prerequisites

::: info
Either use Devbox as [stated above](#local-setup) or check in `devbox.json` for the version you should use for.
:::

- NodeJS
- pnpm

### Make the UI run

1. At the project root, run `pnpm install`.
1. `pnpm dev:ui`.
1. Access http://localhost:5173/ 🎉


## Contribute on the server

### Prerequisites

::: info
Either use Devbox as [stated above](#local-setup) or check in `devbox.json` for the version you should use for.
:::

- A local Kubernetes cluster running, we recommend [minikube](https://minikube.sigs.k8s.io/docs/). 🚨 This step should be done regardless of the use of Devbox.
- [kubectl](https://kubernetes.io/docs/reference/kubectl/) and [helm](https://helm.sh/docs/intro/install/) CLI.
- [Golang](https://go.dev/).
- [Tilt](https://tilt.dev/) and [ko](https://ko.build/).
- UI [prerequisites](#prerequisites).
- An OpenID provider with an application set up. It could be any OIDC compatible provider, like a local [Keycloak](https://www.keycloak.org/securing-apps/oidc-layers) or even [Google](https://developers.google.com/identity/openid-connect/openid-connect).

### Run Kubebrowser locally

Set up your environment
1. Create a new OIDC application following your provider's documentation (you can set the redirect address to `http://localhost:8080/auth/callback`).
1. Copy the `./dev/values.example.yaml` to a new file named `./dev/values.yaml`.
1. Update `client_id`, `client_secret` and `issuer_url` in the newly created `./dev/values.yaml` (from your OIDC application in step 1).
1. Make sure your cluster is running and your context is correct. Run `minikube start` if you're using minikube.
1. Add Helm repository, it should be named "avisto" for Tilt to work correctly `helm repo add avisto https://avistotelecom.github.io/charts/`.
1. Run `tilt up`.
1. Access http://localhost:8080 🎉

::: info
Frontend is compiled once at startup. If you want to work heavily on the UI, please follow the [Contribute on the UI](#contribute-on-the-ui) instead.
:::

## Contribute on the documentation

The documentation is like a VueJS project using [Vitepress](https://vitepress.dev/).

###  Prerequisites

Same as UI [prerequisites](#prerequisites).

### Make the UI run

1. At the project root, run `pnpm install`.
1. `pnpm dev:docs`.
1. Access http://localhost:5173/ 🎉
