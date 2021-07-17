package yiuAll

import "testing"

func TestYiuOsDoOpenFileManagerByParent(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常测试",
			args: args{
				p: "E:\\学习文件\\前端\\Prismjs",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := YiuOsDoOpenFileManagerByParent(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("YiuOsDoOpenFileManagerByParent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
