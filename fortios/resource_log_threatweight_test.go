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

func TestAccFortiOSLogThreatWeight_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogThreatWeight_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogThreatWeightConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogThreatWeightExists("fortios_log_threatweight.trname"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "blocked_connection", "high"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "failed_connection", "low"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "url_block_detected", "high"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "application.0.category", "2"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "application.0.id", "1"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "application.0.level", "low"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "application.1.category", "6"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "application.1.id", "2"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "application.1.level", "medium"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "ips.0.critical_severity", "critical"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "ips.0.high_severity", "high"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "ips.0.info_severity", "disable"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "ips.0.low_severity", "low"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "ips.0.medium_severity", "medium"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "level.0.critical", "50"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "level.0.high", "30"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "level.0.low", "5"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "level.0.medium", "10"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "malware.0.botnet_connection", "critical"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "malware.0.command_blocked", "disable"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "malware.0.content_disarm", "medium"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "malware.0.file_blocked", "low"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "malware.0.mimefragmented", "disable"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "malware.0.oversized", "disable"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "malware.0.switch_proto", "disable"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "malware.0.virus_file_type_executable", "medium"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "malware.0.virus_infected", "critical"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "malware.0.virus_outbreak_prevention", "critical"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "malware.0.virus_scan_error", "high"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.0.category", "26"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.0.id", "1"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.0.level", "high"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.1.category", "61"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.1.id", "2"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.1.level", "high"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.2.category", "86"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.2.id", "3"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.2.level", "high"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.3.category", "1"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.3.id", "4"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.3.level", "medium"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.4.category", "3"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.4.id", "5"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.4.level", "medium"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.5.category", "4"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.5.id", "6"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.5.level", "medium"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.6.category", "5"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.6.id", "7"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.6.level", "medium"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.7.category", "6"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.7.id", "8"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.7.level", "medium"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.8.category", "12"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.8.id", "9"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.8.level", "medium"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.9.category", "59"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.9.id", "10"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.9.level", "medium"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.10.category", "62"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.10.id", "11"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.10.level", "medium"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.11.category", "83"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.11.id", "12"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.11.level", "medium"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.12.category", "72"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.12.id", "13"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.12.level", "low"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.13.category", "14"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.13.id", "14"),
					resource.TestCheckResourceAttr("fortios_log_threatweight.trname", "web.13.level", "low"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogThreatWeightExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogThreatWeight: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogThreatWeight is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogThreatWeight(i)

		if err != nil {
			return fmt.Errorf("Error reading LogThreatWeight: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogThreatWeight: %s", n)
		}

		return nil
	}
}

func testAccCheckLogThreatWeightDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_log_threatweight" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogThreatWeight(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogThreatWeight %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogThreatWeightConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_log_threatweight" "trname" {
  blocked_connection = "high"
  failed_connection  = "low"
  status             = "enable"
  url_block_detected = "high"

  application {
    category = 2
    id       = 1
    level    = "low"
  }
  application {
    category = 6
    id       = 2
    level    = "medium"
  }

  ips {
    critical_severity = "critical"
    high_severity     = "high"
    info_severity     = "disable"
    low_severity      = "low"
    medium_severity   = "medium"
  }

  level {
    critical = 50
    high     = 30
    low      = 5
    medium   = 10
  }

  malware {
    botnet_connection          = "critical"
    command_blocked            = "disable"
    content_disarm             = "medium"
    file_blocked               = "low"
    mimefragmented             = "disable"
    oversized                  = "disable"
    switch_proto               = "disable"
    virus_file_type_executable = "medium"
    virus_infected             = "critical"
    virus_outbreak_prevention  = "critical"
    virus_scan_error           = "high"
  }

  web {
    category = 26
    id       = 1
    level    = "high"
  }
  web {
    category = 61
    id       = 2
    level    = "high"
  }
  web {
    category = 86
    id       = 3
    level    = "high"
  }
  web {
    category = 1
    id       = 4
    level    = "medium"
  }
  web {
    category = 3
    id       = 5
    level    = "medium"
  }
  web {
    category = 4
    id       = 6
    level    = "medium"
  }
  web {
    category = 5
    id       = 7
    level    = "medium"
  }
  web {
    category = 6
    id       = 8
    level    = "medium"
  }
  web {
    category = 12
    id       = 9
    level    = "medium"
  }
  web {
    category = 59
    id       = 10
    level    = "medium"
  }
  web {
    category = 62
    id       = 11
    level    = "medium"
  }
  web {
    category = 83
    id       = 12
    level    = "medium"
  }
  web {
    category = 72
    id       = 13
    level    = "low"
  }
  web {
    category = 14
    id       = 14
    level    = "low"
  }
}

`)
}
