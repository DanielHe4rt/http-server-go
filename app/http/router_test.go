package http

import (
	"github.com/codecrafters-io/http-server-starter-go/app/http/controllers"
	"github.com/codecrafters-io/http-server-starter-go/app/http/request"
	"testing"
)

func Test_getController(t *testing.T) {
	type args struct {
		r request.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "With a hardcoded URL 'danielhe4rt'",
			args: args{
				r: request.Request{
					Verb:    "GET",
					Version: "HTTP/1.1",
					Path:    "/echo/danielhe4rt",
					Params:  map[string]string{},
				},
			},
			want: "MeController",
		},
		{
			name: "With an URL variable 'danielhe4rts'",
			args: args{
				r: request.Request{
					Verb:    "GET",
					Version: "HTTP/1.1",
					Path:    "/echo/danielhe4rts",
					Headers: nil,
					Params:  map[string]string{},
				},
			},
			want: "EchoController",
		},
		{
			name: "Invalid URL brings 404 NotFoundController",
			args: args{
				r: request.Request{
					Verb:    "GET",
					Version: "HTTP/1.1",
					Path:    "/something-cool",
					Params:  map[string]string{},
				},
			},
			want: "NotFoundController",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := getController(tt.args.r); got != tt.want {
				t.Errorf("getController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRouteAction(t *testing.T) {
	type args struct {
		controller string
		request    request.Request
	}
	tests := []struct {
		name string
		args args
		want controllers.BaseController
	}{
		{
			name: "With a hardcoded URL 'danielhe4rt'",
			args: args{
				controller: "MeController",
				request: request.Request{
					Verb:    "GET",
					Version: "HTTP/1.1",
					Path:    "/echo/danielhe4rt",
				},
			},
			want: controllers.MeController{},
		},
		{
			name: "With an URL variable 'danielhe4rts'",
			args: args{
				controller: "EchoController",
				request: request.Request{
					Verb:    "GET",
					Version: "HTTP/1.1",
					Path:    "/echo/danielhe4rts",
				},
			},
			want: controllers.EchoController{},
		},
		{
			name: "Invalid URL brings 404 NotFoundController",
			args: args{
				controller: "",
				request: request.Request{
					Verb:    "GET",
					Version: "HTTP/1.1",
					Path:    "/something-cool",
				},
			},
			want: controllers.NotFoundController{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRouteAction(tt.args.controller); got != tt.want {
				t.Errorf("getRouteAction() = %v, want %v", got, tt.want)
			}
		})
	}
}
