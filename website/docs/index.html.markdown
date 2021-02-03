---
layout: "fortios"
page_title: "Provider: FortiOS"
sidebar_current: "docs-fortios-index"
description: |-
  The FortiOS provider interacts with FortiGate or FortiManager.
---

# FortiOS Provider

The FortiOS provider is used to interact with the resources supported by FortiOS and FortiManager. We need to configure the provider with the proper credentials before it can be used.

Two products are supported now: FortiGate and FortiManager, please use the navigation on the left to read more details about the available resources.

## Configuration for FortiGate

### Example Usage

```hcl
# Configure the FortiOS Provider for FortiGate
provider "fortios" {
  hostname     = "192.168.52.177"
  token        = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"
}

# Create a Static Route Item
resource "fortios_networking_route_static" "test1" {
  dst     = "110.2.2.122/32"
  gateway = "2.2.2.2"
  # ...
}
```

If it is used for testing, you can set `insecure` to "true" and unset `cabundlefile` to quickly set the provider up, for example:

```hcl
provider "fortios" {
  hostname = "192.168.52.177"
  token    = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
  insecure = "true"
}
```

Please refer to the Argument Reference below for more help on `insecure` and `cabundlefile`.

### Authentication

The FortiOS provider offers a means of providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables

#### Static credentials

Static credentials can be provided by adding a `token` key in-line in the FortiOS provider block.

Usage:

```hcl
provider "fortios" {
  hostname     = "192.168.52.177"
  token        = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"
}
```

#### Generate an API token for FortiOS

See the left navigation: `Guides` -> `Generate an API token for FortiOS`.

#### Environment variables

You can provide your credentials via the `FORTIOS_ACCESS_HOSTNAME`, `FORTIOS_ACCESS_TOKEN`, `FORTIOS_INSECURE` and `FORTIOS_CA_CABUNDLE` environment variables. Note that setting your FortiOS credentials using static credentials variables will override the environment variables.

Usage:

```shell
$ export "FORTIOS_ACCESS_HOSTNAME"="192.168.52.177"
$ export "FORTIOS_INSECURE"="false"
$ export "FORTIOS_ACCESS_TOKEN"="09m441wrwc10yGyrtQ4nk6mjbqcfz9"
$ export "FORTIOS_CA_CABUNDLE"="/path/yourCA.crt"
```

Then configure the FortiOS Provider as following:

```hcl
provider "fortios" {}

# Create a Static Route Item
resource "fortios_networking_route_static" "test1" {
  dst       = "110.2.2.122/32"
  gateway   = "2.2.2.2"
  blackhole = "disable"
  distance  = "22"
  weight    = "3"
  # …
}
```

### VDOM

If the FortiGate unit is running in VDOM mode, the `vdom` configuration needs to be added.

Usage:

```hcl
provider "fortios" {
  hostname     = "192.168.52.177"
  token        = "q3Hs49jxts195gkd9Hjsxnjtmr6k39"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"
  vdom         = "vdomtest"
}

resource "fortios_networking_route_static" "test1" {
  dst       = "120.2.2.122/32"
  gateway   = "2.2.2.2"
  blackhole = "disable"
  distance  = "22"
  weight    = "3"
  priority  = "3"
  device    = "lbforvdomtest"
  comment   = "Terraform test"
}
```

### Argument Reference

The following arguments are supported:

* `hostname` - (Optional) The hostname or IP address of FortiOS unit. It must be provided, but it can also be sourced from the `FORTIOS_ACCESS_HOSTNAME` environment variable.

* `token` - (Optional) The token of FortiOS unit. It must be provided, but it can also be sourced from the `FORTIOS_ACCESS_TOKEN` environment variable.

* `insecure` - (Optional) Control whether the Provider to perform insecure SSL requests. If omitted, the `FORTIOS_INSECURE` environment variable is used. If neither is set, default value is `false`.

* `cabundlefile` - (Optional) The path of a custom CA bundle file. You can specify a path to the file, or you can specify it by the `FORTIOS_CA_CABUNDLE` environment variable.

* `vdom` - (Optional) If the FortiGate unit is running in VDOM mode, you can use this argument to specify the name of the vdom to be set .



## Configuration for FortiManager

### Example Usage

```hcl
provider "fortios" {
  fmg_hostname     = "192.168.88.100"
  fmg_username     = "APIUser"
  fmg_passwd       = "admin"
  fmg_insecure     = false
  fmg_cabundlefile = "/path/yourCA.crt"
}

resource "fortios_fmg_system_dns" "test1" {
  primary   = "208.91.112.52"
  secondary = "208.91.112.54"
}
```

