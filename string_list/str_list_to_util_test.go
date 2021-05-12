package YiuStrListUtil

import "testing"

func TestToStr(t *testing.T) {
	type args struct {
		list []string
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
				list: []string{"Hello", "Yiu"},
				sep:  " ",
			},
			want: "Hello Yiu",
		},
		{
			name: "正常测试2",
			args: args{
				list: []string{"123", "654", "963"},
				sep:  "-",
			},
			want: "123-654-963",
		},
		{
			name: "空切片测试",
			args: args{
				list: []string{},
				sep:  "-",
			},
			want: "",
		},
		{
			name: "单切片测试",
			args: args{
				list: []string{"Hello"},
				sep:  "=",
			},
			want: "Hello",
		},
		{
			name: "空分隔符测试",
			args: args{
				list: []string{"Hello", "Yiu"},
				sep:  "",
			},
			want: "HelloYiu",
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
