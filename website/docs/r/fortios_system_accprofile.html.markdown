---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_accprofile"
description: |-
  Configure access profiles for system administrators.
---

# fortios_system_accprofile
Configure access profiles for system administrators.

## Example Usage

```hcl
resource "fortios_system_accprofile" "test12" {
  name                  = "tst1"
  admintimeout          = 10
  admintimeout_override = "disable"
  authgrp               = "read-write"
  ftviewgrp             = "read-write"
  fwgrp                 = "custom"
  loggrp                = "read-write"
  netgrp                = "read-write"
  scope                 = "vdom"
  secfabgrp             = "read-write"
  sysgrp                = "read-write"
  utmgrp                = "custom"
  vpngrp                = "read-write"
  wanoptgrp             = "read-write"
  wifi                  = "read-write"

  fwgrp_permission {
    address  = "read-write"
    policy   = "read-write"
    schedule = "none"
    service  = "none"
  }

  loggrp_permission {
    config        = "none"
    data_access   = "none"
    report_access = "none"
    threat_weight = "none"
  }

  netgrp_permission {
    cfg            = "none"
    packet_capture = "none"
    route_cfg      = "none"
  }

  sysgrp_permission {
    admin = "none"
    cfg   = "none"
    mnt   = "none"
    upd   = "none"
  }

  utmgrp_permission {
    antivirus            = "read-write"
    application_control  = "none"
    data_loss_prevention = "none"
    dnsfilter            = "none"
    endpoint_control     = "none"
    icap                 = "none"
    ips                  = "read-write"
    voip                 = "none"
    waf                  = "none"
    webfilter            = "none"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Profile name.
* `scope` - Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
* `comments` - Comment.
* `secfabgrp` - Security Fabric. Valid values: `none`, `read`, `read-write`.
* `ftviewgrp` - FortiView. Valid values: `none`, `read`, `read-write`.
* `authgrp` - Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
* `sysgrp` - System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
* `netgrp` - Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
* `loggrp` - Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
* `fwgrp` - Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
* `vpngrp` - Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
* `utmgrp` - Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
* `wanoptgrp` - Administrator access to WAN Opt & Cache. Valid values: `none`, `read`, `read-write`.
* `wifi` - Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
* `netgrp_permission` - Custom network permission. The structure of `netgrp_permission` block is documented below.
* `sysgrp_permission` - Custom system permission. The structure of `sysgrp_permission` block is documented below.
* `fwgrp_permission` - Custom firewall permission. The structure of `fwgrp_permission` block is documented below.
* `loggrp_permission` - Custom Log & Report permission. The structure of `loggrp_permission` block is documented below.
* `utmgrp_permission` - Custom Security Profile permissions. The structure of `utmgrp_permission` block is documented below.
* `admintimeout_override` - Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
* `admintimeout` - Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
* `system_diagnostics` - Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `netgrp_permission` block supports:

* `cfg` - Network Configuration. Valid values: `none`, `read`, `read-write`.
* `packet_capture` - Packet Capture Configuration. Valid values: `none`, `read`, `read-write`.
* `route_cfg` - Router Configuration. Valid values: `none`, `read`, `read-write`.

The `sysgrp_permission` block supports:

* `admin` - Administrator Users. Valid values: `none`, `read`, `read-write`.
* `upd` - FortiGuard Updates. Valid values: `none`, `read`, `read-write`.
* `cfg` - System Configuration. Valid values: `none`, `read`, `read-write`.
* `mnt` - Maintenance. Valid values: `none`, `read`, `read-write`.

The `fwgrp_permission` block supports:

* `policy` - Policy Configuration. Valid values: `none`, `read`, `read-write`.
* `address` - Address Configuration. Valid values: `none`, `read`, `read-write`.
* `service` - Service Configuration. Valid values: `none`, `read`, `read-write`.
* `schedule` - Schedule Configuration. Valid values: `none`, `read`, `read-write`.
* `others` - Other Firewall Configuration. Valid values: `none`, `read`, `read-write`.

The `loggrp_permission` block supports:

* `config` - Log & Report configuration. Valid values: `none`, `read`, `read-write`.
* `data_access` - Log & Report Data Access. Valid values: `none`, `read`, `read-write`.
* `report_access` - Log & Report Report Access. Valid values: `none`, `read`, `read-write`.
* `threat_weight` - Log & Report Threat Weight. Valid values: `none`, `read`, `read-write`.

The `utmgrp_permission` block supports:

* `antivirus` - Antivirus profiles and settings. Valid values: `none`, `read`, `read-write`.
* `ips` - IPS profiles and settings. Valid values: `none`, `read`, `read-write`.
* `webfilter` - Web Filter profiles and settings. Valid values: `none`, `read`, `read-write`.
* `emailfilter` - AntiSpam filter and settings. Valid values: `none`, `read`, `read-write`.
* `spamfilter` - AntiSpam filter and settings. Valid values: `none`, `read`, `read-write`.
* `data_loss_prevention` - DLP profiles and settings. Valid values: `none`, `read`, `read-write`.
* `file_filter` - File-filter profiles and settings. Valid values: `none`, `read`, `read-write`.
* `application_control` - Application Control profiles and settings. Valid values: `none`, `read`, `read-write`.
* `icap` - ICAP profiles and settings. Valid values: `none`, `read`, `read-write`.
* `voip` - VoIP profiles and settings. Valid values: `none`, `read`, `read-write`.
* `waf` - Web Application Firewall profiles and settings. Valid values: `none`, `read`, `read-write`.
* `dnsfilter` - DNS Filter profiles and settings. Valid values: `none`, `read`, `read-write`.
* `endpoint_control` - FortiClient Profiles. Valid values: `none`, `read`, `read-write`.
* `videofilter` - Video filter profiles and settings. Valid values: `none`, `read`, `read-write`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Accprofile can be imported using any of these accepted formats:
```
$ terraform import fortios_system_accprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_accprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
