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
* `feature_set` - Flow/proxy feature set. Valid values: `flow`, `proxy`.
* `replacemsg_group` - Replacement message group.
* `inspection_mode` - Web filtering inspection mode. Valid values: `proxy`, `flow-based`.
* `options` - Options.
* `file_filter` - File filter. The structure of `file_filter` block is documented below.
* `https_replacemsg` - Enable replacement messages for HTTPS. Valid values: `enable`, `disable`.
* `ovrd_perm` - Permitted override types. Valid values: `bannedword-override`, `urlfilter-override`, `fortiguard-wf-override`, `contenttype-check-override`.
* `post_action` - Action taken for HTTP POST traffic. Valid values: `normal`, `block`.
* `override` - Web Filter override settings. The structure of `override` block is documented below.
* `web` - Web content filtering settings. The structure of `web` block is documented below.
* `youtube_channel_status` - YouTube channel filter status.
* `youtube_channel_filter` - YouTube channel filter. The structure of `youtube_channel_filter` block is documented below.
* `ftgd_wf` - FortiGuard Web Filter settings. The structure of `ftgd_wf` block is documented below.
* `antiphish` - AntiPhishing profile. The structure of `antiphish` block is documented below.
* `wisp` - Enable/disable web proxy WISP. Valid values: `enable`, `disable`.
* `wisp_servers` - WISP servers. The structure of `wisp_servers` block is documented below.
* `wisp_algorithm` - WISP server selection algorithm. Valid values: `primary-secondary`, `round-robin`, `auto-learning`.
* `log_all_url` - Enable/disable logging all URLs visited. Valid values: `enable`, `disable`.
* `web_content_log` - Enable/disable logging logging blocked web content. Valid values: `enable`, `disable`.
* `web_filter_activex_log` - Enable/disable logging ActiveX. Valid values: `enable`, `disable`.
* `web_filter_command_block_log` - Enable/disable logging blocked commands. Valid values: `enable`, `disable`.
* `web_filter_cookie_log` - Enable/disable logging cookie filtering. Valid values: `enable`, `disable`.
* `web_filter_applet_log` - Enable/disable logging Java applets. Valid values: `enable`, `disable`.
* `web_filter_jscript_log` - Enable/disable logging JScripts. Valid values: `enable`, `disable`.
* `web_filter_js_log` - Enable/disable logging Java scripts. Valid values: `enable`, `disable`.
* `web_filter_vbs_log` - Enable/disable logging VBS scripts. Valid values: `enable`, `disable`.
* `web_filter_unknown_log` - Enable/disable logging unknown scripts. Valid values: `enable`, `disable`.
* `web_filter_referer_log` - Enable/disable logging referrers. Valid values: `enable`, `disable`.
* `web_filter_cookie_removal_log` - Enable/disable logging blocked cookies. Valid values: `enable`, `disable`.
* `web_url_log` - Enable/disable logging URL filtering. Valid values: `enable`, `disable`.
* `web_invalid_domain_log` - Enable/disable logging invalid domain names. Valid values: `enable`, `disable`.
* `web_ftgd_err_log` - Enable/disable logging rating errors. Valid values: `enable`, `disable`.
* `web_ftgd_quota_usage` - Enable/disable logging daily quota usage. Valid values: `enable`, `disable`.
* `extended_log` - Enable/disable extended logging for web filtering. Valid values: `enable`, `disable`.
* `web_extended_all_action_log` - Enable/disable extended any filter action logging for web filtering. Valid values: `enable`, `disable`.
* `web_antiphishing_log` - Enable/disable logging of AntiPhishing checks. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `file_filter` block supports:

* `status` - Enable/disable file filter. Valid values: `enable`, `disable`.
* `log` - Enable/disable file filter logging. Valid values: `enable`, `disable`.
* `scan_archive_contents` - Enable/disable file filter archive contents scan. Valid values: `enable`, `disable`.
* `entries` - File filter entries. The structure of `entries` block is documented below.

The `entries` block supports:

* `filter` - Add a file filter.
* `comment` - Comment.
* `protocol` - Protocols to apply with. Valid values: `http`, `ftp`.
* `action` - Action taken for matched file. Valid values: `log`, `block`.
* `direction` - Match files transmitted in the session's originating or reply direction. Valid values: `incoming`, `outgoing`, `any`.
* `password_protected` - Match password-protected files. Valid values: `yes`, `any`.
* `file_type` - Select file type. The structure of `file_type` block is documented below.

The `file_type` block supports:

* `name` - File type name.

The `override` block supports:

