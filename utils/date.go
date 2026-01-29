//- utils/date.go

package utils

import (
	"time"

	"gorm.io/datatypes"
)

func StringToDate(dateStr string) (datatypes.Date, error) {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return datatypes.Date{}, err
	}

	return datatypes.Date(t), nil
}
