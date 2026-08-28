package workflow36

import (
	"returnbook/internal/archive"
	"returnbook/internal/intake"
	"returnbook/internal/review"
	"returnbook/internal/storage"
	"sync"
)

type Workflow struct {
	Intake  *intake.Service
	Review  *review.Service
	Archive *archive.Service
}

func New(s *storage.Store) *Workflow { return &Workflow{intake.New(s), review.New(s), archive.New(s)} }
func (w *Workflow) Run(ids []string) error {
	done := make(chan string)
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(v string) { defer wg.Done(); done <- v }(id)
	}
	go func() { wg.Wait(); close(done) }()
	for range done {
	}
	return nil
}
