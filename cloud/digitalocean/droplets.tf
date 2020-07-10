resource "digitalocean_droplet" "plaything" {
  name     = "Go-Training"
  size     = "s-2vcpu-4gb"
  image    = "ubuntu-20-04-x64"
  region   = "nyc3"
  vpc_uuid = digitalocean_vpc.vpc_101.id
  ssh_keys = [digitalocean_ssh_key.big_mac.id]
}
