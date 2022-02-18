package excel_serial_date

import (
	"time"
)

// Time converts excel serial date to time.Time
//
// Serial date references:
//
// 1) 1 : serial date of Jan 01, 1900 (excel)
//
// 2) 25569 : serial date of Jan 01, 1970 (go time.Time)
func Time(serial int64) time.Time {
	val := time.Unix((serial-25569)*86400, 0)
	if serial < 25569 {
		return val.AddDate(0, 0, 1)
	}
	return val
}
