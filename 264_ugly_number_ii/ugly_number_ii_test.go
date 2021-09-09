package ugly_number_ii

import "testing"

func Test_nthUglyNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	name: "test1",
		//	args: args{10},
		//	want: 12,
		//},
		//{
		//	name: "test2",
		//	args: args{1},
		//	want: 1,
		//},
		//{
		//	name: "test3",
		//	args: args{11},
		//	want: 15,
		//},
		//{
		//	name: "test4",
		//	args: args{402},
		//	want: 314928,
		//},
		//{
		//	name: "test4",
		//	args: args{1690},
		//	want: 314928,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.args.n); got != tt.want {
				t.Errorf("nthUglyNumber(%d) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
