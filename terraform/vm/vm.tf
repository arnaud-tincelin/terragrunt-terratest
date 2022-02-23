resource "azurerm_network_interface" "demo" {
  name                = "demo-${var.name}-${var.env}"
  location            = var.location
  resource_group_name = var.rg_name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_linux_virtual_machine" "demo" {
  name                = "demo-${var.name}-${var.env}"
  resource_group_name = var.rg_name
  location            = var.location

  size                            = "Standard_B1s"
  admin_username                  = "adminuser"
  admin_password                  = "LeP@ssw0rd123"
  disable_password_authentication = false
  custom_data                     = data.template_cloudinit_config.vm.rendered

  network_interface_ids = [
    azurerm_network_interface.demo.id,
  ]

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }

  identity {
    type = "SystemAssigned"
  }
}

resource "azurerm_role_assignment" "vm_access_storage_account" {
  scope                = var.storage_account.id
  role_definition_name = "Storage Blob Data Contributor"
  principal_id         = azurerm_linux_virtual_machine.demo.identity.0.principal_id
}

output "vm_name" {
  value = azurerm_linux_virtual_machine.demo.name
}
