# Kubernetes

Concepts: Nodes, Pod, Service [iptables]

Controllers: Deployment => ReplicaSet (replicas of Pods), StatefulSet, DaemonSet, CronJob => Job (Pods), ...

External Controller: Ingress Controller (nginx, istio, ...) {Self: ClusterIP => DNS; Cloud: ELB/ALB => DNS, ...}; Certificate Manager (Let's encrypt) [Custom Resource Definition], ...

Custom Controller: Prometheus Operator

---

Helm Chart vs Operator (CRDS => backups, ...)?

- installing, updating, deletion

---

Group Version Kind

apps  v1  Deployment

batch.k8s.algogrit.com v1 CronJob

---

User {kubectl} => Kubernetes Cluster {apiServer {webhooks}, controller manager (Deployment, Custom Controller)}

apply {create, update}

create
update
delete

kubectl edit => kubectl apply

---

cronJob => [
  Job1 at 5:55 PM (Running)
  Job2 at 5:50 PM (Failed)
  Job3 at 5:45 PM (Running)
  Job4 at 5:40 PM (Succeeded)
  Job5 at 5:40 PM (Succeeded)
  ...
]

cronJob => []
