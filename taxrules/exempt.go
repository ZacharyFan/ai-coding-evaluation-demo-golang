package taxrules

import "strings"

func IsExemptCategory(category string) bool {
	normalized := strings.ToLower(strings.TrimSpace(category))
	return normalized == "food" || normalized == "medicine"
}
