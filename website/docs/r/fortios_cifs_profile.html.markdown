---
subcategory: "FortiGate CIFS"
layout: "fortios"
page_title: "FortiOS: fortios_cifs_profile"
description: |-
  Configure CIFS profile.
---

# fortios_cifs_profile
Configure CIFS profile. Applies to FortiOS Version `<= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `name` - Profile name.
* `server_credential_type` - CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
* `file_filter` - File filter. The structure of `file_filter` block is documented below.
* `domain_controller` - Domain for which to decrypt CIFS traffic.
* `server_keytab` - Server keytab. The structure of `server_keytab` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `file_filter` block supports:

* `status` - Enable/disable file filter. Valid values: `enable`, `disable`.
* `log` - Enable/disable file filter logging. Valid values: `enable`, `disable`.
* `entries` - File filter entries. The structure of `entries` block is documented below.

The `entries` block supports:

* `filter` - Add a file filter.
* `comment` - Comment.
* `action` - Action taken for matched file. Valid values: `log`, `block`.
* `direction` - Match files transmitted in the session's originating or reply direction. Valid values: `incoming`, `outgoing`, `any`.
* `file_type` - Select file type. The structure of `file_type` block is documented below.

The `file_type` block supports:

* `name` - File type name.

The `server_keytab` block supports:

* `principal` - Service principal.  For example, "host/cifsserver.example.com@example.com".
* `keytab` - Base64 encoded keytab file containing credential of the server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Cifs Profile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_cifs_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
