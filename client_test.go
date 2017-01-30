package sophosutm_test

import (
	"testing"

	"github.com/DanielRis/sophosutm"
)

var (
	ServerURL string
	APIUser   string
	APIKey    string
)

func init() {
	ServerURL = "https://192.168.178.200:4444"
	APIUser = "admin"
	APIKey = "IxWbdQjLHDosImUEsOhmuVxbLNmNnRmV"
}

func TestCreateClient(t *testing.T) {
	_, err := sophosutm.NewClient(ServerURL, APIUser, APIKey)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
}
