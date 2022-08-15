package fileutil

import "testing"

func TestAccess(t *testing.T) {
	type args struct {
		fromXlsPath string
		toXlsxPath  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{fromXlsPath: "from.xls", toXlsxPath: "to.xlsx"}},
		{name: "test2", args: args{fromXlsPath: "中文.xls", toXlsxPath: "to中文.xlsx"}},
		{name: "test3", args: args{fromXlsPath: "inexistent.xls", toXlsxPath: "inexistent.xlsx"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Example(tt.args.fromXlsPath, tt.args.toXlsxPath)
		})
	}
}
