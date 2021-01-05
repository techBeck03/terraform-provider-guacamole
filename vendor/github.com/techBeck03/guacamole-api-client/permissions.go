package guacamole

import "github.com/techBeck03/guacamole-api-client/types"

const (
	// ConnectionPermissionsBasePath defines base path for connection permissions
	ConnectionPermissionsBasePath = "/connectionPermissions"
	// ConnectionGroupPermissionsBasePath defines base path for connection group permissions
	ConnectionGroupPermissionsBasePath = "/connectionGroupPermissions"
)

var validSystemPermissions = types.StrSlice{
	"ADMINISTER",
	"CREATE_USER",
	"CREATE_CONNECTION",
	"CREATE_CONNECTION_GROUP",
	"CREATE_SHARING_PROFILE",
}

// NewRemoveGroupMemberPermission creates a formatted guac permission item for removing a group member
func (c *Client) NewRemoveGroupMemberPermission(identifier string) types.GuacPermissionItem {
	return types.GuacPermissionItem{
		Op:    "remove",
		Path:  "/",
		Value: identifier,
	}
}

// NewAddGroupMemberPermission creates a formatted guac permission item for adding a group member
func (c *Client) NewAddGroupMemberPermission(identifier string) types.GuacPermissionItem {
	return types.GuacPermissionItem{
		Op:    "add",
		Path:  "/",
		Value: identifier,
	}
}

// NewAddSystemPermission creates a formatted guac permission item for system permissions
func (c *Client) NewAddSystemPermission(permission string) types.GuacPermissionItem {
	return types.GuacPermissionItem{
		Op:    "add",
		Path:  "/systemPermissions",
		Value: permission,
	}
}

// NewRemoveSystemPermission creates a formatted guac permission item for system permissions
func (c *Client) NewRemoveSystemPermission(permission string) types.GuacPermissionItem {
	return types.GuacPermissionItem{
		Op:    "remove",
		Path:  "/systemPermissions",
		Value: permission,
	}
}
