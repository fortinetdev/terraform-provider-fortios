// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSSystemNdProxy_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemNdProxy_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemNdProxyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemNdProxyExists("fortios_system_ndproxy.trname"),
					resource.TestCheckResourceAttr("fortios_system_ndproxy.trname", "status", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemNdProxyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemNdProxy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemNdProxy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemNdProxy(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading SystemNdProxy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemNdProxy: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemNdProxyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_ndproxy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemNdProxy(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemNdProxy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemNdProxyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_ndproxy" "trname" {
  status = "disable"
}
`)
}
