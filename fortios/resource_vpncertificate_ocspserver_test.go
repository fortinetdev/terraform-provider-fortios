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

func TestAccFortiOSVpnCertificateOcspServer_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSVpnCertificateOcspServer_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVpnCertificateOcspServerConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVpnCertificateOcspServerExists("fortios_vpncertificate_ocspserver.trname"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_ocspserver.trname", "cert", "ACCVRAIZ1"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_ocspserver.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_vpncertificate_ocspserver.trname", "source_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_ocspserver.trname", "unavail_action", "revoke"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_ocspserver.trname", "url", "www.tetserv.com"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVpnCertificateOcspServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnCertificateOcspServer: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnCertificateOcspServer is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnCertificateOcspServer(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnCertificateOcspServer: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnCertificateOcspServer: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnCertificateOcspServerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpncertificate_ocspserver" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnCertificateOcspServer(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnCertificateOcspServer %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnCertificateOcspServerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpncertificate_ocspserver" "trname" {
  cert           = "ACCVRAIZ1"
  name           = "%[1]s"
  source_ip      = "0.0.0.0"
  unavail_action = "revoke"
  url            = "www.tetserv.com"
}
`, name)
}
