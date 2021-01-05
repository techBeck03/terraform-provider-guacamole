package guacamole

import (
	"fmt"
	"net/http"

	"github.com/techBeck03/guacamole-api-client/types"
)

const (
	userGroupsBasePath = "userGroups"
)

// CreateUserGroup creates a guacamole user group
func (c *Client) CreateUserGroup(userGroup *types.GuacUserGroup) error {
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.baseURL, userGroupsBasePath), userGroup)

	if err != nil {
		return err
	}

	err = c.Call(request, &userGroup)
	if err != nil {
		return err
	}

	return nil
}

// ReadUserGroup gets a user group by name
func (c *Client) ReadUserGroup(name string) (types.GuacUserGroup, error) {
	var ret types.GuacUserGroup

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.baseURL, userGroupsBasePath, name), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// UpdateUserGroup updates a user group by username
func (c *Client) UpdateUserGroup(group *types.GuacUserGroup) error {
	request, err := c.CreateJSONRequest(http.MethodPut, fmt.Sprintf("%s/%s/%s", c.baseURL, userGroupsBasePath, group.Identifier), group)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserGroup deletes a user group by name
func (c *Client) DeleteUserGroup(name string) error {
	request, err := c.CreateJSONRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s", c.baseURL, userGroupsBasePath, name), nil)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListUserGroups lists all user groups
func (c *Client) ListUserGroups() ([]types.GuacUserGroup, error) {
	var ret map[string]types.GuacUserGroup
	var userGroupList []types.GuacUserGroup

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.baseURL, userGroupsBasePath), nil)

	if err != nil {
		return userGroupList, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return userGroupList, err
	}

	for i := range ret {
		userGroupList = append(userGroupList, ret[i])
	}

	return userGroupList, nil
}

// GetUserGroupPermissions gets a user group's permissions by group name
func (c *Client) GetUserGroupPermissions(identifier string) (types.GuacPermissionData, error) {
	var ret types.GuacPermissionData

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/permissions", c.baseURL, userGroupsBasePath, identifier), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// SetUserGroupConnectionPermissions adds connection permissions to a user group
func (c *Client) SetUserGroupConnectionPermissions(group string, permissionItems *[]types.GuacPermissionItem) error {

	request, err := c.CreateJSONRequest(http.MethodPatch, fmt.Sprintf("%s/%s/%s/permissions", c.baseURL, userGroupsBasePath, group), permissionItems)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// GetUserGroupParentGroups retrieves parent groups of a given group
func (c *Client) GetUserGroupParentGroups(group string) ([]string, error) {
	var ret []string
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/userGroups", c.baseURL, userGroupsBasePath, group), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// GetUserGroupMemberGroups retrieves member groups of a given group
func (c *Client) GetUserGroupMemberGroups(group string) ([]string, error) {
	var ret []string
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/memberUserGroups", c.baseURL, userGroupsBasePath, group), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// SetUserGroupParentGroups defines parent groups of a given group
func (c *Client) SetUserGroupParentGroups(group string, permissionItems *[]types.GuacPermissionItem) error {
	request, err := c.CreateJSONRequest(http.MethodPatch, fmt.Sprintf("%s/%s/%s/userGroups", c.baseURL, userGroupsBasePath, group), permissionItems)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// SetUserGroupMemberGroups defines child member groups of a given group
func (c *Client) SetUserGroupMemberGroups(group string, permissionItems *[]types.GuacPermissionItem) error {
	request, err := c.CreateJSONRequest(http.MethodPatch, fmt.Sprintf("%s/%s/%s/memberUserGroups", c.baseURL, userGroupsBasePath, group), permissionItems)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// SetUserGroupPermissions defines child member groups of a given group
func (c *Client) SetUserGroupPermissions(group string, permissionItems *[]types.GuacPermissionItem) error {
	request, err := c.CreateJSONRequest(http.MethodPatch, fmt.Sprintf("%s/%s/%s/permissions", c.baseURL, userGroupsBasePath, group), permissionItems)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// GetUserGroupUsers retrieves member groups of a given group
func (c *Client) GetUserGroupUsers(group string) ([]string, error) {
	var ret []string
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/memberUsers", c.baseURL, userGroupsBasePath, group), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// SetUserGroupUsers defines users of a given group
func (c *Client) SetUserGroupUsers(group string, permissionItems *[]types.GuacPermissionItem) error {
	request, err := c.CreateJSONRequest(http.MethodPatch, fmt.Sprintf("%s/%s/%s/memberUsers", c.baseURL, userGroupsBasePath, group), permissionItems)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}
