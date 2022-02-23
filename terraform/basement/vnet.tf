resource "azurerm_virtual_network" "demo" {
  name                = "demo-${var.env}"
  location            = var.location
  resource_group_name = azurerm_resource_group.demo.name
  address_space       = ["10.0.0.0/24"]
}

resource "azurerm_subnet" "default" {
  name                                           = "default"
  resource_group_name                            = azurerm_resource_group.demo.name
  virtual_network_name                           = azurerm_virtual_network.demo.name
  address_prefixes                               = ["10.0.0.0/24"]
  enforce_private_link_endpoint_network_policies = true
}

output "subnet_id" {
  value = azurerm_subnet.default.id
}
