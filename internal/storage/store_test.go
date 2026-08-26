package storage

import "testing"

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/db"
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.Put("customers", "x", map[string]string{"id": "x"}); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	var v map[string]string
	if e = s.Get("customers", "x", &v); e != nil || v["id"] != "x" {
		t.Fatalf("%v %v", e, v)
	}
}
