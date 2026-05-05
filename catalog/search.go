package catalog

import "sort"

func FindAvailableByTag(products []Product, tag string) []Product {
	result := []Product{}
	for _, product := range products {
		if !AvailableForSale(product) {
			continue
		}
		for _, current := range product.Tags {
			if current == tag {
				result = append(result, product)
				break
			}
		}
	}
	sort.Slice(result, func(i, j int) bool { return NormalizeSKU(result[i].SKU) < NormalizeSKU(result[j].SKU) })
	return result
}
