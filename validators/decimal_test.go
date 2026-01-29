//- validators/decimal_test.go

package validators

import (
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestDecimalGreaterThan(t *testing.T) {
	minVal := decimal.NewFromInt(100)
	customMsg := "Nominal harus lebih dari 100"

	tests := []struct {
		name    string
		rule    validation.Rule
		input   interface{}
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Valid: Nilai lebih besar",
			rule:    DecimalGreaterThan(minVal, ""),
			input:   decimal.NewFromInt(150),
			wantErr: false,
		},
		{
			name:    "Invalid: Nilai sama dengan minimum (default msg)",
			rule:    DecimalGreaterThan(minVal, ""),
			input:   decimal.NewFromInt(100),
			wantErr: true,
			errMsg:  "value must be greater than 100",
		},
		{
			name:    "Invalid: Nilai lebih kecil (custom msg)",
			rule:    DecimalGreaterThan(minVal, customMsg),
			input:   decimal.NewFromInt(50),
			wantErr: true,
			errMsg:  customMsg,
		},
		{
			name:    "Invalid: Tipe data bukan decimal",
			rule:    DecimalGreaterThan(minVal, ""),
			input:   "150", // string, bukan decimal.Decimal
			wantErr: true,
			errMsg:  "invalid decimal value",
		},
		{
			name:    "Invalid: Nilai nol",
			rule:    DecimalGreaterThan(decimal.Zero, ""),
			input:   decimal.NewFromInt(-10),
			wantErr: true,
			errMsg:  "value must be greater than 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Memanggil internal function dari ozzo-validation
			err := tt.rule.Validate(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Equal(t, tt.errMsg, err.Error())
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDecimalMin(t *testing.T) {
	minimum := decimal.NewFromInt(1000)
	customMsg := "Saldo minimal adalah 1000"

	tests := []struct {
		name    string
		rule    validation.Rule
		input   interface{}
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Valid: Nilai pas dengan minimum (1000)",
			rule:    DecimalMin(minimum, ""),
			input:   decimal.NewFromInt(1000),
			wantErr: false,
		},
		{
			name:    "Valid: Nilai lebih besar dari minimum (1500)",
			rule:    DecimalMin(minimum, ""),
			input:   decimal.NewFromInt(1500),
			wantErr: false,
		},
		{
			name:    "Invalid: Nilai lebih kecil (default message)",
			rule:    DecimalMin(minimum, ""),
			input:   decimal.NewFromInt(999),
			wantErr: true,
			errMsg:  "value must be >= 1000",
		},
		{
			name:    "Invalid: Nilai lebih kecil (custom message)",
			rule:    DecimalMin(minimum, customMsg),
			input:   decimal.NewFromInt(500),
			wantErr: true,
			errMsg:  customMsg,
		},
		{
			name:    "Ignore: Tipe data bukan decimal (return nil sesuai kode)",
			rule:    DecimalMin(minimum, ""),
			input:   "bukan decimal",
			wantErr: false,
		},
		{
			name:    "Ignore: Input nil",
			rule:    DecimalMin(minimum, ""),
			input:   nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
