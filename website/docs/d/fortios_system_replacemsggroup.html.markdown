---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_replacemsggroup"
description: |-
  Get information on an fortios system replacemsggroup.
---

# Data Source: fortios_system_replacemsggroup
Use this data source to get information on an fortios system replacemsggroup

## Argument Reference

* `name` - (Required) Specify the name of the desired system replacemsggroup.

## Attribute Reference

The following attributes are exported:

* `name` - Group name.
* `comment` - Comment.
* `group_type` - Group type.
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

The `mail` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `http` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `webproxy` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `ftp` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `nntp` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `fortiguard_wf` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `spam` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `alertmail` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `admin` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `auth` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `sslvpn` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `ec` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `device_detection_portal` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `nac_quar` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `traffic_quota` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `utm` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `custom_message` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `icap` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

The `automation` block contains:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.

