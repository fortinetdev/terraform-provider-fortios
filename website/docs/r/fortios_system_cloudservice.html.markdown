---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_cloudservice"
description: |-
  Configure system cloud service.
---

# fortios_system_cloudservice
Configure system cloud service. Applies to FortiOS Version `>= 7.6.3`.

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `vendor` - Cloud service vendor. Valid values: `unknown`, `google-cloud-kms`.
* `traffic_vdom` - Vdom used to communicate with cloud service.
* `gck_service_account` - Service account (e.g. "account-id@sampledomain.com").
* `gck_private_key` - Service account private key in PEM format (e.g. "-----BEGIN PRIVATE KEY-----
...").
* `gck_keyid` - Key id, also referred as "kid".
* `gck_access_token_lifetime` - Lifetime of automatically generated access tokens in minutes (default is 60 minutes).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CloudService can be imported using any of these accepted formats:
```
$ terraform import fortios_system_cloudservice.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_cloudservice.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
