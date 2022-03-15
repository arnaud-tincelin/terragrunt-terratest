skip = true

terraform_version_constraint = ">= 1.0.0"

remote_state {
  backend = "azurerm"
  config = merge(local.remote_state, {
    key = "demo/${local.env.env}/${basename(get_terragrunt_dir())}.tfstate"
  })
  generate = {
    path      = "backend.tf"
    if_exists = "overwrite_terragrunt"
  }
}

generate "providers" {
  path      = "providers.tf"
  if_exists = "overwrite"
  contents  = <<EOF
provider "azurerm" {
  features {}
  subscription_id = "${local.env.subscription_id}"
}
EOF
}

locals {
  env          = read_terragrunt_config(find_in_parent_folders("env.hcl")).locals
  remote_state = yamldecode(file(find_in_parent_folders("state.yaml")))
}

inputs = local.env
