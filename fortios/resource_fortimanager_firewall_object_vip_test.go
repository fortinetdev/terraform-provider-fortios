package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFortiManagerFirewallObjectVIp(t *testing.T) {
	name := "fmg-firewall-object-vip-" + acctest.RandString(2)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGFirewallObjectVIpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerFirewallObjectVIpConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerFirewallObjectVIpExists("fortios_fmg_firewall_object_vip.test1"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_vip.test1", "name", name),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_vip.test1", "type", "static-nat"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_vip.test1", "comment", "test obj vip"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_vip.test1", "arp_reply", "enable"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_vip.test1", "ext_ip", "2.2.2.2"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_vip.test1", "ext_intf", "any"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_vip.test1", "mapped_ip", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_fmg_firewall_object_vip.test1", "config_default", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerFirewallObjectVIpExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Firewall Object VIp: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Firewall Object VIp is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectVip("root", i)

		if err != nil {
			return fmt.Errorf("Error reading Firewall Object VIp: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Firewall Object VIp: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGFirewallObjectVIpDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_fmg_firewall_object_vip" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectVip("root", i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Firewall Object Virtual Ip %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiManagerFirewallObjectVIpConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_fmg_firewall_object_vip" "test1" {
    name = "%s"
	comment = "test obj vip"
    type = "static-nat"
    arp_reply = "enable"
    ext_ip = "2.2.2.2"
    ext_intf = "any"
    mapped_ip = "1.1.1.1"
    config_default = "enable"
}
`, name)
}
