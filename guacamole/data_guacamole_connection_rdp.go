package guacamole

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	guac "github.com/techBeck03/guacamole-api-client"
	types "github.com/techBeck03/guacamole-api-client/types"
)

func dataSourceConnectionRDP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConnectionSSHRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the guacamole connection",
				Computed:    true,
			},
			"identifier": {
				Type:        schema.TypeString,
				Description: "Numeric identifier of the guacamole connection",
				Optional:    true,
			},
			"path": {
				Type:        schema.TypeString,
				Description: "Path of connection",
				Optional:    true,
			},
			"parent_identifier": {
				Type:        schema.TypeString,
				Description: "Parent identifier of the guacamole connection",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "Protocol type of the guacamole connection",
				Computed:    true,
			},
			"active_connections": {
				Type:        schema.TypeInt,
				Description: "Active connection count for the guacamole connection",
				Computed:    true,
			},
			"attributes": {
				Type:        schema.TypeList,
				Description: "Guacamole connection attributes",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"guacd_hostname": {
							Type:        schema.TypeString,
							Description: "Guacd proxy hostname",
							Computed:    true,
						},
						"guacd_port": {
							Type:        schema.TypeString,
							Description: "Guacd proxy port",
							Computed:    true,
						},
						"guacd_encryption": {
							Type:        schema.TypeString,
							Description: "Guacd proxy encryption type",
							Computed:    true,
						},
						"failover_only": {
							Type:        schema.TypeBool,
							Description: "Use load balancing for failover only",
							Computed:    true,
						},
						"weight": {
							Type:        schema.TypeString,
							Description: "Load balancing connection weight",
							Computed:    true,
						},
						"max_connections": {
							Type:        schema.TypeString,
							Description: "Maximum concurrent total connections",
							Computed:    true,
						},
						"max_connections_per_user": {
							Type:        schema.TypeString,
							Description: "Maximum concurrent connections per user",
							Computed:    true,
						},
					},
				},
			},
			"parameters": {
				Type:        schema.TypeList,
				Description: "Guacamole connection parameters",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": {
							Type:        schema.TypeString,
							Description: "Hostname of target",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeString,
							Description: "Port for target connection",
							Computed:    true,
						},
						"username": {
							Type:        schema.TypeString,
							Description: "Username for rdp connection",
							Computed:    true,
						},
						"password": {
							Type:        schema.TypeString,
							Description: "Password for rdp connection",
							Computed:    true,
						},
						"domain": {
							Type:        schema.TypeString,
							Description: "Domain name of rdp connection",
							Computed:    true,
						},
						"security_mode": {
							Type:        schema.TypeString,
							Description: "RDP security mode",
							Computed:    true,
						},
						"disable_authentication": {
							Type:        schema.TypeBool,
							Description: "Disable rdp authentication",
							Computed:    true,
						},
						"ignore_cert": {
							Type:        schema.TypeBool,
							Description: "Ignore domain certificate warnings",
							Computed:    true,
						},
						"gateway_hostname": {
							Type:        schema.TypeString,
							Description: "RDS gateway hostname",
							Computed:    true,
						},
						"gateway_port": {
							Type:        schema.TypeString,
							Description: "RDS gateway port",
							Computed:    true,
						},
						"gateway_username": {
							Type:        schema.TypeString,
							Description: "RDS gateway username",
							Computed:    true,
						},
						"gateway_password": {
							Type:        schema.TypeString,
							Description: "RDS gateway password",
							Computed:    true,
						},
						"gateway_domain": {
							Type:        schema.TypeString,
							Description: "RDS gateway domain",
							Computed:    true,
						},
						"initial_program": {
							Type:        schema.TypeString,
							Description: "Initial program for rdp connection",
							Computed:    true,
						},
						"client_name": {
							Type:        schema.TypeString,
							Description: "Client name for rdp connection",
							Computed:    true,
						},
						"keyboard_layout": {
							Type:        schema.TypeString,
							Description: "Keyboard layout for rdp connection",
							Computed:    true,
						},
						"timezone": {
							Type:        schema.TypeString,
							Description: "Timezone/Locale for rdp connection",
							Computed:    true,
						},
						"administrator_console": {
							Type:        schema.TypeBool,
							Description: "Enable administrator console",
							Computed:    true,
						},
						"width": {
							Type:        schema.TypeString,
							Description: "Screen width (px)",
							Computed:    true,
						},
						"height": {
							Type:        schema.TypeString,
							Description: "Screen height (px)",
							Computed:    true,
						},
						"dpi": {
							Type:        schema.TypeString,
							Description: "Resolution (DPI) of rdp connection",
							Computed:    true,
						},
						"color_depth": {
							Type:        schema.TypeString,
							Description: "Color depth of rdp connection",
							Computed:    true,
						},
						"resize_method": {
							Type:        schema.TypeString,
							Description: "Resize method rdp connection",
							Computed:    true,
						},
						"readonly": {
							Type:        schema.TypeBool,
							Description: "Display is readonly",
							Computed:    true,
						},
						"disable_copy": {
							Type:        schema.TypeBool,
							Description: "Disable copying from terminal",
							Computed:    true,
						},
						"disable_paste": {
							Type:        schema.TypeBool,
							Description: "Disable pasting from client",
							Computed:    true,
						},
						"console_audio": {
							Type:        schema.TypeBool,
							Description: "Support audio in console",
							Computed:    true,
						},
						"disable_audio": {
							Type:        schema.TypeBool,
							Description: "Disable audio",
							Computed:    true,
						},
						"enable_audio_input": {
							Type:        schema.TypeBool,
							Description: "Enable audio input (microphone)",
							Computed:    true,
						},
						"enable_printing": {
							Type:        schema.TypeBool,
							Description: "Enable printing",
							Computed:    true,
						},
						"printer_name": {
							Type:        schema.TypeString,
							Description: "Redirected printer name",
							Computed:    true,
						},
						"enable_drive": {
							Type:        schema.TypeBool,
							Description: "Enable drive for device redirection",
							Computed:    true,
						},
						"drive_name": {
							Type:        schema.TypeString,
							Description: "Drive name for device redirection",
							Computed:    true,
						},
						"disable_file_download": {
							Type:        schema.TypeBool,
							Description: "Disable file download for device redirection",
							Computed:    true,
						},
						"disable_file_upload": {
							Type:        schema.TypeBool,
							Description: "Disable file upload for device redirection",
							Computed:    true,
						},
						"drive_path": {
							Type:        schema.TypeString,
							Description: "Drive path for device redirection",
							Computed:    true,
						},
						"create_drive_path": {
							Type:        schema.TypeBool,
							Description: "Create drive path for device redirection",
							Computed:    true,
						},
						"static_channels": {
							Type:        schema.TypeString,
							Description: "Static channel names",
							Computed:    true,
						},
						"enable_wallpaper": {
							Type:        schema.TypeBool,
							Description: "Enable wallpaper",
							Computed:    true,
						},
						"enable_theming": {
							Type:        schema.TypeBool,
							Description: "Enable theming",
							Computed:    true,
						},
						"enable_font_smoothing": {
							Type:        schema.TypeBool,
							Description: "Enable font smoothing",
							Computed:    true,
						},
						"enable_full_window_drag": {
							Type:        schema.TypeBool,
							Description: "Enable full window drag",
							Computed:    true,
						},
						"enable_desktop_composition": {
							Type:        schema.TypeBool,
							Description: "Enable desktop composition",
							Computed:    true,
						},
						"enable_menu_animations": {
							Type:        schema.TypeBool,
							Description: "Enable menu animations",
							Computed:    true,
						},
						"disable_bitmap_caching": {
							Type:        schema.TypeBool,
							Description: "Disable bitmap caching",
							Computed:    true,
						},
						"disable_offscreen_caching": {
							Type:        schema.TypeBool,
							Description: "Disable off-screen caching",
							Computed:    true,
						},
						"disable_glyph_caching": {
							Type:        schema.TypeBool,
							Description: "Disable glyph caching",
							Computed:    true,
						},
						"remote_app": {
							Type:        schema.TypeString,
							Description: "Remote App program",
							Computed:    true,
						},
						"remote_app_working_directory": {
							Type:        schema.TypeString,
							Description: "Remote App working directory",
							Computed:    true,
						},
						"remote_app_parameters": {
							Type:        schema.TypeString,
							Description: "Remote App parameters",
							Computed:    true,
						},
						"preconnection_id": {
							Type:        schema.TypeString,
							Description: "RDP source ID",
							Computed:    true,
						},
						"preconnection_blob": {
							Type:        schema.TypeString,
							Description: "Preconnection BLOB (VM ID)",
							Computed:    true,
						},
						"load_balance_info": {
							Type:        schema.TypeString,
							Description: "Load balance info/cookie",
							Computed:    true,
						},
						"recording_path": {
							Type:        schema.TypeString,
							Description: "Screen recording path",
							Computed:    true,
						},
						"recording_name": {
							Type:        schema.TypeString,
							Description: "Screen recording name",
							Computed:    true,
						},
						"recording_exclude_output": {
							Type:        schema.TypeBool,
							Description: "Exclude graphics/streams",
							Computed:    true,
						},
						"recording_exclude_mouse": {
							Type:        schema.TypeBool,
							Description: "Exclude mouse",
							Computed:    true,
						},
						"recording_include_keys": {
							Type:        schema.TypeBool,
							Description: "Include key events",
							Computed:    true,
						},
						"recording_auto_create_path": {
							Type:        schema.TypeBool,
							Description: "Auto create recording path",
							Computed:    true,
						},
						"sftp_enable": {
							Type:        schema.TypeBool,
							Description: "Enable sftp",
							Computed:    true,
						},
						"sftp_root_directory": {
							Type:        schema.TypeString,
							Description: "File browser root directory",
							Computed:    true,
						},
						"sftp_hostname": {
							Type:        schema.TypeString,
							Description: "SFTP server hostname",
							Computed:    true,
						},
						"sftp_port": {
							Type:        schema.TypeString,
							Description: "SFTP server port",
							Computed:    true,
						},
						"sftp_host_key": {
							Type:        schema.TypeString,
							Description: "SFTP server public host key (Base64)",
							Computed:    true,
						},
						"sftp_username": {
							Type:        schema.TypeString,
							Description: "SFTP server username",
							Computed:    true,
						},
						"sftp_password": {
							Type:        schema.TypeString,
							Description: "SFTP server password",
							Computed:    true,
						},
						"sftp_private_key": {
							Type:        schema.TypeString,
							Description: "SFTP server private key",
							Computed:    true,
						},
						"sftp_passphrase": {
							Type:        schema.TypeString,
							Description: "SFTP server private key passphrase",
							Computed:    true,
						},
						"sftp_upload_directory": {
							Type:        schema.TypeString,
							Description: "SFTP default upload directory",
							Computed:    true,
						},
						"sftp_keepalive_interval": {
							Type:        schema.TypeString,
							Description: "SFTP keepalive interval",
							Computed:    true,
						},
						"sftp_disable_file_download": {
							Type:        schema.TypeBool,
							Description: "Disable file download",
							Computed:    true,
						},
						"sftp_disable_file_upload": {
							Type:        schema.TypeBool,
							Description: "Disable file upload",
							Computed:    true,
						},
						"wol_send_packet": {
							Type:        schema.TypeBool,
							Description: "Send WoL packet",
							Computed:    true,
						},
						"wol_mac_address": {
							Type:        schema.TypeString,
							Description: "MAC address of the remote host",
							Computed:    true,
						},
						"wol_broadcast_address": {
							Type:        schema.TypeString,
							Description: "Broadcast address for WoL packet",
							Computed:    true,
						},
						"wol_boot_wait_time": {
							Type:        schema.TypeString,
							Description: "Host boot wait time",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceConnectionRDPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	identifier := d.Get("identifier").(string)
	path := d.Get("path").(string)

	if path == "" && identifier == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Missing required parameter"),
			Detail:   "Either `identifier` or `path` must be specified",
		})
		return diags
	}

	if path != "" && identifier != "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Identifier and Path are mutually exclusive"),
			Detail:   "Either `identifier` or `path` must be specified but not both",
		})
		return diags
	}

	// get connection
	var connection types.GuacConnection
	if identifier != "" {
		c, err := client.ReadConnection(identifier)
		if err != nil {
			return diag.FromErr(err)
		}
		connection = c
	} else if path != "" {
		c, err := client.ReadConnectionByPath(path)
		if err != nil {
			return diag.FromErr(err)
		}
		connection = c
	}

	check := convertGuacConnectionRDPToResourceData(d, &connection)

	if check.HasError() {
		return check
	}

	d.SetId(connection.Identifier)

	return diags
}
