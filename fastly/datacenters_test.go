package fastly

import (
	"fmt"
	"testing"
)

func TestClient_Datacenter(t *testing.T) {
	t.Parallel()

	var dcs []*Datacenter
	var err error

	record(t, "datacenters/list", func(c *Client) {
		dcs, err = c.ListDatacenters()
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(dcs) == 0 {
		t.Fatal("missing datacenters")
	}

	for _, dc := range dcs {
		fmt.Printf("code: %s\tshield: %s\n", dc.Code, dc.Shield)
	}
}
