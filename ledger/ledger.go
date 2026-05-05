package ledger

import "fmt"

type Entry struct {
	ID          string
	AmountCents int
}

type MemoryLedger struct {
	Entries []Entry
	Fail    bool
}

func (l *MemoryLedger) Record(entry Entry) error {
	if l.Fail {
		return fmt.Errorf("ledger unavailable")
	}
	l.Entries = append(l.Entries, entry)
	return nil
}
