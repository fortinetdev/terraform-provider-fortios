---
subcategory: "FortiGate WANOpt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_remotestorage"
description: |-
  Configure a remote cache device as Web cache storage.
---

# fortios_wanopt_remotestorage
Configure a remote cache device as Web cache storage.

## Example Usage

```hcl
resource "fortios_wanopt_remotestorage" "trname" {
  remote_cache_ip = "0.0.0.0"
  status          = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
* `local_cache_id` - ID that this device uses to connect to the remote device.
* `remote_cache_id` - ID of the remote device to which the device connects.
* `remote_cache_ip` - IP address of the remote device to which the device connects.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Wanopt RemoteStorage can be imported using any of these accepted formats:
```
$ terraform import fortios_wanopt_remotestorage.labelname WanoptRemoteStorage

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wanopt_remotestorage.labelname WanoptRemoteStorage
$ unset "FORTIOS_IMPORT_TABLE"
```
