package validator

import "testing"

func TestContainChinese(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{"中文"}, want: true},
		{name: "test2", args: args{"中123"}, want: true},
		{name: "test3", args: args{"123"}, want: false},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainChinese(tt.args.s); got != tt.want {
				t.Errorf("ContainChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainLetter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{"aa"}, want: true},
		{name: "test2", args: args{"aa123"}, want: true},
		{name: "test3", args: args{"123"}, want: false},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainLetter(tt.args.s); got != tt.want {
				t.Errorf("ContainLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainLower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{"aa"}, want: true},
		{name: "test2", args: args{"aa123"}, want: true},
		{name: "test3", args: args{"123"}, want: false},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainLower(tt.args.s); got != tt.want {
				t.Errorf("ContainLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{"aa"}, want: false},
		{name: "test2", args: args{"aa123"}, want: true},
		{name: "test3", args: args{"123"}, want: true},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainNumber(tt.args.s); got != tt.want {
				t.Errorf("ContainNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.

		{name: "test1", args: args{"AA"}, want: true},
		{name: "test2", args: args{"AA123"}, want: true},
		{name: "test3", args: args{"123"}, want: false},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainUpper(tt.args.s); got != tt.want {
				t.Errorf("ContainUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAllAlph(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.

		{name: "test1", args: args{"aa"}, want: true},
		{name: "test2", args: args{"aa123"}, want: false},
		{name: "test3", args: args{"123"}, want: false},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllAlph(tt.args.s); got != tt.want {
				t.Errorf("IsAllAlph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAllLower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.

		{name: "test1", args: args{"aa"}, want: true},
		{name: "test2", args: args{"aa123"}, want: false},
		{name: "test3", args: args{"123"}, want: false},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllLower(tt.args.s); got != tt.want {
				t.Errorf("IsAllLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAllNum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.

		{name: "test1", args: args{"aa"}, want: false},
		{name: "test2", args: args{"aa123"}, want: false},
		{name: "test3", args: args{"123"}, want: true},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllNum(tt.args.s); got != tt.want {
				t.Errorf("IsAllNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAllUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.

		{name: "test1", args: args{"AA"}, want: true},
		{name: "test2", args: args{"aa123"}, want: false},
		{name: "test3", args: args{"123"}, want: false},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllUpper(tt.args.s); got != tt.want {
				t.Errorf("IsAllUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsChinesePhone(t *testing.T) {
	type args struct {
		phoneStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.

		{name: "test1", args: args{"13536914725"}, want: true},
		{name: "test2", args: args{"123456789"}, want: false},
		{name: "test3", args: args{"23567890456"}, want: false},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChinesePhone(tt.args.phoneStr); got != tt.want {
				t.Errorf("IsChinesePhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	type args struct {
		emailStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{"123@qq.com"}, want: true},
		{name: "test2", args: args{"jeff@gmail.com"}, want: true},
		{name: "test3", args: args{"notexist-@gamil.com"}, want: false},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmail(tt.args.emailStr); got != tt.want {
				t.Errorf("IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFloatStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{"1356789.0456"}, want: true},
		{name: "test2", args: args{"123456789"}, want: true},
		{name: "test3", args: args{".7654"}, want: true},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFloatStr(tt.args.s); got != tt.want {
				t.Errorf("IsFloatStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIntegerStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{"13567890456"}, want: true},
		{name: "test2", args: args{"123456.789"}, want: false},
		{name: "test3", args: args{"23567890tt456"}, want: false},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIntegerStr(tt.args.s); got != tt.want {
				t.Errorf("IsIntegerStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIp(t *testing.T) {
	type args struct {
		ipStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{"192.168.7.23"}, want: true},
		{name: "test2", args: args{"fe80::794b:e81:5c92"}, want: true},
		{name: "test3", args: args{"23567890456"}, want: false},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIp(tt.args.ipStr); got != tt.want {
				t.Errorf("IsIp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIpV4(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{"192.168.7.23"}, want: true},
		{name: "test2", args: args{"fe80::794b:e81:5c92:db78%8"}, want: false},
		{name: "test3", args: args{"23567890456"}, want: false},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIpV4(tt.args.s); got != tt.want {
				t.Errorf("IsIpV4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIpV6(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{"192.168.7.23"}, want: false},
		{name: "test2", args: args{"fe80::794b:e81:5c92"}, want: true},
		{name: "test3", args: args{"23567890456"}, want: false},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIpV6(tt.args.s); got != tt.want {
				t.Errorf("IsIpV6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPort(t *testing.T) {
	type args struct {
		portStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{"8080"}, want: true},
		{name: "test2", args: args{"0"}, want: false},
		{name: "test3", args: args{"65536"}, want: false},
		{name: "test4", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPort(tt.args.portStr); got != tt.want {
				t.Errorf("IsPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUrl(t *testing.T) {
	type args struct {
		urlStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{"https://github.com"}, want: true},
		{name: "test2", args: args{"www.baidu.com"}, want: true},
		{name: "test3", args: args{"github.com"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAbousoluteUrl(tt.args.urlStr); got != tt.want {
				t.Errorf("IsUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegMatcher(t *testing.T) {
	type args struct {
		pattern string
		target  string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{pattern: "[0-9]+", target: "test1"}, want: true},
		{name: "test2", args: args{pattern: "[0-9]+", target: "test"}, want: false},
		{name: "test3", args: args{pattern: "[0-9]+", target: "1111"}, want: true},
		{name: "test4", args: args{pattern: "", target: "test"}, want: true},
		{name: "test4", args: args{pattern: "[0-9]+", target: ""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RegMatcher(tt.args.target, tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegMatcher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RegMatcher() got = %v, want %v", got, tt.want)
			}
		})
	}
}
