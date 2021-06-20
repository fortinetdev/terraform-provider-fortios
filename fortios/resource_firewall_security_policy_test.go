package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSFirewallSecurityPolicy_basic(t *testing.T) {
	rname := "fspb_" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallSecurityPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallSecurityPolicyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallSecurityPolicyExists("fortios_firewall_security_policy.test1"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "internet_service", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "internet_service_src", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "schedule", "always"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "action", "accept"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "utm_status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "logtraffic", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "logtraffic_start", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "capture_packet", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "ippool", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "comments", "Terraform Test"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "av_profile", ""),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "webfilter_profile", ""),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "dnsfilter_profile", "default"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "ips_sensor", ""),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "application_list", ""),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "ssl_ssh_profile", "certificate-inspection"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "nat", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_security_policy.test1", "profile_protocol_options", "default"),
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
		o, err := c.ReadFirewallSecurityPolicy(i, "root")

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
		o, err := c.ReadFirewallSecurityPolicy(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Firewall Security Policy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallSecurityPolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_security_policy" "test1" {
	name = "%s"
	srcintf = ["port2"]
	dstintf = ["port3"]
	srcaddr = ["all"]
	dstaddr = ["all"]
	internet_service = "disable"
	internet_service_id = []
	internet_service_src = "disable"
	internet_service_src_id = []
	users = []
	status = "enable"
	schedule = "always"
	service = ["ALL"]
	action = "accept"
	utm_status = "enable"
	logtraffic = "all"
	logtraffic_start = "enable"
	capture_packet = "enable"
	ippool = "disable"
	poolname = []
	groups = []
	devices = []
	comments = "Terraform Test"
	av_profile = ""
	webfilter_profile = ""
	dnsfilter_profile = "default"
	ips_sensor = ""
	application_list = ""
	ssl_ssh_profile = "certificate-inspection"
	nat = "enable"
	profile_protocol_options = "default"
}
`, name)
}
