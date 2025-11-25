#  What is Kubebrowser ?

Kubebrowser is a self-service platform that enables users to obtain their own Kubeconfig credentials using OIDC authentication.

![Kubebrowser UI](/0.8/kubebrowser-ui.webp)

<div class="tip custom-block" style="padding-top: 8px">

Just want to try it out? Skip to the [Quickstart](./getting-started).

</div>

## Use case

Whenever you have *self-hosted* or *managed* Kubernetes clusters on cloud provider with weak IAM (e.g. [OVH](https://www.ovh.com/)) it becomes hard to control who has access to what.

Kubebrowser let you leverage the [Kubernetes API server OIDC configuration](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#openid-connect-tokens) and solves the *how* to give people (developers, testers, admins, etc.) individual Kubeconfig.

Using Kubebrowser, administrators manage a list of clusters and, if needed, a whitelist of *who* should have access to such cluster. Note that Kubebrowser only offers a way to **distribute** Kubeconfigs (used to *authenticate* to clusters). *Authorization* still relies on native Kubernetes RBAC.

Other ways to manage such *authentication* could be to install [permission manager](https://github.com/sighupio/permission-manager) or [pinniped](https://pinniped.dev/), however, for mid-size businesses, the first one can be limiting, and the latter could be challenging to operate.

Another alternative is the [kubelogin](https://github.com/int128/kubelogin) plugin which enables OIDC authentication directly via `kubectl`. However, using `kubelogin` requires each user to install and configure the plugin on their local machine. In contrast, Kubebrowser does not require any additional plugins for users and simplifies Kubeconfig distribution, supporting access control through easy-to-manage whitelists.
