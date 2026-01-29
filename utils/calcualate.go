//- utils/calculate.go

package utils

import "github.com/shopspring/decimal"

func CalculateTaxAmount(
	subtotal decimal.Decimal,
	discountAmount decimal.Decimal,
) decimal.Decimal {
	taxRatePercent := 10
	taxableAmount := subtotal.Sub(discountAmount)
	taxAmount := taxableAmount.Mul(decimal.NewFromFloat(float64(taxRatePercent) / 100))

	return taxAmount
}

func CalculateGrandTotal(
	subtotal decimal.Decimal,
	discountAmount decimal.Decimal,
	taxAmount decimal.Decimal,
) decimal.Decimal {
	return subtotal.Sub(discountAmount).Add(taxAmount)
}

func CalculateItemAmount(
	qty decimal.Decimal,
	price decimal.Decimal,
) decimal.Decimal {
	return qty.Mul(price)
}
