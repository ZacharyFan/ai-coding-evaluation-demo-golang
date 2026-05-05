package account

import "testing"

func TestAvailableCredit(t *testing.T) {
	got := AvailableCredit(Account{CreditLimitCents: 10000, BalanceCents: 2500})
	if got != 7500 {
		t.Fatalf("AvailableCredit = %d, want 7500", got)
	}
}

func TestCanUseCredit(t *testing.T) {
	if !CanUseCredit(Account{CreditLimitCents: 10000, BalanceCents: 2500}, 5000) {
		t.Fatal("expected account to have usable credit")
	}
}

func TestMaskEmail(t *testing.T) {
	got, err := MaskEmail("alex@example.com")
	if err != nil {
		t.Fatalf("MaskEmail returned error: %v", err)
	}
	if got != "a***@example.com" {
		t.Fatalf("MaskEmail = %q, want masked address", got)
	}
}
