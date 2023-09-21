package timekit

import (
	"testing"
	"time"
)

func TestFormatMmDd(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatMmDd",
			args: args{d: d},
			want: "04-13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatMmDd(tt.args.d); got != tt.want {
				t.Errorf("FormatMmDd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatYyyyMm(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatYyyyMm",
			args: args{d: d},
			want: "2022-04",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatYyyyMm(tt.args.d); got != tt.want {
				t.Errorf("FormatYyyyMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatYyyyMmDd(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatYyyyMmDd",
			args: args{d: d},
			want: "2022-04-13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatYyyyMmDd(tt.args.d); got != tt.want {
				t.Errorf("FormatYyyyMmDd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatMmDdHhMm(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatMmDdHhMm",
			args: args{d: d},
			want: "04-13 10:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatMmDdHhMm(tt.args.d); got != tt.want {
				t.Errorf("FormatMmDdHhMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatMmDdHhMmSs(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSs, "2022-04-13 10:00:00", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatMmDdHhMmSs",
			args: args{d: d},
			want: "04-13 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatMmDdHhMmSs(tt.args.d); got != tt.want {
				t.Errorf("FormatMmDdHhMmSs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatMmDdHhMmSsSss(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSsSss, "2022-04-13 10:00:00.999", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatMmDdHhMmSsSss",
			args: args{d: d},
			want: "04-13 10:00:00.999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatMmDdHhMmSsSss(tt.args.d); got != tt.want {
				t.Errorf("FormatMmDdHhMmSsSss() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatYyyyMmDdHhMm(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSsSss, "2022-04-13 10:00:00.999", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatYyyyMmDdHhMm",
			args: args{d: d},
			want: "2022-04-13 10:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatYyyyMmDdHhMm(tt.args.d); got != tt.want {
				t.Errorf("FormatYyyyMmDdHhMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatYyyyMmDdHhMmSs(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSsSss, "2022-04-13 10:00:00.999", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatYyyyMmDdHhMmSs",
			args: args{d: d},
			want: "2022-04-13 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatYyyyMmDdHhMmSs(tt.args.d); got != tt.want {
				t.Errorf("FormatYyyyMmDdHhMmSs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatYyyyMmDdHhMmSsSss(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSsSss, "2022-04-13 10:00:00.999", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatYyyyMmDdHhMmSsSss",
			args: args{d: d},
			want: "2022-04-13 10:00:00.999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatYyyyMmDdHhMmSsSss(tt.args.d); got != tt.want {
				t.Errorf("FormatYyyyMmDdHhMmSsSss() = %v, want %v", got, tt.want)
			}
		})
	}
}
