---
subcategory: "FortiGate WebFilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_searchengine"
description: |-
  Configure web filter search engines.
---

# fortios_webfilter_searchengine
Configure web filter search engines.

## Example Usage

```hcl
resource "fortios_webfilter_searchengine" "trname" {
  charset    = "utf-8"
  hostname   = "sg.eiwuc.com"
  name       = "sg"
  query      = "sc="
  safesearch = "disable"
  url        = "^\\/f"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Search engine name.
* `hostname` - Hostname (regular expression).
* `url` - URL (regular expression).
* `query` - Code used to prefix a query (must end with an equals character).
* `safesearch` - Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
* `charset` - Search engine charset.
* `safesearch_str` - Safe search parameter used in the URL.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Webfilter SearchEngine can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webfilter_searchengine.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
