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
							Default:     "",
						},
						"full_name": {
							Type:        schema.TypeString,
							Description: "Full name of user",
							Optional:    true,
							Default:     "",
						},
						"email": {
							Type:        schema.TypeString,
							Description: "Email of user",
							Optional:    true,
							Default:     "",
						},
						"expired": {
							Type:        schema.TypeBool,
							Description: "Whether the user is expired",
							Optional:    true,
							Default:     false,
						},
						"timezone": {
							Type:        schema.TypeString,
							Description: "Timezone of user",
							Optional:    true,
							Default:     "",
						},
						"access_window_start": {
							Type:        schema.TypeString,
							Description: "Access window start time for user",
							Optional:    true,
							Default:     "",
						},
						"access_window_end": {
							Type:        schema.TypeString,
							Description: "Access window end time for user",
							Optional:    true,
							Default:     "",
						},
						"disabled": {
							Type:        schema.TypeBool,
							Description: "Whether account is disabled",
							Optional:    true,
							Default:     false,
						},
						"valid_from": {
							Type:        schema.TypeString,
							Description: "Start date for when user is valid",
							Optional:    true,
							Default:     "",
						},
						"valid_until": {
							Type:        schema.TypeString,
							Description: "End date for when user is valid",
							Optional:    true,
							Default:     "",
						},
					},
				},
			},
			"group_membership": {
				Type:        schema.TypeSet,
				Description: "Groups this user is a member of",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"system_permissions": {
				Type:        schema.TypeSet,
				Description: "System permissions assigned to user",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*guac.Client)

	user, err := convertResourceDataToGuacUser(d)

	if err != nil {
		return diag.FromErr(err)
	}

	err = client.CreateUser(&user)

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
		err = client.SetUserGroupMembership(user.Username, &permissionItems)
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
			err = client.SetUserPermissions(user.Username, &permissionItems)
			if err != nil {
				diags = append(diags, diag.FromErr(err)...)
				goto Cleanup
			}
		}
	}

	d.SetId(user.Username)
	return resourceUserRead(ctx, d, m)
Cleanup:
	d.SetId(user.Username)
	check := resourceUserDelete(ctx, d, m)
	if check.HasError() {
		diags = append(diags, check...)
	}
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

	// Read group membership
	groups, err := client.GetUserGroupMembership(userID)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("group_membership", groups)

	// Read system permissions
	permissions, err := client.GetUserPermissions(userID)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("system_permissions", permissions.SystemPermissions)

	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*guac.Client)

	if d.HasChanges("username", "last_active", "attributes") {
		user, err := convertResourceDataToGuacUser(d)
		if err != nil {
			return diag.FromErr(err)
		}
		err = client.UpdateUser(&user)

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
			err := client.SetUserGroupMembership(d.Id(), &permissionItems)
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
			err := client.SetUserPermissions(d.Id(), &permissionItems)
			if err != nil {
				return diag.FromErr(err)
			}
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

func validateGroups(client *guac.Client, groups []string) diag.Diagnostics {
	var diags diag.Diagnostics
	var invalidUserGroups []string

	userGroups, err := client.ListUserGroups()
	if err != nil {
		return diag.FromErr(err)
	}
	for _, group := range groups {
		matchFlag := false
		for _, g := range userGroups {
			if group == g.Identifier {
				matchFlag = true
				break
			}
		}
		if !matchFlag {
			invalidUserGroups = append(invalidUserGroups, group)
		}
	}
	if len(invalidUserGroups) > 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Invalid user group(s) supplied"),
			Detail:   fmt.Sprintf("The following groups are invalid for group_membership: %s", strings.Join(invalidUserGroups[:], ", ")),
		})
		return diags
	}
	return diags
}
