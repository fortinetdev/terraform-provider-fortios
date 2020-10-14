---
subcategory: ""
layout: "fortios"
page_title: "To move the policy"
description: |-
  Methods used to move the position of a policy, relative to another policy, in the sequence order of how policies are applied.
---

# Move a policy

Methods used to move the position of a policy, relative to another policy, in the sequence order of how policies are applied.

## Option I: Move with fortios_firewall_security_policyseq
* See resource [fortios_firewall_security_policyseq](https://registry.terraform.io/providers/fortinetdev/fortios/latest/docs/resources/fortios_firewall_security_policyseq) for further information.

## Option II: Quickly move with curl
1. List existing policies
    ```sh
    curl "https://192.168.52.177/api/v2/cmdb/firewall/policy?access_token=rGqsgj9Qmh3dwfQdc8hd3t3G6xG3N5&format=policyid|name|action"
    ```
    The access_token here is the same as the API token for FortiOS.

    It will return:

    ```{
      "http_method":"GET",
      "revision":"110.0.0.9541577349505417509.1602662160",
      "results":[
        {
          "q_origin_key":5,
          "policyid":5,
          "name":"22",
          "action":"accept"
        },
        {
          "q_origin_key":4,
          "policyid":4,
          "name":"fdscer435",
          "action":"accept"
        },
        {
          "q_origin_key":6,
          "policyid":6,
          "name":"e234552",
          "action":"accept"
        },
        {
          "q_origin_key":7,
          "policyid":7,
          "name":"fdsafcew222",
          "action":"accept"
        },
        {
          "q_origin_key":8,
          "policyid":8,
          "name":"e3232",
          "action":"accept"
        },
        {
          "q_origin_key":1,
          "policyid":1,
          "name":"policys1",
          "action":"accept"
        }
      ],
      "vdom":"root",
      "path":"firewall",
      "name":"policy",
      "status":"success",
      "http_status":200,
      "serial":"FGVM02TM20003062",
      "version":"v6.2.0",
      "build":776
    }
    ```

2. Move Policy 4 after Policy 7

    ```sh
    curl -X PUT "https://192.168.52.177/api/v2/cmdb/firewall/policy/4?access_token=rGqsgj9Qmh3dwfQdc8hd3t3G6xG3N5&action=move&after=7"
    ```

    It will return:

    ```{
      "http_method":"PUT",
      "revision":"110.0.0.9541577349505417509.1602662160",
      "mkey":4,
      "status":"success",
      "http_status":200,
      "vdom":"root",
      "path":"firewall",
      "name":"policy",
      "action":"move",
      "serial":"FGVM02TM20003062",
      "version":"v6.2.0",
      "build":776
    }
    ```
    -> 	**Note** If policy 4 needs to be moved to before 7, then only need to change `after=` to `before=`.

3. List existing policies again to check the result
    ```sh
    curl -k "https://192.168.52.177/api/v2/cmdb/firewall/polic?access_token=rGqsgj9Qmh3dwfQdc8hd3t3G6xG3N5&format=policyid|name|action"
    ```

    It will return:

    ```
    {
      "http_method":"GET",
      "revision":"112.0.0.9541577349505417509.1602662160",
      "results":[
        {
          "q_origin_key":5,
          "policyid":5,
          "name":"22",
          "action":"accept"
        },
        {
          "q_origin_key":6,
          "policyid":6,
          "name":"e234552",
          "action":"accept"
        },
        {
          "q_origin_key":7,
          "policyid":7,
          "name":"fdsafcew222",
          "action":"accept"
        },
        {
          "q_origin_key":4,
          "policyid":4,
          "name":"fdscer435",
          "action":"accept"
        },
        {
          "q_origin_key":8,
          "policyid":8,
          "name":"e3232",
          "action":"accept"
        },
        {
          "q_origin_key":1,
          "policyid":1,
          "name":"policys1",
          "action":"accept"
        }
      ],
      "vdom":"root",
      "path":"firewall",
      "name":"policy",
      "status":"success",
      "http_status":200,
      "serial":"FGVM02TM20003062",
      "version":"v6.2.0",
      "build":776
    }
    ```
    We can find that policy 4 has been moved to after 7.

## Option III: Move with fortios_json_generic_api

1. List existing policies
    ```hcl
    resource "fortios_json_generic_api" "test1" {
      path   = "/api/v2/cmdb/firewall/policy/"
      specialparams = "format=policyid|name|action"
      method = "GET"
      json   = ""
    }

    output "response1" {
      value = "${fortios_json_generic_api.test1.response}"
    }
    ```

    terraform apply will return:
    ```
    Apply complete! Resources: 0 added, 1 changed, 0 destroyed.

    Outputs:

    response1 = {
      "http_method":"GET",
      "revision":"109.0.0.9541577349505417509.1602662160",
      "results":[
        {
          "q_origin_key":5,
          "policyid":5,
          "name":"22",
          "action":"accept"
        },
        {
          "q_origin_key":4,
          "policyid":4,
          "name":"fdscer435",
          "action":"accept"
        },
        {
          "q_origin_key":6,
          "policyid":6,
          "name":"e234552",
          "action":"accept"
        },
        {
          "q_origin_key":7,
          "policyid":7,
          "name":"fdsafcew222",
          "action":"accept"
        },
        {
          "q_origin_key":8,
          "policyid":8,
          "name":"e3232",
          "action":"accept"
        },
        {
          "q_origin_key":1,
          "policyid":1,
          "name":"policys1",
          "action":"accept"
        }
      ],
      "vdom":"root",
      "path":"firewall",
      "name":"policy",
      "status":"success",
      "http_status":200,
      "serial":"FGVM02TM20003062",
      "version":"v6.2.0",
      "build":776
    }
    ```


2. Move Policy 4 after Policy 7
    ```hcl
    resource "fortios_json_generic_api" "test1" {
      path   = "/api/v2/cmdb/firewall/policy/4"
      specialparams = "action=move&after=7"
      method = "PUT"
      json   = ""
    }

    output "response1" {
      value = "${fortios_json_generic_api.test1.response}"
    }
    ```

    terraform apply will return:

    ```
    Apply complete! Resources: 1 added, 0 changed, 0 destroyed.

    Outputs:

    response1 = {
      "http_method":"PUT",
      "revision":"109.0.0.9541577349505417509.1602662160",
      "mkey":4,
      "status":"success",
      "http_status":200,
      "vdom":"root",
      "path":"firewall",
      "name":"policy",
      "action":"move",
      "serial":"FGVM02TM20003062",
      "version":"v6.2.0",
      "build":776
    }
    ```
    -> 	**Note** If policy 4 needs to be moved to before 7, then only need to change `after=` to `before=` in *specialparams* field .

3. List existing policies again to check the result
    ```hcl
    resource "fortios_json_generic_api" "test1" {
      path   = "/api/v2/cmdb/firewall/policy/"
      specialparams = "format=policyid|name|action"
      method = "GET"
      json   = ""
    }

    output "response1" {
      value = "${fortios_json_generic_api.test1.response}"
    }
    ```

    terraform apply will return:
    ```
    Apply complete! Resources: 0 added, 1 changed, 0 destroyed.

    Outputs:

    response1 = {
      "http_method":"GET",
      "revision":"110.0.0.9541577349505417509.1602662160",
      "results":[
        {
          "q_origin_key":5,
          "policyid":5,
          "name":"22",
          "action":"accept"
        },
        {
          "q_origin_key":6,
          "policyid":6,
          "name":"e234552",
          "action":"accept"
        },
        {
          "q_origin_key":7,
          "policyid":7,
          "name":"fdsafcew222",
          "action":"accept"
        },
        {
          "q_origin_key":4,
          "policyid":4,
          "name":"fdscer435",
          "action":"accept"
        },
        {
          "q_origin_key":8,
          "policyid":8,
          "name":"e3232",
          "action":"accept"
        },
        {
          "q_origin_key":1,
          "policyid":1,
          "name":"policys1",
          "action":"accept"
        }
      ],
      "vdom":"root",
      "path":"firewall",
      "name":"policy",
      "status":"success",
      "http_status":200,
      "serial":"FGVM02TM20003062",
      "version":"v6.2.0",
      "build":776
    }
    ```
    We can find that policy 4 has been moved to after 7.

## Last words
If you need to move a policy urgently, option II is the better choice. If you want to be consistent with Terraform, then option I is the better choice. Option III just provides an another additional method. You can choose one of these methods to move a policy according to your needs.