* `ovrd_cookie` - Allow/deny browser-based (cookie) overrides. Valid values: `allow`, `deny`.
* `ovrd_scope` - Override scope. Valid values: `user`, `user-group`, `ip`, `browser`, `ask`.
* `profile_type` - Override profile type. Valid values: `list`, `radius`.
* `ovrd_dur_mode` - Override duration mode. Valid values: `constant`, `ask`.
* `ovrd_dur` - Override duration.
* `profile_attribute` - Profile attribute to retrieve from the RADIUS server. Valid values: `User-Name`, `NAS-IP-Address`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Filter-Id`, `Login-IP-Host`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `Class`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Zone`, `Acct-Session-Id`, `Acct-Multi-Session-Id`.
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
* `blocklist` - Enable/disable automatic addition of URLs detected by FortiSandbox to blocklist. Valid values: `enable`, `disable`.
* `allowlist` - FortiGuard allowlist settings. Valid values: `exempt-av`, `exempt-webcontent`, `exempt-activex-java-cookie`, `exempt-dlp`, `exempt-rangeblock`, `extended-log-others`.
* `blacklist` - Enable/disable automatic addition of URLs detected by FortiSandbox to blacklist. Valid values: `enable`, `disable`.
* `whitelist` - FortiGuard whitelist settings. Valid values: `exempt-av`, `exempt-webcontent`, `exempt-activex-java-cookie`, `exempt-dlp`, `exempt-rangeblock`, `extended-log-others`.
* `safe_search` - Safe search type. Valid values: `url`, `header`.
* `youtube_restrict` - YouTube EDU filter level. Valid values: `none`, `strict`, `moderate`.
* `log_search` - Enable/disable logging all search phrases. Valid values: `enable`, `disable`.
* `keyword_match` - Search keywords to log when match is found. The structure of `keyword_match` block is documented below.

The `keyword_match` block supports:

* `pattern` - Pattern/keyword to search for.

The `youtube_channel_filter` block supports:

* `id` - ID.
* `channel_id` - YouTube channel ID to be filtered.
* `comment` - Comment.

The `ftgd_wf` block supports:

* `options` - Options for FortiGuard Web Filter. Valid values: `error-allow`, `rate-server-ip`, `connect-request-bypass`, `ftgd-disable`.
* `exempt_quota` - Do not stop quota for these categories.
* `ovrd` - Allow web filter profile overrides.
* `filters` - FortiGuard filters. The structure of `filters` block is documented below.
* `quota` - FortiGuard traffic quota settings. The structure of `quota` block is documented below.
* `max_quota_timeout` - Maximum FortiGuard quota used by single page view in seconds (excludes streams).
* `rate_image_urls` - Enable/disable rating images by URL. Valid values: `disable`, `enable`.
* `rate_javascript_urls` - Enable/disable rating JavaScript by URL. Valid values: `disable`, `enable`.
* `rate_css_urls` - Enable/disable rating CSS by URL. Valid values: `disable`, `enable`.
* `rate_crl_urls` - Enable/disable rating CRL by URL. Valid values: `disable`, `enable`.

The `filters` block supports:

* `id` - ID number.
* `category` - Categories and groups the filter examines.
* `action` - Action to take for matches. Valid values: `block`, `authenticate`, `monitor`, `warning`.
* `warn_duration` - Duration of warnings.
* `auth_usr_grp` - Groups with permission to authenticate. The structure of `auth_usr_grp` block is documented below.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `override_replacemsg` - Override replacement message.
* `warning_prompt` - Warning prompts in each category or each domain. Valid values: `per-domain`, `per-category`.
* `warning_duration_type` - Re-display warning after closing browser or after a timeout. Valid values: `session`, `timeout`.

The `auth_usr_grp` block supports:

* `name` - User group name.

The `quota` block supports:

* `id` - ID number.
* `category` - FortiGuard categories to apply quota to (category action must be set to monitor).
* `type` - Quota type. Valid values: `time`, `traffic`.
* `unit` - Traffic quota unit of measurement. Valid values: `B`, `KB`, `MB`, `GB`.
* `value` - Traffic quota value.
* `duration` - Duration of quota.
* `override_replacemsg` - Override replacement message.

The `antiphish` block supports:

* `status` - Toggle AntiPhishing functionality. Valid values: `enable`, `disable`.
* `domain_controller` - Domain for which to verify received credentials against.
* `default_action` - Action to be taken when there is no matching rule. Valid values: `exempt`, `log`, `block`.
* `check_uri` - Enable/disable checking of GET URI parameters for known credentials. Valid values: `enable`, `disable`.
* `check_basic_auth` - Enable/disable checking of HTTP Basic Auth field for known credentials. Valid values: `enable`, `disable`.
* `max_body_len` - Maximum size of a POST body to check for credentials.
* `inspection_entries` - AntiPhishing entries. The structure of `inspection_entries` block is documented below.
* `custom_patterns` - Custom username and password regex patterns. The structure of `custom_patterns` block is documented below.

The `inspection_entries` block supports:

* `name` - Inspection target name.
* `fortiguard_category` - FortiGuard category to match.
* `action` - Action to be taken upon an AntiPhishing match. Valid values: `exempt`, `log`, `block`.

The `custom_patterns` block supports:

* `pattern` - Target pattern.
* `category` - Category that the pattern matches. Valid values: `username`, `password`.

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
