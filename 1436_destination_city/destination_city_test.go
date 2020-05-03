package destination_city

import "testing"

func Test_destCity(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{[][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}},
			want: "Sao Paulo",
		},
		{
			name: "test2",
			args: args{[][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}},
			want: "A",
		},
		{
			name: "test3",
			args: args{[][]string{{"A", "Z"}}},
			want: "Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := destCity(tt.args.paths); got != tt.want {
				t.Errorf("destCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
