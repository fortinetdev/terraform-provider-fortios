---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_netflow"
description: |-
  Configure NetFlow.
---

# fortios_system_netflow
Configure NetFlow.

## Example Usage

```hcl
resource "fortios_system_netflow" "trname" {
  active_flow_timeout   = 30
  collector_ip          = "0.0.0.0"
  collector_port        = 2055
  inactive_flow_timeout = 15
  source_ip             = "0.0.0.0"
  template_tx_counter   = 20
  template_tx_timeout   = 30
}
```

## Argument Reference


The following arguments are supported:

* `collector_ip` - Collector IP.
* `collector_port` - NetFlow collector port number.
* `source_ip` - Source IP address for communication with the NetFlow agent.
* `active_flow_timeout` - Timeout to report active flows (1 - 60 min, default = 30).
* `inactive_flow_timeout` - Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
* `template_tx_timeout` - Timeout for periodic template flowset transmission (1 - 1440 min, default = 30).
* `template_tx_counter` - Counter of flowset records before resending a template flowset record.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Netflow can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_netflow.labelname SystemNetflow
$ unset "FORTIOS_IMPORT_TABLE"
```
