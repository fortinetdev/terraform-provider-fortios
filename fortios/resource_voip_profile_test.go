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

func TestAccFortiOSVoipProfile_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSVoipProfile_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVoipProfileConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVoipProfileExists("fortios_voip_profile.trname"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "comment", "test"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sccp.0.block_mcast", "disable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sccp.0.log_call_summary", "disable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sccp.0.log_violations", "disable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sccp.0.max_calls", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sccp.0.status", "enable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sccp.0.verify_header", "disable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.ack_rate", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.bye_rate", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.call_keepalive", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.cancel_rate", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.contact_fixup", "enable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.hnt_restrict_source_ip", "disable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.hosted_nat_traversal", "disable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.info_rate", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.invite_rate", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.ips_rtp", "enable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.log_call_summary", "enable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.log_violations", "disable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.max_body_length", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.max_dialogs", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.max_idle_dialogs", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.max_line_length", "998"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.message_rate", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.nat_trace", "enable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.no_sdp_fixup", "disable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.notify_rate", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.open_contact_pinhole", "enable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.open_record_route_pinhole", "enable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.open_register_pinhole", "enable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.open_via_pinhole", "disable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.options_rate", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.prack_rate", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.preserve_override", "disable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.provisional_invite_expiry_time", "210"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.publish_rate", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.refer_rate", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.register_contact_trace", "disable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.register_rate", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.rfc2543_branch", "disable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.rtp", "enable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.ssl_algorithm", "high"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.ssl_client_renegotiation", "allow"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.ssl_max_version", "tls-1.2"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.ssl_min_version", "tls-1.1"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.ssl_mode", "off"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.ssl_pfs", "allow"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.ssl_send_empty_frags", "enable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.status", "enable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.strict_register", "enable"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.subscribe_rate", "0"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.unknown_header", "pass"),
					resource.TestCheckResourceAttr("fortios_voip_profile.trname", "sip.0.update_rate", "0"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVoipProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VoipProfile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VoipProfile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVoipProfile(i)

		if err != nil {
			return fmt.Errorf("Error reading VoipProfile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VoipProfile: %s", n)
		}

		return nil
	}
}

func testAccCheckVoipProfileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_voip_profile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVoipProfile(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VoipProfile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVoipProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_voip_profile" "trname" {
  comment = "test"
  name    = "%[1]s"

  sccp {
    block_mcast      = "disable"
    log_call_summary = "disable"
    log_violations   = "disable"
    max_calls        = 0
    status           = "enable"
    verify_header    = "disable"
  }

  sip {
    ack_rate                       = 0
    bye_rate                       = 0
    call_keepalive                 = 0
    cancel_rate                    = 0
    contact_fixup                  = "enable"
    hnt_restrict_source_ip         = "disable"
    hosted_nat_traversal           = "disable"
    info_rate                      = 0
    invite_rate                    = 0
    ips_rtp                        = "enable"
    log_call_summary               = "enable"
    log_violations                 = "disable"
    max_body_length                = 0
    max_dialogs                    = 0
    max_idle_dialogs               = 0
    max_line_length                = 998
    message_rate                   = 0
    nat_trace                      = "enable"
    no_sdp_fixup                   = "disable"
    notify_rate                    = 0
    open_contact_pinhole           = "enable"
    open_record_route_pinhole      = "enable"
    open_register_pinhole          = "enable"
    open_via_pinhole               = "disable"
    options_rate                   = 0
    prack_rate                     = 0
    preserve_override              = "disable"
    provisional_invite_expiry_time = 210
    publish_rate                   = 0
    refer_rate                     = 0
    register_contact_trace         = "disable"
    register_rate                  = 0
    rfc2543_branch                 = "disable"
    rtp                            = "enable"
    ssl_algorithm                  = "high"
    ssl_client_renegotiation       = "allow"
    ssl_max_version                = "tls-1.2"
    ssl_min_version                = "tls-1.1"
    ssl_mode                       = "off"
    ssl_pfs                        = "allow"
    ssl_send_empty_frags           = "enable"
    status                         = "enable"
    strict_register                = "enable"
    subscribe_rate                 = 0
    unknown_header                 = "pass"
    update_rate                    = 0
  }
}
`, name)
}
