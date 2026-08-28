package model

func ValidVisitStatus(s string) bool {
	switch s {
	case "pending", "completed", "cancelled":
		return true
	default:
		return false
	}
}
func ValidDecision(s string) bool { return s == "approved" || s == "rejected" }
