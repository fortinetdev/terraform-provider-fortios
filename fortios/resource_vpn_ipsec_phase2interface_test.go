package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSVPNIPsecPhase2Interface_basic1(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVPNIPsecPhase2InterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVPNIPsecPhase2InterfaceConfig1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVPNIPsecPhase2InterfaceExists("fortios_vpn_ipsec_phase2interface.test1"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test1", "name", "002Test11"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test1", "phase1name", "001AccTest1"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test1", "proposal", "aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test1", "comments", "VPN 001Test P2"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test1", "src_addr_type", "name"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test1", "dst_addr_type", "name"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test1", "src_name", "all"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test1", "dst_name", "all"),
				),
			},
		},
	})
}

func TestAccFortiOSVPNIPsecPhase2Interface_basic2(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVPNIPsecPhase2InterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVPNIPsecPhase2InterfaceConfig2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVPNIPsecPhase2InterfaceExists("fortios_vpn_ipsec_phase2interface.test2"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test2", "name", "002Test12"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test2", "phase1name", "001AccTest2"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test2", "proposal", "aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test2", "comments", "VPN 001Test P2"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test2", "src_addr_type", "subnet"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test2", "src_start_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test2", "src_end_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test2", "src_subnet", "0.0.0.0 0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test2", "dst_addr_type", "subnet"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test2", "dst_start_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test2", "dst_end_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpn_ipsec_phase2interface.test2", "dst_subnet", "0.0.0.0 0.0.0.0"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVPNIPsecPhase2InterfaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VPN IPsec Phase2Interface: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VPN IPsec Phase2Interface is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVPNIPsecPhase2Interface(i)

		if err != nil {
			return fmt.Errorf("Error reading VPN IPsec Phase2Interface: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VPN IPsec Phase2Interface: %s", n)
		}

		return nil
	}
}

func testAccCheckVPNIPsecPhase2InterfaceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpn_ipsec_phase2interface" {
			continue
		}

		i := rs.Primary.ID
		_, err := c.ReadVPNIPsecPhase2Interface(i)

		if err == nil {
			return fmt.Errorf("Error VPN IPsec Phase2Interface %s still exists", rs.Primary.ID)
		}

		return nil
	}

	return nil
}

const testAccFortiOSVPNIPsecPhase2InterfaceConfig1 = `
resource "fortios_vpn_ipsec_phase1interface" "test1" {
	name = "001AccTest1"
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

resource "fortios_vpn_ipsec_phase2interface" "test1" {
	name = "002Test11"
	phase1name = "${fortios_vpn_ipsec_phase1interface.test1.name}"
	proposal = "aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"
	comments = "VPN 001Test P2"
	src_addr_type = "name"
	dst_addr_type = "name"
	src_name = "all"
	dst_name = "all"
}
`

const testAccFortiOSVPNIPsecPhase2InterfaceConfig2 = `
resource "fortios_vpn_ipsec_phase1interface" "test1" {
	name = "001AccTest2"
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
resource "fortios_vpn_ipsec_phase2interface" "test2" {
	name = "002Test12"
	phase1name = "${fortios_vpn_ipsec_phase1interface.test1.name}"
	proposal = "aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"
	comments = "VPN 001Test P2"
	src_addr_type = "subnet"
	src_start_ip = "0.0.0.0"
	src_end_ip = "0.0.0.0"
	src_subnet = "0.0.0.0 0.0.0.0"
	dst_addr_type = "subnet"
	dst_start_ip = "0.0.0.0"
	dst_end_ip = "0.0.0.0"
	dst_subnet = "0.0.0.0 0.0.0.0"
}
`
