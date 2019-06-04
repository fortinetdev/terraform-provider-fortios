---
layout: "fortios"
page_title: "Provider: FortiOS"
sidebar_current: "docs-fortios-index"
description: |-
  The FortiOS provider interacts with FortiGate.
---

# FortiOS Provider

The FortiOS provider is used to interact with the resources supported by FortiOS. We need to configure the provider with with the proper credentials before it can be used.

## Example Usage

```hcl
# Configure the FortiOS Provider
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
}

# Create a Static Route Item
resource "fortios_networking_route_static" "test1" {
	dst = "110.2.2.122/32"
	gateway = "2.2.2.2"
	# ...
}
```

## Authentication

The FortiOS provider offers a means of providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables

### Static credentials
Static credentials can be provided by adding a `token` key in-line in the FortiOS provider block.

Usage:
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
}
```
### Environment variables
You can provide your credentials via the `FORTIOS_ACCESS_HOSTNAME` and `FORTIOS_ACCESS_TOKEN` environment variables. Note that setting your FortiOS credentials using static credentials variables will override the environment variables.

Usage:

```hcl
$ export FORTIOS_ACCESS_HOSTNAME=192.168.52.177
$ export FORTIOS_ACCESS_TOKEN=q3Hs49jxts195gkd9Hjsxnjtmr6k39
```

Then configure the FortiOS Provider as following:

```hcl
provider "fortios" { }

# Create a Static Route Item
resource "fortios_networking_route_static" "test1" {
	dst = "110.2.2.122/32"
	gateway = "2.2.2.2"
	blackhole = "disable"
	distance = "22"
	weight = "3"
	# …
}
```
## VDOM
If the FortiGate unit is running in VDOM mode, the `vdom` configuration needs to be added.

Usage:

```hcl
provider "fortios" {
	hostname = "192.168.52.177"
	token = "q3Hs49jxts195gkd9Hjsxnjtmr6k39"
	vdom = "vdomtest"
}

resource "fortios_networking_route_static" "test1" {
	dst = "120.2.2.122/32"
	gateway = "2.2.2.2"
	blackhole = "disable"
	distance = "22"
	weight = "3"
	priority = "3"
	device = "lbforvdomtest"
	comment = "Terraform test"
}
```

## Argument Reference
The following arguments are supported:

* `hostname` - (Optional) This is the hostname or IP address of FortiOS unit. It must be provided, but it can also be sourced from the `FORTIOS_ACCESS_HOSTNAME` environment variable.

* `token` - (Optional) This is the token of FortiOS unit. It must be provided, but it can also be sourced from the `FORTIOS_ACCESS_TOKEN` environment variable.

* `insecure` - (Optional) This is used to control whether the Provider to perform insecure SSL requests. If omitted, the `FORTIOS_INSECURE` environment variable is used. If neither is set, default value is `false`.

* `cabundlefile` - (Optional) The path of a custom CA bundle file. You can specify a path to the file, or you can specify it by the `FORTIOS_CA_CABUNDLE` environment variable.

* `vdom` - (Optional) If the FortiGate unit is running in VDOM mode, you can use this argument to specify the name of the vdom to be set .

## Versioning

The provider can cover both FortiOS 6.0 and 6.2 versions.
