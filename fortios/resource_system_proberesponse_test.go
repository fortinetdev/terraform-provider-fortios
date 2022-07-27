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

func TestAccFortiOSSystemProbeResponse_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemProbeResponse_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemProbeResponseConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemProbeResponseExists("fortios_system_proberesponse.trname"),
					resource.TestCheckResourceAttr("fortios_system_proberesponse.trname", "http_probe_value", "OK"),
					resource.TestCheckResourceAttr("fortios_system_proberesponse.trname", "mode", "none"),
					resource.TestCheckResourceAttr("fortios_system_proberesponse.trname", "port", "8008"),
					resource.TestCheckResourceAttr("fortios_system_proberesponse.trname", "security_mode", "none"),
					resource.TestCheckResourceAttr("fortios_system_proberesponse.trname", "timeout", "300"),
					resource.TestCheckResourceAttr("fortios_system_proberesponse.trname", "ttl_mode", "retain"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemProbeResponseExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemProbeResponse: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemProbeResponse is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemProbeResponse(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemProbeResponse: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemProbeResponse: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemProbeResponseDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_proberesponse" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemProbeResponse(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemProbeResponse %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemProbeResponseConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_proberesponse" "trname" {
  http_probe_value = "OK"
  mode             = "none"
  port             = 8008
  security_mode    = "none"
  timeout          = 300
  ttl_mode         = "retain"
}
`)
}
