package YiuStrListUtil

import (
	"reflect"
	"testing"
)

func TestGetDeleteEmptyEl(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "正常测试1",
			args: args{
				list: []string{"", "Hello", "", "Yiu"},
			},
			want: []string{"Hello", "Yiu"},
		},
		{
			name: "正常测试2",
			args: args{
				list: []string{" ", "Hello", "", "Yiu"},
			},
			want: []string{" ", "Hello", "Yiu"},
		},
		{
			name: "空测试",
			args: args{
				list: []string{},
			},
			want: []string{},
		},
		{
			name: "nil测试",
			args: args{
				list: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDeleteEmptyEl(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeleteEmptyEl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDeleteBlankEl(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "正常测试1",
			args: args{
				list: []string{"", "Hello", "", "Yiu"},
			},
			want: []string{"Hello", "Yiu"},
		},
		{
			name: "正常测试2",
			args: args{
				list: []string{" ", "Hello", "", "Yiu"},
			},
			want: []string{"Hello", "Yiu"},
		},
		{
			name: "空测试",
			args: args{
				list: []string{},
			},
			want: []string{},
		},
		{
			name: "nil测试",
			args: args{
				list: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDeleteBlankEl(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeleteBlankEl() = %v, want %v", got, tt.want)
			}
		})
	}
}
