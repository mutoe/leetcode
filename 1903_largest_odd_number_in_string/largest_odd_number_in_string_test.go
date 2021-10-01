package largest_odd_number_in_string

import "testing"

func Test_largestOddNumber(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{"52"},
			want: "5",
		},
		{
			name: "test2",
			args: args{"4206"},
			want: "",
		},
		{
			name: "test3",
			args: args{"35427"},
			want: "35427",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestOddNumber(tt.args.num); got != tt.want {
				t.Errorf("largestOddNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
