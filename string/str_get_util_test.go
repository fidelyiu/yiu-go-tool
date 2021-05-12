package YiuStrUtil

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
