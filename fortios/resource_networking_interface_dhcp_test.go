package fortios

import (
	"fmt"
	forticlient "github.com/fgtdev/fortios-sdk-go/sdkcore"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSNetworkingInterfaceDHCP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFortiOSNetworkingInterfaceDHCPDestroy,
		Steps: []resource.TestStep{
			{
				// Setup Interface to test with, Not working
				Config:  testAccFortiOSNetworkingInterfacePortConfigDHCP,
				Destroy: false,
				//Check: resource.ComposeTestCheckFunc(
				//	testAccCheckFortiOSNetworkingInterfacePortExists("fortios_networking_interface_port.test99"),
				//	resource.TestCheckResourceAttr("fortios_networking_interface_port.test99", "name", "TestIntf99"),
				//),
			},

			{
				Config: testAccFortiOSNetworkingInterfaceDHCPConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSNetworkingInterfaceDHCPExists("fortios_networking_interface_dhcp.test1"),
					resource.TestCheckResourceAttr("fortios_networking_interface_dhcp.test1", "interface", "TestIntf"),
					resource.TestCheckResourceAttr("fortios_networking_interface_dhcp.test1", "ip_range", "10.10.10.10-10.10.10.100"),
					resource.TestCheckResourceAttr("fortios_networking_interface_dhcp.test1", "default_gw", "10.10.10.1"),
					resource.TestCheckResourceAttr("fortios_networking_interface_dhcp.test1", "netmask", "255.255.255.0"),
					resource.TestCheckResourceAttr("fortios_networking_interface_dhcp.test1", "vdom", "root"),
					resource.TestCheckResourceAttr("fortios_networking_interface_dhcp.test1", "dns_service", "default"),
					resource.TestCheckResourceAttr("fortios_networking_interface_dhcp.test1", "dns_server_1", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_networking_interface_dhcp.test1", "dns_server_2", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_networking_interface_dhcp.test1", "dns_server_3", "0.0.0.0"),
				),
			},
		},
	})
}

// TODO Network interface needs to be created for test case to work.

const testAccFortiOSNetworkingInterfaceDHCPConfig = `
resource "fortios_networking_interface_dhcp" "test1" {
	interface = "TestIntf"
	ip_range = "10.10.10.10-10.10.10.100"
	default_gw = "10.10.10.1"
	netmask = "255.255.255.0"
	vdom = "root"
	dns_service = "default"
	dns_server_1 = "0.0.0.0"
	dns_server_2 = "0.0.0.0"
	dns_server_3 = "0.0.0.0"
}
`
const testAccFortiOSNetworkingInterfacePortConfigDHCP = `
resource "fortios_networking_interface_port" "test99" {
	name = "TestIntf99"
	allowaccess = "ping"
	ip = "1.1.1.1 255.255.255.0"
	description = "Terraform Test"
	vdom = "root"
	role = "lan"
	mode = "static"
	interface = "port1"
	vlanid = "99"
	distance = "5"
	defaultgw = "enable"
	type = "vlan"
}
`

func testAccCheckFortiOSNetworkingInterfaceDHCPExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Networking Interface DHCP resource named: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Networking Interface DHCP Server is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadNetworkingInterfaceDHCP(i)

		if err != nil {
			return fmt.Errorf("Error reading Networking Interface DHCP: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Networking Interface DHCP: %s", n)
		}

		return nil
	}
}

func testAccCheckFortiOSNetworkingInterfaceDHCPDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_networking_interface_dhcp" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadNetworkingInterfaceDHCP(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Networking Interface DHCP %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccCheckFortiOSNetworkingInterfaceDHCPPreStep() (mkey string, err error) {
	c := testAccProvider.Meta().(*FortiClient).Client

	in := &forticlient.JSONNetworkingInterfacePort{
		Ipf:                  "93.133.133.110 255.255.255.0",
		Alias:                "physicalport8",
		Status:               "up",
		DeviceIdentification: "enable",
		TCPMss:               "3232",
		Speed:                "auto",
		MtuOverride:          "enable",
		Mtu:                  "1500",
		Role:                 "lan",
		Allowaccess:          "ping https",
		Mode:                 "static",
		DNSServerOverride:    "enable",
		Defaultgw:            "enable",
		Distance:             "5",
		Description:          "Terraform Test",
		Type:                 "physical",
		Name:                 "port8",
	}

	o, err := c.CreateNetworkingInterfacePort(in)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return o.Mkey, nil
}

func testAccCheckFortiOSNetworkingInterfaceDHCPPostStep(mkey string) (err error) {
	c := testAccProvider.Meta().(*FortiClient).Client
	err = c.DeleteNetworkingInterfacePort(mkey)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
