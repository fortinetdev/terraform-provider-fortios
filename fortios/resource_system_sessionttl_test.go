// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSSystemSessionTtl_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemSessionTtl_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemSessionTtlConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemSessionTtlExists("fortios_system_sessionttl.trname"),
					resource.TestCheckResourceAttr("fortios_system_sessionttl.trname", "default", "3600"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemSessionTtlExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemSessionTtl: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemSessionTtl is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemSessionTtl(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemSessionTtl: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemSessionTtl: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemSessionTtlDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_sessionttl" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSessionTtl(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemSessionTtl %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemSessionTtlConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_sessionttl" "trname" {
  default = "3600"
}
`)
}
