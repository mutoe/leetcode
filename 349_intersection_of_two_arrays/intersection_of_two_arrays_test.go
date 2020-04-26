package intersection_of_two_arrays

import (
	"reflect"
	"sort"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "test1",
			args: args{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			wantResult: []int{2},
		},
		{
			name: "test2",
			args: args{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			wantResult: []int{9, 4},
		},
		{
			name: "test3",
			args: args{
				nums1: []int{1},
				nums2: []int{},
			},
			wantResult: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := intersection(tt.args.nums1, tt.args.nums2)
			sort.Ints(gotResult)
			sort.Ints(tt.wantResult)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("intersection() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
