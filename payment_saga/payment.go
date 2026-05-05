package payment_saga

import (
	"context"
	"fmt"
)

var ErrTransient = fmt.Errorf("transient payment error")

type PaymentGateway interface {
	Capture(context.Context, string, int) (string, error)
	Refund(context.Context, string, int) (string, error)
}

type LedgerClient interface {
	RecordCapture(context.Context, string, string, int) error
	RecordRefund(context.Context, string, string, int) error
}

type Workflow struct {
	Gateway  PaymentGateway
	Ledger   LedgerClient
	captures map[string]string
}

func (w *Workflow) SettleInvoice(ctx context.Context, invoiceID string, amountCents int) error {
	if invoiceID == "" {
		return fmt.Errorf("invoice id is required")
	}
	if w.captures == nil {
		w.captures = map[string]string{}
	}
	captureID, ok := w.captures[invoiceID]
	var err error
	if !ok {
		captureID, err = w.Gateway.Capture(ctx, invoiceID, amountCents)
	}
	if err != nil {
		if err == ErrTransient {
			captureID, err = w.Gateway.Capture(ctx, invoiceID, amountCents)
		}
		if err != nil {
			return err
		}
	}
	if err := w.Ledger.RecordCapture(ctx, invoiceID, captureID, amountCents); err != nil {
		w.captures[invoiceID] = captureID
		return err
	}
	w.captures[invoiceID] = captureID
	return nil
}
