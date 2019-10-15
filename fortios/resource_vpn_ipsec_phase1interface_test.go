package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSVPNIPsecPhase1Interface_basic(t *testing.T) {
	rname := "i1" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVPNIPsecPhase1InterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVPNIPsecPhase1InterfaceConfig1(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVPNIPsecPhase1InterfaceExists("fortios_vpn_ipsec_phase1interface.test1"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "name", rname),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "type", "static"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "interface", "port3"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "peertype", "any"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "proposal", "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "comments", "VPN 001Test P1"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "wizard_type", "static-fortigate"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "remote_gw", "5.2.2.2"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "psksecret", "testscecret112233445566778899"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "authmethod", "psk"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "mode_cfg", "disable"),
				),
			},
		},
	})
}

func TestAccFortiOSVPNIPsecPhase1Interface_basic2(t *testing.T) {
	rname := "i2" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVPNIPsecPhase1InterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVPNIPsecPhase1InterfaceConfig2(rname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "name", rname),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "type", "dynamic"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "interface", "port3"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "peertype", "any"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "proposal", "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "comments", "VPN 001Test P2"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "wizard_type", "dialup-forticlient"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase1interface.test1", "remote_gw", "0.0.0.0"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVPNIPsecPhase1InterfaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VPN IPsec Phase1Interface: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VPN IPsec Phase1Interface is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVPNIPsecPhase1Interface(i)

		if err != nil {
			return fmt.Errorf("Error reading VPN IPsec Phase1Interface: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VPN IPsec Phase1Interface: %s", n)
		}

		return nil
	}
}

func testAccCheckVPNIPsecPhase1InterfaceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpn_ipsec_phase1interface" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVPNIPsecPhase1Interface(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VPN IPsec Phase1Interface %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVPNIPsecPhase1InterfaceConfig1(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpn_ipsec_phase1interface" "test1" {
	name = "%s"
	type = "static"
	interface = "port3"
	peertype = "any"
	proposal = "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"
	comments = "VPN 001Test P1"
	wizard_type = "static-fortigate"
	remote_gw = "5.2.2.2"
	psksecret = "testscecret112233445566778899"
	authmethod = "psk"
	authmethod_remote = ""
	mode_cfg = "disable"
}
`, name)
}

func testAccFortiOSVPNIPsecPhase1InterfaceConfig2(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpn_ipsec_phase1interface" "test1" {
	name = "%s"
	type = "dynamic"
	interface = "port3"
	peertype = "any"
	proposal = "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"
	comments = "VPN 001Test P2"
	wizard_type = "dialup-forticlient"
	remote_gw = "0.0.0.0"
	psksecret = "testscecret112233445566778899"
	ipv4_split_include = ""
	split_include_service = ""
	ipv4_split_exclude = ""
}
`, name)
}
