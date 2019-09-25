package fortios

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"fortios": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("FORTIOS_ACCESS_HOSTNAME"); v == "" {
		t.Fatal("FORTIOS_ACCESS_HOSTNAME must be set for acceptance tests")
	}
	if v := os.Getenv("FORTIOS_ACCESS_TOKEN"); v == "" {
		t.Fatal("FORTIOS_ACCESS_TOKEN must be set for acceptance tests")
	}
	if v := os.Getenv("FORTIOS_INSECURE"); v == "" {
		t.Fatal("FORTIOS_INSECURE must be set for acceptance tests")
	}
}
