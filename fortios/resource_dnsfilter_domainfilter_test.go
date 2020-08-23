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

func TestAccFortiOSDnsfilterDomainFilter_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSDnsfilterDomainFilter_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSDnsfilterDomainFilterConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSDnsfilterDomainFilterExists("fortios_dnsfilter_domainfilter.trname"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_domainfilter.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_domainfilter.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_dnsfilter_domainfilter.trname", "entries.0.action", "block"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_domainfilter.trname", "entries.0.domain", "bac.com"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_domainfilter.trname", "entries.0.id", "1"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_domainfilter.trname", "entries.0.status", "enable"),
					resource.TestCheckResourceAttr("fortios_dnsfilter_domainfilter.trname", "entries.0.type", "simple"),
				),
			},
		},
	})
}

func testAccCheckFortiOSDnsfilterDomainFilterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found DnsfilterDomainFilter: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No DnsfilterDomainFilter is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadDnsfilterDomainFilter(i)

		if err != nil {
			return fmt.Errorf("Error reading DnsfilterDomainFilter: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating DnsfilterDomainFilter: %s", n)
		}

		return nil
	}
}

func testAccCheckDnsfilterDomainFilterDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_dnsfilter_domainfilter" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadDnsfilterDomainFilter(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error DnsfilterDomainFilter %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSDnsfilterDomainFilterConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_dnsfilter_domainfilter" "trname" {
  fosid = 1
  name  = "%[1]s"

  entries {
    action = "block"
    domain = "bac.com"
    id     = 1
    status = "enable"
    type   = "simple"
  }
}
`, name)
}
