package timekit

import (
	"time"
)

// ParseMmDd returns time.Time with '01-02' format value，others location is default value.
func ParseMmDd(dStr string) (time.Time, error) {
	return time.ParseInLocation(MmDd, dStr, time.Local)
}

// ParseYyyyMm returns time.Time with '2006-01' format value，others location is default value.
func ParseYyyyMm(dStr string) (time.Time, error) {
	return time.ParseInLocation(YyyyMm, dStr, time.Local)
}

// ParseYyyyMmDd returns time.Time with '2006-01-02' format value，others location is default value.
func ParseYyyyMmDd(dStr string) (time.Time, error) {
	return time.ParseInLocation(YyyyMmDd, dStr, time.Local)
}

// ParseMmDdHhMm returns time.Time with '01-02 15:04' format value，others location is default value.
func ParseMmDdHhMm(dStr string) (time.Time, error) {
	return time.ParseInLocation(MmDdHhMm, dStr, time.Local)
}

// ParseMmDdHhMmSs returns time.Time with '01-02 15:04:05' format value，others location is default value.
func ParseMmDdHhMmSs(dStr string) (time.Time, error) {
	return time.ParseInLocation(MmDdHhMmSs, dStr, time.Local)
}

// ParseMmDdHhMmSsSss returns time.Time with '01-02 15:04:05.000' format value，others location is default value.
func ParseMmDdHhMmSsSss(dStr string) (time.Time, error) {
	return time.ParseInLocation(MmDdHhMmSsSss, dStr, time.Local)
}

// ParseYyyyMmDdHhMm returns time.Time with '2006-01-02 15:04' format value，others location is default value.
func ParseYyyyMmDdHhMm(dStr string) (time.Time, error) {
	return time.ParseInLocation(YyyyMmDdHhMm, dStr, time.Local)
}

// ParseYyyyMmDdHhMmSs returns time.Time with '2006-01-02 15:04:05' format value，others location is default value.
func ParseYyyyMmDdHhMmSs(dStr string) (time.Time, error) {
	return time.ParseInLocation(YyyyMmDdHhMmSs, dStr, time.Local)
}

// ParseYyyyMmDdHhMmSsSss returns time.Time with '2006-01-02 15:04:05.000' format value，others location is default value.
func ParseYyyyMmDdHhMmSsSss(dStr string) (time.Time, error) {
	return time.ParseInLocation(YyyyMmDdHhMmSsSss, dStr, time.Local)
}
