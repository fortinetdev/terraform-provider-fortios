package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSFirewallSecurityPolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallSecurityPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallSecurityPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallSecurityPolicyExists("fortios_firewall_security_policy.test1"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "name", "policy1"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "internet_service", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "action", "accept"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "schedule", "always"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "utm_status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "logtraffic", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "logtraffic_start", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "capture_packet", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "ippool", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "comments", "Terraform Test"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "av_profile", "wifi-default"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "webfilter_profile", "monitor-all"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "dnsfilter_profile", "default"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "ips_sensor", "protect_client"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "application_list", "block-high-risk"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "ssl_ssh_profile", "certificate-inspection"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "nat", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallSecurityPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Firewall Security Policy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Firewall Security Policy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallSecurityPolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading Firewall Security Policy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Firewall Security Policy: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallSecurityPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_security_policy" {
			continue
		}

		i := rs.Primary.ID
		_, err := c.ReadFirewallSecurityPolicy(i)

		if err == nil {
			return fmt.Errorf("Error Firewall Security Policy %s still exists", rs.Primary.ID)
		}

		return nil
	}

	return nil
}

const testAccFortiOSFirewallSecurityPolicyConfig = `
resource "fortios_firewall_security_policy" "test1" {
	name = "policy1"
	srcintf = ["port2"]
	dstintf = ["port3"]
	srcaddr = ["swscan.apple.com", "google-play"]    
	dstaddr = ["swscan.apple.com", "update.microsoft.com"]
	internet_service = "disable"
	internet_service_id = []
	schedule = "always"
	service = ["ALL_ICMP", "FTP"]
	action = "accept"	
	utm_status = "enable"
	logtraffic = "all"
	logtraffic_start = "enable"
	capture_packet = "enable"
	ippool = "enable"
	poolname = ["IPPool1fortest"]
	groups = ["Guest-group", "SSO_Guest_Users"]
	devices = ["android-phone", "android-tablet"]
	comments = "Terraform Test"
	av_profile = "wifi-default"
	webfilter_profile = "monitor-all"
	dnsfilter_profile = "default"
	ips_sensor = "protect_client"
	application_list = "block-high-risk"
	ssl_ssh_profile = "certificate-inspection"
	nat = "enable"
}
`
