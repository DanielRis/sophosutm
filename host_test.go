package sophosutm_test

import (
	"github/DanielRis/sophosutm"
	"testing"
)

var (
	host sophosutm.Host
)

func TestCreateHost(t *testing.T) {
	client, err := sophosutm.NewClient(ServerURL, APIUser, APIKey)

	if err != nil {
		t.Fatalf("err: %v", err)
	}

	host.Address = "192.168.200.1"

	host, err := client.CreateHost(host)

	if err != nil {
		t.Fatalf("err: %v", err)
	}

	t.Logf("Host Ref: %s", host.Ref)
	t.Logf("Host Name: %s", host.Name)
	t.Logf("Host Comment: %s", host.Comment)
	t.Logf("Host Address: %s", host.Address)
}

func TestGetHost(t *testing.T) {
	client, err := sophosutm.NewClient(ServerURL, APIUser, APIKey)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	host, err := client.GetHost("REF_NetworkLocalhost")

	if err != nil {
		t.Fatalf("err: %v", err)
	}

	t.Logf("Host Name: %s", host.Name)
}

func TestDeleteHost(t *testing.T) {
	client, err := sophosutm.NewClient(ServerURL, APIUser, APIKey)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	err = client.DeleteHost("REF_NetNetSophoGoClien")
}
