package guacamole

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	guac "github.com/techBeck03/guacamole-api-client"
	types "github.com/techBeck03/guacamole-api-client/types"
)

func guacamoleConnectionGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConnectionGroupCreate,
		ReadContext:   resourceConnectionGroupRead,
		UpdateContext: resourceConnectionGroupUpdate,
		DeleteContext: resourceConnectionGroupDelete,
		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:        schema.TypeString,
				Description: "Identifier of guacamole connection group",
				Computed:    true,
			},
			"parent_identifier": {
				Type:        schema.TypeString,
				Description: "Parent Identifier of guacamole connection group",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of guacamole connection group",
				Required:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type of guacamole connection group",
				Optional:    true,
				Default:     "ORGANIZATIONAL",
				StateFunc: func(val interface{}) string {
					return strings.ToUpper(val.(string))
				},
			},
			"active_connections": {
				Type:        schema.TypeInt,
				Description: "Active connections of guacamole connection group",
				Computed:    true,
			},
			"attributes": {
				Type:        schema.TypeList,
				Description: "Attributes of guacamole connection group",
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_connections": {
							Type:        schema.TypeString,
							Description: "Maximum number of total simultaneous connections allowed",
							Optional:    true,
							Computed:    true,
						},
						"max_connections_per_user": {
							Type:        schema.TypeString,
							Description: "Maximum number of simultaneous connections allowed per user",
							Optional:    true,
							Computed:    true,
						},
						"enable_session_affinity": {
							Type:        schema.TypeBool,
							Description: "Enable session affinity",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceConnectionGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	validate := validateConnectionGroup(d, client)

	if validate.HasError() {
		return validate
	}

	group, check := convertResourceDataToGuacConnectionGroup(d)

	if check.HasError() {
		return check
	}

	err := client.CreateConnectionGroup(&group)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("identifier", group.Identifier)
	d.SetId(group.Identifier)

	if diags.HasError() {
		return diags
	}

	return resourceConnectionGroupRead(ctx, d, m)
}

func resourceConnectionGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	identifier := d.Id()

	group, err := client.ReadConnectionGroup(identifier)

	if err != nil {
		return diag.FromErr(err)
	}

	check := convertGuacConnectionGroupToResourceData(d, &group)
	if check.HasError() {
		return check
	}

	d.SetId(identifier)

	return diags
}

func resourceConnectionGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)
	var diags diag.Diagnostics

	if d.HasChanges("name", "identifier", "parent_identifier", "type", "attributes") {
		validate := validateConnectionGroup(d, client)

		if validate.HasError() {
			return validate
		}

		group, check := convertResourceDataToGuacConnectionGroup(d)

		if check.HasError() {
			return check
		}

		err := client.UpdateConnectionGroup(&group)

		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId(group.Identifier)

	} else {
		d.SetId(d.Id())
	}

	if diags.HasError() {
		return diags
	}

	return resourceConnectionGroupRead(ctx, d, m)
}

func resourceConnectionGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	identifier := d.Id()

	err := client.DeleteConnectionGroup(identifier)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func convertResourceDataToGuacConnectionGroup(d *schema.ResourceData) (types.GuacConnectionGroup, diag.Diagnostics) {
	var group types.GuacConnectionGroup
	var diags diag.Diagnostics

	group.Identifier = d.Get("identifier").(string)
	group.ParentIdentifier = d.Get("parent_identifier").(string)
	group.Name = d.Get("name").(string)
	group.Type = strings.ToUpper(d.Get("type").(string))

	attributeList := d.Get("attributes").([]interface{})

	if len(attributeList) > 0 {
		attributes := attributeList[0].(map[string]interface{})
		group.Attributes = types.GuacConnectionGroupAttributes{
			MaxConnections:        attributes["max_connections"].(string),
			MaxConnectionsPerUser: attributes["max_connections_per_user"].(string),
			EnableSessionAffinity: boolToString(attributes["enable_session_affinity"].(bool)),
		}
	}

	return group, diags
}

func convertGuacConnectionGroupToResourceData(d *schema.ResourceData, group *types.GuacConnectionGroup) diag.Diagnostics {
	d.Set("identifier", group.Identifier)
	d.Set("parent_identifier", group.ParentIdentifier)
	d.Set("name", group.Name)
	d.Set("type", group.Type)

	attributes := map[string]interface{}{
		"max_connections":          group.Attributes.MaxConnections,
		"max_connections_per_user": group.Attributes.MaxConnectionsPerUser,
		"enable_session_affinity":  stringToBool(group.Attributes.EnableSessionAffinity),
	}
	var attributeList []map[string]interface{}

	attributeList = append(attributeList, attributes)

	d.Set("attributes", attributeList)

	return nil
}

func validateConnectionGroup(d *schema.ResourceData, client *guac.Client) diag.Diagnostics {
	var diags diag.Diagnostics

	connectionGroupInterface := types.GuacConnectionGroup{}
	// validate connection group restricted values
	connectionGroupRestricedValueParameters := map[string][]string{
		"type": connectionGroupInterface.ValidTypes(),
	}
	// parameters that need be to forced to upper case
	forceUpperCaseParameters := []string{
		"type",
	}

	for k, v := range connectionGroupRestricedValueParameters {
		var value string
		capitalizationCheck := stringInSlice(forceUpperCaseParameters, []string{k})
		if !capitalizationCheck.HasError() {
			value = strings.ToUpper(d.Get(k).(string))
		} else {
			value = d.Get(k).(string)
		}
		check := stringInSlice(v, []string{value})
		if check.HasError() {
			diags = append(diags, check...)
		}
	}

	// validate attributes
	attributeList := d.Get("attributes").([]interface{})

	stringIntAttributes := []string{
		"max_connections",
		"max_connections_per_user",
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
	}

	return diags
}
