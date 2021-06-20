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

func TestAccFortiOSSystemArpTable_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemArpTable_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemArpTableConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemArpTableExists("fortios_system_arptable.trname"),
					resource.TestCheckResourceAttr("fortios_system_arptable.trname", "fosid", "11"),
					resource.TestCheckResourceAttr("fortios_system_arptable.trname", "interface", "port2"),
					resource.TestCheckResourceAttr("fortios_system_arptable.trname", "ip", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_system_arptable.trname", "mac", "08:00:27:1c:a3:8b"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemArpTableExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemArpTable: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemArpTable is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemArpTable(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading SystemArpTable: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemArpTable: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemArpTableDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_arptable" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemArpTable(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemArpTable %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemArpTableConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_arptable" "trname" {
  fosid     = 11
  interface = "port2"
  ip        = "1.1.1.1"
  mac       = "08:00:27:1c:a3:8b"
}
`)
}
