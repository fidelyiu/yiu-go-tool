package YiuInt

import "testing"

func TestIsIntersect(t *testing.T) {
	type args struct {
		a         int
		b         int
		c         int
		d         int
		isContain bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试01",
			args: args{
				a:         1,
				b:         1,
				c:         1,
				d:         1,
				isContain: true,
			},
			want: true,
		},
		{
			name: "测试02",
			args: args{
				a:         1,
				b:         6,
				c:         1,
				d:         0,
				isContain: false,
			},
			want: false,
		},
		{
			name: "测试03",
			args: args{
				a:         1,
				b:         6,
				c:         1,
				d:         0,
				isContain: true,
			},
			want: true,
		},
		{
			name: "测试04",
			args: args{
				a:         1,
				b:         6,
				c:         3,
				d:         7,
				isContain: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIntersect(tt.args.a, tt.args.b, tt.args.c, tt.args.d, tt.args.isContain); got != tt.want {
				t.Errorf("IsIntersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
