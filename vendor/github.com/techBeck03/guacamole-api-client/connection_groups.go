package guacamole

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

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

// flatten Flattens the
func flatten(nested []types.GuacConnectionGroup) ([]types.GuacConnection, []types.GuacConnectionGroup, error) {
	flatConns := []types.GuacConnection{}
	flatGrps := []types.GuacConnectionGroup{}
	for _, groups := range nested {
		flatGrps = append(flatGrps, groups)
		if len(groups.ChildGroups) > 0 {
			conns, subgrps, err := flatten(groups.ChildGroups)
			if err != nil {
				return nil, nil, err
			}
			for _, c := range conns {
				flatConns = append(flatConns, c)
			}
			for _, g := range subgrps {
				flatGrps = append(flatGrps, g)
			}
		}
		if len(groups.ChildConnections) > 0 {
			for _, c := range groups.ChildConnections {
				flatConns = append(flatConns, c)
			}
		}
	}
	return flatConns, flatGrps, nil
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
	var parentIdentifier string

	splitPath := strings.Split(path, "/")
	groups, err := c.ListConnectionGroups()

	if err != nil {
		return ret, err
	}

	if strings.ToUpper(splitPath[0]) == "ROOT" {
		parentIdentifier = "ROOT"
	} else {
		for group := range groups {
			if groups[group].Name == splitPath[0] {
				parentIdentifier = groups[group].Identifier
				break
			}
		}
	}

	if parentIdentifier == "" {
		return ret, fmt.Errorf("No connection group found for parent with name: %s", splitPath[0])
	}

	for _, group := range groups {
		if (group.ParentIdentifier == parentIdentifier) && (group.Name == splitPath[1]) {
			readGroup, err := c.ReadConnectionGroup(group.Identifier)
			if err != nil {
				return ret, err
			}
			ret = readGroup
			break
		}
	}

	if ret.Identifier == "" {
		return ret, fmt.Errorf("No connection group found with parentIdentifier = %s\tname = %s", parentIdentifier, splitPath[1])
	}

	return ret, nil
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
