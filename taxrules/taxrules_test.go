package taxrules

import "testing"

func TestApplyRule(t *testing.T) {
	got, err := ApplyRule(10000, Rule{Region: "CA", Category: "general", RateBasisPoints: 825})
	if err != nil {
		t.Fatalf("ApplyRule returned error: %v", err)
	}
	if got != 10825 {
		t.Fatalf("ApplyRule = %d, want 10825", got)
	}
}
