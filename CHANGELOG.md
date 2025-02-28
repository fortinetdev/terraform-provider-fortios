## 1.23.0 (Unreleased)


## 1.22.0 (Feb 28, 2025)
BUG FIXES:

* Fix default value could not be set issue;
* Fix IP format convert issue;

IMPROVEMENTS:

* Support FortiOS version 7.0.16, 7.0.17, 7.4.6, 7.4.7, 7.6.1, 7.6.2;
* Add write_only variable for resource system_license_fortiflex;

FEATURES:

* **New Resource:** `fortios_casb_attributematch`
* **New Resource:** `fortios_system_automationcondition`
* **New Resource:** `fortios_system_healthcheckfortiguard`
* **New Resource:** `fortios_system_ngfwsettings`
* **New Resource:** `fortios_system_sdnvpn`
* **New Resource:** `fortios_systemsecurityrating_controls`
* **New Resource:** `fortios_systemsecurityrating_settings`
* **New Resource:** `fortios_webproxy_isolatorserver`
* **New Resource:** `fortios_webfilter_ftgdlocalrisk`
* **New Resource:** `fortios_webfilter_ftgdrisklevel`
* **New Resource:** `fortios_ztna_reverseconnector`
* **New Resource:** `fortios_ztna_webportal`
* **New Resource:** `fortios_ztna_webportalbookmark`
* **New Resource:** `fortios_ztna_webproxy`
* **New Resource:** `fortios_firewall_localinpolicy_sort`
* **New Resource:** `fortios_firewall_localinpolicy_move`
* **New Resource:** `fortios_firewall_localinpolicy6_sort`
* **New Resource:** `fortios_firewall_localinpolicy6_move`

## 1.21.1 (Nov 1, 2024)
BUG FIXES:

* Fix destroy issue of resource system_ftmpush;
* Fix issue of timezone could not been unset of resource system_global;
* Fix auto-generated id always show changes issue;
* Fix ip format issue;
* Fix error of set zero value of resource system_interface;
* Fix token issue of 7.4.5;

IMPROVEMENTS:

* Support FortiOS version 7.2.10, 7.4.5;


## 1.21.0 (Sep 13, 2024)
BUG FIXES:

* Fix global scope issue;
* Fix issue of type mismatch from API response for integer type of variables;
* Fix variables could not been removed/reset issue;

IMPROVEMENTS:

* Support FortiOS version 7.2.9, 7.6.0;

FEATURES:

* **New Resource:** `fortios_system_vneinterface`
* **New Resource:** `fortios_systemsnmp_rmonstat`
* **New Resource:** `fortios_user_scim`
* **New Resource:** `fortios_ztna_trafficforwardproxyreverseservice`
* **New Resource:** `fortios_ztna_trafficforwardproxy`
* **New Resource:** `fortios_firewall_shapingpolicy_move`
* **New Resource:** `fortios_firewall_shapingpolicy_sort`

## 1.20.0 (Jun 21, 2024)
BUG FIXES:

* Fix extra quote issue;
* Fix issue of multiple argument conflict caused by version differences for resource system_automationstitch;

IMPROVEMENTS:

* Support FortiOS version 7.4.4;

FEATURES:

* **New Resource:** `fortios_extensioncontroller_extendervap`
* **New Resource:** `fortios_firewall_ondemandsniffer`
* **New Resource:** `fortios_system_sshconfig`

## 1.19.1 (Apr 29, 2024)
BUG FIXES:

* Fix issue of merged argument element contains comma;
* Fix length limitation issue of merged argument;
* Fix issue of vdomparam cause force replacement even using the same vdom value;
* Fix data type error of API return empty string for integer data type;

IMPROVEMENTS:

* Support FortiOS version 6.4.15, 7.0.14, 7.0.15, 7.2.7, 7.2.8, 7.4.3;
* Update version check functionality;

FEATURES:

* **New Resource:** `fortios_system_license_fortiflex`

## 1.19.0 (Jan 30, 2024)
BUG FIXES:

* Fix issue of remove BGP redistribute route map;
* Set argument of ca to sensitive;
* Fix wrong required arguments on resource firewall_shapingpolicy;
* Fix crash issue of resource firewall_centralsnatmap_sort;
* Fix extra double quotes issue;
* Change doc of resource fortios_system_ha to use its own resource;

IMPROVEMENTS:

* Support FortiOS v7.4.2;
* Update doc for the different description in different FOS versions;

FEATURES:

* **New Resource:** `fortios_diameterfilter_profile`
* **New Resource:** `fortios_dlp_exactdatamatch`
* **New Resource:** `fortios_rule_fmwp`
* **New Resource:** `fortios_user_externalidentityprovider`
* **New Resource:** `fortios_videofilter_keyword`
* **New Resource:** `fortios_vpn_qkd`

## 1.18.1 (Nov 15, 2023)
BUG FIXES:

