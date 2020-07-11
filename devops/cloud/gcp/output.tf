output compute-instance {
  value = google_compute_instance.gcp_vm_101.network_interface.0.access_config.0.nat_ip
}
