---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_list"
description: |-
  Configure application control lists.
---

# fortios_application_list
Configure application control lists.

## Example Usage

```hcl
resource "fortios_application_list" "trname" {
  app_replacemsg             = "enable"
  deep_app_inspection        = "enable"
  enforce_default_app_port   = "disable"
  extended_log               = "disable"
  name                       = "1"
  options                    = "allow-dns"
  other_application_action   = "pass"
  other_application_log      = "disable"
  unknown_application_action = "pass"
  unknown_application_log    = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) List name.
* `comment` - comments
* `replacemsg_group` - Replacement message group.
* `extended_log` - Enable/disable extended logging. Valid values: `enable`, `disable`.
* `other_application_action` - Action for other applications. Valid values: `pass`, `block`.
* `app_replacemsg` - Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
* `other_application_log` - Enable/disable logging for other applications. Valid values: `disable`, `enable`.
* `enforce_default_app_port` - Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
* `force_inclusion_ssl_di_sigs` - Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
* `unknown_application_action` - Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
* `unknown_application_log` - Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
* `p2p_block_list` - P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
* `p2p_black_list` - P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
* `deep_app_inspection` - Enable/disable deep application inspection. Valid values: `disable`, `enable`.
* `options` - Basic application protocol signatures allowed by default. Valid values: `allow-dns`, `allow-icmp`, `allow-http`, `allow-ssl`, `allow-quic`.
* `entries` - Application list entries. The structure of `entries` block is documented below.
* `control_default_network_services` - Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
* `default_network_services` - Default network service entries. The structure of `default_network_services` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `id` - Entry ID.
* `risk` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
* `category` - Category ID list. The structure of `category` block is documented below.
* `sub_category` - Application Sub-category ID list. The structure of `sub_category` block is documented below.
* `application` - ID of allowed applications. The structure of `application` block is documented below.
* `protocols` - Application protocol filter.
* `vendor` - Application vendor filter.
* `technology` - Application technology filter.
* `behavior` - Application behavior filter.
* `popularity` - Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
* `exclusion` - ID of excluded applications. The structure of `exclusion` block is documented below.
* `parameters` - Application parameters. The structure of `parameters` block is documented below.
* `action` - Pass or block traffic, or reset connection for traffic from this application. Valid values: `pass`, `block`, `reset`.
* `log` - Enable/disable logging for this application list. Valid values: `disable`, `enable`.
* `log_packet` - Enable/disable packet logging. Valid values: `disable`, `enable`.
* `rate_count` - Count of the rate.
* `rate_duration` - Duration (sec) of the rate.
* `rate_mode` - Rate limit mode. Valid values: `periodical`, `continuous`.
* `rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`, `dhcp-client-mac`, `dns-domain`.
* `session_ttl` - Session TTL (0 = default).
* `shaper` - Traffic shaper.
* `shaper_reverse` - Reverse traffic shaper.
* `per_ip_shaper` - Per-IP traffic shaper.
* `quarantine` - Quarantine method. Valid values: `none`, `attacker`.
* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging. Valid values: `disable`, `enable`.

The `risk` block supports:

* `level` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).

The `category` block supports:

* `id` - Application category ID.

The `sub_category` block supports:

* `id` - Application sub-category ID.

The `application` block supports:

* `id` - Application IDs.

The `exclusion` block supports:

* `id` - Excluded application IDs.

The `parameters` block supports:

* `id` - Parameter ID.
* `members` - Parameter tuple members. The structure of `members` block is documented below.
* `value` - Parameter value.

The `members` block supports:

* `id` - Parameter.
* `name` - Parameter name.
* `value` - Parameter value.

The `default_network_services` block supports:

* `id` - Entry ID.
* `port` - Port number.
* `services` - Network protocols. Valid values: `http`, `ssh`, `telnet`, `ftp`, `dns`, `smtp`, `pop3`, `imap`, `snmp`, `nntp`, `https`.
* `violation_action` - Action for protocols not white listed under selected port. Valid values: `pass`, `monitor`, `block`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Application List can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_application_list.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
