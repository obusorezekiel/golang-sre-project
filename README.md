Here's a comprehensive `README.md` file for your Golang REST API project:

---

# Golang REST API Project

## Project Overview

This project demonstrates a complete workflow for building, containerizing, and deploying a Golang REST API with its dependent services. It covers everything from local development setup to production-grade deployments using Kubernetes, Helm Charts, and ArgoCD for continuous deployment.

### Key Achievements:
1. **Build a Simple REST API Webserver**
2. **Containerize REST API**
3. **Setup One-click Local Development Setup**
4. **Establish a CI Pipeline**
5. **Deploy REST API and Dependent Services on Bare Metal**
6. **Setup a Kubernetes Cluster**
7. **Deploy REST API and Dependent Services in Kubernetes**
8. **Deploy REST API and Dependent Services Using Helm Charts**
9. **Setup One-click Deployments Using ArgoCD**

---

## Table of Contents

1. [Project Structure](#project-structure)
2. [Tech Stack](#tech-stack)
3. [Setup Guide](#setup-guide)
4. [CI Pipeline](#ci-pipeline)
5. [Deployment](#deployment)
   - [Bare Metal](#bare-metal-deployment)
   - [Kubernetes](#kubernetes-deployment)
   - [Helm Deployment](#helm-deployment)
   - [ArgoCD Deployment](#argocd-deployment)
6. [Testing](#testing)
7. [Contributing](#contributing)

---

## Project Structure

```
├── api/
│   ├── controllers/
│   ├── models/
│   ├── routers/
├── config/
│   ├── database.go
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── kubernetes/
│   ├── manifests/
│   │   ├── api-deployment.yaml
│   │   ├── db-deployment.yaml
│   │   ├── service.yaml
├── helm/
│   ├── charts/
│   ├── values.yaml
├── argo-cd/
│   ├── applications/
│   ├── secrets/
├── .github/
│   ├── workflows/
│   │   ├── ci-pipeline.yaml
└── README.md
```

---

## Tech Stack

- **Language:** Golang
- **Frameworks/Tools:**
  - REST API: Gin
  - Database: PostgreSQL
  - Containerization: Docker
  - Kubernetes Orchestration: K8s
  - CI/CD: GitHub Actions
  - Continuous Delivery: ArgoCD
  - Helm Charts for K8s deployments
- **Infrastructure:**
  - Bare Metal Server
  - Minikube for local K8s testing
  - Prometheus & Grafana for monitoring

---

## Setup Guide

### Prerequisites

- Golang 1.22+
- Docker & Docker Compose
- Kubernetes (Minikube or a managed cluster)
- Helm
- ArgoCD
- GitHub repository for CI

### Local Development

1. **Clone the repository:**

   ```bash
   git clone https://github.com/obusorezekiel/golang-sre-project
   cd golang-sre-project
   ```

2. **Environment Setup:**
   Copy the `.env.example` file to `.env` and update the necessary values:

   ```bash
   cp .env.example .env
   ```

3. **Start local development using Docker Compose:**

   ```bash
   docker-compose up --build
   ```

4. **Access the API:**
   The API should be running on `http://localhost:8888/api/v1/students`. You can use tools like Postman or Curl to test the endpoints.

---

## CI Pipeline

The CI pipeline is configured to run the following stages:

1. **Linting and Formatting** using `golangci-lint`
2. **Unit Tests** using `go test`
3. **Build Docker Image** and push to Docker Hub
4. **Run Database Migrations**
5. **Deployment on Kubernetes**

The pipeline is defined in the `.github/workflows/ci.yaml` file and is triggered automatically on commits to the `main` branch.

---

## Deployment

### Bare Metal Deployment

You can deploy the API and its dependent services directly to bare metal:

1. SSH into the server:
   ```bash
   ssh user@server-ip
   ```

2. Use Docker Compose to run the services:
   ```bash
   docker-compose -f docker-compose.prod.yml up -d
   ```

### Kubernetes Deployment

1. **Setup Kubernetes Cluster:** Use Minikube for local testing or your preferred Kubernetes service for production.

   ```bash
   minikube start
   ```

2. **Apply Kubernetes manifests:**

   ```bash
   kubectl apply -f kubernetes/manifests/
   ```

3. **Verify the Deployment:**

   ```bash
   kubectl get pods
   kubectl get services
   ```

4. The API will be accessible via the Kubernetes service on the exposed LoadBalancer.

### Helm Deployment

1. **Deploy via Helm:**

   Install the Helm chart from the `helm/` directory:

   ```bash
   helm install rest-api helm/my-go/
   ```

2. **Monitor the Deployment:**

   ```bash
   kubectl get pods
   ```

3. **Custom Values:**

   You can override the default configuration by editing `values.yaml` or passing custom values:

   ```bash
   helm install rest-api helm/my0go/ -f custom-values.yaml
   ```

### ArgoCD Deployment

1. **Setup ArgoCD:** Follow the instructions [here](https://argo-cd.readthedocs.io/en/stable/getting_started/) to install ArgoCD on your Kubernetes cluster.

2. **Deploy the Application:**

   Create an ArgoCD application manifest in the `kubernetes/argocd.yaml/` file, then apply it:

   ```bash
   kubectl apply -f kubernetes/arogcd/
   ```

3. **Monitor Application Sync:**

   ```bash
   argocd app get golang-are-project
   ```

---

## Testing

1. **Run Unit Tests:**

   To run tests locally:

   ```bash
   go test ./...
   ```

---

## Contributing

1. Fork the repository.
2. Create a feature branch: `git checkout -b my-feature`
3. Commit your changes: `git commit -m 'Add some feature'`
4. Push to the branch: `git push origin my-feature`
5. Open a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

## Contact

For any inquiries or support, feel free to reach out to me via [ezekiel.umesi@gmail.com](mailto:ezekiel.umesi@gmail.com) or visit my [LinkedIn](https://linkedin.com/in/ezekiel-umesi).

--- 