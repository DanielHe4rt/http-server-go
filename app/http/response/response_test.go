package response

import "testing"

func TestResponse_Build(t *testing.T) {
	type fields struct {
		body       string
		status     int
		statusText string
		headers    map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Successful Request with 200 and 4 length",
			fields: fields{
				body:       "test",
				status:     200,
				statusText: "OK",
				headers:    map[string]string{"Content-Type": "text/plain", "Content-Length": "4"},
			},
			want: "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 4\r\n\r\ntest",
		},
		{
			name: "NotFound Request with 404 and 0 length",
			fields: fields{
				body:       "",
				status:     404,
				statusText: "Not Found",
				headers:    nil,
			},
			want: "HTTP/1.1 404 Not Found\r\n\r\n",
		},
		{
			name: "Created new resource",
			fields: fields{
				body:       "",
				status:     201,
				statusText: "Created",
				headers:    nil,
			},
			want: "HTTP/1.1 201 Created\r\n\r\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Response{
				body:       tt.fields.body,
				status:     tt.fields.status,
				statusText: tt.fields.statusText,
				headers:    tt.fields.headers,
			}
			if got := res.Build(); got != tt.want {
				t.Errorf("Build() = %v (%v), want %v (%v)", got, len(got), tt.want, len(tt.want))
			}
		})
	}
}
