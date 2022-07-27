package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFortiOSFirewallObjectServiceCategory(t *testing.T) {
	rname := "fosfq_" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirewallObjectServiceCategoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallObjectServiceCategoryConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallObjectServiceCategoryExists("fortios_firewall_object_servicecategory.test"),
					resource.TestCheckResourceAttr("fortios_firewall_object_servicecategory.test", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_object_servicecategory.test", "comment", "Terraform Test"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallObjectServiceCategoryExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Firewall Object Service Category: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Firewall Object Service Category is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectServiceCategory(i)

		if err != nil {
			return fmt.Errorf("Error reading Firewall Object Service: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Firewall Object Service: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallObjectServiceCategoryDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_object_servicecategory" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallObjectServiceCategory(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Firewall Object Service Category %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallObjectServiceCategoryConfig(name string) string {
	return fmt.Sprintf(`
		resource "fortios_firewall_object_servicecategory" "test" {
		name = "%s"
		comment = "Terraform Test"
		}
		`, name)
}
