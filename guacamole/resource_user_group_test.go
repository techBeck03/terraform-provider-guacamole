package guacamole

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	guac "github.com/techBeck03/guacamole-api-client"
	types "github.com/techBeck03/guacamole-api-client/types"
)

var testProviderUserGroup = map[string]interface{}{
	"identifier":         "testProviderUserGroup",
	"system_permissions": types.SystemPermissions{}.ValidChoices(),
}

func TestAccGuacamoleUserGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGuacamoleUserGroupConfigBasic(toHclString(testProviderUserGroup, true)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuacamoleUserGroupExists("guacamole_user_group.new"),
					resource.TestCheckResourceAttr("guacamole_user_group.new", "identifier", testProviderUserGroup["identifier"].(string)),
					testAccCheckTestSliceVals("guacamole_user_group.new", "system_permissions", types.SystemPermissions{}.ValidChoices()),
				),
			},
		},
	})
}

func testAccCheckGuacamoleUserGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*guac.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "guacamole_user_group" {
			continue
		}

		username := rs.Primary.ID

		err := c.DeleteUser(username)
		if err != nil {
			return err
		}
	}

	return nil
}

func testAccCheckGuacamoleUserGroupConfigBasic(definition string) string {
	return fmt.Sprintf(`
	resource "guacamole_user_group" "new" %s
	`, definition)
}

func testAccCheckGuacamoleUserGroupExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No group identifier set")
		}

		return nil
	}
}
