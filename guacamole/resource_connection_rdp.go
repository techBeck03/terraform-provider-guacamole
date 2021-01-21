package guacamole

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	guac "github.com/techBeck03/guacamole-api-client"
	types "github.com/techBeck03/guacamole-api-client/types"
)

func guacamoleConnectionRDP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConnectionRDPCreate,
		ReadContext:   resourceConnectionRDPRead,
		UpdateContext: resourceConnectionRDPUpdate,
		DeleteContext: resourceConnectionRDPDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the guacamole connection",
				Required:    true,
			},
			"identifier": {
				Type:        schema.TypeString,
				Description: "Numeric identifier of the guacamole connection",
				Computed:    true,
			},
			"parent_identifier": {
				Type:        schema.TypeString,
				Description: "Parent identifier of the guacamole connection",
				Optional:    true,
				Default:     "ROOT",
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
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"guacd_hostname": {
							Type:        schema.TypeString,
							Description: "Guacd proxy hostname",
							Optional:    true,
							Computed:    true,
						},
						"guacd_port": {
							Type:        schema.TypeString,
							Description: "Guacd proxy port",
							Optional:    true,
							Computed:    true,
						},
						"guacd_encryption": {
							Type:        schema.TypeString,
							Description: "Guacd proxy encryption type",
							Optional:    true,
							Computed:    true,
						},
						"failover_only": {
							Type:        schema.TypeBool,
							Description: "Use load balancing for failover only",
							Optional:    true,
							Computed:    true,
						},
						"weight": {
							Type:        schema.TypeString,
							Description: "Load balancing connection weight",
							Optional:    true,
							Computed:    true,
						},
						"max_connections": {
							Type:        schema.TypeString,
							Description: "Maximum concurrent total connections",
							Optional:    true,
							Computed:    true,
						},
						"max_connections_per_user": {
							Type:        schema.TypeString,
							Description: "Maximum concurrent connections per user",
							Optional:    true,
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
							Required:    true,
						},
						"port": {
							Type:        schema.TypeString,
							Description: "Port for target connection",
							Optional:    true,
							Computed:    true,
						},
						"username": {
							Type:        schema.TypeString,
							Description: "Username for rdp connection",
							Required:    true,
						},
						"password": {
							Type:        schema.TypeString,
							Description: "Password for rdp connection",
							Optional:    true,
							Computed:    true,
						},
						"domain": {
							Type:        schema.TypeString,
							Description: "Domain name of rdp connection",
							Optional:    true,
							Computed:    true,
						},
						"security_mode": {
							Type:        schema.TypeString,
							Description: "RDP security mode",
							Optional:    true,
							Computed:    true,
						},
						"disable_authentication": {
							Type:        schema.TypeBool,
							Description: "Disable rdp authentication",
							Optional:    true,
							Computed:    true,
						},
						"ignore_cert": {
							Type:        schema.TypeBool,
							Description: "Ignore domain certificate warnings",
							Optional:    true,
							Computed:    true,
						},
						"gateway_hostname": {
							Type:        schema.TypeString,
							Description: "RDS gateway hostname",
							Optional:    true,
							Computed:    true,
						},
						"gateway_port": {
							Type:        schema.TypeString,
							Description: "RDS gateway port",
							Optional:    true,
							Computed:    true,
						},
						"gateway_username": {
							Type:        schema.TypeString,
							Description: "RDS gateway username",
							Optional:    true,
							Computed:    true,
						},
						"gateway_password": {
							Type:        schema.TypeString,
							Description: "RDS gateway password",
							Optional:    true,
							Computed:    true,
						},
						"gateway_domain": {
							Type:        schema.TypeString,
							Description: "RDS gateway domain",
							Optional:    true,
							Computed:    true,
						},
						"initial_program": {
							Type:        schema.TypeString,
							Description: "Initial program for rdp connection",
							Optional:    true,
							Computed:    true,
						},
						"client_name": {
							Type:        schema.TypeString,
							Description: "Client name for rdp connection",
							Optional:    true,
							Computed:    true,
						},
						"keyboard_layout": {
							Type:        schema.TypeString,
							Description: "Keyboard layout for rdp connection",
							Optional:    true,
							Computed:    true,
						},
						"timezone": {
							Type:        schema.TypeString,
							Description: "Timezone/Locale for rdp connection",
							Optional:    true,
							Computed:    true,
						},
						"administrator_console": {
							Type:        schema.TypeBool,
							Description: "Enable administrator console",
							Optional:    true,
							Computed:    true,
						},
						"width": {
							Type:        schema.TypeString,
							Description: "Screen width (px)",
							Optional:    true,
							Computed:    true,
						},
						"height": {
							Type:        schema.TypeString,
							Description: "Screen height (px)",
							Optional:    true,
							Computed:    true,
						},
						"dpi": {
							Type:        schema.TypeString,
							Description: "Resolution (DPI) of rdp connection",
							Optional:    true,
							Computed:    true,
						},
						"color_depth": {
							Type:        schema.TypeString,
							Description: "Color depth of rdp connection",
							Optional:    true,
							Computed:    true,
						},
						"resize_method": {
							Type:        schema.TypeString,
							Description: "Resize method rdp connection",
							Optional:    true,
							Computed:    true,
						},
						"readonly": {
							Type:        schema.TypeBool,
							Description: "Display is readonly",
							Optional:    true,
							Computed:    true,
						},
						"disable_copy": {
							Type:        schema.TypeBool,
							Description: "Disable copying from terminal",
							Optional:    true,
							Computed:    true,
						},
						"disable_paste": {
							Type:        schema.TypeBool,
							Description: "Disable pasting from client",
							Optional:    true,
							Computed:    true,
						},
						"console_audio": {
							Type:        schema.TypeBool,
							Description: "Support audio in console",
							Optional:    true,
							Computed:    true,
						},
						"disable_audio": {
							Type:        schema.TypeBool,
							Description: "Disable audio",
							Optional:    true,
							Computed:    true,
						},
						"enable_audio_input": {
							Type:        schema.TypeBool,
							Description: "Enable audio input (microphone)",
							Optional:    true,
							Computed:    true,
						},
						"enable_printing": {
							Type:        schema.TypeBool,
							Description: "Enable printing",
							Optional:    true,
							Computed:    true,
						},
						"printer_name": {
							Type:        schema.TypeString,
							Description: "Redirected printer name",
							Optional:    true,
							Computed:    true,
						},
						"enable_drive": {
							Type:        schema.TypeBool,
							Description: "Enable drive for device redirection",
							Optional:    true,
							Computed:    true,
						},
						"drive_name": {
							Type:        schema.TypeString,
							Description: "Drive name for device redirection",
							Optional:    true,
							Computed:    true,
						},
						"disable_file_download": {
							Type:        schema.TypeBool,
							Description: "Disable file download for device redirection",
							Optional:    true,
							Computed:    true,
						},
						"disable_file_upload": {
							Type:        schema.TypeBool,
							Description: "Disable file upload for device redirection",
							Optional:    true,
							Computed:    true,
						},
						"drive_path": {
							Type:        schema.TypeString,
							Description: "Drive path for device redirection",
							Optional:    true,
							Computed:    true,
						},
						"create_drive_path": {
							Type:        schema.TypeBool,
							Description: "Create drive path for device redirection",
							Optional:    true,
							Computed:    true,
						},
						"static_channels": {
							Type:        schema.TypeString,
							Description: "Static channel names",
							Optional:    true,
							Computed:    true,
						},
						"enable_wallpaper": {
							Type:        schema.TypeBool,
							Description: "Enable wallpaper",
							Optional:    true,
							Computed:    true,
						},
						"enable_theming": {
							Type:        schema.TypeBool,
							Description: "Enable theming",
							Optional:    true,
							Computed:    true,
						},
						"enable_font_smoothing": {
							Type:        schema.TypeBool,
							Description: "Enable font smoothing",
							Optional:    true,
							Computed:    true,
						},
						"enable_full_window_drag": {
							Type:        schema.TypeBool,
							Description: "Enable full window drag",
							Optional:    true,
							Computed:    true,
						},
						"enable_desktop_composition": {
							Type:        schema.TypeBool,
							Description: "Enable desktop composition",
							Optional:    true,
							Computed:    true,
						},
						"enable_menu_animations": {
							Type:        schema.TypeBool,
							Description: "Enable menu animations",
							Optional:    true,
							Computed:    true,
						},
						"disable_bitmap_caching": {
							Type:        schema.TypeBool,
							Description: "Disable bitmap caching",
							Optional:    true,
							Computed:    true,
						},
						"disable_offscreen_caching": {
							Type:        schema.TypeBool,
							Description: "Disable off-screen caching",
							Optional:    true,
							Computed:    true,
						},
						"disable_glyph_caching": {
							Type:        schema.TypeBool,
							Description: "Disable glyph caching",
							Optional:    true,
							Computed:    true,
						},
						"remote_app": {
							Type:        schema.TypeString,
							Description: "Remote App program",
							Optional:    true,
							Computed:    true,
						},
						"remote_app_working_directory": {
							Type:        schema.TypeString,
							Description: "Remote App working directory",
							Optional:    true,
							Computed:    true,
						},
						"remote_app_parameters": {
							Type:        schema.TypeString,
							Description: "Remote App parameters",
							Optional:    true,
							Computed:    true,
						},
						"preconnection_id": {
							Type:        schema.TypeString,
							Description: "RDP source ID",
							Optional:    true,
							Computed:    true,
						},
						"preconnection_blob": {
							Type:        schema.TypeString,
							Description: "Preconnection BLOB (VM ID)",
							Optional:    true,
							Computed:    true,
						},
						"load_balance_info": {
							Type:        schema.TypeString,
							Description: "Load balance info/cookie",
							Optional:    true,
							Computed:    true,
						},
						"recording_path": {
							Type:        schema.TypeString,
							Description: "Screen recording path",
							Optional:    true,
							Computed:    true,
						},
						"recording_name": {
							Type:        schema.TypeString,
							Description: "Screen recording name",
							Optional:    true,
							Computed:    true,
						},
						"recording_exclude_output": {
							Type:        schema.TypeBool,
							Description: "Exclude graphics/streams",
							Optional:    true,
							Computed:    true,
						},
						"recording_exclude_mouse": {
							Type:        schema.TypeBool,
							Description: "Exclude mouse",
							Optional:    true,
							Computed:    true,
						},
						"recording_include_keys": {
							Type:        schema.TypeBool,
							Description: "Include key events",
							Optional:    true,
							Computed:    true,
						},
						"recording_auto_create_path": {
							Type:        schema.TypeBool,
							Description: "Auto create recording path",
							Optional:    true,
							Computed:    true,
						},
						"sftp_enable": {
							Type:        schema.TypeBool,
							Description: "Enable sftp",
							Optional:    true,
							Computed:    true,
						},
						"sftp_root_directory": {
							Type:        schema.TypeString,
							Description: "File browser root directory",
							Optional:    true,
							Computed:    true,
						},
						"sftp_hostname": {
							Type:        schema.TypeString,
							Description: "SFTP server hostname",
							Optional:    true,
							Computed:    true,
						},
						"sftp_port": {
							Type:        schema.TypeString,
							Description: "SFTP server port",
							Optional:    true,
							Computed:    true,
						},
						"sftp_host_key": {
							Type:        schema.TypeString,
							Description: "SFTP server public host key (Base64)",
							Optional:    true,
							Computed:    true,
						},
						"sftp_username": {
							Type:        schema.TypeString,
							Description: "SFTP server username",
							Optional:    true,
							Computed:    true,
						},
						"sftp_password": {
							Type:        schema.TypeString,
							Description: "SFTP server password",
							Optional:    true,
							Computed:    true,
						},
						"sftp_private_key": {
							Type:        schema.TypeString,
							Description: "SFTP server private key",
							Optional:    true,
							Computed:    true,
						},
						"sftp_passphrase": {
							Type:        schema.TypeString,
							Description: "SFTP server private key passphrase",
							Optional:    true,
							Computed:    true,
						},
						"sftp_upload_directory": {
							Type:        schema.TypeString,
							Description: "SFTP default upload directory",
							Optional:    true,
							Computed:    true,
						},
						"sftp_keepalive_interval": {
							Type:        schema.TypeString,
							Description: "SFTP keepalive interval",
							Optional:    true,
							Computed:    true,
						},
						"sftp_disable_file_download": {
							Type:        schema.TypeBool,
							Description: "Disable file download",
							Optional:    true,
							Computed:    true,
						},
						"sftp_disable_file_upload": {
							Type:        schema.TypeBool,
							Description: "Disable file upload",
							Optional:    true,
							Computed:    true,
						},
						"wol_send_packet": {
							Type:        schema.TypeBool,
							Description: "Send WoL packet",
							Optional:    true,
							Computed:    true,
						},
						"wol_mac_address": {
							Type:        schema.TypeString,
							Description: "MAC address of the remote host",
							Optional:    true,
							Computed:    true,
						},
						"wol_broadcast_address": {
							Type:        schema.TypeString,
							Description: "Broadcast address for WoL packet",
							Optional:    true,
							Computed:    true,
						},
						"wol_boot_wait_time": {
							Type:        schema.TypeString,
							Description: "Host boot wait time",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceConnectionRDPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	identifier := d.Id()

	connection, err := client.ReadConnection(identifier)

	if err != nil {
		return diag.FromErr(err)
	}

	check := convertGuacConnectionRDPToResourceData(d, &connection)
	if check.HasError() {
		return check
	}

	d.SetId(identifier)

	return diags
}

func resourceConnectionRDPCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	validate := validateConnectionRDP(d, client)

	if validate.HasError() {
		return validate
	}

	connection, check := convertResourceDataToGuacConnectionRDP(d)

	if check.HasError() {
		return check
	}

	err := client.CreateConnection(&connection)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("identifier", connection.Identifier)
	d.SetId(connection.Identifier)

	if diags.HasError() {
		return diags
	}

	return resourceConnectionRDPRead(ctx, d, m)
}

func resourceConnectionRDPUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.HasChanges("name", "identifier", "parent_identifier", "attributes", "parameters") {
		validate := validateConnectionRDP(d, client)

		if validate.HasError() {
			return validate
		}

		connection, check := convertResourceDataToGuacConnectionRDP(d)

		if check.HasError() {
			return check
		}

		err := client.UpdateConnection(&connection)

		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId(connection.Identifier)

	} else {
		d.SetId(d.Id())
	}

	if diags.HasError() {
		return diags
	}

	return resourceConnectionRDPRead(ctx, d, m)
}

func resourceConnectionRDPDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := client.DeleteConnection(d.Id())

	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func convertGuacConnectionRDPToResourceData(d *schema.ResourceData, connection *types.GuacConnection) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	d.Set("name", connection.Name)
	d.Set("identifier", connection.Identifier)
	d.Set("parent_identifier", connection.ParentIdentifier)
	d.Set("protocol", connection.Protocol)
	d.Set("active_connections", connection.ActiveConnections)

	attributes := map[string]interface{}{
		"guacd_hostname":           connection.Attributes.GuacdHostname,
		"guacd_port":               connection.Attributes.GuacdPort,
		"guacd_encryption":         connection.Attributes.GuacdEncryption,
		"failover_only":            stringToBool(connection.Attributes.FailoverOnly),
		"weight":                   connection.Attributes.Weight,
		"max_connections":          connection.Attributes.MaxConnections,
		"max_connections_per_user": connection.Attributes.MaxConnectionsPerUser,
	}
	var attributeList []map[string]interface{}

	attributeList = append(attributeList, attributes)

	d.Set("attributes", attributeList)

	parameters := map[string]interface{}{
		"hostname":                     connection.Parameters.Hostname,
		"port":                         connection.Parameters.Port,
		"username":                     connection.Parameters.Username,
		"password":                     connection.Parameters.Password,
		"domain":                       connection.Parameters.Domain,
		"security_mode":                connection.Parameters.Security,
		"disable_authentication":       stringToBool(connection.Parameters.DisableAuthentication),
		"ignore_cert":                  stringToBool(connection.Parameters.IgnoreCert),
		"gateway_hostname":             connection.Parameters.GatewayHostname,
		"gateway_port":                 connection.Parameters.GatewayPort,
		"gateway_username":             connection.Parameters.GatewayUsername,
		"gateway_password":             connection.Parameters.GatewayPassword,
		"gateway_domain":               connection.Parameters.GatewayDomain,
		"initial_program":              connection.Parameters.InitialProgram,
		"client_name":                  connection.Parameters.ClientName,
		"keyboard_layout":              connection.Parameters.KeyboardLayout,
		"timezone":                     connection.Parameters.Timezone,
		"administrator_console":        stringToBool(connection.Parameters.AdministratorConsole),
		"width":                        connection.Parameters.Width,
		"height":                       connection.Parameters.Height,
		"dpi":                          connection.Parameters.DPI,
		"color_depth":                  connection.Parameters.ColorDepth,
		"resize_method":                connection.Parameters.ResizeMethod,
		"readonly":                     stringToBool(connection.Parameters.ReadOnly),
		"disable_copy":                 stringToBool(connection.Parameters.DisableCopy),
		"disable_paste":                stringToBool(connection.Parameters.DisablePaste),
		"console_audio":                stringToBool(connection.Parameters.ConsoleAudio),
		"disable_audio":                stringToBool(connection.Parameters.DisableAudio),
		"enable_audio_input":           stringToBool(connection.Parameters.EnableAudioInput),
		"enable_printing":              stringToBool(connection.Parameters.EnablePrinting),
		"printer_name":                 connection.Parameters.PrinterName,
		"enable_drive":                 stringToBool(connection.Parameters.EnableDrive),
		"drive_name":                   connection.Parameters.DriveName,
		"disable_file_download":        stringToBool(connection.Parameters.DisableFileDownload),
		"disable_file_upload":          stringToBool(connection.Parameters.DisableFileUpload),
		"drive_path":                   connection.Parameters.DrivePath,
		"create_drive_path":            stringToBool(connection.Parameters.CreateDrivePath),
		"static_channels":              connection.Parameters.StaticChannels,
		"enable_wallpaper":             stringToBool(connection.Parameters.EnableWallpaper),
		"enable_theming":               stringToBool(connection.Parameters.EnableTheming),
		"enable_font_smoothing":        stringToBool(connection.Parameters.EnableFontSmoothing),
		"enable_full_window_drag":      stringToBool(connection.Parameters.EnableFullWindowDrag),
		"enable_desktop_composition":   stringToBool(connection.Parameters.EnableDesktopComposition),
		"enable_menu_animations":       stringToBool(connection.Parameters.EnableMenuAnimations),
		"disable_bitmap_caching":       stringToBool(connection.Parameters.DisableBitmapCaching),
		"disable_offscreen_caching":    stringToBool(connection.Parameters.DisableOffscreenCaching),
		"disable_glyph_caching":        stringToBool(connection.Parameters.DisableGlyphCaching),
		"remote_app":                   connection.Parameters.RemoteApp,
		"remote_app_working_directory": connection.Parameters.RemoteAppWorkingDirectory,
		"remote_app_parameters":        connection.Parameters.RemoteAppParameters,
		"preconnection_id":             connection.Parameters.PreconnectionID,
		"preconnection_blob":           connection.Parameters.PreconnectionBLOB,
		"load_balance_info":            connection.Parameters.LoadBalanceInfo,
		"recording_path":               connection.Parameters.RecordingPath,
		"recording_name":               connection.Parameters.RecordingName,
		"recording_exclude_output":     stringToBool(connection.Parameters.RecordingExcludeOutput),
		"recording_exclude_mouse":      stringToBool(connection.Parameters.RecordingExcludeMouse),
		"recording_include_keys":       stringToBool(connection.Parameters.RecordingIncludeKeys),
		"recording_auto_create_path":   stringToBool(connection.Parameters.CreateRecordingPath),
		"sftp_enable":                  stringToBool(connection.Parameters.EnableSFTP),
		"sftp_hostname":                connection.Parameters.SFTPHostname,
		"sftp_port":                    connection.Parameters.SFTPPort,
		"sftp_host_key":                connection.Parameters.SFTPHostKey,
		"sftp_username":                connection.Parameters.SFTPUsername,
		"sftp_password":                connection.Parameters.SFTPPassword,
		"sftp_private_key":             connection.Parameters.SFTPPrivateKey,
		"sftp_passphrase":              connection.Parameters.SFTPPassphrase,
		"sftp_root_directory":          connection.Parameters.SFTPRootDirectory,
		"sftp_upload_directory":        connection.Parameters.SFTPUploadDirectory,
		"sftp_keepalive_interval":      connection.Parameters.SFTPKeepAliveInterval,
		"sftp_disable_file_download":   stringToBool(connection.Parameters.SFTPDisableFileDownload),
		"sftp_disable_file_upload":     stringToBool(connection.Parameters.SFTPDisableFileUpload),
		"wol_send_packet":              stringToBool(connection.Parameters.WOLSendPacket),
		"wol_mac_address":              connection.Parameters.WOLMacAddress,
		"wol_broadcast_address":        connection.Parameters.WOLBroadcastAddress,
		"wol_boot_wait_time":           connection.Parameters.WOLBootWaitTime,
	}
	var parameterList []map[string]interface{}

	parameterList = append(parameterList, parameters)

	d.Set("parameters", parameterList)

	return diags
}

