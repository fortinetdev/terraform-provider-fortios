---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceextension"
description: |-
  Configure Internet Services Extension.
---

# fortios_firewall_internetserviceextension
Configure Internet Services Extension.

## Example Usage

```hcl
resource "fortios_firewall_internetserviceextension" "trname" {
  comment = "EIWE"
  fosid   = 65536
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - Internet Service ID in the Internet Service database.
* `comment` - Comment.
* `entry` - Entries added to the Internet Service extension database.
* `disable_entry` - Disable entries in the Internet Service database.

The `entry` block supports:

* `id` - Entry ID(1-255).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port_range` - Port ranges in the custom entry.
* `dst` - Destination address or address group name.

The `port_range` block supports:

* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 to 65535).
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 to 65535).

The `dst` block supports:

* `name` - Select the destination address or address group object from available options.

The `disable_entry` block supports:

* `id` - Disable entry ID.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port` - Integer value for the TCP/IP port (0 - 65535).
* `ip_range` - IP ranges in the disable entry.

The `ip_range` block supports:

* `id` - Disable entry range ID.
* `start_ip` - Start IP address.
* `end_ip` - End IP address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceExtension can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_internetserviceextension.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```