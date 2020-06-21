resource "digitalocean_kubernetes_cluster" "training_cluster" {
  name    = "training-cluster"
  region  = "nyc3"
  # Grab the latest version slug from `doctl kubernetes options versions`
  version = "1.17.5-do.0"
  tags = [ "101-workshop" ]
  vpc_uuid = digitalocean_vpc.vpc_101.id

  node_pool {
    name       = "worker-pool"
    size       = "s-2vcpu-4gb"
    node_count = 3
  }
}
