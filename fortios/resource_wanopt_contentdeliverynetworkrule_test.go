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

func TestAccFortiOSWanoptContentDeliveryNetworkRule_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWanoptContentDeliveryNetworkRule_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWanoptContentDeliveryNetworkRuleConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWanoptContentDeliveryNetworkRuleExists("fortios_wanopt_contentdeliverynetworkrule.trname"),
					resource.TestCheckResourceAttr("fortios_wanopt_contentdeliverynetworkrule.trname", "category", "vcache"),
					resource.TestCheckResourceAttr("fortios_wanopt_contentdeliverynetworkrule.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_wanopt_contentdeliverynetworkrule.trname", "request_cache_control", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_contentdeliverynetworkrule.trname", "response_cache_control", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_contentdeliverynetworkrule.trname", "response_expires", "enable"),
					resource.TestCheckResourceAttr("fortios_wanopt_contentdeliverynetworkrule.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_wanopt_contentdeliverynetworkrule.trname", "text_response_vcache", "enable"),
					resource.TestCheckResourceAttr("fortios_wanopt_contentdeliverynetworkrule.trname", "updateserver", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_contentdeliverynetworkrule.trname", "host_domain_name_suffix.0.name", "kaf.com"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWanoptContentDeliveryNetworkRuleExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WanoptContentDeliveryNetworkRule: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WanoptContentDeliveryNetworkRule is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWanoptContentDeliveryNetworkRule(i)

		if err != nil {
			return fmt.Errorf("Error reading WanoptContentDeliveryNetworkRule: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WanoptContentDeliveryNetworkRule: %s", n)
		}

		return nil
	}
}

func testAccCheckWanoptContentDeliveryNetworkRuleDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wanopt_contentdeliverynetworkrule" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWanoptContentDeliveryNetworkRule(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WanoptContentDeliveryNetworkRule %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWanoptContentDeliveryNetworkRuleConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wanopt_contentdeliverynetworkrule" "trname" {
  category               = "vcache"
  name                   = "%[1]s"
  request_cache_control  = "disable"
  response_cache_control = "disable"
  response_expires       = "enable"
  status                 = "enable"
  text_response_vcache   = "enable"
  updateserver           = "disable"

  host_domain_name_suffix {
    name = "kaf.com"
  }
}

`, name)
}
