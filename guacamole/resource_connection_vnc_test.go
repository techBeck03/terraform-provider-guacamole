package guacamole

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGuacamoleConnectionVNCBasic(t *testing.T) {
	testProviderConnectionVNC := map[string]interface{}{
		"name": "testProviderConnectionVNC",
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
			"hostname":                   "hostname.example.com",
			"port":                       "22",
			"username":                   "user",
			"password":                   "password",
			"readonly":                   true,
			"swap_red_blue":              true,
			"cursor":                     "local",
			"color_depth":                "24",
			"clipboard_encoding":         "ISO8859-1",
			"disable_copy":               true,
			"disable_paste":              true,
			"destination_host":           "destination host",
			"destination_port":           "8443",
			"recording_path":             "recording path",
			"recording_name":             "recording name",
			"recording_exclude_output":   true,
			"recording_exclude_mouse":    true,
			"recording_include_keys":     true,
			"recording_auto_create_path": true,
			"sftp_enable":                true,
			"sftp_root_directory":        "sftp/root/directory",
			"sftp_hostname":              "sftp.example.com",
			"sftp_port":                  "22",
			"sftp_host_key":              "host key",
			"sftp_username":              "sftp username",
			"sftp_password":              "sftp password",
			"sftp_private_key":           "sftp private key",
			"sftp_passphrase":            "aggies",
			"sftp_upload_directory":      "sftp default upload",
			"sftp_keepalive_interval":    "60",
			"sftp_disable_file_download": true,
			"sftp_disable_file_upload":   true,
			"wol_send_packet":            true,
			"wol_mac_address":            "00:11:22:33:44",
			"wol_broadcast_address":      "255.255.255.254",
			"wol_boot_wait_time":         "5",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGuacamoleConnectionVNCConfigBasic(toHclString(testProviderConnectionVNC, true)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuacamoleConnectionVNCExists("guacamole_connection_vnc.new"),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "name", testProviderConnectionVNC["name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "attributes.0.guacd_hostname", testProviderConnectionVNC["attributes"].(map[string]interface{})["guacd_hostname"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "attributes.0.guacd_port", testProviderConnectionVNC["attributes"].(map[string]interface{})["guacd_port"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "attributes.0.guacd_encryption", testProviderConnectionVNC["attributes"].(map[string]interface{})["guacd_encryption"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "attributes.0.failover_only", boolToString(testProviderConnectionVNC["attributes"].(map[string]interface{})["failover_only"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "attributes.0.weight", testProviderConnectionVNC["attributes"].(map[string]interface{})["weight"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "attributes.0.max_connections", testProviderConnectionVNC["attributes"].(map[string]interface{})["max_connections"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "attributes.0.max_connections_per_user", testProviderConnectionVNC["attributes"].(map[string]interface{})["max_connections_per_user"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.hostname", testProviderConnectionVNC["parameters"].(map[string]interface{})["hostname"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.port", testProviderConnectionVNC["parameters"].(map[string]interface{})["port"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.username", testProviderConnectionVNC["parameters"].(map[string]interface{})["username"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.password", testProviderConnectionVNC["parameters"].(map[string]interface{})["password"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.readonly", boolToString(testProviderConnectionVNC["parameters"].(map[string]interface{})["readonly"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.swap_red_blue", boolToString(testProviderConnectionVNC["parameters"].(map[string]interface{})["swap_red_blue"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.cursor", testProviderConnectionVNC["parameters"].(map[string]interface{})["cursor"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.color_depth", testProviderConnectionVNC["parameters"].(map[string]interface{})["color_depth"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.clipboard_encoding", testProviderConnectionVNC["parameters"].(map[string]interface{})["clipboard_encoding"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.disable_copy", boolToString(testProviderConnectionVNC["parameters"].(map[string]interface{})["disable_copy"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.disable_paste", boolToString(testProviderConnectionVNC["parameters"].(map[string]interface{})["disable_paste"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.destination_host", testProviderConnectionVNC["parameters"].(map[string]interface{})["destination_host"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.destination_port", testProviderConnectionVNC["parameters"].(map[string]interface{})["destination_port"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.recording_path", testProviderConnectionVNC["parameters"].(map[string]interface{})["recording_path"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.recording_name", testProviderConnectionVNC["parameters"].(map[string]interface{})["recording_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.recording_exclude_output", boolToString(testProviderConnectionVNC["parameters"].(map[string]interface{})["recording_exclude_output"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.recording_exclude_mouse", boolToString(testProviderConnectionVNC["parameters"].(map[string]interface{})["recording_exclude_mouse"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.recording_include_keys", boolToString(testProviderConnectionVNC["parameters"].(map[string]interface{})["recording_include_keys"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.recording_auto_create_path", boolToString(testProviderConnectionVNC["parameters"].(map[string]interface{})["recording_auto_create_path"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.sftp_enable", boolToString(testProviderConnectionVNC["parameters"].(map[string]interface{})["sftp_enable"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.sftp_root_directory", testProviderConnectionVNC["parameters"].(map[string]interface{})["sftp_root_directory"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.sftp_hostname", testProviderConnectionVNC["parameters"].(map[string]interface{})["sftp_hostname"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.sftp_port", testProviderConnectionVNC["parameters"].(map[string]interface{})["sftp_port"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.sftp_host_key", testProviderConnectionVNC["parameters"].(map[string]interface{})["sftp_host_key"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.sftp_username", testProviderConnectionVNC["parameters"].(map[string]interface{})["sftp_username"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.sftp_password", testProviderConnectionVNC["parameters"].(map[string]interface{})["sftp_password"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.sftp_private_key", testProviderConnectionVNC["parameters"].(map[string]interface{})["sftp_private_key"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.sftp_passphrase", testProviderConnectionVNC["parameters"].(map[string]interface{})["sftp_passphrase"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.sftp_upload_directory", testProviderConnectionVNC["parameters"].(map[string]interface{})["sftp_upload_directory"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.sftp_keepalive_interval", testProviderConnectionVNC["parameters"].(map[string]interface{})["sftp_keepalive_interval"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.sftp_disable_file_download", boolToString(testProviderConnectionVNC["parameters"].(map[string]interface{})["sftp_disable_file_download"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.sftp_disable_file_upload", boolToString(testProviderConnectionVNC["parameters"].(map[string]interface{})["sftp_disable_file_upload"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.wol_send_packet", boolToString(testProviderConnectionVNC["parameters"].(map[string]interface{})["wol_send_packet"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.wol_mac_address", testProviderConnectionVNC["parameters"].(map[string]interface{})["wol_mac_address"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.wol_broadcast_address", testProviderConnectionVNC["parameters"].(map[string]interface{})["wol_broadcast_address"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_vnc.new", "parameters.0.wol_boot_wait_time", testProviderConnectionVNC["parameters"].(map[string]interface{})["wol_boot_wait_time"].(string)),
				),
			},
		},
	})
}

func testAccCheckGuacamoleConnectionVNCConfigBasic(connection string) string {
	return fmt.Sprintf(`
	resource "guacamole_connection_vnc" "new" %s
	`, connection)
}

func testAccCheckGuacamoleConnectionVNCExists(n string) resource.TestCheckFunc {
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
