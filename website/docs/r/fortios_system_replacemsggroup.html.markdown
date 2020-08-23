---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_replacemsggroup"
description: |-
  Configure replacement message groups.
---

# fortios_system_replacemsggroup
Configure replacement message groups.

## Example Usage

```hcl
resource "fortios_system_replacemsggroup" "trname" {
  comment    = "sgh"
  group_type = "utm"
  name       = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Group name.
* `comment` - Comment.
* `group_type` - (Required) Group type.
* `mail` - Replacement message table entries.
* `http` - Replacement message table entries.
* `webproxy` - Replacement message table entries.
* `ftp` - Replacement message table entries.
* `nntp` - Replacement message table entries.
* `fortiguard_wf` - Replacement message table entries.
* `spam` - Replacement message table entries.
* `alertmail` - Replacement message table entries.
* `admin` - Replacement message table entries.
* `auth` - Replacement message table entries.
* `sslvpn` - Replacement message table entries.
* `ec` - Replacement message table entries.
* `device_detection_portal` - Replacement message table entries.
* `nac_quar` - Replacement message table entries.
* `traffic_quota` - Replacement message table entries.
* `utm` - Replacement message table entries.
* `custom_message` - Replacement message table entries.
* `icap` - Replacement message table entries.

The `mail` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `http` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `webproxy` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `ftp` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `nntp` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `fortiguard_wf` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `spam` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `alertmail` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `admin` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `auth` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `sslvpn` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `ec` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `device_detection_portal` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `nac_quar` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `traffic_quota` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `utm` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `custom_message` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `icap` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System ReplacemsgGroup can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_replacemsggroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
