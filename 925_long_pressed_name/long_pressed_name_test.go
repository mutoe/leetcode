package long_pressed_name

import "testing"

func Test_isLongPressedName(t *testing.T) {
	type args struct {
		name  string
		typed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{"alex", "aaleex"},
			want: true,
		},
		{
			name: "test2",
			args: args{"alex", "aaleelx"},
			want: false,
		},
		{
			name: "test3",
			args: args{"alex", "alexx"},
			want: true,
		},
		{
			name: "test4",
			args: args{"alex", "alexxr"},
			want: false,
		},
		{
			name: "test5",
			args: args{"saeed", "ssaaedd"},
			want: false,
		},
		{
			name: "test6",
			args: args{"leelee", "lleeelee"},
			want: true,
		},
		{
			name: "test7",
			args: args{"laiden", "laiden"},
			want: true,
		},
		{
			name: "test8",
			args: args{"kikcxmvzi", "kiikcxxmmvvzz"},
			want: false,
		},
		{
			name: "test9",
			args: args{"pyplrz", "ppyypllr"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLongPressedName(tt.args.name, tt.args.typed); got != tt.want {
				t.Errorf("isLongPressedName(%s, %s) = %v, want %v", tt.args.name, tt.args.typed, got, tt.want)
			}
		})
	}
}
