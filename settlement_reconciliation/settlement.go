package settlement_reconciliation

import (
	"context"
	"fmt"
)

type Payout struct {
	ExternalID  string
	AmountCents int
}

type LedgerEntry struct {
	ExternalID  string
	AmountCents int
}

type Discrepancy struct {
	ExternalID string
	Reason     string
}

type PayoutClient interface {
	ListPayouts(context.Context, string) ([]Payout, error)
}

type LedgerClient interface {
	ListEntries(context.Context, string) ([]LedgerEntry, error)
}

type NotificationClient interface {
	Notify(context.Context, string, string) error
}

type Reconciler struct {
	Payouts       PayoutClient
	Ledger        LedgerClient
	Notifications NotificationClient
	notified      map[string]struct{}
}

func (r *Reconciler) Reconcile(ctx context.Context, batchID string) ([]Discrepancy, error) {
	payouts, err := r.Payouts.ListPayouts(ctx, batchID)
	if err != nil {
		return nil, err
	}
	entries, err := r.Ledger.ListEntries(ctx, batchID)
	if err != nil {
		return nil, err
	}
	discrepancies := []Discrepancy{}
	entryByID := map[string]LedgerEntry{}
	for _, entry := range entries {
		entryByID[entry.ExternalID] = entry
	}
	for _, payout := range payouts {
		entry, ok := entryByID[payout.ExternalID]
		if !ok {
			discrepancies = append(discrepancies, Discrepancy{ExternalID: payout.ExternalID, Reason: "missing ledger"})
			continue
		}
		if entry.AmountCents != payout.AmountCents {
			discrepancies = append(discrepancies, Discrepancy{ExternalID: payout.ExternalID, Reason: "amount mismatch"})
		}
	}
	for _, discrepancy := range discrepancies {
		if err := r.Notifications.Notify(ctx, batchID, fmt.Sprintf("%s:%s", discrepancy.ExternalID, discrepancy.Reason)); err != nil {
			return discrepancies, err
		}
	}
	return discrepancies, nil
}
