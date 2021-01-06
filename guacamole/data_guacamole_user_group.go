package guacamole

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	guac "github.com/techBeck03/guacamole-api-client"
)

func dataSourceUserGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserGroupRead,
		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:        schema.TypeString,
				Description: "Identifier of guacamole user group",
				Required:    true,
			},
			"attributes": {
				Type:        schema.TypeList,
				Description: "Attributes of guacamole user group",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disabled": {
							Type:        schema.TypeBool,
							Description: "Whether group is disabled",
							Optional:    true,
						},
					},
				},
			},
			"parent_groups": {
				Type:        schema.TypeList,
				Description: "Member groups of a guacamole user group",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"member_groups": {
				Type:        schema.TypeList,
				Description: "Member groups of a guacamole user group",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"member_users": {
				Type:        schema.TypeList,
				Description: "Member users of a guacamole user group",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"system_permissions": {
				Type:        schema.TypeList,
				Description: "Member users of a guacamole user group",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceUserGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	identifier := d.Get("identifier").(string)

	group, err := client.ReadUserGroup(identifier)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Error reading guacamole user: %s", identifier),
			Detail:   err.Error(),
		})

		return diags
	}

	err = convertGuacUserGroupToResourceData(d, &group)

	if err != nil {
		return diag.FromErr(err)
	}

	parentGroups, err := client.GetUserGroupParentGroups(identifier)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("parent_groups", parentGroups)

	memberGroups, err := client.GetUserGroupMemberGroups(identifier)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("member_groups", memberGroups)

	members, err := client.GetUserGroupUsers(identifier)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("member_users", members)

	permissions, err := client.GetUserGroupPermissions(identifier)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("system_permissions", permissions.SystemPermissions)

	d.SetId(identifier)

	return diags
}