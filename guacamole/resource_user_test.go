package guacamole

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	guac "github.com/techBeck03/guacamole-api-client"
)

func TestAccGuacamoleUserBasic(t *testing.T) {
	username := "testProviderUser"
	attributes := map[string]interface{}{
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
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGuacamoleUserConfigBasic(username, toHclString(attributes, false)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuacamoleUserExists("guacamole_user.new"),
					resource.TestCheckResourceAttr("guacamole_user.new", "username", username),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.organizational_role", attributes["organizational_role"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.full_name", attributes["full_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.email", attributes["email"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.expired", boolToString(attributes["expired"].(bool))),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.timezone", attributes["timezone"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.access_window_start", attributes["access_window_start"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.access_window_end", attributes["access_window_end"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.disabled", boolToString(attributes["disabled"].(bool))),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.valid_from", attributes["valid_from"].(string)),
					resource.TestCheckResourceAttr("guacamole_user.new", "attributes.0.valid_until", attributes["valid_until"].(string)),
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

		err := c.DeleteUser(username)
		if err != nil {
			return err
		}
	}

	return nil
}

func testAccCheckGuacamoleUserConfigBasic(username string, attributes string) string {
	return fmt.Sprintf(`
	resource "guacamole_user" "new" {
		username = "%s"
		attributes %s
	}
	`, username, attributes)
}

func testAccCheckGuacamoleUserExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Println("Testing")
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No OrderID set")
		}

		return nil
	}
}
