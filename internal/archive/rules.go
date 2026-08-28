package archive

func validReason(r string) bool { return len(r) >= 3 }
func archiveState(reason string) string {
	if validReason(reason) {
		return "archived"
	}
	return "pending"
}
