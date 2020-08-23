
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSSystemDnsDatabase_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemDnsDatabase_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemDnsDatabaseConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemDnsDatabaseExists("fortios_system_dnsdatabase.trname"),
                    resource.TestCheckResourceAttr("fortios_system_dnsdatabase.trname", "authoritative", "enable"),
                    resource.TestCheckResourceAttr("fortios_system_dnsdatabase.trname", "contact", "hostmaster"),
                    resource.TestCheckResourceAttr("fortios_system_dnsdatabase.trname", "domain", "s.com"),
                    resource.TestCheckResourceAttr("fortios_system_dnsdatabase.trname", "ip_master", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_system_dnsdatabase.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_system_dnsdatabase.trname", "primary_name", "dns"),
                    resource.TestCheckResourceAttr("fortios_system_dnsdatabase.trname", "source_ip", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_system_dnsdatabase.trname", "status", "enable"),
                    resource.TestCheckResourceAttr("fortios_system_dnsdatabase.trname", "ttl", "86400"),
                    resource.TestCheckResourceAttr("fortios_system_dnsdatabase.trname", "type", "master"),
                    resource.TestCheckResourceAttr("fortios_system_dnsdatabase.trname", "view", "shadow"),
                    resource.TestCheckResourceAttr("fortios_system_dnsdatabase.trname", "dns_entry.0.type", "MX"),
                    resource.TestCheckResourceAttr("fortios_system_dnsdatabase.trname", "dns_entry.0.ttl", "3"),
                    resource.TestCheckResourceAttr("fortios_system_dnsdatabase.trname", "dns_entry.0.hostname", "sghsgh.com"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemDnsDatabaseExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemDnsDatabase: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemDnsDatabase is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemDnsDatabase(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemDnsDatabase: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemDnsDatabase: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemDnsDatabaseDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_dnsdatabase" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemDnsDatabase(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemDnsDatabase %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemDnsDatabaseConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_dnsdatabase" "trname" {
  authoritative = "enable"
  contact       = "hostmaster"
  domain        = "s.com"
  ip_master     = "0.0.0.0"
  name          = "%[1]s"
  primary_name  = "dns"
  source_ip     = "0.0.0.0"
  status        = "enable"
  ttl           = 86400
  type          = "master"
  view          = "shadow"
  dns_entry {
    type     = "MX"
    ttl      = 3
    hostname = "sghsgh.com"
  }
}
`, name)
}
