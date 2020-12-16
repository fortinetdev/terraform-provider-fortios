---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_wtp"
description: |-
  Configure Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.
---

# fortios_wirelesscontroller_wtp
Configure Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.

## Argument Reference

The following arguments are supported:

* `wtp_id` - WTP ID.
* `index` - Index (0 - 4294967295).
* `admin` - Configure how the FortiGate operating as a wireless controller discovers and manages this WTP, AP or FortiAP.
* `name` - WTP, AP or FortiAP configuration name.
* `location` - Field for describing the physical location of the WTP, AP or FortiAP.
* `region` - Region name WTP is associated with.
* `region_x` - Relative horizontal region coordinate (between 0 and 1).
* `region_y` - Relative vertical region coordinate (between 0 and 1).
* `wtp_profile` - (Required) WTP profile name to apply to this WTP, AP or FortiAP.
* `wtp_mode` - WTP, AP, or FortiAP operating mode; normal (by default) or remote. A tunnel mode SSID can be assigned to an AP in normal mode but not remote mode, while a local-bridge mode SSID can be assigned to an AP in either normal mode or remote mode.
* `bonjour_profile` - Bonjour profile name.
* `override_led_state` - Enable to override the profile LED state setting for this FortiAP. You must enable this option to use the led-state command to turn off the FortiAP's LEDs.
* `led_state` - Enable to allow the FortiAPs LEDs to light. Disable to keep the LEDs off. You may want to keep the LEDs off so they are not distracting in low light areas etc.
* `override_wan_port_mode` - Enable/disable overriding the wan-port-mode in the WTP profile.
* `wan_port_mode` - Enable/disable using the FortiAP WAN port as a LAN port.
* `override_ip_fragment` - Enable/disable overriding the WTP profile IP fragment prevention setting.
* `ip_fragment_preventing` - Method by which IP fragmentation is prevented for CAPWAP tunneled control and data packets (default = tcp-mss-adjust).
* `tun_mtu_uplink` - Uplink tunnel maximum transmission unit (MTU) in octets (eight-bit bytes). Set the value to either 0 (by default), 576, or 1500.
* `tun_mtu_downlink` - Downlink tunnel MTU in octets. Set the value to either 0 (by default), 576, or 1500.
* `override_split_tunnel` - Enable/disable overriding the WTP profile split tunneling setting.
* `split_tunneling_acl_path` - Split tunneling ACL path is local/tunnel.
* `split_tunneling_acl_local_ap_subnet` - Enable/disable automatically adding local subnetwork of FortiAP to split-tunneling ACL (default = disable).
* `split_tunneling_acl` - Split tunneling ACL filter list. The structure of `split_tunneling_acl` block is documented below.
* `override_lan` - Enable to override the WTP profile LAN port setting.
* `lan` - WTP LAN port mapping. The structure of `lan` block is documented below.
* `override_allowaccess` - Enable to override the WTP profile management access configuration.
* `allowaccess` - Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space.
* `override_login_passwd_change` - Enable to override the WTP profile login-password (administrator password) setting.
* `login_passwd_change` - Change or reset the administrator password of a managed WTP, FortiAP or AP (yes, default, or no, default = no).
* `login_passwd` - Set the managed WTP, FortiAP, or AP's administrator password.
* `radio_1` - Configuration options for radio 1. The structure of `radio_1` block is documented below.
* `radio_2` - Configuration options for radio 2. The structure of `radio_2` block is documented below.
* `image_download` - Enable/disable WTP image download.
* `mesh_bridge_enable` - Enable/disable mesh Ethernet bridge when WTP is configured as a mesh branch/leaf AP.
* `coordinate_latitude` - WTP latitude coordinate.
* `coordinate_longitude` - WTP longitude coordinate.

The `split_tunneling_acl` block supports:

* `id` - ID.
* `dest_ip` - Destination IP and mask for the split-tunneling subnet.

The `lan` block supports:

* `port_mode` - LAN port mode.
* `port_ssid` - Bridge LAN port to SSID.
* `port1_mode` - LAN port 1 mode.
* `port1_ssid` - Bridge LAN port 1 to SSID.
* `port2_mode` - LAN port 2 mode.
* `port2_ssid` - Bridge LAN port 2 to SSID.
* `port3_mode` - LAN port 3 mode.
* `port3_ssid` - Bridge LAN port 3 to SSID.
* `port4_mode` - LAN port 4 mode.
* `port4_ssid` - Bridge LAN port 4 to SSID.
* `port5_mode` - LAN port 5 mode.
* `port5_ssid` - Bridge LAN port 5 to SSID.
* `port6_mode` - LAN port 6 mode.
* `port6_ssid` - Bridge LAN port 6 to SSID.
* `port7_mode` - LAN port 7 mode.
* `port7_ssid` - Bridge LAN port 7 to SSID.
* `port8_mode` - LAN port 8 mode.
* `port8_ssid` - Bridge LAN port 8 to SSID.

The `radio_1` block supports:

* `radio_id` - radio-id
* `override_band` - Enable to override the WTP profile band setting.
* `band` - WiFi band that Radio 1 operates on.
* `override_analysis` - Enable to override the WTP profile spectrum analysis configuration.
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.
* `override_txpower` - Enable to override the WTP profile power level configuration.
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).
* `auto_power_high` - Automatic transmission power high limit in decibels (dB) of the measured power referenced to one milliwatt (mW), or dBm (10 - 17 dBm, default = 17).
* `auto_power_low` - Automatic transmission power low limit in dBm (the actual range of transmit power depends on the AP platform type).
* `power_level` - Radio power level as a percentage of the maximum transmit power (0 - 100, default = 100).
* `override_vaps` - Enable to override WTP profile Virtual Access Point (VAP) settings.
* `vap_all` - Enable/disable the automatic inheritance of all Virtual Access Points (VAPs) (default = enable).
* `vaps` - Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.
* `override_channel` - Enable to override WTP profile channel settings.
* `channel` - Selected list of wireless radio channels. The structure of `channel` block is documented below.

The `vaps` block supports:

* `name` - Virtual Access Point (VAP) name.

The `channel` block supports:

* `chan` - Channel number.

The `radio_2` block supports:

* `radio_id` - radio-id
* `override_band` - Enable to override the WTP profile band setting.
* `band` - WiFi band that Radio 1 operates on.
* `override_analysis` - Enable to override the WTP profile spectrum analysis configuration.
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.
* `override_txpower` - Enable to override the WTP profile power level configuration.
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).
* `auto_power_high` - Automatic transmission power high limit in decibels (dB) of the measured power referenced to one milliwatt (mW), or dBm (10 - 17 dBm, default = 17).
* `auto_power_low` - Automatic transmission power low limit in dBm (the actual range of transmit power depends on the AP platform type).
* `power_level` - Radio power level as a percentage of the maximum transmit power (0 - 100, default = 100).
* `override_vaps` - Enable to override WTP profile Virtual Access Point (VAP) settings.
* `vap_all` - Enable/disable the automatic inheritance of all Virtual Access Points (VAPs) (default = enable).
* `vaps` - Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.
* `override_channel` - Enable to override WTP profile channel settings.
* `channel` - Selected list of wireless radio channels. The structure of `channel` block is documented below.

The `vaps` block supports:

* `name` - Virtual Access Point (VAP) name.

The `channel` block supports:

* `chan` - Channel number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{wtp_id}}.

## Import

WirelessController Wtp can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_wtp.labelname {{wtp_id}}
$ unset "FORTIOS_IMPORT_TABLE"
```
