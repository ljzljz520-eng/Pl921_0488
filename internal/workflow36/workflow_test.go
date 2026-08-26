package workflow36

import (
	"returnbook/internal/storage"
	"testing"
)

func TestWorkflow36BusinessInvariant(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/db")
	defer s.Close()
	if e := New(s).Run([]string{"1", "2", "3"}); e != nil {
		t.Fatal(e)
	}
}
