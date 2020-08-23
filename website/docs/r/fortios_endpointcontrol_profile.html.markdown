---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_profile"
description: |-
  Configure FortiClient endpoint control profiles.
---

# fortios_endpointcontrol_profile
Configure FortiClient endpoint control profiles.

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
* `forticlient_winmac_settings` - FortiClient settings for Windows/Mac platform.
* `forticlient_android_settings` - FortiClient settings for Android platform.
* `forticlient_ios_settings` - FortiClient settings for iOS platform.
* `description` - Description.
* `src_addr` - Source addresses.
* `device_groups` - Device groups.
* `users` - Users.
* `user_groups` - User groups.
* `on_net_addr` - Addresses for on-net detection.
* `replacemsg_override_group` - Select an endpoint control replacement message override group from available options.

The `forticlient_winmac_settings` block supports:

* `forticlient_registration_compliance_action` - FortiClient registration compliance action.
* `forticlient_ems_compliance` - Enable/disable FortiClient Enterprise Management Server (EMS) compliance.
* `forticlient_ems_compliance_action` - FortiClient EMS compliance action.
* `forticlient_ems_entries` - FortiClient EMS entries.
* `forticlient_security_posture` - Enable/disable FortiClient security posture check options.
* `forticlient_security_posture_compliance_action` - FortiClient security posture compliance action.
* `forticlient_av` - Enable/disable FortiClient AntiVirus scanning.
* `av_realtime_protection` - Enable/disable FortiClient AntiVirus real-time protection.
* `av_signature_up_to_date` - Enable/disable FortiClient AV signature updates.
* `sandbox_analysis` - Enable/disable sending files to FortiSandbox for analysis.
* `sandbox_address` - FortiSandbox address.
* `os_av_software_installed` - Enable/disable checking for OS recognized AntiVirus software.
* `forticlient_application_firewall` - Enable/disable the FortiClient application firewall.
* `forticlient_application_firewall_list` - FortiClient application firewall rule list.
* `forticlient_wf` - Enable/disable FortiClient web filtering.
* `forticlient_wf_profile` - The FortiClient web filter profile to apply.
* `forticlient_system_compliance` - Enable/disable enforcement of FortiClient system compliance.
* `forticlient_system_compliance_action` - Block or warn clients not compliant with FortiClient requirements.
* `forticlient_minimum_software_version` - Enable/disable requiring clients to run FortiClient with a minimum software version number.
* `forticlient_win_ver` - Minimum FortiClient Windows version.
* `forticlient_mac_ver` - Minimum FortiClient Mac OS version.
* `forticlient_linux_ver` - Minimum FortiClient Linux version.
* `forticlient_operating_system` - FortiClient operating system.
* `forticlient_running_app` - Use FortiClient to verify if the listed applications are running on the client.
* `forticlient_registry_entry` - FortiClient registry entry.
* `forticlient_own_file` - Checking the path and filename of the FortiClient application.
* `forticlient_log_upload` - Enable/disable uploading FortiClient logs.
* `forticlient_log_upload_level` - Select the FortiClient logs to upload.
* `forticlient_log_upload_server` - IP address or FQDN of the server to which to upload FortiClient logs.
* `forticlient_vuln_scan` - Enable/disable FortiClient vulnerability scanning.
* `forticlient_vuln_scan_compliance_action` - FortiClient vulnerability compliance action.
* `forticlient_vuln_scan_enforce` - Configure the level of the vulnerability found that causes a FortiClient vulnerability compliance action.
* `forticlient_vuln_scan_enforce_grace` - FortiClient vulnerability scan enforcement grace period (0 - 30 days, default = 1).
* `forticlient_vuln_scan_exempt` - Enable/disable compliance exemption for vulnerabilities that cannot be patched automatically.

The `forticlient_ems_entries` block supports:

* `name` - FortiClient EMS name.

The `forticlient_operating_system` block supports:

* `id` - Operating system entry ID.
* `os_type` - Operating system type.
* `os_name` - Customize operating system name or Mac OS format:x.x.x

The `forticlient_running_app` block supports:

* `id` - Application ID.
* `app_name` - Application name.
* `application_check_rule` - Application check rule.
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

* `forticlient_wf` - Enable/disable FortiClient web filtering.
* `forticlient_wf_profile` - The FortiClient web filter profile to apply.
* `disable_wf_when_protected` - Enable/disable FortiClient web category filtering when protected by FortiGate.
* `forticlient_vpn_provisioning` - Enable/disable FortiClient VPN provisioning.
* `forticlient_advanced_vpn` - Enable/disable advanced FortiClient VPN configuration.
* `forticlient_advanced_vpn_buffer` - Advanced FortiClient VPN configuration.
* `forticlient_vpn_settings` - FortiClient VPN settings.

The `forticlient_vpn_settings` block supports:

* `name` - VPN name.
* `type` - VPN type (IPsec or SSL VPN).
* `remote_gw` - IP address or FQDN of the remote VPN gateway.
* `sslvpn_access_port` - SSL VPN access port (1 - 65535).
* `sslvpn_require_certificate` - Enable/disable requiring SSL VPN client certificate.
* `auth_method` - Authentication method.
* `preshared_key` - Pre-shared secret for PSK authentication.

The `forticlient_ios_settings` block supports:

* `forticlient_wf` - Enable/disable FortiClient web filtering.
* `forticlient_wf_profile` - The FortiClient web filter profile to apply.
* `disable_wf_when_protected` - Enable/disable FortiClient web category filtering when protected by FortiGate.
* `client_vpn_provisioning` - FortiClient VPN provisioning.
* `client_vpn_settings` - FortiClient VPN settings.
* `distribute_configuration_profile` - Enable/disable configuration profile (.mobileconfig file) distribution.
* `configuration_name` - Name of configuration profile.
* `configuration_content` - Content of configuration profile.

The `client_vpn_settings` block supports:

* `name` - VPN name.
* `type` - VPN type (IPsec or SSL VPN).
* `vpn_configuration_name` - Name of VPN configuration.
* `vpn_configuration_content` - Content of VPN configuration.
* `remote_gw` - IP address or FQDN of the remote VPN gateway.
* `sslvpn_access_port` - SSL VPN access port (1 - 65535).
* `sslvpn_require_certificate` - Enable/disable requiring SSL VPN client certificate.
* `auth_method` - Authentication method.
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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_endpointcontrol_profile.labelname {{profile_name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
