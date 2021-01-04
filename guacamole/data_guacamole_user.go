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
				Optional:    true,
			},
			"attributes": {
				Type:        schema.TypeList,
				Description: "Attributes of guacamole user",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"organizational_role": {
							Type:        schema.TypeString,
							Description: "Organizational role of user",
							Optional:    true,
						},
						"full_name": {
							Type:        schema.TypeString,
							Description: "Full name of user",
							Optional:    true,
						},
						"email": {
							Type:        schema.TypeString,
							Description: "Email of user",
							Optional:    true,
						},
						"expired": {
							Type:        schema.TypeBool,
							Description: "Whether the user is expired",
							Optional:    true,
						},
						"timezone": {
							Type:        schema.TypeString,
							Description: "Timezone of user",
							Optional:    true,
						},
						"access_window_start": {
							Type:        schema.TypeString,
							Description: "Access window start time for user",
							Optional:    true,
						},
						"access_window_end": {
							Type:        schema.TypeString,
							Description: "Access window end time for user",
							Optional:    true,
						},
						"disabled": {
							Type:        schema.TypeBool,
							Description: "Whether account is disabled",
							Optional:    true,
						},
						"valid_from": {
							Type:        schema.TypeString,
							Description: "Start date for when user is valid",
							Optional:    true,
						},
						"valid_until": {
							Type:        schema.TypeString,
							Description: "End date for when user is valid",
							Optional:    true,
						},
					},
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

	err = convertGuacUserToResourceData(d, &user)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
