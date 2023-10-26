package alert

import (
	"reflect"
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	type args struct {
		timeStr string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "test-ok",
			args: args{
				timeStr: "2021-01-01T00:00:00Z",
			},
			want: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "test-err",
			args: args{
				timeStr: "2x21-01-01T00:00:00Z",
			},
			want: time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseTime(tt.args.timeStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTime() = %v, want %v, %d", got, tt.want, got.Unix())
			}
		})
	}
}
