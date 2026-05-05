package customerprofile

import (
	"fmt"
	"strings"
)

type Profile struct {
	ID             string
	Email          string
	FirstName      string
	LastName       string
	MarketingOptIn bool
	Deleted        bool
}

func NormalizeEmail(email string) (string, error) {
	normalized := strings.ToLower(strings.TrimSpace(email))
	parts := strings.Split(normalized, "@")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", fmt.Errorf("invalid email")
	}
	return normalized, nil
}

func Validate(profile Profile) error {
	if strings.TrimSpace(profile.ID) == "" {
		return fmt.Errorf("id is required")
	}
	_, err := NormalizeEmail(profile.Email)
	return err
}

func CanSendMarketing(profile Profile) bool {
	return profile.MarketingOptIn && !profile.Deleted
}
