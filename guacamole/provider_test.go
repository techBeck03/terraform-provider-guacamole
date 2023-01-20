package guacamole

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"guacamole": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("GUACAMOLE_URL"); err == "" {
		t.Fatal("GUACAMOLE_URL must be set for acceptance tests")
	}
	if err := os.Getenv("GUACAMOLE_PASSWORD"); err == "" {
		if os.Getenv("GUACAMOLE_TOKEN") == "" {
			t.Fatal("GUACAMOLE_PASSWORD or GUACAMOLE_TOKEN must be set for acceptance tests")
		}
	} else {
		if os.Getenv("GUACAMOLE_USERNAME") == "" {
			t.Fatal("GUACAMOLE_USERNAME must be set for acceptance tests when GUACAMOLE_PASSWORD is set")
		}
	}
	if err := os.Getenv("GUACAMOLE_TOKEN"); err != "" {
		if os.Getenv("GUACAMOLE_DATA_SOURCE") == "" {
			t.Fatal("GUACAMOLE_DATA_SOURCE must be set for acceptance tests when GUACAMOLE_TOKEN is provided")
		}
	}
}
