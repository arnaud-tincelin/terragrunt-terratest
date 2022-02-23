variable "env" {
  type = string
}

variable "location" {
  type = string
}

variable "storage_allowed_pip" {
  type = list(string)
  default = []
}
