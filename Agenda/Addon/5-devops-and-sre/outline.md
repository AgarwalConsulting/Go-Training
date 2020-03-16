# DevOps & SRE - Outline

## Kubernetes

- Brief overview of Kubernetes
  ![Kubernetes](https://upload.wikimedia.org/wikipedia/commons/b/be/Kubernetes.png)

> Kubernetes offers a declarative API for creating desired state. Controllers incrementally move the current state toward that desired state. This process is often referred to as reconciliation.

- High-level Concepts
  - [kube-controller-manager](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-controller-manager/)
  - [Controller](https://kubernetes.io/docs/concepts/architecture/controller/)
  - [CustomResourceDefinition](https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/)
  - [Operator Pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)

- [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/) vs [Operators](https://coreos.com/operators/)

> A kubernetes operator is the name of a pattern that consists of a kubernetes controller that adds new objects to the Kubernetes API, in order to configure and manage an application, such as Prometheus or etcd.

- Writing your own operator using [Operator Framework](https://github.com/operator-framework/getting-started)

---

## Pre-requisites

### Participant background

- Familiarity with Docker
- Familiarity with Go
- Used Kubernetes in a Production environment
- Nice to have: Managed db & servers

### Machine

- Linux
  - Or Docker on Mac/Windows
- Access to a k8s cluster
  - Recommended: Local setup with [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
- Minimum Specs
  - 8 GB RAM
  - 4-core i5 or higher

## Code

- [K8s Sample Controllers](https://github.com/kubernetes/sample-controller)

## References

- [Extending your Kubernetes Cluster](https://kubernetes.io/docs/concepts/extend-kubernetes/extend-cluster/)
- [Controllers vs Operators](https://stackoverflow.com/questions/47848258/kubernetes-controller-vs-kubernetes-operator)
