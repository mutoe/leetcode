package employee_importance

import "testing"

func Test_getImportance(t *testing.T) {
	type args struct {
		employees []*Employee
		id        int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name: "test1",
			args: args{
				employees: []*Employee{
					{1, 5, []int{2, 3}},
					{2, 3, nil},
					{3, 3, nil},
				},
				id: 1,
			},
			wantSum: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := getImportance(tt.args.employees, tt.args.id); gotSum != tt.wantSum {
				t.Errorf("getImportance() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
