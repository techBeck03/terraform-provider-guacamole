package guacamole

import (
	"context"
	"fmt"
	"log"
	"strconv"

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
	log.Printf("[DEBUG] Looking for guacamole user: %s\n", username)

	user, err := client.ReadUser(username)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Error reading guacamole user: %s", username),
			Detail:   err.Error(),
		})

		return diags
	}

	d.Set("username", user.Username)
	d.Set("last_active", strconv.Itoa(user.LastActive))

	attributes := map[string]interface{}{
		"organizational_role": user.Attributes.GuacOrganizationalRole,
		"full_name":           user.Attributes.GuacFullName,
		"expired":             stringToBool(user.Attributes.Expired),
		"timezone":            user.Attributes.Timezone,
		"access_window_start": user.Attributes.AccessWindowStart,
		"access_window_end":   user.Attributes.AccessWindowEnd,
		"disabled":            stringToBool(user.Attributes.Disabled),
		"valid_from":          user.Attributes.ValidFrom,
		"valid_until":         user.Attributes.ValidUntil,
	}

	var attributeList []map[string]interface{}

	attributeList = append(attributeList, attributes)

	d.Set("attributes", attributeList)
	d.SetId(user.Username)

	return diags
}
