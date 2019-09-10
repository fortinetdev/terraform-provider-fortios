package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiManagerFirewallSecurityPolicy(t *testing.T) {
	name := "fmg-security-policy" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGFirewallSecurityPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerFirewallSecurityPolicyConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerFirewallSecurityPolicyExists("fortios_fortimanager_firewall_security_policy.test1"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "name", name),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "action", "accept"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "internet_service", "enable"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "internet_service_src", "disable"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "rsso", "disable"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "fsso", "enable"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "logtraffic", "all"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "logtraffic_start", "enable"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "capture_packet", "enable"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "nat", "enable"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "ippool", "disable"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "fixedport", "enable"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "utm_status", "enable"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "profile_type", "single"),
					resource.TestCheckResourceAttr("fortios_fortimanager_firewall_security_policy.test1", "comments", "policy test"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerFirewallSecurityPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found firewall security policy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No firewall security policy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		i := rs.Primary.ID
		o, err := c.ReadFirewallSecurityPolicy(i, "default")

		if err != nil {
			return fmt.Errorf("Error reading firewall security policy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating firewall security policy: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGFirewallSecurityPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_fortimanager_firewall_security_policy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallSecurityPolicy(i, "default")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error firewall security policy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiManagerFirewallSecurityPolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_fortimanager_firewall_security_policy" "test1" {
	name = "%s"
	srcaddr = ["all"]
    srcintf = ["any"]
    dstaddr = ["all"]
    dstintf = ["any"]
    service = ["ALL"]
    action = "accept"
    schedule = ["always"]
    internet_service = "enable"
    internet_service_id = ["Alibaba.Web", "AOL.Web"]
    internet_service_src = "disable"
    users = ["guest"]
    groups = ["Guest-group"]
    rsso = "disable"
    fsso = "enable"
    logtraffic = "all"
    logtraffic_start = "enable"
    capture_packet = "enable"
    nat = "enable"
    ippool = "disable"
    fixedport = "enable"
    utm_status = "enable"
    profile_type = "single"
    av_profile = ["g-default"]
    dnsfilter_profile = ["default"]
    traffic_shaper = ["high-priority"]
    comments = "policy test"

}
`, name)
}
