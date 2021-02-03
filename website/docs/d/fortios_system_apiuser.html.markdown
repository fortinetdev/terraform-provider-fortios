---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_apiuser"
description: |-
  Get information on an fortios system apiuser.
---

# Data Source: fortios_system_apiuser
Use this data source to get information on an fortios system apiuser

## Argument Reference

* `name` - (Required) Specify the name of the desired system apiuser.

## Attribute Reference

The following attributes are exported:

* `name` - User name.
* `comments` - Comment.
* `api_key` - Admin user password.
* `accprofile` - Admin user access profile.
* `vdom` - Virtual domains. The structure of `vdom` block is documented below.
* `schedule` - Schedule name.
* `cors_allow_origin` - Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
* `peer_auth` - Enable/disable peer authentication.
* `peer_group` - Peer group name.
* `trusthost` - Trusthost. The structure of `trusthost` block is documented below.

The `vdom` block contains:

* `name` - Virtual domain name.

The `trusthost` block contains:

* `id` - Table ID.
* `type` - Trusthost type.
* `ipv4_trusthost` - IPv4 trusted host address.
* `ipv6_trusthost` - IPv6 trusted host address.

