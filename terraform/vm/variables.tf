variable "env" {
  type = string
}

variable "location" {
  type = string
}

variable "name" {
  type = string
}

variable "rg_name" {
  type = string
}

variable "subnet_id" {
  type = string
}

variable "storage_account" {
  type = object({
    id             = string
    name           = string
    container_name = string
  })
}
