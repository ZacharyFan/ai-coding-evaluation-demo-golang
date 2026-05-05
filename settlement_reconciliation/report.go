package settlement_reconciliation

import "sort"

type Report struct {
	Discrepancies []Discrepancy
}

func BuildReport(payouts []Payout, entries []LedgerEntry) Report {
	discrepancies := matchPayoutsToLedger(payouts, entries)
	sort.Slice(discrepancies, func(i, j int) bool { return discrepancies[i].ExternalID < discrepancies[j].ExternalID })
	return Report{Discrepancies: discrepancies}
}

func matchPayoutsToLedger(payouts []Payout, entries []LedgerEntry) []Discrepancy {
	result := []Discrepancy{}
	entryByID := map[string]LedgerEntry{}
	for _, entry := range entries {
		entryByID[entry.ExternalID] = entry
	}
	for _, payout := range payouts {
		entry, ok := entryByID[payout.ExternalID]
		if !ok {
			result = append(result, Discrepancy{ExternalID: payout.ExternalID, Reason: "missing ledger"})
			continue
		}
		if entry.AmountCents != payout.AmountCents {
			result = append(result, Discrepancy{ExternalID: payout.ExternalID, Reason: "amount mismatch"})
		}
	}
	return result
}
