---
subcategory: "FortiGate Llm"
layout: "fortios"
page_title: "FortiOS: fortios_llm_profile"
description: |-
  Configure LLM Proxy profiles.
---

# fortios_llm_profile
Configure LLM Proxy profiles. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `name` - LLM Proxy profile name.
* `unknown_api` - Support unknown API. Valid values: `enable`, `disable`.
* `log` - Log option. Valid values: `none`, `blocked`, `all`.
* `chat` - LLM Proxy chat completions API (/v1/chat/completions). The structure of `chat` block is documented below.
* `image_gen` - LLM Proxy image generation API (/v1/images/generations). The structure of `image_gen` block is documented below.
* `list_models` - LLM Proxy list models API (/v1/models). The structure of `list_models` block is documented below.
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.

The `chat` block supports:

* `status` - Support chat completions API. Valid values: `enable`, `disable`.
* `max_completion_tokens` - Maximum number of completion tokens (0 - 10000000, default = 0, means no limit).
* `stream` - support chat completions stream mode. Valid values: `bypass`, `block`.
* `max_req_len` - Max size of chat completions request body, (0 - 65535 KiB, default = 1024, 0 means no limit).
* `system_prompt_mode` - System prompt processing mode. Valid values: `bypass`, `replace`, `prepend`, `append`.
* `system_prompt` - Replace/Append chat completions system prompt.

The `image_gen` block supports:

* `status` - Support image generation API. Valid values: `enable`, `disable`.

The `list_models` block supports:

* `status` - Support list models API. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Llm Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_llm_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_llm_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
