# Kubeconfig Specification

```yaml
apiVersion: kubebrowser.io/v1alpha1
kind: Kubeconfig
metadata:
  name: cluster-name # Kubernetes resource name
spec:
  name: "Friendly name" # Name on the UI
  whitelist: # Optional. If no whitelist is set, this Kubeconfig will be displayed on the UI regardless of the user connecting.
    groups: # Groups are fetched from the claim `groups` of the OIDC ID Token.
    - administrators
    users: # User is fetched from the claim `email` of the OIDC ID Token.
    - alan.turing@example.com
    - ada.lovelace@example.com
  kubeconfig: # Content of a regular kubeconfig (usually given at cluster creation). You should remove `users` and `preferences`.
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
