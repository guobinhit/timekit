package auxo

import (
	"testing"
	"time"
)

func TestFormatCnOfMmDd(t *testing.T) {
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
			name: "FormatCnOfMmDd",
			args: args{d: d},
			want: "04月13日",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCnOfMmDd(tt.args.d); got != tt.want {
				t.Errorf("FormatCnOfMmDd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatCnOfYyyyMm(t *testing.T) {
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
			name: "FormatCnOfYyyyMm",
			args: args{d: d},
			want: "2022年04月",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCnOfYyyyMm(tt.args.d); got != tt.want {
				t.Errorf("FormatCnOfYyyyMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatCnOfYyyyMmDd(t *testing.T) {
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
			name: "FormatCnOfYyyyMmDd",
			args: args{d: d},
			want: "2022年04月13日",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCnOfYyyyMmDd(tt.args.d); got != tt.want {
				t.Errorf("FormatCnOfYyyyMmDd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatCnOfMmDdHhMm(t *testing.T) {
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
			name: "FormatCnOfMmDdHhMm",
			args: args{d: d},
			want: "04月13日 10:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCnOfMmDdHhMm(tt.args.d); got != tt.want {
				t.Errorf("FormatCnOfMmDdHhMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatCnOfMmDdHhMmSs(t *testing.T) {
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
			name: "FormatCnOfMmDdHhMmSs",
			args: args{d: d},
			want: "04月13日 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCnOfMmDdHhMmSs(tt.args.d); got != tt.want {
				t.Errorf("FormatCnOfMmDdHhMmSs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatCnOfYyyyMmDdHhMm(t *testing.T) {
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
			name: "FormatCnOfYyyyMmDdHhMm",
			args: args{d: d},
			want: "2022年04月13日 10:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCnOfYyyyMmDdHhMm(tt.args.d); got != tt.want {
				t.Errorf("FormatCnOfYyyyMmDdHhMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatCnOfYyyyMmDdHhMmSs(t *testing.T) {
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
			name: "FormatCnOfYyyyMmDdHhMmSs",
			args: args{d: d},
			want: "2022年04月13日 10:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCnOfYyyyMmDdHhMmSs(tt.args.d); got != tt.want {
				t.Errorf("FormatCnOfYyyyMmDdHhMmSs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatCnOfYyyyMmDdHhMmSsSss(t *testing.T) {
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
			name: "FormatCnOfYyyyMmDdHhMmSsSss",
			args: args{d: d},
			want: "2022年04月13日 10:00:00.999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatCnOfYyyyMmDdHhMmSsSss(tt.args.d); got != tt.want {
				t.Errorf("FormatCnOfYyyyMmDdHhMmSsSss() = %v, want %v", got, tt.want)
			}
		})
	}
}
