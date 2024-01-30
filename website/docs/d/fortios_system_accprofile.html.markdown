---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_accprofile"
description: |-
  Get information on an fortios system accprofile.
---

# Data Source: fortios_system_accprofile
Use this data source to get information on an fortios system accprofile

## Argument Reference

* `name` - (Required) Specify the name of the desired system accprofile.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Profile name.
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
* `netgrp_permission` - Custom network permission. The structure of `netgrp_permission` block is documented below.
* `sysgrp_permission` - Custom system permission. The structure of `sysgrp_permission` block is documented below.
* `fwgrp_permission` - Custom firewall permission. The structure of `fwgrp_permission` block is documented below.
* `loggrp_permission` - Custom Log & Report permission. The structure of `loggrp_permission` block is documented below.
* `utmgrp_permission` - Custom Security Profile permissions. The structure of `utmgrp_permission` block is documented below.
* `admintimeout_override` - Enable/disable overriding the global administrator idle timeout.
* `admintimeout` - Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
* `cli_diagnose` - Enable/disable permission to run diagnostic commands.
* `cli_get` - Enable/disable permission to run get commands.
* `cli_show` - Enable/disable permission to run show commands.
* `cli_exec` - Enable/disable permission to run execute commands.
* `cli_config` - Enable/disable permission to run config commands.
* `system_diagnostics` - Enable/disable permission to run system diagnostic commands.
* `system_execute_ssh` - Enable/disable permission to execute SSH commands.
* `system_execute_telnet` - Enable/disable permission to execute TELNET commands.

The `netgrp_permission` block contains:

* `cfg` - Network Configuration.
* `packet_capture` - Packet Capture Configuration.
* `route_cfg` - Router Configuration.

The `sysgrp_permission` block contains:

* `admin` - Administrator Users.
* `upd` - FortiGuard Updates.
* `cfg` - System Configuration.
* `mnt` - Maintenance.

The `fwgrp_permission` block contains:

* `policy` - Policy Configuration.
* `address` - Address Configuration.
* `service` - Service Configuration.
* `schedule` - Schedule Configuration.
* `others` - Other Firewall Configuration.

The `loggrp_permission` block contains:

* `config` - Log & Report configuration.
* `data_access` - Log & Report Data Access.
* `report_access` - Log & Report Report Access.
* `threat_weight` - Log & Report Threat Weight.

The `utmgrp_permission` block contains:

* `antivirus` - Antivirus profiles and settings.
* `ips` - IPS profiles and settings.
* `webfilter` - Web Filter profiles and settings.
* `emailfilter` - AntiSpam filter and settings.
* `data_leak_prevention` - DLP profiles and settings.
* `spamfilter` - AntiSpam filter and settings.
* `data_loss_prevention` - DLP profiles and settings.
* `file_filter` - File-filter profiles and settings.
* `application_control` - Application Control profiles and settings.
* `icap` - ICAP profiles and settings.
* `voip` - VoIP profiles and settings.
* `waf` - Web Application Firewall profiles and settings.
* `dnsfilter` - DNS Filter profiles and settings.
* `endpoint_control` - FortiClient Profiles.
* `videofilter` - Video filter profiles and settings.
* `virtual_patch` - Virtual patch profiles and settings.
* `casb` - Inline CASB filter profile and settings

