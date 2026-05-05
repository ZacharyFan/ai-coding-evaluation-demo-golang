package subscriptions

import (
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/account"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/billing"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/entitlement"
	"testing"
)

func TestActivateCreatesInvoiceAndEntitlement(t *testing.T) {
	invoices := &billing.Store{}
	entitlements := &entitlement.Store{}
	err := Activate(ActivationRequest{Account: account.Account{ID: "acct-1"}, Plan: "pro", Amount: 2000}, invoices, entitlements)
	if err != nil {
		t.Fatalf("Activate returned error: %v", err)
	}
	if len(invoices.Invoices) != 1 || !entitlements.IsActive("acct-1") {
		t.Fatalf("activation side effects invoices=%d entitlement=%v, want 1/true", len(invoices.Invoices), entitlements.IsActive("acct-1"))
	}
}
