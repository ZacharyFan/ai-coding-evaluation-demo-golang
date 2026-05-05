package returns

import (
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/orders"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/refunds"
)

type Preview struct {
	Eligible    bool
	RefundCents int
	Reason      string
}

func PreviewReturn(request Request) Preview {
	if !orders.CanReturn(request.Order) {
		return Preview{Reason: "order is not returnable"}
	}
	amount, err := refunds.Amount(request.Order.TotalCents, request.RestockingFeeCents)
	if err != nil {
		return Preview{Reason: err.Error()}
	}
	return Preview{Eligible: true, RefundCents: amount}
}
