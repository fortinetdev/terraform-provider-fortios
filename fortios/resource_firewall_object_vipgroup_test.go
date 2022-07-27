package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFortiOSFirewallObjectVipGroup_basic(t *testing.T) {
	rname := "fovgb_" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectVipGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectVipGroupConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectVipGroupExists("fortios_firewall_object_vipgroup.test1"),
					resource.TestCheckResourceAttr("fortios_firewall_object_vipgroup.test1", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_object_vipgroup.test1", "comments", "Terraform Test"),
					resource.TestCheckResourceAttr("fortios_firewall_object_vipgroup.test1", "interface", "port3"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallObjectVipGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Firewall Object VipGroup: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Firewall Object VipGroup is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectVipGroup(i)

		if err != nil {
			return fmt.Errorf("Error reading Firewall Object VipGroup: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Firewall Object VipGroup: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallObjectVipGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_object_vipgroup" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectVipGroup(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Firewall Object VipGroup %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallObjectVipGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_object_vip" "viptest1" {
	name = "vip1fortest%s"
	comment = "Terraform Test"
	extip = "1.1.15.0-1.1.16.0"
	mappedip = ["1.1.18.0-1.1.19.0"]
	extintf = "port3"
	portforward = "enable"
	protocol = "tcp"
	extport = "2-3"
	mappedport = "4-5"
}

resource "fortios_firewall_object_vip" "viptest2" {
	name = "vip2fortest%s"
	comment = "Terraform Test"
	extip = "1.1.17.0-1.1.18.0"
	mappedip = ["1.1.20.0-1.1.21.0"]
	extintf = "port3"
	portforward = "enable"
	protocol = "tcp"
	extport = "2-3"
	mappedport = "4-5"
}


resource "fortios_firewall_object_vipgroup" "test1" {
	name = "%s"
	interface = "port3"
	comments = "Terraform Test"
	member = ["${fortios_firewall_object_vip.viptest1.name}", "${fortios_firewall_object_vip.viptest2.name}"]
}
`, name, name, name)
}
