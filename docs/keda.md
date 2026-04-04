# What is KEDA?

KEDA (Kubernetes Event-Driven Autoscaling) is an open-source component that enables automatic scaling of applications based on external events.

Unlike traditional scaling, which relies mainly on CPU or memory usage, KEDA allows Kubernetes workloads to scale based on:

- Message queues (e.g., Prometheus, Kafka)
- HTTP requests
- Custom metrics
- External systems

This makes applications more responsive to real-world demand.

## Why Use KEDA?

In a default Kubernetes setup:

Autoscaling is typically based on:
- CPU usage
- Memory usage

This can be limiting because:

- It reacts after load increases
- It doesn’t reflect business-level events
- It may lead to:
- Slow scaling

### How KEDA Improves This

KEDA introduces event-driven scaling, allowing workloads to scale based on real triggers.

For example:

Scale when:
- HTTP traffic increases
- Queue length grows
- Events are received

This leads to:

- Faster response to demand
- Better resource efficiency
- More production-like behavior


## How KEDA Works

KEDA integrates with Kubernetes using two main concepts:

🔹 1. ScaledObject

Defines:

- What workload to scale (Deployment)
- Minimum and maximum replicas
- Scaling triggers

🔹 2. Triggers

Triggers define when scaling happens, such as:

- CPU or memory thresholds
- HTTP request rate
- Queue length
- External metrics


## Why I Used KEDA in This Project

This project uses KEDA to simulate real-world dynamic workloads.

Benefits:
- Automatically scales backend services under load
- Responds quickly to traffic spikes
- Demonstrates production-ready scaling behavior
- When load increases pods scale up automatically
- When load decreases Pods scale down

## Installation & Configuration Steps
 ```
# 1. Install Keda using Helm
   helm repo add kedacore https://kedacore.github.io/charts
   helm repo update
   helm install keda kedacore/keda   --namespace keda   --create-namespace

# 2. apply  scaled object resource
   kubectl apply -f Kubernetes/keda/scaledobject-go.yaml

# 3. Run Stress test script
   chmod +x Kubernetes/keda/stresstest.sh
   ./Kubernetes/keda/stresstest.sh # it sends 500 request to backend

# 4. Monitort the resources 
   kubectl get hpa -w
   kubectl get pods -w   
 ```