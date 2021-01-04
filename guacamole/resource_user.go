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

func guacamoleUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Description: "Username of guacamole user",
				Required:    true,
				ForceNew:    true,
			},
			"last_active": {
				Type:        schema.TypeString,
				Description: "Epoch time string of last user activity",
				Computed:    true,
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

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	user, err := convertResourceDataToGuacUser(d)

	if err != nil {
		return diag.FromErr(err)
	}

	err = client.CreateUser(&user)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(user.Username)
	resourceUserRead(ctx, d, m)

	return diags
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	userID := d.Id()
	user, err := client.ReadUser(userID)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Error reading guacamole user: %s", userID),
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

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	if d.HasChange("username") || d.HasChange("last_active") || d.HasChange("attributes") {
		user, err := convertResourceDataToGuacUser(d)
		if err != nil {
			return diag.FromErr(err)
		}
		err = client.UpdateUser(&user)

		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceUserRead(ctx, d, m)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	userID := d.Id()

	err := client.DeleteUser(userID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func convertResourceDataToGuacUser(d *schema.ResourceData) (types.GuacUser, error) {
	var user types.GuacUser

	user.Username = d.Get("username").(string)
	if d.Get("last_active").(string) != "" {
		lastActive, err := strconv.Atoi(d.Get("last_active").(string))
		if err != nil {
			return user, err
		}
		user.LastActive = lastActive
	}

	attributeList := d.Get("attributes").([]interface{})

	if len(attributeList) > 0 {
		attributes := attributeList[0].(map[string]interface{})
		user.Attributes = types.GuacUserAttributes{
			GuacOrganizationalRole: attributes["organizational_role"].(string),
			GuacFullName:           attributes["full_name"].(string),
			Email:                  attributes["email"].(string),
			Expired:                boolToString(attributes["expired"].(bool)),
			Timezone:               attributes["timezone"].(string),
			AccessWindowStart:      attributes["access_window_start"].(string),
			AccessWindowEnd:        attributes["access_window_end"].(string),
			Disabled:               boolToString(attributes["disabled"].(bool)),
			ValidFrom:              attributes["valid_from"].(string),
			ValidUntil:             attributes["valid_until"].(string),
		}
	}

	return user, nil
}

func convertGuacUserToResourceData(d *schema.ResourceData, user *types.GuacUser) error {
	d.Set("username", user.Username)
	d.Set("last_active", strconv.Itoa(user.LastActive))

	attributes := map[string]interface{}{
		"organizational_role": user.Attributes.GuacOrganizationalRole,
		"full_name":           user.Attributes.GuacFullName,
		"email":               user.Attributes.Email,
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

	return nil
}
