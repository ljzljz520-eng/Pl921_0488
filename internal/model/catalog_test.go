package model

import "testing"

func TestCatalogLookup(t *testing.T) {
	if CatalogLabel("S0001") == "unknown service" {
		t.Fatal()
	}
}