If it is used for testing, you can set `insecure` to true and unset `cabundlefile` to quickly set the provider up, for example:

```hcl
provider "fortios" {
  fmg_hostname = "192.168.88.100"
  fmg_username = "APIUser"
  fmg_passwd   = "admin"
  fmg_insecure = true
}
```

Please refer to the Argument Reference below for more help on `insecure` and `cabundlefile`.

### Authentication

As the same to provider for FortiGate, the following two methods are supported:

- Static credentials
- Environment variables

#### Static credentials

Static credentials can be provided by adding the `fmg_hostname`, `fmg_username` and `fmg_passwd` key in-line in the FortiOS provider block.

Usage:

```hcl
provider "fortios" {
  fmg_hostname     = "192.168.88.100"
  fmg_username     = "APIUser"
  fmg_passwd       = "admin"
  fmg_insecure     = false
  fmg_cabundlefile = "/path/yourCA.crt"
}
```

#### Environment variables

You can provide your credentials via the `FORTIOS_FMG_HOSTNAME`, `FORTIOS_FMG_USERNAME`, `FORTIOS_FMG_PASSWORD`, `FORTIOS_FMG_INSECURE` and `FORTIOS_FMG_CABUNDLE` environment variables. Note that setting your FortiOS credentials using static credentials variables will override the environment variables.

Usage:

```shell
$ export "FORTIOS_FMG_HOSTNAME"="192.168.88.100"
$ export "FORTIOS_FMG_USERNAME"="admin"
$ export "FORTIOS_FMG_PASSWORD"="admin"
$ export "FORTIOS_FMG_INSECURE"="false"
$ export "FORTIOS_FMG_CABUNDLE"="/path/yourCA.crt"
```

Then configure the FortiOS Provider as following:

```hcl
provider "fortios" {}

resource "fortios_fmg_system_dns" "test1" {
  primary   = "208.91.112.33"
  secondary = "208.91.112.44"
}
```

### Multi-Adom

Multi-Adom feature is supported in case of using FortiManager, just take the following example as a reference:

```hcl
provider "fortios" {
  fmg_hostname = "192.168.88.200"
  fmg_username = "APIUser"
  fmg_passwd   = "admin"
  fmg_product  = "fortimanager"
  fmg_insecure = true
}

resource "fortios_fmg_devicemanager_script" "test1" {
  name        = "config-intf3"
  description = "configure interface3"
  content     = "config system interface \n edit port3 \n\t set vdom \"root\"\n\t set ip 10.10.0.200 255.255.0.0 \n\t set allowaccess ping http https\n\t next \n end"
  target      = "device_database"
  adom        = "test-adom"
}

resource "fortios_fmg_devicemanager_script_execute" "test1" {
  script_name    = fortios_fmg_devicemanager_script.test1.name
  target_devname = "FGVM64-test"
  adom           = "test-adom"
  vdom           = "root"
}

resource "fortios_fmg_devicemanager_install_device" "test1" {
  target_devname = fortios_fmg_devicemanager_script_execute.test1.target_devname
  adom           = "test-adom"
  vdom           = "root"
}
```

This will install the script from the FMG(test-adom) to the FGT(root).

Note that one resource supports Multi-Adom feature if it has 'adom' argument.

### Argument Reference

The following arguments are supported:

* `fmg_hostname` - (Optional) The hostname or IP address of FortiManager. It must be provided, but it can also be sourced from the `FORTIOS_FMG_HOSTNAME` environment variable.

* `fmg_username` - (Optional) The username of FortiManager. It must be provided, but it can also be sourced from the `FORTIOS_FMG_USERNAME` environment variable.

* `fmg_passwd` - (Optional) The password of FortiManager, it can also be sourced from the `FORTIOS_FMG_PASSWORD` environment variable.

* `fmg_insecure` - (Optional) Control whether the Provider to perform insecure SSL requests. If omitted, the `FORTIOS_FMG_INSECURE` environment variable is used. If neither is set, default value is `false`.

* `fmg_cabundlefile` - (Optional) The path of a custom CA bundle file. You can specify a path to the file, or you can specify it by the `FORTIOS_FMG_CABUNDLE` environment variable.

## Release
Check out the FortiOS provider release notes and additional information from: [the FortiOS provider releases](https://github.com/fortinetdev/terraform-provider-fortios/releases).

## Versioning

The provider can cover FortiOS 6.0, 6.2, 6.4, 6.6 versions, the configuration of all parameters should be based on the relevant FortiOS version manual. The provider can cover FortiManager 6.0 and 6.2 versions. When using FortiManager, make sure the versions of FortiManager and the FortiGates controlled by it are the same.
