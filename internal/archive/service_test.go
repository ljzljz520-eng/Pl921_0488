package archive

import (
	"returnbook/internal/storage"
	"testing"
)

func TestArchiveRecord(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/db")
	defer s.Close()
	x := New(s)
	if _, e := x.Archive("a1", "v1", "done"); e != nil {
		t.Fatal(e)
	}
	if _, e := x.Load("a1"); e != nil {
		t.Fatal(e)
	}
}
