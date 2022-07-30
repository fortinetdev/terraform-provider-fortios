---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_profile"
description: |-
  Configure FortiClient endpoint control profiles.
---

# fortios_endpointcontrol_profile
Configure FortiClient endpoint control profiles. Applies to FortiOS Version `<= 6.2.0`.

## Example Usage

```hcl
resource "fortios_endpointcontrol_profile" "trname" {
  profile_name = "1"

  device_groups {
    name = "Mobile Devices"
  }

  forticlient_android_settings {
    disable_wf_when_protected    = "enable"
    forticlient_advanced_vpn     = "disable"
    forticlient_vpn_provisioning = "disable"
    forticlient_wf               = "disable"
  }

  forticlient_ios_settings {
    client_vpn_provisioning          = "disable"
    disable_wf_when_protected        = "enable"
    distribute_configuration_profile = "disable"
    forticlient_wf                   = "disable"
  }

  forticlient_winmac_settings {
    av_realtime_protection                         = "disable"
    av_signature_up_to_date                        = "disable"
    forticlient_application_firewall               = "disable"
    forticlient_av                                 = "disable"
    forticlient_ems_compliance                     = "disable"
    forticlient_ems_compliance_action              = "warning"
    forticlient_linux_ver                          = "5.4.1"
    forticlient_log_upload                         = "enable"
    forticlient_log_upload_level                   = "traffic vulnerability event"
    forticlient_mac_ver                            = "5.4.1"
    forticlient_minimum_software_version           = "disable"
    forticlient_registration_compliance_action     = "warning"
    forticlient_security_posture                   = "disable"
    forticlient_security_posture_compliance_action = "warning"
    forticlient_system_compliance                  = "enable"
    forticlient_system_compliance_action           = "warning"
    forticlient_vuln_scan                          = "enable"
    forticlient_vuln_scan_compliance_action        = "warning"
    forticlient_vuln_scan_enforce                  = "high"
    forticlient_vuln_scan_enforce_grace            = 1
    forticlient_vuln_scan_exempt                   = "disable"
    forticlient_wf                                 = "disable"
    forticlient_win_ver                            = "5.4.1"
    os_av_software_installed                       = "disable"
    sandbox_analysis                               = "disable"
  }

  on_net_addr {
    name = "all"
  }

  users {
    name = "guest"
  }
}
```

## Argument Reference

The following arguments are supported:

* `profile_name` - Profile name.
* `forticlient_winmac_settings` - FortiClient settings for Windows/Mac platform. The structure of `forticlient_winmac_settings` block is documented below.
* `forticlient_android_settings` - FortiClient settings for Android platform. The structure of `forticlient_android_settings` block is documented below.
* `forticlient_ios_settings` - FortiClient settings for iOS platform. The structure of `forticlient_ios_settings` block is documented below.
* `description` - Description.
* `src_addr` - Source addresses. The structure of `src_addr` block is documented below.
* `device_groups` - Device groups. The structure of `device_groups` block is documented below.
* `users` - Users. The structure of `users` block is documented below.
* `user_groups` - User groups. The structure of `user_groups` block is documented below.
* `on_net_addr` - Addresses for on-net detection. The structure of `on_net_addr` block is documented below.
* `replacemsg_override_group` - Select an endpoint control replacement message override group from available options.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `forticlient_winmac_settings` block supports:

