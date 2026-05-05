package payment_saga

import (
	"context"
	"testing"
)

func TestSettleInvoiceRetriesTransientGatewayError(t *testing.T) {
	gateway := &transientGateway{}
	ledger := &captureLedger{}
	workflow := Workflow{Gateway: gateway, Ledger: ledger}
	if err := workflow.SettleInvoice(context.Background(), "inv-1", 1000); err != nil {
		t.Fatalf("SettleInvoice returned error: %v", err)
	}
	if gateway.calls != 2 || ledger.records != 1 {
		t.Fatalf("gateway calls=%d ledger records=%d, want 2/1", gateway.calls, ledger.records)
	}
}

type transientGateway struct{ calls int }

func (g *transientGateway) Capture(context.Context, string, int) (string, error) {
	g.calls++
	if g.calls == 1 {
		return "", ErrTransient
	}
	return "cap-1", nil
}
func (g *transientGateway) Refund(context.Context, string, int) (string, error) { return "ref-1", nil }

type captureLedger struct{ records int }

func (l *captureLedger) RecordCapture(context.Context, string, string, int) error {
	l.records++
	return nil
}
func (l *captureLedger) RecordRefund(context.Context, string, string, int) error { return nil }
