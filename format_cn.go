package timekit

import (
	"time"
)

// FormatCnOfMmDd returns '01月02日' format date string.
func FormatCnOfMmDd(d time.Time) string {
	return d.Format(CnOfMmDd)
}

// FormatCnOfYyyyMm returns '2006年01月' format date string.
func FormatCnOfYyyyMm(d time.Time) string {
	return d.Format(CnOfYyyyMm)
}

// FormatCnOfYyyyMmDd returns '2006年01月02日' format date string.
func FormatCnOfYyyyMmDd(d time.Time) string {
	return d.Format(CnOfYyyyMmDd)
}

// FormatCnOfMmDdHhMm returns '01月02日 15:04' format date string.
func FormatCnOfMmDdHhMm(d time.Time) string {
	return d.Format(CnOfMmDdHhMm)
}

// FormatCnOfMmDdHhMmSs returns '01月02日 15:04:05' format date string.
func FormatCnOfMmDdHhMmSs(d time.Time) string {
	return d.Format(CnOfMmDdHhMmSs)
}

// FormatCnOfYyyyMmDdHhMm returns '2006年01月02日 15:04' format date string.
func FormatCnOfYyyyMmDdHhMm(d time.Time) string {
	return d.Format(CnOfYyyyMmDdHhMm)
}

// FormatCnOfYyyyMmDdHhMmSs returns '2006年01月02日 15:04:05' format date string.
func FormatCnOfYyyyMmDdHhMmSs(d time.Time) string {
	return d.Format(CnOfYyyyMmDdHhMmSs)
}

// FormatCnOfYyyyMmDdHhMmSsSss returns '2006年01月02日 15:04:05.000' format date string.
func FormatCnOfYyyyMmDdHhMmSsSss(d time.Time) string {
	return d.Format(CnOfYyyyMmDdHhMmSsSss)
}
