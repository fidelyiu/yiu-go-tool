package YiuStr

import "testing"

func TestIsLetter(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试1",
			args: args{
				str: "",
			},
			want: false,
		},
		{
			name: "测试2",
			args: args{
				str: "A",
			},
			want: true,
		},
		{
			name: "测试3",
			args: args{
				str: "你",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLetter(tt.args.str); got != tt.want {
				t.Errorf("IsLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
