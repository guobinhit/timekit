package format

import (
	"time"

	"github.com/guobinhit/auxo"
)

// GetMmDd returns '01-02' format date string.
func GetMmDd(d time.Time) string {
	return d.Format(auxo.MmDd)
}

// GetYyyyMm returns '2006-01' format date string.
func GetYyyyMm(d time.Time) string {
	return d.Format(auxo.YyyyMm)
}

// GetYyyyMmDd returns '2006-01-02' format date string.
func GetYyyyMmDd(d time.Time) string {
	return d.Format(auxo.YyyyMmDd)
}

// GetMmDdHhMm returns '01-02 15:04' format date string.
func GetMmDdHhMm(d time.Time) string {
	return d.Format(auxo.MmDdHhMm)
}

// GetMmDdHhMmSs returns '01-02 15:04:05' format date string.
func GetMmDdHhMmSs(d time.Time) string {
	return d.Format(auxo.MmDdHhMmSs)
}

// GetMmDdHhMmSsSss returns '01-02 15:04:05.000' format date string.
func GetMmDdHhMmSsSss(d time.Time) string {
	return d.Format(auxo.MmDdHhMmSsSss)
}

// GetYyyyMmDdHhMm returns '2006-01-02 15:04' format date string.
func GetYyyyMmDdHhMm(d time.Time) string {
	return d.Format(auxo.YyyyMmDdHhMm)
}

// GetYyyyMmDdHhMmSs returns '2006-01-02 15:04:05' format date string.
func GetYyyyMmDdHhMmSs(d time.Time) string {
	return d.Format(auxo.YyyyMmDdHhMmSs)
}

// GetYyyyMmDdHhMmSsSss returns '2006-01-02 15:04:05.000' format date string.
func GetYyyyMmDdHhMmSsSss(d time.Time) string {
	return d.Format(auxo.YyyyMmDdHhMmSsSss)
}
