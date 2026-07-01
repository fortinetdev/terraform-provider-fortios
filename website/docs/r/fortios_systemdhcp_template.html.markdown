---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemdhcp_template"
description: |-
  Configure DHCP server templates.
---

# fortios_systemdhcp_template
Configure DHCP server templates. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `name` - DHCP server template name.
* `lease_time` - Lease time in seconds, 0 means unlimited.
* `mac_acl_default_action` - MAC access control default action (allow or block assigning IP settings). Valid values: `assign`, `block`.
* `forticlient_on_net_status` - Enable/disable FortiClient-On-Net service for this DHCP server. Valid values: `disable`, `enable`.
* `dns_service` - Options for assigning DNS servers to DHCP clients. Valid values: `local`, `default`, `specify`.
* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `dns_server3` - DNS server 3.
* `dns_server4` - DNS server 4.
* `wifi_ac_service` - Options for assigning WiFi access controllers to DHCP clients. Valid values: `specify`, `local`.
* `wifi_ac1` - WiFi Access Controller 1 IP address (DHCP option 138, RFC 5417).
* `wifi_ac2` - WiFi Access Controller 2 IP address (DHCP option 138, RFC 5417).
* `wifi_ac3` - WiFi Access Controller 3 IP address (DHCP option 138, RFC 5417).
* `ntp_service` - Options for assigning Network Time Protocol (NTP) servers to DHCP clients. Valid values: `local`, `default`, `specify`.
* `ntp_server1` - NTP server 1.
* `ntp_server2` - NTP server 2.
* `ntp_server3` - NTP server 3.
* `domain` - Domain name suffix for the IP addresses that the DHCP server assigns to clients.
* `wins_server1` - WINS server 1.
* `wins_server2` - WINS server 2.
* `next_server` - IP address of a server, such as a TFTP server, from which DHCP clients can download a boot file.
* `reserve_extra_addresses` - Enable/disable reservation of the extra IP addresses in the subnet. Valid values: `disable`, `enable`.
* `ip_range` - DHCP IP range configuration. The structure of `ip_range` block is documented below.
* `timezone_option` - Options for the DHCP server to set the client's time zone. Valid values: `disable`, `default`, `specify`.
* `timezone` - Select the time zone to be assigned to DHCP clients.
* `tftp_server` - One or more hostnames or IP addresses of the TFTP servers in quotes separated by spaces. The structure of `tftp_server` block is documented below.
* `filename` - Name of the boot file on the TFTP server.
* `options` - DHCP options. The structure of `options` block is documented below.
* `server_type` - DHCP server can be a normal DHCP server or an IPsec DHCP server. Valid values: `regular`, `ipsec`.
* `conflicted_ip_timeout` - Time in seconds to wait after a conflicted IP address is removed from the DHCP range before it can be reused.
* `ipsec_lease_hold` - DHCP over IPsec leases expire this many seconds after tunnel down (0 to disable forced-expiry).
* `auto_configuration` - Enable/disable auto configuration. Valid values: `disable`, `enable`.
* `ddns_update` - Enable/disable DDNS update for DHCP. Valid values: `disable`, `enable`.
* `ddns_update_override` - Enable/disable DDNS update override for DHCP. Valid values: `disable`, `enable`.
* `ddns_server_ip` - DDNS server IP.
* `ddns_zone` - Zone of your domain name (ex. DDNS.com).
* `ddns_auth` - DDNS authentication mode. Valid values: `disable`, `tsig`.
* `ddns_keyname` - DDNS update key name.
* `ddns_key` - DDNS update key (base 64 encoding).
* `ddns_ttl` - TTL.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served. Valid values: `disable`, `enable`.
* `vci_string` - One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
* `exclude_range` - Exclude one or more ranges of IP addresses from being assigned to clients. The structure of `exclude_range` block is documented below.
* `shared_subnet` - Enable/disable shared subnet. Valid values: `disable`, `enable`.
* `relay_agent` - Relay agent IP.
* `reserved_address` - Options for the DHCP server to assign IP settings to specific MAC addresses. The structure of `reserved_address` block is documented below.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `fabric_force_sync` - Enable/disable forced synchronization of configuration objects from the root FortiGate unit to the downstream devices.  Configuration conflict check is skipped. Valid values: `enable`, `disable`.
* `fabric_object_source` - Source of truth for fabric object. Valid values: `member`, `local`, `root`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.

