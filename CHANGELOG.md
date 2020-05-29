## 1.2.0 (Unreleased)

IMPROVEMENTS:

* Support Terraform Plugin SDK ([#50](https://github.com/terraform-providers/terraform-provider-fortios/pull/50))

FEATURES:

* **New Resource:** ` fortios_fmg_devicemanager_device`
* **New Resource:** ` fortios_fmg_devicemanager_installdevice`
* **New Resource:** ` fortios_fmg_devicemanager_installpackage`
* **New Resource:** ` fortios_fmg_devicemanager_script_execute`
* **New Resource:** ` fortios_fmg_devicemanager_script`
* **New Resource:** ` fortios_fmg_firewall_object_address`
* **New Resource:** ` fortios_fmg_firewall_object_ippool`
* **New Resource:** ` fortios_fmg_firewall_object_service`
* **New Resource:** ` fortios_fmg_firewall_object_vip`
* **New Resource:** ` fortios_fmg_firewall_security_policy`
* **New Resource:** ` fortios_fmg_firewall_security_policypackage`
* **New Resource:** ` fortios_fmg_object_adom_revision`
* **New Resource:** ` fortios_fmg_system_admin`
* **New Resource:** ` fortios_fmg_system_admin_profiles`
* **New Resource:** ` fortios_fmg_system_admin_user`
* **New Resource:** ` fortios_fmg_system_adom`
* **New Resource:** ` fortios_fmg_system_dns`
* **New Resource:** ` fortios_fmg_system_global`
* **New Resource:** ` fortios_fmg_system_license_forticare`
* **New Resource:** ` fortios_fmg_system_license_vm`
* **New Resource:** ` fortios_fmg_system_network_interface`
* **New Resource:** ` fortios_fmg_system_network_route`
* **New Resource:** ` fortios_fmg_system_ntp`
* **New Resource:** ` fortios_fmg_system_syslogserver`
* **New Resource:** ` fortios_firewall_security_policyseq`
* **New Resource:** ` fortios_firewall_object_servicecategory`


## 1.1.0 (June 25, 2019)

IMPROVEMENTS:

* Add support for Terraform v0.12.2 ([#7](https://github.com/terraform-providers/terraform-provider-fortios/issues/7))
* Add tips for prohibiting modification of the key value ([#20](https://github.com/terraform-providers/terraform-provider-fortios/issues/20))
* Add support for CA certificate

BUG FIXES:

* resource/object_address: Fix incorrect behavior with slash in address object ([#1](https://github.com/terraform-providers/terraform-provider-fortios/issues/1))
* resource/object_service: Fix terraform crashes ([#5](https://github.com/terraform-providers/terraform-provider-fortios/issues/5))
* resource/object_address: Fix getting diff ([#6](https://github.com/terraform-providers/terraform-provider-fortios/issues/6))
* resource/object_service: Fix getting diff in service port section ([#12](https://github.com/terraform-providers/terraform-provider-fortios/issues/12))
* resource/object_address: Fix wrong prefix handling in firewall addresses ([#13](https://github.com/terraform-providers/terraform-provider-fortios/issues/13))


## 1.0.0 (April 18, 2019)

FEATURES:

* **New Resource:** `fortios_networking_route_static`
* **New Resource:** `fortios_networking_interface_port`
* **New Resource:** `fortios_system_admin_profiles`
* **New Resource:** `fortios_system_admin_administrator`
* **New Resource:** `fortios_firewall_object_address`
* **New Resource:** `fortios_firewall_object_addressgroup`
* **New Resource:** `fortios_firewall_object_service`
* **New Resource:** `fortios_firewall_object_servicegroup`
* **New Resource:** `fortios_firewall_object_vip`
* **New Resource:** `fortios_firewall_object_vipgroup`
* **New Resource:** `fortios_firewall_object_ippool`
* **New Resource:** `fortios_firewall_security_policy`
* **New Resource:** `fortios_system_setting_global`
* **New Resource:** `fortios_system_setting_dns`
* **New Resource:** `fortios_system_setting_ntp`
* **New Resource:** `fortios_log_syslog_setting`
* **New Resource:** `fortios_log_fortianalyzer_setting`
* **New Resource:** `fortios_system_apiuser_setting`
* **New Resource:** `fortios_system_vdom_setting`
* **New Resource:** `fortios_system_license_forticare`
* **New Resource:** `fortios_system_license_vdom`
* **New Resource:** `fortios_system_license_vm`
* **New Resource:** `fortios_vpn_ipsec_phase1interface`
* **New Resource:** `fortios_vpn_ipsec_phase2interface`
