package guacamole

import (
	"fmt"
	"net/http"

	"github.com/techBeck03/guacamole-api-client/types"
)

const (
	protocolsBasePath = "schema/protocols"
)

// GetProtocolChoices gets the valid protocol choices for a connection
func (c *Client) GetProtocolChoices() ([]string, error) {
	var ret map[string]types.ProtocolSchema
	var protocols []string

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.baseURL, protocolsBasePath), nil)

	if err != nil {
		return protocols, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return protocols, err
	}

	for protocol := range ret {
		protocols = append(protocols, ret[protocol].Name)
	}
	return protocols, nil
}
