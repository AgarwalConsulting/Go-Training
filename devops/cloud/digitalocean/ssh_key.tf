resource "digitalocean_ssh_key" "big_mac" {
  name       = "Big Mac"
  public_key = file("/Users/gaurav/.ssh/id_ed25519.pub")
}
