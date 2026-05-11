
# ITEMS API - PRODUCTION MODEL #

## 1 - PROJECT OVERVIEW

This application is developed in Golang, serving as a reference model for high-performance APIs.

**1.1 - ITEMS API**

- Efficient item listing API.
- Structured Logs implementation.
- Optimized Dockerization.
- Kubernetes Manifests for orchestration.
- **Local MySQL Integration** (Simplified from RDS).

**1.2 - Infrastructure**
- Infrastructure as Code (Terraform) for AWS EKS (optional).
- Local development with Kubernetes (Kind).
- Task automation with Taskfile.

## 2 - LOCAL DEVELOPMENT

### Prerequisites
- Go 1.20+
- Docker
- Kind
- Task (go-task)

### Setup and Execution
We use a `Taskfile` to simplify local development.

1. **Start Local Cluster:**
   ```bash
   task cluster-up
   ```

2. **Deploy Application:**
   This will deploy MySQL and the API for different regions (Argentina, Brazil, Colombia, Mexico).
   ```bash
   task deploy
   ```

3. **Verify Installation:**
   ```bash
   task verify
   ```

4. **Teardown:**
   ```bash
   task cluster-down
   ```

## 3 - MONITORING AND OBSERVABILITY

Integrated with NewRelic for monitoring (configured via environment variables).

## CI/CD - GITHUB ACTIONS

Automated pipeline for Build and Publish of the Docker image.

##### CONTACTS
- Linkedin: (https://www.linkedin.com/in/marcosouzatech/)
- Youtube: https://www.youtube.com/@maistalkmenosshow
- Email: marcos.souza@luvtech.com.br
