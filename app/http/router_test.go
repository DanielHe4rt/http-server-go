package http

import (
	"github.com/codecrafters-io/http-server-starter-go/app/http/controllers"
	"github.com/codecrafters-io/http-server-starter-go/app/http/request"
	"github.com/codecrafters-io/http-server-starter-go/app/http/response"
	"os"
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
			name: "Insert new File",
			args: args{
				r: request.Request{
					Verb:    "GET",
					Version: "HTTP/1.1",
					Path:    "/files/{fileName}",
					Headers: nil,
					Params:  map[string]string{},
				},
			},
			want: "FilesController",
		},
		{
			name: "Insert new File",
			args: args{
				r: request.Request{
					Verb:    "POST",
					Version: "HTTP/1.1",
					Path:    "/files/{fileName}",
					Headers: nil,
					Params:  map[string]string{},
				},
			},
			want: "FileUploaderController",
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
			},
			want: controllers.MeController{},
		},
		{
			name: "With an URL variable 'danielhe4rts'",
			args: args{
				controller: "EchoController",
			},
			want: controllers.EchoController{},
		},
		{
			name: "Retrieve FilesUploadController",
			args: args{
				controller: "FileUploaderController",
			},
			want: controllers.FileUploaderController{},
		},
		{
			name: "Invalid URL brings 404 NotFoundController",
			args: args{
				controller: "",
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

func TestProcessRequestFiles(t *testing.T) {
	os.Args = append(os.Args, "--directory")
	os.Args = append(os.Args, "/tmp/")
	_ = os.WriteFile("/tmp/fodase", []byte(""), 0644)
	resDownload := response.Response{}
	resUpload := response.Response{}
	type args struct {
		r request.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "File Download",
			args: args{
				r: request.Request{
					Verb:    "GET",
					Version: "HTTP/1.1",
					Path:    "/files/fodase",
					Headers: nil,
					Params:  map[string]string{},
					Body:    "",
				},
			},
			want: resDownload.Download("/tmp/", "fodase").Build(),
		},
		{
			name: "File Upload",
			args: args{
				r: request.Request{
					Verb:    "POST",
					Version: "HTTP/1.1",
					Path:    "/files/123",
					Headers: map[string]string{
						"Content-Length": "3",
						"Content-Type":   "application/octet-stream",
						"User-Agent":     "curl/8.5.0",
						"Accept":         "*/*",
					},
					Params: map[string]string{"fileName": "123"},
					Body:   "123",
				},
			},
			want: resUpload.Created().Build(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.r.Verb == request.VerbPost {
				_, err := os.Stat("/tmp/123")

				if err != nil {
					t.Errorf("File Not found at /tmp/")
				}
			}

			if got := ProcessRequest(tt.args.r); got != tt.want {
				t.Errorf("ProcessRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessRequestCompression(t *testing.T) {
	//os.Args = append(os.Args, "--directory")
	//os.Args = append(os.Args, "/tmp/")
	//_ = os.WriteFile("/tmp/fodase", []byte(""), 0644)
	type args struct {
		r request.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Compression Headers with Gzip",
			args: args{
				r: request.Request{
					Verb:    "GET",
					Version: "HTTP/1.1",
					Path:    "/echo/fodase",
					Headers: map[string]string{
						"Accept-Encoding": "gzip",
					},
					Params: map[string]string{
						"message": "fodase",
					},
					Body: "",
				},
			},
			want: response.New().Compress("gzip", "fodase").Success().Build(),
		},
		{
			name: "Compression Headers with Gzip",
			args: args{
				r: request.Request{
					Verb:    "GET",
					Version: "HTTP/1.1",
					Path:    "/echo/fodase",
					Headers: map[string]string{
						"Accept-Encoding": "gzip, fodase",
					},
					Params: map[string]string{
						"message": "fodase",
					},
					Body: "",
				},
			},
			want: response.New().Compress("gzip, fodase, caguei", "fodase").Success().Build(),
		},
		{
			name: "Unsupported Compression Header",
			args: args{
				r: request.Request{
					Verb:    "GET",
					Version: "HTTP/1.1",
					Path:    "/echo/fodase",
					Headers: map[string]string{
						"Accept-Encoding": "fodase",
					},
					Params: map[string]string{
						"message": "fodase",
					},
					Body: "",
				},
			},
			want: response.New().Compress("fodase", "fodase").Success().Build(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.r.Verb == request.VerbPost {
				_, err := os.Stat("/tmp/123")

				if err != nil {
					t.Errorf("File Not found at /tmp/")
				}
			}

			if got := ProcessRequest(tt.args.r); got != tt.want {
				t.Errorf("ProcessRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
