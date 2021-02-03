---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemlldp_networkpolicy"
description: |-
  Configure LLDP network policy.
---

# fortios_systemlldp_networkpolicy
Configure LLDP network policy.

## Example Usage

```hcl
resource "fortios_systemlldp_networkpolicy" "trname" {
  comment = "test"
  name    = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) LLDP network policy name.
* `comment` - Comment.
* `voice` - Voice. The structure of `voice` block is documented below.
* `voice_signaling` - Voice signaling. The structure of `voice_signaling` block is documented below.
* `guest` - Guest. The structure of `guest` block is documented below.
* `guest_voice_signaling` - Guest Voice Signaling. The structure of `guest_voice_signaling` block is documented below.
* `softphone` - Softphone. The structure of `softphone` block is documented below.
* `video_conferencing` - Video Conferencing. The structure of `video_conferencing` block is documented below.
* `streaming_video` - Streaming Video. The structure of `streaming_video` block is documented below.
* `video_signaling` - Video Signaling. The structure of `video_signaling` block is documented below.

The `voice` block supports:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

The `voice_signaling` block supports:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

The `guest` block supports:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

The `guest_voice_signaling` block supports:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

The `softphone` block supports:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

The `video_conferencing` block supports:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

The `streaming_video` block supports:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

The `video_signaling` block supports:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemLldp NetworkPolicy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_systemlldp_networkpolicy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
