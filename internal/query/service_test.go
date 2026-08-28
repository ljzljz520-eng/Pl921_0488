package query

import (
	"returnbook/internal/intake"
	"returnbook/internal/storage"
	"testing"
)

func TestCustomerQuery(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/db")
	defer s.Close()
	intake.New(s).RegisterCustomer("c", "N", "p", "a")
	q := New(s)
	cs, e := q.Customers()
	if e != nil || len(cs) != 1 {
		t.Fatalf("%v %d", e, len(cs))
	}
}
