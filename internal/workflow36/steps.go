package workflow36

func stepName(i int) string {
	switch i {
	case 0:
		return "intake"
	case 1:
		return "review"
	case 2:
		return "query"
	default:
		return "archive"
	}
}
func allSteps() []string { return []string{stepName(0), stepName(1), stepName(2), stepName(3)} }
