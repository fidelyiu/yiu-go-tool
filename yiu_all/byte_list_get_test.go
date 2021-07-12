package yiuAll

import (
	"reflect"
	"strings"
	"testing"
)

func TestYiuByteListGetDeduplicate(t *testing.T) {
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
				list: YiuStrToByteList("lllHello"),
			},
			want: YiuStrToByteList("lHeo"),
		},
		{
			name: "去重测试2",
			args: args{
				list: YiuStrToByteList("lllHHHHelloOOO"),
			},
			want: YiuStrToByteList("lHeoO"),
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
				list: YiuStrToByteList("Hello Yiu!"),
			},
			want: YiuStrToByteList("Helo Yiu!"),
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
			if got := YiuByteListGetDeduplicate(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YiuByteListGetDeduplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteListGetDeleteByIndex(t *testing.T) {
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
				list:     YiuStrToByteList("Hello Yiu!"),
				delIndex: 0,
			},
			want: YiuStrToByteList("ello Yiu!"),
		},
		{
			name: "删除测试2-正常删除",
			args: args{
				list:     YiuStrToByteList("Hello Yiu!"),
				delIndex: len("Hello Yiu!") - 1,
			},
			want: YiuStrToByteList("Hello Yiu"),
		},
		{
			name: "删除测试3-正常删除",
			args: args{
				list:     YiuStrToByteList("Hello Yiu!"),
				delIndex: strings.Index("Hello Yiu!", " "),
			},
			want: YiuStrToByteList("HelloYiu!"),
		},
		{
			name: "删除测试4-负值删除",
			args: args{
				list:     YiuStrToByteList("Hello Yiu!"),
				delIndex: -1,
			},
			want: YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "删除测试5-越界删除",
			args: args{
				list:     YiuStrToByteList("Hello Yiu!"),
				delIndex: 999,
			},
			want: YiuStrToByteList("Hello Yiu!"),
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
			if got := YiuByteListGetDeleteByIndex(tt.args.list, tt.args.delIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YiuByteListGetDeleteByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteListGetDeleteByRangeIndex(t *testing.T) {
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
				list:       YiuStrToByteList("Hello Yiu!"),
				startIndex: 0,
				endIndex:   0,
			},
			want: YiuStrToByteList("ello Yiu!"),
		},
		{
			name: "范围删除2-正常删除",
			args: args{
				list:       YiuStrToByteList("Hello Yiu!"),
				startIndex: 0,
				endIndex:   1,
			},
			want: YiuStrToByteList("llo Yiu!"),
		},
		{
			name: "范围删除3-正常删除",
			args: args{
				list:       YiuStrToByteList("Hello Yiu!"),
				startIndex: strings.Index("Hello Yiu!", " "),
				endIndex:   strings.Index("Hello Yiu!", "u"),
			},
			want: YiuStrToByteList("Hello!"),
		},
		{
			name: "范围删除4-负值删除",
			args: args{
				list:       YiuStrToByteList("Hello Yiu!"),
				startIndex: -1,
				endIndex:   strings.Index("Hello Yiu!", "u"),
			},
			want: YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "范围删除5-负值删除",
			args: args{
				list:       YiuStrToByteList("Hello Yiu!"),
				startIndex: -1,
				endIndex:   -13,
			},
			want: YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "范围删除6-负值删除",
			args: args{
				list:       YiuStrToByteList("Hello Yiu!"),
				startIndex: 2,
				endIndex:   -13,
			},
			want: YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "范围删除7-越界删除",
			args: args{
				list:       YiuStrToByteList("Hello Yiu!"),
				startIndex: 999,
				endIndex:   3,
			},
			want: YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "范围删除8-越界删除",
			args: args{
				list:       YiuStrToByteList("Hello Yiu!"),
				startIndex: 3,
				endIndex:   999,
			},
			want: YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "范围删除9-越界删除",
			args: args{
				list:       YiuStrToByteList("Hello Yiu!"),
				startIndex: 998,
				endIndex:   999,
			},
			want: YiuStrToByteList("Hello Yiu!"),
		},
		{
			name: "范围删除10-开始索引大于结束索引删除",
			args: args{
				list:       YiuStrToByteList("Hello Yiu!"),
				startIndex: 3,
				endIndex:   2,
			},
			want: YiuStrToByteList("Hello Yiu!"),
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
			if got := YiuByteListGetDeleteByRangeIndex(tt.args.list, tt.args.startIndex, tt.args.endIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YiuByteListGetDeleteByRangeIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteListGetFilter(t *testing.T) {
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
				list: YiuStrToByteList("Hello Yiu!"),
				keep: func(x byte) bool {
					return x == 'Y' || x == 'i' || x == 'u'
				},
			},
			want: YiuStrToByteList("Yiu"),
		},
		{
			name: "过滤测试2",
			args: args{
				list: YiuStrToByteList("Hello Yiu!"),
				keep: func(x byte) bool {
					return false
				},
			},
			want: YiuStrToByteList(""),
		},
		{
			name: "过滤测试3",
			args: args{
				list: YiuStrToByteList("Hello Yiu!"),
				keep: func(x byte) bool {
					return true
				},
			},
			want: YiuStrToByteList("Hello Yiu!"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YiuByteListGetFilter(tt.args.list, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YiuByteListGetFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteListGetPop(t *testing.T) {
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
				list: YiuStrToByteList("Hello Yiu!"),
			},
			want:    '!',
			want1:   YiuStrToByteList("Hello Yiu"),
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
			got, got1, err := YiuByteListGetPop(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("YiuByteListGetPop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("YiuByteListGetPop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("YiuByteListGetPop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestYiuByteListGetReverse(t *testing.T) {
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
				list: YiuStrToByteList("Yiu!"),
			},
			want: YiuStrToByteList("!uiY"),
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
			if got := YiuByteListGetReverse(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YiuByteListGetReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteListGetShuffle(t *testing.T) {
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
			if got := YiuByteListGetShuffle(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YiuByteListGetShuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteListGetMap(t *testing.T) {
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
			if got := YiuByteListGetMap(tt.args.list, tt.args.opFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YiuByteListGetMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteListGetCopy(t *testing.T) {
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
				list: YiuStrToByteList("Hello Yiu!"),
			},
			want: YiuStrToByteList("Hello Yiu!"),
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
			if got := YiuByteListGetCopy(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YiuByteListGetCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteListGetIndexByEl(t *testing.T) {
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
				list: YiuStrToByteList("Hello Yiu!"),
				el:   'l',
			},
			want: strings.Index("Hello Yiu!", "l"),
		},
		{
			name: "索引元素测试2-无值索引",
			args: args{
				list: YiuStrToByteList("Hello Yiu!"),
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
			if got := YiuByteListGetIndexByEl(tt.args.list, tt.args.el); got != tt.want {
				t.Errorf("YiuByteListGetIndexByEl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteListGetIndexByList(t *testing.T) {
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
				list:    YiuStrToByteList("Hello Yiu!"),
				subList: YiuStrToByteList("ell"),
			},
			want: 1,
		},
		{
			name: "多值索引2",
			args: args{
				list:    YiuStrToByteList("Hello Yiu!"),
				subList: YiuStrToByteList("Yiu"),
			},
			want: 6,
		},
		{
			name: "多值索引3-空索引",
			args: args{
				list:    []byte{},
				subList: YiuStrToByteList("Yiu"),
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
				subList: YiuStrToByteList("Yiu"),
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
				list:    YiuStrToByteList("Yiu"),
				subList: []byte{},
			},
			want: -1,
		},
		{
			name: "多值索引8-索引nil",
			args: args{
				list:    YiuStrToByteList("Yiu"),
				subList: nil,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YiuByteListGetIndexByList(tt.args.list, tt.args.subList); got != tt.want {
				t.Errorf("YiuByteListGetIndexByList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteListGetIndexBySList(t *testing.T) {
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
				list: YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					YiuStrToByteList("iu"),
					YiuStrToByteList("ll"),
					YiuStrToByteList("!"),
				},
			},
			want: 2,
		},
		{
			name: "多值多组索引2",
			args: args{
				list: YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					YiuStrToByteList("iu"),
					YiuStrToByteList("Yiu"),
					YiuStrToByteList("Yiu"),
					YiuStrToByteList("!"),
				},
			},
			want: 6,
		},
		{
			name: "多值多组索引3-索引空",
			args: args{
				list: YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					YiuStrToByteList("iu"),
					YiuStrToByteList("Yiu"),
					YiuStrToByteList(""),
					{},
					YiuStrToByteList("Yiu"),
					YiuStrToByteList("!"),
				},
			},
			want: 6,
		},
		{
			name: "多值多组索引4-索引nil",
			args: args{
				list: YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					YiuStrToByteList("iu"),
					YiuStrToByteList("Yiu"),
					YiuStrToByteList(""),
					nil,
					YiuStrToByteList("Yiu"),
					YiuStrToByteList("!"),
				},
			},
			want: 6,
		},
		{
			name: "多值多组索引5-空索引",
			args: args{
				list: []byte{},
				subListArr: [][]byte{
					YiuStrToByteList("iu"),
					YiuStrToByteList("Yiu"),
					YiuStrToByteList("!"),
				},
			},
			want: -1,
		},
		{
			name: "多值多组索引6-nil索引",
			args: args{
				list: nil,
				subListArr: [][]byte{
					YiuStrToByteList("iu"),
					YiuStrToByteList("Yiu"),
					YiuStrToByteList("!"),
				},
			},
			want: -1,
		},
		{
			name: "多值多组索引7-nil索引",
			args: args{
				list: YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					YiuStrToByteList("fan"),
					YiuStrToByteList("Yiu"),
					YiuStrToByteList("!"),
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YiuByteListGetIndexBySList(tt.args.list, tt.args.subListArr...); got != tt.want {
				t.Errorf("YiuByteListGetIndexBySList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteListGetIndexAndSubBySList(t *testing.T) {
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
				list: YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					YiuStrToByteList("iu"),
					YiuStrToByteList("ll"),
					YiuStrToByteList("!"),
				},
			},
			want:  2,
			want1: YiuStrToByteList("ll"),
		},
		{
			name: "多值多组索引元素测试2",
			args: args{
				list: YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					YiuStrToByteList("iu"),
					YiuStrToByteList("Yiu"),
					YiuStrToByteList("Yiu"),
					YiuStrToByteList("!"),
				},
			},
			want:  6,
			want1: YiuStrToByteList("Yiu"),
		},
		{
			name: "多值多组索引元素测试3-索引空",
			args: args{
				list: YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					YiuStrToByteList("iu"),
					YiuStrToByteList("Yiu"),
					YiuStrToByteList(""),
					{},
					YiuStrToByteList("Yiu"),
					YiuStrToByteList("!"),
				},
			},
			want:  6,
			want1: YiuStrToByteList("Yiu"),
		},
		{
			name: "多值多组索引元素测试4-索引nil",
			args: args{
				list: YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					YiuStrToByteList("iu"),
					YiuStrToByteList("Yiu"),
					YiuStrToByteList(""),
					nil,
					YiuStrToByteList("Yiu"),
					YiuStrToByteList("!"),
				},
			},
			want:  6,
			want1: YiuStrToByteList("Yiu"),
		},
		{
			name: "多值多组索引元素测试5-空索引",
			args: args{
				list: []byte{},
				subListArr: [][]byte{
					YiuStrToByteList("iu"),
					YiuStrToByteList("Yiu"),
					YiuStrToByteList("!"),
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
					YiuStrToByteList("iu"),
					YiuStrToByteList("Yiu"),
					YiuStrToByteList("!"),
				},
			},
			want:  -1,
			want1: nil,
		},
		{
			name: "多值多组索引7-nil索引",
			args: args{
				list: YiuStrToByteList("Hello Yiu!"),
				subListArr: [][]byte{
					YiuStrToByteList("fan"),
					YiuStrToByteList("Yiu"),
					YiuStrToByteList("!"),
				},
			},
			want:  6,
			want1: YiuStrToByteList("Yiu"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := YiuByteListGetIndexAndSubBySList(tt.args.list, tt.args.subListArr...)
			if got != tt.want {
				t.Errorf("YiuByteListGetIndexAndSubBySList() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("YiuByteListGetIndexAndSubBySList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestYiuByteListGetElByIndex(t *testing.T) {
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
			if got := YiuByteListGetElByIndex(tt.args.list, tt.args.index); got != tt.want {
				t.Errorf("YiuByteListGetElByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteListGetMax(t *testing.T) {
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
			got, err := YiuByteListGetMax(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("YiuByteListGetMax() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("YiuByteListGetMax() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteListGetMin(t *testing.T) {
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
			got, err := YiuByteListGetMin(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("YiuByteListGetMin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("YiuByteListGetMin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYiuByteListGetSum(t *testing.T) {
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
			if got := YiuByteListGetSum(tt.args.list); got != tt.want {
				t.Errorf("YiuByteListGetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
