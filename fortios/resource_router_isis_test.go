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

func TestAccFortiOSRouterIsis_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSRouterIsis_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSRouterIsisConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSRouterIsisExists("fortios_router_isis.trname"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "adjacency_check", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "adjacency_check6", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "adv_passive_only", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "adv_passive_only6", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "auth_mode_l1", "password"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "auth_mode_l2", "password"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "auth_sendonly_l1", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "auth_sendonly_l2", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "default_originate", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "default_originate6", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "dynamic_hostname", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "ignore_lsp_errors", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "is_type", "level-1-2"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "lsp_gen_interval_l1", "30"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "lsp_gen_interval_l2", "30"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "lsp_refresh_interval", "900"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "max_lsp_lifetime", "1200"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "metric_style", "narrow"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "overload_bit", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "redistribute6_l1", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "redistribute6_l2", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "redistribute_l1", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "redistribute_l2", "disable"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "spf_interval_exp_l1", "500 50000"),
					resource.TestCheckResourceAttr("fortios_router_isis.trname", "spf_interval_exp_l2", "500 50000"),
				),
			},
		},
	})
}

func testAccCheckFortiOSRouterIsisExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterIsis: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterIsis is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterIsis(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterIsis: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterIsis: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterIsisDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_isis" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterIsis(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterIsis %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterIsisConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_isis" "trname" {
  adjacency_check      = "disable"
  adjacency_check6     = "disable"
  adv_passive_only     = "disable"
  adv_passive_only6    = "disable"
  auth_mode_l1         = "password"
  auth_mode_l2         = "password"
  auth_sendonly_l1     = "disable"
  auth_sendonly_l2     = "disable"
  default_originate    = "disable"
  default_originate6   = "disable"
  dynamic_hostname     = "disable"
  ignore_lsp_errors    = "disable"
  is_type              = "level-1-2"
  lsp_gen_interval_l1  = 30
  lsp_gen_interval_l2  = 30
  lsp_refresh_interval = 900
  max_lsp_lifetime     = 1200
  metric_style         = "narrow"
  overload_bit         = "disable"
  redistribute6_l1     = "disable"
  redistribute6_l2     = "disable"
  redistribute_l1      = "disable"
  redistribute_l2      = "disable"
  spf_interval_exp_l1  = "500 50000"
  spf_interval_exp_l2  = "500 50000"

}
`)
}
