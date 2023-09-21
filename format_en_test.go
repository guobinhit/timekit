package timekit

import (
	"testing"
	"time"
)

func TestFormatEnOfMmDd(t *testing.T) {
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
			name: "FormatEnOfMmDd",
			args: args{d: d},
			want: "04/13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatEnOfMmDd(tt.args.d); got != tt.want {
				t.Errorf("FormatEnOfMmDd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatEnOfYyyyMm(t *testing.T) {
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
			name: "FormatEnOfYyyyMm",
			args: args{d: d},
			want: "2022/04",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatEnOfYyyyMm(tt.args.d); got != tt.want {
				t.Errorf("FormatEnOfYyyyMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatEnOfYyyyMmDd(t *testing.T) {
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
			name: "FormatEnOfYyyyMmDd",
			args: args{d: d},
			want: "2022/04/13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatEnOfYyyyMmDd(tt.args.d); got != tt.want {
				t.Errorf("FormatEnOfYyyyMmDd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatEnOfMmDdHhMm(t *testing.T) {
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
			name: "FormatEnOfMmDdHhMm",
			args: args{d: d},
			want: "04/13 10:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatEnOfMmDdHhMm(tt.args.d); got != tt.want {
				t.Errorf("FormatEnOfMmDdHhMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatEnOfMmDdHhMmSs(t *testing.T) {
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
			name: "FormatEnOfMmDdHhMmSs",
			args: args{d: d},
			want: "04/13 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatEnOfMmDdHhMmSs(tt.args.d); got != tt.want {
				t.Errorf("FormatEnOfMmDdHhMmSs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatEnOfYyyyMmDdHhMm(t *testing.T) {
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
			name: "FormatEnOfYyyyMmDdHhMm",
			args: args{d: d},
			want: "2022/04/13 10:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatEnOfYyyyMmDdHhMm(tt.args.d); got != tt.want {
				t.Errorf("FormatEnOfYyyyMmDdHhMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatEnOfYyyyMmDdHhMmSs(t *testing.T) {
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
			name: "FormatEnOfYyyyMmDdHhMmSs",
			args: args{d: d},
			want: "2022/04/13 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatEnOfYyyyMmDdHhMmSs(tt.args.d); got != tt.want {
				t.Errorf("FormatEnOfYyyyMmDdHhMmSs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatEnOfYyyyMmDdHhMmSsSss(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSsSss, "2022-04-13 10:20:59.999", time.Local)
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FormatEnOfYyyyMmDdHhMmSsSss",
			args: args{d: d},
			want: "2022/04/13 10:20:59.999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatEnOfYyyyMmDdHhMmSsSss(tt.args.d); got != tt.want {
				t.Errorf("FormatEnOfYyyyMmDdHhMmSsSss() = %v, want %v", got, tt.want)
			}
		})
	}
}
