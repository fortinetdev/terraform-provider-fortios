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

func TestAccFortiOSSystemApiUser_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSSystemApiUser_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemApiUserConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemApiUserExists("fortios_system_apiuser.test2"),
					resource.TestCheckResourceAttr("fortios_system_apiuser.test2", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_apiuser.test2", "accprofile", var0),
					resource.TestCheckResourceAttr("fortios_system_apiuser.test2", "vdom.0.name", "root"),
					resource.TestCheckResourceAttr("fortios_system_apiuser.test2", "trusthost.0.type", "ipv4-trusthost"),
					resource.TestCheckResourceAttr("fortios_system_apiuser.test2", "trusthost.0.ipv4_trusthost", "2.0.0.0/23"),
					resource.TestCheckResourceAttr("fortios_system_apiuser.test2", "trusthost.1.type", "ipv4-trusthost"),
					resource.TestCheckResourceAttr("fortios_system_apiuser.test2", "trusthost.1.ipv4_trusthost", "2.0.0.0 255.255.255.0"),
					resource.TestCheckResourceAttr("fortios_system_apiuser.test2", "trusthost.2.type", "ipv6-trusthost"),
					resource.TestCheckResourceAttr("fortios_system_apiuser.test2", "trusthost.2.ipv6_trusthost", "101:101:ffff:ffff::/0"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemApiUserExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemApiUser: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemApiUser is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemApiUser(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemApiUser: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemApiUser: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemApiUserDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_apiuser" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemApiUser(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemApiUser %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemApiUserConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_accprofile" "test2" {
  name                  = "var0%[1]s"
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

resource "fortios_system_apiuser" "test2" {
  name       = "%[1]s"
  accprofile = fortios_system_accprofile.test2.name
  vdom {
    name = "root"
  }

  trusthost {
    type           = "ipv4-trusthost"
    ipv4_trusthost = "2.0.0.0/23"
  }

  trusthost {
    type           = "ipv4-trusthost"
    ipv4_trusthost = "2.0.0.0 255.255.255.0"
  }

  trusthost {
    type           = "ipv6-trusthost"
    ipv6_trusthost = "101:101:ffff:ffff::/0"
  }
}
`, name)
}
