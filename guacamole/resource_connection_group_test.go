package guacamole

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testProviderConnectionGroup = map[string]interface{}{
	"name":              "testProviderConnectionGroup",
	"parent_identifier": "ROOT",
	"type":              "ORGANIZATIONAL",
	"attributes": map[string]interface{}{
		"max_connections":          "4",
		"max_connections_per_user": "2",
		"enable_session_affinity":  true,
	},
}

func TestAccGuacamoleConnectionGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGuacamoleConnectionGroupConfigBasic(toHclString(testProviderConnectionGroup, true)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuacamoleUserGroupExists("guacamole_connection_group.new"),
					resource.TestCheckResourceAttr("guacamole_connection_group.new", "parent_identifier", testProviderConnectionGroup["parent_identifier"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_group.new", "attributes.0.max_connections", testProviderConnectionGroup["attributes"].(map[string]interface{})["max_connections"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_group.new", "attributes.0.max_connections_per_user", testProviderConnectionGroup["attributes"].(map[string]interface{})["max_connections_per_user"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_group.new", "attributes.0.enable_session_affinity", boolToString(testProviderConnectionGroup["attributes"].(map[string]interface{})["enable_session_affinity"].(bool))),
				),
			},
		},
	})
}

func testAccCheckGuacamoleConnectionGroupConfigBasic(definition string) string {
	return fmt.Sprintf(`
	resource "guacamole_connection_group" "new" %s
	`, definition)
}

func testAccCheckGuacamoleConnectionGroupExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No OrderID set")
		}

		return nil
	}
}
