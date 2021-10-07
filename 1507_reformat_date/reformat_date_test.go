package reformat_date

import "testing"

func Test_reformatDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{"20th Oct 2052"},
			want: "2052-10-20",
		},
		{
			name: "test2",
			args: args{"6th Jun 1933"},
			want: "1933-06-06",
		},
		{
			name: "test3",
			args: args{"26th May 1960"},
			want: "1960-05-26",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reformatDate(tt.args.date); got != tt.want {
				t.Errorf("reformatDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
