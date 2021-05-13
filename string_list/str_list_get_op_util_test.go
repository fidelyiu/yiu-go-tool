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

func TestGetIndexByList(t *testing.T) {
	type args struct {
		list    []string
		subList []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "正常测试",
			args: args{
				list:    []string{"a", "b", "c", "d"},
				subList: []string{"b", "c"},
			},
			want: 1,
		},
		{
			name: "-1测试",
			args: args{
				list:    []string{"a", "b", "c", "d"},
				subList: []string{"y", "c"},
			},
			want: -1,
		},
		{
			name: "空List测试",
			args: args{
				list:    []string{},
				subList: []string{"y", "c"},
			},
			want: -1,
		},
		{
			name: "nilList测试",
			args: args{
				list:    nil,
				subList: []string{"y", "c"},
			},
			want: -1,
		},
		{
			name: "空subList测试",
			args: args{
				list:    []string{"a", "b"},
				subList: []string{},
			},
			want: -1,
		},
		{
			name: "nilSubList测试",
			args: args{
				list:    []string{"a", "b"},
				subList: nil,
			},
			want: -1,
		},
		{
			name: "全空测试",
			args: args{
				list:    []string{},
				subList: []string{},
			},
			want: -1,
		},
		{
			name: "全nil测试",
			args: args{
				list:    nil,
				subList: nil,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIndexByList(tt.args.list, tt.args.subList); got != tt.want {
				t.Errorf("GetIndexByList() = %v, want %v", got, tt.want)
			}
		})
	}
}
