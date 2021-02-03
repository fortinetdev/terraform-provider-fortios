---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_externalresource"
description: |-
  Get information on an fortios system externalresource.
---

# Data Source: fortios_system_externalresource
Use this data source to get information on an fortios system externalresource

## Argument Reference

* `name` - (Required) Specify the name of the desired system externalresource.

## Attribute Reference

The following attributes are exported:

* `name` - External resource name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `status` - Enable/disable user resource.
* `type` - User resource type.
* `category` - User resource category.
* `username` - HTTP basic authentication user name.
* `password` - HTTP basic authentication password.
* `comments` - Comment.
* `resource` - URI of external resource.
* `user_agent` - Override HTTP User-Agent header used when retrieving this external resource.
* `refresh_rate` - Time interval to refresh external resource (1 - 43200 min, default = 5 min).
* `source_ip` - Source IPv4 address used to communicate with server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.

