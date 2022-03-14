terraform {
  source = "${path_relative_from_include()}/../terraform//basement"
}

include {
  path = find_in_parent_folders()
}

inputs = {}
