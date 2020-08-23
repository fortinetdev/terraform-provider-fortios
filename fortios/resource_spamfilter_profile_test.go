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

func TestAccFortiOSSpamfilterProfile_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSpamfilterProfile_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSpamfilterProfileConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSpamfilterProfileExists("fortios_spamfilter_profile.trname"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "comment", "terraform test"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "external", "disable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "flow_based", "disable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "spam_bwl_table", "0"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "spam_bword_table", "0"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "spam_bword_threshold", "10"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "spam_filtering", "disable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "spam_iptrust_table", "0"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "spam_log", "enable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "spam_log_fortiguard_response", "disable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "spam_mheader_table", "0"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "spam_rbl_table", "0"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "gmail.0.log", "disable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "imap.0.action", "tag"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "imap.0.log", "disable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "imap.0.tag_msg", "Spam"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "imap.0.tag_type", "subject spaminfo"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "mapi.0.action", "discard"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "mapi.0.log", "disable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "msn_hotmail.0.log", "disable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "pop3.0.action", "tag"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "pop3.0.log", "disable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "pop3.0.tag_msg", "Spam"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "pop3.0.tag_type", "subject spaminfo"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "smtp.0.action", "discard"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "smtp.0.hdrip", "disable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "smtp.0.local_override", "disable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "smtp.0.log", "disable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "smtp.0.tag_msg", "Spam"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "smtp.0.tag_type", "subject spaminfo"),
					resource.TestCheckResourceAttr("fortios_spamfilter_profile.trname", "yahoo_mail.0.log", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSpamfilterProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SpamfilterProfile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SpamfilterProfile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSpamfilterProfile(i)

		if err != nil {
			return fmt.Errorf("Error reading SpamfilterProfile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SpamfilterProfile: %s", n)
		}

		return nil
	}
}

func testAccCheckSpamfilterProfileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_spamfilter_profile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSpamfilterProfile(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SpamfilterProfile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSpamfilterProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_spamfilter_profile" "trname" {
  comment                      = "terraform test"
  external                     = "disable"
  flow_based                   = "disable"
  name                         = "%[1]s"
  spam_bwl_table               = 0
  spam_bword_table             = 0
  spam_bword_threshold         = 10
  spam_filtering               = "disable"
  spam_iptrust_table           = 0
  spam_log                     = "enable"
  spam_log_fortiguard_response = "disable"
  spam_mheader_table           = 0
  spam_rbl_table               = 0

  gmail {
    log = "disable"
  }

  imap {
    action   = "tag"
    log      = "disable"
    tag_msg  = "Spam"
    tag_type = "subject spaminfo"
  }

  mapi {
    action = "discard"
    log    = "disable"
  }

  msn_hotmail {
    log = "disable"
  }

  pop3 {
    action   = "tag"
    log      = "disable"
    tag_msg  = "Spam"
    tag_type = "subject spaminfo"
  }

  smtp {
    action         = "discard"
    hdrip          = "disable"
    local_override = "disable"
    log            = "disable"
    tag_msg        = "Spam"
    tag_type       = "subject spaminfo"
  }

  yahoo_mail {
    log = "disable"
  }
}


`, name)
}
