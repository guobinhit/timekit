package timekit

import (
	"time"
)

// FormatMmDd returns '01-02' format date string.
func FormatMmDd(d time.Time) string {
	return d.Format(MmDd)
}

// FormatYyyyMm returns '2006-01' format date string.
func FormatYyyyMm(d time.Time) string {
	return d.Format(YyyyMm)
}

// FormatYyyyMmDd returns '2006-01-02' format date string.
func FormatYyyyMmDd(d time.Time) string {
	return d.Format(YyyyMmDd)
}

// FormatMmDdHhMm returns '01-02 15:04' format date string.
func FormatMmDdHhMm(d time.Time) string {
	return d.Format(MmDdHhMm)
}

// FormatMmDdHhMmSs returns '01-02 15:04:05' format date string.
func FormatMmDdHhMmSs(d time.Time) string {
	return d.Format(MmDdHhMmSs)
}

// FormatMmDdHhMmSsSss returns '01-02 15:04:05.000' format date string.
func FormatMmDdHhMmSsSss(d time.Time) string {
	return d.Format(MmDdHhMmSsSss)
}

// FormatYyyyMmDdHhMm returns '2006-01-02 15:04' format date string.
func FormatYyyyMmDdHhMm(d time.Time) string {
	return d.Format(YyyyMmDdHhMm)
}

// FormatYyyyMmDdHhMmSs returns '2006-01-02 15:04:05' format date string.
func FormatYyyyMmDdHhMmSs(d time.Time) string {
	return d.Format(YyyyMmDdHhMmSs)
}

// FormatYyyyMmDdHhMmSsSss returns '2006-01-02 15:04:05.000' format date string.
func FormatYyyyMmDdHhMmSsSss(d time.Time) string {
	return d.Format(YyyyMmDdHhMmSsSss)
}
