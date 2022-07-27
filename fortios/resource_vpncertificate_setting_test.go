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

func TestAccFortiOSVpnCertificateSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSVpnCertificateSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVpnCertificateSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVpnCertificateSettingExists("fortios_vpncertificate_setting.trname"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "certname_dsa1024", "Fortinet_SSL_DSA1024"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "certname_dsa2048", "Fortinet_SSL_DSA2048"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "certname_ecdsa256", "Fortinet_SSL_ECDSA256"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "certname_ecdsa384", "Fortinet_SSL_ECDSA384"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "certname_rsa1024", "Fortinet_SSL_RSA1024"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "certname_rsa2048", "Fortinet_SSL_RSA2048"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "check_ca_cert", "enable"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "check_ca_chain", "disable"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "cmp_save_extra_certs", "disable"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "cn_match", "substring"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "ocsp_option", "server"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "ocsp_status", "disable"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "ssl_min_proto_version", "default"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "strict_crl_check", "disable"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "strict_ocsp_check", "disable"),
					resource.TestCheckResourceAttr("fortios_vpncertificate_setting.trname", "subject_match", "substring"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVpnCertificateSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnCertificateSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnCertificateSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnCertificateSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnCertificateSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnCertificateSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnCertificateSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpncertificate_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnCertificateSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnCertificateSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnCertificateSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpncertificate_setting" "trname" {
  certname_dsa1024      = "Fortinet_SSL_DSA1024"
  certname_dsa2048      = "Fortinet_SSL_DSA2048"
  certname_ecdsa256     = "Fortinet_SSL_ECDSA256"
  certname_ecdsa384     = "Fortinet_SSL_ECDSA384"
  certname_rsa1024      = "Fortinet_SSL_RSA1024"
  certname_rsa2048      = "Fortinet_SSL_RSA2048"
  check_ca_cert         = "enable"
  check_ca_chain        = "disable"
  cmp_save_extra_certs  = "disable"
  cn_match              = "substring"
  ocsp_option           = "server"
  ocsp_status           = "disable"
  ssl_min_proto_version = "default"
  strict_crl_check      = "disable"
  strict_ocsp_check     = "disable"
  subject_match         = "substring"
}
`)
}
