package sophosutm_test

import (
	"github/DanielRis/sophosutm"
	"testing"
)

func TestGetStatus(t *testing.T) {
	client, err := sophosutm.NewClient(ServerURL, APIUser, APIKey)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	status, err := client.GetStatus()

	if err != nil {
		t.Fatalf("err: %v", err)
	}

	t.Logf("UTM Version: %s", status.Utm)
	t.Logf("RestD Version: %s", status.Restd)
}
