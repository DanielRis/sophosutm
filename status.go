package sophosutm

import "encoding/json"

type Status struct {
	Utm   string `json:"utm"`
	Restd string `json:"restd"`
}

// GetStatus returns some version numbers of the UTM software
func (c *Client) GetStatus() (Status, error) {
	getObjectResponse, err := c.GetObject("/status/version")
	var status Status

	err = json.Unmarshal([]byte(getObjectResponse), &status)
	if err != nil {
		return status, err
	}
	return status, nil
}
