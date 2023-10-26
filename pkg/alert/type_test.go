package alert

import (
	"database/sql/driver"
	"reflect"
	"testing"
)

func TestKV_String(t *testing.T) {
	tests := []struct {
		name string
		l    KV
		want string
	}{
		{
			name: "test-KV-String",
			l:    KV{"a": "b"},
			want: `{"a":"b"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKV_Byte(t *testing.T) {
	tests := []struct {
		name string
		l    KV
		want []byte
	}{
		{
			name: "test-KV-Byte",
			l:    KV{"a": "b"},
			want: []byte(`{"a":"b"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Byte(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKV_Equal(t *testing.T) {
	type args struct {
		r KV
	}
	tests := []struct {
		name string
		l    KV
		args args
		want bool
	}{
		{
			name: "test-KV-Equal-true",
			l:    KV{"a": "b"},
			args: args{KV{"a": "b"}},
			want: true,
		},
		{
			name: "test-KV-Equal-false",
			l:    KV{"a": "b"},
			args: args{KV{"a": "c"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Equal(tt.args.r); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKV_Scan(t *testing.T) {
	type args struct {
		src any
	}

	tests := []struct {
		name    string
		l       *KV
		args    args
		wantErr bool
		wantSrc *KV
	}{
		{
			name:    "test-KV-Scan",
			l:       &KV{},
			args:    args{[]byte(`{"a":"b"}`)},
			wantErr: false,
			wantSrc: &KV{"a": "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.Scan(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.l, tt.wantSrc) {
				t.Errorf("Scan() got = %v, want %v", tt.l, tt.wantSrc)
			}
		})
	}
}

func TestKV_Value(t *testing.T) {
	tests := []struct {
		name    string
		l       KV
		want    driver.Value
		wantErr bool
	}{
		{
			name:    "test-KV-Value",
			l:       KV{"a": "b"},
			want:    `{"a":"b"}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() got = %v, want %v", got, tt.want)
			}
		})
	}
}
