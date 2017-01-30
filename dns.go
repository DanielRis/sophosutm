package sophosutm

import (
	"encoding/json"
	"fmt"
)

// GetDNSAllowedNetworks gets a list of allowed Networks
func (c *Client) GetDNSAllowedNetworks() ([]string, error) {
	getObjectResponse, err := c.GetObject(fmt.Sprintf("/nodes/dns.allowed_networks"))
	var allowedNetworks []string

	err = json.Unmarshal([]byte(getObjectResponse), &allowedNetworks)
	if err != nil {
		return nil, err
	}
	return allowedNetworks, nil
}
