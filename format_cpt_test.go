package timekit

import (
	"testing"
	"time"
)

func TestFormatCptOfYyyyMm(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSs, "2022-04-30 10:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatCptOfYyyyMm",
			args: args{d: d},
			want: "202204",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCptOfYyyyMm(tt.args.d); got != tt.want {
				t.Errorf("FormatCptOfYyyyMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatCptOfYyyyMmDd(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSs, "2022-04-01 10:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatCptOfYyyyMmDd",
			args: args{d: d},
			want: "20220401",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCptOfYyyyMmDd(tt.args.d); got != tt.want {
				t.Errorf("FormatCptOfYyyyMmDd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatCptOfYyyyMmDdHh(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSs, "2022-04-13 23:59:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatCptOfYyyyMmDdHh",
			args: args{d: d},
			want: "2022041323",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCptOfYyyyMmDdHh(tt.args.d); got != tt.want {
				t.Errorf("FormatCptOfYyyyMmDdHh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatCptOfYyyyMmDdHhMm(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSsSss, "2022-04-13 10:59:00.999", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatCptOfYyyyMmDdHhMm",
			args: args{d: d},
			want: "202204131059",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCptOfYyyyMmDdHhMm(tt.args.d); got != tt.want {
				t.Errorf("FormatCptOfYyyyMmDdHhMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatCptOfYyyyMmDdHhMmSs(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSsSss, "2022-04-13 13:20:01.123", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatCptOfYyyyMmDdHhMmSs",
			args: args{d: d},
			want: "20220413132001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCptOfYyyyMmDdHhMmSs(tt.args.d); got != tt.want {
				t.Errorf("FormatCptOfYyyyMmDdHhMmSs() = %v, want %v", got, tt.want)
			}
		})
	}
}
