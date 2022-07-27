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

func TestAccFortiOSEndpointControlProfile_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSEndpointControlProfile_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSEndpointControlProfileConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSEndpointControlProfileExists("fortios_endpointcontrol_profile.trname"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "profile_name", "1"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "device_groups.0.name", "Mobile Devices"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_android_settings.0.disable_wf_when_protected", "enable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_android_settings.0.forticlient_advanced_vpn", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_android_settings.0.forticlient_vpn_provisioning", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_android_settings.0.forticlient_wf", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_ios_settings.0.client_vpn_provisioning", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_ios_settings.0.disable_wf_when_protected", "enable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_ios_settings.0.distribute_configuration_profile", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_ios_settings.0.forticlient_wf", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.av_realtime_protection", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.av_signature_up_to_date", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_application_firewall", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_av", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_ems_compliance", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_ems_compliance_action", "warning"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_linux_ver", "5.4.1"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_log_upload", "enable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_log_upload_level", "traffic vulnerability event"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_mac_ver", "5.4.1"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_minimum_software_version", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_registration_compliance_action", "warning"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_security_posture", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_security_posture_compliance_action", "warning"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_system_compliance", "enable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_system_compliance_action", "warning"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_vuln_scan", "enable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_vuln_scan_compliance_action", "warning"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_vuln_scan_enforce", "high"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_vuln_scan_enforce_grace", "1"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_vuln_scan_exempt", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_wf", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.forticlient_win_ver", "5.4.1"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.os_av_software_installed", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "forticlient_winmac_settings.0.sandbox_analysis", "disable"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "on_net_addr.0.name", "all"),
					resource.TestCheckResourceAttr("fortios_endpointcontrol_profile.trname", "users.0.name", "guest"),
				),
			},
		},
	})
}

func testAccCheckFortiOSEndpointControlProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found EndpointControlProfile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No EndpointControlProfile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadEndpointControlProfile(i)

		if err != nil {
			return fmt.Errorf("Error reading EndpointControlProfile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating EndpointControlProfile: %s", n)
		}

		return nil
	}
}

func testAccCheckEndpointControlProfileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_endpointcontrol_profile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadEndpointControlProfile(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error EndpointControlProfile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSEndpointControlProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_endpointcontrol_profile" "trname" {
  profile_name = "1"

  device_groups {
    name = "Mobile Devices"
  }

  forticlient_android_settings {
    disable_wf_when_protected    = "enable"
    forticlient_advanced_vpn     = "disable"
    forticlient_vpn_provisioning = "disable"
    forticlient_wf               = "disable"
  }

  forticlient_ios_settings {
    client_vpn_provisioning          = "disable"
    disable_wf_when_protected        = "enable"
    distribute_configuration_profile = "disable"
    forticlient_wf                   = "disable"
  }

  forticlient_winmac_settings {
    av_realtime_protection                         = "disable"
    av_signature_up_to_date                        = "disable"
    forticlient_application_firewall               = "disable"
    forticlient_av                                 = "disable"
    forticlient_ems_compliance                     = "disable"
    forticlient_ems_compliance_action              = "warning"
    forticlient_linux_ver                          = "5.4.1"
    forticlient_log_upload                         = "enable"
    forticlient_log_upload_level                   = "traffic vulnerability event"
    forticlient_mac_ver                            = "5.4.1"
    forticlient_minimum_software_version           = "disable"
    forticlient_registration_compliance_action     = "warning"
    forticlient_security_posture                   = "disable"
    forticlient_security_posture_compliance_action = "warning"
    forticlient_system_compliance                  = "enable"
    forticlient_system_compliance_action           = "warning"
    forticlient_vuln_scan                          = "enable"
    forticlient_vuln_scan_compliance_action        = "warning"
    forticlient_vuln_scan_enforce                  = "high"
    forticlient_vuln_scan_enforce_grace            = 1
    forticlient_vuln_scan_exempt                   = "disable"
    forticlient_wf                                 = "disable"
    forticlient_win_ver                            = "5.4.1"
    os_av_software_installed                       = "disable"
    sandbox_analysis                               = "disable"
  }

  on_net_addr {
    name = "all"
  }

  users {
    name = "guest"
  }
}
`)
}
