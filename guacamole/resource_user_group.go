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
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disabled": {
							Type:        schema.TypeBool,
							Description: "Whether group is disabled",
							Optional:    true,
							Computed:    true,
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
			"connections": {
				Type:        schema.TypeSet,
				Description: "Connections identifiers a user has permission to read",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"connection_groups": {
				Type:        schema.TypeSet,
				Description: "Connection Group identifiers a user has permission to read",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
            "member_groups": {
                Type:        schema.TypeSet,
                Description: "User Group identifiers which are part of this group",
                Optional:    true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "member_users": {
                Type:        schema.TypeSet,
                Description: "Usernames which are part of this group",
                Optional:    true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
		},
                Importer: &schema.ResourceImporter{
                        StateContext: schema.ImportStatePassthroughContext,
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
        memberGroupSet, ok := d.GetOk("member_groups")
        var memberGroups []string
        for _, memberGroup := range memberGroupSet.(*schema.Set).List() {
            memberGroups = append(memberGroups, memberGroup.(string))
        }
        if ok && len(memberGroups) > 0 {
            check := validateGroups(client, memberGroups)
            if check.HasError() {
                diags = append(diags, check...)
                goto Cleanup
            }
            var permissionItems []types.GuacPermissionItem
            for _, group := range memberGroups {
                permissionItems = append(permissionItems, client.NewAddGroupMemberPermission(group))
            }
            err = client.SetUserGroupMemberGroups(group.Identifier, &permissionItems)
            if err != nil {
                diags = append(diags, diag.FromErr(err)...)
                goto Cleanup
            }
        }
    }

    if !diags.HasError() {
        memberUserSet, ok := d.GetOk("member_users")
        var memberUsers []string
        for _, memberUser := range memberUserSet.(*schema.Set).List() {
            memberUsers = append(memberUsers, memberUser.(string))
        }
        if ok && len(memberUsers) > 0 {
            var permissionItems []types.GuacPermissionItem
            for _, user := range memberUsers {
                permissionItems = append(permissionItems, client.NewAddGroupMemberPermission(user))
            }
            err = client.SetUserGroupUsers(group.Identifier, &permissionItems)
            if err != nil {
                diags = append(diags, diag.FromErr(err)...)
                goto Cleanup
            }
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

	if !diags.HasError() {
		var connectionPermissionItems []types.GuacPermissionItem
		connectionSet, ok := d.GetOk("connections")
		var connections []string
		for _, connection := range connectionSet.(*schema.Set).List() {
			connections = append(connections, connection.(string))
		}
		if ok && len(connections) > 0 {
			for _, connection := range connections {
				connectionPermissionItems = append(connectionPermissionItems, client.NewAddConnectionPermission(connection))
			}
		}

		connectionGroupSet, ok := d.GetOk("connection_groups")
		var connectionGroups []string
		for _, connectionGroup := range connectionGroupSet.(*schema.Set).List() {
			connectionGroups = append(connectionGroups, connectionGroup.(string))
		}
		if ok && len(connectionGroups) > 0 {
			for _, connectionGroup := range connectionGroups {
				connectionPermissionItems = append(connectionPermissionItems, client.NewAddConnectionGroupPermission(connectionGroup))
			}
		}

		if len(connectionPermissionItems) > 0 {
			err = client.SetUserGroupPermissions(group.Identifier, &connectionPermissionItems)
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

	// Get connections
	var connections []string
	for connection := range permissions.ConnectionPermissions {
		connections = append(connections, connection)
	}

	d.Set("connections", connections)

	// Get connection groups
	var connectionGroups []string
	for group := range permissions.ConnectionGroupPermissions {
		connectionGroups = append(connectionGroups, group)
	}

	d.Set("connection_groups", connectionGroups)

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

	if d.HasChange("connections") {
		var permissionItems []types.GuacPermissionItem
		old, new := d.GetChange("connections")
		var oldConnections, newConnnections []string

		for _, connection := range old.(*schema.Set).List() {
			oldConnections = append(oldConnections, connection.(string))
		}

		for _, connection := range new.(*schema.Set).List() {
			newConnnections = append(newConnnections, connection.(string))
		}

		removeConnections := sliceDiff(oldConnections, newConnnections, false)
		if len(removeConnections) > 0 {
			for _, connection := range removeConnections {
				permissionItems = append(permissionItems, client.NewRemoveConnectionPermission(connection))
			}
		}

		addConnections := sliceDiff(newConnnections, oldConnections, false)
		if len(addConnections) > 0 {
			for _, connection := range addConnections {
				permissionItems = append(permissionItems, client.NewAddConnectionPermission(connection))
			}
		}
		if len(permissionItems) > 0 {
			err := client.SetUserGroupPermissions(d.Id(), &permissionItems)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if d.HasChange("connection_groups") {
		var permissionItems []types.GuacPermissionItem
		old, new := d.GetChange("connection_groups")
		var oldConnectionGroups, newConnectionGroups []string

		for _, connection := range old.(*schema.Set).List() {
			oldConnectionGroups = append(oldConnectionGroups, connection.(string))
		}

		for _, connection := range new.(*schema.Set).List() {
			newConnectionGroups = append(newConnectionGroups, connection.(string))
		}

		removeConnectionGroups := sliceDiff(oldConnectionGroups, newConnectionGroups, false)
		if len(removeConnectionGroups) > 0 {
			for _, connection := range removeConnectionGroups {
				permissionItems = append(permissionItems, client.NewRemoveConnectionGroupPermission(connection))
			}
		}

		addConnectionGroups := sliceDiff(newConnectionGroups, oldConnectionGroups, false)
		if len(addConnectionGroups) > 0 {
			for _, connection := range addConnectionGroups {
				permissionItems = append(permissionItems, client.NewAddConnectionGroupPermission(connection))
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
