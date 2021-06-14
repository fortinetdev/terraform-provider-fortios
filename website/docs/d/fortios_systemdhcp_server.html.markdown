---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemdhcp_server"
description: |-
  Get information on an fortios systemdhcp server.
---

# Data Source: fortios_systemdhcp_server
Use this data source to get information on an fortios systemdhcp server

## Argument Reference

* `fosid` - (Required) Specify the fosid of the desired systemdhcp server.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `fosid` - ID.
* `status` - Enable/disable this DHCP configuration.
* `lease_time` - Lease time in seconds, 0 means unlimited.
* `mac_acl_default_action` - MAC access control default action (allow or block assigning IP settings).
* `forticlient_on_net_status` - Enable/disable FortiClient-On-Net service for this DHCP server.
* `dns_service` - Options for assigning DNS servers to DHCP clients.
* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `dns_server3` - DNS server 3.
* `dns_server4` - DNS server 4.
* `wifi_ac_service` - Options for assigning WiFi Access Controllers to DHCP clients
* `wifi_ac1` - WiFi Access Controller 1 IP address (DHCP option 138, RFC 5417).
* `wifi_ac2` - WiFi Access Controller 2 IP address (DHCP option 138, RFC 5417).
* `wifi_ac3` - WiFi Access Controller 3 IP address (DHCP option 138, RFC 5417).
* `ntp_service` - Options for assigning Network Time Protocol (NTP) servers to DHCP clients.
* `ntp_server1` - NTP server 1.
* `ntp_server2` - NTP server 2.
* `ntp_server3` - NTP server 3.
* `domain` - Domain name suffix for the IP addresses that the DHCP server assigns to clients.
* `wins_server1` - WINS server 1.
* `wins_server2` - WINS server 2.
* `default_gateway` - Default gateway IP address assigned by the DHCP server.
* `next_server` - IP address of a server (for example, a TFTP sever) that DHCP clients can download a boot file from.
* `netmask` - Netmask assigned by the DHCP server.
* `interface` - DHCP server can assign IP configurations to clients connected to this interface.
* `ip_range` - DHCP IP range configuration. The structure of `ip_range` block is documented below.
* `timezone_option` - Options for the DHCP server to set the client's time zone.
* `timezone` - Select the time zone to be assigned to DHCP clients.
* `tftp_server` - One or more hostnames or IP addresses of the TFTP servers in quotes separated by spaces. The structure of `tftp_server` block is documented below.
* `filename` - Name of the boot file on the TFTP server.
* `options` - DHCP options. The structure of `options` block is documented below.
* `server_type` - DHCP server can be a normal DHCP server or an IPsec DHCP server.
* `ip_mode` - Method used to assign client IP.
* `conflicted_ip_timeout` - Time in seconds to wait after a conflicted IP address is removed from the DHCP range before it can be reused.
* `ipsec_lease_hold` - DHCP over IPsec leases expire this many seconds after tunnel down (0 to disable forced-expiry).
* `auto_configuration` - Enable/disable auto configuration.
* `dhcp_settings_from_fortiipam` - Enable/disable populating of DHCP server settings from FortiIPAM.
* `auto_managed_status` - Enable/disable use of this DHCP server once this interface has been assigned an IP address from FortiIPAM.
* `ddns_update` - Enable/disable DDNS update for DHCP.
* `ddns_update_override` - Enable/disable DDNS update override for DHCP.
* `ddns_server_ip` - DDNS server IP.
* `ddns_zone` - Zone of your domain name (ex. DDNS.com).
* `ddns_auth` - DDNS authentication mode.
* `ddns_keyname` - DDNS update key name.
* `ddns_key` - DDNS update key (base 64 encoding).
* `ddns_ttl` - TTL.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served.
* `vci_string` - One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
* `exclude_range` - Exclude one or more ranges of IP addresses from being assigned to clients. The structure of `exclude_range` block is documented below.
* `reserved_address` - Options for the DHCP server to assign IP settings to specific MAC addresses. The structure of `reserved_address` block is documented below.

The `ip_range` block contains:

* `id` - ID.
* `start_ip` - Start of IP range.
* `end_ip` - End of IP range.

The `tftp_server` block contains:

* `tftp_server` - TFTP server.

The `options` block contains:

* `id` - ID.
* `code` - DHCP option code.
* `type` - DHCP option type.
* `value` - DHCP option value.
* `ip` - DHCP option IPs.

The `vci_string` block contains:

* `vci_string` - VCI strings.

The `exclude_range` block contains:

* `id` - ID.
* `start_ip` - Start of IP range.
* `end_ip` - End of IP range.

The `reserved_address` block contains:

* `id` - ID.
* `type` - DHCP reserved-address type.
* `ip` - IP address to be reserved for the MAC address.
* `mac` - MAC address of the client that will get the reserved IP address.
* `action` - Options for the DHCP server to configure the client with the reserved MAC address.
* `circuit_id_type` - DHCP option type.
* `circuit_id` - Option 82 circuit-ID of the client that will get the reserved IP address.
* `remote_id_type` - DHCP option type.
* `remote_id` - Option 82 remote-ID of the client that will get the reserved IP address.
* `description` - Description.

