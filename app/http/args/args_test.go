package args

import (
	"os"
	"reflect"
	"testing"
)

func TestGetArgs(t *testing.T) {

	tests := []struct {
		name string
		args []string
		want RunnerArgs
	}{
		{
			name: "ok",
			args: []string{"--Directory", "/tmp/"},
			want: RunnerArgs{
				Directory: "/tmp/",
			},
		},
		{
			name: "error",
			args: []string{"--Directory", "/fodase"},
			want: RunnerArgs{
				Directory: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, value := range tt.args {
				os.Args = append(os.Args, value)
			}

			if got := GetArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
