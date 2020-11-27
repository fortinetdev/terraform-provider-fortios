---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ha"
description: |-
  Configure HA.
---

# fortios_system_ha
Configure HA.

Due to current FortiOS API limitations, the feature is temporarily unavailable. Please use the following resource configuration as an alternative.

## Example1

```
resource "fortios_system_autoscript" "trname" {
  interval    = 1
  name        = "1"
  output_size = 10
  repeat      = 1
  script      = <<EOF
config system ha
    set session-pickup enable
    set session-pickup-connectionless enable
    set session-pickup-expectation enable
    set session-pickup-nat enable
    set override disable
end

EOF
  start       = "auto"
}

```

## Example2

```
resource "fortios_system_autoscript" "trname1" {
  interval    = 1
  name        = "1"
  output_size = 10
  repeat      = 1
  script      = <<EOF
config system ha
	set group-name "AzureHA"
	set mode a-p
	set hbdev "port3" 100
	set session-pickup enable
	set session-pickup-connectionless enable
	set override disable
	set priority 255
	set unicast-hb enable
	set unicast-hb-peerip 10.0.3.5
end

EOF
  start       = "auto"
}

resource "fortios_system_autoscript" "trname2" {
  interval    = 1
  name        = format("%s%s", fortios_system_autoscript.trname1.name, "extend")
  output_size = 10
  repeat      = 1
  script      = <<EOF
config system ha
	set ha-mgmt-status enable
	config ha-mgmt-interfaces
		edit 1
			set interface "port4"
			set gateway 10.0.4.1
		end
	end
end

EOF
  start       = "auto"
}
```

-> FortiOS has a limit of 255 characters for system autoscript->script, if the configuration content exceeds 255 characters, it can be implemented in multiple stages. In the first stage, configure the necessary parameters. In the subsequent stages, configure other optional parameters. Whenever possible, delete spaces and tabs at the beginning of each line. The goal is to try to reach the 255 limit. The dependencies between several stages need to be ensured.

