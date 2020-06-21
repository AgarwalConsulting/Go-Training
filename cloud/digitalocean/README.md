# Setup infra on Digital Ocean

## Setup infra

```bash
terraform apply
./scripts/setup.sh
```

## Destroy infra

```bash
terraform destroy
```

## Outputs

```bash
# To connect to the k8s cluster
terraform output training_cluster_raw_config > ~/.kube/config
# To connect to the machine: `ssh root@<ip>`
terraform output plaything_ip
```
