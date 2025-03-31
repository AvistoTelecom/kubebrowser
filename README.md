# Kubebrowser - Your Kubernetes catalog with OIDC

## Overview
Kubebrowser simplifies Kubernetes access management by providing a centralized catalog where users can easily generate their own Kubeconfigs using OpenID Connect (OIDC). Designed for organizations using multiple cloud providers or self-managed Kubernetes clusters, Kubebrowser reduces the complexity of managing permissions and Kubeconfig distribution.

## Features
Managing access across multiple Kubernetes clusters can be challenging, especially when dealing with different teams, roles, and providers. Kubebrowser streamlines this process with:

🚀 Self-Service Access: Users can generate their own Kubeconfigs without manual intervention.

🔐 OIDC Integration: Secure authentication with identity providers.

🌍 Multi-Cluster Support: Manage access across multiple Kubernetes clusters easily.

✅ Whitelist Management: Allow administrators to set a whitelist on each Kubeconfig.

## Getting started
To deploy Kubebrowser and start using it:

1. **Install Kubebrowser** using Helm or Kubernetes manifests.
2. **Configure OIDC authentication** with your identity provider.
3. **Add Kubernetes clusters** to the catalog.
4. **Grant appropriate permissions** to teams and users.
