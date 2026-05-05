package orders

import (
	"fmt"
	"strings"
)

type Line struct {
	SKU             string
	UnitCents       int
	Quantity        int
	DiscountPercent int
}

type Coupon struct {
	Code                 string
	PercentOff           int
	MinimumSubtotalCents int
}

type Order struct {
	ID         string
	Status     string
	TotalCents int
}

type Promotion struct {
	SKU        string
	PercentOff int
}

func Subtotal(lines []Line) (int, error) {
	total := 0
	for _, line := range lines {
		if line.UnitCents < 0 {
			return 0, fmt.Errorf("unit cents must be non-negative")
		}
		if line.Quantity <= 0 {
			return 0, fmt.Errorf("quantity must be positive")
		}
		if line.DiscountPercent < 0 || line.DiscountPercent > 100 {
			return 0, fmt.Errorf("discount percent must be between 0 and 100")
		}
		lineTotal := line.UnitCents * line.Quantity
		total += lineTotal * (100 - line.DiscountPercent) / 100
	}
	return total, nil
}

func ApplyCoupon(subtotalCents int, coupon Coupon) (int, error) {
	if subtotalCents < 0 {
		return 0, fmt.Errorf("subtotal cents must be non-negative")
	}
	if coupon.PercentOff < 0 || coupon.PercentOff > 100 {
		return 0, fmt.Errorf("coupon percent must be between 0 and 100")
	}
	return subtotalCents * (100 - coupon.PercentOff) / 100, nil
}

func CanTransition(from string, to string) bool {
	if from == "pending" && (to == "paid" || to == "canceled") {
		return true
	}
	if from == "paid" && (to == "shipped" || to == "refunded") {
		return true
	}
	if from == "shipped" && to == "delivered" {
		return true
	}
	return false
}

func PromotionDiscountCents(line Line, promotions []Promotion) (int, error) {
	if line.UnitCents < 0 || line.Quantity <= 0 {
		return 0, fmt.Errorf("line must have non-negative unit cents and positive quantity")
	}
	lineTotal := line.UnitCents * line.Quantity
	discount := 0
	for _, promotion := range promotions {
		if promotion.SKU == line.SKU {
			if promotion.PercentOff < 0 || promotion.PercentOff > 100 {
				return 0, fmt.Errorf("promotion percent must be between 0 and 100")
			}
			discount += lineTotal * promotion.PercentOff / 100
		}
	}
	return discount, nil
}

func ReceiptText(order Order, lines []Line) string {
	var builder strings.Builder
	builder.WriteString("Order ")
	builder.WriteString(order.ID)
	builder.WriteString("\n")
	builder.WriteString("Status: ")
	builder.WriteString(order.Status)
	builder.WriteString("\n")
	for _, line := range lines {
		builder.WriteString(line.SKU)
		builder.WriteString("\n")
	}
	return builder.String()
}
