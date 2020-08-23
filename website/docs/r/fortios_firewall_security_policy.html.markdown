---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_security_policy"
sidebar_current: "docs-fortios-resource-firewall-security-policy"
subcategory: "FortiGate OldVersion"
description: |-
  Provides a resource to configure firewall policies of FortiOS.
---

# fortios_firewall_security_policy
Provides a resource to configure firewall policies of FortiOS.

~> **Warning:** The resource will be deprecated and replaced by `fortios_firewall_policy`.

## Example Usage 1
```hcl
resource "fortios_firewall_security_policy" "test1" {
  name                = "ap11"
  srcintf             = ["port2"]
  dstintf             = ["port1"]
  srcaddr             = ["swscan.apple.com", "google-play"]
  dstaddr             = ["swscan.apple.com", "update.microsoft.com"]
  internet_service    = "disable"
  internet_service_id = []
  schedule            = "always"
  service             = ["ALL_ICMP", "FTP"]
  action              = "accept"
  utm_status          = "enable"
  logtraffic          = "all"
  logtraffic_start    = "enable"
  capture_packet      = "enable"
  ippool              = "enable"
  poolname            = ["rewq", "rbb"]
  groups              = ["Guest-group", "SSO_Guest_Users"]
  devices             = ["android-phone", "android-tablet"]
  comments            = "security policy"
  av_profile          = "wifi-default"
  webfilter_profile   = "monitor-all"
  dnsfilter_profile   = "default"
  ips_sensor          = "protect_client"
  application_list    = "block-high-risk"
  ssl_ssh_profile     = "certificate-inspection"
  nat                 = "enable"
}
```

## Example Usage 2
```hcl
resource "fortios_firewall_security_policy" "test2" {
  name                = "ap21"
  srcintf             = ["port2"]
  dstintf             = ["port1"]
  srcaddr             = ["swscan.apple.com", "google-play"]
  dstaddr             = ["swscan.apple.com", "update.microsoft.com"]
  internet_service    = "enable"
  internet_service_id = [917520, 6881402, 393219]
  schedule            = "always"
  service             = []
  action              = "accept"
  utm_status          = "enable"
  logtraffic          = "all"
  logtraffic_start    = "enable"
  capture_packet      = "enable"
  ippool              = "enable"
  poolname            = ["rewq", "rbb"]
  groups              = ["Guest-group", "SSO_Guest_Users"]
  devices             = ["android-phone", "android-tablet"]
  comments            = "security policy"
  av_profile          = "wifi-default"
  webfilter_profile   = "monitor-all"
  dnsfilter_profile   = "default"
  ips_sensor          = "protect_client"
  application_list    = "block-high-risk"
  ssl_ssh_profile     = "certificate-inspection"
  nat                 = "enable"
}
```

## Example Usage 3
```hcl
resource "fortios_firewall_security_policy" "test1" {
  name                     = "ap12221"
  srcintf                  = ["port3"]
  dstintf                  = ["port4"]
  srcaddr                  = []
  dstaddr                  = []
  internet_service         = "enable"
  internet_service_id      = [5242880]
  internet_service_src     = "enable"
  internet_service_src_id  = [65643]
  users                    = ["guest"]
  status                   = "enable"
  schedule                 = "always"
  service                  = []
  action                   = "accept"
  utm_status               = "enable"
  logtraffic               = "all"
  logtraffic_start         = "enable"
  capture_packet           = "enable"
  ippool                   = "disable"
  poolname                 = []
  groups                   = ["Guest-group", "SSO_Guest_Users"]
  devices                  = []
  comments                 = "security policy"
  av_profile               = "wifi-default"
  webfilter_profile        = "monitor-all"
  dnsfilter_profile        = "default"
  ips_sensor               = "protect_client"
  application_list         = "block-high-risk"
  ssl_ssh_profile          = "certificate-inspection"
  nat                      = "enable"
  profile_protocol_options = "default"
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Policy name.
* `srcintf` - (Required) Incoming (ingress) interface.
* `dstintf` - (Required) Outgoing (egress) interface.
* `srcaddr` - (Required) Source address and address group names.
* `dstaddr` - (Required) Destination address and address group names.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.
* `internet_service_id` - Internet Service ID.
* `action` - (Required) Policy action.
* `schedule` - (Required) Schedule name.
* `service` - (Required) Service and service group names..
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions.
* `logtraffic_start` - Record logs when a session starts and ends.
* `capture_packet` - Enable/disable capture packets.
* `ippool` - Enable to use IP Pools for source NAT.
* `poolname` - IP Pool names.
* `groups` - Names of user groups that can authenticate with this policy.
* `devices` - Device type category.
* `comments` - Comment.
* `av_profile` - Name of an existing Antivirus profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `ips_sensor` - Name of an existing IPS sensor.
* `application_list` - Name of an existing Application list.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `nat` - Enable/disable source NAT.
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.
* `internet_service_src_id` - Internet Service source ID.
* `users` - Names of individual users that can authenticate with this policy.
* `status` - Enable or disable this policy.
* `profile_protocol_options` - Name of an existing Protocol options profile.

## Attributes Reference
The following attributes are exported:

* `id` - The ID of the firewall policy item.
* `name` - Policy name.
* `srcintf` - Incoming (ingress) interface.
* `dstintf` - Outgoing (egress) interface.
* `srcaddr` - Source address and address group names.
* `dstaddr` - Destination address and address group names.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.
* `internet_service_id` - Internet Service ID.
* `action` - Policy action.
* `schedule` - Schedule name.
* `service` - Service and service group names..
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions.
* `logtraffic_start` - Record logs when a session starts and ends.
* `capture_packet` - Enable/disable capture packets.
* `ippool` - Enable to use IP Pools for source NAT.
* `poolname` - IP Pool names.
* `groups` - Names of user groups that can authenticate with this policy.
* `devices` - Device type category.
* `comments` - Comment.
* `av_profile` - Name of an existing Antivirus profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `ips_sensor` - Name of an existing IPS sensor.
* `application_list` - Name of an existing Application list.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `nat` - Enable/disable source NAT.
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.
* `internet_service_src_id` - Internet Service source ID.
* `users` - Names of individual users that can authenticate with this policy.
* `status` - Enable or disable this policy.
* `profile_protocol_options` - Name of an existing Protocol options profile.

