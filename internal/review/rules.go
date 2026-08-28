package review

func canApprove(reviewer, comment string) bool { return reviewer != "" && comment != "" }
func outcome(decision string) string {
	if decision == "approved" {
		return "accepted"
	}
	return "needs_changes"
}
