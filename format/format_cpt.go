package format

import (
	"github.com/guobinhit/auxo"
	"time"
)

// GetCptOfYyyyMm returns '200601' format date string.
func GetCptOfYyyyMm(d time.Time) string {
	return d.Format(auxo.CptOfYyyyMm)
}

// GetCptOfYyyyMmDd returns '20060102' format date string.
func GetCptOfYyyyMmDd(d time.Time) string {
	return d.Format(auxo.CptOfYyyyMmDd)
}

// GetCptOfYyyyMmDdHh returns '2006010215' format date string.
func GetCptOfYyyyMmDdHh(d time.Time) string {
	return d.Format(auxo.CptOfYyyyMmDdHh)
}

// GetCptOfYyyyMmDdHhMm returns '200601021504' format date string.
func GetCptOfYyyyMmDdHhMm(d time.Time) string {
	return d.Format(auxo.CptOfYyyyMmDdHhMm)
}

// GetCptOfYyyyMmDdHhMmSs returns '20060102150405' format date string.
func GetCptOfYyyyMmDdHhMmSs(d time.Time) string {
	return d.Format(auxo.CptOfYyyyMmDdHhMmSs)
}
