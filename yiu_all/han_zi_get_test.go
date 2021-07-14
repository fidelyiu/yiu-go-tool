package yiuAll

import (
	"reflect"
	"testing"
)

func TestYiuHanZiGetPinYinByRune(t *testing.T) {
	type args struct {
		i rune
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "正常测试",
			args: args{
				i: 13367,
			},
			want: []string{"mà", "mián", "mǎ"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YiuHanZiGetPinYinByRune(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YiuHanZiGetPinYinByRune() = %v, want %v", got, tt.want)
			}
		})
	}
}
