package curl

import (
	"context"
	"testing"
)

func TestCurl(t *testing.T) {
	type args struct {
		ctx  context.Context
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test-baidu",
			args: args{
				ctx:  context.Background(),
				path: "https://www.baidu.com",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Curl(tt.args.ctx, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Curl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Curl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
