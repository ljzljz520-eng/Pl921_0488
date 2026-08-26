package workflow36

import (
	"returnbook/internal/storage"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/db")
	defer s.Close()
	if e := New(s).Run([]string{"a"}); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/db")
	defer s.Close()
	if e := New(s).Run([]string{"b"}); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/db")
	defer s.Close()
	if e := New(s).Run([]string{"c"}); e != nil {
		t.Fatal(e)
	}
}
