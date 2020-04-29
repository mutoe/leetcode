package excel_sheet_column_number

import "testing"

func Test_titleToNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{"A"},
			want: 1,
		},
		{
			name: "test",
			args: args{"AB"},
			want: 28,
		},
		{
			name: "test",
			args: args{"ZY"},
			want: 701,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.s); got != tt.want {
				t.Errorf("titleToNumber(%s) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
