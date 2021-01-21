package guacamole

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	guac "github.com/techBeck03/guacamole-api-client"
	types "github.com/techBeck03/guacamole-api-client/types"
)

func guacamoleConnectionVNC() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConnectionVNCCreate,
		ReadContext:   resourceConnectionVNCRead,
		UpdateContext: resourceConnectionVNCUpdate,
		DeleteContext: resourceConnectionVNCDelete,
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
							Description: "Username for vnc connection",
							Required:    true,
						},
						"password": {
							Type:        schema.TypeString,
							Description: "Password for vnc connection",
							Optional:    true,
							Computed:    true,
						},
						"readonly": {
							Type:        schema.TypeBool,
							Description: "Display is readonly",
							Optional:    true,
							Computed:    true,
						},
						"swap_red_blue": {
							Type:        schema.TypeBool,
							Description: "Swap red/blue Components",
							Optional:    true,
							Computed:    true,
						},
						"cursor": {
							Type:        schema.TypeString,
							Description: "Local or remote cursor",
							Optional:    true,
							Computed:    true,
						},
						"color_depth": {
							Type:        schema.TypeString,
							Description: "Color depth",
							Optional:    true,
							Computed:    true,
						},
						"clipboard_encoding": {
							Type:        schema.TypeString,
							Description: "Clipboard encoding",
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
						"destination_host": {
							Type:        schema.TypeString,
							Description: "VNC repeater destination host",
							Optional:    true,
							Computed:    true,
						},
						"destination_port": {
							Type:        schema.TypeString,
							Description: "VN repeater destination port",
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
						"enable_audio": {
							Type:        schema.TypeBool,
							Description: "Enable audio",
							Optional:    true,
							Computed:    true,
						},
						"audio_server_name": {
							Type:        schema.TypeString,
							Description: "Audio server name",
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

func resourceConnectionVNCRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	identifier := d.Id()

	connection, err := client.ReadConnection(identifier)

	if err != nil {
		return diag.FromErr(err)
	}

	check := convertGuacConnectionVNCToResourceData(d, &connection)
	if check.HasError() {
		return check
	}

	d.SetId(identifier)

	return diags
}

func resourceConnectionVNCCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	validate := validateConnectionVNC(d, client)

	if validate.HasError() {
		return validate
	}

	connection, check := convertResourceDataToGuacConnectionVNC(d)

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

	return resourceConnectionVNCRead(ctx, d, m)
}

func resourceConnectionVNCUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.HasChanges("name", "identifier", "parent_identifier", "attributes", "parameters") {
		validate := validateConnectionVNC(d, client)

		if validate.HasError() {
			return validate
		}

		connection, check := convertResourceDataToGuacConnectionVNC(d)

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

	return resourceConnectionVNCRead(ctx, d, m)
}

func resourceConnectionVNCDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := client.DeleteConnection(d.Id())

	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func convertGuacConnectionVNCToResourceData(d *schema.ResourceData, connection *types.GuacConnection) diag.Diagnostics {
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
		"hostname":                   connection.Parameters.Hostname,
		"port":                       connection.Parameters.Port,
		"username":                   connection.Parameters.Username,
		"password":                   connection.Parameters.Password,
		"readonly":                   stringToBool(connection.Parameters.ReadOnly),
		"swap_red_blue":              stringToBool(connection.Parameters.SwapRedBlue),
		"cursor":                     connection.Parameters.Cursor,
		"color_depth":                connection.Parameters.ColorDepth,
		"clipboard_encoding":         connection.Parameters.ClipboardEncoding,
		"disable_copy":               stringToBool(connection.Parameters.DisableCopy),
		"disable_paste":              stringToBool(connection.Parameters.DisablePaste),
		"destination_host":           connection.Parameters.DestinationHost,
		"destination_port":           connection.Parameters.DestinationPort,
		"recording_path":             connection.Parameters.RecordingPath,
		"recording_name":             connection.Parameters.RecordingName,
		"recording_exclude_output":   stringToBool(connection.Parameters.RecordingExcludeOutput),
		"recording_exclude_mouse":    stringToBool(connection.Parameters.RecordingExcludeMouse),
		"recording_include_keys":     stringToBool(connection.Parameters.RecordingIncludeKeys),
		"recording_auto_create_path": stringToBool(connection.Parameters.CreateRecordingPath),
		"sftp_enable":                stringToBool(connection.Parameters.EnableSFTP),
		"sftp_root_directory":        connection.Parameters.SFTPRootDirectory,
		"sftp_hostname":              connection.Parameters.SFTPHostname,
		"sftp_port":                  connection.Parameters.SFTPPort,
		"sftp_host_key":              connection.Parameters.SFTPHostKey,
		"sftp_username":              connection.Parameters.SFTPUsername,
		"sftp_password":              connection.Parameters.SFTPPassword,
		"sftp_private_key":           connection.Parameters.SFTPPrivateKey,
		"sftp_passphrase":            connection.Parameters.SFTPPassphrase,
		"sftp_upload_directory":      connection.Parameters.SFTPUploadDirectory,
		"sftp_keepalive_interval":    connection.Parameters.SFTPKeepAliveInterval,
		"sftp_disable_file_download": stringToBool(connection.Parameters.SFTPDisableFileDownload),
		"sftp_disable_file_upload":   stringToBool(connection.Parameters.SFTPDisableFileUpload),
		"enable_audio":               stringToBool(connection.Parameters.EnableAudio),
		"audio_server_name":          connection.Parameters.AudioServerName,
		"wol_send_packet":            stringToBool(connection.Parameters.WOLSendPacket),
		"wol_mac_address":            connection.Parameters.WOLMacAddress,
		"wol_broadcast_address":      connection.Parameters.WOLBroadcastAddress,
		"wol_boot_wait_time":         connection.Parameters.WOLBootWaitTime,
	}
	var parameterList []map[string]interface{}

	parameterList = append(parameterList, parameters)

	d.Set("parameters", parameterList)

	return diags
}

