package query

import (
	"returnbook/internal/model"
	"returnbook/internal/storage"
)

type Service struct{ Store *storage.Store }

func New(s *storage.Store) *Service { return &Service{Store: s} }
func (s *Service) Customers() ([]model.Customer, error) {
	raw, e := s.Store.List("customers")
	if e != nil {
		return nil, e
	}
	out := make([]model.Customer, 0, len(raw))
	for _, d := range raw {
		var c model.Customer
		if modelErr := jsonUnmarshal(d, &c); modelErr == nil {
			out = append(out, c)
		}
	}
	return out, nil
}
func jsonUnmarshal(d []byte, v any) error { return storageDecode(d, v) }
