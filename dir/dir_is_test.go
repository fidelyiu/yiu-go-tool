package YiuDir

import (
	"testing"
)

func TestIsExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试1",
			args: args{
				path: "F:\\GoCode\\src\\yiu\\yiu-go\\dir\\dir_is.go",
			},
			want: false,
		},
		{
			name: "测试2",
			args: args{
				path: "F:\\GoCode\\src\\yiu\\yiu-go\\dir",
			},
			want: true,
		},
		{
			name: "测试2",
			args: args{
				path: "F:\\GoCode\\src\\yiu\\yiu-go\\dir1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsExists(tt.args.path); got != tt.want {
				t.Errorf("IsExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
