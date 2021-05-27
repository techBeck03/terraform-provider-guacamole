package guacamole

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/techBeck03/guacamole-api-client/types"
)

const (
	connectionsBasePath = "connections"
)

// CreateConnection creates a guacamole connection
func (c *Client) CreateConnection(connection *types.GuacConnection) error {
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.baseURL, connectionsBasePath), connection)

	if err != nil {
		return err
	}

	err = c.Call(request, &connection)
	if err != nil {
		return err
	}
	return nil
}

// ReadConnection gets a connection by identifier
func (c *Client) ReadConnection(identifier string) (types.GuacConnection, error) {
	var ret types.GuacConnection
	var retParams types.GuacConnectionParameters

	// Get connection base details
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.baseURL, connectionsBasePath, url.QueryEscape(identifier)), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	if ret.Identifier != "" {
		// Get connection parameters
		request, err = c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/parameters", c.baseURL, connectionsBasePath, identifier), nil)

		if err != nil {
			return ret, err
		}

		err = c.Call(request, &retParams)
		if err != nil {
			return ret, err
		}
	}

	ret.Parameters = retParams

	return ret, nil
}

// ReadConnectionByPath gets a connection by path (Parent/Name)
func (c *Client) ReadConnectionByPath(path string) (types.GuacConnection, error) {
	var ret types.GuacConnection
	var parentIdentifier string

	splitPath := strings.Split(path, "/")
	groups, err := c.ListConnectionGroups()

	if err != nil {
		return ret, err
	}

	if strings.ToUpper(splitPath[0]) == "ROOT" {
		parentIdentifier = "ROOT"
	} else {
		for _, group := range groups {
			if group.Name == splitPath[0] {
				parentIdentifier = group.Identifier
				break
			}
		}
	}

	if parentIdentifier == "" {
		return ret, fmt.Errorf("No connection group found for parent with name: %s", splitPath[0])
	}

	connections, err := c.ListConnections()

	if err != nil {
		return ret, err
	}

	for _, connection := range connections {
		if (connection.ParentIdentifier == parentIdentifier) && (connection.Name == splitPath[1]) {
			ret, err = c.ReadConnection(connection.Identifier)
			if err != nil {
				return ret, err
			}
			break
		}
	}

	if ret.Identifier == "" {
		return ret, fmt.Errorf("No connection group found with parentIdentifier = %s\tname = %s", parentIdentifier, splitPath[1])
	}

	return ret, nil
}

// UpdateConnection updates a connection by identifier
func (c *Client) UpdateConnection(connection *types.GuacConnection) error {
	request, err := c.CreateJSONRequest(http.MethodPut, fmt.Sprintf("%s/%s/%s", c.baseURL, connectionsBasePath, url.QueryEscape(connection.Identifier)), connection)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// DeleteConnection deletes a connection by identifier
func (c *Client) DeleteConnection(identifier string) error {
	request, err := c.CreateJSONRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s", c.baseURL, connectionsBasePath, url.QueryEscape(identifier)), nil)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListConnections lists all connections
func (c *Client) ListConnections() ([]types.GuacConnection, error) {
	var ret []types.GuacConnection
	var connectionList map[string]types.GuacConnection

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.baseURL, connectionsBasePath), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &connectionList)
	if err != nil {
		return ret, err
	}

	for _, connection := range connectionList {
		ret = append(ret, connection)
	}
	return ret, nil
}
