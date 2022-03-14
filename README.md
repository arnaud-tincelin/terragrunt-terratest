# Example of Terragrunt + Terratest

This repository aims to demonstrate how to use Terragrunt and Terratest.  
It deploys a VNET and a Storage Account with a Private Endpoint (module [basement](./terraform/basement/) and a VM (module [VM](./terraform/vm/).  
The test validates the VM can access the SA through its PE: the VM shall delete a blob from the SA.

## Required tooling

- [Terraform](https://github.com/hashicorp/terraform)
- [Terragrunt](https://github.com/gruntwork-io/terragrunt)
- [Golang](https://go.dev/doc/install)

## Deploy / Destroy with Terragrunt

1. create a storage account & and storage container on Azure to store the remote state
1. create the `state.yaml` file at the root of the [terragrunt](./terragrunt) folder from the following template:

  ```yaml
  subscription_id: <subscription where the storage account exists>
  resource_group_name: <RG where the storage account exists>
  storage_account_name: <name of the SA created at the step1>
  container_name: <container within the SA>
  ```

1. create the `env.hcl` file in the folder of your choice: [dev](./terragrunt/dev) or [prod](./terragrunt/dev) from the following template:

    ```hcl
    locals {
      subscription_id     = <subscription id where to deploy the resources>
      env                 = <name of your env. ex: "dev">
      location            = <ex: "west europe">
      storage_allowed_pip = <array of public IP allowed to reach the storage account ex: ["1.1.1.1"]>
    }
    ```

1. Login on Azure:

    - option 1: run `az login`
    - option 2: export the following environment variables:
      - `ARM_CLIENT_ID`
      - `ARM_CLIENT_SECRET`
      - `ARM_TENANT_ID`
      - `ARM_SUBSCRIPTION_ID`

1. Deploy:

    ```bash
    terragrunt run-all apply
    ```

1. Destroy:

    ```bash
    terragrunt run-all apply
    ```

## Execute the tests

1. Run `cd ./terratest`
1. Set the following environment variables:

    - `AZURE_CLIENT_ID`
    - `AZURE_TENANT_ID`
    - `AZURE_CLIENT_SECRET`

1. Run `go test *.go -v -timeout 60m`