* `forticlient_registration_compliance_action` - FortiClient registration compliance action. Valid values: `block`, `warning`.
* `forticlient_ems_compliance` - Enable/disable FortiClient Enterprise Management Server (EMS) compliance. Valid values: `enable`, `disable`.
* `forticlient_ems_compliance_action` - FortiClient EMS compliance action. Valid values: `block`, `warning`.
* `forticlient_ems_entries` - FortiClient EMS entries. The structure of `forticlient_ems_entries` block is documented below.
* `forticlient_security_posture` - Enable/disable FortiClient security posture check options. Valid values: `enable`, `disable`.
* `forticlient_security_posture_compliance_action` - FortiClient security posture compliance action. Valid values: `block`, `warning`.
* `forticlient_av` - Enable/disable FortiClient AntiVirus scanning. Valid values: `enable`, `disable`.
* `av_realtime_protection` - Enable/disable FortiClient AntiVirus real-time protection. Valid values: `enable`, `disable`.
* `av_signature_up_to_date` - Enable/disable FortiClient AV signature updates. Valid values: `enable`, `disable`.
* `sandbox_analysis` - Enable/disable sending files to FortiSandbox for analysis. Valid values: `enable`, `disable`.
* `sandbox_address` - FortiSandbox address.
* `os_av_software_installed` - Enable/disable checking for OS recognized AntiVirus software. Valid values: `enable`, `disable`.
* `forticlient_application_firewall` - Enable/disable the FortiClient application firewall. Valid values: `enable`, `disable`.
* `forticlient_application_firewall_list` - FortiClient application firewall rule list.
* `forticlient_wf` - Enable/disable FortiClient web filtering. Valid values: `enable`, `disable`.
* `forticlient_wf_profile` - The FortiClient web filter profile to apply.
* `forticlient_system_compliance` - Enable/disable enforcement of FortiClient system compliance. Valid values: `enable`, `disable`.
* `forticlient_system_compliance_action` - Block or warn clients not compliant with FortiClient requirements. Valid values: `block`, `warning`.
* `forticlient_minimum_software_version` - Enable/disable requiring clients to run FortiClient with a minimum software version number. Valid values: `enable`, `disable`.
* `forticlient_win_ver` - Minimum FortiClient Windows version.
* `forticlient_mac_ver` - Minimum FortiClient Mac OS version.
* `forticlient_linux_ver` - Minimum FortiClient Linux version.
* `forticlient_operating_system` - FortiClient operating system. The structure of `forticlient_operating_system` block is documented below.
* `forticlient_running_app` - Use FortiClient to verify if the listed applications are running on the client. The structure of `forticlient_running_app` block is documented below.
* `forticlient_registry_entry` - FortiClient registry entry. The structure of `forticlient_registry_entry` block is documented below.
* `forticlient_own_file` - Checking the path and filename of the FortiClient application. The structure of `forticlient_own_file` block is documented below.
* `forticlient_log_upload` - Enable/disable uploading FortiClient logs. Valid values: `enable`, `disable`.
* `forticlient_log_upload_level` - Select the FortiClient logs to upload. Valid values: `traffic`, `vulnerability`, `event`.
* `forticlient_log_upload_server` - IP address or FQDN of the server to which to upload FortiClient logs.
* `forticlient_vuln_scan` - Enable/disable FortiClient vulnerability scanning. Valid values: `enable`, `disable`.
* `forticlient_vuln_scan_compliance_action` - FortiClient vulnerability compliance action. Valid values: `block`, `warning`.
* `forticlient_vuln_scan_enforce` - Configure the level of the vulnerability found that causes a FortiClient vulnerability compliance action. Valid values: `critical`, `high`, `medium`, `low`, `info`.
* `forticlient_vuln_scan_enforce_grace` - FortiClient vulnerability scan enforcement grace period (0 - 30 days, default = 1).
* `forticlient_vuln_scan_exempt` - Enable/disable compliance exemption for vulnerabilities that cannot be patched automatically. Valid values: `enable`, `disable`.

The `forticlient_ems_entries` block supports:

* `name` - FortiClient EMS name.

The `forticlient_operating_system` block supports:

* `id` - Operating system entry ID.
* `os_type` - Operating system type. Valid values: `custom`, `mac-os`, `win-7`, `win-80`, `win-81`, `win-10`, `win-2000`, `win-home-svr`, `win-svr-10`, `win-svr-2003`, `win-svr-2003-r2`, `win-svr-2008`, `win-svr-2008-r2`, `win-svr-2012`, `win-svr-2012-r2`, `win-sto-svr-2003`, `win-vista`, `win-xp`, `ubuntu-linux`, `centos-linux`, `redhat-linux`, `fedora-linux`.
* `os_name` - Customize operating system name or Mac OS format:x.x.x

The `forticlient_running_app` block supports:

