package YiuStrList

import (
	"fmt"
	YiuInt "github.com/fidelyiu/yiu-go/int"
	"testing"
)

func TestOpMap(t *testing.T) {
	type args struct {
		list   *[]string
		opFunc func(int, string) string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "正常测试",
			args: args{
				list: &[]string{"Hello", "Fidel", "Yiu"},
				opFunc: func(i int, s string) string {
					return s + "(" + YiuInt.ToStr(i) + ")"
				},
			},
		},
		{
			name: "空测试",
			args: args{
				list: &[]string{},
				opFunc: func(i int, s string) string {
					return s + "(" + YiuInt.ToStr(i) + ")"
				},
			},
		},
		{
			name: "nil测试",
			args: args{
				list: nil,
				opFunc: func(i int, s string) string {
					return s + "(" + YiuInt.ToStr(i) + ")"
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name, tt.args.list)
			OpMap(tt.args.list, tt.args.opFunc)
			fmt.Println(tt.name, tt.args.list)
		})
	}
}
