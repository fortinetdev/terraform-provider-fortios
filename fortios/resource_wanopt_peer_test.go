
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSWanoptPeer_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSWanoptPeer_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSWanoptPeerConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSWanoptPeerExists("fortios_wanopt_peer.trname"),
                    resource.TestCheckResourceAttr("fortios_wanopt_peer.trname", "ip", "1.1.1.1"),
                    resource.TestCheckResourceAttr("fortios_wanopt_peer.trname", "peer_host_id", "1"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSWanoptPeerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WanoptPeer: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WanoptPeer is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWanoptPeer(i)

		if err != nil {
			return fmt.Errorf("Error reading WanoptPeer: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WanoptPeer: %s", n)
		}

		return nil
	}
}

func testAccCheckWanoptPeerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wanopt_peer" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWanoptPeer(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WanoptPeer %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWanoptPeerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wanopt_peer" "trname" {
  ip           = "1.1.1.1"
  peer_host_id = "1"
}
`)
}
