package intake

import "strings"

func normalize(v string) string  { return strings.TrimSpace(v) }
func validPhone(v string) bool   { return len(normalize(v)) >= 3 }
func validAddress(v string) bool { return normalize(v) != "" }
