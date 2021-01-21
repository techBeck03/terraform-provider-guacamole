package guacamole

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGuacamoleConnectionRDPBasic(t *testing.T) {
	testProviderConnectionRDP := map[string]interface{}{
		"name": "testProviderConnectionRDP",
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
			"hostname":                     "hostname.example.com",
			"port":                         "22",
			"username":                     "user",
			"password":                     "password",
			"domain":                       "example.domain.com",
			"security_mode":                "nla",
			"disable_authentication":       true,
			"ignore_cert":                  true,
			"gateway_hostname":             "gw hostname",
			"gateway_port":                 "8443",
			"gateway_username":             "gw username",
			"gateway_password":             "gw password",
			"gateway_domain":               "gw.domain.com",
			"initial_program":              "powershell",
			"client_name":                  "client name",
			"keyboard_layout":              "en-us-qwerty",
			"timezone":                     "America/Chicago",
			"administrator_console":        true,
			"width":                        "2560",
			"height":                       "1440",
			"dpi":                          "96",
			"color_depth":                  "24",
			"resize_method":                "display-update",
			"readonly":                     true,
			"disable_copy":                 true,
			"disable_paste":                true,
			"console_audio":                true,
			"disable_audio":                true,
			"enable_audio_input":           true,
			"enable_printing":              true,
			"printer_name":                 "printer name",
			"enable_drive":                 true,
			"drive_name":                   "drive name",
			"disable_file_download":        true,
			"disable_file_upload":          true,
			"drive_path":                   "drive path",
			"create_drive_path":            true,
			"static_channels":              "channel-1, channel-2",
			"enable_wallpaper":             true,
			"enable_theming":               true,
			"enable_font_smoothing":        true,
			"enable_full_window_drag":      true,
			"enable_desktop_composition":   true,
			"enable_menu_animations":       true,
			"disable_bitmap_caching":       true,
			"disable_offscreen_caching":    true,
			"disable_glyph_caching":        true,
			"remote_app":                   "remote app",
			"remote_app_working_directory": "remote app working directory",
			"remote_app_parameters":        "remote app parameters",
			"preconnection_id":             "2",
			"preconnection_blob":           "preconnection blob",
			"load_balance_info":            "load balance info",
			"recording_path":               "recording path",
			"recording_name":               "recording name",
			"recording_exclude_output":     true,
			"recording_exclude_mouse":      true,
			"recording_include_keys":       true,
			"recording_auto_create_path":   true,
			"sftp_enable":                  true,
			"sftp_root_directory":          "sftp/root/directory",
			"sftp_hostname":                "sftp.example.com",
			"sftp_port":                    "22",
			"sftp_host_key":                "host key",
			"sftp_username":                "sftp username",
			"sftp_password":                "sftp password",
			"sftp_private_key":             "sftp private key",
			"sftp_passphrase":              "aggies",
			"sftp_upload_directory":        "sftp default upload",
			"sftp_keepalive_interval":      "60",
			"sftp_disable_file_download":   true,
			"sftp_disable_file_upload":     true,
			"wol_send_packet":              true,
			"wol_mac_address":              "00:11:22:33:44",
			"wol_broadcast_address":        "255.255.255.254",
			"wol_boot_wait_time":           "5",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGuacamoleConnectionRDPConfigBasic(toHclString(testProviderConnectionRDP, true)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuacamoleConnectionRDPExists("guacamole_connection_rdp.new"),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "name", testProviderConnectionRDP["name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "attributes.0.guacd_hostname", testProviderConnectionRDP["attributes"].(map[string]interface{})["guacd_hostname"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "attributes.0.guacd_port", testProviderConnectionRDP["attributes"].(map[string]interface{})["guacd_port"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "attributes.0.guacd_encryption", testProviderConnectionRDP["attributes"].(map[string]interface{})["guacd_encryption"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "attributes.0.failover_only", boolToString(testProviderConnectionRDP["attributes"].(map[string]interface{})["failover_only"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "attributes.0.weight", testProviderConnectionRDP["attributes"].(map[string]interface{})["weight"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "attributes.0.max_connections", testProviderConnectionRDP["attributes"].(map[string]interface{})["max_connections"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "attributes.0.max_connections_per_user", testProviderConnectionRDP["attributes"].(map[string]interface{})["max_connections_per_user"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.hostname", testProviderConnectionRDP["parameters"].(map[string]interface{})["hostname"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.port", testProviderConnectionRDP["parameters"].(map[string]interface{})["port"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.username", testProviderConnectionRDP["parameters"].(map[string]interface{})["username"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.password", testProviderConnectionRDP["parameters"].(map[string]interface{})["password"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.domain", testProviderConnectionRDP["parameters"].(map[string]interface{})["domain"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.security_mode", testProviderConnectionRDP["parameters"].(map[string]interface{})["security_mode"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.disable_authentication", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["disable_authentication"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.ignore_cert", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["ignore_cert"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.gateway_hostname", testProviderConnectionRDP["parameters"].(map[string]interface{})["gateway_hostname"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.gateway_port", testProviderConnectionRDP["parameters"].(map[string]interface{})["gateway_port"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.gateway_username", testProviderConnectionRDP["parameters"].(map[string]interface{})["gateway_username"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.gateway_password", testProviderConnectionRDP["parameters"].(map[string]interface{})["gateway_password"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.gateway_domain", testProviderConnectionRDP["parameters"].(map[string]interface{})["gateway_domain"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.initial_program", testProviderConnectionRDP["parameters"].(map[string]interface{})["initial_program"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.client_name", testProviderConnectionRDP["parameters"].(map[string]interface{})["client_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.keyboard_layout", testProviderConnectionRDP["parameters"].(map[string]interface{})["keyboard_layout"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.timezone", testProviderConnectionRDP["parameters"].(map[string]interface{})["timezone"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.administrator_console", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["administrator_console"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.width", testProviderConnectionRDP["parameters"].(map[string]interface{})["width"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.height", testProviderConnectionRDP["parameters"].(map[string]interface{})["height"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.dpi", testProviderConnectionRDP["parameters"].(map[string]interface{})["dpi"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.color_depth", testProviderConnectionRDP["parameters"].(map[string]interface{})["color_depth"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.resize_method", testProviderConnectionRDP["parameters"].(map[string]interface{})["resize_method"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.readonly", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["readonly"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.disable_copy", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["disable_copy"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.disable_paste", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["disable_paste"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.enable_audio_input", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["enable_audio_input"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.enable_printing", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["enable_printing"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.printer_name", testProviderConnectionRDP["parameters"].(map[string]interface{})["printer_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.enable_drive", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["enable_drive"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.drive_name", testProviderConnectionRDP["parameters"].(map[string]interface{})["drive_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.disable_file_download", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["disable_file_download"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.disable_file_upload", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["disable_file_upload"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.drive_path", testProviderConnectionRDP["parameters"].(map[string]interface{})["drive_path"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.create_drive_path", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["create_drive_path"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.static_channels", testProviderConnectionRDP["parameters"].(map[string]interface{})["static_channels"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.enable_wallpaper", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["enable_wallpaper"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.enable_theming", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["enable_theming"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.enable_font_smoothing", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["enable_font_smoothing"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.enable_full_window_drag", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["enable_full_window_drag"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.enable_desktop_composition", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["enable_desktop_composition"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.enable_menu_animations", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["enable_menu_animations"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.disable_bitmap_caching", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["disable_bitmap_caching"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.disable_offscreen_caching", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["disable_offscreen_caching"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.disable_glyph_caching", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["disable_glyph_caching"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.remote_app", testProviderConnectionRDP["parameters"].(map[string]interface{})["remote_app"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.remote_app_working_directory", testProviderConnectionRDP["parameters"].(map[string]interface{})["remote_app_working_directory"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.remote_app_parameters", testProviderConnectionRDP["parameters"].(map[string]interface{})["remote_app_parameters"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.preconnection_id", testProviderConnectionRDP["parameters"].(map[string]interface{})["preconnection_id"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.preconnection_blob", testProviderConnectionRDP["parameters"].(map[string]interface{})["preconnection_blob"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.load_balance_info", testProviderConnectionRDP["parameters"].(map[string]interface{})["load_balance_info"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.recording_path", testProviderConnectionRDP["parameters"].(map[string]interface{})["recording_path"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.recording_name", testProviderConnectionRDP["parameters"].(map[string]interface{})["recording_name"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.recording_exclude_output", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["recording_exclude_output"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.recording_exclude_mouse", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["recording_exclude_mouse"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.recording_include_keys", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["recording_include_keys"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.recording_auto_create_path", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["recording_auto_create_path"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.sftp_enable", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["sftp_enable"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.sftp_root_directory", testProviderConnectionRDP["parameters"].(map[string]interface{})["sftp_root_directory"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.sftp_hostname", testProviderConnectionRDP["parameters"].(map[string]interface{})["sftp_hostname"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.sftp_port", testProviderConnectionRDP["parameters"].(map[string]interface{})["sftp_port"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.sftp_host_key", testProviderConnectionRDP["parameters"].(map[string]interface{})["sftp_host_key"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.sftp_username", testProviderConnectionRDP["parameters"].(map[string]interface{})["sftp_username"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.sftp_password", testProviderConnectionRDP["parameters"].(map[string]interface{})["sftp_password"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.sftp_private_key", testProviderConnectionRDP["parameters"].(map[string]interface{})["sftp_private_key"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.sftp_passphrase", testProviderConnectionRDP["parameters"].(map[string]interface{})["sftp_passphrase"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.sftp_upload_directory", testProviderConnectionRDP["parameters"].(map[string]interface{})["sftp_upload_directory"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.sftp_keepalive_interval", testProviderConnectionRDP["parameters"].(map[string]interface{})["sftp_keepalive_interval"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.sftp_disable_file_download", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["sftp_disable_file_download"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.sftp_disable_file_upload", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["sftp_disable_file_upload"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.wol_send_packet", boolToString(testProviderConnectionRDP["parameters"].(map[string]interface{})["wol_send_packet"].(bool))),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.wol_mac_address", testProviderConnectionRDP["parameters"].(map[string]interface{})["wol_mac_address"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.wol_broadcast_address", testProviderConnectionRDP["parameters"].(map[string]interface{})["wol_broadcast_address"].(string)),
					resource.TestCheckResourceAttr("guacamole_connection_rdp.new", "parameters.0.wol_boot_wait_time", testProviderConnectionRDP["parameters"].(map[string]interface{})["wol_boot_wait_time"].(string)),
				),
			},
		},
	})
}

func testAccCheckGuacamoleConnectionRDPConfigBasic(connection string) string {
	return fmt.Sprintf(`
	resource "guacamole_connection_rdp" "new" %s
	`, connection)
}

func testAccCheckGuacamoleConnectionRDPExists(n string) resource.TestCheckFunc {
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
