package intake

import (
	"fmt"
	"returnbook/internal/model"
	"returnbook/internal/storage"
	"time"
)

type Service struct{ Store *storage.Store }

func New(s *storage.Store) *Service { return &Service{Store: s} }
func (s *Service) RegisterCustomer(id, name, phone, address string) (model.Customer, error) {
	if id == "" || name == "" {
		return model.Customer{}, fmt.Errorf("missing identity")
	}
	c := model.NewCustomer(id, name, phone, address)
	return c, s.Store.Put("customers", id, c)
}
func (s *Service) ScheduleVisit(id, cid, staff, notes string, at time.Time) (model.Visit, error) {
	if at.IsZero() {
		return model.Visit{}, fmt.Errorf("schedule required")
	}
	v := model.NewVisit(id, cid, staff, notes, at)
	if !v.IsReady() {
		return v, fmt.Errorf("incomplete visit")
	}
	return v, s.Store.Put("visits", id, v)
}
func (s *Service) GetCustomer(id string) (model.Customer, error) {
	var c model.Customer
	return c, s.Store.Get("customers", id, &c)
}