* Fix move resources always show changes issue (#289);
* Fix extra spaces issue (#299);

IMPROVEMENTS:

* Support FortiOS v7.0.13, 7.2.6;
* Set non-order variables to TypeSet;
* Set variable 'name' as required of resource fortios_user_ldap (#298);

FEATURES:

* **New Data source:** `data_source_vpnssl_settings`

## 1.18.0 (Sep 11, 2023)
BUG FIXES:

* Fix crash issue (#291);

IMPROVEMENTS:

* Support FortiOS v6.4.13, v6.4.14, v7.0.12, v7.4.1;

FEATURES:

* **New Resource:** `fortios_casb_useractivity`
* **New Resource:** `fortios_casb_saasapplication`
* **New Resource:** `fortios_casb_profile`
* **New Resource:** `fortios_rule_otvp`
* **New Resource:** `fortios_rule_otdt`
* **New Resource:** `fortios_switchcontrollerptp_profile`
* **New Resource:** `fortios_switchcontrollerptp_interfacepolicy`
* **New Resource:** `fortios_system_speedtestsetting`
* **New Resource:** `fortios_virtualpatch_profile`
* **New Resource:** `fortios_webproxy_fastfallback`
* **New Resource:** `fortios_firewall_policy_sort`
* **New Resource:** `fortios_firewall_policy_move`
* **New Resource:** `fortios_firewall_securitypolicy_sort`
* **New Resource:** `fortiof_firewall_securitypolicy_move`

## 1.17.0 (Jun 22, 2023)
BUG FIXES:

* Fix issue of can't update admin user (#248);
* Fix issue of no option to not use HTTP proxy (#253);
* Fix crash issue with insufficient permissions (#257);
* Fix type change issue (#258)
* Fix issue of configuration been destroyed after second apply (#265);
* Fix range issue (#284)
* Remove extra quotes in the conversion of argument from block to string;
* Fix duplicate argument issue; 

IMPROVEMENTS:

* Support FortiOS v6.4.11, v6.4.12, v7.0.7, v7.0.8, v7.0.9, v7.0.10, v7.0.11, v7.2.3, v7.2.4, v7.4.0;
* Update mergeable arguments with type changes; 
* Update version check function;
* Add variable get_all_tables to determine whether get all table or complex items when refresh the state file;
* Support username/password login;

FEATURES:

* **New Resource:** `fortios_antivirus_exemptlist`
* **New Resource:** `fortios_endpointcontrol_fctemsoverride`
* **New Resource:** `fortios_firewall_internetservicesubapp`
* **New Resource:** `fortios_router_extcommunitylist`
* **New Resource:** `fortios_switchcontrolleracl_ingress`
* **New Resource:** `fortios_switchcontrolleracl_group`
* **New Resource:** `fortios_system_deviceupgrade`
* **New Resource:** `fortios_system_evpn`
* **New Resource:** `fortios_system_fabricvpn`
* **New Resource:** `fortios_system_pcpserver`
* **New Resource:** `fortios_system_sdnproxy`
* **New Resource:** `fortios_system_ssofortigatecloudadmin`
* **New Resource:** `fortios_vpn_kmipserver`

## 1.16.0 (Oct 7, 2022)
BUG FIXES:

* Improve requirement of argument fosid for resource fortios_systemdhcp_cerver (#221);
* Fix policyid always forcing replacement issue (#243);
* Fix issue of static block always shows need to be removed (#245);
* Fix crash issue caused by mkey type convert;
* Fix issue of predefined object always shows need to be deleted;

IMPROVEMENTS:

* Support FortiOS v6.4.10, v7.2.1, v7.2.2;
* Improve certificate parameters' format issue;

FEATURES:

* **New Resource:** `fortios_extensioncontroller_fortigateprofile`
* **New Resource:** `fortios_extensioncontroller_fortigate`
* **New Resource:** `fortios_extensioncontroller_extenderprofile`
* **New Resource:** `fortios_extensioncontroller_extender`
* **New Resource:** `fortios_extensioncontroller_dataplan`
* **New Resource:** `fortios_firewall_global`
* **New Resource:** `fortios_firewall_networkservicedynamic`

# 1.15.0 (Aug 5, 2022)

BUG FIXES:

* Improve sort blocks by adding options of alphabetical sorting and natural sorting;
* Fix issue of parameter `split_dns` could not be deleted (#228);
* Fix issue of parameter `certificate` always needs to be updated (#230);
* Fix issue of parameter `associated_interface` could not be unset (#233);
* Fix issue of SSH key needs extra quotes (#234);
* Fix issue of parameter `list` always shows need to be updated (#235);
* Fix issue of filter content could not contains spaces;
* Fix issue of prarmeter `dhcp_relay_ip` has extra quotes (#239);

IMPROVEMENTS:

* Support FortiOS v7.0.5, v7.0.6, v7.2.0;
* Add parameter `http_proxy` to provider configuration (#229);
* Add parameter `cabundlecontent` to provider configuration (#238);

FEATURES:

* **New Resource:** `fortios_automation_setting`
* **New Resource:** `fortios_dlp_dictionary`
* **New Resource:** `fortios_dlp_datatype`
* **New Resource:** `fortios_dlp_profile`
* **New Resource:** `fortios_icap_servergroup`
* **New Resource:** `fortios_system_fortindr`
* **New Resource:** `fortios_systemsnmp_mibview`

# 1.14.1 (Apr 25, 2022)

BUG FIXES:

* Fix filter issue;
* Fix incorrect requirement issue;
* Improve delete operation for non-deletable resources;
* Improve import process;
* Fix capital letters issue;
* Fix reset argument issue for resource fortios_firewall_policy;
* Improve certificate argument for API format issue;
* Fix argument type convert issue;

## 1.14.0 (Feb 1, 2022)

BUG FIXES:

* Fix parameter value's range limitation issue #190, #195
* Fix argument trailing space issue #184

IMPROVEMENTS:

* Support FortiOS v7.0.1 - v7.0.4 #201
* Add helpful error information from HTTP response to Terraform output #193

FEATURES:

* **New Resource:** `fortios_extendercontroller_extenderprofile`
* **New Resource:** `fortios_firewall_accessproxy`
* **New Resource:** `fortios_firewall_accessproxysshclientcert`
* **New Resource:** `fortios_firewall_accessproxyvirtualhost`
* **New Resource:** `fortios_firewall_accessproxy6`
* **New Resource:** `fortios_logtacacsaccounting_filter`
* **New Resource:** `fortios_logtacacsaccounting_setting`
* **New Resource:** `fortios_logtacacsaccounting2_filter`
* **New Resource:** `fortios_logtacacsaccounting2_setting`
* **New Resource:** `fortios_logtacacsaccounting3_filter"`
* **New Resource:** `fortios_logtacacsaccounting3_setting`
* **New Resource:** `fortios_sctpfilter_profile`
* **New Resource:** `fortios_switchcontroller_dynamicportpolicy`
* **New Resource:** `fortios_switchcontroller_fortilinksettings`
* **New Resource:** `fortios_system_acme`
* **New Resource:** `fortios_system_dns64`
* **New Resource:** `fortios_system_fortiai`
* **New Resource:** `fortios_system_ike`
* **New Resource:** `fortios_system_ipam`
* **New Resource:** `fortios_system_ltemodem`
* **New Resource:** `fortios_system_modem`
* **New Resource:** `fortios_system_npu`
* **New Resource:** `fortios_system_physicalswitch`
* **New Resource:** `fortios_system_ssoforticloudadmin`
* **New Resource:** `fortios_system_stp`
* **New Resource:** `fortios_system3gmodem_custom`
* **New Resource:** `fortios_user_certificate`
* **New Resource:** `fortios_videofilter_profile`
* **New Resource:** `fortios_videofilter_youtubechannelfilter`
* **New Resource:** `fortios_videofilter_youtubekey`
* **New Resource:** `fortios_vpnipsec_fec`
* **New Resource:** `fortios_vpnssl_client`
* **New Resource:** `fortios_wirelesscontroller_nacprofile`
* **New Resource:** `fortios_wirelesscontroller_ssidpolicy`
* **New Resource:** `fortios_wirelesscontroller_syslogprofile`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_anqpvenueurl`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_h2qpadviceofcharge`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_h2qposuprovidernai`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_h2qptermsandconditions`

## 1.13.2 (Sep 28, 2021)

BUG FIXES:

* Fix parameter value's length limitation issue

## 1.13.0 (Jun 12, 2021)

IMPROVEMENTS:

* Support ipmask/cidr for system/interface/ip
* Support vdom for each resource

FEATURES:

* **New Resource:** `fortios_firewall_securitypolicy`
* **New Resource:** `fortios_system_virtualswitch`
* **New Resource:** `fortios_routerbgp_network6`
* **New Resource:** `fortios_routerospf_ospfinterface`
* **New Resource:** `fortios_routerospf_network`
* **New Resource:** `fortios_routerospf_neighbor`
* **New Resource:** `fortios_routerospf6_ospf6interface`
* **New Resource:** `resource_firewall_centralsnatmap_move`
* **New Resource:** `resource_firewall_centralsnatmap_sort`
* **New Resource:** `resource_firewall_proxypolicy_move`
* **New Resource:** `resource_firewall_proxypolicy_sort`
* **New Data Source:** `fortios_ipmask_cidr`
* **New Data Source:** `fortios_firewall_centralsnatmap`
* **New Data Source:** `fortios_firewall_centralsnatmaplist`
* **New Data Source:** `fortios_firewall_proxypolicy`
* **New Data Source:** `fortios_firewall_proxypolicylist`

## 1.12.1 (Jun 5, 2021)

BUG FIXES:

* Fix fortios_certificate_local documentation error

## 1.12.0 (Jun 3, 2021)

IMPROVEMENTS:

* Adjust instructions website/docs/r/fortios_system_ha.html.markdown
* Adjust instructions website/docs/r/fortios_system_fortimanager.html.markdown
* Adjust instructions website/docs/r/fortios_router_accesslist.html.markdown
* Adjust instructions website/docs/r/fortios_certificate_local.html.markdown
* Adjust function structure for vendor/github.com/fortinetdev/forti-sdk-go/fortios/sdkcore/sdkfos.go
* Adjust function structure for fortios/resource_routerbgp_network.go
* Update READM.md
* Update policy seq/sort docs
* Remove redundant statement #148
* Fix wrong spelling in fgt_policysort.html.md

FEATURES:

* **New Resource:** `fortios_routerbgp_network`

## 1.11.0 (Mar 15, 2021)

IMPROVEMENTS:

* Add example for fortios_firewall_sslsshprofile #147
* Add options for fortios_alertemail_setting
* Add options for fortios_antivirus_heuristic
* Add options for fortios_antivirus_profile
* Add options for fortios_antivirus_quarantine
* Add options for fortios_antivirus_settings
* Add options for fortios_application_group
* Add options for fortios_application_list
* Add options for fortios_authentication_rule
* Add options for fortios_authentication_scheme
* Add options for fortios_authentication_setting
* Add options for fortios_certificate_ca
* Add options for fortios_certificate_crl
* Add options for fortios_certificate_remote
* Add options for fortios_cifs_profile
* Add options for fortios_dlp_filepattern
* Add options for fortios_dlp_fpdocsource
* Add options for fortios_dlp_sensor
* Add options for fortios_dlp_settings
* Add options for fortios_dnsfilter_domainfilter
* Add options for fortios_dnsfilter_profile
* Add options for fortios_dpdk_global
* Add options for fortios_emailfilter_blockallowlist
* Add options for fortios_emailfilter_bwl
* Add options for fortios_emailfilter_bword
* Add options for fortios_emailfilter_dnsbl
* Add options for fortios_emailfilter_fortishield
* Add options for fortios_emailfilter_iptrust
* Add options for fortios_emailfilter_mheader
* Add options for fortios_emailfilter_profile
* Add options for fortios_endpointcontrol_fctems
* Add options for fortios_endpointcontrol_forticlientems
* Add options for fortios_endpointcontrol_profile
* Add options for fortios_endpointcontrol_settings
* Add options for fortios_extendercontroller_dataplan
* Add options for fortios_extendercontroller_extender
* Add options for fortios_extendercontroller_extender1
* Add options for fortios_filefilter_profile
* Add options for fortios_firewall_DoSpolicy
* Add options for fortios_firewall_DoSpolicy6
* Add options for fortios_firewall_address
* Add options for fortios_firewall_address6
* Add options for fortios_firewall_address6template
* Add options for fortios_firewall_addrgrp
* Add options for fortios_firewall_addrgrp6
* Add options for fortios_firewall_centralsnatmap
* Add options for fortios_firewall_decryptedtrafficmirror
* Add options for fortios_firewall_interfacepolicy
* Add options for fortios_firewall_interfacepolicy6
* Add options for fortios_firewall_internetservice
* Add options for fortios_firewall_internetservicegroup
* Add options for fortios_firewall_internetservicename
* Add options for fortios_firewall_ippool
* Add options for fortios_firewall_iptranslation
* Add options for fortios_firewall_ipv6ehfilter
* Add options for fortios_firewall_localinpolicy
* Add options for fortios_firewall_localinpolicy6
* Add options for fortios_firewall_multicastaddress
* Add options for fortios_firewall_multicastaddress6
* Add options for fortios_firewall_multicastpolicy
* Add options for fortios_firewall_multicastpolicy6
* Add options for fortios_firewall_policy
* Add options for fortios_firewall_policy46
* Add options for fortios_firewall_policy6
* Add options for fortios_firewall_policy64
* Add options for fortios_firewall_profileprotocoloptions
* Add options for fortios_firewall_proxyaddress
* Add options for fortios_firewall_proxyaddrgrp
* Add options for fortios_firewall_proxypolicy
* Add options for fortios_firewall_shapingpolicy
* Add options for fortios_firewall_shapingprofile
* Add options for fortios_firewall_sniffer
* Add options for fortios_firewall_sslserver
* Add options for fortios_firewall_sslsshprofile
* Add options for fortios_firewall_ttlpolicy
* Add options for fortios_firewall_vip
* Add options for fortios_firewall_vip46
* Add options for fortios_firewall_vip6
* Add options for fortios_firewall_vip64
* Add options for fortios_firewallconsolidated_policy
* Add options for fortios_firewallipmacbinding_setting
* Add options for fortios_firewallipmacbinding_table
* Add options for fortios_firewallschedule_recurring
* Add options for fortios_firewallservice_custom
* Add options for fortios_firewallservice_group
* Add options for fortios_firewallshaper_peripshaper
* Add options for fortios_firewallshaper_trafficshaper
* Add options for fortios_firewallssh_hostkey
* Add options for fortios_firewallssh_localca
* Add options for fortios_firewallssh_localkey
* Add options for fortios_firewallssh_setting
* Add options for fortios_firewallssl_setting
* Add options for fortios_firewallwildcardfqdn_custom
* Add options for fortios_firewallwildcardfqdn_group
* Add options for fortios_ftpproxy_explicit
* Add options for fortios_icap_profile
* Add options for fortios_icap_server
* Add options for fortios_ips_custom
* Add options for fortios_ips_global
* Add options for fortios_ips_rule
* Add options for fortios_ips_sensor
* Add options for fortios_log_eventfilter
* Add options for fortios_log_guidisplay
* Add options for fortios_log_setting
* Add options for fortios_log_threatweight
* Add options for fortios_logdisk_filter
* Add options for fortios_logdisk_setting
* Add options for fortios_logfortianalyzer2_filter
* Add options for fortios_logfortianalyzer2_overridefilter
* Add options for fortios_logfortianalyzer2_overridesetting
* Add options for fortios_logfortianalyzer2_setting
* Add options for fortios_logfortianalyzer3_filter
* Add options for fortios_logfortianalyzer3_overridefilter
* Add options for fortios_logfortianalyzer3_overridesetting
* Add options for fortios_logfortianalyzer3_setting
* Add options for fortios_logfortianalyzer_filter
* Add options for fortios_logfortianalyzer_overridefilter
* Add options for fortios_logfortianalyzer_overridesetting
* Add options for fortios_logfortianalyzer_setting
* Add options for fortios_logfortianalyzercloud_filter
* Add options for fortios_logfortianalyzercloud_overridefilter
* Add options for fortios_logfortianalyzercloud_overridesetting
* Add options for fortios_logfortianalyzercloud_setting
* Add options for fortios_logfortiguard_filter
* Add options for fortios_logfortiguard_overridefilter
* Add options for fortios_logfortiguard_overridesetting
* Add options for fortios_logfortiguard_setting
* Add options for fortios_logmemory_filter
* Add options for fortios_logmemory_setting
* Add options for fortios_lognulldevice_filter
* Add options for fortios_lognulldevice_setting
* Add options for fortios_logsyslogd2_filter
* Add options for fortios_logsyslogd2_overridefilter
* Add options for fortios_logsyslogd2_overridesetting
* Add options for fortios_logsyslogd2_setting
* Add options for fortios_logsyslogd3_filter
* Add options for fortios_logsyslogd3_overridefilter
* Add options for fortios_logsyslogd3_overridesetting
* Add options for fortios_logsyslogd3_setting
* Add options for fortios_logsyslogd4_filter
* Add options for fortios_logsyslogd4_overridefilter
* Add options for fortios_logsyslogd4_overridesetting
* Add options for fortios_logsyslogd4_setting
* Add options for fortios_logsyslogd_filter
* Add options for fortios_logsyslogd_overridefilter
* Add options for fortios_logsyslogd_overridesetting
* Add options for fortios_logsyslogd_setting
* Add options for fortios_logwebtrends_filter
* Add options for fortios_logwebtrends_setting
* Add options for fortios_nsxt_setting
* Add options for fortios_report_chart
* Add options for fortios_report_dataset
* Add options for fortios_report_layout
* Add options for fortios_report_setting
* Add options for fortios_report_style
* Add options for fortios_report_theme
* Add options for fortios_router_accesslist6
* Add options for fortios_router_aspathlist
* Add options for fortios_router_bgp
* Add options for fortios_router_communitylist
* Add options for fortios_router_isis
* Add options for fortios_router_multicast
* Add options for fortios_router_multicast6
* Add options for fortios_router_ospf
* Add options for fortios_router_ospf6
* Add options for fortios_router_policy
* Add options for fortios_router_policy6
* Add options for fortios_router_prefixlist
* Add options for fortios_router_prefixlist6
* Add options for fortios_router_rip
* Add options for fortios_router_ripng
* Add options for fortios_router_routemap
* Add options for fortios_router_static
* Add options for fortios_router_static6
* Add options for fortios_routerbgp_neighbor
* Add options for fortios_spamfilter_bwl
* Add options for fortios_spamfilter_bword
* Add options for fortios_spamfilter_dnsbl
* Add options for fortios_spamfilter_fortishield
* Add options for fortios_spamfilter_iptrust
* Add options for fortios_spamfilter_mheader
* Add options for fortios_spamfilter_profile
* Add options for fortios_sshfilter_profile
* Add options for fortios_switchcontroller_8021Xsettings
* Add options for fortios_switchcontroller_flowtracking
* Add options for fortios_switchcontroller_global
* Add options for fortios_switchcontroller_igmpsnooping
* Add options for fortios_switchcontroller_lldpprofile
* Add options for fortios_switchcontroller_lldpsettings
* Add options for fortios_switchcontroller_location
* Add options for fortios_switchcontroller_managedswitch
* Add options for fortios_switchcontroller_nacdevice
* Add options for fortios_switchcontroller_nacsettings
* Add options for fortios_switchcontroller_networkmonitorsettings
* Add options for fortios_switchcontroller_portpolicy
* Add options for fortios_switchcontroller_quarantine
* Add options for fortios_switchcontroller_remotelog
* Add options for fortios_switchcontroller_snmpcommunity
* Add options for fortios_switchcontroller_snmpsysinfo
* Add options for fortios_switchcontroller_snmpuser
* Add options for fortios_switchcontroller_stormcontrol
* Add options for fortios_switchcontroller_stormcontrolpolicy
* Add options for fortios_switchcontroller_stpsettings
* Add options for fortios_switchcontroller_switchlog
* Add options for fortios_switchcontroller_switchprofile
* Add options for fortios_switchcontroller_system
* Add options for fortios_switchcontroller_trafficpolicy
* Add options for fortios_switchcontroller_trafficsniffer
* Add options for fortios_switchcontroller_vlan
* Add options for fortios_switchcontroller_vlanpolicy
* Add options for fortios_switchcontrollerautoconfig_policy
* Add options for fortios_switchcontrollerinitialconfig_template
* Add options for fortios_switchcontrollerptp_policy
* Add options for fortios_switchcontrollerptp_settings
* Add options for fortios_switchcontrollerqos_dot1pmap
* Add options for fortios_switchcontrollerqos_ipdscpmap
* Add options for fortios_switchcontrollerqos_queuepolicy
* Add options for fortios_switchcontrollersecuritypolicy_8021X
* Add options for fortios_switchcontrollersecuritypolicy_captiveportal
* Add options for fortios_switchcontrollersecuritypolicy_localaccess
* Add options for fortios_system_accprofile
* Add options for fortios_system_admin
* Add options for fortios_system_alarm
* Add options for fortios_system_apiuser
* Add options for fortios_system_autoinstall
* Add options for fortios_system_automationaction
* Add options for fortios_system_automationdestination
* Add options for fortios_system_automationstitch
* Add options for fortios_system_automationtrigger
* Add options for fortios_system_autoscript
* Add options for fortios_system_centralmanagement
* Add options for fortios_system_clustersync
* Add options for fortios_system_console
* Add options for fortios_system_csf
* Add options for fortios_system_ddns
* Add options for fortios_system_dedicatedmgmt
* Add options for fortios_system_dns
* Add options for fortios_system_dnsdatabase
* Add options for fortios_system_dnsserver
* Add options for fortios_system_dscpbasedpriority
* Add options for fortios_system_emailserver
* Add options for fortios_system_externalresource
* Add options for fortios_system_federatedupgrade
* Add options for fortios_system_fipscc
* Add options for fortios_system_fm
* Add options for fortios_system_fortiguard
* Add options for fortios_system_fortisandbox
* Add options for fortios_system_fssopolling
* Add options for fortios_system_ftmpush
* Add options for fortios_system_geneve
* Add options for fortios_system_global
* Add options for fortios_system_gretunnel
* Add options for fortios_system_hamonitor
* Add options for fortios_system_interface
* Add options for fortios_system_ips
* Add options for fortios_system_ipsurlfilterdns
* Add options for fortios_system_ipsurlfilterdns6
* Add options for fortios_system_linkmonitor
* Add options for fortios_system_managementtunnel
* Add options for fortios_system_mobiletunnel
* Add options for fortios_system_nat64
* Add options for fortios_system_ndproxy
* Add options for fortios_system_networkvisibility
* Add options for fortios_system_ntp
* Add options for fortios_system_objecttagging
* Add options for fortios_system_passwordpolicy
* Add options for fortios_system_passwordpolicyguestadmin
* Add options for fortios_system_pppoeinterface
* Add options for fortios_system_proberesponse
* Add options for fortios_system_ptp
* Add options for fortios_system_replacemsggroup
* Add options for fortios_system_replacemsgimage
* Add options for fortios_system_saml
* Add options for fortios_system_sdnconnector
* Add options for fortios_system_sdwan
* Add options for fortios_system_settings
* Add options for fortios_system_speedtestschedule
* Add options for fortios_system_standalonecluster
* Add options for fortios_system_storage
* Add options for fortios_system_switchinterface
* Add options for fortios_system_tosbasedpriority
* Add options for fortios_system_vdomdns
* Add options for fortios_system_vdomexception
* Add options for fortios_system_vdomlink
* Add options for fortios_system_vdomnetflow
* Add options for fortios_system_vdomradiusserver
* Add options for fortios_system_vdomsflow
* Add options for fortios_system_virtualwanlink
* Add options for fortios_system_virtualwirepair
* Add options for fortios_system_vnetunnel
* Add options for fortios_system_vxlan
* Add options for fortios_system_wccp
* Add options for fortios_system_zone
* Add options for fortios_systemautoupdate_pushupdate
* Add options for fortios_systemautoupdate_schedule
* Add options for fortios_systemautoupdate_tunneling
* Add options for fortios_systemdhcp6_server
* Add options for fortios_systemdhcp_server
* Add options for fortios_systemlldp_networkpolicy
* Add options for fortios_systemreplacemsg_admin
* Add options for fortios_systemreplacemsg_alertmail
* Add options for fortios_systemreplacemsg_auth
* Add options for fortios_systemreplacemsg_automation
* Add options for fortios_systemreplacemsg_devicedetectionportal
* Add options for fortios_systemreplacemsg_ec
* Add options for fortios_systemreplacemsg_fortiguardwf
* Add options for fortios_systemreplacemsg_ftp
* Add options for fortios_systemreplacemsg_http
* Add options for fortios_systemreplacemsg_icap
* Add options for fortios_systemreplacemsg_mail
* Add options for fortios_systemreplacemsg_nacquar
* Add options for fortios_systemreplacemsg_nntp
* Add options for fortios_systemreplacemsg_spam
* Add options for fortios_systemreplacemsg_sslvpn
* Add options for fortios_systemreplacemsg_trafficquota
* Add options for fortios_systemreplacemsg_utm
* Add options for fortios_systemreplacemsg_webproxy
* Add options for fortios_systemsnmp_community
* Add options for fortios_systemsnmp_sysinfo
* Add options for fortios_systemsnmp_user
* Add options for fortios_user_device
* Add options for fortios_user_deviceaccesslist
* Add options for fortios_user_exchange
* Add options for fortios_user_fortitoken
* Add options for fortios_user_fsso
* Add options for fortios_user_fssopolling
* Add options for fortios_user_group
* Add options for fortios_user_krbkeytab
* Add options for fortios_user_ldap
* Add options for fortios_user_local
* Add options for fortios_user_nacpolicy
* Add options for fortios_user_passwordpolicy
* Add options for fortios_user_peer
* Add options for fortios_user_pop3
* Add options for fortios_user_quarantine
* Add options for fortios_user_radius
* Add options for fortios_user_setting
* Add options for fortios_user_tacacs
* Add options for fortios_voip_profile
* Add options for fortios_vpn_l2tp
* Add options for fortios_vpn_ocvpn
* Add options for fortios_vpn_pptp
* Add options for fortios_vpncertificate_ca
* Add options for fortios_vpncertificate_crl
* Add options for fortios_vpncertificate_local
* Add options for fortios_vpncertificate_ocspserver
* Add options for fortios_vpncertificate_remote
* Add options for fortios_vpncertificate_setting
* Add options for fortios_vpnipsec_concentrator
* Add options for fortios_vpnipsec_forticlient
* Add options for fortios_vpnipsec_manualkey
* Add options for fortios_vpnipsec_manualkeyinterface
* Add options for fortios_vpnipsec_phase1
* Add options for fortios_vpnipsec_phase1interface
* Add options for fortios_vpnipsec_phase2
* Add options for fortios_vpnipsec_phase2interface
* Add options for fortios_vpnssl_settings
* Add options for fortios_vpnsslweb_hostchecksoftware
* Add options for fortios_vpnsslweb_portal
* Add options for fortios_vpnsslweb_realm
* Add options for fortios_vpnsslweb_userbookmark
* Add options for fortios_vpnsslweb_usergroupbookmark
* Add options for fortios_waf_profile
* Add options for fortios_wanopt_authgroup
* Add options for fortios_wanopt_cacheservice
* Add options for fortios_wanopt_contentdeliverynetworkrule
* Add options for fortios_wanopt_profile
* Add options for fortios_wanopt_remotestorage
* Add options for fortios_wanopt_settings
* Add options for fortios_wanopt_webcache
* Add options for fortios_webfilter_content
* Add options for fortios_webfilter_contentheader
* Add options for fortios_webfilter_fortiguard
* Add options for fortios_webfilter_ftgdlocalcat
* Add options for fortios_webfilter_ftgdlocalrating
* Add options for fortios_webfilter_override
* Add options for fortios_webfilter_profile
* Add options for fortios_webfilter_searchengine
* Add options for fortios_webfilter_urlfilter
* Add options for fortios_webproxy_debugurl
* Add options for fortios_webproxy_explicit
* Add options for fortios_webproxy_forwardserver
* Add options for fortios_webproxy_forwardservergroup
* Add options for fortios_webproxy_global
* Add options for fortios_webproxy_profile
* Add options for fortios_webproxy_urlmatch
* Add options for fortios_wirelesscontroller_accesscontrollist
* Add options for fortios_wirelesscontroller_address
* Add options for fortios_wirelesscontroller_addrgrp
* Add options for fortios_wirelesscontroller_apcfgprofile
* Add options for fortios_wirelesscontroller_apstatus
* Add options for fortios_wirelesscontroller_bleprofile
* Add options for fortios_wirelesscontroller_bonjourprofile
* Add options for fortios_wirelesscontroller_global
* Add options for fortios_wirelesscontroller_intercontroller
* Add options for fortios_wirelesscontroller_log
* Add options for fortios_wirelesscontroller_mpskprofile
* Add options for fortios_wirelesscontroller_qosprofile
* Add options for fortios_wirelesscontroller_region
* Add options for fortios_wirelesscontroller_setting
* Add options for fortios_wirelesscontroller_snmp
* Add options for fortios_wirelesscontroller_timers
* Add options for fortios_wirelesscontroller_utmprofile
* Add options for fortios_wirelesscontroller_vap
* Add options for fortios_wirelesscontroller_wagprofile
* Add options for fortios_wirelesscontroller_widsprofile
* Add options for fortios_wirelesscontroller_wtp
* Add options for fortios_wirelesscontroller_wtpprofile
* Add options for fortios_wirelesscontrollerhotspot20_anqpipaddresstype
* Add options for fortios_wirelesscontrollerhotspot20_anqpnairealm
* Add options for fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype
* Add options for fortios_wirelesscontrollerhotspot20_h2qpconncapability
* Add options for fortios_wirelesscontrollerhotspot20_h2qposuprovider
* Add options for fortios_wirelesscontrollerhotspot20_h2qpwanmetric
* Add options for fortios_wirelesscontrollerhotspot20_hsprofile
* Add options for fortios_wirelesscontrollerhotspot20_icon

BUG FIXES:

* Fix multiple ports in fortios_firewall_sslsshprofile #148

## 1.10.4 (Feb 26, 2021)

IMPROVEMENTS:

* Improve request error message

## 1.10.3 (Feb 24, 2021)

IMPROVEMENTS:

* Improve fgt_filter documentation
* Add Example for fortios_systemreplacemsg_admin

BUG FIXES:

* Add default for fortios_firewall_policy schedule #140
* Make resource_router_static dst optional #141
* Update release version info according to the changes in FOS version
* Fix system replacemsg issues #139

## 1.10.2 (Feb 11, 2021)

IMPROVEMENTS:

* Add buildURLWithoutVdom

## 1.10.1 (Feb 10, 2021)

BUG FIXES:

* Resync

## 1.10.0 (Feb 10, 2021)

FEATURES:

* Support FortiGate terraform access using PKI group
* Support fmg_firewall_security_policy for FMG 6.4
* Support force_recreate for fortios_json_generic_api

## 1.9.0 (Feb 3, 2021)

FEATURES:

* **New Data Source:** `data_source_firewall_DoSpolicy`
* **New Data Source:** `data_source_firewall_DoSpolicy6`
* **New Data Source:** `data_source_firewall_DoSpolicy6list`
* **New Data Source:** `data_source_firewall_DoSpolicylist`
* **New Data Source:** `data_source_firewall_address6template`
* **New Data Source:** `data_source_firewall_address6templatelist`
* **New Data Source:** `data_source_firewall_internetservice`
* **New Data Source:** `data_source_firewall_internetservicecustom`
* **New Data Source:** `data_source_firewall_internetservicecustomgroup`
* **New Data Source:** `data_source_firewall_internetservicecustomgrouplist`
* **New Data Source:** `data_source_firewall_internetservicecustomlist`
* **New Data Source:** `data_source_firewall_internetservicedefinition`
* **New Data Source:** `data_source_firewall_internetservicedefinitionlist`
* **New Data Source:** `data_source_firewall_internetserviceextension`
* **New Data Source:** `data_source_firewall_internetserviceextensionlist`
* **New Data Source:** `data_source_firewall_internetservicegroup`
* **New Data Source:** `data_source_firewall_internetservicegrouplist`
* **New Data Source:** `data_source_firewall_internetservicelist`
* **New Data Source:** `data_source_firewall_ipv6ehfilter`
* **New Data Source:** `data_source_firewall_multicastaddress`
* **New Data Source:** `data_source_firewall_multicastaddress6`
* **New Data Source:** `data_source_firewall_multicastaddress6list`
* **New Data Source:** `data_source_firewall_multicastaddresslist`
* **New Data Source:** `data_source_firewall_policy46`
* **New Data Source:** `data_source_firewall_policy46list`
* **New Data Source:** `data_source_firewall_profileprotocoloptions`
* **New Data Source:** `data_source_firewall_profileprotocoloptionslist`
* **New Data Source:** `data_source_firewall_proxyaddress`
* **New Data Source:** `data_source_firewall_proxyaddresslist`
* **New Data Source:** `data_source_firewall_proxyaddrgrp`
* **New Data Source:** `data_source_firewall_proxyaddrgrplist`
* **New Data Source:** `data_source_firewallconsolidated_policy`
* **New Data Source:** `data_source_firewallconsolidated_policylist`
* **New Data Source:** `data_source_firewallschedule_group`
* **New Data Source:** `data_source_firewallschedule_grouplist`
* **New Data Source:** `data_source_firewallschedule_onetime`
* **New Data Source:** `data_source_firewallschedule_onetimelist`
* **New Data Source:** `data_source_firewallschedule_recurring`
* **New Data Source:** `data_source_firewallschedule_recurringlist`
* **New Data Source:** `data_source_firewallservice_category`
* **New Data Source:** `data_source_firewallservice_categorylist`
* **New Data Source:** `data_source_firewallservice_custom`
* **New Data Source:** `data_source_firewallservice_customlist`
* **New Data Source:** `data_source_firewallservice_group`
* **New Data Source:** `data_source_firewallservice_grouplist`
* **New Data Source:** `data_source_firewallshaper_peripshaper`
* **New Data Source:** `data_source_firewallshaper_peripshaperlist`
* **New Data Source:** `data_source_firewallshaper_trafficshaper`
* **New Data Source:** `data_source_firewallshaper_trafficshaperlist`
* **New Data Source:** `data_source_firewallwildcardfqdn_custom`
* **New Data Source:** `data_source_firewallwildcardfqdn_customlist`
* **New Data Source:** `data_source_firewallwildcardfqdn_group`
* **New Data Source:** `data_source_firewallwildcardfqdn_grouplist`
* **New Data Source:** `data_source_system_accprofile`
* **New Data Source:** `data_source_system_accprofilelist`
* **New Data Source:** `data_source_system_admin`
* **New Data Source:** `data_source_system_adminlist`
* **New Data Source:** `data_source_system_alias`
* **New Data Source:** `data_source_system_aliaslist`
* **New Data Source:** `data_source_system_apiuser`
* **New Data Source:** `data_source_system_apiuserlist`
* **New Data Source:** `data_source_system_arptable`
* **New Data Source:** `data_source_system_arptablelist`
* **New Data Source:** `data_source_system_autoinstall`
* **New Data Source:** `data_source_system_automationaction`
* **New Data Source:** `data_source_system_automationactionlist`
* **New Data Source:** `data_source_system_automationdestination`
* **New Data Source:** `data_source_system_automationdestinationlist`
* **New Data Source:** `data_source_system_automationtrigger`
* **New Data Source:** `data_source_system_automationtriggerlist`
* **New Data Source:** `data_source_system_autoscript`
* **New Data Source:** `data_source_system_autoscriptlist`
* **New Data Source:** `data_source_system_centralmanagement`
* **New Data Source:** `data_source_system_clustersync`
* **New Data Source:** `data_source_system_clustersynclist`
* **New Data Source:** `data_source_system_console`
* **New Data Source:** `data_source_system_csf`
* **New Data Source:** `data_source_system_ddns`
* **New Data Source:** `data_source_system_ddnslist`
* **New Data Source:** `data_source_system_dns`
* **New Data Source:** `data_source_system_dnsdatabase`
* **New Data Source:** `data_source_system_dnsdatabaselist`
* **New Data Source:** `data_source_system_dnsserver`
* **New Data Source:** `data_source_system_dnsserverlist`
* **New Data Source:** `data_source_system_dscpbasedpriority`
* **New Data Source:** `data_source_system_dscpbasedprioritylist`
* **New Data Source:** `data_source_system_emailserver`
* **New Data Source:** `data_source_system_externalresource`
* **New Data Source:** `data_source_system_externalresourcelist`
* **New Data Source:** `data_source_system_fipscc`
* **New Data Source:** `data_source_system_fm`
* **New Data Source:** `data_source_system_fortiguard`
* **New Data Source:** `data_source_system_fortimanager`
* **New Data Source:** `data_source_system_fortisandbox`
* **New Data Source:** `data_source_system_fssopolling`
* **New Data Source:** `data_source_system_ftmpush`
* **New Data Source:** `data_source_system_gretunnel`
* **New Data Source:** `data_source_system_gretunnellist`
* **New Data Source:** `data_source_system_ha`
* **New Data Source:** `data_source_system_hamonitor`
* **New Data Source:** `data_source_system_ipiptunnel`
* **New Data Source:** `data_source_system_ipiptunnellist`
* **New Data Source:** `data_source_system_ipv6neighborcache`
* **New Data Source:** `data_source_system_ipv6neighborcachelist`
* **New Data Source:** `data_source_system_ipv6tunnel`
* **New Data Source:** `data_source_system_ipv6tunnellist`
* **New Data Source:** `data_source_system_linkmonitor`
* **New Data Source:** `data_source_system_linkmonitorlist`
* **New Data Source:** `data_source_system_managementtunnel`
* **New Data Source:** `data_source_system_mobiletunnel`
* **New Data Source:** `data_source_system_mobiletunnellist`
* **New Data Source:** `data_source_system_nat64`
* **New Data Source:** `data_source_system_ndproxy`
* **New Data Source:** `data_source_system_netflow`
* **New Data Source:** `data_source_system_networkvisibility`
* **New Data Source:** `data_source_system_ntp`
* **New Data Source:** `data_source_system_objecttagging`
* **New Data Source:** `data_source_system_objecttagginglist`
* **New Data Source:** `data_source_system_passwordpolicy`
* **New Data Source:** `data_source_system_passwordpolicyguestadmin`
* **New Data Source:** `data_source_system_pppoeinterface`
* **New Data Source:** `data_source_system_pppoeinterfacelist`
* **New Data Source:** `data_source_system_proberesponse`
* **New Data Source:** `data_source_system_proxyarp`
* **New Data Source:** `data_source_system_proxyarplist`
* **New Data Source:** `data_source_system_replacemsggroup`
* **New Data Source:** `data_source_system_replacemsggrouplist`
* **New Data Source:** `data_source_system_replacemsgimage`
* **New Data Source:** `data_source_system_replacemsgimagelist`
* **New Data Source:** `data_source_system_resourcelimits`
* **New Data Source:** `data_source_system_sdnconnector`
* **New Data Source:** `data_source_system_sdnconnectorlist`
* **New Data Source:** `data_source_system_sessionhelper`
* **New Data Source:** `data_source_system_sessionhelperlist`
* **New Data Source:** `data_source_system_sessionttl`
* **New Data Source:** `data_source_system_sflow`
* **New Data Source:** `data_source_system_sittunnel`
* **New Data Source:** `data_source_system_sittunnellist`
* **New Data Source:** `data_source_system_smsserver`
* **New Data Source:** `data_source_system_smsserverlist`
* **New Data Source:** `data_source_system_tosbasedpriority`
* **New Data Source:** `data_source_system_tosbasedprioritylist`
* **New Data Source:** `data_source_system_vdomexception`
* **New Data Source:** `data_source_system_vdomexceptionlist`
* **New Data Source:** `data_source_system_vdomnetflow`
* **New Data Source:** `data_source_system_vdomsflow`
* **New Data Source:** `data_source_system_virtualwanlink`
* **New Data Source:** `data_source_system_vxlan`
* **New Data Source:** `data_source_system_vxlanlist`
* **New Data Source:** `data_source_system_wccp`
* **New Data Source:** `data_source_system_wccplist`
* **New Data Source:** `data_source_system_zone`
* **New Data Source:** `data_source_system_zonelist`
* **New Data Source:** `data_source_systemautoupdate_pushupdate`
* **New Data Source:** `data_source_systemautoupdate_schedule`
* **New Data Source:** `data_source_systemautoupdate_tunneling`
* **New Data Source:** `data_source_systemdhcp_server`
* **New Data Source:** `data_source_systemdhcp_serverlist`
* **New Data Source:** `data_source_systemlldp_networkpolicy`
* **New Data Source:** `data_source_systemlldp_networkpolicylist`
* **New Data Source:** `data_source_systemsnmp_community`
* **New Data Source:** `data_source_systemsnmp_communitylist`
* **New Data Source:** `data_source_systemsnmp_sysinfo`
* **New Data Source:** `data_source_systemsnmp_user`
* **New Data Source:** `data_source_systemsnmp_userlist`
* **New Data Source:** `data_source_user_saml`
* **New Data Source:** `data_source_user_samllist`
* **New Resource:** `resource_certificate_remote`
* **New Resource:** `resource_cifs_domaincontroller`
* **New Resource:** `resource_cifs_profile`
* **New Resource:** `resource_credentialstore_domaincontroller`
* **New Resource:** `resource_dlp_sensitivity`
* **New Resource:** `resource_dpdk_cpus`
* **New Resource:** `resource_dpdk_global`
* **New Resource:** `resource_emailfilter_blockallowlist`
* **New Resource:** `resource_emailfilter_bwl`
* **New Resource:** `resource_emailfilter_bword`
* **New Resource:** `resource_emailfilter_dnsbl`
* **New Resource:** `resource_emailfilter_fortishield`
* **New Resource:** `resource_emailfilter_iptrust`
* **New Resource:** `resource_emailfilter_mheader`
* **New Resource:** `resource_emailfilter_options`
* **New Resource:** `resource_emailfilter_profile`
* **New Resource:** `resource_endpointcontrol_fctems`
* **New Resource:** `resource_extendercontroller_dataplan`
* **New Resource:** `resource_filefilter_profile`
* **New Resource:** `resource_firewall_city`
* **New Resource:** `resource_firewall_country`
* **New Resource:** `resource_firewall_decryptedtrafficmirror`
* **New Resource:** `resource_firewall_internetserviceaddition`
* **New Resource:** `resource_firewall_internetserviceappend`
* **New Resource:** `resource_firewall_internetservicebotnet`
* **New Resource:** `resource_firewall_internetserviceipblreason`
* **New Resource:** `resource_firewall_internetserviceipblvendor`
* **New Resource:** `resource_firewall_internetservicelist`
* **New Resource:** `resource_firewall_internetservicename`
* **New Resource:** `resource_firewall_internetserviceowner`
* **New Resource:** `resource_firewall_internetservicereputation`
* **New Resource:** `resource_firewall_region`
* **New Resource:** `resource_firewall_trafficclass`
* **New Resource:** `resource_firewall_vendormac`
* **New Resource:** `resource_ips_viewmap`
* **New Resource:** `resource_logfortianalyzercloud_filter`
* **New Resource:** `resource_logfortianalyzercloud_overridefilter`
* **New Resource:** `resource_logfortianalyzercloud_overridesetting`
* **New Resource:** `resource_logfortianalyzercloud_setting`
* **New Resource:** `resource_nsxt_servicechain`
* **New Resource:** `resource_nsxt_setting`
* **New Resource:** `resource_switchcontroller_flowtracking`
* **New Resource:** `resource_switchcontroller_location`
* **New Resource:** `resource_switchcontroller_nacdevice`
* **New Resource:** `resource_switchcontroller_nacsettings`
* **New Resource:** `resource_switchcontroller_portpolicy`
* **New Resource:** `resource_switchcontroller_remotelog`
* **New Resource:** `resource_switchcontroller_snmpcommunity`
* **New Resource:** `resource_switchcontroller_snmpsysinfo`
* **New Resource:** `resource_switchcontroller_snmptrapthreshold`
* **New Resource:** `resource_switchcontroller_snmpuser`
* **New Resource:** `resource_switchcontroller_stormcontrolpolicy`
* **New Resource:** `resource_switchcontroller_stpinstance`
* **New Resource:** `resource_switchcontroller_trafficsniffer`
* **New Resource:** `resource_switchcontroller_vlanpolicy`
* **New Resource:** `resource_switchcontrollerinitialconfig_template`
* **New Resource:** `resource_switchcontrollerinitialconfig_vlans`
* **New Resource:** `resource_switchcontrollerptp_policy`
* **New Resource:** `resource_switchcontrollerptp_settings`
* **New Resource:** `resource_switchcontrollersecuritypolicy_localaccess`
* **New Resource:** `resource_system_federatedupgrade`
* **New Resource:** `resource_system_geoipcountry`
* **New Resource:** `resource_system_ips`
* **New Resource:** `resource_system_ipsurlfilterdns`
* **New Resource:** `resource_system_ipsurlfilterdns6`
* **New Resource:** `resource_system_sdwan`
* **New Resource:** `resource_system_speedtestschedule`
* **New Resource:** `resource_system_speedtestserver`
* **New Resource:** `resource_system_standalonecluster`
* **New Resource:** `resource_system_vnetunnel`
* **New Resource:** `resource_systemreplacemsg_automation`
* **New Resource:** `resource_user_exchange`
* **New Resource:** `resource_user_nacpolicy`
* **New Resource:** `resource_vpn_ocvpn`
* **New Resource:** `resource_wirelesscontroller_accesscontrollist`
* **New Resource:** `resource_wirelesscontroller_address`
* **New Resource:** `resource_wirelesscontroller_addrgrp`
* **New Resource:** `resource_wirelesscontroller_apcfgprofile`
* **New Resource:** `resource_wirelesscontroller_arrpprofile`
* **New Resource:** `resource_wirelesscontroller_log`
* **New Resource:** `resource_wirelesscontroller_mpskprofile`
* **New Resource:** `resource_wirelesscontroller_snmp`
* **New Resource:** `resource_wirelesscontroller_wagprofile`

IMPROVEMENTS:

* Update data_source_firewall_address to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_firewall_address6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_firewall_addrgrp to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_firewall_policy to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_firewall_policy6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_firewall_policy64 to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_router_bgp to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_router_multicast to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_router_ospf to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_router_ospf6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_router_policy to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_router_routemap to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_router_static to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_router_static6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_routerbgp_neighbor to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_system_global to support FortiOS 6.0 6.2 6.4 6.6.
* Update data_source_system_interface to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_alertemail_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_antivirus_heuristic to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_antivirus_profile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_antivirus_quarantine to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_antivirus_settings to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_application_custom to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_application_group to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_application_list to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_application_name to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_application_rulesettings to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_authentication_rule to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_authentication_scheme to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_authentication_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_certificate_ca to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_certificate_crl to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_certificate_local to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_dlp_filepattern to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_dlp_fpdocsource to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_dlp_fpsensitivity to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_dlp_sensor to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_dlp_settings to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_dnsfilter_domainfilter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_dnsfilter_profile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_endpointcontrol_client to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_endpointcontrol_forticlientems to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_endpointcontrol_forticlientregistrationsync to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_endpointcontrol_profile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_endpointcontrol_registeredforticlient to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_endpointcontrol_settings to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_extendercontroller_extender to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_extendercontroller_extender1 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_DoSpolicy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_DoSpolicy6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_address to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_address6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_address6template to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_addrgrp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_addrgrp6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_authportal to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_centralsnatmap to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_dnstranslation to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_identitybasedroute to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_interfacepolicy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_interfacepolicy6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_internetservice to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_internetservicecustom to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_internetservicecustomgroup to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_internetservicedefinition to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_internetserviceextension to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_internetservicegroup to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_ippool to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_ippool6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_iptranslation to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_ipv6ehfilter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_ldbmonitor to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_localinpolicy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_localinpolicy6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_multicastaddress to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_multicastaddress6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_multicastpolicy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_multicastpolicy6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_policy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_policy46 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_policy6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_policy64 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_profilegroup to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_profileprotocoloptions to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_proxyaddress to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_proxyaddrgrp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_proxypolicy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_shapingpolicy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_shapingprofile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_sniffer to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_sslserver to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_sslsshprofile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_ttlpolicy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_vip to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_vip46 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_vip6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_vip64 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_vipgrp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_vipgrp46 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_vipgrp6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewall_vipgrp64 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallconsolidated_policy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallipmacbinding_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallipmacbinding_table to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallschedule_group to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallschedule_onetime to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallschedule_recurring to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallservice_category to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallservice_custom to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallservice_group to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallshaper_peripshaper to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallshaper_trafficshaper to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallssh_hostkey to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallssh_localca to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallssh_localkey to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallssh_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallssl_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallwildcardfqdn_custom to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_firewallwildcardfqdn_group to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_ftpproxy_explicit to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_icap_profile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_icap_server to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_ips_custom to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_ips_decoder to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_ips_global to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_ips_rule to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_ips_rulesettings to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_ips_sensor to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_ips_settings to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_log_customfield to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_log_eventfilter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_log_guidisplay to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_log_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_log_threatweight to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logdisk_filter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logdisk_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortianalyzer2_filter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortianalyzer2_overridefilter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortianalyzer2_overridesetting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortianalyzer2_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortianalyzer3_filter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortianalyzer3_overridefilter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortianalyzer3_overridesetting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortianalyzer3_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortianalyzer_filter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortianalyzer_overridefilter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortianalyzer_overridesetting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortianalyzer_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortiguard_filter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortiguard_overridefilter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortiguard_overridesetting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logfortiguard_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logmemory_filter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logmemory_globalsetting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logmemory_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_lognulldevice_filter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_lognulldevice_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd2_filter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd2_overridefilter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd2_overridesetting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd2_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd3_filter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd3_overridefilter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd3_overridesetting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd3_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd4_filter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd4_overridefilter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd4_overridesetting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd4_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd_filter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd_overridefilter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd_overridesetting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logsyslogd_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logwebtrends_filter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_logwebtrends_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_report_chart to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_report_dataset to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_report_layout to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_report_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_report_style to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_report_theme to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_accesslist to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_accesslist6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_aspathlist to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_authpath to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_bfd to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_bfd6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_bgp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_communitylist to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_isis to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_keychain to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_multicast to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_multicast6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_multicastflow to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_ospf to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_ospf6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_policy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_policy6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_prefixlist to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_prefixlist6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_rip to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_ripng to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_routemap to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_static to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_router_static6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_routerbgp_neighbor to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_spamfilter_bwl to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_spamfilter_bword to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_spamfilter_dnsbl to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_spamfilter_fortishield to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_spamfilter_iptrust to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_spamfilter_mheader to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_spamfilter_options to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_spamfilter_profile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_sshfilter_profile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_8021Xsettings to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_customcommand to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_global to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_igmpsnooping to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_lldpprofile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_lldpsettings to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_macsyncsettings to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_managedswitch to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_networkmonitorsettings to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_quarantine to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_sflow to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_stormcontrol to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_stpsettings to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_switchgroup to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_switchinterfacetag to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_switchlog to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_switchprofile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_system to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_trafficpolicy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_virtualportpool to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontroller_vlan to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontrollerautoconfig_custom to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontrollerautoconfig_default to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontrollerautoconfig_policy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontrollerqos_dot1pmap to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontrollerqos_ipdscpmap to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontrollerqos_qospolicy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontrollerqos_queuepolicy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontrollersecuritypolicy_8021X to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_switchcontrollersecuritypolicy_captiveportal to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_accprofile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_admin to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_affinityinterrupt to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_affinitypacketredistribution to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_alarm to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_alias to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_apiuser to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_arptable to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_autoinstall to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_automationaction to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_automationdestination to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_automationstitch to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_automationtrigger to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_autoscript to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_centralmanagement to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_clustersync to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_console to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_csf to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_customlanguage to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_ddns to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_dedicatedmgmt to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_dns to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_dnsdatabase to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_dnsserver to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_dscpbasedpriority to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_emailserver to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_externalresource to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_fipscc to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_fm to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_fortiguard to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_fortimanager to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_fortisandbox to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_fssopolling to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_ftmpush to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_geneve to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_geoipoverride to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_global to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_gretunnel to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_ha to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_hamonitor to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_interface to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_ipiptunnel to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_ipsecaggregate to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_ipv6neighborcache to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_ipv6tunnel to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_linkmonitor to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_macaddresstable to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_managementtunnel to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_mobiletunnel to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_nat64 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_ndproxy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_netflow to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_networkvisibility to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_ntp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_objecttagging to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_passwordpolicy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_passwordpolicyguestadmin to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_pppoeinterface to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_proberesponse to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_proxyarp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_ptp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_replacemsggroup to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_replacemsgimage to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_resourcelimits to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_saml to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_sdnconnector to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_sessionhelper to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_sessionttl to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_settings to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_sflow to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_sittunnel to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_smsserver to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_ssoadmin to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_storage to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_switchinterface to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_tosbasedpriority to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_vdom to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_vdomdns to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_vdomexception to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_vdomlink to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_vdomnetflow to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_vdomproperty to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_vdomradiusserver to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_vdomsflow to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_virtualwanlink to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_virtualwirepair to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_vxlan to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_wccp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_system_zone to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemautoupdate_pushupdate to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemautoupdate_schedule to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemautoupdate_tunneling to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemdhcp6_server to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemdhcp_server to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemlldp_networkpolicy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_admin to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_alertmail to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_auth to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_devicedetectionportal to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_ec to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_fortiguardwf to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_ftp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_http to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_icap to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_mail to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_nacquar to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_nntp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_spam to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_sslvpn to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_trafficquota to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_utm to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemreplacemsg_webproxy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemsnmp_community to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemsnmp_sysinfo to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_systemsnmp_user to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_adgrp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_device to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_deviceaccesslist to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_devicecategory to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_devicegroup to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_domaincontroller to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_fortitoken to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_fsso to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_fssopolling to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_group to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_krbkeytab to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_ldap to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_local to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_passwordpolicy to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_peer to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_peergrp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_pop3 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_quarantine to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_radius to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_saml to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_securityexemptlist to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_user_tacacs to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_voip_profile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpn_l2tp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpn_pptp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpncertificate_ca to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpncertificate_crl to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpncertificate_local to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpncertificate_ocspserver to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpncertificate_remote to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpncertificate_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpnipsec_concentrator to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpnipsec_forticlient to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpnipsec_manualkey to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpnipsec_manualkeyinterface to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpnipsec_phase1 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpnipsec_phase1interface to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpnipsec_phase2 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpnipsec_phase2interface to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpnssl_settings to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpnsslweb_hostchecksoftware to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpnsslweb_portal to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpnsslweb_realm to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpnsslweb_userbookmark to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_vpnsslweb_usergroupbookmark to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_waf_mainclass to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_waf_profile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_waf_signature to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_waf_subclass to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wanopt_authgroup to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wanopt_cacheservice to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wanopt_contentdeliverynetworkrule to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wanopt_peer to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wanopt_profile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wanopt_remotestorage to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wanopt_settings to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wanopt_webcache to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webfilter_content to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webfilter_contentheader to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webfilter_fortiguard to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webfilter_ftgdlocalcat to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webfilter_ftgdlocalrating to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webfilter_ipsurlfiltercachesetting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webfilter_ipsurlfiltersetting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webfilter_ipsurlfiltersetting6 to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webfilter_override to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webfilter_profile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webfilter_searchengine to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webfilter_urlfilter to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webproxy_debugurl to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webproxy_explicit to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webproxy_forwardserver to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webproxy_forwardservergroup to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webproxy_global to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webproxy_profile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webproxy_urlmatch to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_webproxy_wisp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_apstatus to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_bleprofile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_bonjourprofile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_global to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_intercontroller to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_qosprofile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_region to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_setting to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_timers to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_utmprofile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_vap to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_vapgroup to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_widsprofile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_wtp to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_wtpgroup to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontroller_wtpprofile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontrollerhotspot20_anqp3gppcellular to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontrollerhotspot20_anqpipaddresstype to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontrollerhotspot20_anqpnairealm to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontrollerhotspot20_anqpnetworkauthtype to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontrollerhotspot20_anqproamingconsortium to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontrollerhotspot20_anqpvenuename to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontrollerhotspot20_h2qpconncapability to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontrollerhotspot20_h2qpoperatorname to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontrollerhotspot20_h2qposuprovider to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontrollerhotspot20_h2qpwanmetric to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontrollerhotspot20_hsprofile to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontrollerhotspot20_icon to support FortiOS 6.0 6.2 6.4 6.6.
* Update resource_wirelesscontrollerhotspot20_qosmap to support FortiOS 6.0 6.2 6.4 6.6.
* Support vdom_mode for resource fortios_system_global.

## 1.8.0 (Jan 25, 2021)

IMPROVEMENTS:

* Support dynamic_sort_subtable for for_each + toset #132
* Filter doc for datasource list

BUG FIXES:

* Clear fortios_vpnssl_settings sub tables when terraform destroy #133

FEATURES:

* **New Data Source:** `fortios_system_interfacelist`
* **New Data Source:** `fortios_system_interface`
* **New Data Source:** `fortios_system_global`
* **New Data Source:** `fortios_routerbgp_neighborlist`
* **New Data Source:** `fortios_routerbgp_neighbor`
* **New Data Source:** `fortios_router_staticlist`
* **New Data Source:** `fortios_router_static6list`
* **New Data Source:** `fortios_router_static6`
* **New Data Source:** `fortios_router_static`
* **New Data Source:** `fortios_router_setting`
* **New Data Source:** `fortios_router_routemaplist`
* **New Data Source:** `fortios_router_routemap`
* **New Data Source:** `fortios_router_ripng`
* **New Data Source:** `fortios_router_rip`
* **New Data Source:** `fortios_router_prefixlistlist`
* **New Data Source:** `fortios_router_prefixlist6list`
* **New Data Source:** `fortios_router_prefixlist6`
* **New Data Source:** `fortios_router_prefixlist`
* **New Data Source:** `fortios_router_policylist`
* **New Data Source:** `fortios_router_policy6list`
* **New Data Source:** `fortios_router_policy6`
* **New Data Source:** `fortios_router_policy`
* **New Data Source:** `fortios_router_ospf6`
* **New Data Source:** `fortios_router_ospf`
* **New Data Source:** `fortios_router_multicastflowlist`
* **New Data Source:** `fortios_router_multicastflow`
* **New Data Source:** `fortios_router_multicast6`
* **New Data Source:** `fortios_router_multicast`
* **New Data Source:** `fortios_router_keychainlist`
* **New Data Source:** `fortios_router_keychain`
* **New Data Source:** `fortios_router_isis`
* **New Data Source:** `fortios_router_communitylistlist`
* **New Data Source:** `fortios_router_communitylist`
* **New Data Source:** `fortios_router_bgp`
* **New Data Source:** `fortios_router_bfd6`
* **New Data Source:** `fortios_router_bfd`
* **New Data Source:** `fortios_router_authpathlist`
* **New Data Source:** `fortios_router_authpath`
* **New Data Source:** `fortios_router_aspathlistlist`
* **New Data Source:** `fortios_router_aspathlist`
* **New Data Source:** `fortios_router_accesslistlist`
* **New Data Source:** `fortios_router_accesslist6list`
* **New Data Source:** `fortios_router_accesslist6`
* **New Data Source:** `fortios_router_accesslist`
* **New Data Source:** `fortios_firewall_policylist`
* **New Data Source:** `fortios_firewall_policy6list`
* **New Data Source:** `fortios_firewall_policy64list`
* **New Data Source:** `fortios_firewall_policy64`
* **New Data Source:** `fortios_firewall_policy6`
* **New Data Source:** `fortios_firewall_policy`
* **New Data Source:** `fortios_firewall_addrgrplist`
* **New Data Source:** `fortios_firewall_addrgrp6list`
* **New Data Source:** `fortios_firewall_addrgrp6`
* **New Data Source:** `fortios_firewall_addrgrp`
* **New Data Source:** `fortios_firewall_addresslist`
* **New Data Source:** `fortios_firewall_address6list`
* **New Data Source:** `fortios_firewall_address6`
* **New Data Source:** `fortios_firewall_address`

## 1.7.0 (Jan 8, 2021)

IMPROVEMENTS:

* Support internet_service_name for fortios_firewall_policy #128
* Support when resource_router_ripng is destroyed, clean its child table
* Support when resource_router_rip is destroyed, clean its child table
* Support when resource_router_multicast(6) is destroyed, clean its child table
* Support when resource_router_isis is destroyed, clean its child table
* Support when resource_router_bfd(6) is destroyed, clean its child table
* Support when resource_router_ospf/resource_router_ospf6 are destroyed, clean its child table #129
* Support when resource_router_bgp is destroyed, clean its child table #127

BUG FIXES:

* Improve documentation of fortios_firewall_security_policysort
* Update fortios_firewall_security_policysort.html.markdown to fix the alignment problem of the example format

FEATURES:

* **New Data Source:** `fortios_json_generic_api`
* **New Resource:** `fortios_routerbgp_neighbor`

## 1.6.18 (Dec 23, 2020)

IMPROVEMENTS:

* Support sort security policies by policy name
* Update documentation for fortios_system_autoscript
* Support radius server for fortios_fmg_system_admin_user #125

## 1.6.17 (Dec 16, 2020)

BUG FIXES:

* Remove ForceNew of firewall objects that don't include policy objects

## 1.6.16 (Dec 6, 2020)

IMPROVEMENTS:

* Support anti-replay, geoip-anycast, emailfilter-profile, cifs-profile, auto-asic-offload, fsso-groups, email-collect, match-vip-only for firewall policy

## 1.6.15 (Dec 2, 2020)

IMPROVEMENTS:

* Improve fortios_system_dnsdatabase.html.markdown #119
* Improve fortios_certificate_local.html.markdown #118

BUG FIXES:

* Fix terraform omit 0 value problem
* Fix fortios_vpnssl_settings auth_timeout cannot be set to 0 #120
* Make mkey of all table resources as ForceNew and Required #105
* Fix unable to change interface name on VpnIpsecPhase1Interface #105

## 1.6.13 (Nov 27, 2020)

BUG FIXES:

* Resync

## 1.6.12 (Nov 27, 2020)

IMPROVEMENTS:

* Improve/fix documentation for fortios_router_accesslist, fortios_certificate_local, fortios_system_fortimanager and fortios_system_ha

## 1.6.11 (Nov 26, 2020)

IMPROVEMENTS:

* Improve documentation for block #116

BUG FIXES:

* Support range for firewall policy

## 1.6.10 (Nov 19, 2020)

BUG FIXES:

* Resync

## 1.6.9 (Nov 19, 2020)

IMPROVEMENTS:

* Improve the documentation of fortios_system_interface

BUG FIXES:

* Fix systemdhcp_server lease_time can not be 0 #112

FEATURES:

* **New Resource:** `fortios_extendercontroller_extender1` (for FortiOS Version >= 6.4.2)
* **New Resource:** `fortios_user_saml`

## 1.6.8 (Oct 23, 2020)

IMPROVEMENTS:

* Add documentation for sorting security policies with terraform depends_on or fortios_firewall_security_policysort

FEATURES:

* **New Resource:** `fortios_firewall_security_policysort`

## 1.6.7 (Oct 21, 2020)

IMPROVEMENTS:

* Support the configuration of the interfaces automatically created by FGT #97

BUG FIXES:

* Make fortios_vpnipsec_phase1interface net_device optional #101

## 1.6.6 (Oct 16, 2020)

IMPROVEMENTS:

* Improve documentation of 'Generate an API token for FortiOS'

BUG FIXES:

* Support package for fortios_fmg_devicemanager_script_execute #95

## 1.6.5 (Oct 15, 2020)

BUG FIXES:

* Fix the errors in the document descriptions of 'To move a policy' and 'fortios_firewall_security_policyseq'

## 1.6.4 (Oct 14, 2020)

IMPROVEMENTS:

* Improve the documentation description for the resources that will be deprecated

FEATURES:

* Add documentation for moving a policy

BUG FIXES:

* Make fortios_router_static device optional #93

## 1.6.3 (Oct 8, 2020)

IMPROVEMENTS:

* Improve FMG jsonrpc resource

FEATURES:

* Support inspection_mode, http_policy_redirect, ssh_policy_redirect, webproxy_profile for fortios_firewall_policy #92

## 1.6.2 (Sep 24, 2020)

IMPROVEMENTS:

* Mark password and key arguments as sensitive #86

FEATURES:

* **New Resource:** `fortios_system_ssoadmin`

## 1.6.1 (Sep 16, 2020)

IMPROVEMENTS:

* Support detecting changes to firewall policy sequence performed outside of terraform for fortios_firewall_security_policyseq #77
* Improve the optional/required option of parameteres for all resources

BUG FIXES:

* Make fortios_system_admin password optional #83
* Fix optional arguments problem for fortios_system_sdnconnector #85 and fortios_firewall_vip #84

FEATURES:

* Support ssl_max_proto_ver, ssl_min_proto_ver for fortios_vpnssl_settings #82

## 1.6.0 (Sep 7, 2020)

FEATURES:

* **New Resource:** `fortios_application_custom`
* **New Resource:** `fortios_application_name`
* **New Resource:** `fortios_certificate_ca`
* **New Resource:** `fortios_certificate_crl`
* **New Resource:** `fortios_certificate_local`
* **New Resource:** `fortios_endpointcontrol_client`
* **New Resource:** `fortios_endpointcontrol_forticlientems`
* **New Resource:** `fortios_endpointcontrol_registeredforticlient`
* **New Resource:** `fortios_firewall_DoSpolicy`
* **New Resource:** `fortios_firewall_DoSpolicy6`
* **New Resource:** `fortios_firewall_internetservice`
* **New Resource:** `fortios_firewall_internetservicecustom`
* **New Resource:** `fortios_firewall_internetservicecustomgroup`
* **New Resource:** `fortios_firewall_internetservicedefinition`
* **New Resource:** `fortios_firewall_internetservicegroup`
* **New Resource:** `fortios_firewall_ipv6ehfilter`
* **New Resource:** `fortios_firewall_policy46`
* **New Resource:** `fortios_firewall_profileprotocoloptions`
* **New Resource:** `fortios_firewall_proxyaddrgrp`
* **New Resource:** `fortios_firewallconsolidated_policy`
* **New Resource:** `fortios_ips_custom`
* **New Resource:** `fortios_ips_decoder`
* **New Resource:** `fortios_ips_rule`
* **New Resource:** `fortios_ips_rulesettings`
* **New Resource:** `fortios_ips_sensor`
* **New Resource:** `fortios_logsyslogd2_overridesetting`
* **New Resource:** `fortios_logsyslogd3_overridesetting`
* **New Resource:** `fortios_logsyslogd4_overridesetting`
* **New Resource:** `fortios_logsyslogd_overridesetting`
* **New Resource:** `fortios_spamfilter_iptrust`
* **New Resource:** `fortios_switchcontroller_global`
* **New Resource:** `fortios_switchcontroller_macsyncsettings`
* **New Resource:** `fortios_switchcontroller_managedswitch`
* **New Resource:** `fortios_switchcontroller_quarantine`
* **New Resource:** `fortios_switchcontroller_stormcontrol`
* **New Resource:** `fortios_switchcontroller_switchgroup`
* **New Resource:** `fortios_switchcontroller_system`
* **New Resource:** `fortios_switchcontroller_vlan`
* **New Resource:** `fortios_switchcontrollerautoconfig_custom`
* **New Resource:** `fortios_switchcontrollerautoconfig_default`
* **New Resource:** `fortios_switchcontrollerautoconfig_policy`
* **New Resource:** `fortios_switchcontrollerqos_qospolicy`
* **New Resource:** `fortios_switchcontrollersecuritypolicy_captiveportal`
* **New Resource:** `fortios_system_affinityinterrupt`
* **New Resource:** `fortios_system_affinitypacketredistribution`
* **New Resource:** `fortios_system_alarm`
* **New Resource:** `fortios_system_automationstitch`
* **New Resource:** `fortios_system_customlanguage`
* **New Resource:** `fortios_system_dedicatedmgmt`
* **New Resource:** `fortios_system_geneve`
* **New Resource:** `fortios_system_geoipoverride`
* **New Resource:** `fortios_system_macaddresstable`
* **New Resource:** `fortios_system_ptp`
* **New Resource:** `fortios_system_saml`
* **New Resource:** `fortios_system_storage`
* **New Resource:** `fortios_system_vdomdns`
* **New Resource:** `fortios_system_vdomproperty`
* **New Resource:** `fortios_system_vdomradiusserver`
* **New Resource:** `fortios_systemdhcp6_server`
* **New Resource:** `fortios_systemreplacemsg_admin`
* **New Resource:** `fortios_systemreplacemsg_alertmail`
* **New Resource:** `fortios_systemreplacemsg_auth`
* **New Resource:** `fortios_systemreplacemsg_devicedetectionportal`
* **New Resource:** `fortios_systemreplacemsg_ec`
* **New Resource:** `fortios_systemreplacemsg_fortiguardwf`
* **New Resource:** `fortios_systemreplacemsg_ftp`
* **New Resource:** `fortios_systemreplacemsg_http`
* **New Resource:** `fortios_systemreplacemsg_icap`
* **New Resource:** `fortios_systemreplacemsg_mail`
* **New Resource:** `fortios_systemreplacemsg_nacquar`
* **New Resource:** `fortios_systemreplacemsg_nntp`
* **New Resource:** `fortios_systemreplacemsg_spam`
* **New Resource:** `fortios_systemreplacemsg_sslvpn`
* **New Resource:** `fortios_systemreplacemsg_trafficquota`
* **New Resource:** `fortios_systemreplacemsg_utm`
* **New Resource:** `fortios_systemreplacemsg_webproxy`
* **New Resource:** `fortios_user_devicecategory`
* **New Resource:** `fortios_user_fortitoken`
* **New Resource:** `fortios_user_fssopolling`
* **New Resource:** `fortios_user_securityexemptlist`
* **New Resource:** `fortios_vpn_l2tp`
* **New Resource:** `fortios_vpn_pptp`
* **New Resource:** `fortios_vpncertificate_ca`
* **New Resource:** `fortios_vpncertificate_crl`
* **New Resource:** `fortios_vpncertificate_local`
* **New Resource:** `fortios_vpncertificate_remote`
* **New Resource:** `fortios_waf_mainclass`
* **New Resource:** `fortios_waf_signature`
* **New Resource:** `fortios_waf_subclass`
* **New Resource:** `fortios_wanopt_profile`
* **New Resource:** `fortios_webproxy_explicit`
* **New Resource:** `fortios_wirelesscontroller_apstatus`
* **New Resource:** `fortios_wirelesscontroller_bleprofile`
* **New Resource:** `fortios_wirelesscontroller_bonjourprofile`
* **New Resource:** `fortios_wirelesscontroller_qosprofile`
* **New Resource:** `fortios_wirelesscontroller_region`
* **New Resource:** `fortios_wirelesscontroller_setting`
* **New Resource:** `fortios_wirelesscontroller_timers`
* **New Resource:** `fortios_wirelesscontroller_utmprofile`
* **New Resource:** `fortios_wirelesscontroller_vap`
* **New Resource:** `fortios_wirelesscontroller_vapgroup`
* **New Resource:** `fortios_wirelesscontroller_widsprofile`
* **New Resource:** `fortios_wirelesscontroller_wtp`
* **New Resource:** `fortios_wirelesscontroller_wtpgroup`
* **New Resource:** `fortios_wirelesscontroller_wtpprofile`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_anqpnairealm`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_h2qposuprovider`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_hsprofile`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_qosmap`

## 1.5.3 (Sep 2, 2020)

BUG FIXES:

* Fix fortios_vpnipsec_phase1interfaceas bugs

## 1.5.2 (Sep 2, 2020)

IMPROVEMENTS:

* Improve doc category
* Improve doc `To change the admin default port`

BUG FIXES:

* Remove the required attribute of remote_gw in fortios_vpnipsec_phase1interfaceas #75

FEATURES:

* **New Resource:** `fortios_system_settings`
* **New Resource:** `fortios_application_rulesettings`
* **New Resource:** `fortios_firewall_sslsshprofile`
* **New Resource:** `fortios_firewallssh_localca`
* **New Resource:** `fortios_firewallssh_localkey`
* **New Resource:** `fortios_router_aspathlist`
* **New Resource:** `fortios_router_bfd`
* **New Resource:** `fortios_router_bfd6`
* **New Resource:** `fortios_system_switchinterface`
* **New Resource:** `fortios_system_vdom`
* **New Resource:** `fortios_system_vdomlink`
* **New Resource:** `system_virtualwirepair`
* **New Resource:** `user_devicegroup`

## 1.5.1 (Aug 25, 2020)

IMPROVEMENTS:

* Improve doc category

## 1.5.0 (Aug 25, 2020)

FEATURES:

* **New Resource:** `fortios_wirelesscontrollerhotspot20_icon`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_h2qpwanmetric`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_h2qpoperatorname`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_h2qpconncapability`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_anqpvenuename`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_anqproamingconsortium`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_anqpipaddresstype`
* **New Resource:** `fortios_wirelesscontrollerhotspot20_anqp3gppcellular`
* **New Resource:** `fortios_wirelesscontroller_intercontroller`
* **New Resource:** `fortios_wirelesscontroller_global`
* **New Resource:** `fortios_webproxy_wisp`
* **New Resource:** `fortios_webproxy_urlmatch`
* **New Resource:** `fortios_webproxy_profile`
* **New Resource:** `fortios_webproxy_global`
* **New Resource:** `fortios_webproxy_forwardservergroup`
* **New Resource:** `fortios_webproxy_forwardserver`
* **New Resource:** `fortios_webproxy_debugurl`
* **New Resource:** `fortios_webfilter_urlfilter`
* **New Resource:** `fortios_webfilter_searchengine`
* **New Resource:** `fortios_webfilter_profile`
* **New Resource:** `fortios_webfilter_override`
* **New Resource:** `fortios_webfilter_ipsurlfiltersetting6`
* **New Resource:** `fortios_webfilter_ipsurlfiltersetting`
* **New Resource:** `fortios_webfilter_ipsurlfiltercachesetting`
* **New Resource:** `fortios_webfilter_ftgdlocalrating`
* **New Resource:** `fortios_webfilter_ftgdlocalcat`
* **New Resource:** `fortios_webfilter_fortiguard`
* **New Resource:** `fortios_webfilter_contentheader`
* **New Resource:** `fortios_webfilter_content`
* **New Resource:** `fortios_wanopt_webcache`
* **New Resource:** `fortios_wanopt_settings`
* **New Resource:** `fortios_wanopt_remotestorage`
* **New Resource:** `fortios_wanopt_peer`
* **New Resource:** `fortios_wanopt_contentdeliverynetworkrule`
* **New Resource:** `fortios_wanopt_cacheservice`
* **New Resource:** `fortios_wanopt_authgroup`
* **New Resource:** `fortios_waf_profile`
* **New Resource:** `fortios_vpnsslweb_usergroupbookmark`
* **New Resource:** `fortios_vpnsslweb_userbookmark`
* **New Resource:** `fortios_vpnsslweb_realm`
* **New Resource:** `fortios_vpnsslweb_portal`
* **New Resource:** `fortios_vpnsslweb_hostchecksoftware`
* **New Resource:** `fortios_vpnssl_settings`
* **New Resource:** `fortios_vpnipsec_phase2interface`
* **New Resource:** `fortios_vpnipsec_phase2`
* **New Resource:** `fortios_vpnipsec_phase1interface`
* **New Resource:** `fortios_vpnipsec_phase1`
* **New Resource:** `fortios_vpnipsec_manualkeyinterface`
* **New Resource:** `fortios_vpnipsec_manualkey`
* **New Resource:** `fortios_vpnipsec_forticlient`
* **New Resource:** `fortios_vpnipsec_concentrator`
* **New Resource:** `fortios_vpncertificate_setting`
* **New Resource:** `fortios_vpncertificate_ocspserver`
* **New Resource:** `fortios_voip_profile`
* **New Resource:** `fortios_user_tacacs`
* **New Resource:** `fortios_user_setting`
* **New Resource:** `fortios_user_radius`
* **New Resource:** `fortios_user_quarantine`
* **New Resource:** `fortios_user_pop3`
* **New Resource:** `fortios_user_peergrp`
* **New Resource:** `fortios_user_peer`
* **New Resource:** `fortios_user_passwordpolicy`
* **New Resource:** `fortios_user_local`
* **New Resource:** `fortios_user_ldap`
* **New Resource:** `fortios_user_krbkeytab`
* **New Resource:** `fortios_user_group`
* **New Resource:** `fortios_user_fsso`
* **New Resource:** `fortios_user_domaincontroller`
* **New Resource:** `fortios_user_deviceaccesslist`
* **New Resource:** `fortios_user_device`
* **New Resource:** `fortios_user_adgrp`
* **New Resource:** `fortios_systemsnmp_user`
* **New Resource:** `fortios_systemsnmp_sysinfo`
* **New Resource:** `fortios_systemsnmp_community`
* **New Resource:** `fortios_systemlldp_networkpolicy`
* **New Resource:** `fortios_systemdhcp_server`
* **New Resource:** `fortios_systemautoupdate_tunneling`
* **New Resource:** `fortios_systemautoupdate_schedule`
* **New Resource:** `fortios_systemautoupdate_pushupdate`
* **New Resource:** `fortios_system_zone`
* **New Resource:** `fortios_system_wccp`
* **New Resource:** `fortios_system_vxlan`
* **New Resource:** `fortios_system_virtualwanlink`
* **New Resource:** `fortios_system_vdomsflow`
* **New Resource:** `fortios_system_vdomnetflow`
* **New Resource:** `fortios_system_vdomexception`
* **New Resource:** `fortios_system_tosbasedpriority`
* **New Resource:** `fortios_system_smsserver`
* **New Resource:** `fortios_system_sittunnel`
* **New Resource:** `fortios_system_sflow`
* **New Resource:** `fortios_system_sessionttl`
* **New Resource:** `fortios_system_sessionhelper`
* **New Resource:** `fortios_system_sdnconnector`
* **New Resource:** `fortios_system_resourcelimits`
* **New Resource:** `fortios_system_replacemsgimage`
* **New Resource:** `fortios_system_replacemsggroup`
* **New Resource:** `fortios_system_proxyarp`
* **New Resource:** `fortios_system_proberesponse`
* **New Resource:** `fortios_system_pppoeinterface`
* **New Resource:** `fortios_system_passwordpolicyguestadmin`
* **New Resource:** `fortios_system_passwordpolicy`
* **New Resource:** `fortios_system_objecttagging`
* **New Resource:** `fortios_system_ntp`
* **New Resource:** `fortios_system_networkvisibility`
* **New Resource:** `fortios_system_netflow`
* **New Resource:** `fortios_system_ndproxy`
* **New Resource:** `fortios_system_nat64`
* **New Resource:** `fortios_system_mobiletunnel`
* **New Resource:** `fortios_system_managementtunnel`
* **New Resource:** `fortios_system_linkmonitor`
* **New Resource:** `fortios_system_ipv6tunnel`
* **New Resource:** `fortios_system_ipv6neighborcache`
* **New Resource:** `fortios_system_ipsecaggregate`
* **New Resource:** `fortios_system_ipiptunnel`
* **New Resource:** `fortios_system_interface`
* **New Resource:** `fortios_system_hamonitor`
* **New Resource:** `fortios_system_ha`
* **New Resource:** `fortios_system_gretunnel`
* **New Resource:** `fortios_system_global`
* **New Resource:** `fortios_system_ftmpush`
* **New Resource:** `fortios_system_fssopolling`
* **New Resource:** `fortios_system_fortisandbox`
* **New Resource:** `fortios_system_fortimanager`
* **New Resource:** `fortios_system_fortiguard`
* **New Resource:** `fortios_system_fm`
* **New Resource:** `fortios_system_fipscc`
* **New Resource:** `fortios_system_externalresource`
* **New Resource:** `fortios_system_emailserver`
* **New Resource:** `fortios_system_dscpbasedpriority`
* **New Resource:** `fortios_system_dnsserver`
* **New Resource:** `fortios_system_dnsdatabase`
* **New Resource:** `fortios_system_dns`
* **New Resource:** `fortios_system_ddns`
* **New Resource:** `fortios_system_csf`
* **New Resource:** `fortios_system_console`
* **New Resource:** `fortios_system_clustersync`
* **New Resource:** `fortios_system_centralmanagement`
* **New Resource:** `fortios_system_autoscript`
* **New Resource:** `fortios_system_automationtrigger`
* **New Resource:** `fortios_system_automationdestination`
* **New Resource:** `fortios_system_automationaction`
* **New Resource:** `fortios_system_autoinstall`
* **New Resource:** `fortios_system_arptable`
* **New Resource:** `fortios_system_apiuser`
* **New Resource:** `fortios_system_alias`
* **New Resource:** `fortios_system_admin`
* **New Resource:** `fortios_system_accprofile`
* **New Resource:** `fortios_switchcontrollersecuritypolicy_8021X`
* **New Resource:** `fortios_switchcontrollerqos_queuepolicy`
* **New Resource:** `fortios_switchcontrollerqos_ipdscpmap`
* **New Resource:** `fortios_switchcontrollerqos_dot1pmap`
* **New Resource:** `fortios_switchcontroller_virtualportpool`
* **New Resource:** `fortios_switchcontroller_trafficpolicy`
* **New Resource:** `fortios_switchcontroller_switchprofile`
* **New Resource:** `fortios_switchcontroller_switchlog`
* **New Resource:** `fortios_switchcontroller_switchinterfacetag`
* **New Resource:** `fortios_switchcontroller_stpsettings`
* **New Resource:** `fortios_switchcontroller_sflow`
* **New Resource:** `fortios_switchcontroller_networkmonitorsettings`
* **New Resource:** `fortios_switchcontroller_lldpsettings`
* **New Resource:** `fortios_switchcontroller_lldpprofile`
* **New Resource:** `fortios_switchcontroller_igmpsnooping`
* **New Resource:** `fortios_switchcontroller_customcommand`
* **New Resource:** `fortios_switchcontroller_8021Xsettings`
* **New Resource:** `fortios_sshfilter_profile`
* **New Resource:** `fortios_spamfilter_profile`
* **New Resource:** `fortios_spamfilter_options`
* **New Resource:** `fortios_spamfilter_mheader`
* **New Resource:** `fortios_spamfilter_fortishield`
* **New Resource:** `fortios_spamfilter_dnsbl`
* **New Resource:** `fortios_spamfilter_bword`
* **New Resource:** `fortios_spamfilter_bwl`
* **New Resource:** `fortios_router_static6`
* **New Resource:** `fortios_router_static`
* **New Resource:** `fortios_router_setting`
* **New Resource:** `fortios_router_routemap`
* **New Resource:** `fortios_router_ripng`
* **New Resource:** `fortios_router_rip`
* **New Resource:** `fortios_router_prefixlist6`
* **New Resource:** `fortios_router_prefixlist`
* **New Resource:** `fortios_router_policy6`
* **New Resource:** `fortios_router_policy`
* **New Resource:** `fortios_router_ospf6`
* **New Resource:** `fortios_router_ospf`
* **New Resource:** `fortios_router_multicastflow`
* **New Resource:** `fortios_router_multicast6`
* **New Resource:** `fortios_router_multicast`
* **New Resource:** `fortios_router_keychain`
* **New Resource:** `fortios_router_isis`
* **New Resource:** `fortios_router_communitylist`
* **New Resource:** `fortios_router_bgp`
* **New Resource:** `fortios_router_authpath`
* **New Resource:** `fortios_router_accesslist6`
* **New Resource:** `fortios_router_accesslist`
* **New Resource:** `fortios_report_theme`
* **New Resource:** `fortios_report_style`
* **New Resource:** `fortios_report_setting`
* **New Resource:** `fortios_report_layout`
* **New Resource:** `fortios_report_dataset`
* **New Resource:** `fortios_report_chart`
* **New Resource:** `fortios_logwebtrends_setting`
* **New Resource:** `fortios_logwebtrends_filter`
* **New Resource:** `fortios_logsyslogd4_setting`
* **New Resource:** `fortios_logsyslogd4_overridefilter`
* **New Resource:** `fortios_logsyslogd4_filter`
* **New Resource:** `fortios_logsyslogd3_setting`
* **New Resource:** `fortios_logsyslogd3_overridefilter`
* **New Resource:** `fortios_logsyslogd3_filter`
* **New Resource:** `fortios_logsyslogd2_setting`
* **New Resource:** `fortios_logsyslogd2_overridefilter`
* **New Resource:** `fortios_logsyslogd2_filter`
* **New Resource:** `fortios_logsyslogd_setting`
* **New Resource:** `fortios_logsyslogd_overridefilter`
* **New Resource:** `fortios_logsyslogd_filter`
* **New Resource:** `fortios_lognulldevice_setting`
* **New Resource:** `fortios_lognulldevice_filter`
* **New Resource:** `fortios_logmemory_setting`
* **New Resource:** `fortios_logmemory_globalsetting`
* **New Resource:** `fortios_logmemory_filter`
* **New Resource:** `fortios_logfortiguard_setting`
* **New Resource:** `fortios_logfortiguard_overridesetting`
* **New Resource:** `fortios_logfortiguard_overridefilter`
* **New Resource:** `fortios_logfortiguard_filter`
* **New Resource:** `fortios_logfortianalyzer3_setting`
* **New Resource:** `fortios_logfortianalyzer3_overridesetting`
* **New Resource:** `fortios_logfortianalyzer3_overridefilter`
* **New Resource:** `fortios_logfortianalyzer3_filter`
* **New Resource:** `fortios_logfortianalyzer2_setting`
* **New Resource:** `fortios_logfortianalyzer2_overridesetting`
* **New Resource:** `fortios_logfortianalyzer2_overridefilter`
* **New Resource:** `fortios_logfortianalyzer2_filter`
* **New Resource:** `fortios_logfortianalyzer_setting`
* **New Resource:** `fortios_logfortianalyzer_overridesetting`
* **New Resource:** `fortios_logfortianalyzer_overridefilter`
* **New Resource:** `fortios_logfortianalyzer_filter`
* **New Resource:** `fortios_logdisk_setting`
* **New Resource:** `fortios_logdisk_filter`
* **New Resource:** `fortios_log_threatweight`
* **New Resource:** `fortios_log_setting`
* **New Resource:** `fortios_log_guidisplay`
* **New Resource:** `fortios_log_eventfilter`
* **New Resource:** `fortios_log_customfield`
* **New Resource:** `fortios_ips_settings`
* **New Resource:** `fortios_ips_global`
* **New Resource:** `fortios_icap_server`
* **New Resource:** `fortios_icap_profile`
* **New Resource:** `fortios_ftpproxy_explicit`
* **New Resource:** `fortios_firewallwildcardfqdn_group`
* **New Resource:** `fortios_firewallwildcardfqdn_custom`
* **New Resource:** `fortios_firewallssl_setting`
* **New Resource:** `fortios_firewallssh_setting`
* **New Resource:** `fortios_firewallssh_hostkey`
* **New Resource:** `fortios_firewallshaper_trafficshaper`
* **New Resource:** `fortios_firewallshaper_peripshaper`
* **New Resource:** `fortios_firewallservice_group`
* **New Resource:** `fortios_firewallservice_custom`
* **New Resource:** `fortios_firewallservice_category`
* **New Resource:** `fortios_firewallschedule_recurring`
* **New Resource:** `fortios_firewallschedule_onetime`
* **New Resource:** `fortios_firewallschedule_group`
* **New Resource:** `fortios_firewallipmacbinding_table`
* **New Resource:** `fortios_firewallipmacbinding_setting`
* **New Resource:** `fortios_firewall_vipgrp64`
* **New Resource:** `fortios_firewall_vipgrp46`
* **New Resource:** `fortios_firewall_vipgrp6`
* **New Resource:** `fortios_firewall_vipgrp`
* **New Resource:** `fortios_firewall_vip64`
* **New Resource:** `fortios_firewall_vip46`
* **New Resource:** `fortios_firewall_vip6`
* **New Resource:** `fortios_firewall_vip`
* **New Resource:** `fortios_firewall_ttlpolicy`
* **New Resource:** `fortios_firewall_sslserver`
* **New Resource:** `fortios_firewall_sniffer`
* **New Resource:** `fortios_firewall_shapingprofile`
* **New Resource:** `fortios_firewall_shapingpolicy`
* **New Resource:** `fortios_firewall_proxypolicy`
* **New Resource:** `fortios_firewall_proxyaddress`
* **New Resource:** `fortios_firewall_profilegroup`
* **New Resource:** `fortios_firewall_policy64`
* **New Resource:** `fortios_firewall_policy6`
* **New Resource:** `fortios_firewall_policy`
* **New Resource:** `fortios_firewall_multicastpolicy6`
* **New Resource:** `fortios_firewall_multicastpolicy`
* **New Resource:** `fortios_firewall_multicastaddress6`
* **New Resource:** `fortios_firewall_multicastaddress`
* **New Resource:** `fortios_firewall_localinpolicy6`
* **New Resource:** `fortios_firewall_localinpolicy`
* **New Resource:** `fortios_firewall_ldbmonitor`
* **New Resource:** `fortios_firewall_iptranslation`
* **New Resource:** `fortios_firewall_ippool6`
* **New Resource:** `fortios_firewall_ippool`
* **New Resource:** `fortios_firewall_internetserviceextension`
* **New Resource:** `fortios_firewall_interfacepolicy6`
* **New Resource:** `fortios_firewall_interfacepolicy`
* **New Resource:** `fortios_firewall_identitybasedroute`
* **New Resource:** `fortios_firewall_dnstranslation`
* **New Resource:** `fortios_firewall_centralsnatmap`
* **New Resource:** `fortios_firewall_authportal`
* **New Resource:** `fortios_firewall_addrgrp6`
* **New Resource:** `fortios_firewall_addrgrp`
* **New Resource:** `fortios_firewall_address6template`
* **New Resource:** `fortios_firewall_address6`
* **New Resource:** `fortios_firewall_address`
* **New Resource:** `fortios_extendercontroller_extender`
* **New Resource:** `fortios_endpointcontrol_settings`
* **New Resource:** `fortios_endpointcontrol_profile`
* **New Resource:** `fortios_endpointcontrol_forticlientregistrationsync`
* **New Resource:** `fortios_dnsfilter_profile`
* **New Resource:** `fortios_dnsfilter_domainfilter`
* **New Resource:** `fortios_dlp_settings`
* **New Resource:** `fortios_dlp_sensor`
* **New Resource:** `fortios_dlp_fpsensitivity`
* **New Resource:** `fortios_dlp_fpdocsource`
* **New Resource:** `fortios_dlp_filepattern`
* **New Resource:** `fortios_authentication_setting`
* **New Resource:** `fortios_authentication_scheme`
* **New Resource:** `fortios_authentication_rule`
* **New Resource:** `fortios_application_list`
* **New Resource:** `fortios_application_group`
* **New Resource:** `fortios_antivirus_settings`
* **New Resource:** `fortios_antivirus_quarantine`
* **New Resource:** `fortios_antivirus_profile`
* **New Resource:** `fortios_antivirus_heuristic`
* **New Resource:** `fortios_alertemail_setting`

## 1.4.1 (July 31, 2020)

IMPROVEMENTS:

* Support doc subcategories in Registry
* Support new doc naming rules in Registry

BUG FIXES:

* Fix SendWithSpecialParams comment error

## 1.4.0 (July 30, 2020)

IMPROVEMENTS:

* Update for Terraform 0.13 and Terraform Registry

## 1.3.0 (June 15, 2020)

BUG FIXES:

* resource/object_service: Creating fortios_firewall_object_service causes Terraform panic crash - Support firewall.service/custom's session-ttl type change from v6.2.3 ([#67](https://github.com/terraform-providers/terraform-provider-fortios/issues/67))


## 1.2.0 (May 29, 2020)

IMPROVEMENTS:

* Support Terraform Plugin SDK ([#50](https://github.com/terraform-providers/terraform-provider-fortios/pull/50))

FEATURES:

* **New Resource:** `fortios_fmg_devicemanager_device`
* **New Resource:** `fortios_fmg_devicemanager_installdevice`
* **New Resource:** `fortios_fmg_devicemanager_installpackage`
* **New Resource:** `fortios_fmg_devicemanager_script_execute`
* **New Resource:** `fortios_fmg_devicemanager_script`
* **New Resource:** `fortios_fmg_firewall_object_address`
* **New Resource:** `fortios_fmg_firewall_object_ippool`
* **New Resource:** `fortios_fmg_firewall_object_service`
* **New Resource:** `fortios_fmg_firewall_object_vip`
* **New Resource:** `fortios_fmg_firewall_security_policy`
* **New Resource:** `fortios_fmg_firewall_security_policypackage`
* **New Resource:** `fortios_fmg_object_adom_revision`
* **New Resource:** `fortios_fmg_system_admin`
* **New Resource:** `fortios_fmg_system_admin_profiles`
* **New Resource:** `fortios_fmg_system_admin_user`
* **New Resource:** `fortios_fmg_system_adom`
* **New Resource:** `fortios_fmg_system_dns`
* **New Resource:** `fortios_fmg_system_global`
* **New Resource:** `fortios_fmg_system_license_forticare`
* **New Resource:** `fortios_fmg_system_license_vm`
* **New Resource:** `fortios_fmg_system_network_interface`
* **New Resource:** `fortios_fmg_system_network_route`
* **New Resource:** `fortios_fmg_system_ntp`
* **New Resource:** `fortios_fmg_system_syslogserver`
* **New Resource:** `fortios_firewall_security_policyseq`
* **New Resource:** `fortios_firewall_object_servicecategory`


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
