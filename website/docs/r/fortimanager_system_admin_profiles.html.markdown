---
layout: "fortios"
page_title: "FortiOS: fortios_fortimanager_system_admin_profiles"
sidebar_current: "docs-fortios-fortimanager-resource-system-admin-profiles"
description: |-
  Provides a resource to configure system admin profiles for FortiManager.
---

# fortios_fortimanager_system_admin_profiles
This resource supports Create/Read/Update/Delete admin profiles for FortiManager

## Example Usage
```hcl
provider "fortios" {
	hostname = "192.168.88.100"
	username = "APIUser"
	passwd = "admin"
	product = "fortimanager"
	insecure = false
	cabundlefile = "/path/yourCA.crt"
}

resource "fortios_fortimanager_system_admin_profiles" "test1" {
	profileid = "terraform-test1"
	description = "11"
	system_setting = "read"
	adom_switch = "read"
	deploy_management = "read"
	import_policy_packages = "read"
	intf_mapping = "read-write"
	device_ap = "none"
	device_forticlient = "read"
	device_fortiswitch = "read"
	vpn_manager = "read"
	log_viewer = "read"
	fortiguard_center = "read"
	fortiguard_center_advanced = "read"
	fortiguard_center_firmware_managerment = "read"
	fortiguard_center_licensing = "read"
	device_manager = "read-write"
	device_operation = "read"
	config_retrieve = "read"
	config_revert = "read"
	device_revision_deletion = "read"
	terminal_access = "read"
	device_config = "read"
	device_profile = "read"
	device_wan_link_load_balance = "read"
	policy_objects = "read-write"
	global_policy_packages = "read-write"
	assignment = "read"
	adom_policy_packages = "read"
	consistency_check = "read-write"
	set_install_targets = "read-write"
}
```

## Argument Reference
The following arguments are supported:

* `profileid` - (Required) Profile name.
* `description` - Description.
* `system_setting` - System Setting.
* `adom_switch` - Administrator Domain.
* `deploy_management` - Install to devices.
* `import_policy_packages` - Import Policy Package.
* `intf_mapping` - Interface Mapping.
* `device_ap` - Manage AP.
* `device_forticlient` - Manage FortiClient.
* `device_fortiswitch` - Manage FortiSwitch.
* `vpn_manager` - VPN Manager.
* `log_viewer` - Log Viewer.
* `fortiguard_center` - FortiGuard Center.
* `fortiguard_center_advanced` - FortiGuard Center Advanced.
* `fortiguard_center_firmware_managerment` - FortiGuard Center Firmware Managerment.
* `fortiguard_center_licensing` - FortiGuard Center Licensing.
* `device_manager` - Device Manager.
* `device_operation` - Device add/delete/edit.
* `config_retrieve` - Configuration Retrieve.
* `config_revert` - Revert Configuration from Revision History.
* `device_revision_deletion` - Delete device revision.
* `terminal_access` - Terminal access.
* `device_config` - Manage device configurations.
* `device_profile` - Device profile permission.
* `device_wan_link_load_balance` - Manage WAN link load balance.
* `policy_objects` - Policy objects permission.
* `global_policy_packages` - Global policy packages.
* `assignment` - Assignment Permission.
* `adom_policy_packages` - Adom policy packages.
* `consistency_check` - Consistency check.
* `set_install_targets` - Edit installation targets.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `profileid` - Profile name.
* `description` - Description.
* `system_setting` - System Setting.
* `adom_switch` - Administrator Domain.
* `deploy_management` - Install to devices.
* `import_policy_packages` - Import Policy Package.
* `intf_mapping` - Interface Mapping.
* `device_ap` - Manage AP.
* `device_forticlient` - Manage FortiClient.
* `device_fortiswitch` - Manage FortiSwitch.
* `vpn_manager` - VPN Manager.
* `log_viewer` - Log Viewer.
* `fortiguard_center` - FortiGuard Center.
* `fortiguard_center_advanced` - FortiGuard Center Advanced.
* `fortiguard_center_firmware_managerment` - FortiGuard Center Firmware Managerment.
* `fortiguard_center_licensing` - FortiGuard Center Licensing.
* `device_manager` - Device Manager.
* `device_operation` - Device add/delete/edit.
* `config_retrieve` - Configuration Retrieve.
* `config_revert` - Revert Configuration from Revision History.
* `device_revision_deletion` - Delete device revision.
* `terminal_access` - Terminal access.
* `device_config` - Manage device configurations.
* `device_profile` - Device profile permission.
* `device_wan_link_load_balance` - Manage WAN link load balance.
* `policy_objects` - Policy objects permission.
* `global_policy_packages` - Global policy packages.
* `assignment` - Assignment Permission.
* `adom_policy_packages` - Adom policy packages.
* `consistency_check` - Consistency check.
* `set_install_targets` - Edit installation targets.
