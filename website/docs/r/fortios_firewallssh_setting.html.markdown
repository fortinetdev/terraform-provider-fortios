---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallssh_setting"
description: |-
  SSH proxy settings.
---

# fortios_firewallssh_setting
SSH proxy settings.

## Example Usage

```hcl
resource "fortios_firewallssh_setting" "trname" {
  caname                = "Fortinet_SSH_CA"
  host_trusted_checking = "enable"
  hostkey_dsa1024       = "Fortinet_SSH_DSA1024"
  hostkey_ecdsa256      = "Fortinet_SSH_ECDSA256"
  hostkey_ecdsa384      = "Fortinet_SSH_ECDSA384"
  hostkey_ecdsa521      = "Fortinet_SSH_ECDSA521"
  hostkey_ed25519       = "Fortinet_SSH_ED25519"
  hostkey_rsa2048       = "Fortinet_SSH_RSA2048"
  untrusted_caname      = "Fortinet_SSH_CA_Untrusted"
}
```

## Argument Reference

The following arguments are supported:

* `caname` - CA certificate used by SSH Inspection.
* `untrusted_caname` - Untrusted CA certificate used by SSH Inspection.
* `hostkey_rsa2048` - RSA certificate used by SSH proxy.
* `hostkey_dsa1024` - DSA certificate used by SSH proxy.
* `hostkey_ecdsa256` - ECDSA nid256 certificate used by SSH proxy.
* `hostkey_ecdsa384` - ECDSA nid384 certificate used by SSH proxy.
* `hostkey_ecdsa521` - ECDSA nid384 certificate used by SSH proxy.
* `hostkey_ed25519` - ED25519 hostkey used by SSH proxy.
* `host_trusted_checking` - Enable/disable host trusted checking. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

FirewallSsh Setting can be imported using any of these accepted formats:
```
$ terraform import fortios_firewallssh_setting.labelname FirewallSshSetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewallssh_setting.labelname FirewallSshSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
