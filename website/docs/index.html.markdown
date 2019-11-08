---
layout: "fortios"
page_title: "Provider: FortiOS"
sidebar_current: "docs-fortios-index"
description: |-
  The FortiOS provider interacts with FortiGate or FortiManager.
---

# FortiOS Provider

The FortiOS provider is used to interact with the resources supported by FortiOS. We need to configure the provider with the proper credentials before it can be used.

Two products are supported now: FortiGate and FortiManager, please use the navigation to the left to read more details about the available resources.

# FortiOS Provider for FortiGate

## Example Usage

```hcl
# Configure the FortiOS Provider for FortiGate
provider "fortios" {
	hostname = "192.168.52.177"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
	insecure = "false"
	cabundlefile = "/path/yourCA.crt"
}

# Create a Static Route Item
resource "fortios_networking_route_static" "test1" {
	dst = "110.2.2.122/32"
	gateway = "2.2.2.2"
	# ...
}
```

If it is used for testing, you can set `insecure` to "true" and unset `cabundlefile` to quickly set the provider up, for example:

```hcl
provider "fortios" {
	hostname = "192.168.52.177"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
	insecure = "true"
}
```

Please refer to the Argument Reference below for more help on `insecure` and `cabundlefile`.

## Authentication

The FortiOS provider offers a means of providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables

### Static credentials

Static credentials can be provided by adding a `token` key in-line in the FortiOS provider block.

Usage:

```hcl
provider "fortios" {
	hostname = "192.168.52.177"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
	insecure = "false"
	cabundlefile = "/path/yourCA.crt"
}
```

For the configuration of token, please refer to the `system->system api-user` and `execute->api-user` chapters of the `FortiOS Handbook - CLI Reference`.

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
	insecure = "false"
	cabundlefile = "/path/yourCA.crt"
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

* `hostname` - (Optional) The hostname or IP address of FortiOS unit. It must be provided, but it can also be sourced from the `FORTIOS_ACCESS_HOSTNAME` environment variable.

* `token` - (Optional) The token of FortiOS unit. It must be provided, but it can also be sourced from the `FORTIOS_ACCESS_TOKEN` environment variable.

* `insecure` - (Optional) Control whether the Provider to perform insecure SSL requests. If omitted, the `FORTIOS_INSECURE` environment variable is used. If neither is set, default value is `false`.

* `cabundlefile` - (Optional) The path of a custom CA bundle file. You can specify a path to the file, or you can specify it by the `FORTIOS_CA_CABUNDLE` environment variable.

* `vdom` - (Optional) If the FortiGate unit is running in VDOM mode, you can use this argument to specify the name of the vdom to be set .


# FortiOS Provider for FortiManager

## Example Usage

```hcl
provider "fortios" {
	hostname = "192.168.88.100"
	username = "APIUser"
	passwd = "admin"
	product = "fortimanager"
	insecure = false
	cabundlefile = "/path/yourCA.crt"
}

resource "fortios_fortimanager_system_dns" "test1" {
	primary = "208.91.112.52"
	secondary = "208.91.112.54"
}
```

If it is used for testing, you can set `insecure` to true and unset `cabundlefile` to quickly set the provider up, for example:

```hcl
provider "fortios" {
	hostname = "192.168.88.100"
	username = "APIUser"
	passwd = "admin"
	product = "fortimanager"
	insecure = true
}
```

Please refer to the Argument Reference below for more help on `insecure` and `cabundlefile`.

## Authentication

As the same to provider for FortiGate, the following two methods are supported:

- Static credentials
- Environment variables

### Static credentials

Static credentials can be provided by adding the `hostname`, `username` and `passwd`  key in-line in the FortiOS provider block.

Usage:

```hcl
provider "fortios" {
	hostname = "192.168.88.100"
	username = "APIUser"
	passwd = "admin"
	product = "fortimanager"
	insecure = false
	cabundlefile = "/path/yourCA.crt"
}
```

### Environment variables

You can provide your credentials via the `FORTIOS_FMG_HOSTNAME`, `FORTIOS_FMG_USERNAME` and `FORTIOS_FMG_PASSWORD` environment variables. Note that setting your FortiOS credentials using static credentials variables will override the environment variables.

Usage:

```hcl
$ export FORTIOS_FMG_HOSTNAME="192.168.88.100"
$ export FORTIOS_FMG_USERNAME="APIUser"
$ export FORTIOS_FMG_PASSWORD="admin"
```

Then configure the FortiOS Provider as following:

```hcl
provider "fortios" {
	product = "fortimanager"
	insecure = false
	cabundlefile = "/path/yourCA.crt"
}

resource "fortios_fortimanager_system_dns" "test1" {
	primary = "208.91.112.33"
	secondary = "208.91.112.44"
}
```

## Argument Reference

The following arguments are supported:

* `hostname` - (Optional) The hostname or IP address of FortiManager. It must be provided, but it can also be sourced from the `FORTIOS_FMG_HOSTNAME` environment variable.

* `username` - (Optional) The username of FortiManager. It must be provided, but it can also be sourced from the `FORTIOS_FMG_USERNAME` environment variable.

* `passwd` - (Optional) The password of FortiManager, it can also be sourced from the `FORTIOS_FMG_PASSWORD` environment variable.

* `product` - (Optional) Which product to use, should be set to "fortimanager" here. default is "fortigate", it can also be sourced from the `FORTIOS_PRODUCT` environment variable.

* `insecure` - (Optional) Control whether the Provider to perform insecure SSL requests. If omitted, the `FORTIOS_FMG_INSECURE` environment variable is used. If neither is set, default value is `false`.

* `cabundlefile` - (Optional) The path of a custom CA bundle file. You can specify a path to the file, or you can specify it by the `FORTIOS_FMG_CABUNDLE` environment variable.

# Versioning

The provider can cover both FortiOS 6.0 and 6.2 versions.

When using for FortiManager, make sure the version of FortiManager and FortiGate should be the same.
