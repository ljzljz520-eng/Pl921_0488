package review

import (
	"returnbook/internal/storage"
	"testing"
)

func TestReviewApproval(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/db")
	defer s.Close()
	x := New(s)
	if _, e := x.Submit("r1", "v1", "u", "approved", "good"); e != nil || !x.IsApproved("r1") {
		t.Fatal(e)
	}
}
