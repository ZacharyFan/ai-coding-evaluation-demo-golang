package returns

import (
	"fmt"

	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/inventory"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/ledger"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/orders"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/refunds"
)

type Request struct {
	Order              orders.Order
	SKU                string
	Quantity           int
	RestockingFeeCents int
}

type Result struct {
	RefundCents int
}

func ProcessReturn(request Request, stock inventory.Snapshot, book *ledger.MemoryLedger) (Result, error) {
	if !orders.CanReturn(request.Order) {
		return Result{}, fmt.Errorf("order is not returnable")
	}
	refundCents, err := refunds.Amount(request.Order.TotalCents, request.RestockingFeeCents)
	if err != nil {
		return Result{}, err
	}
	if err := book.Record(ledger.Entry{ID: request.Order.ID, AmountCents: -refundCents}); err != nil {
		return Result{}, err
	}
	if err := inventory.Restock(stock, request.SKU, request.Quantity); err != nil {
		return Result{}, err
	}
	return Result{RefundCents: refundCents}, nil
}
