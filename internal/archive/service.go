package archive

import (
	"fmt"
	"returnbook/internal/model"
	"returnbook/internal/storage"
	"time"
)

type Service struct{ Store *storage.Store }

func New(s *storage.Store) *Service { return &Service{Store: s} }
func (s *Service) Archive(id, visit, reason string) (model.ArchiveRecord, error) {
	if id == "" || visit == "" {
		return model.ArchiveRecord{}, fmt.Errorf("missing archive data")
	}
	a := model.ArchiveRecord{ID: id, VisitID: visit, Reason: reason, ArchivedAt: time.Now().UTC()}
	return a, s.Store.Put("archives", id, a)
}
func (s *Service) Load(id string) (model.ArchiveRecord, error) {
	var a model.ArchiveRecord
	return a, s.Store.Get("archives", id, &a)
}
