package sophosutm_test

import (
	"github/DanielRis/sophosutm"
	"testing"
)

var (
	network sophosutm.Network
)

func TestCreateNetwork(t *testing.T) {
	client, err := sophosutm.NewClient(ServerURL, APIUser, APIKey)

	if err != nil {
		t.Fatalf("err: %v", err)
	}

	network.Name = "SophosUTM go client test network name"
	network.Comment = "SophosUTM go client test network comment"
	network.Address = "192.168.200.1"
	network.Netmask = "32"

	network, err := client.CreateNetwork(network)

	if err != nil {
		t.Fatalf("err: %v", err)
	}

	t.Logf("Network Ref: %s", network.Ref)
	t.Logf("Network Name: %s", network.Name)
	t.Logf("Network Comment: %s", network.Comment)
	t.Logf("Network Address: %s", network.Address)
	t.Logf("Network Netmask: %v", network.Netmask)
}

func TestGetNetwork(t *testing.T) {
	client, err := sophosutm.NewClient(ServerURL, APIUser, APIKey)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	network, err := client.GetNetwork("REF_NetNetSophoGoClien")

	if err != nil {
		t.Fatalf("err: %v", err)
	}

	t.Logf("Network Name: %s", network.Name)
}

func TestDeleteNetwork(t *testing.T) {
	client, err := sophosutm.NewClient(ServerURL, APIUser, APIKey)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	err = client.DeleteNetwork("REF_NetNetSophoGoClien")
}
