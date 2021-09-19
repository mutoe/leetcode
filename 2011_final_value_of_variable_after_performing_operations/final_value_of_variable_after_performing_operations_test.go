package final_value_of_variable_after_performing_operations

import "testing"

func Test_finalValueAfterOperations(t *testing.T) {
	type args struct {
		operations []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]string{"--X", "X++", "X++"}},
			want: 1,
		},
		{
			name: "test2",
			args: args{[]string{"++X", "++X", "X++"}},
			want: 3,
		},
		{
			name: "test3",
			args: args{[]string{"X++", "++X", "--X", "X--"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalValueAfterOperations(tt.args.operations); got != tt.want {
				t.Errorf("finalValueAfterOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
