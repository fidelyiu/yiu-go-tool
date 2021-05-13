package YiuByteList

import "testing"

func TestIsContainsElList(t *testing.T) {
	type args struct {
		list    []byte
		subList []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "正常测试",
			args: args{
				list:    []byte{'a', 'b', 'c', 'd'},
				subList: []byte{'b', 'c'},
			},
			want: true,
		},
		{
			name: "异常测试",
			args: args{
				list:    []byte{'a', 'b', 'c', 'd'},
				subList: []byte{'y', 'c'},
			},
			want: false,
		},
		{
			name: "空List测试",
			args: args{
				list:    []byte{},
				subList: []byte{'b', 'c'},
			},
			want: false,
		},
		{
			name: "nilList测试",
			args: args{
				list:    nil,
				subList: []byte{'b', 'c'},
			},
			want: false,
		},
		{
			name: "空subList测试",
			args: args{
				list:    []byte{'b', 'c'},
				subList: []byte{},
			},
			want: false,
		},
		{
			name: "nilSubList测试",
			args: args{
				list:    []byte{'b', 'c'},
				subList: nil,
			},
			want: false,
		},
		{
			name: "全空测试",
			args: args{
				list:    []byte{},
				subList: []byte{},
			},
			want: false,
		},
		{
			name: "全nil测试",
			args: args{
				list:    nil,
				subList: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsContainsElList(tt.args.list, tt.args.subList); got != tt.want {
				t.Errorf("IsContainsElList() = %v, want %v", got, tt.want)
			}
		})
	}
}
