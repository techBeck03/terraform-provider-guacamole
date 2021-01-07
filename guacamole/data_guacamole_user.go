package guacamole

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	guac "github.com/techBeck03/guacamole-api-client"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Description: "Username of guacamole user",
				Required:    true,
			},
			"last_active": {
				Type:        schema.TypeString,
				Description: "Epoch time string of last user activity",
				Computed:    true,
			},
			"attributes": {
				Type:        schema.TypeList,
				Description: "Attributes of guacamole user",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"organizational_role": {
							Type:        schema.TypeString,
							Description: "Organizational role of user",
							Computed:    true,
						},
						"full_name": {
							Type:        schema.TypeString,
							Description: "Full name of user",
							Computed:    true,
						},
						"email": {
							Type:        schema.TypeString,
							Description: "Email of user",
							Computed:    true,
						},
						"expired": {
							Type:        schema.TypeBool,
							Description: "Whether the user is expired",
							Computed:    true,
						},
						"timezone": {
							Type:        schema.TypeString,
							Description: "Timezone of user",
							Computed:    true,
						},
						"access_window_start": {
							Type:        schema.TypeString,
							Description: "Access window start time for user",
							Computed:    true,
						},
						"access_window_end": {
							Type:        schema.TypeString,
							Description: "Access window end time for user",
							Computed:    true,
						},
						"disabled": {
							Type:        schema.TypeBool,
							Description: "Whether account is disabled",
							Computed:    true,
						},
						"valid_from": {
							Type:        schema.TypeString,
							Description: "Start date for when user is valid",
							Computed:    true,
						},
						"valid_until": {
							Type:        schema.TypeString,
							Description: "End date for when user is valid",
							Computed:    true,
						},
					},
				},
			},
			"group_membership": {
				Type:        schema.TypeSet,
				Description: "Groups this user is a member of",
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"system_permissions": {
				Type:        schema.TypeSet,
				Description: "System permissions assigned to user",
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	username := d.Get("username").(string)

	user, err := client.ReadUser(username)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Error reading guacamole user: %s", username),
			Detail:   err.Error(),
		})

		return diags
	}

	groups, err := client.GetUserGroupMembership(username)

	if err != nil {
		return diag.FromErr(err)
	}

	err = convertGuacUserToResourceData(d, &user)

	d.Set("group_membership", groups)

	if err != nil {
		return diag.FromErr(err)
	}

	permissions, err := client.GetUserPermissions(username)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("system_permissions", permissions.SystemPermissions)

	d.SetId(username)

	return diags
}
