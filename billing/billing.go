package billing

type Invoice struct {
	AccountID   string
	AmountCents int
	Paid        bool
}

type Store struct {
	Invoices []Invoice
}

func (s *Store) Create(invoice Invoice) {
	s.Invoices = append(s.Invoices, invoice)
}
