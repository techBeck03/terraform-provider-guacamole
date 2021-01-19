package guacamole

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGuacamoleConnectionTelnetBasic(t *testing.T) {
	testProviderConnectionTelnet := map[string]interface{}{
		"name": "testProviderConnectionTelnet",
		"attributes": map[string]interface{}{
			"guacd_hostname":           "example.example.com",
			"guacd_port":               "8443",
			"guacd_encryption":         "ssl",
			"failover_only":            true,
			"weight":                   "10",
			"max_connections":          "4",
			"max_connections_per_user": "2",
		},
		"parameters": map[string]interface{}{
			"hostname":                    "hostname.example.com",
			"port":                        "22",
			"username":                    "user",
			"password":                    "password",
			"username_regex":              "[Uu]sername:",
			"password_regex":              "[Pp]assword:",
			"login_success_regex":         ".*[>]",
			"login_failure_regex":         ".*[:]",
			"color_scheme":                "green-black",
			"font_name":                   "Helvetica, sans-serif",
			"font_size":                   "12",
			"max_scrollback_size":         "200",
			"readonly":                    true,
			"disable_copy":                true,
			"disable_paste":               true,
			"backspace":                   "127",
			"terminal_type":               "vt100",
			"typescript_path":             "typescript path",
			"typescript_name":             "typescript name",
			"typescript_auto_create_path": true,
			"recording_path":              "recording path",
			"recording_name":              "recording name",
			"recording_exclude_output":    true,
			"recording_exclude_mouse":     true,
			"recording_include_keys":      true,
			"recording_auto_create_path":  true,
			"wol_send_packet":             true,
			"wol_mac_address":             "00:11:22:33:44",
			"wol_broadcast_address":       "255.255.255.254",
			"wol_boot_wait_time":          "5",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGuacamoleConnectionTelnetConfigBasic(toHclString(testProviderConnectionTelnet, true)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuacamoleConnectionTelnetExists("guacamole_connection_telnet.new"),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "name", testProviderConnectionTelnet["name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "attributes.0.guacd_hostname", testProviderConnectionTelnet["attributes"].(map[string]interface{})["guacd_hostname"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "attributes.0.guacd_port", testProviderConnectionTelnet["attributes"].(map[string]interface{})["guacd_port"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "attributes.0.guacd_encryption", testProviderConnectionTelnet["attributes"].(map[string]interface{})["guacd_encryption"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "attributes.0.failover_only", boolToString(testProviderConnectionTelnet["attributes"].(map[string]interface{})["failover_only"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "attributes.0.weight", testProviderConnectionTelnet["attributes"].(map[string]interface{})["weight"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "attributes.0.max_connections", testProviderConnectionTelnet["attributes"].(map[string]interface{})["max_connections"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "attributes.0.max_connections_per_user", testProviderConnectionTelnet["attributes"].(map[string]interface{})["max_connections_per_user"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.hostname", testProviderConnectionTelnet["parameters"].(map[string]interface{})["hostname"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.port", testProviderConnectionTelnet["parameters"].(map[string]interface{})["port"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.username", testProviderConnectionTelnet["parameters"].(map[string]interface{})["username"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.password", testProviderConnectionTelnet["parameters"].(map[string]interface{})["password"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.username_regex", testProviderConnectionTelnet["parameters"].(map[string]interface{})["username_regex"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.password_regex", testProviderConnectionTelnet["parameters"].(map[string]interface{})["password_regex"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.login_success_regex", testProviderConnectionTelnet["parameters"].(map[string]interface{})["login_success_regex"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.login_failure_regex", testProviderConnectionTelnet["parameters"].(map[string]interface{})["login_failure_regex"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.color_scheme", testProviderConnectionTelnet["parameters"].(map[string]interface{})["color_scheme"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.font_name", testProviderConnectionTelnet["parameters"].(map[string]interface{})["font_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.font_size", testProviderConnectionTelnet["parameters"].(map[string]interface{})["font_size"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.max_scrollback_size", testProviderConnectionTelnet["parameters"].(map[string]interface{})["max_scrollback_size"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.readonly", boolToString(testProviderConnectionTelnet["parameters"].(map[string]interface{})["readonly"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.disable_copy", boolToString(testProviderConnectionTelnet["parameters"].(map[string]interface{})["disable_copy"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.disable_paste", boolToString(testProviderConnectionTelnet["parameters"].(map[string]interface{})["disable_paste"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.backspace", testProviderConnectionTelnet["parameters"].(map[string]interface{})["backspace"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.terminal_type", testProviderConnectionTelnet["parameters"].(map[string]interface{})["terminal_type"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.typescript_path", testProviderConnectionTelnet["parameters"].(map[string]interface{})["typescript_path"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.typescript_name", testProviderConnectionTelnet["parameters"].(map[string]interface{})["typescript_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.typescript_auto_create_path", boolToString(testProviderConnectionTelnet["parameters"].(map[string]interface{})["typescript_auto_create_path"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.recording_path", testProviderConnectionTelnet["parameters"].(map[string]interface{})["recording_path"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.recording_name", testProviderConnectionTelnet["parameters"].(map[string]interface{})["recording_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.recording_exclude_output", boolToString(testProviderConnectionTelnet["parameters"].(map[string]interface{})["recording_exclude_output"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.recording_exclude_mouse", boolToString(testProviderConnectionTelnet["parameters"].(map[string]interface{})["recording_exclude_mouse"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.recording_include_keys", boolToString(testProviderConnectionTelnet["parameters"].(map[string]interface{})["recording_include_keys"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.recording_auto_create_path", boolToString(testProviderConnectionTelnet["parameters"].(map[string]interface{})["recording_auto_create_path"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.wol_send_packet", boolToString(testProviderConnectionTelnet["parameters"].(map[string]interface{})["wol_send_packet"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.wol_mac_address", testProviderConnectionTelnet["parameters"].(map[string]interface{})["wol_mac_address"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.wol_broadcast_address", testProviderConnectionTelnet["parameters"].(map[string]interface{})["wol_broadcast_address"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_telnet.new", "parameters.0.wol_boot_wait_time", testProviderConnectionTelnet["parameters"].(map[string]interface{})["wol_boot_wait_time"].(string)),
				),
			},
		},
	})
}

func testAccCheckGuacamoleConnectionTelnetConfigBasic(connection string) string {
	return fmt.Sprintf(`
	resource "guacamole_connection_telnet" "new" %s
	`, connection)
}

func testAccCheckGuacamoleConnectionTelnetExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No connection id set")
		}

		return nil
	}
}
