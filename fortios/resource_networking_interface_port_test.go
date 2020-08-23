package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSNetworkingInterfacePort_Physical(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// CheckDestroy: testAccCheckNetworkingInterfacePortDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSNetworkingInterfacePortConfigPhysical,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSNetworkingInterfacePortExists("fortios_networking_interface_port.test1"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "name", "port3"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "ip", "93.133.133.110 255.255.255.0"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "alias", "physicalport3"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "status", "up"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "device_identification", "enable"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "tcp_mss", "3232"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "speed", "auto"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "mtu_override", "enable"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "mtu", "1500"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "role", "lan"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "allowaccess", "ping https"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "mode", "static"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "dns_server_override", "enable"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "defaultgw", "enable"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "distance", "5"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test1", "type", "physical"),
				),
			},
		},
	})
}

func TestAccFortiOSNetworkingInterfacePort_Vlan(t *testing.T) {
	rname := acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkingInterfacePortDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSNetworkingInterfacePortConfigVlan(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSNetworkingInterfacePortExists("fortios_networking_interface_port.test2"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test2", "name", rname),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test2", "role", "lan"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test2", "mode", "static"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test2", "defaultgw", "enable"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test2", "distance", "5"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test2", "type", "vlan"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test2", "vlanid", "3"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test2", "vdom", "root"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test2", "ip", "3.123.33.10 255.255.255.0"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test2", "interface", "port2"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test2", "allowaccess", "ping"),
				),
			},
		},
	})
}

func TestAccFortiOSNetworkingInterfacePort_Loopback(t *testing.T) {
	rname := acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkingInterfacePortDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSNetworkingInterfacePortConfigLoopback(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSNetworkingInterfacePortExists("fortios_networking_interface_port.test3"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test3", "name", rname),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test3", "ip", "23.123.33.10 255.255.255.0"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test3", "allowaccess", "ping https http"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test3", "alias", "loopbackportalias"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test3", "description", "Terraform Test"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test3", "mtu_override", "disable"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test3", "status", "up"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test3", "role", "lan"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test3", "vdom", "root"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test3", "type", "loopback"),
					resource.TestCheckResourceAttr("fortios_networking_interface_port.test3", "mode", "static"),
				),
			},
		},
	})
}

func testAccCheckFortiOSNetworkingInterfacePortExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Networking Interface Port: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Networking Interface Port is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadNetworkingInterfacePort(i)

		if err != nil {
			return fmt.Errorf("Error reading Networking Interface Port: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Networking Interface Port: %s", n)
		}

		return nil
	}
}

func testAccCheckNetworkingInterfacePortDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_networking_interface_port" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadNetworkingInterfacePort(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Networking Interface Port %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

const testAccFortiOSNetworkingInterfacePortConfigPhysical = `
resource "fortios_networking_interface_port" "test1" {
	name = "port3"
	ip = "93.133.133.110 255.255.255.0"
	alias = "physicalport3"
	status = "up"
	device_identification = "enable"
	tcp_mss = "3232"
	speed = "auto"
	mtu_override = "enable"
	mtu = "1500"
	role = "lan"
	allowaccess = "ping https"
	mode = "static"
	dns_server_override = "enable"
	defaultgw = "enable"
	distance = "5"
	type = "physical"
}
`

func testAccFortiOSNetworkingInterfacePortConfigVlan(name string) string {
	return fmt.Sprintf(`
resource "fortios_networking_interface_port" "test2" {
	name = "%s"
	role = "lan"
	mode = "static"
	defaultgw = "enable"
	distance = "5"
	type = "vlan"
	vlanid = "3"
	vdom = "root"
	ip = "3.123.33.10 255.255.255.0"
	interface = "port2"
	allowaccess = "ping"
}
`, name)
}

func testAccFortiOSNetworkingInterfacePortConfigLoopback(name string) string {
	return fmt.Sprintf(`
resource "fortios_networking_interface_port" "test3" {
	name = "%s"
	ip = "23.123.33.10 255.255.255.0"
	allowaccess = "ping https http"
	alias = "loopbackportalias"
	description = "Terraform Test"
	status = "up"
	role = "lan"
	vdom = "root"
	type = "loopback"
	mode = "static"
	mtu_override = "disable"
}
`, name)
}
