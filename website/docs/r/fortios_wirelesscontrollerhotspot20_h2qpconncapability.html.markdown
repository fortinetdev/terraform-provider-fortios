---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qpconncapability"
description: |-
  Configure connection capability.
---

# fortios_wirelesscontrollerhotspot20_h2qpconncapability
Configure connection capability.

## Example Usage

```hcl
resource "fortios_wirelesscontrollerhotspot20_h2qpconncapability" "trname" {
  esp_port      = "unknown"
  ftp_port      = "unknown"
  http_port     = "unknown"
  icmp_port     = "closed"
  ikev2_port    = "unknown"
  ikev2_xx_port = "unknown"
  name          = "1"
  pptp_vpn_port = "unknown"
  ssh_port      = "unknown"
  tls_port      = "unknown"
  voip_tcp_port = "unknown"
  voip_udp_port = "unknown"
}
```

## Argument Reference


The following arguments are supported:

* `name` - Connection capability name.
* `icmp_port` - Set ICMP port service status.
* `ftp_port` - Set FTP port service status.
* `ssh_port` - Set SSH port service status.
* `http_port` - Set HTTP port service status.
* `tls_port` - Set TLS VPN (HTTPS) port service status.
* `pptp_vpn_port` - Set Point to Point Tunneling Protocol (PPTP) VPN port service status.
* `voip_tcp_port` - Set VoIP TCP port service status.
* `voip_udp_port` - Set VoIP UDP port service status.
* `ikev2_port` - Set IKEv2 port service for IPsec VPN status.
* `ikev2_xx_port` - Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status.
* `esp_port` - Set ESP port service (used by IPsec VPNs) status.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 H2QpConnCapability can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontrollerhotspot20_h2qpconncapability.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
