package settlement_reconciliation

import (
	"context"
	"fmt"
	"testing"
)

func TestReconcileReturnsNotificationFailure(t *testing.T) {
	reconciler := Reconciler{Payouts: onePayout{}, Ledger: mismatchLedger{}, Notifications: failingNotification{}}
	_, err := reconciler.Reconcile(context.Background(), "batch-1")
	if err == nil {
		t.Fatal("Reconcile returned nil error; want notification failure")
	}
}

type onePayout struct{}

func (onePayout) ListPayouts(context.Context, string) ([]Payout, error) {
	return []Payout{{ExternalID: "p-1", AmountCents: 100}}, nil
}

type mismatchLedger struct{}

func (mismatchLedger) ListEntries(context.Context, string) ([]LedgerEntry, error) {
	return []LedgerEntry{{ExternalID: "p-1", AmountCents: 90}}, nil
}

type failingNotification struct{}

func (failingNotification) Notify(context.Context, string, string) error {
	return fmt.Errorf("notify failed")
}
