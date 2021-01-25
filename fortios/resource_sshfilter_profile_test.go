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

func TestAccFortiOSSshFilterProfile_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSshFilterProfile_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSshFilterProfileConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSshFilterProfileExists("fortios_sshfilter_profile.trname"),
					resource.TestCheckResourceAttr("fortios_sshfilter_profile.trname", "block", "x11"),
					resource.TestCheckResourceAttr("fortios_sshfilter_profile.trname", "default_command_log", "enable"),
					resource.TestCheckResourceAttr("fortios_sshfilter_profile.trname", "log", "x11"),
					resource.TestCheckResourceAttr("fortios_sshfilter_profile.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSSshFilterProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SshFilterProfile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SshFilterProfile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSshFilterProfile(i)

		if err != nil {
			return fmt.Errorf("Error reading SshFilterProfile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SshFilterProfile: %s", n)
		}

		return nil
	}
}

func testAccCheckSshFilterProfileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_sshfilter_profile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSshFilterProfile(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SshFilterProfile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSshFilterProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_sshfilter_profile" "trname" {
  block               = "x11"
  default_command_log = "enable"
  log                 = "x11"
  name                = "%[1]s"
}
`, name)
}
