---
outline: deep
---

# Getting started
## Installation
### Prerequisites

- Kubectl.
- Helm version 3 or higher.
- Kubernetes cluster, we recommend [minikube](https://minikube.sigs.k8s.io/docs/) for development or test purposes.
  - API server configured to accept OpenID Connect Authentification. For minikube [visit the docs](https://minikube.sigs.k8s.io/docs/tutorials/openid_connect_auth/).
- OpenID Connect provider, such as Microsoft, Google or Keycloak.

### Prepare your OpenID application

Follow the documentation of your OpenID Connect provider in order to register an application. You should retrieve/generate:

- Client ID
- Client Secret
- Issuer URL

These values are mandatory in order to be able to install Kubebrowser in your cluster.

You should also authorize a redirect URI, for test purposes, you can set `http://localhost:8080`.

### Install Kubebrowser in your cluster

First, create a `values.yaml` file.

```yaml
server:
  oidc:
    clientID: <your-client-id>
    clientSecret: <your-client-secret>
    issuerURL: <your-issuer-url>
```
Second, run the command:

```sh
helm repo add avisto https://avistotelecom.github.io/charts/
helm install kubebrowser avisto/kubebrowser -f values.yaml
```

::: info
The chart will install a CRD (CustomResourceDefinition) named `Kubeconfig`.
:::

## Add a Kubeconfig

Because Kubebrowser declares a new resource of kind `Kubeconfig`, adding a cluster to your catalog is as easy as creating a new `Kubeconfig` ressource using `kubectl`.

First, grab your current Kubeconfig.

```sh
kubectl config view --minify --raw > kubeconfig.yaml
```

Open `kubeconfig.yaml` and delete `preferences` and `users`.

```yaml
apiVersion: v1
kind: Config
clusters:
- cluster:
    certificate-authority-data: <base64-encoded>
    server: https://127.0.0.1:32771
  name: cluster
contexts:
- context:
    cluster: cluster
    user: placeholder
  name: context
current-context: context
preferences: {} # [!code --]
users: []       # [!code --]
```

Then embed everything in the `Kubeconfig` CRD.

```yaml
apiVersion: kubebrowser.io/v1alpha1 # [!code ++]
kind: Kubeconfig                    # [!code ++]
metadata:                           # [!code ++]
  name: cluster-name                # [!code ++]
spec:                               # [!code ++]
  name: "Friendly name"             # [!code ++]
  kubeconfig:                       # [!code ++]
    apiVersion: v1
    kind: Config
    clusters:
    - cluster:
        certificate-authority-data: <base64-encoded>
        server: https://127.0.0.1:32771
      name: cluster
    contexts:
    - context:
        cluster: cluster
        user: placeholder
      name: context
    current-context: context
```

Finally, create your Kubeconfig in your cluster.

```sh
kubectl apply -f kubeconfig.yaml
```

## Grab your personnal Kubeconfig

Port forward the application.

```sh
kubectl port-forward services/kubebrowser-server 8080
```

Access your Kubebrowser: http://localhost:8080.

You should be able to copy your personal Kubeconfig and save it locally, or paste it in any tool like FreeLens or Headlamp.

For the rest of the Getting Started, paste the content in a file named `config`.

## Use your fresh Kubeconfig

::: warning
By default, you are authenticated but have no authorization to query any information from the Kubernetes API Server. In the following we'll create some basic permissions to complete the Getting Started.

If you want to know more, read [this documentation](https://kubernetes.io/docs/reference/access-authn-authz/rbac) about RBAC on Kubernetes.
:::

As before, create the following resources in order to grant a permission to your user.
```sh
kubectl apply -f cr.yaml
```
```yaml
# cr.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: get-namespaces-binding
subjects:
- kind: User
  name: your-username    # Replace with your actual username [!code highlight]
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: get-namespaces
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: get-namespaces
rules:
- apiGroups: [""]
  resources: ["namespaces"]
  verbs: ["get"]
```

Finally, use the Kubeconfig fetched in the [previous section](#grab-your-personnal-kubeconfig) !

```sh
kubectl get namespaces --kubeconfig config
```
