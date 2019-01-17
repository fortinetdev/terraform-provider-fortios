provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}


# Attention: system will reboot here if successful
resource "fortios_system_license_vm" "test2" {
	file_content = "LS0tLS1CRUdJTiBGR1QgVk0gTElDRU5TRS0tLS0tDQpRQUFBQUxXaTdCVnVkV2x3QXJZcC92S2J2Yk5zME5YNWluUW9sVldmcFoxWldJQi9pL2g4c01oR0psWWc5Vkl1DQorSlBJRis1aFphMWwyNm9yNHdiOW93RnJDeVZnQUFBQWhxWjliWHFLK1hGN2o3dnB3WTB6QXRTaTdOMVM1ZWNxDQpWYmRRREZyYklUdnRvUWNyRU1jV0ltQzFqWWs5dmVoeGlYTG1OV0MwN25BSitYTTJFNmh2b29DMjE1YUwxK2wrDQovUHl5M0VLVnNTNjJDT2hMZHc3UndXajB3V3RqMmZiWg0KLS0tLS1FTkQgRkdUIFZNIExJQ0VOU0UtLS0tLQ0K"
}

#######################################################################
#
#resource "fortios_system_license_forticare" "test2" {
#	registration_code = "fdasfdsafdsafdsa"
#}
#
#resource "fortios_system_vdom_setting" "test2" {
#	name = "aa1122"
#	short_name = "aa1122"
#	temporary = 0
#}
#
#resource "fortios_system_apiuser_setting" "test2" {
#	name = "cccc"
#	accprofile = "restAPIprofile"
#	vdom = "root"
#	trusthost = [
#		{
#			type = "ipv4-trusthost"
#			ipv4_trusthost = "61.149.0.0 255.255.0.0"
#		},
#		{
#			type = "ipv4-trusthost"
#			ipv4_trusthost = "22.22.0.0 255.255.0.0"
#		}				
#	]
#}
#
#resource "fortios_log_fortianalyzer_setting" "test1" {
#	status = "enable"
#	server = "10.2.2.99"
#	source_ip = "10.2.2.99"
#	upload_option = "realtime"
#	reliable = "enable"
#	hmac_algorithm = "sha256"
#	enc_algorithm = "high-medium"
#}
#resource "fortios_log_syslog_setting" "test2" {
#	status = "enable"
#	server = "2.2.2.2"
#	mode = "udp"
#	port = "514"
#	facility = "local7"
#	source_ip = "10.2.2.199"
#	format = "csv"
#}
#
#resource "fortios_system_setting_dns" "test1" {
#	primary = "208.91.112.53"
#	secondary = "208.91.112.22"
#}
#
#resource "fortios_system_setting_ntp" "test2" {
#	type = "custom"
#	ntpserver = ["1.1.1.1", "3.3.3.3"]
#}
#
#resource "fortios_system_setting_global" "test3" {
#	admintimeout = 65
#	timezone = "04"
#	hostname = "333333"
#	admin_sport = 443
#	admin_ssh_port = 22
#}
#
#
#resource "fortios_firewall_security_policy" "test2" {
#	name = "ap1"
#	srcintf = ["port2"]
#	dstintf = ["port1"]
#	srcaddr = ["swscan.apple.com", "google-play"]    
#	dstaddr = ["swscan.apple.com", "update.microsoft.com"]
#	internet_service = "enable"
#	internet_service_id = [917520, 6881402, 393219]
#	schedule = "always"
#	service = []
#	action = "accept"	
#	utm_status = "enable"
#	logtraffic = "all"
#	logtraffic_start = "enable"
#	capture_packet = "enable"
#	ippool = "enable"
#	poolname = ["rewq", "rbb"]
#	groups = ["Guest-group", "SSO_Guest_Users"]
#	devices = ["android-phone", "android-tablet"]
#	comments = "fdasfdsa"
#	av_profile = "wifi-default"
#	webfilter_profile = "monitor-all"
#	dnsfilter_profile = "default"
#	ips_sensor = "protect_client"
#	application_list = "block-high-risk"
#	ssl_ssh_profile = "certificate-inspection"
#	nat = "enable"
#}
#
#resource "fortios_firewall_security_policy" "test1" {
#	name = "ap2"
#	srcintf = ["port2"]
#	dstintf = ["port1"]
#	srcaddr = ["swscan.apple.com", "google-play"]    
#	dstaddr = ["swscan.apple.com", "update.microsoft.com"]
#	internet_service = "disable"
#	internet_service_id = []
#	schedule = "always"
#	service = ["ALL_ICMP", "FTP"]
#	action = "accept"	
#	utm_status = "enable"
#	logtraffic = "all"
#	logtraffic_start = "enable"
#	capture_packet = "enable"
#	ippool = "enable"
#	poolname = ["rewq", "rbb"]
#	groups = ["Guest-group", "SSO_Guest_Users"]
#	devices = ["android-phone", "android-tablet"]
#	comments = "fdasfdsa"
#	av_profile = "wifi-default"
#	webfilter_profile = "monitor-all"
#	dnsfilter_profile = "default"
#	ips_sensor = "protect_client"
#	application_list = "block-high-risk"
#	ssl_ssh_profile = "certificate-inspection"
#	nat = "enable"
#}
#
#
#resource "fortios_firewall_object_service" "v11" {
#	name = "a23322"
#	category = "General"
#	protocol = "TCP/UDP/SCTP"
#	fqdn = "fdsafdsa"
#	comment = "fdsafdsafdas11"
#}
#
#resource "fortios_firewall_object_service" "v13" {
#	name = "1fdsafd11a"
#	category = "General"
#	protocol = "TCP/UDP/SCTP"
#	iprange = "1.1.1.1-2.2.2.2"
#	comment = "fdsqqqqeweafdsa"
#}
#
#resource "fortios_firewall_object_servicegroup" "v11" {
#	name = "1fdsafd11a"
#	comment = "fdsafdsa"
#	member = ["DCE-RPC", "DNS", "HTTPS"]
#}
#
#resource "fortios_firewall_object_vipgroup" "v11" {
#	name = "1fdsafd11a"
#	interface = "port3"
#	comments = "fdsafdsa"
#	member = ["44", "3"]
#}
#
#resource "fortios_firewall_object_vip" "v11" {
#	name = "dfa"
#	comment = "fdsafdsafds"
#	extip = "11.1.1.1-21.1.1.1"
#	mappedip = ["22.2.2.2-32.2.2.2"]
#	extintf = "port3"
#	portforward = "disable"
#}
#
#
#resource "fortios_firewall_object_ippool" "s1" {
#	name = "ddd"
#	type = "overload"
#	startip = "11.0.0.0"
#	endip = "22.0.0.0"
#	arp_reply = "enable"
#	comments = "fdsaf"
#}
#
#resource "fortios_firewall_object_ippool" "s2" {
#	name = "dd22d"
#	type = "one-to-one"
#	startip = "121.0.0.0"
#	endip = "222.0.0.0"
#	arp_reply = "enable"
#	comments = "fdsaf"
#}
#
#
#resource "fortios_firewall_object_addressgroup" "s1" {
#	name = "s1"
#	member = ["google-play","swscan.apple.com"]
#	comment = "dfdsad"
#}
#
#resource "fortios_firewall_object_address" "s1" {
#	name = "s1"
#	type = "iprange"
#	start_ip = "1.0.0.0"
#	end_ip = "2.0.0.0"
#	comment = "dd"
#}
#resource "fortios_firewall_object_address" "s2" {
#	name = "s2"
#	type = "geography"
#	country = "AO"
#	comment = "dd"
#}
#
#resource "fortios_firewall_object_address" "s3" {
#	name = "s3"
#	type = "fqdn"
#	fqdn = "baid.com"
#	comment = "dd"
#}
# 
#resource "fortios_firewall_object_address" "s4" {
#	name = "s4"
#	type = "ipmask"
#	subnet = "0.0.0.0 0.0.0.0"
#	comment = "dd"
#}
#
#
#resource "fortios_system_admin_administrator" "cdaddf" {
#	name = "addaca"
#	password = "cc37331AC1"
#	trusthost1 = "1.1.1.0 255.255.255.0"
#	trusthost2 = "2.2.2.0 255.255.255.0"
#	trusthost3 = "0.0.0.0 0.0.0.0"
#	trusthost4 = "0.0.0.0 0.0.0.0"
#	trusthost5 = "0.0.0.0 0.0.0.0"
#	trusthost6 = "0.0.0.0 0.0.0.0"
#	trusthost7 = "0.0.0.0 0.0.0.0"
#	trusthost8 = "0.0.0.0 0.0.0.0"
#	trusthost9 = "0.0.0.0 0.0.0.0"
#	trusthost10 = "0.0.0.0 0.0.0.0"
#	accprofile = "3d3"
#	comments = "zzzzzzzzzzzzzzzzz"
#}
#
#resource "fortios_system_admin_profiles" "tst11" {
#	name = "223d3"
#	scope = "vdom"
#	comments = "test"
#	secfabgrp = "read-write"
#	ftviewgrp = "read"
#	authgrp = "none"
#	sysgrp = "read"
#	netgrp = "none"
#	loggrp = "none"
#	fwgrp = "none"
#	vpngrp = "none"
#	utmgrp = "none"
#	wanoptgrp = "none"
#	wifi = "none"
#	admintimeout_override = "disable"
#}
#
#resource "fortios_networking_interface_port" "loopback1" {
#	ip = "23.123.33.10/24"
#	allowaccess = "ping http"
#	alias = "cc1"
#	description = ">fdasdfafdsalfdsaf"
#	status = "up"
#	role = "lan"
#	name = "fddadddddsa"
#	vdom = "root"
#	type = "loopback"
#	mode = "static"
#}
#
#resource "fortios_networking_interface_port" "vlan1" {
#	role = "lan"
#	mode = "static"
#	defaultgw = "enable"
#	distance = "33"
#	type = "vlan"
#	vlanid = "3"
#	name = "fdddsa"
#	vdom = "root"
#	ip = "3.123.33.10/24"
#	interface = "port2"
#	allowaccess = "ping"
#}
#
#resource "fortios_networking_interface_port" "test1" {
#	portname = "port2"
#	ip = "93.133.133.110/24"
#	alias = "dkeeew"
#	description = "description..."
#	status = "up"
#	device_identification = "enable"
#	tcp_mss = "3232"
#	speed = "auto"
#	mtu_override = "enable"
#	mtu = "2933"
#	role = "lan"
#	allowaccess = "ping https"
#	mode = "static"
#	dns_server_override = "enable"
#	defaultgw = "enable"
#	distance = "33"
#	type = "physical"
#}
#
#
#resource "fortios_networking_route_static" "test1" {
#	dst = "110.2.2.122/32"
#	gateway = "2.2.2.2"
#	blackhole = "disable"
#	distance = "22"
#	weight = "3"
#	priority = "3"
#	device = "port2"
#	comment = "Terraform test"
#}
#
#####################################################
#resource "fortios_firewall_policy" "test1" {
#	name = "test_policy1"
#	srcintf = ["port2", "port1"]  
#	dstintf = ["port3"]  
#	srcaddr = ["swscan.apple.com", "google_play"]  
#	dstaddr = ["all"]  
#	schedule = "always"
#	service = ["ALL_ICMP"]  
#	action = "deny"
#}
#
#resource "fortios_system_global" "test1" {
#	hostname = "terraformtest123"
#	#admintimeout = "77"
#}
