package yiuAll

import "testing"

func TestYiuRuneListIsLeByUnicodeAndLowerBeforeUpper(t *testing.T) {
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
				listA: YiuStrToRuneList("yiuVar"),
				listB: YiuStrToRuneList("yiuAll"),
			},
			want: false,
		},
		{
			name: "t2",
			args: args{
				listA: YiuStrToRuneList("yiuvar"),
				listB: YiuStrToRuneList("yiuAll"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YiuRuneListIsLeByUnicodeAndLowerBeforeUpper(tt.args.listA, tt.args.listB); got != tt.want {
				t.Errorf("YiuRuneListIsLeByUnicodeAndLowerBeforeUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuRuneListIsGtByUnicodeAndLowerBeforeUpper(t *testing.T) {
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
				listA: YiuStrToRuneList("yiuVar"),
				listB: YiuStrToRuneList("yiuAll"),
			},
			want: true,
		},
		{
			name: "t2",
			args: args{
				listA: YiuStrToRuneList("yiuvar"),
				listB: YiuStrToRuneList("yiuAll"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YiuRuneListIsGtByUnicodeAndLowerBeforeUpper(tt.args.listA, tt.args.listB); got != tt.want {
				t.Errorf("YiuRuneListIsGtByUnicodeAndLowerBeforeUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuRuneListIsLeByUnicode(t *testing.T) {
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
				listA: YiuStrToRuneList("IsTypeAndroid"),
				listB: YiuStrToRuneList("IsTypeLinux"),
			},
			want: true,
		},
		{
			name: "t1",
			args: args{
				listA: YiuStrToRuneList("IsGoarchMips64p32"),
				listB: YiuStrToRuneList("IsGoarchMips64p32le"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YiuRuneListIsLeByUnicode(tt.args.listA, tt.args.listB); got != tt.want {
				t.Errorf("YiuRuneListIsLeByUnicode() = %v, want %v", got, tt.want)
			}
		})
	}
}
