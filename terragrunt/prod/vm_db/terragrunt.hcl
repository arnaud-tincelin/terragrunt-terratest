terraform {
  source = "${path_relative_from_include()}/../terraform//vm"
}

include {
  path = find_in_parent_folders()
}

dependency "basement" {
  config_path = "../basement"
}

locals {}

inputs = {
  name      = "db"
  rg_name   = dependency.network.outputs.rg_name
  subnet_id = dependency.network.outputs.subnet_id
storage_account = merge(dependency.basement.outputs.storage, {
    container_name = dependency.basement.outputs.storage_container_name
  })
}
