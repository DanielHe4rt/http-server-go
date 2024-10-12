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

func TestRequest_GetKey(t *testing.T) {
	type fields struct {
		Verb    VerbType
		Version string
		Path    string
		headers map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Request{
				Verb:    tt.fields.Verb,
				Version: tt.fields.Version,
				Path:    tt.fields.Path,
				Headers: tt.fields.headers,
			}
			if got := r.GetKey(); got != tt.want {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractRequestLine(t *testing.T) {
	type args struct {
		requestLine []byte
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := extractRequestLine(tt.args.requestLine)
			if got != tt.want {
				t.Errorf("extractRequestLine() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("extractRequestLine() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("extractRequestLine() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
