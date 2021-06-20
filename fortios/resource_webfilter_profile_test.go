// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSWebfilterProfile_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebfilterProfile_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebfilterProfileConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebfilterProfileExists("fortios_webfilter_profile.trname"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "extended_log", "disable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "https_replacemsg", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "inspection_mode", "flow-based"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "log_all_url", "disable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "post_action", "normal"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_content_log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_extended_all_action_log", "disable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_filter_activex_log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_filter_applet_log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_filter_command_block_log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_filter_cookie_log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_filter_cookie_removal_log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_filter_js_log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_filter_jscript_log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_filter_referer_log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_filter_unknown_log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_filter_vbs_log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_ftgd_err_log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_ftgd_quota_usage", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_invalid_domain_log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web_url_log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "wisp", "disable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "wisp_algorithm", "auto-learning"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "youtube_channel_status", "disable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.exempt_quota", "17"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.max_quota_timeout", "300"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.rate_crl_urls", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.rate_css_urls", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.rate_image_urls", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.rate_javascript_urls", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.filters.0.action", "warning"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.filters.0.category", "2"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.filters.0.id", "1"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.filters.0.log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.filters.0.warn_duration", "5m"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.filters.0.warning_duration_type", "timeout"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.filters.0.warning_prompt", "per-category"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.filters.1.action", "warning"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.filters.1.category", "7"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.filters.1.id", "2"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.filters.1.log", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.filters.1.warn_duration", "5m"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.filters.1.warning_duration_type", "timeout"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "ftgd_wf.0.filters.1.warning_prompt", "per-category"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "override.0.ovrd_cookie", "deny"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "override.0.ovrd_dur", "15m"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "override.0.ovrd_dur_mode", "constant"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "override.0.ovrd_scope", "user"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "override.0.profile_attribute", "Login-LAT-Service"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "override.0.profile_type", "list"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web.0.blacklist", "disable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web.0.bword_table", "0"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web.0.bword_threshold", "10"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web.0.content_header_list", "0"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web.0.log_search", "disable"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web.0.urlfilter_table", "0"),
					resource.TestCheckResourceAttr("fortios_webfilter_profile.trname", "web.0.youtube_restrict", "none"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebfilterProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebfilterProfile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebfilterProfile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebfilterProfile(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading WebfilterProfile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebfilterProfile: %s", n)
		}

		return nil
	}
}

func testAccCheckWebfilterProfileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webfilter_profile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebfilterProfile(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebfilterProfile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebfilterProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webfilter_profile" "trname" {
  extended_log                  = "disable"
  https_replacemsg              = "enable"
  inspection_mode               = "flow-based"
  log_all_url                   = "disable"
  name                          = "%[1]s"
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
`, name)
}
