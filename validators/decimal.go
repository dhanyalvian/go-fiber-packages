//- validators/decimal.go

package validators

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/shopspring/decimal"
)

func DecimalGreaterThan(min decimal.Decimal, msg string) validation.Rule {
	return validation.By(func(v interface{}) error {
		val, ok := v.(decimal.Decimal)
		if !ok {
			return errors.New("invalid decimal value")
		}

		if val.LessThanOrEqual(min) {
			if msg != "" {
				return errors.New(msg)
			}
			return errors.New("value must be greater than " + min.String())
		}

		return nil
	})
}

func DecimalMin(min decimal.Decimal, msg string) validation.Rule {
	return validation.By(func(v interface{}) error {
		val, ok := v.(decimal.Decimal)
		if !ok {
			return nil // optional
		}

		if val.LessThan(min) {
			if msg != "" {
				return errors.New(msg)
			}
			return errors.New("value must be >= " + min.String())
		}

		return nil
	})
}
