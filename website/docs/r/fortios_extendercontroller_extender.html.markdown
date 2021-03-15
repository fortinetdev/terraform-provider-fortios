---
subcategory: "FortiGate Extender-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extendercontroller_extender"
description: |-
  Extender controller configuration.
---

# fortios_extendercontroller_extender
Extender controller configuration.

-> The resource applies to FortiOS Version < 6.4.2. For FortiOS Version >= 6.4.2, see `fortios_extendercontroller_extender1`.


## Example Usage

```hcl
resource "fortios_extendercontroller_extender" "trname" {
  admin               = "disable"
  billing_start_day   = 1
  conn_status         = 0
  dial_mode           = "always-connect"
  dial_status         = 0
  ext_name            = "332"
  fosid               = "1"
  initiated_update    = "disable"
  mode                = "standalone"
  modem_type          = "gsm/lte"
  multi_mode          = "auto"
  ppp_auth_protocol   = "auto"
  ppp_echo_request    = "disable"
  quota_limit_mb      = 0
  redial              = "none"
  roaming             = "disable"
  role                = "primary"
  vdom                = 0
  wimax_auth_protocol = "tls"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) FortiExtender serial number.
* `admin` - (Required) FortiExtender Administration (enable or disable). Valid values: `disable`, `discovered`, `enable`.
* `ifname` - FortiExtender interface name.
* `vdom` - VDOM
* `role` - (Required) FortiExtender work role(Primary, Secondary, None). Valid values: `none`, `primary`, `secondary`.
* `mode` - FortiExtender mode. Valid values: `standalone`, `redundant`.
* `dial_mode` - Dial mode (dial-on-demand or always-connect). Valid values: `dial-on-demand`, `always-connect`.
* `redial` - Number of redials allowed based on failed attempts. Valid values: `none`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`.
* `redundant_intf` - Redundant interface.
* `dial_status` - Dial status.
* `conn_status` - Connection status.
* `ext_name` - FortiExtender name.
* `description` - Description.
* `quota_limit_mb` - Monthly quota limit (MB).
* `billing_start_day` - Billing start day.
* `at_dial_script` - Initialization AT commands specific to the MODEM.
* `modem_passwd` - MODEM password.
* `initiated_update` - Allow/disallow network initiated updates to the MODEM. Valid values: `enable`, `disable`.
* `modem_type` - MODEM type (CDMA, GSM/LTE or WIMAX). Valid values: `cdma`, `gsm/lte`, `wimax`.
* `ppp_username` - PPP username.
* `ppp_password` - PPP password.
* `ppp_auth_protocol` - PPP authentication protocol (PAP,CHAP or auto). Valid values: `auto`, `pap`, `chap`.
* `ppp_echo_request` - Enable/disable PPP echo request. Valid values: `enable`, `disable`.
* `wimax_carrier` - WiMax carrier.
* `wimax_realm` - WiMax realm.
* `wimax_auth_protocol` - WiMax authentication protocol(TLS or TTLS). Valid values: `tls`, `ttls`.
* `sim_pin` - SIM PIN.
* `access_point_name` - Access point name(APN).
* `multi_mode` - MODEM mode of operation(3G,LTE,etc). Valid values: `auto`, `auto-3g`, `force-lte`, `force-3g`, `force-2g`.
* `roaming` - Enable/disable MODEM roaming. Valid values: `enable`, `disable`.
* `cdma_nai` - NAI for CDMA MODEMS.
* `aaa_shared_secret` - AAA shared secret.
* `ha_shared_secret` - HA shared secret.
* `primary_ha` - Primary HA.
* `secondary_ha` - Secondary HA.
* `cdma_aaa_spi` - CDMA AAA SPI.
* `cdma_ha_spi` - CDMA HA SPI.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ExtenderController Extender can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_extendercontroller_extender.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
