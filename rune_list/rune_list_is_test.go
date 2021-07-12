package yiuRuneList

import (
	yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"
	"testing"
)

func TestIsLeByUnicodeAndLowerBeforeUpper(t *testing.T) {
	type args struct {
		listA []rune
		listB []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "t1",
			args: args{
				listA: yiuAll.YiuStrToRuneList("yiuVar"),
				listB: yiuAll.YiuStrToRuneList("yiuAll"),
			},
			want: false,
		},
		{
			name: "t2",
			args: args{
				listA: yiuAll.YiuStrToRuneList("yiuvar"),
				listB: yiuAll.YiuStrToRuneList("yiuAll"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLeByUnicodeAndLowerBeforeUpper(tt.args.listA, tt.args.listB); got != tt.want {
				t.Errorf("IsLeByUnicodeAndLowerBeforeUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsGtByUnicodeAndLowerBeforeUpper(t *testing.T) {
	type args struct {
		listA []rune
		listB []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "t1",
			args: args{
				listA: yiuAll.YiuStrToRuneList("yiuVar"),
				listB: yiuAll.YiuStrToRuneList("yiuAll"),
			},
			want: true,
		},
		{
			name: "t2",
			args: args{
				listA: yiuAll.YiuStrToRuneList("yiuvar"),
				listB: yiuAll.YiuStrToRuneList("yiuAll"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsGtByUnicodeAndLowerBeforeUpper(tt.args.listA, tt.args.listB); got != tt.want {
				t.Errorf("IsGtByUnicodeAndLowerBeforeUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLeByUnicode(t *testing.T) {
	type args struct {
		listA []rune
		listB []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "t1",
			args: args{
				listA: yiuAll.YiuStrToRuneList("IsTypeAndroid"),
				listB: yiuAll.YiuStrToRuneList("IsTypeLinux"),
			},
			want: true,
		},
		{
			name: "t1",
			args: args{
				listA: yiuAll.YiuStrToRuneList("IsGoarchMips64p32"),
				listB: yiuAll.YiuStrToRuneList("IsGoarchMips64p32le"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLeByUnicode(tt.args.listA, tt.args.listB); got != tt.want {
				t.Errorf("IsLeByUnicode() = %v, want %v", got, tt.want)
			}
		})
	}
}
