package average_salary_excluding_the_minimum_and_maximum_salary

import "testing"

func Test_average(t *testing.T) {
	type args struct {
		salary []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{[]int{4000, 3000, 1000, 2000}},
			want: 2500,
		},
		{
			name: "test2",
			args: args{[]int{1000, 2000, 3000}},
			want: 2000,
		},
		{
			name: "test3",
			args: args{[]int{6000, 5000, 4000, 3000, 2000, 1000}},
			want: 3500,
		},
		{
			name: "test4",
			args: args{[]int{8000, 9000, 2000, 3000, 6000, 1000}},
			want: 4750,
		},
		{
			name: "test5",
			args: args{[]int{48000, 59000, 99000, 13000, 78000, 45000, 31000, 17000, 39000, 37000, 93000, 77000, 33000, 28000, 4000, 54000, 67000, 6000, 1000, 11000}},
			want: 41111.11111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.args.salary); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
