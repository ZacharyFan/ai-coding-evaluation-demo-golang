package payment_saga

import "context"

type RefundResult struct {
	RefundID string
}

func (w *Workflow) RefundInvoice(ctx context.Context, invoiceID string, amountCents int) (RefundResult, error) {
	refundID, err := w.Gateway.Refund(ctx, invoiceID, amountCents)
	if err != nil {
		return RefundResult{}, err
	}
	if err := w.Ledger.RecordRefund(ctx, invoiceID, refundID, amountCents); err != nil {
		return RefundResult{RefundID: refundID}, err
	}
	return RefundResult{RefundID: refundID}, nil
}
