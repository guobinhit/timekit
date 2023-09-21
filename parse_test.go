package timekit

import (
	"reflect"
	"testing"
	"time"
)

func TestParseMmDd(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSs, "0000-04-13 00:00:00", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "ParseMmDd error",
			args:    args{dStr: "2022-04-13"},
			wantErr: true,
		},
		{
			name: "ParseMmDd success",
			args: args{dStr: "04-13"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMmDd(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMmDd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMmDd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseYyyyMm(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSs, "2022-04-01 00:00:00", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "ParseYyyyMm error",
			args:    args{dStr: "2022-04-13"},
			wantErr: true,
		},
		{
			name: "ParseYyyyMm success",
			args: args{dStr: "2022-04"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseYyyyMm(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseYyyyMm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseYyyyMm() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseYyyyMmDd(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSs, "2022-04-13 00:00:00", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "ParseYyyyMmDd error",
			args:    args{dStr: "2022-04-13 00:00:00"},
			wantErr: true,
		},
		{
			name: "ParseYyyyMmDd success",
			args: args{dStr: "2022-04-13"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseYyyyMmDd(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseYyyyMmDd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseYyyyMmDd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseMmDdHhMm(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSs, "0000-04-13 10:20:00", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "ParseMmDdHhMm error",
			args:    args{dStr: "2022-04-13 00:00:00"},
			wantErr: true,
		},
		{
			name: "ParseMmDdHhMm success",
			args: args{dStr: "04-13 10:20"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMmDdHhMm(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMmDdHhMm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMmDdHhMm() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseMmDdHhMmSs(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSs, "0000-04-13 10:20:30", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "ParseMmDdHhMmSs error",
			args:    args{dStr: "2022-04-13 00:00:00.123"},
			wantErr: true,
		},
		{
			name: "ParseMmDdHhMmSs success",
			args: args{dStr: "04-13 10:20:30"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMmDdHhMmSs(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMmDdHhMmSs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMmDdHhMmSs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseMmDdHhMmSsSss(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSsSss, "0000-04-13 10:20:30.123", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "ParseMmDdHhMmSsSss error",
			args:    args{dStr: "2022-04-13 00:00:00.123"},
			wantErr: true,
		},
		{
			name: "ParseMmDdHhMmSsSss success",
			args: args{dStr: "04-13 10:20:30.123"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMmDdHhMmSsSss(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMmDdHhMmSsSss() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMmDdHhMmSsSss() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseYyyyMmDdHhMm(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSsSss, "2022-04-13 10:20:00.000", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "ParseYyyyMmDdHhMm error",
			args:    args{dStr: "2022-04-13 00:00:00.123"},
			wantErr: true,
		},
		{
			name: "ParseYyyyMmDdHhMm success",
			args: args{dStr: "2022-04-13 10:20"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseYyyyMmDdHhMm(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseYyyyMmDdHhMm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseYyyyMmDdHhMm() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseYyyyMmDdHhMmSs(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSsSss, "2022-04-13 10:20:30.000", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "ParseYyyyMmDdHhMmSs error",
			args:    args{dStr: "04-13 00:00:00.123"},
			wantErr: true,
		},
		{
			name: "ParseYyyyMmDdHhMmSs success",
			args: args{dStr: "2022-04-13 10:20:30"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseYyyyMmDdHhMmSs(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseYyyyMmDdHhMmSs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseYyyyMmDdHhMmSs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseYyyyMmDdHhMmSsSss(t *testing.T) {
	d, _ := time.ParseInLocation(YyyyMmDdHhMmSsSss, "2022-04-13 10:20:30.999", time.Local)
	type args struct {
		dStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "ParseYyyyMmDdHhMmSsSss error",
			args:    args{dStr: "04-13 00:00:00.123"},
			wantErr: true,
		},
		{
			name: "ParseYyyyMmDdHhMmSsSss success",
			args: args{dStr: "2022-04-13 10:20:30.999"},
			want: d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseYyyyMmDdHhMmSsSss(tt.args.dStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseYyyyMmDdHhMmSsSss() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseYyyyMmDdHhMmSsSss() got = %v, want %v", got, tt.want)
			}
		})
	}
}
