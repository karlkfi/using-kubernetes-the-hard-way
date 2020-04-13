# Using Kubernetes The Hard Way

This tutorial walks you through **using** Kubernetes the hard way.

This guide is *not* for people looking to learn how to deploy and operate Kubernetes clusters. If that's you, then check out [Kubernetes The Hard Way](https://github.com/kelseyhightower/kubernetes-the-hard-way).

Usage Kubernetes The Hard Way is optimized for learning, which means taking the long route to ensure you understand each task required to deploy and operate a web app on Kubernetes.

## Prerequisites

- Clone of this GitHub repository
- [Kubernetes](https://kubernetes.io/) cluster
- [Kubernetes user account](https://kubernetes.io/docs/reference/access-authn-authz/authentication/)
- [Vault](https://www.vaultproject.io/) cluster
- [Vault user account](https://learn.hashicorp.com/vault/getting-started/authentication)

## Goals

- Go
  - Write a web app in Go
  - **TODO**: Build a Go binary
  - **TODO**: Go Modules for Dependency Management
  - **TODO**: Unit Tests
  - Health check endpoint (for Kubernetes Readiness Check)
  - Graceful Termination with Connection Draining on SIGTERM
- Linting
  - **TODO**: YAML linting
  - **TODO**: JSON linting
  - **TODO**: Shell linting
  - **TODO**: Goimports (incl go fmt) linting
- Container Image Building
  - **TODO**: Docker Hub Image GitOps using GitHub Actions
  - **TODO**: Image builds for master and version tags
  - **TODO**: Image builds for branches (for end-to-end testing)
  - **TODO**: Image builds for CI (for test tool installation)
  - Multi-stage builds for smaller images
- GitHub Actions
  - GitHub Actions config GitOps
  - **TODO**: Trigger on PRs/branches
  - **TODO**: Integration tests on Kubernetes for master and branches
  - **TODO**: Deployment Promotion Dev -> Staging -> Prod
  - **TODO**: Container Image Garbage Collection (new image on every push)
  - **TODO**: Automated Canary Deployment Testing
  - **TODO**: Terraform Module GitOps
- Vault
  - **TODO**: App Secrets in Vault (Kubernetes read, Team write)
- Kubernetes
  - Zero Downtime Deployments
  - Ingress Load Balancing
  - Pod Disruption Budget
  - Deployment Rollout Strategy
  - Liveness & Readiness Checks
  - **TODO**: App-specific Kubernetes Service Account
  - **TODO**: Daytona sidecar App Secrets injection
  - **TODO**: Resource requests & limits (cpu, memory, ephemeral-storage)
  - **TODO**: Ephemeral Storage EmptyDir Volume
  - **TODO**: Version Rollbacks
  - **TODO**: Horizontal Autoscaling
  - **TODO**: Vertical Autoscaling
- Templating ([envsubst](https://www.gnu.org/software/gettext/manual/html_node/envsubst-Invocation.html))
- DNS
  - Namecheap DNS GitOps
- **TODO**: Blackbox Endpoint Monitoring (ex: [CULA](https://cula.io/))
- **TODO**: Metrics Dashboards
- **TODO**: Monitoring & Alerting
- **TODO**: Distributed Tracing

## Setup

- Create a [Kubernetes namespace](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)
- Configure [Kubernetes RBAC authorization](https://kubernetes.io/docs/reference/access-authn-authz/rbac/) for your user to manage resources in the Kubernetes namespace
- Configure [Vault authorization](https://learn.hashicorp.com/vault/getting-started/policies) for your user to manage secrets in a secrets workspace (ex: `secret/project/${PROJECT}/*`).
- Configure a [Vault Kubernetes Auth Backend](https://www.vaultproject.io/docs/auth/kubernetes) to allow Kubernetes pods to login to Vault using a Kubernetes Service Account.
