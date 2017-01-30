package sophosutm

import (
	"encoding/json"
	"fmt"
)

type Network struct {
	Locked    string      `json:"_locked"`
	Ref       string      `json:"_ref"`
	Type      string      `json:"_type"`
	Name      string      `json:"name"`
	Comment   string      `json:"comment"`
	Address   string      `json:"address"`
	Netmask   json.Number `json:"netmask"`
	Address6  string      `json:"address6"`
	Netmask6  json.Number `json:"netmask6"`
	Interface string      `json:"interface"`
	Resolved  bool        `json:"resolved"`
	Resolved6 bool        `json:"resolved6"`
}

// GetNetwork gets a network defintion
func (c *Client) GetNetwork(ref string) (Network, error) {
	getObjectResponse, err := c.GetObject(fmt.Sprintf("/objects/network/network/%s", ref))
	var network Network

	err = json.Unmarshal([]byte(getObjectResponse), &network)
	if err != nil {
		return network, err
	}
	return network, nil
}

// CreateNetwork creates a new network
func (c *Client) CreateNetwork(network Network) (Network, error) {
	reqBody, _ := json.Marshal(
		network,
	)
	createObjectResponse, err := c.CreateObject("/objects/network/network/", reqBody)
	if err != nil {
		return network, err
	}

	err = json.Unmarshal([]byte(createObjectResponse), &network)
	if err != nil {
		return network, err
	}

	return network, nil
}

// DeleteNetwork deletes a network
func (c *Client) DeleteNetwork(ref string) error {
	_, err := c.DeleteObject(fmt.Sprintf("/objects/network/network/%s", ref))

	if err != nil {
		return err
	}

	return nil
}
