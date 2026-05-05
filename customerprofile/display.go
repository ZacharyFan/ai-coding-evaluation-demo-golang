package customerprofile

import "strings"

func DisplayName(profile Profile) string {
	name := strings.TrimSpace(strings.TrimSpace(profile.FirstName) + " " + strings.TrimSpace(profile.LastName))
	if name != "" {
		return name
	}
	if strings.TrimSpace(profile.Email) != "" {
		return strings.TrimSpace(profile.Email)
	}
	return "anonymous"
}
