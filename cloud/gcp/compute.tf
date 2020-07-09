resource "google_compute_instance" "gcp_vm_101" {
  project      = var.project_id
  zone         = var.zone
  name         = var.vm_name
  machine_type = "n1-standard-4"

  allow_stopping_for_update = true

  boot_disk {
    initialize_params {
      image = "debian-9-stretch-v20181210"
      size  = "20"
    }
  }

  network_interface {
    # A default network is created for all GCP projects
    network       = "default"
    access_config {
    }
  }

  metadata = {
    ssh-keys = "gcp:${file(var.publickeyfile)}"
  }

  tags = [
    "http-server",
  ]
}
