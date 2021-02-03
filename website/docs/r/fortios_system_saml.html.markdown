---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_saml"
description: |-
  Global settings for SAML authentication.
---

# fortios_system_saml
Global settings for SAML authentication.

## Example Usage

```hcl
resource "fortios_system_saml" "trname" {
  default_login_page = "normal"
  default_profile    = "admin_no_access"
  life               = 30
  role               = "service-provider"
  status             = "disable"
  tolerance          = 5
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable SAML authentication (default = disable).
* `role` - SAML role.
* `default_login_page` - Choose default login page.
* `default_profile` - Default profile for new SSO admin.
* `cert` - Certificate to sign SAML messages.
* `portal_url` - SP portal URL.
* `entity_id` - SP entity ID.
* `single_sign_on_url` - SP single sign-on URL.
* `single_logout_url` - SP single logout URL.
* `idp_entity_id` - IDP entity ID.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `idp_single_logout_url` - IDP single logout URL.
* `idp_cert` - IDP certificate name.
* `server_address` - Server address.
* `tolerance` - Tolerance to the range of time when the assertion is valid (in minutes).
* `life` - Length of the range of time when the assertion is valid (in minutes).
* `service_providers` - Authorized service providers. The structure of `service_providers` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `service_providers` block supports:

* `name` - Name.
* `prefix` - Prefix.
* `sp_cert` - SP certificate name.
* `sp_entity_id` - SP entity ID.
* `sp_single_sign_on_url` - SP single sign-on URL.
* `sp_single_logout_url` - SP single logout URL.
* `sp_portal_url` - SP portal URL.
* `idp_entity_id` - IDP entity ID.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `idp_single_logout_url` - IDP single logout URL.
* `assertion_attributes` - Customized SAML attributes to send along with assertion. The structure of `assertion_attributes` block is documented below.

The `assertion_attributes` block supports:

* `name` - Name.
* `type` - Type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Saml can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_saml.labelname SystemSaml
$ unset "FORTIOS_IMPORT_TABLE"
```
