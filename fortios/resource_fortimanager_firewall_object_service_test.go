package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFortiManagerFirewallObjectService(t *testing.T) {
	name := "fmg-firewall-object-service" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGFirewallObjectServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerFirewallObjectServiceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerFirewallObjectServiceExists("fortios_fmg_firewall_object_service.test1"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_service.test1", "name", name),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_service.test1", "protocol", "TCP/UDP/SCTP"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_service.test1", "comment", "test obj service"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_service.test1", "category", "Email"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_service.test1", "proxy", "disable"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_service.test1", "fqdn", "fqdn.google.com"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerFirewallObjectServiceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Firewall Object Service: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Firewall Object Service is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectService("root", i)

		if err != nil {
			return fmt.Errorf("Error reading Firewall Object Service: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Firewall Object Service: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGFirewallObjectServiceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_fmg_firewall_object_service" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectService("root", i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Firewall Object Service %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiManagerFirewallObjectServiceConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_fmg_firewall_object_service" "test1" {
    name = "%s"
    comment = "test obj service"
    protocol = "TCP/UDP/SCTP"
    category = "Email"
    proxy = "disable"
    fqdn = "fqdn.google.com"
}
`, name)
}
