package subscriptions

import (
	"fmt"

	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/account"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/billing"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/entitlement"
)

type ActivationRequest struct {
	Account account.Account
	Plan    string
	Amount  int
}

func Activate(request ActivationRequest, invoices *billing.Store, entitlements *entitlement.Store) error {
	if request.Account.ID == "" || request.Plan == "" {
		return fmt.Errorf("account and plan are required")
	}
	invoices.Create(billing.Invoice{AccountID: request.Account.ID, AmountCents: request.Amount, Paid: true})
	entitlements.Enable(request.Account.ID, request.Plan)
	return nil
}
