---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_settings"
description: |-
  Configure endpoint control settings.
---

# fortios_endpointcontrol_settings
Configure endpoint control settings. Applies to FortiOS Version `6.2.0,6.2.4,6.2.6,7.4.0,7.4.1`.

## Example Usage

```hcl
resource "fortios_endpointcontrol_settings" "trname" {
  download_location                     = "fortiguard"
  forticlient_avdb_update_interval      = 8
  forticlient_dereg_unsupported_client  = "enable"
  forticlient_ems_rest_api_call_timeout = 5000
  forticlient_keepalive_interval        = 60
  forticlient_offline_grace             = "disable"
  forticlient_offline_grace_interval    = 120
  forticlient_reg_key_enforce           = "disable"
  forticlient_reg_timeout               = 7
  forticlient_sys_update_interval       = 720
  forticlient_user_avatar               = "enable"
  forticlient_warning_interval          = 1
}
```

## Argument Reference

The following arguments are supported:

* `override` - Override global EMS table for this VDOM. Valid values: `enable`, `disable`.
* `forticlient_reg_key_enforce` - Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
* `forticlient_reg_key` - FortiClient registration key.
* `forticlient_reg_timeout` - FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
* `download_custom_link` - Customized URL for downloading FortiClient.
* `download_location` - FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
* `forticlient_offline_grace` - Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
* `forticlient_offline_grace_interval` - Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
* `forticlient_keepalive_interval` - Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
* `forticlient_sys_update_interval` - Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
* `forticlient_avdb_update_interval` - Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
* `forticlient_warning_interval` - Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
* `forticlient_user_avatar` - Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
* `forticlient_disconnect_unsupported_client` - Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
* `forticlient_dereg_unsupported_client` - Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
* `forticlient_ems_rest_api_call_timeout` - FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

EndpointControl Settings can be imported using any of these accepted formats:
```
$ terraform import fortios_endpointcontrol_settings.labelname EndpointControlSettings

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_endpointcontrol_settings.labelname EndpointControlSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
