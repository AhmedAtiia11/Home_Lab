# What is Istio?

Istio is an open-source service mesh that manages communication between microservices.
It works by injecting sidecar proxies (Envoy) into each pod, allowing it to control:
- Traffic routing
- Security policies
- Observability

This enables enforcing infrastructure-level concerns without modifying application code.

## Why Use Istio for Security?

In a default Kubernetes environment:

- Service-to-service communication is not encrypted by default  
- There is no strong identity verification between services  

This creates security risks such as:

- Unauthorized access between services
- Traffic interception
- Service impersonation

## Istio Solution

Istio addresses these issues by:

- Enforcing mutual TLS (mTLS) between services
- Assigning unique identities to each workload
- Ensuring that only authenticated and encrypted traffic is allowed


## What is mTLS?

mTLS (Mutual TLS) is a security protocol where:

- Both the client and server authenticate each other
- All communication is encrypted

With Istio:

- Certificates are automatically generated and rotated
- Encryption is enforced transparently
- No application changes are required

## mTLS Modes in Istio

Istio supports different modes for controlling mTLS behavior using PeerAuthentication.

1. PERMISSIVE Mode
- Encrypted (mTLS) traffic
- Plain-text traffic

2. STRICT Mode (Used in This Project)
- Only mTLS-encrypted traffic is allowed
- Plain-text communication is rejected

3. DISABLE Mode
- mTLS is completely disabled
- All traffic is plain-text

Why STRICT Mode?

- This project uses **STRICT mTLS mode** to enforce secure communication between services.

Benefits:
- End-to-end encryption between all pods
- Verified service identity (no impersonation)
- Blocks unauthorized or insecure traffic
- Enables a Zero Trust architecture

With STRICT mode enabled:

Frontend → Backend communication is:
- Encrypted
- Authenticated
- Any unknown or non-mTLS service:
  - Cannot connect
  - Is automatically rejected


## Role of Istio Ingress (Brief)

The Istio Ingress Gateway acts as the secure entry point for external traffic.

Controls how requests enter the cluster
Can enforce TLS and routing rules
Integrates with Istio security policies

## Installation & Configuration Steps
```
# 1. Install Istio with cert rotation settings
istioctl install --set profile=demo \
  --set values.pilot.env.CITADEL_SELF_SIGNED_CA_CERT_TTL=8760h \
  --set values.pilot.env.CITADEL_SELF_SIGNED_ROOT_CERT_CHECK_INTERVAL=1h \
  --set values.pilot.env.CITADEL_SELF_SIGNED_ROOT_CERT_GRACE_PERIOD_PERCENTILE=20

kubectl delete secret istio-ca-secret -n istio-system
kubectl rollout restart deployment istiod -n istio-system
kubectl rollout status deployment istiod -n istio-system

# 2. Verify Istio control plane is running
kubectl get pods -n istio-system

# 3. Move ingressgateway off port 80/443 to avoid Traefik conflict
kubectl edit daemonsets.apps -n kube-system svclb-istio-ingressgateway-e5a375fc
# change 80 → 8080 and 443 → 8443

kubectl edit svc -n istio-system istio-ingressgateway
# change 80 → 8080 and 443 → 8443

# 4. Enable sidecar injection on your namespace
kubectl label namespace default istio-injection=enabled

# 5. Apply your app manifests (deployments, services, gateway, virtualservices, peerauth)
kubectl apply -f Kubernetes/

# 6. Restart pods so sidecars get injected
kubectl rollout restart deployment golang-app react-app

# 7. Verify pods show 2/2
kubectl get pods
# golang-app-xxx   2/2   Running  ✓
# react-app-xxx    2/2   Running  ✓

# 8. Verify all proxies are synced
istioctl proxy-status

# 9. Verify routes are on 8080
istioctl proxy-config routes \
  $(kubectl get pod -n istio-system -l app=istio-ingressgateway \
    -o jsonpath='{.items[0].metadata.name}') -n istio-system
# Should show http.8080 for backend.local and frontend.local

# 10. Test it works
curl -H "Host: backend.local" http://192.168.1.15:8080/health
# Expected: {"msg":"healthy"}

# 11. Install observability addons (Grafana, Kiali, Prometheus, Jaeger)
kubectl apply -f samples/addons/

# 12. Verify addons are running
kubectl get pods -n istio-system
```