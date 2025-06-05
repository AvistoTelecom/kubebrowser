# Expose Kubebrowser

The preferred way to expose Kubebrowser is using an Ingress with a TLS termination.

Here is an example using our [Helm chart](https://artifacthub.io/packages/helm/avisto/kubebrowser).

:::warning
For the TLS to work, something (e.g. cert-manager) should handle the certificate generation.
:::

```yaml
server:
  oidc:
    ...
ingress:
  enabled: true
  hostname: kubebrowser.example.com
  tls: true
```


You can also directly provide your own certificate using a Secret.

```yaml
server:
  oidc:
    ...
ingress:
  enabled: true
  hostname: &hostname kubebrowser.example.com
  extraTls:
    - hosts:
        - *hostname
      secretName: example.com-tls
```
