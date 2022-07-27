// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

func TestAccFortiOSSystemReplacemsgAdmin_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemReplacemsgAdmin_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemReplacemsgAdminConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemReplacemsgAdminExists("fortios_systemreplacemsg_admin.trname"),
					resource.TestCheckResourceAttr("fortios_systemreplacemsg_admin.trname", "buffer", "CSD"),
					resource.TestCheckResourceAttr("fortios_systemreplacemsg_admin.trname", "format", "text"),
					resource.TestCheckResourceAttr("fortios_systemreplacemsg_admin.trname", "header", "none"),
					resource.TestCheckResourceAttr("fortios_systemreplacemsg_admin.trname", "msg_type", "post_admin-disclaimer-text"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemReplacemsgAdminExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemReplacemsgAdmin: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemReplacemsgAdmin is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemReplacemsgAdmin(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemReplacemsgAdmin: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemReplacemsgAdmin: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemReplacemsgAdminDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_systemreplacemsg_admin" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemReplacemsgAdmin(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemReplacemsgAdmin %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemReplacemsgAdminConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_systemreplacemsg_admin" "trname" {
  buffer   = "CSD"
  format   = "text"
  header   = "none"
  msg_type = "post_admin-disclaimer-text"
}
`)
}
