package guacamole

import (
	"fmt"
	"net/http"
	"net/url"

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

	for i, p := range tree.Connections {
		if p == path {
			conn, err := c.ReadConnection(i)
			conn.Path = path
			if err != nil {
				return ret, err
			}
			return conn, nil
		}
	}

	return ret, fmt.Errorf("no connection found with path: %s", path)
}

// GetConnectionPathById gets a connection group path by identifier
func (c *Client) GetConnectionPathById(identifier string) (string, error) {
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
	return tree.Connections[identifier], nil
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