func validateConnectionVNC(d *schema.ResourceData, client *guac.Client) diag.Diagnostics {
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
		"destination_port",
		"sftp_port",
		"sftp_keepalive_interval",
		"wol_boot_wait_time",
	}

	var parameterInterface types.GuacConnectionParameters
	restrictedValueParameters := map[string][]string{
		"cursor":             parameterInterface.ValidCursors(),
		"color_depth":        parameterInterface.ValidColorDepths(),
		"clipboard_encoding": parameterInterface.ValidClipboardEncodings(),
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
	}

	return diags
}

func convertResourceDataToGuacConnectionVNC(d *schema.ResourceData) (types.GuacConnection, diag.Diagnostics) {
	var diags diag.Diagnostics
	var connection types.GuacConnection

	connection.Name = d.Get("name").(string)
	connection.Identifier = d.Get("identifier").(string)
	connection.ParentIdentifier = d.Get("parent_identifier").(string)
	connection.Protocol = "vnc"

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
			Hostname:                attributes["hostname"].(string),
			Port:                    attributes["port"].(string),
			Username:                attributes["username"].(string),
			Password:                attributes["password"].(string),
			ReadOnly:                boolToString(attributes["readonly"].(bool)),
			SwapRedBlue:             boolToString(attributes["swap_red_blue"].(bool)),
			Cursor:                  attributes["cursor"].(string),
			ColorDepth:              attributes["color_depth"].(string),
			ClipboardEncoding:       attributes["clipboard_encoding"].(string),
			DisableCopy:             boolToString(attributes["disable_copy"].(bool)),
			DisablePaste:            boolToString(attributes["disable_paste"].(bool)),
			DestinationHost:         attributes["destination_host"].(string),
			DestinationPort:         attributes["destination_port"].(string),
			RecordingPath:           attributes["recording_path"].(string),
			RecordingName:           attributes["recording_name"].(string),
			RecordingExcludeOutput:  boolToString(attributes["recording_exclude_output"].(bool)),
			RecordingExcludeMouse:   boolToString(attributes["recording_exclude_mouse"].(bool)),
			RecordingIncludeKeys:    boolToString(attributes["recording_include_keys"].(bool)),
			CreateRecordingPath:     boolToString(attributes["recording_auto_create_path"].(bool)),
			EnableSFTP:              boolToString(attributes["sftp_enable"].(bool)),
			SFTPHostname:            attributes["sftp_hostname"].(string),
			SFTPPort:                attributes["sftp_port"].(string),
			SFTPHostKey:             attributes["sftp_host_key"].(string),
			SFTPUsername:            attributes["sftp_username"].(string),
			SFTPPassword:            attributes["sftp_password"].(string),
			SFTPPrivateKey:          attributes["sftp_private_key"].(string),
			SFTPPassphrase:          attributes["sftp_passphrase"].(string),
			SFTPRootDirectory:       attributes["sftp_root_directory"].(string),
			SFTPUploadDirectory:     attributes["sftp_upload_directory"].(string),
			SFTPKeepAliveInterval:   attributes["sftp_keepalive_interval"].(string),
			SFTPDisableFileDownload: boolToString(attributes["sftp_disable_file_download"].(bool)),
			SFTPDisableFileUpload:   boolToString(attributes["sftp_disable_file_upload"].(bool)),
			EnableAudio:             boolToString(attributes["enable_audio"].(bool)),
			AudioServerName:         attributes["audio_server_name"].(string),
			WOLSendPacket:           boolToString(attributes["wol_send_packet"].(bool)),
			WOLMacAddress:           attributes["wol_mac_address"].(string),
			WOLBroadcastAddress:     attributes["wol_broadcast_address"].(string),
			WOLBootWaitTime:         attributes["wol_boot_wait_time"].(string),
		}
	}

	return connection, diags
}
