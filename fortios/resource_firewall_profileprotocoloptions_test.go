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

func TestAccFortiOSFirewallProfileProtocolOptions_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallProfileProtocolOptions_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallProfileProtocolOptionsConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallProfileProtocolOptionsExists("fortios_firewall_profileprotocoloptions.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "oversize_log", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "rpc_over_http", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "switching_protocols_log", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "dns.0.ports", "53"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "dns.0.status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "ftp.0.comfort_amount", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "ftp.0.comfort_interval", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "ftp.0.inspect_all", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "ftp.0.options", "splice"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "ftp.0.oversize_limit", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "ftp.0.ports", "21"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "ftp.0.scan_bzip2", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "ftp.0.status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "ftp.0.uncompressed_nest_limit", "12"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "ftp.0.uncompressed_oversize_limit", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.block_page_status_code", "403"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.comfort_amount", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.comfort_interval", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.fortinet_bar", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.fortinet_bar_port", "8011"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.http_policy", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.inspect_all", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.oversize_limit", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.ports", "80"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.range_block", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.retry_count", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.scan_bzip2", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.streaming_content_bypass", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.strip_x_forwarded_for", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.switching_protocols", "bypass"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.uncompressed_nest_limit", "12"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "http.0.uncompressed_oversize_limit", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "imap.0.inspect_all", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "imap.0.options", "fragmail"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "imap.0.oversize_limit", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "imap.0.ports", "143"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "imap.0.scan_bzip2", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "imap.0.status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "imap.0.uncompressed_nest_limit", "12"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "imap.0.uncompressed_oversize_limit", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "mail_signature.0.status", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "mapi.0.options", "fragmail"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "mapi.0.oversize_limit", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "mapi.0.ports", "135"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "mapi.0.scan_bzip2", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "mapi.0.status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "mapi.0.uncompressed_nest_limit", "12"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "mapi.0.uncompressed_oversize_limit", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "nntp.0.inspect_all", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "nntp.0.options", "splice"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "nntp.0.oversize_limit", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "nntp.0.ports", "119"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "nntp.0.scan_bzip2", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "nntp.0.status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "nntp.0.uncompressed_nest_limit", "12"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "nntp.0.uncompressed_oversize_limit", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "pop3.0.inspect_all", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "pop3.0.options", "fragmail"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "pop3.0.oversize_limit", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "pop3.0.ports", "110"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "pop3.0.scan_bzip2", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "pop3.0.status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "pop3.0.uncompressed_nest_limit", "12"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "pop3.0.uncompressed_oversize_limit", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "smtp.0.inspect_all", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "smtp.0.options", "fragmail splice"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "smtp.0.oversize_limit", "10"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "smtp.0.ports", "25"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "smtp.0.scan_bzip2", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "smtp.0.server_busy", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "smtp.0.status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "smtp.0.uncompressed_nest_limit", "12"),
					resource.TestCheckResourceAttr("fortios_firewall_profileprotocoloptions.trname", "smtp.0.uncompressed_oversize_limit", "10"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallProfileProtocolOptionsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallProfileProtocolOptions: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallProfileProtocolOptions is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallProfileProtocolOptions(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallProfileProtocolOptions: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallProfileProtocolOptions: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallProfileProtocolOptionsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_profileprotocoloptions" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallProfileProtocolOptions(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallProfileProtocolOptions %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallProfileProtocolOptionsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_profileprotocoloptions" "trname" {
  name                    = "%[1]s"
  oversize_log            = "disable"
  rpc_over_http           = "disable"
  switching_protocols_log = "disable"

  dns {
    ports  = 53
    status = "enable"
  }

  ftp {
    comfort_amount              = 1
    comfort_interval            = 10
    inspect_all                 = "disable"
    options                     = "splice"
    oversize_limit              = 10
    ports                       = 21
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  http {
    block_page_status_code      = 403
    comfort_amount              = 1
    comfort_interval            = 10
    fortinet_bar                = "disable"
    fortinet_bar_port           = 8011
    http_policy                 = "disable"
    inspect_all                 = "disable"
    oversize_limit              = 10
    ports                       = 80
    range_block                 = "disable"
    retry_count                 = 0
    scan_bzip2                  = "enable"
    status                      = "enable"
    streaming_content_bypass    = "enable"
    strip_x_forwarded_for       = "disable"
    switching_protocols         = "bypass"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  imap {
    inspect_all                 = "disable"
    options                     = "fragmail"
    oversize_limit              = 10
    ports                       = 143
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  mail_signature {
    status = "disable"
  }

  mapi {
    options                     = "fragmail"
    oversize_limit              = 10
    ports                       = 135
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  nntp {
    inspect_all                 = "disable"
    options                     = "splice"
    oversize_limit              = 10
    ports                       = 119
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  pop3 {
    inspect_all                 = "disable"
    options                     = "fragmail"
    oversize_limit              = 10
    ports                       = 110
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  smtp {
    inspect_all                 = "disable"
    options                     = "fragmail splice"
    oversize_limit              = 10
    ports                       = 25
    scan_bzip2                  = "enable"
    server_busy                 = "disable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }
}
`, name)
}
