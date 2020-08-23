
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

func TestAccFortiOSSystemAccprofile_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemAccprofile_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemAccprofileConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemAccprofileExists("fortios_system_accprofile.test12"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "name", rname),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "admintimeout", "10"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "admintimeout_override", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "authgrp", "read-write"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "ftviewgrp", "read-write"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "fwgrp", "custom"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "loggrp", "read-write"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "netgrp", "read-write"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "scope", "vdom"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "secfabgrp", "read-write"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "sysgrp", "read-write"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "utmgrp", "custom"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "vpngrp", "read-write"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "wanoptgrp", "read-write"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "wifi", "read-write"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "fwgrp_permission.0.address", "read-write"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "fwgrp_permission.0.policy", "read-write"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "fwgrp_permission.0.schedule", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "fwgrp_permission.0.service", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "loggrp_permission.0.config", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "loggrp_permission.0.data_access", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "loggrp_permission.0.report_access", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "loggrp_permission.0.threat_weight", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "netgrp_permission.0.cfg", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "netgrp_permission.0.packet_capture", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "netgrp_permission.0.route_cfg", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "sysgrp_permission.0.admin", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "sysgrp_permission.0.cfg", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "sysgrp_permission.0.mnt", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "sysgrp_permission.0.upd", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "utmgrp_permission.0.antivirus", "read-write"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "utmgrp_permission.0.application_control", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "utmgrp_permission.0.data_loss_prevention", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "utmgrp_permission.0.dnsfilter", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "utmgrp_permission.0.endpoint_control", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "utmgrp_permission.0.icap", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "utmgrp_permission.0.ips", "read-write"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "utmgrp_permission.0.voip", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "utmgrp_permission.0.waf", "none"),
                    resource.TestCheckResourceAttr("fortios_system_accprofile.test12", "utmgrp_permission.0.webfilter", "none"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemAccprofileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemAccprofile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemAccprofile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemAccprofile(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemAccprofile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemAccprofile: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemAccprofileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_accprofile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAccprofile(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemAccprofile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemAccprofileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_accprofile" "test12" {
  name                  = "%[1]s"
  admintimeout          = 10
  admintimeout_override = "disable"
  authgrp               = "read-write"
  ftviewgrp             = "read-write"
  fwgrp                 = "custom"
  loggrp                = "read-write"
  netgrp                = "read-write"
  scope                 = "vdom"
  secfabgrp             = "read-write"
  sysgrp                = "read-write"
  utmgrp                = "custom"
  vpngrp                = "read-write"
  wanoptgrp             = "read-write"
  wifi                  = "read-write"

  fwgrp_permission {
    address  = "read-write"
    policy   = "read-write"
    schedule = "none"
    service  = "none"
  }

  loggrp_permission {
    config        = "none"
    data_access   = "none"
    report_access = "none"
    threat_weight = "none"
  }

  netgrp_permission {
    cfg            = "none"
    packet_capture = "none"
    route_cfg      = "none"
  }

  sysgrp_permission {
    admin = "none"
    cfg   = "none"
    mnt   = "none"
    upd   = "none"
  }

  utmgrp_permission {
    antivirus            = "read-write"
    application_control  = "none"
    data_loss_prevention = "none"
    dnsfilter            = "none"
    endpoint_control     = "none"
    icap                 = "none"
    ips                  = "read-write"
    voip                 = "none"
    waf                  = "none"
    webfilter            = "none"
  }
}



`, name)
}