func validateConnectionRDP(d *schema.ResourceData, client *guac.Client) diag.Diagnostics {
	var diags diag.Diagnostics

	// validate attributes
	attributeList := d.Get("attributes").([]interface{})

	stringIntAttributes := []string{
		"guacd_port",
		"weight",
		"max_connections",
		"max_connections_per_user",
	}

	var attributeInterface types.GuacConnectionAttributes
	restrictedValueAttributes := map[string][]string{
		"guacd_encryption": attributeInterface.ValidEncryptionTypes(),
	}

	if len(attributeList) > 0 {
		attributes := attributeList[0].(map[string]interface{})
		// validate string integer values
		for _, v := range stringIntAttributes {
			if attributes[v].(string) != "" {
				_, err := strconv.Atoi(attributes[v].(string))
				if err != nil {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Invalid entry",
						Detail:   fmt.Sprintf("Expected string integer for attribute key: %s but was unable to convert: %s to integer", v, attributes[v].(string)),
					})
				}
			}
		}

		// validate restricted value fields
		for k, v := range restrictedValueAttributes {
			if attributes[k].(string) != "" {
				check := stringInSlice(v, []string{attributes[k].(string)})
				if check.HasError() {
					diags = append(diags, check...)
				}
			}
		}
	}

	// validate parameters
	parameterList := d.Get("parameters").([]interface{})

	stringIntparameters := []string{
		"port",
		"gateway_port",
		"width",
		"height",
		"dpi",
		"preconnection_id",
		"sftp_port",
		"sftp_keepalive_interval",
		"wol_boot_wait_time",
	}

	var parameterInterface types.GuacConnectionParameters
	restrictedValueParameters := map[string][]string{
		"security_mode":   parameterInterface.ValidSecurityModes(),
		"keyboard_layout": parameterInterface.ValidKeyboardLayouts(),
		"color_depth":     parameterInterface.ValidColorDepths(),
		"resize_method":   parameterInterface.ValidResizeMethods(),
	}

	if len(parameterList) > 0 {
		parameters := parameterList[0].(map[string]interface{})
		// validate string integer values
		for _, v := range stringIntparameters {
			if parameters[v].(string) != "" {
				_, err := strconv.Atoi(parameters[v].(string))
				if err != nil {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Invalid entry",
						Detail:   fmt.Sprintf("Expected string integer for parameter key: %s but was unable to convert: %s to integer", v, parameters[v].(string)),
					})
				}
			}
		}

		// validate restricted value fields
		for k, v := range restrictedValueParameters {
			if parameters[k].(string) != "" {
				check := stringInSlice(v, []string{parameters[k].(string)})
				if check.HasError() {
					diags = append(diags, check...)
				}
			}
		}

		// validate timezone
		timezone := parameters["timezone"].(string)
		_, err := time.LoadLocation(timezone)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Invalid timezone",
				Detail:   fmt.Sprintf("Unable to process timezone string: %s", timezone),
			})
		}
	}

	return diags
}

