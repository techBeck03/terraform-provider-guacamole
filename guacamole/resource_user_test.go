package guacamole

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	guac "github.com/techBeck03/guacamole-api-client"
	types "github.com/techBeck03/guacamole-api-client/types"
)

func TestAccGuacamoleUserBasic(t *testing.T) {
	testProviderUser := map[string]interface{}{
		"username": "testProviderUser",
		"attributes": map[string]interface{}{
			"organizational_role": "testRole",
			"full_name":           "Guac Provider Test",
			"email":               "guacProvider@example.com",
			"expired":             true,
			"timezone":            "America/Chicago",
			"access_window_start": "09:00:00",
			"access_window_end":   "21:00:00",
			"disabled":            true,
			"valid_from":          "2021-01-01",
			"valid_until":         "2022-01-01",
		},
		"system_permissions": types.SystemPermissions{}.ValidChoices(),
		"group_membership":   []string{testProviderUserGroup["identifier"].(string)},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGuacamoleUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGuacamoleUserConfigBasic(toHclString(testProviderUserGroup, true), toHclString(testProviderUser, true)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuacamoleUserGroupExists("guacamole_user_group.new"),
					resource.TestCheckResourceAttr("guacamole_user_group.new", "identifier", testProviderUserGroup["identifier"].(string)),
					testAccCheckGuacamoleUserExists("guacamole_user.new"),
					resource.TestCheckResourceAttr("guacamole_user.new", "username", testProviderUser["username"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.organizational_role", testProviderUser["attributes"].(map[string]interface{})["organizational_role"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.full_name", testProviderUser["attributes"].(map[string]interface{})["full_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.email", testProviderUser["attributes"].(map[string]interface{})["email"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.expired", boolToString(testProviderUser["attributes"].(map[string]interface{})["expired"].(bool))),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.timezone", testProviderUser["attributes"].(map[string]interface{})["timezone"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.access_window_start", testProviderUser["attributes"].(map[string]interface{})["access_window_start"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.access_window_end", testProviderUser["attributes"].(map[string]interface{})["access_window_end"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.disabled", boolToString(testProviderUser["attributes"].(map[string]interface{})["disabled"].(bool))),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.valid_from", testProviderUser["attributes"].(map[string]interface{})["valid_from"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.valid_until", testProviderUser["attributes"].(map[string]interface{})["valid_until"].(string)),
					testAccCheckTestSliceVals("guacamole_user.new", "system_permissions", types.SystemPermissions{}.ValidChoices()),
					testAccCheckTestSliceVals("guacamole_user.new", "group_membership", []string{testProviderUserGroup["identifier"].(string)}),
				),
			},
		},
	})
}

func testAccCheckGuacamoleUserDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*guac.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "guacamole_user" {
			continue
		}

		username := rs.Primary.ID

		user, err := c.ReadUser(username)
		if err != nil {
			return nil
		}
		if user.Username != "" {
			return fmt.Errorf("Username %s still exists in guacamole database", username)
		}
	}

	return nil
}

func testAccCheckGuacamoleUserConfigBasic(group string, user string) string {
	return fmt.Sprintf(`
	resource "guacamole_user_group" "new" %s
	resource "guacamole_user" "new" %s
	`, group, user)
}

func testAccCheckGuacamoleUserExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No username set")
		}

		return nil
	}
}
