package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiManagerObjectAdomRevision(t *testing.T) {
	name := "fmg-object-adom-revision" + acctest.RandString(12)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGObjectAdomRevisionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerObjectAdomRevisionConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerObjectAdomRevisionExists("fortios_fmg_object_adom_revision.test1"),
					resource.TestCheckResourceAttr("fortios_fmg_object_adom_revision.test1", "name", name),
					resource.TestCheckResourceAttr("fortios_fmg_object_adom_revision.test1", "description", "adom revision"),
					resource.TestCheckResourceAttr("fortios_fmg_object_adom_revision.test1", "created_by", "fortinet"),
					resource.TestCheckResourceAttr("fortios_fmg_object_adom_revision.test1", "locked", "0"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerObjectAdomRevisionExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Object Adom Revision: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Object Adom Revision is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		i := rs.Primary.ID
		o, err := c.ReadObjectAdomRevision("root", i)

		if err != nil {
			return fmt.Errorf("Error reading Object Adom Revision: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Object Adom Revision: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGObjectAdomRevisionDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_fmg_object_adom_revision" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadObjectAdomRevision("root", i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Object Adom Revision %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiManagerObjectAdomRevisionConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_fmg_object_adom_revision" "test1" {
	name = "%s"
    description = "adom revision"
    created_by = "fortinet"
    locked = 0
}
`, name)
}
