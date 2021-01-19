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

func guacamoleConnectionKubernetes() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConnectionKubernetesCreate,
		ReadContext:   resourceConnectionKubernetesRead,
		UpdateContext: resourceConnectionKubernetesUpdate,
		DeleteContext: resourceConnectionKubernetesDelete,
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
						"use_ssl": {
							Type:        schema.TypeBool,
							Description: "Use SSL/TLS",
							Optional:    true,
							Computed:    true,
						},
						"ignore_cert": {
							Type:        schema.TypeBool,
							Description: "Ignore certificate errors",
							Optional:    true,
							Computed:    true,
						},
						"ca_cert": {
							Type:        schema.TypeString,
							Description: "Certificate authority certificate",
							Optional:    true,
							Computed:    true,
						},
						"pod": {
							Type:        schema.TypeString,
							Description: "Pod name",
							Optional:    true,
							Computed:    true,
						},
						"namespace": {
							Type:        schema.TypeString,
							Description: "Namespace name",
							Optional:    true,
							Computed:    true,
						},
						"container": {
							Type:        schema.TypeString,
							Description: "Container name",
							Optional:    true,
							Computed:    true,
						},
						"client_cert": {
							Type:        schema.TypeString,
							Description: "Client certificate",
							Optional:    true,
							Computed:    true,
						},
						"client_key": {
							Type:        schema.TypeString,
							Description: "Client key",
							Optional:    true,
							Computed:    true,
						},
						"color_scheme": {
							Type:        schema.TypeString,
							Description: "Display color scheme",
							Optional:    true,
							Computed:    true,
						},
						"font_name": {
							Type:        schema.TypeString,
							Description: "Display font name",
							Optional:    true,
							Computed:    true,
						},
						"font_size": {
							Type:        schema.TypeString,
							Description: "Display font size",
							Optional:    true,
							Computed:    true,
						},
						"max_scrollback_size": {
							Type:        schema.TypeString,
							Description: "Display maximum scrollback",
							Optional:    true,
							Computed:    true,
						},
						"readonly": {
							Type:        schema.TypeBool,
							Description: "Display is readonly",
							Optional:    true,
							Computed:    true,
						},
						"backspace": {
							Type:        schema.TypeString,
							Description: "Backspace key sends",
							Optional:    true,
							Computed:    true,
						},
						"typescript_path": {
							Type:        schema.TypeString,
							Description: "Typescript path",
							Optional:    true,
							Computed:    true,
						},
						"typescript_name": {
							Type:        schema.TypeString,
							Description: "Typescript name",
							Optional:    true,
							Computed:    true,
						},
						"typescript_auto_create_path": {
							Type:        schema.TypeBool,
							Description: "Automatically create typescript path",
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
					},
				},
			},
		},
	}
}

func resourceConnectionKubernetesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	identifier := d.Id()

	connection, err := client.ReadConnection(identifier)

	if err != nil {
		return diag.FromErr(err)
	}

	check := convertGuacConnectionKubernetesToResourceData(d, &connection)
	if check.HasError() {
		return check
	}

	d.SetId(identifier)

	return diags
}

func resourceConnectionKubernetesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	validate := validateConnectionKubernetes(d, client)

	if validate.HasError() {
		return validate
	}

	connection, check := convertResourceDataToGuacConnectionKubernetes(d)

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

	return resourceConnectionKubernetesRead(ctx, d, m)
}

func resourceConnectionKubernetesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.HasChanges("name", "identifier", "parent_identifier", "attributes", "parameters") {
		validate := validateConnectionKubernetes(d, client)

		if validate.HasError() {
			return validate
		}

		connection, check := convertResourceDataToGuacConnectionKubernetes(d)

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

	return resourceConnectionKubernetesRead(ctx, d, m)
}

func resourceConnectionKubernetesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := client.DeleteConnection(d.Id())

	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func convertGuacConnectionKubernetesToResourceData(d *schema.ResourceData, connection *types.GuacConnection) diag.Diagnostics {
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
		"hostname":                    connection.Parameters.Hostname,
		"port":                        connection.Parameters.Port,
		"use_ssl":                     stringToBool(connection.Parameters.UseSSL),
		"ignore_cert":                 stringToBool(connection.Parameters.IgnoreCert),
		"ca_cert":                     connection.Parameters.CACert,
		"namespace":                   connection.Parameters.Namespace,
		"pod":                         connection.Parameters.Pod,
		"container":                   connection.Parameters.Container,
		"client_cert":                 connection.Parameters.ClientCert,
		"client_key":                  connection.Parameters.ClientKey,
		"color_scheme":                connection.Parameters.ColorScheme,
		"font_name":                   connection.Parameters.FontName,
		"font_size":                   connection.Parameters.FontSize,
		"max_scrollback_size":         connection.Parameters.Scrollback,
		"readonly":                    stringToBool(connection.Parameters.ReadOnly),
		"backspace":                   connection.Parameters.Backspace,
		"typescript_path":             connection.Parameters.TypescriptPath,
		"typescript_name":             connection.Parameters.TypescriptName,
		"typescript_auto_create_path": stringToBool(connection.Parameters.CreateTypescriptPath),
		"recording_path":              connection.Parameters.RecordingPath,
		"recording_name":              connection.Parameters.RecordingName,
		"recording_exclude_output":    stringToBool(connection.Parameters.RecordingExcludeOutput),
		"recording_exclude_mouse":     stringToBool(connection.Parameters.RecordingExcludeMouse),
		"recording_include_keys":      stringToBool(connection.Parameters.RecordingIncludeKeys),
		"recording_auto_create_path":  stringToBool(connection.Parameters.CreateRecordingPath),
	}
	var parameterList []map[string]interface{}

	parameterList = append(parameterList, parameters)

	d.Set("parameters", parameterList)

	return diags
}

func validateConnectionKubernetes(d *schema.ResourceData, client *guac.Client) diag.Diagnostics {
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
		"max_scrollback_size",
		"wol_boot_wait_time",
	}

	var parameterInterface types.GuacConnectionParameters
	restrictedValueParameters := map[string][]string{
		"color_scheme": parameterInterface.ValidColorSchemes(),
		"font_size":    parameterInterface.ValidFontSizes(),
		"backspace":    parameterInterface.ValidBackspaceCodes(),
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

func convertResourceDataToGuacConnectionKubernetes(d *schema.ResourceData) (types.GuacConnection, diag.Diagnostics) {
	var diags diag.Diagnostics
	var connection types.GuacConnection

	connection.Name = d.Get("name").(string)
	connection.Identifier = d.Get("identifier").(string)
	connection.ParentIdentifier = d.Get("parent_identifier").(string)
	connection.Protocol = "kubernetes"

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
			Hostname:               attributes["hostname"].(string),
			Port:                   attributes["port"].(string),
			UseSSL:                 boolToString(attributes["use_ssl"].(bool)),
			IgnoreCert:             boolToString(attributes["ignore_cert"].(bool)),
			CACert:                 attributes["ca_cert"].(string),
			Namespace:              attributes["namespace"].(string),
			Pod:                    attributes["pod"].(string),
			Container:              attributes["container"].(string),
			ClientCert:             attributes["client_cert"].(string),
			ClientKey:              attributes["client_key"].(string),
			ColorScheme:            attributes["color_scheme"].(string),
			FontName:               attributes["font_name"].(string),
			FontSize:               attributes["font_size"].(string),
			Scrollback:             attributes["max_scrollback_size"].(string),
			ReadOnly:               boolToString(attributes["readonly"].(bool)),
			Backspace:              attributes["backspace"].(string),
			TypescriptPath:         attributes["typescript_path"].(string),
			TypescriptName:         attributes["typescript_name"].(string),
			CreateTypescriptPath:   boolToString(attributes["typescript_auto_create_path"].(bool)),
			RecordingPath:          attributes["recording_path"].(string),
			RecordingName:          attributes["recording_name"].(string),
			RecordingExcludeOutput: boolToString(attributes["recording_exclude_output"].(bool)),
			RecordingExcludeMouse:  boolToString(attributes["recording_exclude_mouse"].(bool)),
			RecordingIncludeKeys:   boolToString(attributes["recording_include_keys"].(bool)),
			CreateRecordingPath:    boolToString(attributes["recording_auto_create_path"].(bool)),
		}
	}

	return connection, diags
}
