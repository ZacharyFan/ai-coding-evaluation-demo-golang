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
	if err := activationDecision(request.Account, request.Plan); err != nil {
		return err
	}
	invoices.Create(billing.Invoice{AccountID: request.Account.ID, AmountCents: request.Amount, Paid: true})
	entitlements.Enable(request.Account.ID, request.Plan)
	return nil
}

func activationDecision(account account.Account, plan string) error {
	if account.ID == "" || plan == "" {
		return fmt.Errorf("account and plan are required")
	}
	return nil
}
