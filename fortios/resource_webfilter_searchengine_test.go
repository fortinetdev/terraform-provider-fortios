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

func TestAccFortiOSWebfilterSearchEngine_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebfilterSearchEngine_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebfilterSearchEngineConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebfilterSearchEngineExists("fortios_webfilter_searchengine.trname"),
					resource.TestCheckResourceAttr("fortios_webfilter_searchengine.trname", "charset", "utf-8"),
					resource.TestCheckResourceAttr("fortios_webfilter_searchengine.trname", "hostname", "sg.eiwuc.com"),
					resource.TestCheckResourceAttr("fortios_webfilter_searchengine.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_webfilter_searchengine.trname", "query", "sc="),
					resource.TestCheckResourceAttr("fortios_webfilter_searchengine.trname", "safesearch", "disable"),
					resource.TestCheckResourceAttr("fortios_webfilter_searchengine.trname", "url", "^\\/f"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebfilterSearchEngineExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebfilterSearchEngine: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebfilterSearchEngine is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebfilterSearchEngine(i)

		if err != nil {
			return fmt.Errorf("Error reading WebfilterSearchEngine: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebfilterSearchEngine: %s", n)
		}

		return nil
	}
}

func testAccCheckWebfilterSearchEngineDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webfilter_searchengine" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebfilterSearchEngine(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebfilterSearchEngine %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebfilterSearchEngineConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webfilter_searchengine" "trname" {
  charset    = "utf-8"
  hostname   = "sg.eiwuc.com"
  name       = "%[1]s"
  query      = "sc="
  safesearch = "disable"
  url        = "^\\/f"
}
`, name)
}
