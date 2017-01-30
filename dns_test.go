package sophosutm_test

import (
	"github/DanielRis/sophosutm"
	"testing"
)

func TestGetDNSAllowedNetworks(t *testing.T) {
	client, err := sophosutm.NewClient(ServerURL, APIUser, APIKey)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	dnsAllowedNetworks, err := client.GetDNSAllowedNetworks()

	if err != nil {
		t.Fatalf("err: %v", err)
	}

	t.Logf("Host Name: %v", dnsAllowedNetworks)
}
