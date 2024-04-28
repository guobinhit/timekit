package timekit

import (
	"time"
)

// ParseMmDd return time.Time with '01-02' format value，others location is default value.
func ParseMmDd(dStr string) (time.Time, error) {
	return time.ParseInLocation(MmDd, dStr, time.Local)
}

// ParseYyyyMm return time.Time with '2006-01' format value，others location is default value.
func ParseYyyyMm(dStr string) (time.Time, error) {
	return time.ParseInLocation(YyyyMm, dStr, time.Local)
}

// ParseYyyyMmDd return time.Time with '2006-01-02' format value，others location is default value.
func ParseYyyyMmDd(dStr string) (time.Time, error) {
	return time.ParseInLocation(YyyyMmDd, dStr, time.Local)
}

// ParseMmDdHhMm return time.Time with '01-02 15:04' format value，others location is default value.
func ParseMmDdHhMm(dStr string) (time.Time, error) {
	return time.ParseInLocation(MmDdHhMm, dStr, time.Local)
}

// ParseMmDdHhMmSs return time.Time with '01-02 15:04:05' format value，others location is default value.
func ParseMmDdHhMmSs(dStr string) (time.Time, error) {
	return time.ParseInLocation(MmDdHhMmSs, dStr, time.Local)
}

// ParseMmDdHhMmSsSss return time.Time with '01-02 15:04:05.000' format value，others location is default value.
func ParseMmDdHhMmSsSss(dStr string) (time.Time, error) {
	return time.ParseInLocation(MmDdHhMmSsSss, dStr, time.Local)
}

// ParseYyyyMmDdHhMm return time.Time with '2006-01-02 15:04' format value，others location is default value.
func ParseYyyyMmDdHhMm(dStr string) (time.Time, error) {
	return time.ParseInLocation(YyyyMmDdHhMm, dStr, time.Local)
}

// ParseYyyyMmDdHhMmSs return time.Time with '2006-01-02 15:04:05' format value，others location is default value.
func ParseYyyyMmDdHhMmSs(dStr string) (time.Time, error) {
	return time.ParseInLocation(YyyyMmDdHhMmSs, dStr, time.Local)
}

// ParseYyyyMmDdHhMmSsSss return time.Time with '2006-01-02 15:04:05.000' format value，others location is default value.
func ParseYyyyMmDdHhMmSsSss(dStr string) (time.Time, error) {
	return time.ParseInLocation(YyyyMmDdHhMmSsSss, dStr, time.Local)
}

// ParseTimestampUnix return time.Time by parse timeUnix seconds.
func ParseTimestampUnix(timeUnix int64) time.Time {
	return time.Unix(timeUnix, 0)
}

// ParseTimestampUnixMilli return time.Time by parse timeUnixMilli milliseconds.
func ParseTimestampUnixMilli(timeUnixMilli int64) time.Time {
	return time.UnixMilli(timeUnixMilli)
}
