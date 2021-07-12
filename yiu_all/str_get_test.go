package yiuAll

import "testing"

func TestYiuStrGetTrimLeftSStr(t *testing.T) {
	type args struct {
		str          string
		targetStrArr []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "正常测试1",
			args: args{
				str: "	◎ 测试字符串",
				targetStrArr: []string{
					" ",
				},
			},
			want: "\t◎ 测试字符串",
		},
		{
			name: "正常测试2",
			args: args{
				str: "	◎ 测试字符串",
				targetStrArr: []string{
					" ",
					"\t",
				},
			},
			want: "◎ 测试字符串",
		},
		{
			name: "正常测试3",
			args: args{
				str: "	◎ 测试字符串",
				targetStrArr: []string{
					" ",
					"◎",
					"\t",
				},
			},
			want: "测试字符串",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YiuStrGetTrimLeftSStr(tt.args.str, tt.args.targetStrArr...); got != tt.want {
				t.Errorf("YiuStrGetTrimLeftSStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
