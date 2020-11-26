---
subcategory: "FortiGate WANOpt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_contentdeliverynetworkrule"
description: |-
  Configure WAN optimization content delivery network rules.
---

# fortios_wanopt_contentdeliverynetworkrule
Configure WAN optimization content delivery network rules.

## Example Usage

```hcl
resource "fortios_wanopt_contentdeliverynetworkrule" "trname" {
  category               = "vcache"
  name                   = "contentdeliverynetworkrules1"
  request_cache_control  = "disable"
  response_cache_control = "disable"
  response_expires       = "enable"
  status                 = "enable"
  text_response_vcache   = "enable"
  updateserver           = "disable"

  host_domain_name_suffix {
    name = "kaf.com"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name of table.
* `comment` - Comment about this CDN-rule.
* `status` - Enable/disable WAN optimization content delivery network rules.
* `host_domain_name_suffix` - Suffix portion of the fully qualified domain name (eg. fortinet.com in "www.fortinet.com"). The structure of `host_domain_name_suffix` block is documented below.
* `category` - Content delivery network rule category.
* `request_cache_control` - Enable/disable HTTP request cache control.
* `response_cache_control` - Enable/disable HTTP response cache control.
* `response_expires` - Enable/disable HTTP response cache expires.
* `text_response_vcache` - Enable/disable caching of text responses.
* `updateserver` - Enable/disable update server.
* `rules` - WAN optimization content delivery network rule entries. The structure of `rules` block is documented below.

The `host_domain_name_suffix` block supports:

* `name` - Suffix portion of the fully qualified domain name.

The `rules` block supports:

* `name` - WAN optimization content delivery network rule name.
* `match_mode` - Match criteria for collecting content ID.
* `skip_rule_mode` - Skip mode when evaluating skip-rules.
* `match_entries` - List of entries to match. The structure of `match_entries` block is documented below.
* `skip_entries` - List of entries to skip. The structure of `skip_entries` block is documented below.
* `content_id` - Content ID settings. The structure of `content_id` block is documented below.

The `match_entries` block supports:

* `id` - Rule ID.
* `target` - Option in HTTP header or URL parameter to match.
* `pattern` - Pattern string for matching target (Referrer or URL pattern, eg. "a", "a*c", "*a*", "a*c*e", and "*"). The structure of `pattern` block is documented below.

The `pattern` block supports:

* `string` - Pattern strings.

The `skip_entries` block supports:

* `id` - Rule ID.
* `target` - Option in HTTP header or URL parameter to match.
* `pattern` - Pattern string for matching target (Referrer or URL pattern, eg. "a", "a*c", "*a*", "a*c*e", and "*"). The structure of `pattern` block is documented below.

The `pattern` block supports:

* `string` - Pattern strings.

The `content_id` block supports:

* `target` - Option in HTTP header or URL parameter to match.
* `start_str` - String from which to start search.
* `start_skip` - Number of characters in URL to skip after start-str has been matched.
* `start_direction` - Search direction from start-str match.
* `end_str` - String from which to end search.
* `end_skip` - Number of characters in URL to skip after end-str has been matched.
* `end_direction` - Search direction from end-str match.
* `range_str` - Name of content ID within the start string and end string.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Wanopt ContentDeliveryNetworkRule can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wanopt_contentdeliverynetworkrule.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
