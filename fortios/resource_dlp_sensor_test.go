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

func TestAccFortiOSDlpSensor_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSDlpSensor_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSDlpSensorConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSDlpSensorExists("fortios_dlp_sensor.trname"),
					resource.TestCheckResourceAttr("fortios_dlp_sensor.trname", "dlp_log", "enable"),
					resource.TestCheckResourceAttr("fortios_dlp_sensor.trname", "extended_log", "disable"),
					resource.TestCheckResourceAttr("fortios_dlp_sensor.trname", "flow_based", "enable"),
					resource.TestCheckResourceAttr("fortios_dlp_sensor.trname", "nac_quar_log", "disable"),
					resource.TestCheckResourceAttr("fortios_dlp_sensor.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_dlp_sensor.trname", "summary_proto", "smtp pop3"),
				),
			},
		},
	})
}

func testAccCheckFortiOSDlpSensorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found DlpSensor: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No DlpSensor is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadDlpSensor(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading DlpSensor: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating DlpSensor: %s", n)
		}

		return nil
	}
}

func testAccCheckDlpSensorDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_dlp_sensor" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadDlpSensor(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error DlpSensor %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSDlpSensorConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_dlp_sensor" "trname" {
  dlp_log       = "enable"
  extended_log  = "disable"
  flow_based    = "enable"
  nac_quar_log  = "disable"
  name          = "%[1]s"
  summary_proto = "smtp pop3"
}
`, name)
}
