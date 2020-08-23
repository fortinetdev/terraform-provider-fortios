
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSDlpFpDocSource_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSDlpFpDocSource_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSDlpFpDocSourceConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSDlpFpDocSourceExists("fortios_dlp_fpdocsource.trname"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "date", "1"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "file_path", "/"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "file_pattern", "*"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "keep_modified", "enable"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "period", "none"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "remove_deleted", "enable"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "scan_on_creation", "enable"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "scan_subdirectories", "enable"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "server", "1.1.1.1"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "server_type", "samba"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "tod_hour", "1"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "tod_min", "0"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "username", "sgh"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "vdom", "mgmt"),
                    resource.TestCheckResourceAttr("fortios_dlp_fpdocsource.trname", "weekday", "sunday"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSDlpFpDocSourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found DlpFpDocSource: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No DlpFpDocSource is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadDlpFpDocSource(i)

		if err != nil {
			return fmt.Errorf("Error reading DlpFpDocSource: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating DlpFpDocSource: %s", n)
		}

		return nil
	}
}

func testAccCheckDlpFpDocSourceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_dlp_fpdocsource" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadDlpFpDocSource(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error DlpFpDocSource %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSDlpFpDocSourceConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_dlp_fpdocsource" "trname" {
  date                = 1
  file_path           = "/"
  file_pattern        = "*"
  keep_modified       = "enable"
  name                = "%[1]s"
  period              = "none"
  remove_deleted      = "enable"
  scan_on_creation    = "enable"
  scan_subdirectories = "enable"
  server              = "1.1.1.1"
  server_type         = "samba"
  tod_hour            = 1
  tod_min             = 0
  username            = "sgh"
  vdom                = "mgmt"
  weekday             = "sunday"
}
`, name)
}
