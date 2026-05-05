package subscriptions

import (
	"fmt"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/account"
	"github.com/ZacharyFan/ai-coding-evaluation-demo-golang/entitlement"
)

type PlanChangePreview struct {
	ProratedCreditCents int
	NextChargeCents     int
}

func PreviewPlanChange(account account.Account, entitlements *entitlement.Store, nextPlan string, nextAmountCents int, currentCreditCents int) (PlanChangePreview, error) {
	if account.Hold {
		return PlanChangePreview{}, fmt.Errorf("account is on hold")
	}
	if !entitlements.IsActive(account.ID) {
		return PlanChangePreview{}, fmt.Errorf("active entitlement is required")
	}
	next := nextAmountCents - currentCreditCents
	if next < 0 {
		next = 0
	}
	return PlanChangePreview{ProratedCreditCents: currentCreditCents, NextChargeCents: next}, nil
}
