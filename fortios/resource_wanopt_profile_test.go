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

func TestAccFortiOSWanoptProfile_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWanoptProfile_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWanoptProfileConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWanoptProfileExists("fortios_wanopt_profile.trname"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "comments", "test"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "transparent", "enable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "cifs.0.byte_caching", "enable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "cifs.0.log_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "cifs.0.port", "445"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "cifs.0.prefer_chunking", "fix"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "cifs.0.secure_tunnel", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "cifs.0.status", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "cifs.0.tunnel_sharing", "private"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "ftp.0.byte_caching", "enable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "ftp.0.log_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "ftp.0.port", "21"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "ftp.0.prefer_chunking", "fix"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "ftp.0.secure_tunnel", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "ftp.0.status", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "ftp.0.tunnel_sharing", "private"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "http.0.byte_caching", "enable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "http.0.log_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "http.0.port", "80"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "http.0.prefer_chunking", "fix"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "http.0.secure_tunnel", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "http.0.ssl", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "http.0.ssl_port", "443"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "http.0.status", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "http.0.tunnel_non_http", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "http.0.tunnel_sharing", "private"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "http.0.unknown_http_version", "tunnel"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "mapi.0.byte_caching", "enable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "mapi.0.log_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "mapi.0.port", "135"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "mapi.0.secure_tunnel", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "mapi.0.status", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "mapi.0.tunnel_sharing", "private"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "tcp.0.byte_caching", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "tcp.0.byte_caching_opt", "mem-only"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "tcp.0.log_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "tcp.0.port", "1-65535"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "tcp.0.secure_tunnel", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "tcp.0.ssl", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "tcp.0.ssl_port", "443"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "tcp.0.status", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_profile.trname", "tcp.0.tunnel_sharing", "private"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWanoptProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WanoptProfile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WanoptProfile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWanoptProfile(i)

		if err != nil {
			return fmt.Errorf("Error reading WanoptProfile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WanoptProfile: %s", n)
		}

		return nil
	}
}

func testAccCheckWanoptProfileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wanopt_profile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWanoptProfile(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WanoptProfile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWanoptProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wanopt_profile" "trname" {
  comments    = "test"
  name        = "%[1]s"
  transparent = "enable"

  cifs {
    byte_caching    = "enable"
    log_traffic     = "enable"
    port            = 445
    prefer_chunking = "fix"
    secure_tunnel   = "disable"
    status          = "disable"
    tunnel_sharing  = "private"
  }

  ftp {
    byte_caching    = "enable"
    log_traffic     = "enable"
    port            = 21
    prefer_chunking = "fix"
    secure_tunnel   = "disable"
    status          = "disable"
    tunnel_sharing  = "private"
  }

  http {
    byte_caching         = "enable"
    log_traffic          = "enable"
    port                 = 80
    prefer_chunking      = "fix"
    secure_tunnel        = "disable"
    ssl                  = "disable"
    ssl_port             = 443
    status               = "disable"
    tunnel_non_http      = "disable"
    tunnel_sharing       = "private"
    unknown_http_version = "tunnel"
  }

  mapi {
    byte_caching   = "enable"
    log_traffic    = "enable"
    port           = 135
    secure_tunnel  = "disable"
    status         = "disable"
    tunnel_sharing = "private"
  }

  tcp {
    byte_caching     = "disable"
    byte_caching_opt = "mem-only"
    log_traffic      = "enable"
    port             = "1-65535"
    secure_tunnel    = "disable"
    ssl              = "disable"
    ssl_port         = 443
    status           = "disable"
    tunnel_sharing   = "private"
  }
}
`, name)
}
