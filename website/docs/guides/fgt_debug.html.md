---
subcategory: ""
layout: "fortios"
page_title: "Debugging for FortiGate"
description: |-
  Debugging for FortiGate.
---

# Debugging for FortiGate	

Verbose output can help with debugging. To enable verbose debugging, use the following commands in the FortiGate CLI:

```shell
$ diagnose debug enable
$ diagnose debug application httpsd -1
$ diagnose debug cli 8
```

Debug messages will be displayed for 30 minutes and will include debug messages for all requests to/from the FortiOS web interface. CLI commands for any configuration changes will be logged to the console.

When the REST API for IPv4 policy traffic statistics is queried, REST API related debug messages will be similar to the following:

```hcl
[httpsd 4137 - 1588331103     info] ap_invoke_handler[593] -- new request (handler='api_cmdb_v2-handler', uri='/api/v2/cmdb/system/global?access_token=******************************', method='PUT')
[httpsd 4137 - 1588331103     info] ap_invoke_handler[597] -- User-Agent: Go-http-client/1.1
[httpsd 4137 - 1588331103     info] ap_invoke_handler[600] -- Source: 192.168.52.1:64379 Destination: 192.168.52.177:443
[httpsd 4137 - 1588331103     info] api_cmdb_v2_handler[2048] -- received api_cmdb_v2_request from '192.168.52.1'
[httpsd 4137 - 1588331103  warning] api_access_check_for_api_key[957] -- API Key request authorized for test-apiuser from 192.168.52.1.
[httpsd 4137 - 1588331103     info] api_store_parameter[227] -- add API parameter 'access_token': '********' (type=string)
[httpsd 4137 - 1588331103     info] api_store_parameter[227] -- add API parameter 'admintimeout': '"65"' (type=string)
[httpsd 4137 - 1588331103     info] api_store_parameter[227] -- add API parameter 'timezone': '"04"' (type=string)
[httpsd 4137 - 1588331103     info] api_store_parameter[227] -- add API parameter 'hostname': '"mytestfirewall"' (type=string)
[httpsd 4137 - 1588331103     info] api_store_parameter[227] -- add API parameter 'admin-sport': '"443"' (type=string)
[httpsd 4137 - 1588331103     info] api_store_parameter[227] -- add API parameter 'admin-ssh-port': '"22"' (type=string)
[httpsd 4137 - 1588331103     info] handle_cli_req_v2_vdom[1950] -- new CMDB API request (vdom='root',user='test-apiuser')
[httpsd 4137 - 1588331103     info] api_cmdb_request_init_by_path[1351] -- new CMDB query (path='system',name='global')
[httpsd 4137 - 1588331103     info] api_return_cmdb_revision[779] -- ETag check for system.global
[httpsd 4137 - 1588331103     info] _api_cmdb_v2_config[1136] -- editing CLI object (append=0, auto_key=0, path=system, name=global, mkey=(null), flags=0)
[httpsd 4137 - 1588331103     info] api_set_cmdb_attr[1685] -- 'admintimeout': '65'
[httpsd 4137 - 1588331103     info] api_set_cmdb_attr[1685] -- 'timezone': '04'
[httpsd 4137 - 1588331103     info] api_set_cmdb_attr[1685] -- 'hostname': 'mytestfirewall'
[httpsd 4137 - 1588331103     info] api_set_cmdb_attr[1685] -- 'admin-sport': '443'
[httpsd 4137 - 1588331103     info] api_set_cmdb_attr[1685] -- 'admin-ssh-port': '22'
0: config system global
0: set hostname "mytestfirewall"
0: end
[httpsd 4137 - 1588331103     info] cmdb_save_with_children[269] -- appended main node (nret=0, is_new=0)
[httpsd 4137 - 1588331103     info] ap_invoke_handler[616] -- request completed (handler='api_cmdb_v2-handler' result==0)
[httpsd 4137 - 1588331103     info] ap_invoke_handler[593] -- new request (handler='api_cmdb_v2-handler', uri='/api/v2/cmdb/system/global?access_token=******************************', method='GET')
[httpsd 4137 - 1588331103     info] ap_invoke_handler[597] -- User-Agent: Go-http-client/1.1
[httpsd 4137 - 1588331103     info] ap_invoke_handler[600] -- Source: 192.168.52.1:64379 Destination: 192.168.52.177:443
[httpsd 4137 - 1588331103     info] api_cmdb_v2_handler[2048] -- received api_cmdb_v2_request from '192.168.52.1'
[httpsd 4137 - 1588331103  warning] api_access_check_for_api_key[957] -- API Key request authorized for test-apiuser from 192.168.52.1.
[httpsd 4137 - 1588331103     info] api_store_parameter[227] -- add API parameter 'access_token': '********' (type=string)
[httpsd 4137 - 1588331103     info] handle_cli_req_v2_vdom[1950] -- new CMDB API request (vdom='root',user='test-apiuser')
[httpsd 4137 - 1588331103     info] api_cmdb_request_init_by_path[1351] -- new CMDB query (path='system',name='global')
[httpsd 4137 - 1588331103     info] api_cmdb_select_etag[2062] -- ETag check for system.global
[httpsd 4137 - 1588331103     info] api_return_cmdb_revision[779] -- ETag check for system.global
[httpsd 4137 - 1588331103     info] api_add_etag[864] -- no If-None-Match header
[httpsd 4137 - 1588331103     info] api_return_cmdb_revision[779] -- ETag check for system.global
[httpsd 4137 - 1588331103     info] ap_invoke_handler[616] -- request completed (handler='api_cmdb_v2-handler' result==0)
write config file success, prepare to save in flash
[__create_file_new_version:257] the new version config file '/data/./config/sys_global.conf.gz.v000000172' is created
[symlink_config_file:324] a new version of '/data/./config/sys_global.conf.gz' is created: /data/./config/sys_global.conf.gz.v000000172
[symlink_config_file:367] the old version '/data/./config/sys_global.conf.gz.v000000171' is deleted
[symlink_config_file:370] '/data/./config/sys_global.conf.gz' has been symlink'ed to the new version '/data/./config/sys_global.conf.gz.v000000172'. The old version '/data/./config/sys_global.conf.gz.v000000171' has been deleted
zip config file /data/./config/sys_global.conf.gz success!
```

