package model

import "time"

type Customer struct {
	ID, Name, Phone, Address string
	Active                   bool
	CreatedAt                time.Time
}
type Visit struct {
	ID, CustomerID, Staff, Notes string
	ScheduledAt                  time.Time
	Status                       string
}
type Review struct {
	ID, VisitID, Reviewer, Decision, Comment string
	ApprovedAt                               time.Time
}
type ArchiveRecord struct {
	ID, VisitID, Reason string
	ArchivedAt          time.Time
}

func NewCustomer(id, name, phone, address string) Customer {
	return Customer{ID: id, Name: name, Phone: phone, Address: address, Active: true, CreatedAt: time.Now().UTC()}
}
func NewVisit(id, cid, staff, notes string, at time.Time) Visit {
	return Visit{ID: id, CustomerID: cid, Staff: staff, Notes: notes, ScheduledAt: at, Status: "pending"}
}
func (v Visit) IsReady() bool   { return v.CustomerID != "" && v.Staff != "" && v.Notes != "" }
func (r Review) Accepted() bool { return r.Decision == "approved" }
