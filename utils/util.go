//- utils/util.go

package utils

import (
	"database/sql/driver"
	"time"
)

type DateOnly string

// MarshalJSON: Mengubah tipe data menjadi "YYYY-MM-DD" saat dikirim ke Frontend (JSON)
func (d DateOnly) MarshalJSON() ([]byte, error) {
	str := string(d)
	if len(str) >= 10 {
		return []byte(`"` + str[:10] + `"`), nil
	}
	return []byte(`"` + str + `"`), nil
}

// Scan: Mengambil data dari Database dan mengubahnya menjadi DateOnly
func (d *DateOnly) Scan(value interface{}) error {
	if value == nil {
		*d = ""
		return nil
	}

	var t time.Time
	switch v := value.(type) {
	case time.Time:
		t = v
	case []byte:
		// Coba parsing jika database mengirim string/bytes
		t, _ = time.Parse("2006-01-02", string(v[:10]))
	case string:
		t, _ = time.Parse("2006-01-02", v[:10])
	}

	*d = DateOnly(t.Format("2006-01-02"))
	return nil
}

// Value: Memastikan GORM bisa menyimpan kembali tipe ini ke database jika diperlukan
func (d DateOnly) Value() (driver.Value, error) {
	if d == "" {
		return nil, nil
	}
	return string(d), nil
}
