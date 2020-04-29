package excel_sheet_column_title

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{1},
			want: "A",
		},
		{
			name: "test",
			args: args{28},
			want: "AB",
		},
		{
			name: "test",
			args: args{26},
			want: "Z",
		},
		{
			name: "test",
			args: args{701},
			want: "ZY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.n); got != tt.want {
				t.Errorf("convertToTitle(%d) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
