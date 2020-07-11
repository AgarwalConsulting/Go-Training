k8s-kind-create:
	kind create cluster --config devops/config/kind/multi-node.yaml

k8s-kind-delete:
	kind delete cluster --name kind
