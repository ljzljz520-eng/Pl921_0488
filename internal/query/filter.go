package query

import "returnbook/internal/model"

func ActiveCustomers(in []model.Customer) []model.Customer {
	out := make([]model.Customer, 0)
	for _, c := range in {
		if c.Active {
			out = append(out, c)
		}
	}
	return out
}
func FindCustomer(in []model.Customer, id string) (model.Customer, bool) {
	for _, c := range in {
		if c.ID == id {
			return c, true
		}
	}
	return model.Customer{}, false
}