The `ip_range` block supports:

* `id` - ID.
* `ip_count` - Number of IP addresses to include in the range.
* `reserve` - Enable/disable address reservation for use without DHCP. Valid values: `disable`, `enable`.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range. Valid values: `disable`, `enable`.
* `vci_string` - One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
* `uci_match` - Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this range. Valid values: `disable`, `enable`.
* `uci_string` - One or more UCI strings in quotes separated by spaces. The structure of `uci_string` block is documented below.
* `oui_match` - Enable/disable organizationally unique identifier (OUI) matching. When enabled only DHCP requests with a matching OUI are served with this range. Valid values: `disable`, `enable`.
* `oui_string` - One or more OUI strings in quotes separated by spaces (in format of xx:xx:xx). The structure of `oui_string` block is documented below.
* `lease_time` - Lease time in seconds, 0 means default lease time.
* `vendor` - Vendor this ip-range will be assigned to.

The `vci_string` block supports:

* `vci_string` - VCI strings.

The `uci_string` block supports:

* `uci_string` - UCI strings.

The `oui_string` block supports:

* `oui_string` - MAC OUI strings.

The `tftp_server` block supports:

* `tftp_server` - TFTP server.

The `options` block supports:

* `id` - ID.
* `code` - DHCP option code.
* `type` - DHCP option type. Valid values: `hex`, `string`, `ip`, `fqdn`.
* `value` - DHCP option value.
* `ip` - DHCP option IPs.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this option. Valid values: `disable`, `enable`.
* `vci_string` - One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
* `uci_match` - Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this option. Valid values: `disable`, `enable`.
* `uci_string` - One or more UCI strings in quotes separated by spaces. The structure of `uci_string` block is documented below.

The `vci_string` block supports:

* `vci_string` - VCI strings.

The `uci_string` block supports:

* `uci_string` - UCI strings.

The `vci_string` block supports:

* `vci_string` - VCI strings.

The `exclude_range` block supports:

* `id` - ID.
* `start_ip_index` - Start of IP range.
* `ip_count` - Number of IP addresses to include in the range.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range. Valid values: `disable`, `enable`.
* `vci_string` - One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
* `uci_match` - Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this range. Valid values: `disable`, `enable`.
* `uci_string` - One or more UCI strings in quotes separated by spaces. The structure of `uci_string` block is documented below.
* `oui_match` - Enable/disable organizationally unique identifier (OUI) matching. When enabled only DHCP requests with a matching OUI are served with this range. Valid values: `disable`, `enable`.
* `oui_string` - One or more OUI strings in quotes separated by spaces (in format of xx:xx:xx). The structure of `oui_string` block is documented below.
* `lease_time` - Lease time in seconds, 0 means default lease time.
* `vendor` - Vendor this ip-range will be assigned to.

The `vci_string` block supports:

* `vci_string` - VCI strings.

The `uci_string` block supports:

* `uci_string` - UCI strings.

The `oui_string` block supports:

* `oui_string` - MAC OUI strings.

The `reserved_address` block supports:

* `id` - ID.
* `type` - DHCP reserved-address type. Valid values: `mac`, `option82`.
* `ip_index` - Index of IP address to be reserved for the MAC address.
* `mac` - MAC address of the client that will get the reserved IP address.
* `action` - Options for the DHCP server to configure the client with the reserved MAC address. Valid values: `assign`, `block`, `reserved`.
* `circuit_id_type` - DHCP option type. Valid values: `hex`, `string`.
* `circuit_id` - Option 82 circuit-ID of the client that will get the reserved IP address.
* `remote_id_type` - DHCP option type. Valid values: `hex`, `string`.
* `remote_id` - Option 82 remote-ID of the client that will get the reserved IP address.
* `description` - Description.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemDhcp Template can be imported using any of these accepted formats:
```
$ terraform import fortios_systemdhcp_template.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_systemdhcp_template.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