* `id` - Application ID.
* `app_name` - Application name.
* `application_check_rule` - Application check rule. Valid values: `present`, `absent`.
* `process_name` - Process name.
* `app_sha256_signature` - App's SHA256 signature.
* `process_name2` - Process name.
* `app_sha256_signature2` - App's SHA256 Signature.
* `process_name3` - Process name.
* `app_sha256_signature3` - App's SHA256 Signature.
* `process_name4` - Process name.
* `app_sha256_signature4` - App's SHA256 Signature.

The `forticlient_registry_entry` block supports:

* `id` - Registry entry ID.
* `registry_entry` - Registry entry.

The `forticlient_own_file` block supports:

* `id` - File ID.
* `file` - File path and name.

The `forticlient_android_settings` block supports:

* `forticlient_wf` - Enable/disable FortiClient web filtering. Valid values: `enable`, `disable`.
* `forticlient_wf_profile` - The FortiClient web filter profile to apply.
* `disable_wf_when_protected` - Enable/disable FortiClient web category filtering when protected by FortiGate. Valid values: `enable`, `disable`.
* `forticlient_vpn_provisioning` - Enable/disable FortiClient VPN provisioning. Valid values: `enable`, `disable`.
* `forticlient_advanced_vpn` - Enable/disable advanced FortiClient VPN configuration. Valid values: `enable`, `disable`.
* `forticlient_advanced_vpn_buffer` - Advanced FortiClient VPN configuration.
* `forticlient_vpn_settings` - FortiClient VPN settings. The structure of `forticlient_vpn_settings` block is documented below.

The `forticlient_vpn_settings` block supports:

* `name` - VPN name.
* `type` - VPN type (IPsec or SSL VPN). Valid values: `ipsec`, `ssl`.
* `remote_gw` - IP address or FQDN of the remote VPN gateway.
* `sslvpn_access_port` - SSL VPN access port (1 - 65535).
* `sslvpn_require_certificate` - Enable/disable requiring SSL VPN client certificate. Valid values: `enable`, `disable`.
* `auth_method` - Authentication method. Valid values: `psk`, `certificate`.
* `preshared_key` - Pre-shared secret for PSK authentication.

The `forticlient_ios_settings` block supports:

* `forticlient_wf` - Enable/disable FortiClient web filtering. Valid values: `enable`, `disable`.
* `forticlient_wf_profile` - The FortiClient web filter profile to apply.
* `disable_wf_when_protected` - Enable/disable FortiClient web category filtering when protected by FortiGate. Valid values: `enable`, `disable`.
* `client_vpn_provisioning` - FortiClient VPN provisioning. Valid values: `enable`, `disable`.
* `client_vpn_settings` - FortiClient VPN settings. The structure of `client_vpn_settings` block is documented below.
* `distribute_configuration_profile` - Enable/disable configuration profile (.mobileconfig file) distribution. Valid values: `enable`, `disable`.
* `configuration_name` - Name of configuration profile.
* `configuration_content` - Content of configuration profile.

The `client_vpn_settings` block supports:

* `name` - VPN name.
* `type` - VPN type (IPsec or SSL VPN). Valid values: `ipsec`, `ssl`.
* `vpn_configuration_name` - Name of VPN configuration.
* `vpn_configuration_content` - Content of VPN configuration.
* `remote_gw` - IP address or FQDN of the remote VPN gateway.
* `sslvpn_access_port` - SSL VPN access port (1 - 65535).
* `sslvpn_require_certificate` - Enable/disable requiring SSL VPN client certificate. Valid values: `enable`, `disable`.
* `auth_method` - Authentication method. Valid values: `psk`, `certificate`.
* `preshared_key` - Pre-shared secret for PSK authentication.

The `src_addr` block supports:

* `name` - Address object from available options.

The `device_groups` block supports:

* `name` - Device group object from available options.

The `users` block supports:

* `name` - User name.

The `user_groups` block supports:

* `name` - User group name.

The `on_net_addr` block supports:

* `name` - Address object from available options.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{profile_name}}.

## Import

EndpointControl Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_endpointcontrol_profile.labelname {{profile_name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_endpointcontrol_profile.labelname {{profile_name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
