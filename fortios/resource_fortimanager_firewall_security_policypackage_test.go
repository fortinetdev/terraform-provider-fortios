package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiManagerFirewallSecurityPolicyPackage(t *testing.T) {
	name := "fmg-security-policypackage-" + acctest.RandString(4)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGFirewallSecurityPolicyPackageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerFirewallSecurityPolicyPackageConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerFirewallSecurityPolicyPackageExists("fortios_fortimanager_firewall_security_policypackage.test1"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policypackage.test1", "name", name),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policypackage.test1", "target", "FGVM64-test"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerFirewallSecurityPolicyPackageExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found firewall security policypackage: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No firewall security policypackage is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		i := rs.Primary.ID
		o, err := c.ReadFirewallSecurityPolicyPackage(i)

		if err != nil {
			return fmt.Errorf("Error reading firewall security policypackage: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating firewall security policypackage: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGFirewallSecurityPolicyPackageDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_fortimanager_firewall_security_policypackage" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallSecurityPolicyPackage(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error firewall security policypackage %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiManagerFirewallSecurityPolicyPackageConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_fortimanager_firewall_security_policypackage" "test1" {
	name = "%s"
	target = "FGVM64-test"
}
`, name)
}
