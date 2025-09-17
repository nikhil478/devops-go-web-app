# devops-go-web-app

A complete template to build, containerize, and deploy a **Go web application** with a modern DevOps stack on **DigitalOcean Kubernetes (DOKS)**.

## Features Covered

### Containerization
- **Dockerfile** to build and run the Go application.

### Kubernetes
- **Deployment**, **Service**, and **Ingress** manifests.
- **Ingress Controller** integrated with a **DigitalOcean Load Balancer** and DNS.

### Continuous Integration (CI)
Implemented with **GitHub Actions**:
- Build and test the Go application.
- Static code analysis.
- Build and push Docker images.
- Update Helm chart automatically.

### Continuous Delivery (CD)
Implemented with **Argo CD**:
- Pulls the Helm chart from GitHub.
- Deploys the application to Kubernetes.

### Helm
- Environment-based Helm charts for:
  - **dev**
  - **qa**
  - **prod**

### DigitalOcean Kubernetes (DOKS)
- Managed Kubernetes cluster hosting the entire setup.

---

## Quick Start

### Build and Run Locally
```bash
go build -o main .
./main