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

func TestAccFortiOSIcapProfile_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSIcapProfile_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSIcapProfileConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSIcapProfileExists("fortios_icap_profile.trname"),
					resource.TestCheckResourceAttr("fortios_icap_profile.trname", "methods", "delete get head options post put trace other"),
					resource.TestCheckResourceAttr("fortios_icap_profile.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_icap_profile.trname", "request", "disable"),
					resource.TestCheckResourceAttr("fortios_icap_profile.trname", "request_failure", "error"),
					resource.TestCheckResourceAttr("fortios_icap_profile.trname", "response", "disable"),
					resource.TestCheckResourceAttr("fortios_icap_profile.trname", "response_failure", "error"),
					resource.TestCheckResourceAttr("fortios_icap_profile.trname", "response_req_hdr", "disable"),
					resource.TestCheckResourceAttr("fortios_icap_profile.trname", "streaming_content_bypass", "disable"),
					resource.TestCheckResourceAttr("fortios_icap_profile.trname", "icap_headers.0.base64_encoding", "disable"),
					resource.TestCheckResourceAttr("fortios_icap_profile.trname", "icap_headers.0.content", "$user"),
					resource.TestCheckResourceAttr("fortios_icap_profile.trname", "icap_headers.0.name", "X-Authenticated-User"),
				),
			},
		},
	})
}

func testAccCheckFortiOSIcapProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found IcapProfile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No IcapProfile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadIcapProfile(i)

		if err != nil {
			return fmt.Errorf("Error reading IcapProfile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating IcapProfile: %s", n)
		}

		return nil
	}
}

func testAccCheckIcapProfileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_icap_profile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadIcapProfile(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error IcapProfile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSIcapProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_icap_profile" "trname" {
  methods                  = "delete get head options post put trace other"
  name                     = "%[1]s"
  request                  = "disable"
  request_failure          = "error"
  response                 = "disable"
  response_failure         = "error"
  response_req_hdr         = "disable"
  streaming_content_bypass = "disable"

  icap_headers {
    base64_encoding = "disable"
    content         = "$user"
    name            = "X-Authenticated-User"
  }
}
`, name)
}
