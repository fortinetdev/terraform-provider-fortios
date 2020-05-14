package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSFirewallObjectServiceGroup_basic(t *testing.T) {
	rname := "fosgb_" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectServiceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectServiceGroupConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectServiceGroupExists("fortios_firewall_object_servicegroup.test1"),
					resource.TestCheckResourceAttr("fortios_firewall_object_servicegroup.test1", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_object_servicegroup.test1", "comment", "Terraform Test")),
			},
		},
	})
}

func testAccCheckFortiOSFirewallObjectServiceGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Firewall Object Service Group is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client
		i := rs.Primary.ID

		o, err := c.ReadFirewallObjectServiceGroup(i)
		if err != nil {
			return fmt.Errorf("Error reading Firewall Object Service Group: %s", err)
		}

		if o == nil {
			return fmt.Errorf("firewall object service group %s was not created", n)
		}

		return nil
	}
}

func testAccCheckFirewallObjectServiceGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_object_servicegroup" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectServiceGroup(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Firewall Object ServiceGroup %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallObjectServiceGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_object_servicegroup" "test1" {
	name = "%s"
	comment = "Terraform Test"
	member = ["DCE-RPC", "DNS", "HTTPS"]
}
`, name)
}
