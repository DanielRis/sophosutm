package sophosutm

import (
	"encoding/json"
	"fmt"
)


type Host struct {
	Locked     string   `json:"_locked"`
	Ref        string   `json:"_ref"`
	Type       string   `json:"_type"`
	Name       string   `json:"name"`
	Comment    string   `json:"comment"`
	Address    string   `json:"address"`
	Address6   string   `json:"address6"`
	Hostnames  []string `json:"hostnames"`
	DUIDS      []string `json:"duids"`
	MACS       []string `json:"macs"`
	Interface  string   `json:"interface"`
	Resolved   bool     `json:"resolved"`
	Resolved6  bool     `json:"resolved6"`
	ReverseDNS bool     `json:"reverse_dns"`
}

// GetHost gets a host defintion
func (c *Client) GetHost(ref string) (Host, error) {
	getObjectResponse, err := c.GetObject(fmt.Sprintf("/objects/network/host/%s", ref))
	var host Host

	err = json.Unmarshal([]byte(getObjectResponse), &host)
	if err != nil {
		return host, err
	}
	return host, nil
}

// CreateHost creates a new host
func (c *Client) CreateHost(host Host) (Host, error) {
	reqBody, _ := json.Marshal(
		host,
	)
	createObjectResponse, err := c.CreateObject("/objects/network/host/", reqBody)
	if err != nil {
		return host, err
	}

	err = json.Unmarshal([]byte(createObjectResponse), &host)
	if err != nil {
		return host, err
	}

	return host, nil
}

// DeleteHost deletes a host
func (c *Client) DeleteHost(ref string) error {
	_, err := c.DeleteObject(fmt.Sprintf("/objects/network/host/%s", ref))

	if err != nil {
		return err
	}

	return nil
}
