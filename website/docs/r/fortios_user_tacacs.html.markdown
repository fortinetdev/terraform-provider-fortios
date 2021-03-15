---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_tacacs"
description: |-
  Configure TACACS+ server entries.
---

# fortios_user_tacacs
Configure TACACS+ server entries.

## Example Usage

```hcl
resource "fortios_user_tacacs" "trname" {
  authen_type   = "auto"
  authorization = "disable"
  name          = "eiew"
  port          = 2342
  server        = "1.1.1.1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - TACACS+ server entry name.
* `server` - Primary TACACS+ server CN domain name or IP address.
* `secondary_server` - Secondary TACACS+ server CN domain name or IP address.
* `tertiary_server` - Tertiary TACACS+ server CN domain name or IP address.
* `port` - Port number of the TACACS+ server.
* `key` - Key to access the primary server.
* `secondary_key` - Key to access the secondary server.
* `tertiary_key` - Key to access the tertiary server.
* `authen_type` - Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
* `authorization` - Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
* `source_ip` - source IP for communications to TACACS+ server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Tacacs can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_tacacs.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
