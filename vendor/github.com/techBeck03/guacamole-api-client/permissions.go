package guacamole

import (
	"fmt"

	"github.com/techBeck03/guacamole-api-client/types"
)

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

// NewRemoveConnectionPermission creates a formatted guac permission item for removing a user connection permission
func (c *Client) NewRemoveConnectionPermission(identifier string) types.GuacPermissionItem {
	return types.GuacPermissionItem{
		Op:    "remove",
		Path:  fmt.Sprintf("%s/%s", ConnectionPermissionsBasePath, identifier),
		Value: "READ",
	}
}

// NewAddConnectionPermission creates a formatted guac permission item for adding a user connection permission
func (c *Client) NewAddConnectionPermission(identifier string) types.GuacPermissionItem {
	return types.GuacPermissionItem{
		Op:    "add",
		Path:  fmt.Sprintf("%s/%s", ConnectionPermissionsBasePath, identifier),
		Value: "READ",
	}
}

// NewRemoveConnectionGroupPermission creates a formatted guac permission item for removing a user connection permission
func (c *Client) NewRemoveConnectionGroupPermission(identifier string) types.GuacPermissionItem {
	return types.GuacPermissionItem{
		Op:    "remove",
		Path:  fmt.Sprintf("%s/%s", ConnectionGroupPermissionsBasePath, identifier),
		Value: "READ",
	}
}

// NewAddConnectionGroupPermission creates a formatted guac permission item for adding a user connection permission
func (c *Client) NewAddConnectionGroupPermission(identifier string) types.GuacPermissionItem {
	return types.GuacPermissionItem{
		Op:    "add",
		Path:  fmt.Sprintf("%s/%s", ConnectionGroupPermissionsBasePath, identifier),
		Value: "READ",
	}
}
