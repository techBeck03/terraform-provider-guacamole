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
							Default:     false,
						},
					},
				},
			},
			"group_membership": {
				Type:        schema.TypeSet,
				Description: "Groups this user group is a member of",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"system_permissions": {
				Type:        schema.TypeSet,
				Description: "System permissions assigned to user group",
				Optional:    true,
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

	groupMembershipSet, ok := d.GetOk("group_membership")
	var groupMembership []string
	for _, group := range groupMembershipSet.(*schema.Set).List() {
		groupMembership = append(groupMembership, group.(string))
	}
	if ok && len(groupMembership) > 0 {
		check := validateGroups(client, groupMembership)
		if check.HasError() {
			diags = append(diags, check...)
			goto Cleanup
		}
		var permissionItems []types.GuacPermissionItem
		for _, group := range groupMembership {
			permissionItems = append(permissionItems, client.NewAddGroupMemberPermission(group))
		}
		err = client.SetUserGroupMemberGroups(group.Identifier, &permissionItems)
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
			goto Cleanup
		}
	}

	if !diags.HasError() {
		systemPermissionsSet, ok := d.GetOk("system_permissions")
		var systemPermissions []string
		for _, group := range systemPermissionsSet.(*schema.Set).List() {
			systemPermissions = append(systemPermissions, group.(string))
		}
		if ok && len(systemPermissions) > 0 {
			check := stringInSlice(types.SystemPermissions{}.ValidChoices(), systemPermissions)
			if check.HasError() {
				diags = append(diags, check...)
				goto Cleanup
			}
			var permissionItems []types.GuacPermissionItem
			for _, permission := range systemPermissions {
				permissionItems = append(permissionItems, client.NewAddSystemPermission(permission))
			}
			err = client.SetUserGroupPermissions(group.Identifier, &permissionItems)
			if err != nil {
				diags = append(diags, diag.FromErr(err)...)
				goto Cleanup
			}
		}
	}

	d.SetId(group.Identifier)
	resourceUserGroupRead(ctx, d, m)

	return diags
Cleanup:
	d.SetId(group.Identifier)
	check := resourceUserGroupDelete(ctx, d, m)
	if check.HasError() {
		diags = append(diags, check...)
	}
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

	// Read group membership
	groups, err := client.GetUserGroupMemberGroups(identifier)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("group_membership", groups)

	// Read system permissions
	permissions, err := client.GetUserGroupPermissions(identifier)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("system_permissions", permissions.SystemPermissions)

	return diags
}

func resourceUserGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	if d.HasChanges("username", "attributes") {
		group, err := convertResourceDataToGuacUserGroup(d)
		if err != nil {
			return diag.FromErr(err)
		}
		err = client.UpdateUserGroup(&group)

		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("group_membership") {
		var permissionItems []types.GuacPermissionItem
		var oldGroups, newGroups []string
		old, new := d.GetChange("group_membership")
		for _, group := range old.(*schema.Set).List() {
			oldGroups = append(oldGroups, group.(string))
		}

		for _, group := range new.(*schema.Set).List() {
			newGroups = append(newGroups, group.(string))
		}

		removeGroups := sliceDiff(oldGroups, newGroups, false)
		if len(removeGroups) > 0 {
			for _, group := range removeGroups {
				permissionItems = append(permissionItems, client.NewRemoveGroupMemberPermission(group))
			}
		}

		addGroups := sliceDiff(newGroups, oldGroups, false)
		if len(addGroups) > 0 {
			check := validateGroups(client, addGroups)
			if check.HasError() {
				return check
			}
			check = checkForDuplicates(addGroups)
			if check.HasError() {
				return check
			}
			for _, group := range addGroups {
				permissionItems = append(permissionItems, client.NewAddGroupMemberPermission(group))
			}
		}
		if len(permissionItems) > 0 {
			err := client.SetUserGroupMemberGroups(d.Id(), &permissionItems)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if d.HasChange("system_permissions") {
		var permissionItems []types.GuacPermissionItem
		old, new := d.GetChange("system_permissions")
		var oldPermissions, newPermissions []string

		for _, permission := range old.(*schema.Set).List() {
			oldPermissions = append(oldPermissions, permission.(string))
		}

		for _, permission := range new.(*schema.Set).List() {
			newPermissions = append(newPermissions, permission.(string))
		}

		removePermissions := sliceDiff(oldPermissions, newPermissions, false)
		if len(removePermissions) > 0 {
			for _, permission := range removePermissions {
				permissionItems = append(permissionItems, client.NewRemoveSystemPermission(permission))
			}
		}

		addPermissions := sliceDiff(newPermissions, oldPermissions, false)
		if len(addPermissions) > 0 {
			check := stringInSlice(types.SystemPermissions{}.ValidChoices(), addPermissions)
			if check.HasError() {
				return check
			}
			for _, permission := range addPermissions {
				permissionItems = append(permissionItems, client.NewAddSystemPermission(permission))
			}
		}
		if len(permissionItems) > 0 {
			err := client.SetUserGroupPermissions(d.Id(), &permissionItems)
			if err != nil {
				return diag.FromErr(err)
			}
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
