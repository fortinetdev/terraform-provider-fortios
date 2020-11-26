---
subcategory: "FortiGate WebFilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_profile"
description: |-
  Configure Web filter profiles.
---

# fortios_webfilter_profile
Configure Web filter profiles.

## Example Usage

```hcl
resource "fortios_webfilter_profile" "trname" {
  extended_log                  = "disable"
  https_replacemsg              = "enable"
  inspection_mode               = "flow-based"
  log_all_url                   = "disable"
  name                          = "1"
  post_action                   = "normal"
  web_content_log               = "enable"
  web_extended_all_action_log   = "disable"
  web_filter_activex_log        = "enable"
  web_filter_applet_log         = "enable"
  web_filter_command_block_log  = "enable"
  web_filter_cookie_log         = "enable"
  web_filter_cookie_removal_log = "enable"
  web_filter_js_log             = "enable"
  web_filter_jscript_log        = "enable"
  web_filter_referer_log        = "enable"
  web_filter_unknown_log        = "enable"
  web_filter_vbs_log            = "enable"
  web_ftgd_err_log              = "enable"
  web_ftgd_quota_usage          = "enable"
  web_invalid_domain_log        = "enable"
  web_url_log                   = "enable"
  wisp                          = "disable"
  wisp_algorithm                = "auto-learning"
  youtube_channel_status        = "disable"

  ftgd_wf {
    exempt_quota         = "17"
    max_quota_timeout    = 300
    rate_crl_urls        = "enable"
    rate_css_urls        = "enable"
    rate_image_urls      = "enable"
    rate_javascript_urls = "enable"

    filters {
      action                = "warning"
      category              = 2
      id                    = 1
      log                   = "enable"
      warn_duration         = "5m"
      warning_duration_type = "timeout"
      warning_prompt        = "per-category"
    }
    filters {
      action                = "warning"
      category              = 7
      id                    = 2
      log                   = "enable"
      warn_duration         = "5m"
      warning_duration_type = "timeout"
      warning_prompt        = "per-category"
    }
  }

  override {
    ovrd_cookie       = "deny"
    ovrd_dur          = "15m"
    ovrd_dur_mode     = "constant"
    ovrd_scope        = "user"
    profile_attribute = "Login-LAT-Service"
    profile_type      = "list"
  }

  web {
    blacklist           = "disable"
    bword_table         = 0
    bword_threshold     = 10
    content_header_list = 0
    log_search          = "disable"
    urlfilter_table     = 0
    youtube_restrict    = "none"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Profile name.
* `comment` - Optional comments.
* `replacemsg_group` - Replacement message group.
* `inspection_mode` - Web filtering inspection mode.
* `options` - Options.
* `https_replacemsg` - Enable replacement messages for HTTPS.
* `ovrd_perm` - Permitted override types.
* `post_action` - Action taken for HTTP POST traffic.
* `override` - Web Filter override settings. The structure of `override` block is documented below.
* `web` - Web content filtering settings. The structure of `web` block is documented below.
* `youtube_channel_status` - YouTube channel filter status.
* `youtube_channel_filter` - YouTube channel filter. The structure of `youtube_channel_filter` block is documented below.
* `ftgd_wf` - FortiGuard Web Filter settings. The structure of `ftgd_wf` block is documented below.
* `wisp` - Enable/disable web proxy WISP.
* `wisp_servers` - WISP servers. The structure of `wisp_servers` block is documented below.
* `wisp_algorithm` - WISP server selection algorithm.
* `log_all_url` - Enable/disable logging all URLs visited.
* `web_content_log` - Enable/disable logging logging blocked web content.
* `web_filter_activex_log` - Enable/disable logging ActiveX.
* `web_filter_command_block_log` - Enable/disable logging blocked commands.
* `web_filter_cookie_log` - Enable/disable logging cookie filtering.
* `web_filter_applet_log` - Enable/disable logging Java applets.
* `web_filter_jscript_log` - Enable/disable logging JScripts.
* `web_filter_js_log` - Enable/disable logging Java scripts.
* `web_filter_vbs_log` - Enable/disable logging VBS scripts.
* `web_filter_unknown_log` - Enable/disable logging unknown scripts.
* `web_filter_referer_log` - Enable/disable logging referrers.
* `web_filter_cookie_removal_log` - Enable/disable logging blocked cookies.
* `web_url_log` - Enable/disable logging URL filtering.
* `web_invalid_domain_log` - Enable/disable logging invalid domain names.
* `web_ftgd_err_log` - Enable/disable logging rating errors.
* `web_ftgd_quota_usage` - Enable/disable logging daily quota usage.
* `extended_log` - Enable/disable extended logging for web filtering.
* `web_extended_all_action_log` - Enable/disable extended any filter action logging for web filtering.

The `override` block supports:

* `ovrd_cookie` - Allow/deny browser-based (cookie) overrides.
* `ovrd_scope` - Override scope.
* `profile_type` - Override profile type.
* `ovrd_dur_mode` - Override duration mode.
* `ovrd_dur` - Override duration.
* `profile_attribute` - Profile attribute to retrieve from the RADIUS server.
* `ovrd_user_group` - User groups with permission to use the override. The structure of `ovrd_user_group` block is documented below.
* `profile` - Web filter profile with permission to create overrides. The structure of `profile` block is documented below.

The `ovrd_user_group` block supports:

* `name` - User group name.

The `profile` block supports:

* `name` - Web profile.

The `web` block supports:

* `bword_threshold` - Banned word score threshold.
* `bword_table` - Banned word table ID.
* `urlfilter_table` - URL filter table ID.
* `content_header_list` - Content header list.
* `blacklist` - Enable/disable automatic addition of URLs detected by FortiSandbox to blacklist.
* `whitelist` - FortiGuard whitelist settings.
* `safe_search` - Safe search type.
* `youtube_restrict` - YouTube EDU filter level.
* `log_search` - Enable/disable logging all search phrases.
* `keyword_match` - Search keywords to log when match is found. The structure of `keyword_match` block is documented below.

The `keyword_match` block supports:

* `pattern` - Pattern/keyword to search for.

The `youtube_channel_filter` block supports:

* `id` - ID.
* `channel_id` - YouTube channel ID to be filtered.
* `comment` - Comment.

The `ftgd_wf` block supports:

* `options` - Options for FortiGuard Web Filter.
* `exempt_quota` - Do not stop quota for these categories.
* `ovrd` - Allow web filter profile overrides.
* `filters` - FortiGuard filters. The structure of `filters` block is documented below.
* `quota` - FortiGuard traffic quota settings. The structure of `quota` block is documented below.
* `max_quota_timeout` - Maximum FortiGuard quota used by single page view in seconds (excludes streams).
* `rate_image_urls` - Enable/disable rating images by URL.
* `rate_javascript_urls` - Enable/disable rating JavaScript by URL.
* `rate_css_urls` - Enable/disable rating CSS by URL.
* `rate_crl_urls` - Enable/disable rating CRL by URL.

The `filters` block supports:

* `id` - ID number.
* `category` - Categories and groups the filter examines.
* `action` - Action to take for matches.
* `warn_duration` - Duration of warnings.
* `auth_usr_grp` - Groups with permission to authenticate. The structure of `auth_usr_grp` block is documented below.
* `log` - Enable/disable logging.
* `override_replacemsg` - Override replacement message.
* `warning_prompt` - Warning prompts in each category or each domain.
* `warning_duration_type` - Re-display warning after closing browser or after a timeout.

The `auth_usr_grp` block supports:

* `name` - User group name.

The `quota` block supports:

* `id` - ID number.
* `category` - FortiGuard categories to apply quota to (category action must be set to monitor).
* `type` - Quota type.
* `unit` - Traffic quota unit of measurement.
* `value` - Traffic quota value.
* `duration` - Duration of quota.
* `override_replacemsg` - Override replacement message.

The `wisp_servers` block supports:

* `name` - Server name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Webfilter Profile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webfilter_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
