package timekit

import (
	"time"
)

// FormatCptOfYyyyMm returns '200601' format date string.
func FormatCptOfYyyyMm(d time.Time) string {
	return d.Format(CptOfYyyyMm)
}

// FormatCptOfYyyyMmDd returns '20060102' format date string.
func FormatCptOfYyyyMmDd(d time.Time) string {
	return d.Format(CptOfYyyyMmDd)
}

// FormatCptOfYyyyMmDdHh returns '2006010215' format date string.
func FormatCptOfYyyyMmDdHh(d time.Time) string {
	return d.Format(CptOfYyyyMmDdHh)
}

// FormatCptOfYyyyMmDdHhMm returns '200601021504' format date string.
func FormatCptOfYyyyMmDdHhMm(d time.Time) string {
	return d.Format(CptOfYyyyMmDdHhMm)
}

// FormatCptOfYyyyMmDdHhMmSs returns '20060102150405' format date string.
func FormatCptOfYyyyMmDdHhMmSs(d time.Time) string {
	return d.Format(CptOfYyyyMmDdHhMmSs)
}
