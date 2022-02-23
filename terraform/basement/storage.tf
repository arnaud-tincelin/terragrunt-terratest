resource "azurerm_storage_account" "demo" {
  name                     = "satidemo${var.env}"
  resource_group_name      = azurerm_resource_group.demo.name
  location                 = var.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
  allow_blob_public_access = false

  network_rules {
    default_action = "Deny"
    ip_rules       = var.storage_allowed_pip
  }
}

output "storage" {
  value = {
    name = azurerm_storage_account.demo.name
    id   = azurerm_storage_account.demo.id
  }
}

resource "azurerm_storage_container" "demo" {
  name                  = "demo"
  storage_account_name  = azurerm_storage_account.demo.name
  container_access_type = "private"
}

output "storage_container_name" {
  value = azurerm_storage_container.demo.name
}

resource "azurerm_storage_blob" "demo" {
  name                   = "demo.txt"
  storage_account_name   = azurerm_storage_account.demo.name
  storage_container_name = azurerm_storage_container.demo.name
  type                   = "Block"
  source_content         = <<EOF
  This is an example blob
  EOF
}

resource "azurerm_private_dns_zone" "privatelink_blob" {
  name                = "privatelink.blob.core.windows.net"
  resource_group_name = azurerm_resource_group.demo.name
}

resource "azurerm_private_dns_zone_virtual_network_link" "blob" {
  name                  = "blob"
  resource_group_name   = azurerm_resource_group.demo.name
  private_dns_zone_name = azurerm_private_dns_zone.privatelink_blob.name
  virtual_network_id    = azurerm_virtual_network.demo.id
}

resource "azurerm_private_endpoint" "storage" {
  name                = azurerm_storage_account.demo.name
  resource_group_name = azurerm_resource_group.demo.name
  location            = var.location
  subnet_id           = azurerm_subnet.default.id

  private_dns_zone_group {
    name                 = "privatelink-blob-core-windows-net"
    private_dns_zone_ids = [azurerm_private_dns_zone.privatelink_blob.id]
  }

  private_service_connection {
    name                           = azurerm_storage_account.demo.name
    private_connection_resource_id = azurerm_storage_account.demo.id
    subresource_names              = ["blob"]
    is_manual_connection           = false
  }
}
