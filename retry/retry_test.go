package retry

import "testing"

func TestExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Example()
		})
	}
}
