package yiuAll

import "testing"

func TestYiuByteIsLetter(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试1",
			args: args{
				b: 'a',
			},
			want: true,
		},
		{
			name: "测试2",
			args: args{
				b: 'b',
			},
			want: true,
		},
		{
			name: "测试3",
			args: args{
				b: 'z',
			},
			want: true,
		},
		{
			name: "测试4",
			args: args{
				b: 'A',
			},
			want: true,
		},
		{
			name: "测试5",
			args: args{
				b: 'Z',
			},
			want: true,
		},
		{
			name: "测试6",
			args: args{
				b: '1',
			},
			want: false,
		},
		{
			name: "测试7",
			args: args{
				b: '*',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YiuByteIsLetter(tt.args.b); got != tt.want {
				t.Errorf("YiuByteIsLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteIsLowerLetter(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试1",
			args: args{
				b: 'a',
			},
			want: true,
		},
		{
			name: "测试2",
			args: args{
				b: 'b',
			},
			want: true,
		},
		{
			name: "测试3",
			args: args{
				b: 'z',
			},
			want: true,
		},
		{
			name: "测试4",
			args: args{
				b: 'A',
			},
			want: false,
		},
		{
			name: "测试5",
			args: args{
				b: 'Z',
			},
			want: false,
		},
		{
			name: "测试6",
			args: args{
				b: '1',
			},
			want: false,
		},
		{
			name: "测试7",
			args: args{
				b: '*',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YiuByteIsLowerLetter(tt.args.b); got != tt.want {
				t.Errorf("YiuByteIsLowerLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteIsUpperLetter(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试1",
			args: args{
				b: 'a',
			},
			want: false,
		},
		{
			name: "测试2",
			args: args{
				b: 'b',
			},
			want: false,
		},
		{
			name: "测试3",
			args: args{
				b: 'z',
			},
			want: false,
		},
		{
			name: "测试4",
			args: args{
				b: 'A',
			},
			want: true,
		},
		{
			name: "测试5",
			args: args{
				b: 'Z',
			},
			want: true,
		},
		{
			name: "测试6",
			args: args{
				b: '1',
			},
			want: false,
		},
		{
			name: "测试7",
			args: args{
				b: '*',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YiuByteIsUpperLetter(tt.args.b); got != tt.want {
				t.Errorf("YiuByteIsUpperLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteIsNotLetter(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试1",
			args: args{
				b: 'a',
			},
			want: false,
		},
		{
			name: "测试2",
			args: args{
				b: 'b',
			},
			want: false,
		},
		{
			name: "测试3",
			args: args{
				b: 'z',
			},
			want: false,
		},
		{
			name: "测试4",
			args: args{
				b: 'A',
			},
			want: false,
		},
		{
			name: "测试5",
			args: args{
				b: 'Z',
			},
			want: false,
		},
		{
			name: "测试6",
			args: args{
				b: '1',
			},
			want: true,
		},
		{
			name: "测试7",
			args: args{
				b: '*',
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YiuByteIsNotLetter(tt.args.b); got != tt.want {
				t.Errorf("YiuByteIsNotLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
