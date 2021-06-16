---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_threatweight"
description: |-
  Configure threat weight settings.
---

# fortios_log_threatweight
Configure threat weight settings.

## Example Usage

```hcl
resource "fortios_log_threatweight" "trname" {
  blocked_connection = "high"
  failed_connection  = "low"
  status             = "enable"
  url_block_detected = "high"

  application {
    category = 2
    id       = 1
    level    = "low"
  }
  application {
    category = 6
    id       = 2
    level    = "medium"
  }

  ips {
    critical_severity = "critical"
    high_severity     = "high"
    info_severity     = "disable"
    low_severity      = "low"
    medium_severity   = "medium"
  }

  level {
    critical = 50
    high     = 30
    low      = 5
    medium   = 10
  }

  malware {
    botnet_connection          = "critical"
    command_blocked            = "disable"
    content_disarm             = "medium"
    file_blocked               = "low"
    mimefragmented             = "disable"
    oversized                  = "disable"
    switch_proto               = "disable"
    virus_file_type_executable = "medium"
    virus_infected             = "critical"
    virus_outbreak_prevention  = "critical"
    virus_scan_error           = "high"
  }

  web {
    category = 26
    id       = 1
    level    = "high"
  }
  web {
    category = 61
    id       = 2
    level    = "high"
  }
  web {
    category = 86
    id       = 3
    level    = "high"
  }
  web {
    category = 1
    id       = 4
    level    = "medium"
  }
  web {
    category = 3
    id       = 5
    level    = "medium"
  }
  web {
    category = 4
    id       = 6
    level    = "medium"
  }
  web {
    category = 5
    id       = 7
    level    = "medium"
  }
  web {
    category = 6
    id       = 8
    level    = "medium"
  }
  web {
    category = 12
    id       = 9
    level    = "medium"
  }
  web {
    category = 59
    id       = 10
    level    = "medium"
  }
  web {
    category = 62
    id       = 11
    level    = "medium"
  }
  web {
    category = 83
    id       = 12
    level    = "medium"
  }
  web {
    category = 72
    id       = 13
    level    = "low"
  }
  web {
    category = 14
    id       = 14
    level    = "low"
  }
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable the threat weight feature. Valid values: `enable`, `disable`.
* `level` - Score mapping for threat weight levels. The structure of `level` block is documented below.
* `blocked_connection` - Threat weight score for blocked connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `failed_connection` - Threat weight score for failed connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `url_block_detected` - Threat weight score for URL blocking. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `botnet_connection_detected` - Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `malware` - Anti-virus malware threat weight settings. The structure of `malware` block is documented below.
* `ips` - IPS threat weight settings. The structure of `ips` block is documented below.
* `web` - Web filtering threat weight settings. The structure of `web` block is documented below.
* `geolocation` - Geolocation-based threat weight settings. The structure of `geolocation` block is documented below.
* `application` - Application-control threat weight settings. The structure of `application` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `level` block supports:

* `low` - Low level score value (1 - 100).
* `medium` - Medium level score value (1 - 100).
* `high` - High level score value (1 - 100).
* `critical` - Critical level score value (1 - 100).

The `malware` block supports:

* `virus_infected` - Threat weight score for virus (infected) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `file_blocked` - Threat weight score for blocked file detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `command_blocked` - Threat weight score for blocked command detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `oversized` - Threat weight score for oversized file detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `virus_scan_error` - Threat weight score for virus (scan error) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `switch_proto` - Threat weight score for switch proto detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `mimefragmented` - Threat weight score for mimefragmented detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `virus_file_type_executable` - Threat weight score for virus (filetype executable) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `virus_outbreak_prevention` - Threat weight score for virus (outbreak prevention) event. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `botnet_connection` - Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `content_disarm` - Threat weight score for virus (content disarm) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `malware_list` - Threat weight score for virus (malware list) detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `fsa_malicious` - Threat weight score for FortiSandbox malicious malware detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `fsa_high_risk` - Threat weight score for FortiSandbox high risk malware detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `fsa_medium_risk` - Threat weight score for FortiSandbox medium risk malware detected. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

The `ips` block supports:

* `info_severity` - Threat weight score for IPS info severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `low_severity` - Threat weight score for IPS low severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `medium_severity` - Threat weight score for IPS medium severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `high_severity` - Threat weight score for IPS high severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
* `critical_severity` - Threat weight score for IPS critical severity events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

The `web` block supports:

* `id` - Entry ID.
* `category` - Threat weight score for web category filtering matches.
* `level` - Threat weight score for web category filtering matches. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

The `geolocation` block supports:

* `id` - Entry ID.
* `country` - Country code.
* `level` - Threat weight score for Geolocation-based events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.

The `application` block supports:

* `id` - Entry ID.
* `category` - Application category.
* `level` - Threat weight score for Application events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log ThreatWeight can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_log_threatweight.labelname LogThreatWeight
$ unset "FORTIOS_IMPORT_TABLE"
```
