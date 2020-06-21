output plaything_ip {
  value = digitalocean_droplet.plaything.ipv4_address
}

output training_cluster_config {
  value = digitalocean_kubernetes_cluster.training_cluster.kube_config.0
}

output training_cluster_raw_config {
  value = digitalocean_kubernetes_cluster.training_cluster.kube_config.0.raw_config
}
