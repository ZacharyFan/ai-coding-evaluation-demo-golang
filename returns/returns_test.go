package returns

import (
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/inventory"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/ledger"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/orders"
	"testing"
)

func TestProcessReturnRecordsLedgerAndRestocksOnSuccess(t *testing.T) {
	stock := inventory.Snapshot{"sku-1": 2}
	book := &ledger.MemoryLedger{}
	result, err := ProcessReturn(Request{Order: orders.Order{ID: "ord-1", Status: orders.StatusDelivered, TotalCents: 5000}, SKU: "sku-1", Quantity: 1, RestockingFeeCents: 500}, stock, book)
	if err != nil {
		t.Fatalf("ProcessReturn returned error: %v", err)
	}
	if result.RefundCents != 4500 || stock["sku-1"] != 3 || len(book.Entries) != 1 || book.Entries[0].AmountCents != -4500 {
		t.Fatalf("return result=%#v stock=%d ledger=%#v, want refund 4500 stock 3 ledger -4500", result, stock["sku-1"], book.Entries)
	}
}
