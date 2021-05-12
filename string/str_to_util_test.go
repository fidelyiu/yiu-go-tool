package YiuStrUtil

import (
	"reflect"
	"testing"
)

func TestToStrListByNumber(t *testing.T) {
	type args struct {
		str      string
		indexArr []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "常规测试",
			args: args{
				str:      "Hello Yiu!",
				indexArr: []int{1, 2, 3},
			},
			want: []string{"H", "el", "lo "},
		},
		{
			name: "刚好好切割完测试",
			args: args{
				str:      "Hello Yiu!",
				indexArr: []int{1, 2, 3, 4},
			},
			want: []string{"H", "el", "lo ", "Yiu!"},
		},
		{
			name: "超出切割测试",
			args: args{
				str:      "Hello Yiu!",
				indexArr: []int{1, 2, 3, 4, 5},
			},
			want: []string{"H", "el", "lo ", "Yiu!"},
		},
		{
			name: "0切割测试",
			args: args{
				str:      "Hello Yiu!",
				indexArr: []int{0, 1, 0, 2, 3, 4, 5},
			},
			want: []string{"", "H", "", "el", "lo ", "Yiu!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStrListByNumber(tt.args.str, tt.args.indexArr...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStrListByNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToByteList(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "正常测试",
			args: args{
				str: "Hello Yiu!",
			},
			want: []byte{'H', 'e', 'l', 'l', 'o', ' ', 'Y', 'i', 'u', '!'},
		},
		{
			name: "空测试",
			args: args{
				str: "",
			},
			want: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToByteList(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToByteList() = %v, want %v", got, tt.want)
			}
		})
	}
}
