package review

import (
	"fmt"
	"returnbook/internal/model"
	"returnbook/internal/storage"
	"time"
)

type Service struct{ Store *storage.Store }

func New(s *storage.Store) *Service { return &Service{Store: s} }
func (s *Service) Submit(id, visit, reviewer, decision, comment string) (model.Review, error) {
	if decision != "approved" && decision != "rejected" {
		return model.Review{}, fmt.Errorf("invalid decision")
	}
	r := model.Review{ID: id, VisitID: visit, Reviewer: reviewer, Decision: decision, Comment: comment, ApprovedAt: time.Now().UTC()}
	return r, s.Store.Put("reviews", id, r)
}
func (s *Service) Load(id string) (model.Review, error) {
	var r model.Review
	return r, s.Store.Get("reviews", id, &r)
}
func (s *Service) IsApproved(id string) bool { r, e := s.Load(id); return e == nil && r.Accepted() }
