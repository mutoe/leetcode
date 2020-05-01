package backspace_string_compare

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{"ab#c", "ad#c"},
			want: true,
		},
		{
			name: "test2",
			args: args{"ab##", "c#d#"},
			want: true,
		},
		{
			name: "test3",
			args: args{"a##c", "#a#c"},
			want: true,
		},
		{
			name: "test4",
			args: args{"a#c", "b"},
			want: false,
		},
		{
			name: "test5",
			args: args{"bxj##tw", "bxj###tw"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompare(%s, %s) = %v, want %v", tt.args.S, tt.args.T, got, tt.want)
			}
		})
	}
}
