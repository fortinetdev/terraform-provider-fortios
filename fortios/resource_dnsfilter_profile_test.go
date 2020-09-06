// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSDnsfilterProfile_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSDnsfilterProfile_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSDnsfilterProfileConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSDnsfilterProfileExists("fortios_dnsfilter_profile.trname"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "block_action", "redirect"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "block_botnet", "disable"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "log_all_domain", "disable"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "redirect_portal", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "safe_search", "disable"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "sdns_domain_log", "enable"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "sdns_ftgd_err_log", "enable"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "youtube_restrict", "strict"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "domain_filter.0.domain_filter_table", "0"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.0.action", "block"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.0.category", "26"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.0.id", "1"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.0.log", "enable"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.1.action", "block"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.1.category", "61"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.1.id", "2"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.1.log", "enable"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.2.action", "block"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.2.category", "86"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.2.id", "3"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.2.log", "enable"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.3.action", "block"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.3.category", "88"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.3.id", "4"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_profile.trname", "ftgd_dns.0.filters.3.log", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSDnsfilterProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found DnsfilterProfile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No DnsfilterProfile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadDnsfilterProfile(i)

		if err != nil {
			return fmt.Errorf("Error reading DnsfilterProfile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating DnsfilterProfile: %s", n)
		}

		return nil
	}
}

func testAccCheckDnsfilterProfileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_dnsfilter_profile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadDnsfilterProfile(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error DnsfilterProfile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSDnsfilterProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_dnsfilter_profile" "trname" {
  block_action      = "redirect"
  block_botnet      = "disable"
  log_all_domain    = "disable"
  name              = "%[1]s"
  redirect_portal   = "0.0.0.0"
  safe_search       = "disable"
  sdns_domain_log   = "enable"
  sdns_ftgd_err_log = "enable"
  youtube_restrict  = "strict"

  domain_filter {
    domain_filter_table = 0
  }

  ftgd_dns {
    filters {
      action   = "block"
      category = 26
      id       = 1
      log      = "enable"
    }
    filters {
      action   = "block"
      category = 61
      id       = 2
      log      = "enable"
    }
    filters {
      action   = "block"
      category = 86
      id       = 3
      log      = "enable"
    }
    filters {
      action   = "block"
      category = 88
      id       = 4
      log      = "enable"
    }
  }
}
`, name)
}
