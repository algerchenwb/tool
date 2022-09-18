package sort

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeapSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
