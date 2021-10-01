package minimum_time_to_type_word_using_special_typewriter

import "testing"

func Test_minTimeToType(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{"abc"},
			want: 5,
		},
		{
			name: "test2",
			args: args{"bza"},
			want: 7,
		},
		{
			name: "test3",
			args: args{"zjpc"},
			want: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeToType(tt.args.word); got != tt.want {
				t.Errorf("minTimeToType(%v) = %v, want %v", tt.args.word, got, tt.want)
			}
		})
	}
}
