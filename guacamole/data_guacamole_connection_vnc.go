package guacamole

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	guac "github.com/techBeck03/guacamole-api-client"
	types "github.com/techBeck03/guacamole-api-client/types"
)

func dataSourceConnectionVNC() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConnectionVNCRead,
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
							Optional:    true,
							Computed:    true,
						},
						"username": {
							Type:        schema.TypeString,
							Description: "Username for vnc connection",
							Required:    true,
						},
						"password": {
							Type:        schema.TypeString,
							Description: "Password for vnc connection",
							Computed:    true,
						},
						"readonly": {
							Type:        schema.TypeBool,
							Description: "Display is readonly",
							Computed:    true,
						},
						"swap_red_blue": {
							Type:        schema.TypeBool,
							Description: "Swap red/blue Components",
							Computed:    true,
						},
						"cursor": {
							Type:        schema.TypeString,
							Description: "Local or remote cursor",
							Computed:    true,
						},
						"color_depth": {
							Type:        schema.TypeString,
							Description: "Color depth",
							Computed:    true,
						},
						"clipboard_encoding": {
							Type:        schema.TypeString,
							Description: "Clipboard encoding",
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
						"destination_host": {
							Type:        schema.TypeString,
							Description: "VNC repeater destination host",
							Computed:    true,
						},
						"destination_port": {
							Type:        schema.TypeString,
							Description: "VN repeater destination port",
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
						"enable_audio": {
							Type:        schema.TypeBool,
							Description: "Enable audio",
							Computed:    true,
						},
						"audio_server_name": {
							Type:        schema.TypeString,
							Description: "Audio server name",
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

func dataSourceConnectionVNCRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	check := convertGuacConnectionVNCToResourceData(d, &connection)

	if check.HasError() {
		return check
	}

	d.SetId(connection.Identifier)

	return diags
}
