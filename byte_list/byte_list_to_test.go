package YiuByteList

import "testing"

func TestToStr(t *testing.T) {
	type args struct {
		list []byte
		sep  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "正常测试1",
			args: args{
				list: []byte{'a', 'b'},
				sep:  " ",
			},
			want: "a b",
		},
		{
			name: "正常测试2",
			args: args{
				list: []byte{'1', '6', '9'},
				sep:  "-",
			},
			want: "1-6-9",
		},
		{
			name: "空切片测试",
			args: args{
				list: []byte{},
				sep:  "-",
			},
			want: "",
		},
		{
			name: "单切片测试",
			args: args{
				list: []byte{'H'},
				sep:  "=",
			},
			want: "H",
		},
		{
			name: "空分隔符测试",
			args: args{
				list: []byte{'a', 'b'},
				sep:  "",
			},
			want: "ab",
		},
		{
			name: "nil测试",
			args: args{
				list: nil,
				sep:  "=",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStr(tt.args.list, tt.args.sep); got != tt.want {
				t.Errorf("ToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
