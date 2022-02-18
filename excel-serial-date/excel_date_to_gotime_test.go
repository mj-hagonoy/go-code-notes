package excel_serial_date

import (
	"testing"
)

var testDates = map[int64]string{
	1:     "01/01/1900",
	25569: "01/01/1970",
	32874: "01/01/1990",
	36526: "01/01/2000",
	43831: "01/01/2020",
	44197: "01/01/2021",
	44562: "01/01/2022",
}

func TestSerialToTime(t *testing.T) {
	for serial, expected := range testDates {
		actual := (Time(serial)).Format("01/02/2006")
		if actual != expected {
			t.Errorf("expected %s got %s", expected, actual)
		}
	}
}
