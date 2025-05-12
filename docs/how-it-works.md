# How it works

## KubeBrowser is an Operator

KubeBrowser is an [Operator](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/), hence it has been designed to be hosted on Kubernetes. KubeBrowser watches its CustomResourceDefinition and always renders an up-to-date view on the UI.

## OpenID Connect

KubeBrowser leverages your OIDC provider to:

1. Show only the Kubeconfigs your user has been whitelisted to see.
1. Fill the Kubeconfig `users` section so it is ready to use.

::: warning
Currently, KubeBrowser uses a unique OIDC configuration for all the Kubeconfigs, this could change in the future depending on the community feedback and needs.
:::

```d2
vars: {
  d2-config: {
    theme-id: 105
    layout-engine: tala
  }
}

Cluster 1 {
  API Server
  KubeBrowser
  CRD.Kubeconfigs {
    Cluster 1
    Cluster 2
  }
}

oidc: OIDC Provider

Cluster 1.KubeBrowser -> Cluster 1.CRD: watch

Cluster 2.API Server -> oidc: authenticate requests
Cluster 1.API Server -> oidc: authenticate requests
Cluster 1.KubeBrowser -> oidc: authenticate users
```
