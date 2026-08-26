package intake

import (
	"returnbook/internal/storage"
	"testing"
	"time"
)

func TestCustomerAndVisitEntry(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/db")
	defer s.Close()
	x := New(s)
	if _, e := x.RegisterCustomer("c1", "Ana", "1", "A"); e != nil {
		t.Fatal(e)
	}
	if _, e := x.ScheduleVisit("v1", "c1", "Sam", "ok", time.Unix(10, 0)); e != nil {
		t.Fatal(e)
	}
}
