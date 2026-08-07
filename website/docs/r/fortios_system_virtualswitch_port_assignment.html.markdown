---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_virtualswitch_port_assignment"
description: |-
  Manage port assignments for FortiOS system virtual switches.
---

# fortios_system_virtualswitch_port_assignment
Manage port assignments for FortiOS system virtual switches. This resource does not create or delete virtual switches; it only assigns ports to existing virtual switches. It handles port moves between virtual switches by using a two-pass update strategy: first removing ports from existing switches, then applying full updates.

~> **Warning:** Do not use `fortios_system_virtualswitch_port_assignment` together with another resource that manages the `port` list of the same virtual switch at the same time. They will conflict with each other and cause inconsistent state. Choose one resource to own the port assignments for a given virtual switch.

## Example Usage

```hcl
resource "fortios_system_virtualswitch_port_assignment" "trname" {
  port_map = {
    "vlan.105" = {
      vdom = "root"
      ports = [
        { name = "port2", alias = "web" },
        { name = "port3", alias = "app" },
      ]
    }
    "vlan.109" = {
      ports = [
        { name = "port4", alias = "db" },
        { name = "port5", alias = "mgmt" },
      ]
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `port_map` - (Required) A map of virtual switch name to an object describing the port assignments for that switch. The key is the virtual switch name, and the value is an object with the following attributes.

The `port_map` value supports:

* `vdom` - (Optional) VDOM to which the virtual switch belongs. If not set, inherits from the provider configuration.
* `ports` - (Required) List of member ports to assign to the virtual switch. The structure of the `ports` block is documented below.

The `ports` block supports:

* `name` - (Optional) Physical interface name.
* `alias` - (Optional) Alias.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:

* `id` - An identifier for the resource. Always set to `system_virtualswitch_port_assignment`.

## Validation

This resource performs the following validations during `terraform plan`:

* **Port Conflict Detection**: Checks whether the same port is assigned to multiple virtual switches. If a conflict is found, the plan fails with an error before any API calls are made.

## Update Strategy

When updating port assignments, this resource uses a three-phase strategy to handle port moves and removals between switches:

1. **Removed Switch Phase**: For switches that have been removed from `port_map` entirely (present in state but absent from the new configuration), clear all of their assigned ports. This must happen before any adds so the freed ports can be reclaimed by other switches.
2. **Port Removal Phase**: For switches that remain in the configuration but are losing some ports, first update with only the kept ports (no new ports added). This frees up ports for other switches.
3. **Full Update Phase**: Apply the complete port list for all switches remaining in the configuration, including adding new ports.

This ensures that a port can be removed from one virtual switch and added to another, or removed entirely by dropping a switch from `port_map`, in the same `terraform apply`.

## Import

This resource does not support `terraform import`. It does not correspond to a single remote object that can be identified by an import ID; instead it manages port assignments across multiple virtual switches described by the `port_map` argument in the configuration. To adopt existing port assignments into Terraform management, define the `port_map` to match the current state and run `terraform import` is not required — `terraform apply` will reconcile the configuration with the device.
