## 1.7.0 (Unreleased)

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
