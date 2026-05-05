package customerprofile

import "testing"

func TestValidateProfile(t *testing.T) {
	err := Validate(Profile{ID: "cus-1", Email: "a@example.com"})
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
}

func TestCanSendMarketing(t *testing.T) {
	if !CanSendMarketing(Profile{MarketingOptIn: true}) {
		t.Fatal("opted-in active profile should receive marketing")
	}
}

func TestCanSendMarketingRejectsDeletedProfiles(t *testing.T) {
	if CanSendMarketing(Profile{MarketingOptIn: true, Deleted: true}) {
		t.Fatal("deleted profile should not receive marketing")
	}
}
