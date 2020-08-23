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
* `scope` - Scope of admin access: global or specific VDOM(s).
* `comments` - Comment.
* `secfabgrp` - Security Fabric.
* `ftviewgrp` - FortiView.
* `authgrp` - Administrator access to Users and Devices.
* `sysgrp` - System Configuration.
* `netgrp` - Network Configuration.
* `loggrp` - Administrator access to Logging and Reporting including viewing log messages.
* `fwgrp` - Administrator access to the Firewall configuration.
* `vpngrp` - Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
* `utmgrp` - Administrator access to Security Profiles.
* `wanoptgrp` - Administrator access to WAN Opt & Cache.
* `wifi` - Administrator access to the WiFi controller and Switch controller.
* `netgrp_permission` - Custom network permission.
* `sysgrp_permission` - Custom system permission.
* `fwgrp_permission` - Custom firewall permission.
* `loggrp_permission` - Custom Log & Report permission.
* `utmgrp_permission` - Custom Security Profile permissions.
* `admintimeout_override` - Enable/disable overriding the global administrator idle timeout.
* `admintimeout` - Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).

The `netgrp_permission` block supports:

* `cfg` - Network Configuration.
* `packet_capture` - Packet Capture Configuration.
* `route_cfg` - Router Configuration.

The `sysgrp_permission` block supports:

* `admin` - Administrator Users.
* `upd` - FortiGuard Updates.
* `cfg` - System Configuration.
* `mnt` - Maintenance.

The `fwgrp_permission` block supports:

* `policy` - Policy Configuration.
* `address` - Address Configuration.
* `service` - Service Configuration.
* `schedule` - Schedule Configuration.

The `loggrp_permission` block supports:

* `config` - Log & Report configuration.
* `data_access` - Log & Report Data Access.
* `report_access` - Log & Report Report Access.
* `threat_weight` - Log & Report Threat Weight.

The `utmgrp_permission` block supports:

* `antivirus` - Antivirus profiles and settings.
* `ips` - IPS profiles and settings.
* `webfilter` - Web Filter profiles and settings.
* `spamfilter` - AntiSpam filter and settings.
* `data_loss_prevention` - DLP profiles and settings.
* `application_control` - Application Control profiles and settings.
* `icap` - ICAP profiles and settings.
* `voip` - VoIP profiles and settings.
* `waf` - Web Application Firewall profiles and settings.
* `dnsfilter` - DNS Filter profiles and settings.
* `endpoint_control` - FortiClient Profiles.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Accprofile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_accprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
