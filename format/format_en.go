package format

import (
	"github.com/guobinhit/auxo"
	"time"
)

// GetEnOfMmDd returns '01/02' format date string.
func GetEnOfMmDd(d time.Time) string {
	return d.Format(auxo.EnOfMmDd)
}

// GetEnOfYyyyMm returns '2006/01' format date string.
func GetEnOfYyyyMm(d time.Time) string {
	return d.Format(auxo.EnOfYyyyMm)
}

// GetEnOfYyyyMmDd returns '2006/01/02' format date string.
func GetEnOfYyyyMmDd(d time.Time) string {
	return d.Format(auxo.EnOfYyyyMmDd)
}

// GetEnOfMmDdHhMm returns '01/02 15:04' format date string.
func GetEnOfMmDdHhMm(d time.Time) string {
	return d.Format(auxo.EnOfMmDdHhMm)
}

// GetEnOfMmDdHhMmSs returns '01/02 15:04:05' format date string.
func GetEnOfMmDdHhMmSs(d time.Time) string {
	return d.Format(auxo.EnOfMmDdHhMmSs)
}

// GetEnOfYyyyMmDdHhMm returns '2006/01/02 15:04' format date string.
func GetEnOfYyyyMmDdHhMm(d time.Time) string {
	return d.Format(auxo.EnOfYyyyMmDdHhMm)
}

// GetEnOfYyyyMmDdHhMmSs returns '2006/01/02 15:04:05' format date string.
func GetEnOfYyyyMmDdHhMmSs(d time.Time) string {
	return d.Format(auxo.EnOfYyyyMmDdHhMmSs)
}

// GetEnOfYyyyMmDdHhMmSsSss returns '2006/01/02 15:04:05.000' format date string.
func GetEnOfYyyyMmDdHhMmSsSss(d time.Time) string {
	return d.Format(auxo.EnOfYyyyMmDdHhMmSsSss)
}
