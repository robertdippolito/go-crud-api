# Simple CRUD API written in GoLang 
A simple API used to simulate various real world scenarios using simple and easy to understand logic.

## Table of Contents
- [About](#-about)
- [Getting Started](#-getting-started)
- [API Endpoints](#-api)
- [_infra Directory](#-infra)
- [Feedback and Contributions](#-feedback-and-contributions)
- [License](#-license)
- [Contacts](#%EF%B8%8F-contacts)

## 🚀 About
This repository is part of a multi-repository experiment designed to test Kubernetes clusters under load under different scenarios. As much as possible, this api was designed to be easily spun up and spun down via Github Actions to optimize cost and useability. Further testing and development on this repository is encouraged, if you notice any issues feel free to open a PR.

Videos related to this repository can be found here:

<a href="youtube.com">![image](https://img.shields.io/badge/YouTube-FF0000?style=for-the-badge&logo=youtube&logoColor=white) </a>

## 🔥 Getting Started
This API connects to MongoDB on start up. In order to run this application locally you will need to have a MongoDB Atlas cluster setup. To do this, see the instructions [here](todo) on how to set up a **free** MongoDB Atlas cluster. In additon to having a MongoDB cluster you also need to have [Go installed](https://go.dev/learn/) on your computer. Once your cluster is up and running and you will need to execute the following steps:

### Local Development
1. Update your ```.env``` (todo add .env.example file) with your Mongo DB credentials
2. Allowlist your IP address in the MongoDB console (ideally, use a set duration so this rule will auto expire when you're done)
3. Run the following command:
   
       go run .
   
5. Verify your endpoints are accessible via ```localhost:8000/{endpoint}```

### Remote Development (i.e. API deployed to EKS)
External-secrets are used to store your credentials to be used for remote development. To do this, you will need to create a secret store in AWS Secret's manager following the instructions provided here [here](todo). Once setup, the secrets will be available to your cluster via the ClusterSecretStore. To deploy this API execute the following steps:
1. Ensure your EKS cluster as outlined [here](https://github.com/robertdippolito/eks-infrastructure-iac) is deployed and healthy.
2. Install the following applications onto your cluster:
    - ArgoCD
    - Prometheus / Grafana
    - External-Secrets
4. Navigate to the Actions tab and run the [Deploy Helm Chart](https://github.com/robertdippolito/go-crud-api/actions/workflows/deploy-helm.yaml) action.
5. Once deployed, update your Kube context and run:

       kubectl get pods -n api 
   verify your api pods are running in the cluster.

6. Verify your endpoints are accessible via:
```https://{your-domain}/{endpoint} ```
   
## 🗒️ API Endpoints
This application has 2 simple endpoints ```/users``` and ```/compute```. The default route ```/``` returns a simple string. 

## 🕸️ _infra Directory
The _infra directory contains the Kubernetes objects used to run the API in the EKS Cluster. The diagram below highlights each of the Kubernetes objects running in the cluster and the text below explains each:
TODO: insert diagram
### Ingress Object
### Deplyment Object
### Service Object
### Horizontal Pod Autoscaler (HPA) Object
### External Secrets Object

## 📃 License 
As much as possible I'd like to encourage participation and collaboration. This code is open-source software licensed under [Apache 2.0 License](https://github.com/gowebly/gowebly/blob/main/LICENSE), created and supported by [Robert D'Ippolito](https://robertdippolito.me) for people and robots.

<a href="https://github.com/gowebly/gowebly/blob/main/LICENSE">![License](https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none)</a>

## 🗨️ Contacts
The best way to get in touch is via a Github issue or PR. You can also get in touch via my blog here: [https://robertdippolito.me](https://robertdippolito.me/)
