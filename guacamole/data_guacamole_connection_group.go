package guacamole

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	guac "github.com/techBeck03/guacamole-api-client"
	"github.com/techBeck03/guacamole-api-client/types"
)

func dataSourceConnectionGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConnectionGroupRead,
		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:        schema.TypeString,
				Description: "Identifier of guacamole connection group",
				Optional:    true,
			},
			"path": {
				Type:        schema.TypeString,
				Description: "Identifier of guacamole connection group",
				Optional:    true,
			},
			"parent_identifier": {
				Type:        schema.TypeString,
				Description: "Parent Identifier of guacamole connection group",
				Optional:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Identifier of guacamole connection group",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Identifier of guacamole connection group",
				Computed:    true,
			},
			"active_connections": {
				Type:        schema.TypeInt,
				Description: "Identifier of guacamole connection group",
				Computed:    true,
			},
			"attributes": {
				Type:        schema.TypeList,
				Description: "Attributes of guacamole connection group",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_connections": {
							Type:        schema.TypeString,
							Description: "Maximum number of total simultaneous connections allowed",
							Computed:    true,
						},
						"max_connections_per_user": {
							Type:        schema.TypeString,
							Description: "Maximum number of simultaneous connections allowed per user",
							Computed:    true,
						},
						"enable_session_affinity": {
							Type:        schema.TypeBool,
							Description: "Enable session affinity",
							Computed:    true,
						},
					},
				},
			},
			"member_connection_groups": {
				Type:        schema.TypeList,
				Description: "Member connection groups of a guacamole connection group",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identifier": {
							Type:        schema.TypeString,
							Description: "Identifier of guacamole connection group",
							Optional:    true,
						},
						"parent_identifier": {
							Type:        schema.TypeString,
							Description: "Parent Identifier of guacamole connection group",
							Optional:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Identifier of guacamole connection group",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Identifier of guacamole connection group",
							Computed:    true,
						},
						"active_connections": {
							Type:        schema.TypeInt,
							Description: "Identifier of guacamole connection group",
							Computed:    true,
						},
					},
				},
			},
			"member_connections": {
				Type:        schema.TypeList,
				Description: "Member connections of a guacamole connection group",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identifier": {
							Type:        schema.TypeString,
							Description: "Guacd proxy port",
							Computed:    true,
						},
						"parent_identifier": {
							Type:        schema.TypeString,
							Description: "Parent Identifier of guacamole connection",
							Optional:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Identifier of guacamole connection group",
							Computed:    true,
						},
						"protocol": {
							Type:        schema.TypeString,
							Description: "Protocol type of the guacamole connection",
							Computed:    true,
						},
						"active_connections": {
							Type:        schema.TypeInt,
							Description: "Identifier of guacamole connection group",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceConnectionGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	// get connection group
	var group types.GuacConnectionGroup
	if identifier != "" {
		g, err := client.ReadConnectionGroup(identifier)
		if err != nil {
			return diag.FromErr(err)
		}
		group = g
	} else if path != "" {
		g, err := client.ReadConnectionGroupByPath(path)
		if err != nil {
			return diag.FromErr(err)
		}
		group = g
	}

	check := convertGuacConnectionGroupToResourceData(d, &group)

	if check.HasError() {
		return check
	}

	var memberGroups []interface{}
	for _, group := range group.ChildGroups {
		memberGroups = append(memberGroups, map[string]interface{}{
			"identifier":         group.Identifier,
			"parent_identifier":  group.ParentIdentifier,
			"name":               group.Name,
			"type":               group.Type,
			"active_connections": group.ActiveConnections,
		})
	}

	d.Set("member_connection_groups", memberGroups)

	var memberConnections []interface{}
	for _, connection := range group.ChildConnections {
		memberConnections = append(memberConnections, map[string]interface{}{
			"identifier":         connection.Identifier,
			"parent_identifier":  connection.ParentIdentifier,
			"name":               connection.Name,
			"protocol":           connection.Protocol,
			"active_connections": connection.ActiveConnections,
		})
	}

	d.Set("member_connections", memberConnections)

	d.SetId(group.Identifier)

	return diags
}
