package auxo

import (
	"time"
)

// FormatEnOfMmDd returns '01/02' format date string.
func FormatEnOfMmDd(d time.Time) string {
	return d.Format(EnOfMmDd)
}

// FormatEnOfYyyyMm returns '2006/01' format date string.
func FormatEnOfYyyyMm(d time.Time) string {
	return d.Format(EnOfYyyyMm)
}

// FormatEnOfYyyyMmDd returns '2006/01/02' format date string.
func FormatEnOfYyyyMmDd(d time.Time) string {
	return d.Format(EnOfYyyyMmDd)
}

// FormatEnOfMmDdHhMm returns '01/02 15:04' format date string.
func FormatEnOfMmDdHhMm(d time.Time) string {
	return d.Format(EnOfMmDdHhMm)
}

// FormatEnOfMmDdHhMmSs returns '01/02 15:04:05' format date string.
func FormatEnOfMmDdHhMmSs(d time.Time) string {
	return d.Format(EnOfMmDdHhMmSs)
}

// FormatEnOfYyyyMmDdHhMm returns '2006/01/02 15:04' format date string.
func FormatEnOfYyyyMmDdHhMm(d time.Time) string {
	return d.Format(EnOfYyyyMmDdHhMm)
}

// FormatEnOfYyyyMmDdHhMmSs returns '2006/01/02 15:04:05' format date string.
func FormatEnOfYyyyMmDdHhMmSs(d time.Time) string {
	return d.Format(EnOfYyyyMmDdHhMmSs)
}

// FormatEnOfYyyyMmDdHhMmSsSss returns '2006/01/02 15:04:05.000' format date string.
func FormatEnOfYyyyMmDdHhMmSsSss(d time.Time) string {
	return d.Format(EnOfYyyyMmDdHhMmSsSss)
}
