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
* `mail` - Replacement message table entries. The structure of `mail` block is documented below.
* `http` - Replacement message table entries. The structure of `http` block is documented below.
* `webproxy` - Replacement message table entries. The structure of `webproxy` block is documented below.
* `ftp` - Replacement message table entries. The structure of `ftp` block is documented below.
* `nntp` - Replacement message table entries. The structure of `nntp` block is documented below.
* `fortiguard_wf` - Replacement message table entries. The structure of `fortiguard_wf` block is documented below.
* `spam` - Replacement message table entries. The structure of `spam` block is documented below.
* `alertmail` - Replacement message table entries. The structure of `alertmail` block is documented below.
* `admin` - Replacement message table entries. The structure of `admin` block is documented below.
* `auth` - Replacement message table entries. The structure of `auth` block is documented below.
* `sslvpn` - Replacement message table entries. The structure of `sslvpn` block is documented below.
* `ec` - Replacement message table entries. The structure of `ec` block is documented below.
* `device_detection_portal` - Replacement message table entries. The structure of `device_detection_portal` block is documented below.
* `nac_quar` - Replacement message table entries. The structure of `nac_quar` block is documented below.
* `traffic_quota` - Replacement message table entries. The structure of `traffic_quota` block is documented below.
* `utm` - Replacement message table entries. The structure of `utm` block is documented below.
* `custom_message` - Replacement message table entries. The structure of `custom_message` block is documented below.
* `icap` - Replacement message table entries. The structure of `icap` block is documented below.
* `automation` - Replacement message table entries. The structure of `automation` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `mail` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `http` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `webproxy` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `ftp` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `nntp` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `fortiguard_wf` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `spam` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `alertmail` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `admin` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `auth` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `sslvpn` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `ec` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

The `device_detection_portal` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `nac_quar` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `traffic_quota` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `utm` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `custom_message` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `icap` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.

The `automation` block supports:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag. Valid values: `none`, `text`, `html`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System ReplacemsgGroup can be imported using any of these accepted formats:
```
$ terraform import fortios_system_replacemsggroup.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_replacemsggroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
