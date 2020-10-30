package add_strings

import "testing"

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{"1", "1"},
			want: "2",
		},
		{
			name: "test2",
			args: args{"100000000000000000000000000", "1"},
			want: "100000000000000000000000001",
		},
		{
			name: "test3",
			args: args{"99999999999", "1"},
			want: "100000000000",
		},
		{
			name: "test4",
			args: args{"1", "99999999999"},
			want: "100000000000",
		},
		{
			name: "test5",
			args: args{"123", "123456"},
			want: "123579",
		},
		{
			name: "test6",
			args: args{"666", "65443"},
			want: "66109",
		},
		{
			name: "test7",
			args: args{"5", "5"},
			want: "10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
