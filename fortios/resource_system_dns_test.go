// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

func TestAccFortiOSSystemDns_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemDns_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemDnsConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemDnsExists("fortios_system_dns.trname"),
					resource.TestCheckResourceAttr("fortios_system_dns.trname", "cache_notfound_responses", "disable"),
					resource.TestCheckResourceAttr("fortios_system_dns.trname", "dns_cache_limit", "5000"),
					resource.TestCheckResourceAttr("fortios_system_dns.trname", "dns_cache_ttl", "1800"),
					resource.TestCheckResourceAttr("fortios_system_dns.trname", "ip6_primary", "::"),
					resource.TestCheckResourceAttr("fortios_system_dns.trname", "ip6_secondary", "::"),
					resource.TestCheckResourceAttr("fortios_system_dns.trname", "primary", "208.91.112.53"),
					resource.TestCheckResourceAttr("fortios_system_dns.trname", "retry", "2"),
					resource.TestCheckResourceAttr("fortios_system_dns.trname", "secondary", "208.91.112.51"),
					resource.TestCheckResourceAttr("fortios_system_dns.trname", "source_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_dns.trname", "timeout", "5"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemDnsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemDns: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemDns is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemDns(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemDns: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemDns: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemDnsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_dns" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemDns(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemDns %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemDnsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_dns" "trname" {
  cache_notfound_responses = "disable"
  dns_cache_limit          = 5000
  dns_cache_ttl            = 1800
  ip6_primary              = "::"
  ip6_secondary            = "::"
  primary                  = "208.91.112.53"
  retry                    = 2
  secondary                = "208.91.112.51"
  source_ip                = "0.0.0.0"
  timeout                  = 5
}
`)
}
