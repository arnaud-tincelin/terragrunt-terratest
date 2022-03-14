data "template_file" "cloudinit" {
  template = file("${path.module}/cloudinit.yaml")

  vars = {
    storage_account = var.storage_account.name
    container_name  = var.storage_account.container_name
    username        = "adminuser"
  }
}

data "template_cloudinit_config" "vm" {
  gzip          = true
  base64_encode = true

  part {
    filename     = "vm.cfg"
    content_type = "text/cloud-config"
    content      = data.template_file.cloudinit.rendered
  }
}
