---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemlldp_networkpolicy"
description: |-
  Get information on an fortios systemlldp networkpolicy.
---

# Data Source: fortios_systemlldp_networkpolicy
Use this data source to get information on an fortios systemlldp networkpolicy

## Argument Reference

* `name` - (Required) Specify the name of the desired systemlldp networkpolicy.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - LLDP network policy name.
* `comment` - Comment.
* `voice` - Voice. The structure of `voice` block is documented below.
* `voice_signaling` - Voice signaling. The structure of `voice_signaling` block is documented below.
* `guest` - Guest. The structure of `guest` block is documented below.
* `guest_voice_signaling` - Guest Voice Signaling. The structure of `guest_voice_signaling` block is documented below.
* `softphone` - Softphone. The structure of `softphone` block is documented below.
* `video_conferencing` - Video Conferencing. The structure of `video_conferencing` block is documented below.
* `streaming_video` - Streaming Video. The structure of `streaming_video` block is documented below.
* `video_signaling` - Video Signaling. The structure of `video_signaling` block is documented below.

The `voice` block contains:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

The `voice_signaling` block contains:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

The `guest` block contains:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

The `guest_voice_signaling` block contains:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

The `softphone` block contains:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

The `video_conferencing` block contains:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

The `streaming_video` block contains:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

The `video_signaling` block contains:

* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.

