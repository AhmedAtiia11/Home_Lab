# Project Overview

This project is a Kubernetes-based homelab environment designed for hands-on practice and in-depth exploration of cloud-native and DevOps tools. The goal is to build a solid understanding of how modern infrastructure components integrate within a real-world microservices architecture.

## Base Application

The platform is built around a simple connection-testing application that demonstrates communication between services:

Frontend: A web interface built with React that displays the connection status
Backend: A Go-based service responsible for handling requests and validating connectivity

When the backend successfully communicates with the frontend, the application reflects an “online” status.

## Technologies & Tools

The following technologies are currently used in this project:

- Go
- Docker
- Kubernetes
- Istio

## 🔐 Istio
<div align="center">
  <img src="docs/images/Istio.png" alt="RTA Project" width="500"/>
</div>


All service-to-service communication inside the cluster is secured using Istio STRICT mTLS:

- Traffic is encrypted using mutual TLS  
- Each service is authenticated via strong identity  
- Plain-text or unauthorized traffic is rejected  

This ensures a Zero Trust networking model inside Kubernetes.
