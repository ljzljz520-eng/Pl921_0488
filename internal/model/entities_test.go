package model

import (
	"testing"
	"time"
)

func TestVisitReadiness(t *testing.T) {
	if !NewVisit("v", "c", "s", "n", time.Now()).IsReady() {
		t.Fatal()
	}
}