func convertResourceDataToGuacConnectionRDP(d *schema.ResourceData) (types.GuacConnection, diag.Diagnostics) {
	var diags diag.Diagnostics
	var connection types.GuacConnection

	connection.Name = d.Get("name").(string)
	connection.Identifier = d.Get("identifier").(string)
	connection.ParentIdentifier = d.Get("parent_identifier").(string)
	connection.Protocol = "rdp"

	attributeList := d.Get("attributes").([]interface{})

	if len(attributeList) > 0 {
		attributes := attributeList[0].(map[string]interface{})
		connection.Attributes = types.GuacConnectionAttributes{
			GuacdHostname:         attributes["guacd_hostname"].(string),
			GuacdPort:             attributes["guacd_port"].(string),
			GuacdEncryption:       attributes["guacd_encryption"].(string),
			FailoverOnly:          boolToString(attributes["failover_only"].(bool)),
			Weight:                attributes["weight"].(string),
			MaxConnections:        attributes["max_connections"].(string),
			MaxConnectionsPerUser: attributes["max_connections_per_user"].(string),
		}
	}

	parameterList := d.Get("parameters").([]interface{})

	if len(parameterList) > 0 {
		attributes := parameterList[0].(map[string]interface{})
		connection.Parameters = types.GuacConnectionParameters{
			Hostname:                  attributes["hostname"].(string),
			Port:                      attributes["port"].(string),
			Username:                  attributes["username"].(string),
			Password:                  attributes["password"].(string),
			Domain:                    attributes["domain"].(string),
			Security:                  attributes["security_mode"].(string),
			DisableAuthentication:     boolToString(attributes["disable_authentication"].(bool)),
			IgnoreCert:                boolToString(attributes["ignore_cert"].(bool)),
			GatewayHostname:           attributes["gateway_hostname"].(string),
			GatewayPort:               attributes["gateway_port"].(string),
			GatewayUsername:           attributes["gateway_username"].(string),
			GatewayPassword:           attributes["gateway_password"].(string),
			GatewayDomain:             attributes["gateway_domain"].(string),
			InitialProgram:            attributes["initial_program"].(string),
			ClientName:                attributes["client_name"].(string),
			KeyboardLayout:            attributes["keyboard_layout"].(string),
			AdministratorConsole:      boolToString(attributes["administrator_console"].(bool)),
			Timezone:                  attributes["timezone"].(string),
			Width:                     attributes["width"].(string),
			Height:                    attributes["height"].(string),
			DPI:                       attributes["dpi"].(string),
			ColorDepth:                attributes["color_depth"].(string),
			ResizeMethod:              attributes["resize_method"].(string),
			ReadOnly:                  boolToString(attributes["readonly"].(bool)),
			DisableCopy:               boolToString(attributes["disable_copy"].(bool)),
			DisablePaste:              boolToString(attributes["disable_paste"].(bool)),
			ConsoleAudio:              boolToString(attributes["console_audio"].(bool)),
			DisableAudio:              boolToString(attributes["disable_audio"].(bool)),
			EnableAudioInput:          boolToString(attributes["enable_audio_input"].(bool)),
			EnablePrinting:            boolToString(attributes["enable_printing"].(bool)),
			PrinterName:               attributes["printer_name"].(string),
			EnableDrive:               boolToString(attributes["enable_drive"].(bool)),
			DriveName:                 attributes["drive_name"].(string),
			DisableFileDownload:       boolToString(attributes["disable_file_download"].(bool)),
			DisableFileUpload:         boolToString(attributes["disable_file_upload"].(bool)),
			DrivePath:                 attributes["drive_path"].(string),
			CreateDrivePath:           boolToString(attributes["create_drive_path"].(bool)),
			StaticChannels:            attributes["static_channels"].(string),
			EnableWallpaper:           boolToString(attributes["enable_wallpaper"].(bool)),
			EnableTheming:             boolToString(attributes["enable_theming"].(bool)),
			EnableFontSmoothing:       boolToString(attributes["enable_font_smoothing"].(bool)),
			EnableFullWindowDrag:      boolToString(attributes["enable_full_window_drag"].(bool)),
			EnableDesktopComposition:  boolToString(attributes["enable_desktop_composition"].(bool)),
			EnableMenuAnimations:      boolToString(attributes["enable_menu_animations"].(bool)),
			DisableBitmapCaching:      boolToString(attributes["disable_bitmap_caching"].(bool)),
			DisableOffscreenCaching:   boolToString(attributes["disable_offscreen_caching"].(bool)),
			DisableGlyphCaching:       boolToString(attributes["disable_glyph_caching"].(bool)),
			RemoteApp:                 attributes["remote_app"].(string),
			RemoteAppWorkingDirectory: attributes["remote_app_working_directory"].(string),
			RemoteAppParameters:       attributes["remote_app_parameters"].(string),
			PreconnectionID:           attributes["preconnection_id"].(string),
			PreconnectionBLOB:         attributes["preconnection_blob"].(string),
			LoadBalanceInfo:           attributes["load_balance_info"].(string),
			RecordingPath:             attributes["recording_path"].(string),
			RecordingName:             attributes["recording_name"].(string),
			RecordingExcludeOutput:    boolToString(attributes["recording_exclude_output"].(bool)),
			RecordingExcludeMouse:     boolToString(attributes["recording_exclude_mouse"].(bool)),
			RecordingIncludeKeys:      boolToString(attributes["recording_include_keys"].(bool)),
			CreateRecordingPath:       boolToString(attributes["recording_auto_create_path"].(bool)),
			EnableSFTP:                boolToString(attributes["sftp_enable"].(bool)),
			SFTPHostname:              attributes["sftp_hostname"].(string),
			SFTPPort:                  attributes["sftp_port"].(string),
			SFTPHostKey:               attributes["sftp_host_key"].(string),
			SFTPUsername:              attributes["sftp_username"].(string),
			SFTPPassword:              attributes["sftp_password"].(string),
			SFTPPrivateKey:            attributes["sftp_private_key"].(string),
			SFTPPassphrase:            attributes["sftp_passphrase"].(string),
			SFTPRootDirectory:         attributes["sftp_root_directory"].(string),
			SFTPUploadDirectory:       attributes["sftp_upload_directory"].(string),
			SFTPKeepAliveInterval:     attributes["sftp_keepalive_interval"].(string),
			SFTPDisableFileDownload:   boolToString(attributes["sftp_disable_file_download"].(bool)),
			SFTPDisableFileUpload:     boolToString(attributes["sftp_disable_file_upload"].(bool)),
			WOLSendPacket:             boolToString(attributes["wol_send_packet"].(bool)),
			WOLMacAddress:             attributes["wol_mac_address"].(string),
			WOLBroadcastAddress:       attributes["wol_broadcast_address"].(string),
			WOLBootWaitTime:           attributes["wol_boot_wait_time"].(string),
		}
	}

	return connection, diags
}
