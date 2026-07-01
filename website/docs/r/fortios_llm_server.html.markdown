---
subcategory: "FortiGate Llm"
layout: "fortios"
page_title: "FortiOS: fortios_llm_server"
description: |-
  Configure LLM Proxy servers.
---

# fortios_llm_server
Configure LLM Proxy servers. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `name` - LLM Proxy server name.
* `display_name` - display name of the LLM Server
* `type` - LLM server type Valid values: `built-in`, `customized`.
* `built_in_server` - built-in LLM server Valid values: `openai`, `azure`, `azure-openai`, `gemini`, `anthropic`, `grok`, `gemini-with-openai-api`, `anthropic-with-openai-api`.
* `azure_resource_name` - Azure resource name.
* `end_point` - Overwrite the default end-point of the vendor.
* `chat_completions_api` - Chat Completions API of this server Valid values: `none`, `openai`, `azure-openai`, `gemini`, `anthropic`.
* `image_gen_api` - Image-Gen API of this server Valid values: `none`, `openai`, `azure-openai`.
* `anthropic_version` - Anthropic version in API
* `azure_api_version` - Azure API version.
* `models` - models of the LLM Server
* `accept_custom_model` - accept custom model Valid values: `enable`, `disable`.
* `custom_model_allow_regex` - custom model allow regex, allow all if empty
* `custom_model_block_regex` - custom model block regex, no block if empty
* `verify_cert` - Enable/disable certificate verification. Valid values: `enable`, `disable`.
* `api_key` - API Keys of the LLM server
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Llm Server can be imported using any of these accepted formats:
```
$ terraform import fortios_llm_server.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_llm_server.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
