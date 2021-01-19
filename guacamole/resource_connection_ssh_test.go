package guacamole

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGuacamoleConnectionSSHBasic(t *testing.T) {
	testProviderConnectionSSH := map[string]interface{}{
		"name": "testProviderConnectionSSH",
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
			"public_host_key":             "public host key",
			"username":                    "user",
			"password":                    "password",
			"private_key":                 "super secret private key",
			"passphrase":                  "gigem",
			"color_scheme":                "green-black",
			"font_name":                   "Helvetica, sans-serif",
			"font_size":                   "12",
			"max_scrollback_size":         "200",
			"readonly":                    true,
			"disable_copy":                true,
			"disable_paste":               true,
			"execute_command":             "run this",
			"timezone":                    "America/Chicago",
			"server_keepalive":            "20",
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
			"sftp_enable":                 true,
			"sftp_root_directory":         "sftp/root/directory",
			"sftp_disable_file_download":  true,
			"sftp_disable_file_upload":    true,
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
				Config: testAccCheckGuacamoleConnectionSSHConfigBasic(toHclString(testProviderConnectionSSH, true)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuacamoleConnectionSSHExists("guacamole_connection_ssh.new"),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "name", testProviderConnectionSSH["name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "attributes.0.guacd_hostname", testProviderConnectionSSH["attributes"].(map[string]interface{})["guacd_hostname"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "attributes.0.guacd_port", testProviderConnectionSSH["attributes"].(map[string]interface{})["guacd_port"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "attributes.0.guacd_encryption", testProviderConnectionSSH["attributes"].(map[string]interface{})["guacd_encryption"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "attributes.0.failover_only", boolToString(testProviderConnectionSSH["attributes"].(map[string]interface{})["failover_only"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "attributes.0.weight", testProviderConnectionSSH["attributes"].(map[string]interface{})["weight"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "attributes.0.max_connections", testProviderConnectionSSH["attributes"].(map[string]interface{})["max_connections"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "attributes.0.max_connections_per_user", testProviderConnectionSSH["attributes"].(map[string]interface{})["max_connections_per_user"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.hostname", testProviderConnectionSSH["parameters"].(map[string]interface{})["hostname"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.port", testProviderConnectionSSH["parameters"].(map[string]interface{})["port"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.public_host_key", testProviderConnectionSSH["parameters"].(map[string]interface{})["public_host_key"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.username", testProviderConnectionSSH["parameters"].(map[string]interface{})["username"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.password", testProviderConnectionSSH["parameters"].(map[string]interface{})["password"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.private_key", testProviderConnectionSSH["parameters"].(map[string]interface{})["private_key"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.passphrase", testProviderConnectionSSH["parameters"].(map[string]interface{})["passphrase"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.color_scheme", testProviderConnectionSSH["parameters"].(map[string]interface{})["color_scheme"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.font_name", testProviderConnectionSSH["parameters"].(map[string]interface{})["font_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.font_size", testProviderConnectionSSH["parameters"].(map[string]interface{})["font_size"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.max_scrollback_size", testProviderConnectionSSH["parameters"].(map[string]interface{})["max_scrollback_size"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.readonly", boolToString(testProviderConnectionSSH["parameters"].(map[string]interface{})["readonly"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.disable_copy", boolToString(testProviderConnectionSSH["parameters"].(map[string]interface{})["disable_copy"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.disable_paste", boolToString(testProviderConnectionSSH["parameters"].(map[string]interface{})["disable_paste"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.execute_command", testProviderConnectionSSH["parameters"].(map[string]interface{})["execute_command"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.timezone", testProviderConnectionSSH["parameters"].(map[string]interface{})["timezone"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.server_keepalive", testProviderConnectionSSH["parameters"].(map[string]interface{})["server_keepalive"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.backspace", testProviderConnectionSSH["parameters"].(map[string]interface{})["backspace"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.terminal_type", testProviderConnectionSSH["parameters"].(map[string]interface{})["terminal_type"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.typescript_path", testProviderConnectionSSH["parameters"].(map[string]interface{})["typescript_path"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.typescript_name", testProviderConnectionSSH["parameters"].(map[string]interface{})["typescript_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.typescript_auto_create_path", boolToString(testProviderConnectionSSH["parameters"].(map[string]interface{})["typescript_auto_create_path"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.recording_path", testProviderConnectionSSH["parameters"].(map[string]interface{})["recording_path"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.recording_name", testProviderConnectionSSH["parameters"].(map[string]interface{})["recording_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.recording_exclude_output", boolToString(testProviderConnectionSSH["parameters"].(map[string]interface{})["recording_exclude_output"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.recording_exclude_mouse", boolToString(testProviderConnectionSSH["parameters"].(map[string]interface{})["recording_exclude_mouse"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.recording_include_keys", boolToString(testProviderConnectionSSH["parameters"].(map[string]interface{})["recording_include_keys"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.recording_auto_create_path", boolToString(testProviderConnectionSSH["parameters"].(map[string]interface{})["recording_auto_create_path"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.sftp_enable", boolToString(testProviderConnectionSSH["parameters"].(map[string]interface{})["sftp_enable"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.sftp_root_directory", testProviderConnectionSSH["parameters"].(map[string]interface{})["sftp_root_directory"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.sftp_disable_file_download", boolToString(testProviderConnectionSSH["parameters"].(map[string]interface{})["sftp_disable_file_download"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.sftp_disable_file_upload", boolToString(testProviderConnectionSSH["parameters"].(map[string]interface{})["sftp_disable_file_upload"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.wol_send_packet", boolToString(testProviderConnectionSSH["parameters"].(map[string]interface{})["wol_send_packet"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.wol_mac_address", testProviderConnectionSSH["parameters"].(map[string]interface{})["wol_mac_address"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.wol_broadcast_address", testProviderConnectionSSH["parameters"].(map[string]interface{})["wol_broadcast_address"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_ssh.new", "parameters.0.wol_boot_wait_time", testProviderConnectionSSH["parameters"].(map[string]interface{})["wol_boot_wait_time"].(string)),
				),
			},
		},
	})
}

func testAccCheckGuacamoleConnectionSSHConfigBasic(connection string) string {
	return fmt.Sprintf(`
	resource "guacamole_connection_ssh" "new" %s
	`, connection)
}

func testAccCheckGuacamoleConnectionSSHExists(n string) resource.TestCheckFunc {
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
