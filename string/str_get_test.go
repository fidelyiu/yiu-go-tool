package YiuStr

import "testing"

func TestGetFirstByte(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		wantErr bool
	}{
		{
			name: "正常测试",
			args: args{
				str: "Hello Yiu!",
			},
			want:    'H',
			wantErr: false,
		},
		{
			name: "空测试",
			args: args{
				str: "",
			},
			want:    ' ',
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFirstByte(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFirstByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFirstByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLastByte(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		wantErr bool
	}{
		{
			name: "正常测试",
			args: args{
				str: "Hello Yiu!",
			},
			want:    '!',
			wantErr: false,
		},
		{
			name: "空测试",
			args: args{
				str: "",
			},
			want:    ' ',
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLastByte(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLastByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLastByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTrimWithoutTarget(t *testing.T) {
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
				str:          "  --1Hello Yiu!+ ",
				targetStrArr: []string{" ", "--", "+", "1"},
			},
			want: "Hello Yiu!",
		},
		{
			name: "正常测试2",
			args: args{
				str:          "Hello Yiu!Hello Yiu!",
				targetStrArr: []string{"Hello", "--", "+", "!"},
			},
			want: " Yiu!Hello Yiu",
		},
		{
			name: "正常测试3",
			args: args{
				str:          " \n\r\n Hello Yiu!  \n\r\n  ",
				targetStrArr: []string{"\n", "\r", " "},
			},
			want: "Hello Yiu!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTrimWithoutTarget(tt.args.str, tt.args.targetStrArr...); got != tt.want {
				t.Errorf("GetTrimWithoutTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDeleteTargetStr(t *testing.T) {
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
			name: "正常测试",
			args: args{
				str:          "Hell\ro- \nYiu!",
				targetStrArr: []string{"\r", "\n", "-"},
			},
			want: "Hello Yiu!",
		},
		{
			name: "空测试",
			args: args{
				str:          "",
				targetStrArr: []string{"\r", "\n", "-"},
			},
			want: "",
		},
		{
			name: "长字符串替换",
			args: args{
				str:          "Hell\ro- \nYiu!",
				targetStrArr: []string{"\r", "\n", "Hello- "},
			},
			want: "Yiu!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDeleteTargetStr(tt.args.str, tt.args.targetStrArr...); got != tt.want {
				t.Errorf("GetDeleteTargetStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIndexAndSubByStrMore(t *testing.T) {
	type args struct {
		str        string
		subListArr []string
	}
	tests := []struct {
		name       string
		args       args
		wantIndex  int
		wantSubStr string
	}{
		{
			name: "正常测试",
			args: args{
				str:        "Hello Yiu!",
				subListArr: []string{"ell", "Yiu", "ello", ""},
			},
			wantIndex:  1,
			wantSubStr: "ell",
		},
		{
			name: "空测试",
			args: args{
				str:        "",
				subListArr: []string{"ell", "Yiu", "ello", ""},
			},
			wantIndex:  -1,
			wantSubStr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, subStr := GetIndexAndSubByStrMore(tt.args.str, tt.args.subListArr...)
			if got != tt.wantIndex {
				t.Errorf("GetIndexAndSubByStrMore() got = %v, want %v", got, tt.wantIndex)
			}
			if subStr != tt.wantSubStr {
				t.Errorf("GetIndexAndSubByStrMore() got1 = %v, want %v", subStr, tt.wantSubStr)
			}
		})
	}
}

func TestGetStrByRuneIndex(t *testing.T) {
	type args struct {
		str string
		i   int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "测试1",
			args: args{
				str: "你好呀，Hello Yiu!",
				i:   2,
			},
			want:    "呀",
			wantErr: false,
		},
		{
			name: "空测试",
			args: args{
				str: "",
				i:   2,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "越界测试",
			args: args{
				str: "1",
				i:   20,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStrByRuneIndex(tt.args.str, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStrByRuneIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetStrByRuneIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}
