package yiuByteList

import (
	yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"
	"reflect"
	"strings"
	"testing"
)

func TestGetDeduplicate(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "去重测试1",
			args: args{
				list: yiuAll.YiuStrToByteList("lllHello"),
			},
			want: yiuAll.YiuStrToByteList("lHeo"),
		},
		{
			name: "去重测试2",
			args: args{
				list: yiuAll.YiuStrToByteList("lllHHHHelloOOO"),
			},
			want: yiuAll.YiuStrToByteList("lHeoO"),
		},
		{
			name: "去重测试3",
			args: args{
				list: []byte{1, 2, 3, 1, 5, 4},
			},
			want: []byte{1, 2, 3, 5, 4},
		},
		{
			name: "去重测试4",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
			},
			want: yiuAll.YiuStrToByteList("Helo Yiu!"),
		},
		{
			name: "去重测试5-空去重",
			args: args{
				list: []byte{},
			},
			want: []byte{},
		},
		{
			name: "去重测试6-nil去重",
			args: args{
				list: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDeduplicate(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeduplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDeleteByIndex(t *testing.T) {
	type args struct {
		list     []byte
		delIndex int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "删除测试1-正常删除",
			args: args{
				list:     yiuAll.YiuStrToByteList("Hello Yiu!"),
				delIndex: 0,
			},
			want: yiuAll.YiuStrToByteList("ello Yiu!"),
		},
		{
			name: "删除测试2-正常删除",
			args: args{
				list:     yiuAll.YiuStrToByteList("Hello Yiu!"),
				delIndex: len("Hello Yiu!") - 1,
			},
			want: yiuAll.YiuStrToByteList("Hello Yiu"),
		},
		{
			name: "删除测试3-正常删除",
			args: args{
				list:     yiuAll.YiuStrToByteList("Hello Yiu!"),
				delIndex: strings.Index("Hello Yiu!", " "),
			},
			want: yiuAll.YiuStrToByteList("HelloYiu!"),
		},
		{
			name: "删除测试4-负值删除",
			args: args{
				list:     yiuAll.YiuStrToByteList("Hello Yiu!"),
				delIndex: -1,
			},
			want: yiuAll.YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "删除测试5-越界删除",
			args: args{
				list:     yiuAll.YiuStrToByteList("Hello Yiu!"),
				delIndex: 999,
			},
			want: yiuAll.YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "删除测试6-空删除",
			args: args{
				list:     []byte{},
				delIndex: 0,
			},
			want: []byte{},
		},
		{
			name: "删除测试7-nil删除",
			args: args{
				list:     nil,
				delIndex: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDeleteByIndex(tt.args.list, tt.args.delIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeleteByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDeleteByRangeIndex(t *testing.T) {
	type args struct {
		list       []byte
		startIndex int
		endIndex   int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "范围删除1-正常删除",
			args: args{
				list:       yiuAll.YiuStrToByteList("Hello Yiu!"),
				startIndex: 0,
				endIndex:   0,
			},
			want: yiuAll.YiuStrToByteList("ello Yiu!"),
		},
		{
			name: "范围删除2-正常删除",
			args: args{
				list:       yiuAll.YiuStrToByteList("Hello Yiu!"),
				startIndex: 0,
				endIndex:   1,
			},
			want: yiuAll.YiuStrToByteList("llo Yiu!"),
		},
		{
			name: "范围删除3-正常删除",
			args: args{
				list:       yiuAll.YiuStrToByteList("Hello Yiu!"),
				startIndex: strings.Index("Hello Yiu!", " "),
				endIndex:   strings.Index("Hello Yiu!", "u"),
			},
			want: yiuAll.YiuStrToByteList("Hello!"),
		},
		{
			name: "范围删除4-负值删除",
			args: args{
				list:       yiuAll.YiuStrToByteList("Hello Yiu!"),
				startIndex: -1,
				endIndex:   strings.Index("Hello Yiu!", "u"),
			},
			want: yiuAll.YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "范围删除5-负值删除",
			args: args{
				list:       yiuAll.YiuStrToByteList("Hello Yiu!"),
				startIndex: -1,
				endIndex:   -13,
			},
			want: yiuAll.YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "范围删除6-负值删除",
			args: args{
				list:       yiuAll.YiuStrToByteList("Hello Yiu!"),
				startIndex: 2,
				endIndex:   -13,
			},
			want: yiuAll.YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "范围删除7-越界删除",
			args: args{
				list:       yiuAll.YiuStrToByteList("Hello Yiu!"),
				startIndex: 999,
				endIndex:   3,
			},
			want: yiuAll.YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "范围删除8-越界删除",
			args: args{
				list:       yiuAll.YiuStrToByteList("Hello Yiu!"),
				startIndex: 3,
				endIndex:   999,
			},
			want: yiuAll.YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "范围删除9-越界删除",
			args: args{
				list:       yiuAll.YiuStrToByteList("Hello Yiu!"),
				startIndex: 998,
				endIndex:   999,
			},
			want: yiuAll.YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "范围删除10-开始索引大于结束索引删除",
			args: args{
				list:       yiuAll.YiuStrToByteList("Hello Yiu!"),
				startIndex: 3,
				endIndex:   2,
			},
			want: yiuAll.YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "范围删除11-空删除",
			args: args{
				list:       []byte{},
				startIndex: 0,
				endIndex:   0,
			},
			want: []byte{},
		},
		{
			name: "范围删除12-nil删除",
			args: args{
				list:       nil,
				startIndex: 0,
				endIndex:   0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDeleteByRangeIndex(tt.args.list, tt.args.startIndex, tt.args.endIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeleteByRangeIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFilter(t *testing.T) {
	type args struct {
		list []byte
		keep func(x byte) bool
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "过滤测试1",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				keep: func(x byte) bool {
					return x == 'Y' || x == 'i' || x == 'u'
				},
			},
			want: yiuAll.YiuStrToByteList("Yiu"),
		},
		{
			name: "过滤测试2",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				keep: func(x byte) bool {
					return false
				},
			},
			want: yiuAll.YiuStrToByteList(""),
		},
		{
			name: "过滤测试3",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				keep: func(x byte) bool {
					return true
				},
			},
			want: yiuAll.YiuStrToByteList("Hello Yiu!"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFilter(tt.args.list, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPop(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		want1   []byte
		wantErr bool
	}{
		{
			name: "出栈测试1-正常出栈",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
			},
			want:    '!',
			want1:   yiuAll.YiuStrToByteList("Hello Yiu"),
			wantErr: false,
		},
		{
			name: "出栈测试2-空出栈",
			args: args{
				list: []byte{},
			},
			want:    0,
			want1:   []byte{},
			wantErr: true,
		},
		{
			name: "出栈测试2-nil出栈",
			args: args{
				list: nil,
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GetPop(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetPop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetReverse(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "反序测试1",
			args: args{
				list: []byte{1, 2, 3, 4, 5, 6},
			},
			want: []byte{6, 5, 4, 3, 2, 1},
		},
		{
			name: "反序测试2",
			args: args{
				list: yiuAll.YiuStrToByteList("Yiu!"),
			},
			want: yiuAll.YiuStrToByteList("!uiY"),
		},
		{
			name: "反序测试3-空反序",
			args: args{
				list: []byte{},
			},
			want: []byte{},
		},
		{
			name: "反序测试4-nil反序",
			args: args{
				list: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReverse(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetShuffle(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "乱序测试1",
			args: args{
				list: []byte{1},
			},
			want: []byte{1},
		},
		{
			name: "乱序测试2-空乱序",
			args: args{
				list: []byte{},
			},
			want: []byte{},
		},
		{
			name: "乱序测试3-nil乱序",
			args: args{
				list: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetShuffle(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetShuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMap(t *testing.T) {
	type args struct {
		list   []byte
		opFunc func(int, byte) byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Map测试1",
			args: args{
				list: []byte{1, 2, 3, 4},
				opFunc: func(i int, b byte) byte {
					return byte(i+1) + b
				},
			},
			want: []byte{2, 4, 6, 8},
		},
		{
			name: "Map测试2-空测试",
			args: args{
				list: []byte{},
				opFunc: func(i int, b byte) byte {
					return byte(i+1) + b
				},
			},
			want: []byte{},
		},
		{
			name: "Map测试2-nil测试",
			args: args{
				list: nil,
				opFunc: func(i int, b byte) byte {
					return byte(i+1) + b
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMap(tt.args.list, tt.args.opFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCopy(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "拷贝测试1",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
			},
			want: yiuAll.YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "拷贝测试2-空测试",
			args: args{
				list: []byte{},
			},
			want: []byte{},
		},
		{
			name: "拷贝测试2-nil测试",
			args: args{
				list: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCopy(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			name: "索引元素测试1-正常索引",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				el:   'l',
			},
			want: strings.Index("Hello Yiu!", "l"),
		},
		{
			name: "索引元素测试2-无值索引",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				el:   '=',
			},
			want: -1,
		},
		{
			name: "索引元素测试3-空索引",
			args: args{
				list: []byte{},
				el:   'l',
			},
			want: -1,
		},
		{
			name: "索引元素测试4-nil索引",
			args: args{
				list: nil,
				el:   'l',
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
			name: "多值索引1",
			args: args{
				list:    yiuAll.YiuStrToByteList("Hello Yiu!"),
				subList: yiuAll.YiuStrToByteList("ell"),
			},
			want: 1,
		},
		{
			name: "多值索引2",
			args: args{
				list:    yiuAll.YiuStrToByteList("Hello Yiu!"),
				subList: yiuAll.YiuStrToByteList("Yiu"),
			},
			want: 6,
		},
		{
			name: "多值索引3-空索引",
			args: args{
				list:    []byte{},
				subList: yiuAll.YiuStrToByteList("Yiu"),
			},
			want: -1,
		},
		{
			name: "多值索引4-空索引",
			args: args{
				list:    []byte{},
				subList: []byte{},
			},
			want: -1,
		},
		{
			name: "多值索引5-nil索引",
			args: args{
				list:    nil,
				subList: yiuAll.YiuStrToByteList("Yiu"),
			},
			want: -1,
		},
		{
			name: "多值索引6-nil索引",
			args: args{
				list:    nil,
				subList: nil,
			},
			want: -1,
		},
		{
			name: "多值索引7-索引空",
			args: args{
				list:    yiuAll.YiuStrToByteList("Yiu"),
				subList: []byte{},
			},
			want: -1,
		},
		{
			name: "多值索引8-索引nil",
			args: args{
				list:    yiuAll.YiuStrToByteList("Yiu"),
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

func TestGetIndexBySList(t *testing.T) {
	type args struct {
		list       []byte
		subListArr [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "多值多组索引1",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					yiuAll.YiuStrToByteList("iu"),
					yiuAll.YiuStrToByteList("ll"),
					yiuAll.YiuStrToByteList("!"),
				},
			},
			want: 2,
		},
		{
			name: "多值多组索引2",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					yiuAll.YiuStrToByteList("iu"),
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList("!"),
				},
			},
			want: 6,
		},
		{
			name: "多值多组索引3-索引空",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					yiuAll.YiuStrToByteList("iu"),
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList(""),
					{},
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList("!"),
				},
			},
			want: 6,
		},
		{
			name: "多值多组索引4-索引nil",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					yiuAll.YiuStrToByteList("iu"),
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList(""),
					nil,
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList("!"),
				},
			},
			want: 6,
		},
		{
			name: "多值多组索引5-空索引",
			args: args{
				list: []byte{},
				subListArr: [][]byte{
					yiuAll.YiuStrToByteList("iu"),
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList("!"),
				},
			},
			want: -1,
		},
		{
			name: "多值多组索引6-nil索引",
			args: args{
				list: nil,
				subListArr: [][]byte{
					yiuAll.YiuStrToByteList("iu"),
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList("!"),
				},
			},
			want: -1,
		},
		{
			name: "多值多组索引7-nil索引",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					yiuAll.YiuStrToByteList("fan"),
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList("!"),
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIndexBySList(tt.args.list, tt.args.subListArr...); got != tt.want {
				t.Errorf("GetIndexBySList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIndexAndSubBySList(t *testing.T) {
	type args struct {
		list       []byte
		subListArr [][]byte
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []byte
	}{
		{
			name: "多值多组索引元素测试1",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					yiuAll.YiuStrToByteList("iu"),
					yiuAll.YiuStrToByteList("ll"),
					yiuAll.YiuStrToByteList("!"),
				},
			},
			want:  2,
			want1: yiuAll.YiuStrToByteList("ll"),
		},
		{
			name: "多值多组索引元素测试2",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					yiuAll.YiuStrToByteList("iu"),
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList("!"),
				},
			},
			want:  6,
			want1: yiuAll.YiuStrToByteList("Yiu"),
		},
		{
			name: "多值多组索引元素测试3-索引空",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					yiuAll.YiuStrToByteList("iu"),
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList(""),
					{},
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList("!"),
				},
			},
			want:  6,
			want1: yiuAll.YiuStrToByteList("Yiu"),
		},
		{
			name: "多值多组索引元素测试4-索引nil",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					yiuAll.YiuStrToByteList("iu"),
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList(""),
					nil,
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList("!"),
				},
			},
			want:  6,
			want1: yiuAll.YiuStrToByteList("Yiu"),
		},
		{
			name: "多值多组索引元素测试5-空索引",
			args: args{
				list: []byte{},
				subListArr: [][]byte{
					yiuAll.YiuStrToByteList("iu"),
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList("!"),
				},
			},
			want:  -1,
			want1: nil,
		},
		{
			name: "多值多组索引元素测试6-nil索引",
			args: args{
				list: nil,
				subListArr: [][]byte{
					yiuAll.YiuStrToByteList("iu"),
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList("!"),
				},
			},
			want:  -1,
			want1: nil,
		},
		{
			name: "多值多组索引7-nil索引",
			args: args{
				list: yiuAll.YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					yiuAll.YiuStrToByteList("fan"),
					yiuAll.YiuStrToByteList("Yiu"),
					yiuAll.YiuStrToByteList("!"),
				},
			},
			want:  6,
			want1: yiuAll.YiuStrToByteList("Yiu"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetIndexAndSubBySList(tt.args.list, tt.args.subListArr...)
			if got != tt.want {
				t.Errorf("GetIndexAndSubBySList() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetIndexAndSubBySList() got1 = %v, want %v", got1, tt.want1)
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
			name: "索引元素测试1",
			args: args{
				list:  []byte{1, 2, 3, 4, 5},
				index: 0,
			},
			want: 1,
		},
		{
			name: "索引元素测试2",
			args: args{
				list:  []byte{1, 2, 3, 4, 5},
				index: 2,
			},
			want: 3,
		},
		{
			name: "索引元素测试3-负值索引",
			args: args{
				list:  []byte{1, 2, 3, 4, 5},
				index: -2,
			},
			want: 0,
		},
		{
			name: "索引元素测试4-越界索引",
			args: args{
				list:  []byte{1, 2, 3, 4, 5},
				index: 999,
			},
			want: 0,
		},
		{
			name: "索引元素测试5-空索引",
			args: args{
				list:  []byte{},
				index: 2,
			},
			want: 0,
		},
		{
			name: "索引元素测试6-nil索引",
			args: args{
				list:  nil,
				index: 2,
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

func TestGetMax(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		wantErr bool
	}{
		{
			name: "最大值测试1",
			args: args{
				list: []byte{1, 2, 3, 4, 5, 6},
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "最大值测试2-空最大值",
			args: args{
				list: []byte{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "最大值测试2-nil最大值",
			args: args{
				list: nil,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMax(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMax() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMax() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMin(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		wantErr bool
	}{
		{
			name: "最小值测试1",
			args: args{
				list: []byte{1, 2, 3, 4, 5, 6},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "最小值测试2-空最小值",
			args: args{
				list: []byte{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "最小值测试3-nil最小值",
			args: args{
				list: nil,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMin(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSum(t *testing.T) {
	type args struct {
		list []byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "求和测试1",
			args: args{
				list: []byte{1, 2, 3, 4, 5, 6},
			},
			want: 21,
		},
		{
			name: "求和测试2-空求和",
			args: args{
				list: []byte{},
			},
			want: 0,
		},
		{
			name: "求和测试3-nil求和",
			args: args{
				list: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSum(tt.args.list); got != tt.want {
				t.Errorf("GetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
