resource "azurerm_resource_group" "demo" {
  name     = "demo-${var.env}"
  location = var.location
}

output "rg_name" {
  value = azurerm_resource_group.demo.name
}
