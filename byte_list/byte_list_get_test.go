package YiuByteList

import "testing"

func TestGetIndexByEl(t *testing.T) {
	type args struct {
		list []byte
		el   byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "正常测试",
			args: args{
				list: []byte{'a', 'b', 'c', 'd'},
				el:   'c',
			},
			want: 2,
		},
		{
			name: "空测试",
			args: args{
				list: []byte{},
				el:   'c',
			},
			want: -1,
		},
		{
			name: "nil测试",
			args: args{
				list: nil,
				el:   'c',
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIndexByEl(tt.args.list, tt.args.el); got != tt.want {
				t.Errorf("GetIndexByEl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIndexByList(t *testing.T) {
	type args struct {
		list    []byte
		subList []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "正常测试",
			args: args{
				list:    []byte{'a', 'b', 'c', 'd'},
				subList: []byte{'b', 'c'},
			},
			want: 1,
		},
		{
			name: "-1测试",
			args: args{
				list:    []byte{'a', 'b', 'c', 'd'},
				subList: []byte{'y', 'c'},
			},
			want: -1,
		},
		{
			name: "空List测试",
			args: args{
				list:    []byte{},
				subList: []byte{'y', 'c'},
			},
			want: -1,
		},
		{
			name: "nilList测试",
			args: args{
				list:    nil,
				subList: []byte{'y', 'c'},
			},
			want: -1,
		},
		{
			name: "空subList测试",
			args: args{
				list:    []byte{'a', 'b', 'c', 'd'},
				subList: []byte{},
			},
			want: -1,
		},
		{
			name: "nilSubList测试",
			args: args{
				list:    []byte{'a', 'b', 'c', 'd'},
				subList: nil,
			},
			want: -1,
		},
		{
			name: "全空测试",
			args: args{
				list:    []byte{},
				subList: []byte{},
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

func TestGetElByIndex(t *testing.T) {
	type args struct {
		list  []byte
		index int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "正常测试",
			args: args{
				list:  []byte{'a', 'b', 'c', 'd'},
				index: 1,
			},
			want: 'b',
		},
		{
			name: "负值测试",
			args: args{
				list:  []byte{'a', 'b', 'c', 'd'},
				index: -1,
			},
			want: 0,
		},
		{
			name: "超索引测试",
			args: args{
				list:  []byte{'a', 'b', 'c', 'd'},
				index: 4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetElByIndex(tt.args.list, tt.args.index); got != tt.want {
				t.Errorf("GetElByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
