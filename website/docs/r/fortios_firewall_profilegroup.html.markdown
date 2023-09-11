---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_profilegroup"
description: |-
  Configure profile groups.
---

# fortios_firewall_profilegroup
Configure profile groups.

## Example Usage

```hcl
resource "fortios_firewall_profilegroup" "trname" {
  name                     = "profilegroup1"
  profile_protocol_options = "default"
  ssl_ssh_profile          = "deep-inspection"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Profile group name.
* `av_profile` - Name of an existing Antivirus profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `emailfilter_profile` - Name of an existing email filter profile.
* `dlp_profile` - Name of an existing DLP profile.
* `spamfilter_profile` - Name of an existing Spam filter profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `file_filter_profile` - Name of an existing file-filter profile.
* `ips_sensor` - Name of an existing IPS sensor.
* `application_list` - Name of an existing Application list.
* `voip_profile` - Name of an existing VoIP profile.
* `ips_voip_filter` - Name of an existing VoIP (ips) profile.
* `sctp_filter_profile` - Name of an existing SCTP filter profile.
* `virtual_patch_profile` - Name of an existing virtual-patch profile.
* `icap_profile` - Name of an existing ICAP profile.
* `cifs_profile` - Name of an existing CIFS profile.
* `videofilter_profile` - Name of an existing VideoFilter profile.
* `waf_profile` - Name of an existing Web application firewall profile.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `casb_profile` - Name of an existing CASB profile.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ProfileGroup can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_profilegroup.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_profilegroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
