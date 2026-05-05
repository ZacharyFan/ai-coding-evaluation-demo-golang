package pricing

import "fmt"

func CalculateTax(subtotalCents int, basisPoints int) (int, error) {
	if subtotalCents < 0 {
		return 0, fmt.Errorf("subtotal cents must be non-negative")
	}
	if basisPoints < 0 {
		return 0, fmt.Errorf("basis points must be non-negative")
	}
	return subtotalCents * basisPoints / 10000, nil
}

func NetAfterReturns(grossCents int, returnCents int) (int, error) {
	if grossCents < 0 {
		return 0, fmt.Errorf("gross cents must be non-negative")
	}
	if returnCents < 0 {
		return 0, fmt.Errorf("return cents must be non-negative")
	}
	if returnCents > grossCents {
		return 0, fmt.Errorf("return cents cannot exceed gross cents")
	}
	return grossCents - returnCents, nil
}

func InvoiceTotal(subtotalCents int, taxBasisPoints int, creditCents int) (int, error) {
	if subtotalCents < 0 {
		return 0, fmt.Errorf("subtotal cents must be non-negative")
	}
	if creditCents < 0 {
		return 0, fmt.Errorf("credit cents must be non-negative")
	}
	if creditCents > subtotalCents {
		return 0, fmt.Errorf("credit cents cannot exceed subtotal cents")
	}
	taxCents, err := CalculateTax(subtotalCents, taxBasisPoints)
	if err != nil {
		return 0, err
	}
	return subtotalCents + taxCents - creditCents, nil
}
