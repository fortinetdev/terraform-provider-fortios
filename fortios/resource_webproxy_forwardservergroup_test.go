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

func TestAccFortiOSWebProxyForwardServerGroup_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSWebProxyForwardServerGroup_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebProxyForwardServerGroupConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebProxyForwardServerGroupExists("fortios_webproxy_forwardservergroup.trname1"),
					resource.TestCheckResourceAttr("fortios_webproxy_forwardservergroup.trname1", "affinity", "disable"),
					resource.TestCheckResourceAttr("fortios_webproxy_forwardservergroup.trname1", "group_down_option", "block"),
					resource.TestCheckResourceAttr("fortios_webproxy_forwardservergroup.trname1", "ldb_method", "weighted"),
					resource.TestCheckResourceAttr("fortios_webproxy_forwardservergroup.trname1", "name", rname),
					resource.TestCheckResourceAttr("fortios_webproxy_forwardservergroup.trname1", "server_list.0.name", var0),
					resource.TestCheckResourceAttr("fortios_webproxy_forwardservergroup.trname1", "server_list.0.weight", "12"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebProxyForwardServerGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebProxyForwardServerGroup: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebProxyForwardServerGroup is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebProxyForwardServerGroup(i)

		if err != nil {
			return fmt.Errorf("Error reading WebProxyForwardServerGroup: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebProxyForwardServerGroup: %s", n)
		}

		return nil
	}
}

func testAccCheckWebProxyForwardServerGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webproxy_forwardservergroup" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebProxyForwardServerGroup(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebProxyForwardServerGroup %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebProxyForwardServerGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webproxy_forwardserver" "trname1" {
  addr_type          = "fqdn"
  healthcheck        = "disable"
  ip                 = "0.0.0.0"
  monitor            = "http://www.google.com"
  name               = "var0%[1]s"
  port               = 1128
  server_down_option = "block"
}

resource "fortios_webproxy_forwardservergroup" "trname1" {
  affinity          = "disable"
  group_down_option = "block"
  ldb_method        = "weighted"
  name              = "%[1]s"

  server_list {
    name   = fortios_webproxy_forwardserver.trname1.name
    weight = 12
  }
}




`, name)
}
