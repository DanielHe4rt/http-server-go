package request

import (
	"reflect"
	"testing"
)

func TestNewRequest(t *testing.T) {
	type args struct {
		payload []byte
	}
	tests := []struct {
		name string
		args args
		want Request
	}{
		{
			name: "Base Request Argless",
			args: args{
				payload: []byte("GET /index.html HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1\r\nAccept: */*\r\n\r\n"),
			},
			want: Request{
				Verb:    VerbGet,
				Version: "HTTP/1.1",
				Path:    "/index.html",
				Headers: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRequest(tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
