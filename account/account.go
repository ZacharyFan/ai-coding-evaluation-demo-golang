package account

import (
	"fmt"
	"strings"
)

type Account struct {
	ID               string
	Email            string
	CreditLimitCents int
	BalanceCents     int
	Hold             bool
	SpendCents       int
	ReturnCents      int
}

func AvailableCredit(account Account) int {
	available := account.CreditLimitCents - account.BalanceCents
	if available < 0 {
		return 0
	}
	return available
}

func CanUseCredit(account Account, amountCents int) bool {
	if amountCents <= 0 {
		return false
	}
	return AvailableCredit(account) >= amountCents
}

func MaskEmail(email string) (string, error) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", fmt.Errorf("invalid email")
	}
	local := parts[0]
	if len(local) == 1 {
		return "*@" + parts[1], nil
	}
	return local[:1] + "***@" + parts[1], nil
}
