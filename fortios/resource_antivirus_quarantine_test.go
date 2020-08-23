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

func TestAccFortiOSAntivirusQuarantine_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSAntivirusQuarantine_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSAntivirusQuarantineConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSAntivirusQuarantineExists("fortios_antivirus_quarantine.trname"),
					resource.TestCheckResourceAttr("fortios_antivirus_quarantine.trname", "agelimit", "0"),
					resource.TestCheckResourceAttr("fortios_antivirus_quarantine.trname", "destination", "disk"),
					resource.TestCheckResourceAttr("fortios_antivirus_quarantine.trname", "lowspace", "ovrw-old"),
					resource.TestCheckResourceAttr("fortios_antivirus_quarantine.trname", "maxfilesize", "0"),
					resource.TestCheckResourceAttr("fortios_antivirus_quarantine.trname", "quarantine_quota", "0"),
					resource.TestCheckResourceAttr("fortios_antivirus_quarantine.trname", "store_blocked", "imap smtp pop3 http ftp nntp imaps smtps pop3s ftps mapi cifs"),
					resource.TestCheckResourceAttr("fortios_antivirus_quarantine.trname", "store_heuristic", "imap smtp pop3 http ftp nntp imaps smtps pop3s https ftps mapi cifs"),
					resource.TestCheckResourceAttr("fortios_antivirus_quarantine.trname", "store_infected", "imap smtp pop3 http ftp nntp imaps smtps pop3s https ftps mapi cifs"),
				),
			},
		},
	})
}

func testAccCheckFortiOSAntivirusQuarantineExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found AntivirusQuarantine: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AntivirusQuarantine is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadAntivirusQuarantine(i)

		if err != nil {
			return fmt.Errorf("Error reading AntivirusQuarantine: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating AntivirusQuarantine: %s", n)
		}

		return nil
	}
}

func testAccCheckAntivirusQuarantineDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_antivirus_quarantine" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadAntivirusQuarantine(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error AntivirusQuarantine %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSAntivirusQuarantineConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_antivirus_quarantine" "trname" {
  agelimit         = 0
  destination      = "disk"
  lowspace         = "ovrw-old"
  maxfilesize      = 0
  quarantine_quota = 0
  store_blocked    = "imap smtp pop3 http ftp nntp imaps smtps pop3s ftps mapi cifs"
  store_heuristic  = "imap smtp pop3 http ftp nntp imaps smtps pop3s https ftps mapi cifs"
  store_infected   = "imap smtp pop3 http ftp nntp imaps smtps pop3s https ftps mapi cifs"
}

`)
}
