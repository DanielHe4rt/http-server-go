package http

import (
	"github.com/codecrafters-io/http-server-starter-go/app/controllers"
	"testing"
)

func Test_getController(t *testing.T) {
	type args struct {
		r controllers.Request
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "With a hardcoded URL 'danielhe4rt'",
			args: args{
				r: controllers.Request{
					Verb:    "GET",
					Version: "HTTP/1.1",
					Path:    "/echo/danielhe4rt",
				},
			},
			want:    "HTTP/1.1 201 OK\r\n\r\n",
			wantErr: false,
		},
		{
			name: "With an URL variable 'danielhe4rts'",
			args: args{
				r: controllers.Request{
					Verb:    "GET",
					Version: "HTTP/1.1",
					Path:    "/echo/danielhe4rts",
				},
			},
			want:    "HTTP/1.1 200 OK\r\n\r\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getController(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("getController() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getController() got = %v, want %v", got, tt.want)
			}
		})
	}
}
