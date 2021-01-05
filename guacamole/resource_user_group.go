package guacamole

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	guac "github.com/techBeck03/guacamole-api-client"
	types "github.com/techBeck03/guacamole-api-client/types"
)

func guacamoleUserGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserGroupCreate,
		ReadContext:   resourceUserGroupRead,
		UpdateContext: resourceUserGroupUpdate,
		DeleteContext: resourceUserGroupDelete,
		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:        schema.TypeString,
				Description: "Identifier of guacamole user group",
				Required:    true,
				ForceNew:    true,
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
			"group_membership": {
				Type:        schema.TypeList,
				Description: "Groups this user group is a member of",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"system_permissions": {
				Type:        schema.TypeList,
				Description: "System permissions assigned to user group",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceUserGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	group, err := convertResourceDataToGuacUserGroup(d)

	if err != nil {
		return diag.FromErr(err)
	}

	err = client.CreateUserGroup(&group)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(group.Identifier)
	resourceUserGroupRead(ctx, d, m)

	return diags
}

func resourceUserGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	identifier := d.Id()
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

	return diags
}

func resourceUserGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	if d.HasChange("username") || d.HasChange("attributes") {
		group, err := convertResourceDataToGuacUserGroup(d)
		if err != nil {
			return diag.FromErr(err)
		}
		err = client.UpdateUserGroup(&group)

		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceUserGroupRead(ctx, d, m)
}

func resourceUserGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	identifier := d.Id()

	err := client.DeleteUserGroup(identifier)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func convertResourceDataToGuacUserGroup(d *schema.ResourceData) (types.GuacUserGroup, error) {
	var group types.GuacUserGroup

	group.Identifier = d.Get("identifier").(string)

	attributeList := d.Get("attributes").([]interface{})

	if len(attributeList) > 0 {
		attributes := attributeList[0].(map[string]interface{})
		group.Attributes = types.GuacUserGroupAttributes{
			Disabled: boolToString(attributes["disabled"].(bool)),
		}
	}

	return group, nil
}

func convertGuacUserGroupToResourceData(d *schema.ResourceData, group *types.GuacUserGroup) error {
	d.Set("identifier", group.Identifier)

	attributes := map[string]interface{}{
		"disabled": stringToBool(group.Attributes.Disabled),
	}

	var attributeList []map[string]interface{}

	attributeList = append(attributeList, attributes)

	d.Set("attributes", attributeList)

	return nil
}
