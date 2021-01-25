// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSEndpointControlForticlientRegistrationSync_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSEndpointControlForticlientRegistrationSync_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSEndpointControlForticlientRegistrationSyncConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSEndpointControlForticlientRegistrationSyncExists("fortios_endpointcontrol_forticlientregistrationsync.trname"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_forticlientregistrationsync.trname", "peer_ip", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_forticlientregistrationsync.trname", "peer_name", "1"),
				),
			},
		},
	})
}

func testAccCheckFortiOSEndpointControlForticlientRegistrationSyncExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found EndpointControlForticlientRegistrationSync: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No EndpointControlForticlientRegistrationSync is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadEndpointControlForticlientRegistrationSync(i)

		if err != nil {
			return fmt.Errorf("Error reading EndpointControlForticlientRegistrationSync: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating EndpointControlForticlientRegistrationSync: %s", n)
		}

		return nil
	}
}

func testAccCheckEndpointControlForticlientRegistrationSyncDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_endpointcontrol_forticlientregistrationsync" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadEndpointControlForticlientRegistrationSync(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error EndpointControlForticlientRegistrationSync %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSEndpointControlForticlientRegistrationSyncConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_endpointcontrol_forticlientregistrationsync" "trname" {
  peer_ip   = "1.1.1.1"
  peer_name = "1"
}
`)
}
