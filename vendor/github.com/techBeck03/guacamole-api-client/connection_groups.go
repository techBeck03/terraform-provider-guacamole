package guacamole

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/techBeck03/guacamole-api-client/types"
)

const (
	connectionGroupsBasePath = "connectionGroups"
)

// GetConnectionTree gets the connection tree starting from ROOT
func (c *Client) GetConnectionTree(identifier string) (types.GuacConnectionGroup, error) {
	var ret types.GuacConnectionGroup
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/tree", c.baseURL, connectionGroupsBasePath, identifier), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// getPathTree generates a map of all connection paths
func (c *Client) getPathTree(nested types.GuacConnectionGroup, results *types.GuacConnectionGroupPathTree) error {
	for _, group := range nested.ChildGroups {
		if nested.Path != "" {
			group.Path = fmt.Sprintf("%s/%s", nested.Path, group.Name)
		} else {
			group.Path = group.Name
		}
		results.Groups[group.Identifier] = group.Path
		err := c.getPathTree(group, results)
		if err != nil {
			return err
		}
	}

	for _, connection := range nested.ChildConnections {
		if nested.Name == "ROOT" {
			results.Connections[connection.Identifier] = connection.Name
		} else {
			results.Connections[connection.Identifier] = fmt.Sprintf("%s/%s", nested.Path, connection.Name)
		}
	}

	return nil
}

// CreateConnectionGroup creates a guacamole connection group
func (c *Client) CreateConnectionGroup(group *types.GuacConnectionGroup) error {
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.baseURL, connectionGroupsBasePath), group)

	if err != nil {
		return err
	}

	err = c.Call(request, &group)
	if err != nil {
		return err
	}
	return nil
}

// ReadConnectionGroup gets a connection group by identifier
func (c *Client) ReadConnectionGroup(identifier string) (types.GuacConnectionGroup, error) {
	var ret types.GuacConnectionGroup
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.baseURL, connectionGroupsBasePath, url.QueryEscape(identifier)), nil)
	if err != nil {
		return ret, err
	}
	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	connectionTree, err := c.GetConnectionTree(identifier)

	if err != nil {
		return ret, err
	}

	for _, group := range connectionTree.ChildGroups {
		ret.ChildGroups = append(ret.ChildGroups, types.GuacConnectionGroup{
			Name:              group.Name,
			Identifier:        group.Identifier,
			ParentIdentifier:  group.ParentIdentifier,
			Type:              group.Type,
			ActiveConnections: group.ActiveConnections,
		})
	}

	ret.ChildConnections = connectionTree.ChildConnections

	return ret, nil
}

// ReadConnectionGroupByPath gets a connection group by path (Parent/Name)
func (c *Client) ReadConnectionGroupByPath(path string) (types.GuacConnectionGroup, error) {
	var ret types.GuacConnectionGroup

	groups, err := c.GetConnectionTree("ROOT")

	if err != nil {
		return ret, err
	}

	var tree types.GuacConnectionGroupPathTree
	tree.Connections = make(map[string]string)
	tree.Groups = make(map[string]string)
	err = c.getPathTree(groups, &tree)

	if err != nil {
		return ret, err
	}

	for i, p := range tree.Groups {
		if p == path {
			grp, err := c.ReadConnectionGroup(i)
			grp.Path = path
			if err != nil {
				return ret, err
			}
			return grp, nil
		}
	}

	return ret, fmt.Errorf("no connection group found with path: %s", path)
}

// GetConnectionGroupPathById gets a connection group path by identifier
func (c *Client) GetConnectionGroupPathById(identifier string) (string, error) {
	groups, err := c.GetConnectionTree("ROOT")
	if err != nil {
		return "", err
	}

	var tree types.GuacConnectionGroupPathTree
	tree.Connections = make(map[string]string)
	tree.Groups = make(map[string]string)
	err = c.getPathTree(groups, &tree)
	if err != nil {
		return "", err
	}
	return tree.Groups[identifier], nil
}

// UpdateConnectionGroup updates a connection group by identifier
func (c *Client) UpdateConnectionGroup(group *types.GuacConnectionGroup) error {
	request, err := c.CreateJSONRequest(http.MethodPut, fmt.Sprintf("%s/%s/%s", c.baseURL, connectionGroupsBasePath, url.QueryEscape(group.Identifier)), group)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// DeleteConnectionGroup deletes a connection group by identifier
func (c *Client) DeleteConnectionGroup(identifier string) error {
	request, err := c.CreateJSONRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s", c.baseURL, connectionGroupsBasePath, url.QueryEscape(identifier)), nil)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListConnectionGroups lists all connections
func (c *Client) ListConnectionGroups() ([]types.GuacConnectionGroup, error) {
	var ret []types.GuacConnectionGroup
	var connectionGroupList map[string]types.GuacConnectionGroup

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.baseURL, connectionGroupsBasePath), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &connectionGroupList)
	if err != nil {
		return ret, err
	}

	for _, group := range connectionGroupList {
		ret = append(ret, group)
	}

	return ret, nil
}
