package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFortiOSFirewallSecurityPolicySeq_basic(t *testing.T) {
	rname := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallSecurityPolicySeqConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallSecurityPolicySeqExists("fortios_firewall_security_policyseq.test1"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallSecurityPolicySeqExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Firewall Security Policy Sequence: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Security Policy Sequence Configure is set")
		}

		return nil
	}
}

func testAccFortiOSFirewallSecurityPolicySeqConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_security_policy" "test1" {
	name = "name-%s"
	srcintf = ["port2"]
	dstintf = ["port3"]
	srcaddr = ["all"]
	dstaddr = ["all"]
	schedule = "always"
	service = ["ALL"]
	action = "accept"
}

resource "fortios_firewall_security_policy" "test2" {
	name = "%s-2"
	srcintf = ["port3"]
	dstintf = ["port2"]
	srcaddr = ["all"]
	dstaddr = ["all"]
	schedule = "always"
	service = ["ALL"]
	action = "accept"
}

resource "fortios_firewall_security_policyseq" "test1" {
        policy_src_id = "${fortios_firewall_security_policy.test1.id}"
        policy_dst_id = "${fortios_firewall_security_policy.test2.id}"
        alter_position = "after"
}
`, name, "${fortios_firewall_security_policy.test1.name}")
}
