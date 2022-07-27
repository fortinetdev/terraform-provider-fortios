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

func TestAccFortiOSLogDiskSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogDiskSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogDiskSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogDiskSettingExists("fortios_logdisk_setting.trname"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "diskfull", "overwrite"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "dlp_archive_quota", "0"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "full_final_warning_threshold", "95"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "full_first_warning_threshold", "75"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "full_second_warning_threshold", "90"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "ips_archive", "enable"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "log_quota", "0"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "max_log_file_size", "20"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "max_policy_packet_capture_size", "100"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "maximum_log_age", "7"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "report_quota", "0"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "roll_day", "sunday"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "roll_schedule", "daily"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "roll_time", "00:00"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "source_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "upload", "disable"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "upload_delete_files", "enable"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "upload_destination", "ftp-server"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "upload_ssl_conn", "default"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "uploadip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "uploadport", "21"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "uploadsched", "disable"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "uploadtime", "00:00"),
					resource.TestCheckResourceAttr("fortios_logdisk_setting.trname", "uploadtype", "traffic event virus webfilter IPS spamfilter dlp-archive anomaly voip dlp app-ctrl waf netscan gtp dns"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogDiskSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogDiskSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogDiskSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogDiskSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading LogDiskSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogDiskSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogDiskSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logdisk_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogDiskSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogDiskSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogDiskSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logdisk_setting" "trname" {
  diskfull                       = "overwrite"
  dlp_archive_quota              = 0
  full_final_warning_threshold   = 95
  full_first_warning_threshold   = 75
  full_second_warning_threshold  = 90
  ips_archive                    = "enable"
  log_quota                      = 0
  max_log_file_size              = 20
  max_policy_packet_capture_size = 100
  maximum_log_age                = 7
  report_quota                   = 0
  roll_day                       = "sunday"
  roll_schedule                  = "daily"
  roll_time                      = "00:00"
  source_ip                      = "0.0.0.0"
  status                         = "enable"
  upload                         = "disable"
  upload_delete_files            = "enable"
  upload_destination             = "ftp-server"
  upload_ssl_conn                = "default"
  uploadip                       = "0.0.0.0"
  uploadport                     = 21
  uploadsched                    = "disable"
  uploadtime                     = "00:00"
  uploadtype                     = "traffic event virus webfilter IPS spamfilter dlp-archive anomaly voip dlp app-ctrl waf netscan gtp dns"
}
`)
}
